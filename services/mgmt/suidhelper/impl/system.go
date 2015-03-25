// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin

package impl

import (
	"log"
	"os"
	"path/filepath"
	"syscall"

	"v.io/v23/verror"
)

var (
	errChownFailed        = verror.Register(pkgPath+".errChownFailed", verror.NoRetry, "{1:}{2:} os.Chown({3}, {4}, {5}) failed{:_}")
	errGetwdFailed        = verror.Register(pkgPath+".errGetwdFailed", verror.NoRetry, "{1:}{2:} os.Getwd failed{:_}")
	errStartProcessFailed = verror.Register(pkgPath+".errStartProcessFailed", verror.NoRetry, "{1:}{2:} syscall.StartProcess({3}) failed{:_}")
	errRemoveAllFailed    = verror.Register(pkgPath+".errRemoveAllFailed", verror.NoRetry, "{1:}{2:} os.RemoveAll({3}) failed{:_}")
)

// Chown is only availabe on UNIX platforms so this file has a build
// restriction.
func (hw *WorkParameters) Chown() error {
	chown := func(path string, _ os.FileInfo, inerr error) error {
		if inerr != nil {
			return inerr
		}
		if hw.dryrun {
			log.Printf("[dryrun] os.Chown(%s, %d, %d)", path, hw.uid, hw.gid)
			return nil
		}
		return os.Chown(path, hw.uid, hw.gid)
	}

	for _, p := range []string{hw.workspace, hw.logDir} {
		// TODO(rjkroege): Ensure that the device manager can read log entries.
		if err := filepath.Walk(p, chown); err != nil {
			return verror.New(errChownFailed, nil, p, hw.uid, hw.gid, err)
		}
	}
	return nil
}

func (hw *WorkParameters) Exec() error {
	attr := new(syscall.ProcAttr)

	if dir, err := os.Getwd(); err != nil {
		log.Printf("error Getwd(): %v\n", err)
		return verror.New(errGetwdFailed, nil, err)
		attr.Dir = dir
	}
	attr.Env = hw.envv
	attr.Files = []uintptr{
		uintptr(syscall.Stdin),
		uintptr(syscall.Stdout),
		uintptr(syscall.Stderr),
	}

	attr.Sys = new(syscall.SysProcAttr)
	attr.Sys.Setsid = true
	if hw.dryrun {
		log.Printf("[dryrun] syscall.Setgid(%d)", hw.gid)
		log.Printf("[dryrun] syscall.Setuid(%d)", hw.uid)
	} else {
		attr.Sys.Credential = new(syscall.Credential)
		attr.Sys.Credential.Gid = uint32(hw.gid)
		attr.Sys.Credential.Uid = uint32(hw.uid)
	}

	_, _, err := syscall.StartProcess(hw.argv0, hw.argv, attr)
	if err != nil {
		if !hw.dryrun {
			log.Printf("StartProcess failed: attr: %#v, attr.Sys: %#v, attr.Sys.Cred: %#v error: %v\n", attr, attr.Sys, attr.Sys.Credential, err)
		} else {
			log.Printf("StartProcess failed: %v", err)
		}
		return verror.New(errStartProcessFailed, nil, hw.argv0, err)
	}
	// TODO(rjkroege): Return the pid to the node manager.
	os.Exit(0)
	return nil // Not reached.
}

func (hw *WorkParameters) Remove() error {
	for _, p := range hw.argv {
		if err := os.RemoveAll(p); err != nil {
			return verror.New(errRemoveAllFailed, nil, p, err)
		}
	}
	return nil
}
