// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: channel.vdl

package channel

import (
	// VDL system imports
	"v.io/v23/vdl"
)

type Request struct {
	Type string
	Seq  uint32
	Body *vdl.Value
}

func (Request) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wspr/internal/channel.Request"
}) {
}

type Response struct {
	ReqSeq uint32
	Err    string // TODO(bprosnitz) change this back to error when it is possible to do so. (issue 368)
	Body   *vdl.Value
}

func (Response) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wspr/internal/channel.Response"
}) {
}

type (
	// Message represents any single field of the Message union type.
	Message interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Message union type.
		__VDLReflect(__MessageReflect)
	}
	// MessageRequest represents field Request of the Message union type.
	MessageRequest struct{ Value Request }
	// MessageResponse represents field Response of the Message union type.
	MessageResponse struct{ Value Response }
	// __MessageReflect describes the Message union type.
	__MessageReflect struct {
		Name  string "v.io/x/ref/services/wspr/internal/channel.Message"
		Type  Message
		Union struct {
			Request  MessageRequest
			Response MessageResponse
		}
	}
)

func (x MessageRequest) Index() int                    { return 0 }
func (x MessageRequest) Interface() interface{}        { return x.Value }
func (x MessageRequest) Name() string                  { return "Request" }
func (x MessageRequest) __VDLReflect(__MessageReflect) {}

func (x MessageResponse) Index() int                    { return 1 }
func (x MessageResponse) Interface() interface{}        { return x.Value }
func (x MessageResponse) Name() string                  { return "Response" }
func (x MessageResponse) __VDLReflect(__MessageReflect) {}

func init() {
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))
	vdl.Register((*Message)(nil))
}