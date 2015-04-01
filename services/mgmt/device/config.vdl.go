// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: config.vdl

package device

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// ConfigClientMethods is the client interface
// containing Config methods.
//
// Config is an RPC API to the config service.
type ConfigClientMethods interface {
	// Set sets the value for key.
	Set(ctx *context.T, key string, value string, opts ...rpc.CallOpt) error
}

// ConfigClientStub adds universal methods to ConfigClientMethods.
type ConfigClientStub interface {
	ConfigClientMethods
	rpc.UniversalServiceMethods
}

// ConfigClient returns a client stub for Config.
func ConfigClient(name string) ConfigClientStub {
	return implConfigClientStub{name}
}

type implConfigClientStub struct {
	name string
}

func (c implConfigClientStub) Set(ctx *context.T, i0 string, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Set", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// ConfigServerMethods is the interface a server writer
// implements for Config.
//
// Config is an RPC API to the config service.
type ConfigServerMethods interface {
	// Set sets the value for key.
	Set(call rpc.ServerCall, key string, value string) error
}

// ConfigServerStubMethods is the server interface containing
// Config methods, as expected by rpc.Server.
// There is no difference between this interface and ConfigServerMethods
// since there are no streaming methods.
type ConfigServerStubMethods ConfigServerMethods

// ConfigServerStub adds universal methods to ConfigServerStubMethods.
type ConfigServerStub interface {
	ConfigServerStubMethods
	// Describe the Config interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ConfigServer returns a server stub for Config.
// It converts an implementation of ConfigServerMethods into
// an object that may be used by rpc.Server.
func ConfigServer(impl ConfigServerMethods) ConfigServerStub {
	stub := implConfigServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implConfigServerStub struct {
	impl ConfigServerMethods
	gs   *rpc.GlobState
}

func (s implConfigServerStub) Set(call rpc.ServerCall, i0 string, i1 string) error {
	return s.impl.Set(call, i0, i1)
}

func (s implConfigServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implConfigServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ConfigDesc}
}

// ConfigDesc describes the Config interface.
var ConfigDesc rpc.InterfaceDesc = descConfig

// descConfig hides the desc to keep godoc clean.
var descConfig = rpc.InterfaceDesc{
	Name:    "Config",
	PkgPath: "v.io/x/ref/services/mgmt/device",
	Doc:     "// Config is an RPC API to the config service.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Set",
			Doc:  "// Set sets the value for key.",
			InArgs: []rpc.ArgDesc{
				{"key", ``},   // string
				{"value", ``}, // string
			},
		},
	},
}
