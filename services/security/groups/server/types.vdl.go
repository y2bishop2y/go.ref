// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

package server

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/services/security/access"
	"v.io/v23/services/security/groups"
)

// groupData represents the persistent state of a group. (The group name is
// persisted as the store entry key.)
type groupData struct {
	ACL     access.TaggedACLMap
	Entries map[groups.BlessingPatternChunk]struct{}
}

func (groupData) __VDLReflect(struct {
	Name string "v.io/core/veyron/services/security/groups/server.groupData"
}) {
}

func init() {
	vdl.Register((*groupData)(nil))
}
