// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: protocol.vdl

package proxy

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
)

// Request is the message sent by a server to request that the proxy route
// traffic intended for the server's RoutingId to the network connection
// between the server and the proxy.
type Request struct {
	// Blessings of the server that wishes to be proxied.
	// Used to authorize the use of the proxy.
	Blessings security.Blessings
	// Discharges required to make Blessings valid.
	Discharges []security.Discharge
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/profiles/internal/rpc/stream/proxy.Request"`
}) {
}

// Response is sent by the proxy to the server after processing Request.
type Response struct {
	// Error is a description of why the proxy refused to proxy the server.
	// A nil error indicates that the proxy will route traffic to the server.
	Error error
	// Endpoint is the string representation of an endpoint that can be
	// used to communicate with the server through the proxy.
	Endpoint string
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/profiles/internal/rpc/stream/proxy.Response"`
}) {
}

func init() {
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))
}
