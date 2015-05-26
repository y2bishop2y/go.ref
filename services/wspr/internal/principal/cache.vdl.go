// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: cache.vdl

package principal

import (
	// VDL system imports
	"v.io/v23/vdl"
)

// Identifier of a blessings cache entry.
type BlessingsId uint32

func (BlessingsId) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsId"`
}) {
}

type BlessingsCacheAddMessage struct {
	CacheId   BlessingsId
	Blessings JsBlessings
}

func (BlessingsCacheAddMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage"`
}) {
}

// Message from Blessings Cache GC to delete a cache entry in Javascript.
type BlessingsCacheDeleteMessage struct {
	CacheId BlessingsId
	// Number of references expected. Javascript should wait until this number
	// has been received before deleting the entry because up until that point
	// messages with further references are expected.
	DeleteAfter uint32
}

func (BlessingsCacheDeleteMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheDeleteMessage"`
}) {
}

type (
	// BlessingsCacheMessage represents any single field of the BlessingsCacheMessage union type.
	BlessingsCacheMessage interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlessingsCacheMessage union type.
		__VDLReflect(__BlessingsCacheMessageReflect)
	}
	// BlessingsCacheMessageAdd represents field Add of the BlessingsCacheMessage union type.
	BlessingsCacheMessageAdd struct{ Value BlessingsCacheAddMessage }
	// BlessingsCacheMessageDelete represents field Delete of the BlessingsCacheMessage union type.
	BlessingsCacheMessageDelete struct{ Value BlessingsCacheDeleteMessage }
	// __BlessingsCacheMessageReflect describes the BlessingsCacheMessage union type.
	__BlessingsCacheMessageReflect struct {
		Name  string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheMessage"`
		Type  BlessingsCacheMessage
		Union struct {
			Add    BlessingsCacheMessageAdd
			Delete BlessingsCacheMessageDelete
		}
	}
)

func (x BlessingsCacheMessageAdd) Index() int                                  { return 0 }
func (x BlessingsCacheMessageAdd) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageAdd) Name() string                                { return "Add" }
func (x BlessingsCacheMessageAdd) __VDLReflect(__BlessingsCacheMessageReflect) {}

func (x BlessingsCacheMessageDelete) Index() int                                  { return 1 }
func (x BlessingsCacheMessageDelete) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageDelete) Name() string                                { return "Delete" }
func (x BlessingsCacheMessageDelete) __VDLReflect(__BlessingsCacheMessageReflect) {}

func init() {
	vdl.Register((*BlessingsId)(nil))
	vdl.Register((*BlessingsCacheAddMessage)(nil))
	vdl.Register((*BlessingsCacheDeleteMessage)(nil))
	vdl.Register((*BlessingsCacheMessage)(nil))
}
