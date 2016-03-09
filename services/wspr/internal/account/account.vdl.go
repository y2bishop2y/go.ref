// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: account.vdl

package account

import (
	"v.io/v23/vdl"
)

// Caveat describes a restriction on the validity of a blessing/discharge.
// TODO Remove this
type Caveat struct {
	Type string
	Args string
}

func (Caveat) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/account.Caveat"`
}) {
}

func (m *Caveat) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_account_v_io_x_ref_services_wspr_internal_account_Caveat == nil || __VDLTypeaccount0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Type")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Type), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Args")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Args), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Caveat) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*Caveat)(nil))
}

var __VDLTypeaccount0 *vdl.Type = vdl.TypeOf((*Caveat)(nil))
var __VDLType_account_v_io_x_ref_services_wspr_internal_account_Caveat *vdl.Type = vdl.TypeOf(Caveat{})

func __VDLEnsureNativeBuilt_account() {
}
