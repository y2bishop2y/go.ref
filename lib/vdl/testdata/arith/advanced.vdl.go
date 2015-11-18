// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: advanced.vdl

package arith

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"

	// VDL user imports
	"v.io/x/ref/lib/vdl/testdata/arith/exp"
)

// TrigonometryClientMethods is the client interface
// containing Trigonometry methods.
//
// Trigonometry is an interface that specifies a couple trigonometric functions.
type TrigonometryClientMethods interface {
	Sine(_ *context.T, angle float64, _ ...rpc.CallOpt) (float64, error)
	Cosine(_ *context.T, angle float64, _ ...rpc.CallOpt) (float64, error)
}

// TrigonometryClientStub adds universal methods to TrigonometryClientMethods.
type TrigonometryClientStub interface {
	TrigonometryClientMethods
	rpc.UniversalServiceMethods
}

// TrigonometryClient returns a client stub for Trigonometry.
func TrigonometryClient(name string) TrigonometryClientStub {
	return implTrigonometryClientStub{name}
}

type implTrigonometryClientStub struct {
	name string
}

func (c implTrigonometryClientStub) Sine(ctx *context.T, i0 float64, opts ...rpc.CallOpt) (o0 float64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sine", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTrigonometryClientStub) Cosine(ctx *context.T, i0 float64, opts ...rpc.CallOpt) (o0 float64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Cosine", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// TrigonometryServerMethods is the interface a server writer
// implements for Trigonometry.
//
// Trigonometry is an interface that specifies a couple trigonometric functions.
type TrigonometryServerMethods interface {
	Sine(_ *context.T, _ rpc.ServerCall, angle float64) (float64, error)
	Cosine(_ *context.T, _ rpc.ServerCall, angle float64) (float64, error)
}

// TrigonometryServerStubMethods is the server interface containing
// Trigonometry methods, as expected by rpc.Server.
// There is no difference between this interface and TrigonometryServerMethods
// since there are no streaming methods.
type TrigonometryServerStubMethods TrigonometryServerMethods

// TrigonometryServerStub adds universal methods to TrigonometryServerStubMethods.
type TrigonometryServerStub interface {
	TrigonometryServerStubMethods
	// Describe the Trigonometry interfaces.
	Describe__() []rpc.InterfaceDesc
}

// TrigonometryServer returns a server stub for Trigonometry.
// It converts an implementation of TrigonometryServerMethods into
// an object that may be used by rpc.Server.
func TrigonometryServer(impl TrigonometryServerMethods) TrigonometryServerStub {
	stub := implTrigonometryServerStub{
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

type implTrigonometryServerStub struct {
	impl TrigonometryServerMethods
	gs   *rpc.GlobState
}

func (s implTrigonometryServerStub) Sine(ctx *context.T, call rpc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Sine(ctx, call, i0)
}

func (s implTrigonometryServerStub) Cosine(ctx *context.T, call rpc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Cosine(ctx, call, i0)
}

func (s implTrigonometryServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implTrigonometryServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{TrigonometryDesc}
}

// TrigonometryDesc describes the Trigonometry interface.
var TrigonometryDesc rpc.InterfaceDesc = descTrigonometry

// descTrigonometry hides the desc to keep godoc clean.
var descTrigonometry = rpc.InterfaceDesc{
	Name:    "Trigonometry",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith",
	Doc:     "// Trigonometry is an interface that specifies a couple trigonometric functions.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Sine",
			InArgs: []rpc.ArgDesc{
				{"angle", ``}, // float64
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // float64
			},
		},
		{
			Name: "Cosine",
			InArgs: []rpc.ArgDesc{
				{"angle", ``}, // float64
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // float64
			},
		},
	},
}

// AdvancedMathClientMethods is the client interface
// containing AdvancedMath methods.
//
// AdvancedMath is an interface for more advanced math than arith.  It embeds
// interfaces defined both in the same file and in an external package; and in
// turn it is embedded by arith.Calculator (which is in the same package but
// different file) to verify that embedding works in all these scenarios.
type AdvancedMathClientMethods interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryClientMethods
	exp.ExpClientMethods
}

// AdvancedMathClientStub adds universal methods to AdvancedMathClientMethods.
type AdvancedMathClientStub interface {
	AdvancedMathClientMethods
	rpc.UniversalServiceMethods
}

// AdvancedMathClient returns a client stub for AdvancedMath.
func AdvancedMathClient(name string) AdvancedMathClientStub {
	return implAdvancedMathClientStub{name, TrigonometryClient(name), exp.ExpClient(name)}
}

type implAdvancedMathClientStub struct {
	name string

	TrigonometryClientStub
	exp.ExpClientStub
}

// AdvancedMathServerMethods is the interface a server writer
// implements for AdvancedMath.
//
// AdvancedMath is an interface for more advanced math than arith.  It embeds
// interfaces defined both in the same file and in an external package; and in
// turn it is embedded by arith.Calculator (which is in the same package but
// different file) to verify that embedding works in all these scenarios.
type AdvancedMathServerMethods interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryServerMethods
	exp.ExpServerMethods
}

// AdvancedMathServerStubMethods is the server interface containing
// AdvancedMath methods, as expected by rpc.Server.
// There is no difference between this interface and AdvancedMathServerMethods
// since there are no streaming methods.
type AdvancedMathServerStubMethods AdvancedMathServerMethods

// AdvancedMathServerStub adds universal methods to AdvancedMathServerStubMethods.
type AdvancedMathServerStub interface {
	AdvancedMathServerStubMethods
	// Describe the AdvancedMath interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AdvancedMathServer returns a server stub for AdvancedMath.
// It converts an implementation of AdvancedMathServerMethods into
// an object that may be used by rpc.Server.
func AdvancedMathServer(impl AdvancedMathServerMethods) AdvancedMathServerStub {
	stub := implAdvancedMathServerStub{
		impl: impl,
		TrigonometryServerStub: TrigonometryServer(impl),
		ExpServerStub:          exp.ExpServer(impl),
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

type implAdvancedMathServerStub struct {
	impl AdvancedMathServerMethods
	TrigonometryServerStub
	exp.ExpServerStub
	gs *rpc.GlobState
}

func (s implAdvancedMathServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAdvancedMathServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AdvancedMathDesc, TrigonometryDesc, exp.ExpDesc}
}

// AdvancedMathDesc describes the AdvancedMath interface.
var AdvancedMathDesc rpc.InterfaceDesc = descAdvancedMath

// descAdvancedMath hides the desc to keep godoc clean.
var descAdvancedMath = rpc.InterfaceDesc{
	Name:    "AdvancedMath",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith",
	Doc:     "// AdvancedMath is an interface for more advanced math than arith.  It embeds\n// interfaces defined both in the same file and in an external package; and in\n// turn it is embedded by arith.Calculator (which is in the same package but\n// different file) to verify that embedding works in all these scenarios.",
	Embeds: []rpc.EmbedDesc{
		{"Trigonometry", "v.io/x/ref/lib/vdl/testdata/arith", "// Trigonometry is an interface that specifies a couple trigonometric functions."},
		{"Exp", "v.io/x/ref/lib/vdl/testdata/arith/exp", ``},
	},
}
