// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: localblobstore

package localblobstore

import (
	"fmt"
	"time"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/x/ref/services/syncbase/server/interfaces"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// A BlobMetadata describes information that syncbase stores for a blob it
// holds, independent of the actual content.  Compare with a Signpost, which
// may be stored for a blob that the current device does not hold (and
// indicates where it may be found).  (See
// v.io/x/ref/services/syncbase/server/interfaces/sync_types.vdl for the
// Signpost definition.)
type BlobMetadata struct {
	OwnerShares interfaces.BlobSharesBySyncgroup // >0 for any group => syncbase must keep blob.
	Referenced  time.Time                        // When structured-store reference to blob last seen.
	Accessed    time.Time                        // Last attempted access.
}

func (BlobMetadata) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/localblobstore.BlobMetadata"`
}) {
}

func (m *BlobMetadata) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var var4 bool
	if len(m.OwnerShares) == 0 {
		var4 = true
	}
	if var4 {
		if err := fieldsTarget1.ZeroField("OwnerShares"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("OwnerShares")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.OwnerShares.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var wireValue5 vdltime.Time
	if err := vdltime.TimeFromNative(&wireValue5, m.Referenced); err != nil {
		return err
	}

	var8 := (wireValue5 == vdltime.Time{})
	if var8 {
		if err := fieldsTarget1.ZeroField("Referenced"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Referenced")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue5.FillVDLTarget(fieldTarget7, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var wireValue9 vdltime.Time
	if err := vdltime.TimeFromNative(&wireValue9, m.Accessed); err != nil {
		return err
	}

	var12 := (wireValue9 == vdltime.Time{})
	if var12 {
		if err := fieldsTarget1.ZeroField("Accessed"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("Accessed")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue9.FillVDLTarget(fieldTarget11, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobMetadata) MakeVDLTarget() vdl.Target {
	return &BlobMetadataTarget{Value: m}
}

type BlobMetadataTarget struct {
	Value             *BlobMetadata
	ownerSharesTarget interfaces.BlobSharesBySyncgroupTarget
	referencedTarget  vdltime.TimeTarget
	accessedTarget    vdltime.TimeTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobMetadataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobMetadata)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobMetadataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "OwnerShares":
		t.ownerSharesTarget.Value = &t.Value.OwnerShares
		target, err := &t.ownerSharesTarget, error(nil)
		return nil, target, err
	case "Referenced":
		t.referencedTarget.Value = &t.Value.Referenced
		target, err := &t.referencedTarget, error(nil)
		return nil, target, err
	case "Accessed":
		t.accessedTarget.Value = &t.Value.Accessed
		target, err := &t.accessedTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *BlobMetadataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobMetadataTarget) ZeroField(name string) error {
	switch name {
	case "OwnerShares":
		t.Value.OwnerShares = interfaces.BlobSharesBySyncgroup(nil)
		return nil
	case "Referenced":
		t.Value.Referenced = time.Time{}
		return nil
	case "Accessed":
		t.Value.Accessed = time.Time{}
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *BlobMetadataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x BlobMetadata) VDLIsZero() bool {
	if len(x.OwnerShares) != 0 {
		return false
	}
	if !x.Referenced.IsZero() {
		return false
	}
	if !x.Accessed.IsZero() {
		return false
	}
	return true
}

func (x BlobMetadata) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobMetadata)(nil)).Elem()); err != nil {
		return err
	}
	if len(x.OwnerShares) != 0 {
		if err := enc.NextField("OwnerShares"); err != nil {
			return err
		}
		if err := x.OwnerShares.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Referenced.IsZero() {
		if err := enc.NextField("Referenced"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.Referenced); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Accessed.IsZero() {
		if err := enc.NextField("Accessed"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.Accessed); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobMetadata) VDLRead(dec vdl.Decoder) error {
	*x = BlobMetadata{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "OwnerShares":
			if err := x.OwnerShares.VDLRead(dec); err != nil {
				return err
			}
		case "Referenced":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.Referenced); err != nil {
				return err
			}
		case "Accessed":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.Accessed); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// A PerSyncgroup is blob-related data stored per syncgroup.
// It includes information that helps syncgroup members decide whether
// a peer makes a better or worse owner of a blob.
type PerSyncgroup struct {
	Priority interfaces.SgPriority
}

func (PerSyncgroup) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/localblobstore.PerSyncgroup"`
}) {
}

func (m *PerSyncgroup) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Priority == interfaces.SgPriority{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Priority"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Priority")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Priority.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *PerSyncgroup) MakeVDLTarget() vdl.Target {
	return &PerSyncgroupTarget{Value: m}
}

type PerSyncgroupTarget struct {
	Value          *PerSyncgroup
	priorityTarget interfaces.SgPriorityTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *PerSyncgroupTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*PerSyncgroup)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *PerSyncgroupTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Priority":
		t.priorityTarget.Value = &t.Value.Priority
		target, err := &t.priorityTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *PerSyncgroupTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *PerSyncgroupTarget) ZeroField(name string) error {
	switch name {
	case "Priority":
		t.Value.Priority = interfaces.SgPriority{}
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *PerSyncgroupTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x PerSyncgroup) VDLIsZero() bool {
	if !x.Priority.VDLIsZero() {
		return false
	}
	return true
}

func (x PerSyncgroup) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*PerSyncgroup)(nil)).Elem()); err != nil {
		return err
	}
	if !x.Priority.VDLIsZero() {
		if err := enc.NextField("Priority"); err != nil {
			return err
		}
		if err := x.Priority.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PerSyncgroup) VDLRead(dec vdl.Decoder) error {
	*x = PerSyncgroup{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Priority":
			if err := x.Priority.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*BlobMetadata)(nil))
	vdl.Register((*PerSyncgroup)(nil))

	return struct{}{}
}
