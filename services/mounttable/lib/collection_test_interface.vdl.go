// This file was auto-generated by the veyron vdl tool.
// Source: collection_test_interface.vdl

package mounttable

import (
	// VDL system imports
	"v.io/core/veyron2"
	"v.io/core/veyron2/context"
	"v.io/core/veyron2/ipc"
)

// CollectionClientMethods is the client interface
// containing Collection methods.
type CollectionClientMethods interface {
	// Export sets the value for a name.  Overwrite controls the behavior when
	// an entry exists, if Overwrite is true, then the binding is replaced,
	// otherwise the call fails with an error.  The Val must be no larger than
	// MaxSize bytes.
	Export(ctx *context.T, Val string, Overwrite bool, opts ...ipc.CallOpt) error
	// Lookup retrieves the value associated with a name.  Returns an error if
	// there is no such binding.
	Lookup(*context.T, ...ipc.CallOpt) ([]byte, error)
}

// CollectionClientStub adds universal methods to CollectionClientMethods.
type CollectionClientStub interface {
	CollectionClientMethods
	ipc.UniversalServiceMethods
}

// CollectionClient returns a client stub for Collection.
func CollectionClient(name string, opts ...ipc.BindOpt) CollectionClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implCollectionClientStub{name, client}
}

type implCollectionClientStub struct {
	name   string
	client ipc.Client
}

func (c implCollectionClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return veyron2.GetClient(ctx)
}

func (c implCollectionClientStub) Export(ctx *context.T, i0 string, i1 bool, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Export", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implCollectionClientStub) Lookup(ctx *context.T, opts ...ipc.CallOpt) (o0 []byte, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Lookup", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// CollectionServerMethods is the interface a server writer
// implements for Collection.
type CollectionServerMethods interface {
	// Export sets the value for a name.  Overwrite controls the behavior when
	// an entry exists, if Overwrite is true, then the binding is replaced,
	// otherwise the call fails with an error.  The Val must be no larger than
	// MaxSize bytes.
	Export(ctx ipc.ServerContext, Val string, Overwrite bool) error
	// Lookup retrieves the value associated with a name.  Returns an error if
	// there is no such binding.
	Lookup(ipc.ServerContext) ([]byte, error)
}

// CollectionServerStubMethods is the server interface containing
// Collection methods, as expected by ipc.Server.
// There is no difference between this interface and CollectionServerMethods
// since there are no streaming methods.
type CollectionServerStubMethods CollectionServerMethods

// CollectionServerStub adds universal methods to CollectionServerStubMethods.
type CollectionServerStub interface {
	CollectionServerStubMethods
	// Describe the Collection interfaces.
	Describe__() []ipc.InterfaceDesc
}

// CollectionServer returns a server stub for Collection.
// It converts an implementation of CollectionServerMethods into
// an object that may be used by ipc.Server.
func CollectionServer(impl CollectionServerMethods) CollectionServerStub {
	stub := implCollectionServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implCollectionServerStub struct {
	impl CollectionServerMethods
	gs   *ipc.GlobState
}

func (s implCollectionServerStub) Export(ctx ipc.ServerContext, i0 string, i1 bool) error {
	return s.impl.Export(ctx, i0, i1)
}

func (s implCollectionServerStub) Lookup(ctx ipc.ServerContext) ([]byte, error) {
	return s.impl.Lookup(ctx)
}

func (s implCollectionServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implCollectionServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{CollectionDesc}
}

// CollectionDesc describes the Collection interface.
var CollectionDesc ipc.InterfaceDesc = descCollection

// descCollection hides the desc to keep godoc clean.
var descCollection = ipc.InterfaceDesc{
	Name:    "Collection",
	PkgPath: "v.io/core/veyron/services/mounttable/lib",
	Methods: []ipc.MethodDesc{
		{
			Name: "Export",
			Doc:  "// Export sets the value for a name.  Overwrite controls the behavior when\n// an entry exists, if Overwrite is true, then the binding is replaced,\n// otherwise the call fails with an error.  The Val must be no larger than\n// MaxSize bytes.",
			InArgs: []ipc.ArgDesc{
				{"Val", ``},       // string
				{"Overwrite", ``}, // bool
			},
		},
		{
			Name: "Lookup",
			Doc:  "// Lookup retrieves the value associated with a name.  Returns an error if\n// there is no such binding.",
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []byte
			},
		},
	},
}
