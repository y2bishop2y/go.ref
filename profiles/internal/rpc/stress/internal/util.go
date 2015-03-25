// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"crypto/md5"
	"encoding/binary"

	"v.io/x/ref/profiles/internal/rpc/stress"
)

// doSum returns the MD5 checksum of the arg.
func doSum(arg stress.Arg) ([]byte, error) {
	h := md5.New()
	if arg.ABool {
		if err := binary.Write(h, binary.LittleEndian, arg.AInt64); err != nil {
			return nil, err
		}
	}
	if _, err := h.Write(arg.AListOfBytes); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
