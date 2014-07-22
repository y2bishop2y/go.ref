// This file was auto-generated by the veyron vdl tool.
// Source: inspector.vdl

package inspector

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

type Details struct {
	Name string
	Size int64
	Mode uint32
	// TODO: add native time format to the idl+vom
	// Seconds since the start of the Unix epoch
	ModUnixSecs int64
	// Nanoseconds in the current second
	ModNano int32
	IsDir   bool
}

// Inspector is the interface the client binds and uses.
// Inspector_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Inspector_ExcludingUniversal interface {
	Ls(ctx _gen_context.T, Glob string, opts ..._gen_ipc.CallOpt) (reply InspectorLsStream, err error)
	LsDetails(ctx _gen_context.T, Glob string, opts ..._gen_ipc.CallOpt) (reply InspectorLsDetailsStream, err error)
}
type Inspector interface {
	_gen_ipc.UniversalServiceMethods
	Inspector_ExcludingUniversal
}

// InspectorService is the interface the server implements.
type InspectorService interface {
	Ls(context _gen_ipc.ServerContext, Glob string, stream InspectorServiceLsStream) (err error)
	LsDetails(context _gen_ipc.ServerContext, Glob string, stream InspectorServiceLsDetailsStream) (err error)
}

// InspectorLsStream is the interface for streaming responses of the method
// Ls in the service interface Inspector.
type InspectorLsStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item string, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the InspectorLsStream interface that is not exported.
type implInspectorLsStream struct {
	clientCall _gen_ipc.Call
}

func (c *implInspectorLsStream) Recv() (item string, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implInspectorLsStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implInspectorLsStream) Cancel() {
	c.clientCall.Cancel()
}

// InspectorServiceLsStream is the interface for streaming responses of the method
// Ls in the service interface Inspector.
type InspectorServiceLsStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item string) error
}

// Implementation of the InspectorServiceLsStream interface that is not exported.
type implInspectorServiceLsStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implInspectorServiceLsStream) Send(item string) error {
	return s.serverCall.Send(item)
}

// InspectorLsDetailsStream is the interface for streaming responses of the method
// LsDetails in the service interface Inspector.
type InspectorLsDetailsStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item Details, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the InspectorLsDetailsStream interface that is not exported.
type implInspectorLsDetailsStream struct {
	clientCall _gen_ipc.Call
}

func (c *implInspectorLsDetailsStream) Recv() (item Details, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implInspectorLsDetailsStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implInspectorLsDetailsStream) Cancel() {
	c.clientCall.Cancel()
}

// InspectorServiceLsDetailsStream is the interface for streaming responses of the method
// LsDetails in the service interface Inspector.
type InspectorServiceLsDetailsStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item Details) error
}

// Implementation of the InspectorServiceLsDetailsStream interface that is not exported.
type implInspectorServiceLsDetailsStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implInspectorServiceLsDetailsStream) Send(item Details) error {
	return s.serverCall.Send(item)
}

// BindInspector returns the client stub implementing the Inspector
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindInspector(name string, opts ..._gen_ipc.BindOpt) (Inspector, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubInspector{client: client, name: name}

	return stub, nil
}

// NewServerInspector creates a new server stub.
//
// It takes a regular server implementing the InspectorService
// interface, and returns a new server stub.
func NewServerInspector(server InspectorService) interface{} {
	return &ServerStubInspector{
		service: server,
	}
}

// clientStubInspector implements Inspector.
type clientStubInspector struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubInspector) Ls(ctx _gen_context.T, Glob string, opts ..._gen_ipc.CallOpt) (reply InspectorLsStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Ls", []interface{}{Glob}, opts...); err != nil {
		return
	}
	reply = &implInspectorLsStream{clientCall: call}
	return
}

func (__gen_c *clientStubInspector) LsDetails(ctx _gen_context.T, Glob string, opts ..._gen_ipc.CallOpt) (reply InspectorLsDetailsStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "LsDetails", []interface{}{Glob}, opts...); err != nil {
		return
	}
	reply = &implInspectorLsDetailsStream{clientCall: call}
	return
}

func (__gen_c *clientStubInspector) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubInspector) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubInspector) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubInspector wraps a server that implements
// InspectorService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubInspector struct {
	service InspectorService
}

func (__gen_s *ServerStubInspector) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Ls":
		return []interface{}{}, nil
	case "LsDetails":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubInspector) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Ls"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Glob", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 3,
	}
	result.Methods["LsDetails"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Glob", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 66,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "Size"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Mode"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "ModUnixSecs"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "ModNano"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "IsDir"},
			},
			"veyron/examples/inspector.Details", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubInspector) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubInspector) Ls(call _gen_ipc.ServerCall, Glob string) (err error) {
	stream := &implInspectorServiceLsStream{serverCall: call}
	err = __gen_s.service.Ls(call, Glob, stream)
	return
}

func (__gen_s *ServerStubInspector) LsDetails(call _gen_ipc.ServerCall, Glob string) (err error) {
	stream := &implInspectorServiceLsDetailsStream{serverCall: call}
	err = __gen_s.service.LsDetails(call, Glob, stream)
	return
}
