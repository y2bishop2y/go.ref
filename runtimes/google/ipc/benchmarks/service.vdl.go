// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// package benchmark provides simple tools to measure the performance of the
// IPC system.
package benchmarks

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// BenchmarkClientMethods is the client interface
// containing Benchmark methods.
type BenchmarkClientMethods interface {
	// Echo returns the payload that it receives.
	Echo(ctx __context.T, Payload []byte, opts ...__ipc.CallOpt) ([]byte, error)
	// EchoStream returns the payload that it receives via the stream.
	EchoStream(__context.T, ...__ipc.CallOpt) (BenchmarkEchoStreamCall, error)
}

// BenchmarkClientStub adds universal methods to BenchmarkClientMethods.
type BenchmarkClientStub interface {
	BenchmarkClientMethods
	__ipc.UniversalServiceMethods
}

// BenchmarkClient returns a client stub for Benchmark.
func BenchmarkClient(name string, opts ...__ipc.BindOpt) BenchmarkClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implBenchmarkClientStub{name, client}
}

type implBenchmarkClientStub struct {
	name   string
	client __ipc.Client
}

func (c implBenchmarkClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implBenchmarkClientStub) Echo(ctx __context.T, i0 []byte, opts ...__ipc.CallOpt) (o0 []byte, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Echo", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implBenchmarkClientStub) EchoStream(ctx __context.T, opts ...__ipc.CallOpt) (ocall BenchmarkEchoStreamCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "EchoStream", nil, opts...); err != nil {
		return
	}
	ocall = &implBenchmarkEchoStreamCall{Call: call}
	return
}

func (c implBenchmarkClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implBenchmarkClientStub) GetMethodTags(ctx __context.T, method string, opts ...__ipc.CallOpt) (o0 []interface{}, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// BenchmarkEchoStreamClientStream is the client stream for Benchmark.EchoStream.
type BenchmarkEchoStreamClientStream interface {
	// RecvStream returns the receiver side of the Benchmark.EchoStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Benchmark.EchoStream client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending, or if Send is called after Close or Cancel.  Blocks if
		// there is no buffer space; will unblock when buffer space is available or
		// after Cancel.
		Send(item []byte) error
		// Close indicates to the server that no more items will be sent; server
		// Recv calls will receive io.EOF after all sent items.  This is an optional
		// call - e.g. a client might call Close if it needs to continue receiving
		// items from the server after it's done sending.  Returns errors
		// encountered while closing, or if Close is called after Cancel.  Like
		// Send, blocks if there is no buffer space available.
		Close() error
	}
}

// BenchmarkEchoStreamCall represents the call returned from Benchmark.EchoStream.
type BenchmarkEchoStreamCall interface {
	BenchmarkEchoStreamClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implBenchmarkEchoStreamCall struct {
	__ipc.Call
	valRecv []byte
	errRecv error
}

func (c *implBenchmarkEchoStreamCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implBenchmarkEchoStreamCallRecv{c}
}

type implBenchmarkEchoStreamCallRecv struct {
	c *implBenchmarkEchoStreamCall
}

func (c implBenchmarkEchoStreamCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implBenchmarkEchoStreamCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implBenchmarkEchoStreamCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implBenchmarkEchoStreamCall) SendStream() interface {
	Send(item []byte) error
	Close() error
} {
	return implBenchmarkEchoStreamCallSend{c}
}

type implBenchmarkEchoStreamCallSend struct {
	c *implBenchmarkEchoStreamCall
}

func (c implBenchmarkEchoStreamCallSend) Send(item []byte) error {
	return c.c.Send(item)
}
func (c implBenchmarkEchoStreamCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implBenchmarkEchoStreamCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// BenchmarkServerMethods is the interface a server writer
// implements for Benchmark.
type BenchmarkServerMethods interface {
	// Echo returns the payload that it receives.
	Echo(ctx __ipc.ServerContext, Payload []byte) ([]byte, error)
	// EchoStream returns the payload that it receives via the stream.
	EchoStream(BenchmarkEchoStreamContext) error
}

// BenchmarkServerStubMethods is the server interface containing
// Benchmark methods, as expected by ipc.Server.
// The only difference between this interface and BenchmarkServerMethods
// is the streaming methods.
type BenchmarkServerStubMethods interface {
	// Echo returns the payload that it receives.
	Echo(ctx __ipc.ServerContext, Payload []byte) ([]byte, error)
	// EchoStream returns the payload that it receives via the stream.
	EchoStream(*BenchmarkEchoStreamContextStub) error
}

// BenchmarkServerStub adds universal methods to BenchmarkServerStubMethods.
type BenchmarkServerStub interface {
	BenchmarkServerStubMethods
	// GetMethodTags will be replaced with DescribeInterfaces.
	GetMethodTags(ctx __ipc.ServerContext, method string) ([]interface{}, error)
	// Signature will be replaced with DescribeInterfaces.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// BenchmarkServer returns a server stub for Benchmark.
// It converts an implementation of BenchmarkServerMethods into
// an object that may be used by ipc.Server.
func BenchmarkServer(impl BenchmarkServerMethods) BenchmarkServerStub {
	stub := implBenchmarkServerStub{
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

type implBenchmarkServerStub struct {
	impl BenchmarkServerMethods
	gs   *__ipc.GlobState
}

func (s implBenchmarkServerStub) Echo(ctx __ipc.ServerContext, i0 []byte) ([]byte, error) {
	return s.impl.Echo(ctx, i0)
}

func (s implBenchmarkServerStub) EchoStream(ctx *BenchmarkEchoStreamContextStub) error {
	return s.impl.EchoStream(ctx)
}

func (s implBenchmarkServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implBenchmarkServerStub) GetMethodTags(ctx __ipc.ServerContext, method string) ([]interface{}, error) {
	// TODO(toddw): Replace with new DescribeInterfaces implementation.
	switch method {
	case "Echo":
		return []interface{}{}, nil
	case "EchoStream":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (s implBenchmarkServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw) Replace with new DescribeInterfaces implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Echo"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "Payload", Type: 66},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 66},
			{Name: "", Type: 67},
		},
	}
	result.Methods["EchoStream"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 67},
		},
		InStream:  66,
		OutStream: 66,
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, __wiretype.SliceType{Elem: 0x41, Name: "", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

// BenchmarkEchoStreamServerStream is the server stream for Benchmark.EchoStream.
type BenchmarkEchoStreamServerStream interface {
	// RecvStream returns the receiver side of the Benchmark.EchoStream server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Benchmark.EchoStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// BenchmarkEchoStreamContext represents the context passed to Benchmark.EchoStream.
type BenchmarkEchoStreamContext interface {
	__ipc.ServerContext
	BenchmarkEchoStreamServerStream
}

// BenchmarkEchoStreamContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements BenchmarkEchoStreamContext.
type BenchmarkEchoStreamContextStub struct {
	__ipc.ServerCall
	valRecv []byte
	errRecv error
}

// Init initializes BenchmarkEchoStreamContextStub from ipc.ServerCall.
func (s *BenchmarkEchoStreamContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// RecvStream returns the receiver side of the Benchmark.EchoStream server stream.
func (s *BenchmarkEchoStreamContextStub) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implBenchmarkEchoStreamContextRecv{s}
}

type implBenchmarkEchoStreamContextRecv struct {
	s *BenchmarkEchoStreamContextStub
}

func (s implBenchmarkEchoStreamContextRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implBenchmarkEchoStreamContextRecv) Value() []byte {
	return s.s.valRecv
}
func (s implBenchmarkEchoStreamContextRecv) Err() error {
	if s.s.errRecv == __io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Benchmark.EchoStream server stream.
func (s *BenchmarkEchoStreamContextStub) SendStream() interface {
	Send(item []byte) error
} {
	return implBenchmarkEchoStreamContextSend{s}
}

type implBenchmarkEchoStreamContextSend struct {
	s *BenchmarkEchoStreamContextStub
}

func (s implBenchmarkEchoStreamContextSend) Send(item []byte) error {
	return s.s.Send(item)
}
