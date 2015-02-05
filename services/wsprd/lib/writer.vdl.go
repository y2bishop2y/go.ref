// This file was auto-generated by the veyron vdl tool.
// Source: writer.vdl

package lib

import (
	// VDL system imports
	"v.io/core/veyron2/vdl"
)

// The response from the javascript server to the proxy.
type ServerRPCReply struct {
	Results []vdl.AnyRep
	Err     error
}

func (ServerRPCReply) __VDLReflect(struct {
	Name string "v.io/wspr/veyron/services/wsprd/lib.ServerRPCReply"
}) {
}

func init() {
	vdl.Register(ServerRPCReply{})
}
