// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: otherfile.vdl

package nativetest

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"time"
	"v.io/v23/vdl/testdata/nativetest"
)

type ignoreme string

func (ignoreme) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.ignoreme"`
}) {
}

func init() {
	vdl.Register((*ignoreme)(nil))
}
