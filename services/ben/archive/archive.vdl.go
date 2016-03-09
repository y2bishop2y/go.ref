// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: archive.vdl

// Package archive defines the RPC interface for archiving benchmark results.
package archive

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/x/ref/services/ben"
)

func __VDLEnsureNativeBuilt_archive() {
}

// BenchmarkArchiverClientMethods is the client interface
// containing BenchmarkArchiver methods.
//
// BenchmarkArchiver is the interface to store microbenchmark results.
type BenchmarkArchiverClientMethods interface {
	// Archive saves results in 'runs' under the assumption that the
	// benchmarks were run on a machine whose configuration is defined by
	// 'scenario' and the were built from source code described by 'code'
	// (a commit hash of a repository, the manifest of a jiri project etc.)
	//
	// Returns a URL that can be used to browse the uploaded benchmark
	// results.
	Archive(_ *context.T, scenario ben.Scenario, code ben.SourceCode, runs []ben.Run, _ ...rpc.CallOpt) (string, error)
}

// BenchmarkArchiverClientStub adds universal methods to BenchmarkArchiverClientMethods.
type BenchmarkArchiverClientStub interface {
	BenchmarkArchiverClientMethods
	rpc.UniversalServiceMethods
}

// BenchmarkArchiverClient returns a client stub for BenchmarkArchiver.
func BenchmarkArchiverClient(name string) BenchmarkArchiverClientStub {
	return implBenchmarkArchiverClientStub{name}
}

type implBenchmarkArchiverClientStub struct {
	name string
}

func (c implBenchmarkArchiverClientStub) Archive(ctx *context.T, i0 ben.Scenario, i1 ben.SourceCode, i2 []ben.Run, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Archive", []interface{}{i0, i1, i2}, []interface{}{&o0}, opts...)
	return
}

// BenchmarkArchiverServerMethods is the interface a server writer
// implements for BenchmarkArchiver.
//
// BenchmarkArchiver is the interface to store microbenchmark results.
type BenchmarkArchiverServerMethods interface {
	// Archive saves results in 'runs' under the assumption that the
	// benchmarks were run on a machine whose configuration is defined by
	// 'scenario' and the were built from source code described by 'code'
	// (a commit hash of a repository, the manifest of a jiri project etc.)
	//
	// Returns a URL that can be used to browse the uploaded benchmark
	// results.
	Archive(_ *context.T, _ rpc.ServerCall, scenario ben.Scenario, code ben.SourceCode, runs []ben.Run) (string, error)
}

// BenchmarkArchiverServerStubMethods is the server interface containing
// BenchmarkArchiver methods, as expected by rpc.Server.
// There is no difference between this interface and BenchmarkArchiverServerMethods
// since there are no streaming methods.
type BenchmarkArchiverServerStubMethods BenchmarkArchiverServerMethods

// BenchmarkArchiverServerStub adds universal methods to BenchmarkArchiverServerStubMethods.
type BenchmarkArchiverServerStub interface {
	BenchmarkArchiverServerStubMethods
	// Describe the BenchmarkArchiver interfaces.
	Describe__() []rpc.InterfaceDesc
}

// BenchmarkArchiverServer returns a server stub for BenchmarkArchiver.
// It converts an implementation of BenchmarkArchiverServerMethods into
// an object that may be used by rpc.Server.
func BenchmarkArchiverServer(impl BenchmarkArchiverServerMethods) BenchmarkArchiverServerStub {
	stub := implBenchmarkArchiverServerStub{
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

type implBenchmarkArchiverServerStub struct {
	impl BenchmarkArchiverServerMethods
	gs   *rpc.GlobState
}

func (s implBenchmarkArchiverServerStub) Archive(ctx *context.T, call rpc.ServerCall, i0 ben.Scenario, i1 ben.SourceCode, i2 []ben.Run) (string, error) {
	return s.impl.Archive(ctx, call, i0, i1, i2)
}

func (s implBenchmarkArchiverServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implBenchmarkArchiverServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{BenchmarkArchiverDesc}
}

// BenchmarkArchiverDesc describes the BenchmarkArchiver interface.
var BenchmarkArchiverDesc rpc.InterfaceDesc = descBenchmarkArchiver

// descBenchmarkArchiver hides the desc to keep godoc clean.
var descBenchmarkArchiver = rpc.InterfaceDesc{
	Name:    "BenchmarkArchiver",
	PkgPath: "v.io/x/ref/services/ben/archive",
	Doc:     "// BenchmarkArchiver is the interface to store microbenchmark results.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Archive",
			Doc:  "// Archive saves results in 'runs' under the assumption that the\n// benchmarks were run on a machine whose configuration is defined by\n// 'scenario' and the were built from source code described by 'code'\n// (a commit hash of a repository, the manifest of a jiri project etc.)\n//\n// Returns a URL that can be used to browse the uploaded benchmark\n// results.",
			InArgs: []rpc.ArgDesc{
				{"scenario", ``}, // ben.Scenario
				{"code", ``},     // ben.SourceCode
				{"runs", ``},     // []ben.Run
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
	},
}
