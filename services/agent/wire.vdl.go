// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: wire.vdl

// Package agent defines an interface to keep a private key in memory, and for
// clients to have access to the private key.
//
// Protocol
//
// The agent starts processes with the VEYRON_AGENT_FD set to one end of a
// unix domain socket. To connect to the agent, a client should create
// a unix domain socket pair. Then send one end of the socket to the agent
// with 1 byte of data. The agent will then serve the Agent service on
// the received socket, using SecurityNone.
//
// The agent also supports an optional mode where it can manage multiple principals.
// Typically this is only used by Device Manager. In this mode, VEYRON_AGENT_FD
// will be 3, and there will be another socket at fd 4.
// Creating a new principal is similar to connecting to to agent: create a socket
// pair and send one end on fd 4 with 1 byte of data.
// Set the data to 1 to request the principal only be stored in memory.
// The agent will create a new principal and respond with a principal handle on fd 4.
// To connect using a previously created principal, create a socket pair and send
// one end with the principal handle as data on fd 4. The agent will not send a
// response on fd 4.
// In either, you can use the normal process to connect to an agent over the
// other end of the pair. Typically you would pass the other end to a child
// process and set VEYRON_AGENT_FD so it knows to connect.
//
// The protocol also has limited support for caching: A client can
// request notification when any other client modifies the principal so it
// can flush the cache. See NotifyWhenChanged for details.
package agent

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"

	// VDL user imports
	"v.io/v23/security"
)

// AgentClientMethods is the client interface
// containing Agent methods.
type AgentClientMethods interface {
	Bless(ctx *context.T, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessSelf(ctx *context.T, name string, caveats []security.Caveat, opts ...rpc.CallOpt) (security.Blessings, error)
	Sign(ctx *context.T, message []byte, opts ...rpc.CallOpt) (security.Signature, error)
	MintDischarge(ctx *context.T, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat, opts ...rpc.CallOpt) (security.Discharge, error)
	PublicKey(*context.T, ...rpc.CallOpt) ([]byte, error)
	BlessingsByName(ctx *context.T, name security.BlessingPattern, opts ...rpc.CallOpt) ([]security.Blessings, error)
	BlessingsInfo(ctx *context.T, blessings security.Blessings, opts ...rpc.CallOpt) (map[string][]security.Caveat, error)
	AddToRoots(ctx *context.T, blessing security.Blessings, opts ...rpc.CallOpt) error
	BlessingStoreSet(ctx *context.T, blessings security.Blessings, forPeers security.BlessingPattern, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreForPeer(ctx *context.T, peerBlessings []string, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreSetDefault(ctx *context.T, blessings security.Blessings, opts ...rpc.CallOpt) error
	BlessingStoreDefault(*context.T, ...rpc.CallOpt) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, ...rpc.CallOpt) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, ...rpc.CallOpt) (string, error)
	BlessingRootsAdd(ctx *context.T, root []byte, pattern security.BlessingPattern, opts ...rpc.CallOpt) error
	BlessingRootsRecognized(ctx *context.T, root []byte, blessing string, opts ...rpc.CallOpt) error
	BlessingRootsDump(*context.T, ...rpc.CallOpt) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, ...rpc.CallOpt) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, ...rpc.CallOpt) (AgentNotifyWhenChangedClientCall, error)
}

// AgentClientStub adds universal methods to AgentClientMethods.
type AgentClientStub interface {
	AgentClientMethods
	rpc.UniversalServiceMethods
}

// AgentClient returns a client stub for Agent.
func AgentClient(name string) AgentClientStub {
	return implAgentClientStub{name}
}

type implAgentClientStub struct {
	name string
}

func (c implAgentClientStub) Bless(ctx *context.T, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Bless", []interface{}{i0, i1, i2, i3, i4}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessSelf(ctx *context.T, i0 string, i1 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessSelf", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) Sign(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 security.Signature, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sign", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) MintDischarge(ctx *context.T, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Discharge, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "MintDischarge", []interface{}{i0, i1, i2}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) PublicKey(ctx *context.T, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "PublicKey", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingsByName(ctx *context.T, i0 security.BlessingPattern, opts ...rpc.CallOpt) (o0 []security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingsByName", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingsInfo(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (o0 map[string][]security.Caveat, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingsInfo", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) AddToRoots(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "AddToRoots", []interface{}{i0}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreSet(ctx *context.T, i0 security.Blessings, i1 security.BlessingPattern, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreSet", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreForPeer(ctx *context.T, i0 []string, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreForPeer", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreSetDefault(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreSetDefault", []interface{}{i0}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreDefault(ctx *context.T, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreDefault", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStorePeerBlessings(ctx *context.T, opts ...rpc.CallOpt) (o0 map[security.BlessingPattern]security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStorePeerBlessings", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreDebugString", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsAdd(ctx *context.T, i0 []byte, i1 security.BlessingPattern, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsAdd", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsRecognized(ctx *context.T, i0 []byte, i1 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsRecognized", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsDump(ctx *context.T, opts ...rpc.CallOpt) (o0 map[security.BlessingPattern][][]byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsDump", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsDebugString", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) NotifyWhenChanged(ctx *context.T, opts ...rpc.CallOpt) (ocall AgentNotifyWhenChangedClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "NotifyWhenChanged", nil, opts...); err != nil {
		return
	}
	ocall = &implAgentNotifyWhenChangedClientCall{ClientCall: call}
	return
}

// AgentNotifyWhenChangedClientStream is the client stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientStream interface {
	// RecvStream returns the receiver side of the Agent.NotifyWhenChanged client stream.
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

// AgentNotifyWhenChangedClientCall represents the call returned from Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientCall interface {
	AgentNotifyWhenChangedClientStream
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

type implAgentNotifyWhenChangedClientCall struct {
	rpc.ClientCall
	valRecv bool
	errRecv error
}

func (c *implAgentNotifyWhenChangedClientCall) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implAgentNotifyWhenChangedClientCallRecv{c}
}

type implAgentNotifyWhenChangedClientCallRecv struct {
	c *implAgentNotifyWhenChangedClientCall
}

func (c implAgentNotifyWhenChangedClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implAgentNotifyWhenChangedClientCallRecv) Value() bool {
	return c.c.valRecv
}
func (c implAgentNotifyWhenChangedClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implAgentNotifyWhenChangedClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// AgentServerMethods is the interface a server writer
// implements for Agent.
type AgentServerMethods interface {
	Bless(ctx *context.T, call rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(ctx *context.T, call rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(ctx *context.T, call rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(ctx *context.T, call rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(*context.T, rpc.ServerCall) ([]byte, error)
	BlessingsByName(ctx *context.T, call rpc.ServerCall, name security.BlessingPattern) ([]security.Blessings, error)
	BlessingsInfo(ctx *context.T, call rpc.ServerCall, blessings security.Blessings) (map[string][]security.Caveat, error)
	AddToRoots(ctx *context.T, call rpc.ServerCall, blessing security.Blessings) error
	BlessingStoreSet(ctx *context.T, call rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(ctx *context.T, call rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(ctx *context.T, call rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(*context.T, rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, rpc.ServerCall) (string, error)
	BlessingRootsAdd(ctx *context.T, call rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(ctx *context.T, call rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDump(*context.T, rpc.ServerCall) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, AgentNotifyWhenChangedServerCall) error
}

// AgentServerStubMethods is the server interface containing
// Agent methods, as expected by rpc.Server.
// The only difference between this interface and AgentServerMethods
// is the streaming methods.
type AgentServerStubMethods interface {
	Bless(ctx *context.T, call rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(ctx *context.T, call rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(ctx *context.T, call rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(ctx *context.T, call rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(*context.T, rpc.ServerCall) ([]byte, error)
	BlessingsByName(ctx *context.T, call rpc.ServerCall, name security.BlessingPattern) ([]security.Blessings, error)
	BlessingsInfo(ctx *context.T, call rpc.ServerCall, blessings security.Blessings) (map[string][]security.Caveat, error)
	AddToRoots(ctx *context.T, call rpc.ServerCall, blessing security.Blessings) error
	BlessingStoreSet(ctx *context.T, call rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(ctx *context.T, call rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(ctx *context.T, call rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(*context.T, rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, rpc.ServerCall) (string, error)
	BlessingRootsAdd(ctx *context.T, call rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(ctx *context.T, call rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDump(*context.T, rpc.ServerCall) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, *AgentNotifyWhenChangedServerCallStub) error
}

// AgentServerStub adds universal methods to AgentServerStubMethods.
type AgentServerStub interface {
	AgentServerStubMethods
	// Describe the Agent interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AgentServer returns a server stub for Agent.
// It converts an implementation of AgentServerMethods into
// an object that may be used by rpc.Server.
func AgentServer(impl AgentServerMethods) AgentServerStub {
	stub := implAgentServerStub{
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

type implAgentServerStub struct {
	impl AgentServerMethods
	gs   *rpc.GlobState
}

func (s implAgentServerStub) Bless(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat) (security.Blessings, error) {
	return s.impl.Bless(ctx, call, i0, i1, i2, i3, i4)
}

func (s implAgentServerStub) BlessSelf(ctx *context.T, call rpc.ServerCall, i0 string, i1 []security.Caveat) (security.Blessings, error) {
	return s.impl.BlessSelf(ctx, call, i0, i1)
}

func (s implAgentServerStub) Sign(ctx *context.T, call rpc.ServerCall, i0 []byte) (security.Signature, error) {
	return s.impl.Sign(ctx, call, i0)
}

func (s implAgentServerStub) MintDischarge(ctx *context.T, call rpc.ServerCall, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat) (security.Discharge, error) {
	return s.impl.MintDischarge(ctx, call, i0, i1, i2)
}

func (s implAgentServerStub) PublicKey(ctx *context.T, call rpc.ServerCall) ([]byte, error) {
	return s.impl.PublicKey(ctx, call)
}

func (s implAgentServerStub) BlessingsByName(ctx *context.T, call rpc.ServerCall, i0 security.BlessingPattern) ([]security.Blessings, error) {
	return s.impl.BlessingsByName(ctx, call, i0)
}

func (s implAgentServerStub) BlessingsInfo(ctx *context.T, call rpc.ServerCall, i0 security.Blessings) (map[string][]security.Caveat, error) {
	return s.impl.BlessingsInfo(ctx, call, i0)
}

func (s implAgentServerStub) AddToRoots(ctx *context.T, call rpc.ServerCall, i0 security.Blessings) error {
	return s.impl.AddToRoots(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreSet(ctx *context.T, call rpc.ServerCall, i0 security.Blessings, i1 security.BlessingPattern) (security.Blessings, error) {
	return s.impl.BlessingStoreSet(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingStoreForPeer(ctx *context.T, call rpc.ServerCall, i0 []string) (security.Blessings, error) {
	return s.impl.BlessingStoreForPeer(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreSetDefault(ctx *context.T, call rpc.ServerCall, i0 security.Blessings) error {
	return s.impl.BlessingStoreSetDefault(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreDefault(ctx *context.T, call rpc.ServerCall) (security.Blessings, error) {
	return s.impl.BlessingStoreDefault(ctx, call)
}

func (s implAgentServerStub) BlessingStorePeerBlessings(ctx *context.T, call rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error) {
	return s.impl.BlessingStorePeerBlessings(ctx, call)
}

func (s implAgentServerStub) BlessingStoreDebugString(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.BlessingStoreDebugString(ctx, call)
}

func (s implAgentServerStub) BlessingRootsAdd(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 security.BlessingPattern) error {
	return s.impl.BlessingRootsAdd(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsRecognized(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 string) error {
	return s.impl.BlessingRootsRecognized(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsDump(ctx *context.T, call rpc.ServerCall) (map[security.BlessingPattern][][]byte, error) {
	return s.impl.BlessingRootsDump(ctx, call)
}

func (s implAgentServerStub) BlessingRootsDebugString(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.BlessingRootsDebugString(ctx, call)
}

func (s implAgentServerStub) NotifyWhenChanged(ctx *context.T, call *AgentNotifyWhenChangedServerCallStub) error {
	return s.impl.NotifyWhenChanged(ctx, call)
}

func (s implAgentServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAgentServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AgentDesc}
}

// AgentDesc describes the Agent interface.
var AgentDesc rpc.InterfaceDesc = descAgent

// descAgent hides the desc to keep godoc clean.
var descAgent = rpc.InterfaceDesc{
	Name:    "Agent",
	PkgPath: "v.io/x/ref/services/agent",
	Methods: []rpc.MethodDesc{
		{
			Name: "Bless",
			InArgs: []rpc.ArgDesc{
				{"key", ``},               // []byte
				{"wit", ``},               // security.Blessings
				{"extension", ``},         // string
				{"caveat", ``},            // security.Caveat
				{"additionalCaveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessSelf",
			InArgs: []rpc.ArgDesc{
				{"name", ``},    // string
				{"caveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "Sign",
			InArgs: []rpc.ArgDesc{
				{"message", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Signature
			},
		},
		{
			Name: "MintDischarge",
			InArgs: []rpc.ArgDesc{
				{"forCaveat", ``},                    // security.Caveat
				{"caveatOnDischarge", ``},            // security.Caveat
				{"additionalCaveatsOnDischarge", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Discharge
			},
		},
		{
			Name: "PublicKey",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
		},
		{
			Name: "BlessingsByName",
			InArgs: []rpc.ArgDesc{
				{"name", ``}, // security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []security.Blessings
			},
		},
		{
			Name: "BlessingsInfo",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[string][]security.Caveat
			},
		},
		{
			Name: "AddToRoots",
			InArgs: []rpc.ArgDesc{
				{"blessing", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreSet",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
				{"forPeers", ``},  // security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreForPeer",
			InArgs: []rpc.ArgDesc{
				{"peerBlessings", ``}, // []string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreSetDefault",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreDefault",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStorePeerBlessings",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[security.BlessingPattern]security.Blessings
			},
		},
		{
			Name: "BlessingStoreDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "BlessingRootsAdd",
			InArgs: []rpc.ArgDesc{
				{"root", ``},    // []byte
				{"pattern", ``}, // security.BlessingPattern
			},
		},
		{
			Name: "BlessingRootsRecognized",
			InArgs: []rpc.ArgDesc{
				{"root", ``},     // []byte
				{"blessing", ``}, // string
			},
		},
		{
			Name: "BlessingRootsDump",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[security.BlessingPattern][][]byte
			},
		},
		{
			Name: "BlessingRootsDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "NotifyWhenChanged",
			Doc:  "// Clients using caching should call NotifyWhenChanged upon connecting to\n// the server. The server will stream back values whenever the client should\n// flush the cache. The streamed value is arbitrary, simply flush whenever\n// recieving a new item.",
		},
	},
}

// AgentNotifyWhenChangedServerStream is the server stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerStream interface {
	// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item bool) error
	}
}

// AgentNotifyWhenChangedServerCall represents the context passed to Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerCall interface {
	rpc.ServerCall
	AgentNotifyWhenChangedServerStream
}

// AgentNotifyWhenChangedServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements AgentNotifyWhenChangedServerCall.
type AgentNotifyWhenChangedServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes AgentNotifyWhenChangedServerCallStub from rpc.StreamServerCall.
func (s *AgentNotifyWhenChangedServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
func (s *AgentNotifyWhenChangedServerCallStub) SendStream() interface {
	Send(item bool) error
} {
	return implAgentNotifyWhenChangedServerCallSend{s}
}

type implAgentNotifyWhenChangedServerCallSend struct {
	s *AgentNotifyWhenChangedServerCallStub
}

func (s implAgentNotifyWhenChangedServerCallSend) Send(item bool) error {
	return s.s.Send(item)
}
