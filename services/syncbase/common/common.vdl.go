// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: common

package common

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrPermsEmpty          = verror.Register("v.io/x/ref/services/syncbase/common.PermsEmpty", verror.NoRetry, "{1:}{2:} permissions cannot be empty")
	ErrPermsNoAdmin        = verror.Register("v.io/x/ref/services/syncbase/common.PermsNoAdmin", verror.NoRetry, "{1:}{2:} permissions must include at least one admin")
	ErrPermsDisallowedTags = verror.Register("v.io/x/ref/services/syncbase/common.PermsDisallowedTags", verror.NoRetry, "{1:}{2:} permissions tags {3} are not allowed; only {4} are allowed")
)

// NewErrPermsEmpty returns an error with the ErrPermsEmpty ID.
func NewErrPermsEmpty(ctx *context.T) error {
	return verror.New(ErrPermsEmpty, ctx)
}

// NewErrPermsNoAdmin returns an error with the ErrPermsNoAdmin ID.
func NewErrPermsNoAdmin(ctx *context.T) error {
	return verror.New(ErrPermsNoAdmin, ctx)
}

// NewErrPermsDisallowedTags returns an error with the ErrPermsDisallowedTags ID.
func NewErrPermsDisallowedTags(ctx *context.T, disallowed []string, allowed []string) error {
	return verror.New(ErrPermsDisallowedTags, ctx, disallowed, allowed)
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPermsEmpty.ID), "{1:}{2:} permissions cannot be empty")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPermsNoAdmin.ID), "{1:}{2:} permissions must include at least one admin")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPermsDisallowedTags.ID), "{1:}{2:} permissions tags {3} are not allowed; only {4} are allowed")

	return struct{}{}
}
