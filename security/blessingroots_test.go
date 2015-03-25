// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"v.io/v23/security"
	"v.io/v23/verror"
)

type rootsTester [3]security.PublicKey

func newRootsTester() *rootsTester {
	var tester rootsTester
	var err error
	for idx := range tester {
		if tester[idx], _, err = NewPrincipalKey(); err != nil {
			panic(err)
		}
	}
	return &tester
}

func (t *rootsTester) add(br security.BlessingRoots) error {
	testdata := []struct {
		root    security.PublicKey
		pattern security.BlessingPattern
	}{
		{t[0], "vanadium"},
		{t[1], "google/foo"},
		{t[0], "google/$"},
	}
	for _, d := range testdata {
		if err := br.Add(d.root, d.pattern); err != nil {
			return fmt.Errorf("Add(%v, %q) failed: %s", d.root, d.pattern, err)
		}
	}
	return nil
}

func (t *rootsTester) testRecognized(br security.BlessingRoots) error {
	testdata := []struct {
		root          security.PublicKey
		recognized    []string
		notRecognized []string
	}{
		{
			root:          t[0],
			recognized:    []string{"vanadium", "vanadium/foo", "vanadium/foo/bar", "google"},
			notRecognized: []string{"google/foo", "foo", "foo/bar"},
		},
		{
			root:          t[1],
			recognized:    []string{"google/foo", "google/foo/bar"},
			notRecognized: []string{"google", "google/bar", "vanadium", "vanadium/foo", "foo", "foo/bar"},
		},
		{
			root:          t[2],
			recognized:    []string{},
			notRecognized: []string{"vanadium", "vanadium/foo", "vanadium/bar", "google", "google/foo", "google/bar", "foo", "foo/bar"},
		},
	}
	for _, d := range testdata {
		for _, b := range d.recognized {
			if err := br.Recognized(d.root, b); err != nil {
				return fmt.Errorf("Recognized(%v, %q): got: %v, want nil", d.root, b, err)
			}
		}
		for _, b := range d.notRecognized {
			if err, want := br.Recognized(d.root, b), security.ErrUnrecognizedRoot.ID; !verror.Is(err, want) {
				return fmt.Errorf("Recognized(%v, %q): Got %v(errorid=%v), want errorid=%v", d.root, b, err, verror.ErrorID(err), want)
			}
		}
	}
	return nil
}

func TestBlessingRoots(t *testing.T) {
	p, err := NewPrincipal()
	if err != nil {
		t.Fatal(err)
	}
	tester := newRootsTester()
	if err := tester.add(p.Roots()); err != nil {
		t.Fatal(err)
	}
	if err := tester.testRecognized(p.Roots()); err != nil {
		t.Fatal(err)
	}
}

func TestBlessingRootsPersistence(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestBlessingRootsPersistence")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	tester := newRootsTester()
	p, err := CreatePersistentPrincipal(dir, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := tester.add(p.Roots()); err != nil {
		t.Error(err)
	}
	if err := tester.testRecognized(p.Roots()); err != nil {
		t.Error(err)
	}
	// Recreate the principal (and thus BlessingRoots)
	p2, err := LoadPersistentPrincipal(dir, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := tester.testRecognized(p2.Roots()); err != nil {
		t.Error(err)
	}
}
