// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package nosql

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security/access"
)

// databaseData represents the persistent state of a Database.
type databaseData struct {
	Name    string
	Version uint64 // covers the fields below
	Perms   access.Permissions
}

func (databaseData) __VDLReflect(struct {
	Name string "v.io/syncbase/x/ref/services/syncbase/server/nosql.databaseData"
}) {
}

// tableData represents the persistent state of a Table.
// TODO(sadovsky): Decide whether to track "empty-prefix" perms here.
type tableData struct {
	Name  string
	Perms access.Permissions
}

func (tableData) __VDLReflect(struct {
	Name string "v.io/syncbase/x/ref/services/syncbase/server/nosql.tableData"
}) {
}

func init() {
	vdl.Register((*databaseData)(nil))
	vdl.Register((*tableData)(nil))
}
