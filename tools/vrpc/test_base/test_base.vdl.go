// This file was auto-generated by the veyron vdl tool.
// Source: test_base.vdl

package test_base

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "v.io/core/veyron2"
	__context "v.io/core/veyron2/context"
	__ipc "v.io/core/veyron2/ipc"
	__vdl "v.io/core/veyron2/vdl"
)

type Struct struct {
	X int32
	Y int32
}

func (Struct) __VDLReflect(struct {
	Name string "v.io/core/veyron/tools/vrpc/test_base.Struct"
}) {
}

type Array2Int [2]int32

func (Array2Int) __VDLReflect(struct {
	Name string "v.io/core/veyron/tools/vrpc/test_base.Array2Int"
}) {
}

func init() {
	__vdl.Register(Struct{})
	__vdl.Register(Array2Int{})
}

// TypeTesterClientMethods is the client interface
// containing TypeTester methods.
//
// TypeTester methods are listed in alphabetical order, to make it easier to
// test Signature output, which sorts methods alphabetically.
type TypeTesterClientMethods interface {
	// Methods to test support for primitive types.
	EchoBool(ctx *__context.T, I1 bool, opts ...__ipc.CallOpt) (O1 bool, err error)
	EchoFloat32(ctx *__context.T, I1 float32, opts ...__ipc.CallOpt) (O1 float32, err error)
	EchoFloat64(ctx *__context.T, I1 float64, opts ...__ipc.CallOpt) (O1 float64, err error)
	EchoInt32(ctx *__context.T, I1 int32, opts ...__ipc.CallOpt) (O1 int32, err error)
	EchoInt64(ctx *__context.T, I1 int64, opts ...__ipc.CallOpt) (O1 int64, err error)
	EchoString(ctx *__context.T, I1 string, opts ...__ipc.CallOpt) (O1 string, err error)
	EchoByte(ctx *__context.T, I1 byte, opts ...__ipc.CallOpt) (O1 byte, err error)
	EchoUint32(ctx *__context.T, I1 uint32, opts ...__ipc.CallOpt) (O1 uint32, err error)
	EchoUint64(ctx *__context.T, I1 uint64, opts ...__ipc.CallOpt) (O1 uint64, err error)
	// Methods to test support for composite types.
	XEchoArray(ctx *__context.T, I1 Array2Int, opts ...__ipc.CallOpt) (O1 Array2Int, err error)
	XEchoMap(ctx *__context.T, I1 map[int32]string, opts ...__ipc.CallOpt) (O1 map[int32]string, err error)
	XEchoSet(ctx *__context.T, I1 map[int32]struct{}, opts ...__ipc.CallOpt) (O1 map[int32]struct{}, err error)
	XEchoSlice(ctx *__context.T, I1 []int32, opts ...__ipc.CallOpt) (O1 []int32, err error)
	XEchoStruct(ctx *__context.T, I1 Struct, opts ...__ipc.CallOpt) (O1 Struct, err error)
	// Methods to test support for different number of arguments.
	YMultiArg(ctx *__context.T, I1 int32, I2 int32, opts ...__ipc.CallOpt) (O1 int32, O2 int32, err error)
	YNoArgs(*__context.T, ...__ipc.CallOpt) error
	// Methods to test support for streaming.
	ZStream(ctx *__context.T, NumStreamItems int32, StreamItem bool, opts ...__ipc.CallOpt) (TypeTesterZStreamCall, error)
}

// TypeTesterClientStub adds universal methods to TypeTesterClientMethods.
type TypeTesterClientStub interface {
	TypeTesterClientMethods
	__ipc.UniversalServiceMethods
}

// TypeTesterClient returns a client stub for TypeTester.
func TypeTesterClient(name string, opts ...__ipc.BindOpt) TypeTesterClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implTypeTesterClientStub{name, client}
}

type implTypeTesterClientStub struct {
	name   string
	client __ipc.Client
}

func (c implTypeTesterClientStub) c(ctx *__context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.GetClient(ctx)
}

func (c implTypeTesterClientStub) EchoBool(ctx *__context.T, i0 bool, opts ...__ipc.CallOpt) (o0 bool, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoBool", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoFloat32(ctx *__context.T, i0 float32, opts ...__ipc.CallOpt) (o0 float32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoFloat32", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoFloat64(ctx *__context.T, i0 float64, opts ...__ipc.CallOpt) (o0 float64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoFloat64", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoInt32(ctx *__context.T, i0 int32, opts ...__ipc.CallOpt) (o0 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoInt32", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoInt64(ctx *__context.T, i0 int64, opts ...__ipc.CallOpt) (o0 int64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoInt64", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoString(ctx *__context.T, i0 string, opts ...__ipc.CallOpt) (o0 string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoString", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoByte(ctx *__context.T, i0 byte, opts ...__ipc.CallOpt) (o0 byte, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoByte", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoUint32(ctx *__context.T, i0 uint32, opts ...__ipc.CallOpt) (o0 uint32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoUint32", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) EchoUint64(ctx *__context.T, i0 uint64, opts ...__ipc.CallOpt) (o0 uint64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoUint64", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) XEchoArray(ctx *__context.T, i0 Array2Int, opts ...__ipc.CallOpt) (o0 Array2Int, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "XEchoArray", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) XEchoMap(ctx *__context.T, i0 map[int32]string, opts ...__ipc.CallOpt) (o0 map[int32]string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "XEchoMap", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) XEchoSet(ctx *__context.T, i0 map[int32]struct{}, opts ...__ipc.CallOpt) (o0 map[int32]struct{}, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "XEchoSet", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) XEchoSlice(ctx *__context.T, i0 []int32, opts ...__ipc.CallOpt) (o0 []int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "XEchoSlice", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) XEchoStruct(ctx *__context.T, i0 Struct, opts ...__ipc.CallOpt) (o0 Struct, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "XEchoStruct", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) YMultiArg(ctx *__context.T, i0 int32, i1 int32, opts ...__ipc.CallOpt) (o0 int32, o1 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "YMultiArg", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &o1, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) YNoArgs(ctx *__context.T, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "YNoArgs", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTypeTesterClientStub) ZStream(ctx *__context.T, i0 int32, i1 bool, opts ...__ipc.CallOpt) (ocall TypeTesterZStreamCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "ZStream", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implTypeTesterZStreamCall{Call: call}
	return
}

// TypeTesterZStreamClientStream is the client stream for TypeTester.ZStream.
type TypeTesterZStreamClientStream interface {
	// RecvStream returns the receiver side of the TypeTester.ZStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() bool
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// TypeTesterZStreamCall represents the call returned from TypeTester.ZStream.
type TypeTesterZStreamCall interface {
	TypeTesterZStreamClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implTypeTesterZStreamCall struct {
	__ipc.Call
	valRecv bool
	errRecv error
}

func (c *implTypeTesterZStreamCall) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implTypeTesterZStreamCallRecv{c}
}

type implTypeTesterZStreamCallRecv struct {
	c *implTypeTesterZStreamCall
}

func (c implTypeTesterZStreamCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implTypeTesterZStreamCallRecv) Value() bool {
	return c.c.valRecv
}
func (c implTypeTesterZStreamCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implTypeTesterZStreamCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// TypeTesterServerMethods is the interface a server writer
// implements for TypeTester.
//
// TypeTester methods are listed in alphabetical order, to make it easier to
// test Signature output, which sorts methods alphabetically.
type TypeTesterServerMethods interface {
	// Methods to test support for primitive types.
	EchoBool(ctx __ipc.ServerContext, I1 bool) (O1 bool, err error)
	EchoFloat32(ctx __ipc.ServerContext, I1 float32) (O1 float32, err error)
	EchoFloat64(ctx __ipc.ServerContext, I1 float64) (O1 float64, err error)
	EchoInt32(ctx __ipc.ServerContext, I1 int32) (O1 int32, err error)
	EchoInt64(ctx __ipc.ServerContext, I1 int64) (O1 int64, err error)
	EchoString(ctx __ipc.ServerContext, I1 string) (O1 string, err error)
	EchoByte(ctx __ipc.ServerContext, I1 byte) (O1 byte, err error)
	EchoUint32(ctx __ipc.ServerContext, I1 uint32) (O1 uint32, err error)
	EchoUint64(ctx __ipc.ServerContext, I1 uint64) (O1 uint64, err error)
	// Methods to test support for composite types.
	XEchoArray(ctx __ipc.ServerContext, I1 Array2Int) (O1 Array2Int, err error)
	XEchoMap(ctx __ipc.ServerContext, I1 map[int32]string) (O1 map[int32]string, err error)
	XEchoSet(ctx __ipc.ServerContext, I1 map[int32]struct{}) (O1 map[int32]struct{}, err error)
	XEchoSlice(ctx __ipc.ServerContext, I1 []int32) (O1 []int32, err error)
	XEchoStruct(ctx __ipc.ServerContext, I1 Struct) (O1 Struct, err error)
	// Methods to test support for different number of arguments.
	YMultiArg(ctx __ipc.ServerContext, I1 int32, I2 int32) (O1 int32, O2 int32, err error)
	YNoArgs(__ipc.ServerContext) error
	// Methods to test support for streaming.
	ZStream(ctx TypeTesterZStreamContext, NumStreamItems int32, StreamItem bool) error
}

// TypeTesterServerStubMethods is the server interface containing
// TypeTester methods, as expected by ipc.Server.
// The only difference between this interface and TypeTesterServerMethods
// is the streaming methods.
type TypeTesterServerStubMethods interface {
	// Methods to test support for primitive types.
	EchoBool(ctx __ipc.ServerContext, I1 bool) (O1 bool, err error)
	EchoFloat32(ctx __ipc.ServerContext, I1 float32) (O1 float32, err error)
	EchoFloat64(ctx __ipc.ServerContext, I1 float64) (O1 float64, err error)
	EchoInt32(ctx __ipc.ServerContext, I1 int32) (O1 int32, err error)
	EchoInt64(ctx __ipc.ServerContext, I1 int64) (O1 int64, err error)
	EchoString(ctx __ipc.ServerContext, I1 string) (O1 string, err error)
	EchoByte(ctx __ipc.ServerContext, I1 byte) (O1 byte, err error)
	EchoUint32(ctx __ipc.ServerContext, I1 uint32) (O1 uint32, err error)
	EchoUint64(ctx __ipc.ServerContext, I1 uint64) (O1 uint64, err error)
	// Methods to test support for composite types.
	XEchoArray(ctx __ipc.ServerContext, I1 Array2Int) (O1 Array2Int, err error)
	XEchoMap(ctx __ipc.ServerContext, I1 map[int32]string) (O1 map[int32]string, err error)
	XEchoSet(ctx __ipc.ServerContext, I1 map[int32]struct{}) (O1 map[int32]struct{}, err error)
	XEchoSlice(ctx __ipc.ServerContext, I1 []int32) (O1 []int32, err error)
	XEchoStruct(ctx __ipc.ServerContext, I1 Struct) (O1 Struct, err error)
	// Methods to test support for different number of arguments.
	YMultiArg(ctx __ipc.ServerContext, I1 int32, I2 int32) (O1 int32, O2 int32, err error)
	YNoArgs(__ipc.ServerContext) error
	// Methods to test support for streaming.
	ZStream(ctx *TypeTesterZStreamContextStub, NumStreamItems int32, StreamItem bool) error
}

// TypeTesterServerStub adds universal methods to TypeTesterServerStubMethods.
type TypeTesterServerStub interface {
	TypeTesterServerStubMethods
	// Describe the TypeTester interfaces.
	Describe__() []__ipc.InterfaceDesc
}

// TypeTesterServer returns a server stub for TypeTester.
// It converts an implementation of TypeTesterServerMethods into
// an object that may be used by ipc.Server.
func TypeTesterServer(impl TypeTesterServerMethods) TypeTesterServerStub {
	stub := implTypeTesterServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implTypeTesterServerStub struct {
	impl TypeTesterServerMethods
	gs   *__ipc.GlobState
}

func (s implTypeTesterServerStub) EchoBool(ctx __ipc.ServerContext, i0 bool) (bool, error) {
	return s.impl.EchoBool(ctx, i0)
}

func (s implTypeTesterServerStub) EchoFloat32(ctx __ipc.ServerContext, i0 float32) (float32, error) {
	return s.impl.EchoFloat32(ctx, i0)
}

func (s implTypeTesterServerStub) EchoFloat64(ctx __ipc.ServerContext, i0 float64) (float64, error) {
	return s.impl.EchoFloat64(ctx, i0)
}

func (s implTypeTesterServerStub) EchoInt32(ctx __ipc.ServerContext, i0 int32) (int32, error) {
	return s.impl.EchoInt32(ctx, i0)
}

func (s implTypeTesterServerStub) EchoInt64(ctx __ipc.ServerContext, i0 int64) (int64, error) {
	return s.impl.EchoInt64(ctx, i0)
}

func (s implTypeTesterServerStub) EchoString(ctx __ipc.ServerContext, i0 string) (string, error) {
	return s.impl.EchoString(ctx, i0)
}

func (s implTypeTesterServerStub) EchoByte(ctx __ipc.ServerContext, i0 byte) (byte, error) {
	return s.impl.EchoByte(ctx, i0)
}

func (s implTypeTesterServerStub) EchoUint32(ctx __ipc.ServerContext, i0 uint32) (uint32, error) {
	return s.impl.EchoUint32(ctx, i0)
}

func (s implTypeTesterServerStub) EchoUint64(ctx __ipc.ServerContext, i0 uint64) (uint64, error) {
	return s.impl.EchoUint64(ctx, i0)
}

func (s implTypeTesterServerStub) XEchoArray(ctx __ipc.ServerContext, i0 Array2Int) (Array2Int, error) {
	return s.impl.XEchoArray(ctx, i0)
}

func (s implTypeTesterServerStub) XEchoMap(ctx __ipc.ServerContext, i0 map[int32]string) (map[int32]string, error) {
	return s.impl.XEchoMap(ctx, i0)
}

func (s implTypeTesterServerStub) XEchoSet(ctx __ipc.ServerContext, i0 map[int32]struct{}) (map[int32]struct{}, error) {
	return s.impl.XEchoSet(ctx, i0)
}

func (s implTypeTesterServerStub) XEchoSlice(ctx __ipc.ServerContext, i0 []int32) ([]int32, error) {
	return s.impl.XEchoSlice(ctx, i0)
}

func (s implTypeTesterServerStub) XEchoStruct(ctx __ipc.ServerContext, i0 Struct) (Struct, error) {
	return s.impl.XEchoStruct(ctx, i0)
}

func (s implTypeTesterServerStub) YMultiArg(ctx __ipc.ServerContext, i0 int32, i1 int32) (int32, int32, error) {
	return s.impl.YMultiArg(ctx, i0, i1)
}

func (s implTypeTesterServerStub) YNoArgs(ctx __ipc.ServerContext) error {
	return s.impl.YNoArgs(ctx)
}

func (s implTypeTesterServerStub) ZStream(ctx *TypeTesterZStreamContextStub, i0 int32, i1 bool) error {
	return s.impl.ZStream(ctx, i0, i1)
}

func (s implTypeTesterServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implTypeTesterServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{TypeTesterDesc}
}

// TypeTesterDesc describes the TypeTester interface.
var TypeTesterDesc __ipc.InterfaceDesc = descTypeTester

// descTypeTester hides the desc to keep godoc clean.
var descTypeTester = __ipc.InterfaceDesc{
	Name:    "TypeTester",
	PkgPath: "v.io/core/veyron/tools/vrpc/test_base",
	Doc:     "// TypeTester methods are listed in alphabetical order, to make it easier to\n// test Signature output, which sorts methods alphabetically.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "EchoBool",
			Doc:  "// Methods to test support for primitive types.",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // bool
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // bool
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoFloat32",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // float32
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // float32
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoFloat64",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // float64
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // float64
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoInt32",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // int32
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoInt64",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // int64
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // int64
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoString",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // string
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // string
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoByte",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // byte
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // byte
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoUint32",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // uint32
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // uint32
				{"err", ``}, // error
			},
		},
		{
			Name: "EchoUint64",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // uint64
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // uint64
				{"err", ``}, // error
			},
		},
		{
			Name: "XEchoArray",
			Doc:  "// Methods to test support for composite types.",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // Array2Int
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // Array2Int
				{"err", ``}, // error
			},
		},
		{
			Name: "XEchoMap",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // map[int32]string
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // map[int32]string
				{"err", ``}, // error
			},
		},
		{
			Name: "XEchoSet",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // map[int32]struct{}
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // map[int32]struct{}
				{"err", ``}, // error
			},
		},
		{
			Name: "XEchoSlice",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // []int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // []int32
				{"err", ``}, // error
			},
		},
		{
			Name: "XEchoStruct",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // Struct
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // Struct
				{"err", ``}, // error
			},
		},
		{
			Name: "YMultiArg",
			Doc:  "// Methods to test support for different number of arguments.",
			InArgs: []__ipc.ArgDesc{
				{"I1", ``}, // int32
				{"I2", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"O1", ``},  // int32
				{"O2", ``},  // int32
				{"err", ``}, // error
			},
		},
		{
			Name: "YNoArgs",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "ZStream",
			Doc:  "// Methods to test support for streaming.",
			InArgs: []__ipc.ArgDesc{
				{"NumStreamItems", ``}, // int32
				{"StreamItem", ``},     // bool
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
	},
}

// TypeTesterZStreamServerStream is the server stream for TypeTester.ZStream.
type TypeTesterZStreamServerStream interface {
	// SendStream returns the send side of the TypeTester.ZStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item bool) error
	}
}

// TypeTesterZStreamContext represents the context passed to TypeTester.ZStream.
type TypeTesterZStreamContext interface {
	__ipc.ServerContext
	TypeTesterZStreamServerStream
}

// TypeTesterZStreamContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements TypeTesterZStreamContext.
type TypeTesterZStreamContextStub struct {
	__ipc.ServerCall
}

// Init initializes TypeTesterZStreamContextStub from ipc.ServerCall.
func (s *TypeTesterZStreamContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the TypeTester.ZStream server stream.
func (s *TypeTesterZStreamContextStub) SendStream() interface {
	Send(item bool) error
} {
	return implTypeTesterZStreamContextSend{s}
}

type implTypeTesterZStreamContextSend struct {
	s *TypeTesterZStreamContextStub
}

func (s implTypeTesterZStreamContextSend) Send(item bool) error {
	return s.s.Send(item)
}
