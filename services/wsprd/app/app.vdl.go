// This file was auto-generated by the vanadium vdl tool.
// Source: app.vdl

// The app package contains the struct that keeps per javascript app state and handles translating
// javascript requests to veyron requests and vice versa.
package app

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/vdlroot/time"
	"v.io/v23/vtrace"
)

type VeyronRPCRequest struct {
	Name        string
	Method      string
	NumInArgs   int32
	NumOutArgs  int32
	IsStreaming bool
	// TODO(bjornick): Change Timeout to use time.WireDeadline instead.
	Deadline     time.Deadline
	TraceRequest vtrace.Request
}

func (VeyronRPCRequest) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/app.VeyronRPCRequest"
}) {
}

type VeyronRPCResponse struct {
	OutArgs       []*vdl.Value
	TraceResponse vtrace.Response
}

func (VeyronRPCResponse) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/app.VeyronRPCResponse"
}) {
}

func init() {
	vdl.Register((*VeyronRPCRequest)(nil))
	vdl.Register((*VeyronRPCResponse)(nil))
}
