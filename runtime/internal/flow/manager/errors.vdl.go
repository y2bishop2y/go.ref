// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package manager

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var (
	ErrLargerThan3ByteUInt       = verror.Register("v.io/x/ref/runtime/internal/flow/manager.LargerThan3ByteUInt", verror.NoRetry, "{1:}{2:} integer too large to represent in 3 bytes")
	ErrUnknownProtocol           = verror.Register("v.io/x/ref/runtime/internal/flow/manager.UnknownProtocol", verror.NoRetry, "{1:}{2:} unknown protocol{:3}")
	ErrManagerClosed             = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ManagerClosed", verror.NoRetry, "{1:}{2:} manager is already closed")
	ErrAcceptFailed              = verror.Register("v.io/x/ref/runtime/internal/flow/manager.AcceptFailed", verror.NoRetry, "{1:}{2:} accept failed{:3}")
	ErrCacheClosed               = verror.Register("v.io/x/ref/runtime/internal/flow/manager.CacheClosed", verror.NoRetry, "{1:}{2:} cache is closed")
	ErrConnKilledToFreeResources = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ConnKilledToFreeResources", verror.NoRetry, "{1:}{2:} Connection killed to free resources.")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrLargerThan3ByteUInt.ID), "{1:}{2:} integer too large to represent in 3 bytes")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownProtocol.ID), "{1:}{2:} unknown protocol{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrManagerClosed.ID), "{1:}{2:} manager is already closed")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAcceptFailed.ID), "{1:}{2:} accept failed{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCacheClosed.ID), "{1:}{2:} cache is closed")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrConnKilledToFreeResources.ID), "{1:}{2:} Connection killed to free resources.")
}

// NewErrLargerThan3ByteUInt returns an error with the ErrLargerThan3ByteUInt ID.
func NewErrLargerThan3ByteUInt(ctx *context.T) error {
	return verror.New(ErrLargerThan3ByteUInt, ctx)
}

// NewErrUnknownProtocol returns an error with the ErrUnknownProtocol ID.
func NewErrUnknownProtocol(ctx *context.T, protocol string) error {
	return verror.New(ErrUnknownProtocol, ctx, protocol)
}

// NewErrManagerClosed returns an error with the ErrManagerClosed ID.
func NewErrManagerClosed(ctx *context.T) error {
	return verror.New(ErrManagerClosed, ctx)
}

// NewErrAcceptFailed returns an error with the ErrAcceptFailed ID.
func NewErrAcceptFailed(ctx *context.T, err error) error {
	return verror.New(ErrAcceptFailed, ctx, err)
}

// NewErrCacheClosed returns an error with the ErrCacheClosed ID.
func NewErrCacheClosed(ctx *context.T) error {
	return verror.New(ErrCacheClosed, ctx)
}

// NewErrConnKilledToFreeResources returns an error with the ErrConnKilledToFreeResources ID.
func NewErrConnKilledToFreeResources(ctx *context.T) error {
	return verror.New(ErrConnKilledToFreeResources, ctx)
}
