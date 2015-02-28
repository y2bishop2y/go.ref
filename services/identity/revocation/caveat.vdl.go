// This file was auto-generated by the veyron vdl tool.
// Source: caveat.vdl

package revocation

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
	"v.io/v23/uniqueid"
)

// NotRevokedCaveat is used to implement revocation.
// It validates iff the parameter is not included in a list of blacklisted
// values.
//
// The third-party discharging service checks this revocation caveat against a
// database of blacklisted (revoked) keys before issuing a discharge.
var NotRevokedCaveat = security.CaveatDescriptor{
	Id: uniqueid.Id{
		75,
		70,
		92,
		86,
		55,
		121,
		209,
		59,
		123,
		163,
		167,
		214,
		165,
		52,
		128,
		0,
	},
	ParamType: vdl.TypeOf([]byte(nil)),
}
