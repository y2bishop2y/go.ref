// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: service.vdl

package cluster

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
)

func __VDLEnsureNativeBuilt_service() {
}

// ClusterAgentClientMethods is the client interface
// containing ClusterAgent methods.
type ClusterAgentClientMethods interface {
	// Retrieves all the blessings associated with a particular secret.
	// The only authorization required to access this method is the secret
	// itself.
	// TODO(rthellend): Consider adding other side-channel authorization
	// mechanisms, e.g. verify that the IP address of the client belongs to
	// an authorized user.
	SeekBlessings(_ *context.T, secret string, _ ...rpc.CallOpt) (security.Blessings, error)
}

// ClusterAgentClientStub adds universal methods to ClusterAgentClientMethods.
type ClusterAgentClientStub interface {
	ClusterAgentClientMethods
	rpc.UniversalServiceMethods
}

// ClusterAgentClient returns a client stub for ClusterAgent.
func ClusterAgentClient(name string) ClusterAgentClientStub {
	return implClusterAgentClientStub{name}
}

type implClusterAgentClientStub struct {
	name string
}

func (c implClusterAgentClientStub) SeekBlessings(ctx *context.T, i0 string, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "SeekBlessings", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// ClusterAgentServerMethods is the interface a server writer
// implements for ClusterAgent.
type ClusterAgentServerMethods interface {
	// Retrieves all the blessings associated with a particular secret.
	// The only authorization required to access this method is the secret
	// itself.
	// TODO(rthellend): Consider adding other side-channel authorization
	// mechanisms, e.g. verify that the IP address of the client belongs to
	// an authorized user.
	SeekBlessings(_ *context.T, _ rpc.ServerCall, secret string) (security.Blessings, error)
}

// ClusterAgentServerStubMethods is the server interface containing
// ClusterAgent methods, as expected by rpc.Server.
// There is no difference between this interface and ClusterAgentServerMethods
// since there are no streaming methods.
type ClusterAgentServerStubMethods ClusterAgentServerMethods

// ClusterAgentServerStub adds universal methods to ClusterAgentServerStubMethods.
type ClusterAgentServerStub interface {
	ClusterAgentServerStubMethods
	// Describe the ClusterAgent interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ClusterAgentServer returns a server stub for ClusterAgent.
// It converts an implementation of ClusterAgentServerMethods into
// an object that may be used by rpc.Server.
func ClusterAgentServer(impl ClusterAgentServerMethods) ClusterAgentServerStub {
	stub := implClusterAgentServerStub{
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

type implClusterAgentServerStub struct {
	impl ClusterAgentServerMethods
	gs   *rpc.GlobState
}

func (s implClusterAgentServerStub) SeekBlessings(ctx *context.T, call rpc.ServerCall, i0 string) (security.Blessings, error) {
	return s.impl.SeekBlessings(ctx, call, i0)
}

func (s implClusterAgentServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implClusterAgentServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ClusterAgentDesc}
}

// ClusterAgentDesc describes the ClusterAgent interface.
var ClusterAgentDesc rpc.InterfaceDesc = descClusterAgent

// descClusterAgent hides the desc to keep godoc clean.
var descClusterAgent = rpc.InterfaceDesc{
	Name:    "ClusterAgent",
	PkgPath: "v.io/x/ref/services/cluster",
	Methods: []rpc.MethodDesc{
		{
			Name: "SeekBlessings",
			Doc:  "// Retrieves all the blessings associated with a particular secret.\n// The only authorization required to access this method is the secret\n// itself.\n// TODO(rthellend): Consider adding other side-channel authorization\n// mechanisms, e.g. verify that the IP address of the client belongs to\n// an authorized user.",
			InArgs: []rpc.ArgDesc{
				{"secret", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
	},
}

// ClusterAgentAdminClientMethods is the client interface
// containing ClusterAgentAdmin methods.
type ClusterAgentAdminClientMethods interface {
	ClusterAgentClientMethods
	// Creates a new "secret" that can be used to retrieve extensions
	// of the blessings granted on this RPC, e.g. with the rpc.Granter
	// ClientCallOpt in Go.
	NewSecret(*context.T, ...rpc.CallOpt) (secret string, _ error)
	// Forgets a secret and its associated blessings.
	ForgetSecret(_ *context.T, secret string, _ ...rpc.CallOpt) error
}

// ClusterAgentAdminClientStub adds universal methods to ClusterAgentAdminClientMethods.
type ClusterAgentAdminClientStub interface {
	ClusterAgentAdminClientMethods
	rpc.UniversalServiceMethods
}

// ClusterAgentAdminClient returns a client stub for ClusterAgentAdmin.
func ClusterAgentAdminClient(name string) ClusterAgentAdminClientStub {
	return implClusterAgentAdminClientStub{name, ClusterAgentClient(name)}
}

type implClusterAgentAdminClientStub struct {
	name string

	ClusterAgentClientStub
}

func (c implClusterAgentAdminClientStub) NewSecret(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "NewSecret", nil, []interface{}{&o0}, opts...)
	return
}

func (c implClusterAgentAdminClientStub) ForgetSecret(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "ForgetSecret", []interface{}{i0}, nil, opts...)
	return
}

// ClusterAgentAdminServerMethods is the interface a server writer
// implements for ClusterAgentAdmin.
type ClusterAgentAdminServerMethods interface {
	ClusterAgentServerMethods
	// Creates a new "secret" that can be used to retrieve extensions
	// of the blessings granted on this RPC, e.g. with the rpc.Granter
	// ClientCallOpt in Go.
	NewSecret(*context.T, rpc.ServerCall) (secret string, _ error)
	// Forgets a secret and its associated blessings.
	ForgetSecret(_ *context.T, _ rpc.ServerCall, secret string) error
}

// ClusterAgentAdminServerStubMethods is the server interface containing
// ClusterAgentAdmin methods, as expected by rpc.Server.
// There is no difference between this interface and ClusterAgentAdminServerMethods
// since there are no streaming methods.
type ClusterAgentAdminServerStubMethods ClusterAgentAdminServerMethods

// ClusterAgentAdminServerStub adds universal methods to ClusterAgentAdminServerStubMethods.
type ClusterAgentAdminServerStub interface {
	ClusterAgentAdminServerStubMethods
	// Describe the ClusterAgentAdmin interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ClusterAgentAdminServer returns a server stub for ClusterAgentAdmin.
// It converts an implementation of ClusterAgentAdminServerMethods into
// an object that may be used by rpc.Server.
func ClusterAgentAdminServer(impl ClusterAgentAdminServerMethods) ClusterAgentAdminServerStub {
	stub := implClusterAgentAdminServerStub{
		impl: impl,
		ClusterAgentServerStub: ClusterAgentServer(impl),
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

type implClusterAgentAdminServerStub struct {
	impl ClusterAgentAdminServerMethods
	ClusterAgentServerStub
	gs *rpc.GlobState
}

func (s implClusterAgentAdminServerStub) NewSecret(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.NewSecret(ctx, call)
}

func (s implClusterAgentAdminServerStub) ForgetSecret(ctx *context.T, call rpc.ServerCall, i0 string) error {
	return s.impl.ForgetSecret(ctx, call, i0)
}

func (s implClusterAgentAdminServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implClusterAgentAdminServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ClusterAgentAdminDesc, ClusterAgentDesc}
}

// ClusterAgentAdminDesc describes the ClusterAgentAdmin interface.
var ClusterAgentAdminDesc rpc.InterfaceDesc = descClusterAgentAdmin

// descClusterAgentAdmin hides the desc to keep godoc clean.
var descClusterAgentAdmin = rpc.InterfaceDesc{
	Name:    "ClusterAgentAdmin",
	PkgPath: "v.io/x/ref/services/cluster",
	Embeds: []rpc.EmbedDesc{
		{"ClusterAgent", "v.io/x/ref/services/cluster", ``},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "NewSecret",
			Doc:  "// Creates a new \"secret\" that can be used to retrieve extensions\n// of the blessings granted on this RPC, e.g. with the rpc.Granter\n// ClientCallOpt in Go.",
			OutArgs: []rpc.ArgDesc{
				{"secret", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
		{
			Name: "ForgetSecret",
			Doc:  "// Forgets a secret and its associated blessings.",
			InArgs: []rpc.ArgDesc{
				{"secret", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}
