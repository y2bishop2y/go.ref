// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: blobtestsvdl

package blobtestsvdl

import (
	"fmt"
	"reflect"
	"v.io/v23/services/syncbase"
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type BlobInfo struct {
	Info string
	Br   syncbase.BlobRef
}

func (BlobInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobInfo"`
}) {
}

func (m *BlobInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Info == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Info"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Br == syncbase.BlobRef(""))
	if var7 {
		if err := fieldsTarget1.ZeroField("Br"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Br")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Br.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobInfo) MakeVDLTarget() vdl.Target {
	return &BlobInfoTarget{Value: m}
}

type BlobInfoTarget struct {
	Value      *BlobInfo
	infoTarget vdl.StringTarget
	brTarget   syncbase.BlobRefTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobInfo)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Br":
		t.brTarget.Value = &t.Value.Br
		target, err := &t.brTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *BlobInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobInfoTarget) ZeroField(name string) error {
	switch name {
	case "Info":
		t.Value.Info = ""
		return nil
	case "Br":
		t.Value.Br = syncbase.BlobRef("")
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *BlobInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x BlobInfo) VDLIsZero() bool {
	return x == BlobInfo{}
}

func (x BlobInfo) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobInfo)(nil)).Elem()); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Br != "" {
		if err := enc.NextField("Br"); err != nil {
			return err
		}
		if err := x.Br.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobInfo) VDLRead(dec vdl.Decoder) error {
	*x = BlobInfo{}
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
		case "Info":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Br":
			if err := x.Br.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

type (
	// BlobUnion represents any single field of the BlobUnion union type.
	BlobUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlobUnion union type.
		__VDLReflect(__BlobUnionReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
		VDLIsZero() bool
		VDLWrite(vdl.Encoder) error
	}
	// BlobUnionNum represents field Num of the BlobUnion union type.
	BlobUnionNum struct{ Value int32 }
	// BlobUnionBi represents field Bi of the BlobUnion union type.
	BlobUnionBi struct{ Value BlobInfo }
	// __BlobUnionReflect describes the BlobUnion union type.
	__BlobUnionReflect struct {
		Name               string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobUnion"`
		Type               BlobUnion
		UnionTargetFactory blobUnionTargetFactory
		Union              struct {
			Num BlobUnionNum
			Bi  BlobUnionBi
		}
	}
)

func (x BlobUnionNum) Index() int                      { return 0 }
func (x BlobUnionNum) Interface() interface{}          { return x.Value }
func (x BlobUnionNum) Name() string                    { return "Num" }
func (x BlobUnionNum) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionNum) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Num")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromInt(int64(m.Value), tt.NonOptional().Field(0).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionNum) MakeVDLTarget() vdl.Target {
	return nil
}

func (x BlobUnionBi) Index() int                      { return 1 }
func (x BlobUnionBi) Interface() interface{}          { return x.Value }
func (x BlobUnionBi) Name() string                    { return "Bi" }
func (x BlobUnionBi) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionBi) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Bi")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(1).Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionBi) MakeVDLTarget() vdl.Target {
	return nil
}

type BlobUnionTarget struct {
	Value     *BlobUnion
	fieldName string

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobUnionTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if ttWant := vdl.TypeOf((*BlobUnion)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}

	return t, nil
}
func (t *BlobUnionTarget) StartField(name string) (key, field vdl.Target, _ error) {
	t.fieldName = name
	switch name {
	case "Num":
		val := int32(0)
		return nil, &vdl.Int32Target{Value: &val}, nil
	case "Bi":
		val := BlobInfo{}
		return nil, &BlobInfoTarget{Value: &val}, nil
	default:
		return nil, nil, fmt.Errorf("field %s not in union v.io/x/ref/services/syncbase/vsync/testdata.BlobUnion", name)
	}
}
func (t *BlobUnionTarget) FinishField(_, fieldTarget vdl.Target) error {
	switch t.fieldName {
	case "Num":
		*t.Value = BlobUnionNum{*(fieldTarget.(*vdl.Int32Target)).Value}
	case "Bi":
		*t.Value = BlobUnionBi{*(fieldTarget.(*BlobInfoTarget)).Value}
	}
	return nil
}
func (t *BlobUnionTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type blobUnionTargetFactory struct{}

func (t blobUnionTargetFactory) VDLMakeUnionTarget(union interface{}) (vdl.Target, error) {
	if typedUnion, ok := union.(*BlobUnion); ok {
		return &BlobUnionTarget{Value: typedUnion}, nil
	}
	return nil, fmt.Errorf("got %T, want *BlobUnion", union)
}

func (x BlobUnionNum) VDLIsZero() bool {
	return x.Value == 0
}

func (x BlobUnionBi) VDLIsZero() bool {
	return false
}

func (x BlobUnionNum) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobUnion)(nil))); err != nil {
		return err
	}
	if err := enc.NextField("Num"); err != nil {
		return err
	}
	if err := enc.StartValue(vdl.Int32Type); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x.Value)); err != nil {
		return err
	}
	if err := enc.FinishValue(); err != nil {
		return err
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x BlobUnionBi) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobUnion)(nil))); err != nil {
		return err
	}
	if err := enc.NextField("Bi"); err != nil {
		return err
	}
	if err := x.Value.VDLWrite(enc); err != nil {
		return err
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func VDLReadBlobUnion(dec vdl.Decoder, x *BlobUnion) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(x), dec.Type()) {
		return fmt.Errorf("incompatible union %T, from %v", x, dec.Type())
	}
	f, err := dec.NextField()
	if err != nil {
		return err
	}
	switch f {
	case "Num":
		var field BlobUnionNum
		if err := dec.StartValue(); err != nil {
			return err
		}
		tmp, err := dec.DecodeInt(32)
		if err != nil {
			return err
		}
		field.Value = int32(tmp)
		if err := dec.FinishValue(); err != nil {
			return err
		}
		*x = field
	case "Bi":
		var field BlobUnionBi
		if err := field.Value.VDLRead(dec); err != nil {
			return err
		}
		*x = field
	case "":
		return fmt.Errorf("missing field in union %T, from %v", x, dec.Type())
	default:
		return fmt.Errorf("field %q not in union %T, from %v", f, x, dec.Type())
	}
	switch f, err := dec.NextField(); {
	case err != nil:
		return err
	case f != "":
		return fmt.Errorf("extra field %q in union %T, from %v", f, x, dec.Type())
	}
	return dec.FinishValue()
}

type BlobSet struct {
	Info string
	Bs   map[syncbase.BlobRef]struct{}
}

func (BlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobSet"`
}) {
}

func (m *BlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Info == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Info"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Bs) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Bs"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Bs")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			setTarget8, err := fieldTarget6.StartSet(tt.NonOptional().Field(1).Type, len(m.Bs))
			if err != nil {
				return err
			}
			for key10 := range m.Bs {
				keyTarget9, err := setTarget8.StartKey()
				if err != nil {
					return err
				}

				if err := key10.FillVDLTarget(keyTarget9, tt.NonOptional().Field(1).Type.Key()); err != nil {
					return err
				}
				if err := setTarget8.FinishKey(keyTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishSet(setTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobSet) MakeVDLTarget() vdl.Target {
	return &BlobSetTarget{Value: m}
}

type BlobSetTarget struct {
	Value      *BlobSet
	infoTarget vdl.StringTarget
	bsTarget   __VDLTarget1_set
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobSet)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bs":
		t.bsTarget.Value = &t.Value.Bs
		target, err := &t.bsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *BlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobSetTarget) ZeroField(name string) error {
	switch name {
	case "Info":
		t.Value.Info = ""
		return nil
	case "Bs":
		t.Value.Bs = map[syncbase.BlobRef]struct{}(nil)
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *BlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[syncbase.BlobRef]struct{}
type __VDLTarget1_set struct {
	Value     *map[syncbase.BlobRef]struct{}
	currKey   syncbase.BlobRef
	keyTarget syncbase.BlobRefTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *__VDLTarget1_set) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {

	if ttWant := vdl.TypeOf((*map[syncbase.BlobRef]struct{})(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[syncbase.BlobRef]struct{})
	return t, nil
}
func (t *__VDLTarget1_set) StartKey() (key vdl.Target, _ error) {
	t.currKey = syncbase.BlobRef("")
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_set) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *__VDLTarget1_set) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

func (x BlobSet) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.Bs) != 0 {
		return false
	}
	return true
}

func (x BlobSet) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobSet)(nil)).Elem()); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Bs) != 0 {
		if err := enc.NextField("Bs"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_set_1(enc, x.Bs); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_set_1(enc vdl.Encoder, x map[syncbase.BlobRef]struct{}) error {
	if err := enc.StartValue(vdl.TypeOf((*map[syncbase.BlobRef]struct{})(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := key.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobSet) VDLRead(dec vdl.Decoder) error {
	*x = BlobSet{}
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
		case "Info":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Bs":
			if err := __VDLReadAnon_set_1(dec, &x.Bs); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_set_1(dec vdl.Decoder, x *map[syncbase.BlobRef]struct{}) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible set %T, from %v", *x, dec.Type())
	}
	var tmpMap map[syncbase.BlobRef]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[syncbase.BlobRef]struct{}, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key syncbase.BlobRef
		{
			if err := key.VDLRead(dec); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[syncbase.BlobRef]struct{})
		}
		tmpMap[key] = struct{}{}
	}
}

type BlobAny struct {
	Info string
	Baa  []*vom.RawBytes
}

func (BlobAny) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobAny"`
}) {
}

func (m *BlobAny) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Info == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Info"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Baa) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Baa"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Baa")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget8, err := fieldTarget6.StartList(tt.NonOptional().Field(1).Type, len(m.Baa))
			if err != nil {
				return err
			}
			for i, elem10 := range m.Baa {
				elemTarget9, err := listTarget8.StartElem(i)
				if err != nil {
					return err
				}

				if err := elem10.FillVDLTarget(elemTarget9, tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget8.FinishElem(elemTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishList(listTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobAny) MakeVDLTarget() vdl.Target {
	return &BlobAnyTarget{Value: m}
}

type BlobAnyTarget struct {
	Value      *BlobAny
	infoTarget vdl.StringTarget
	baaTarget  __VDLTarget2_list
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobAnyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobAny)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobAnyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Baa":
		t.baaTarget.Value = &t.Value.Baa
		target, err := &t.baaTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *BlobAnyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobAnyTarget) ZeroField(name string) error {
	switch name {
	case "Info":
		t.Value.Info = ""
		return nil
	case "Baa":
		t.Value.Baa = []*vom.RawBytes(nil)
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *BlobAnyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []*vom.RawBytes
type __VDLTarget2_list struct {
	Value *[]*vom.RawBytes

	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget2_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[]*vom.RawBytes)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*vom.RawBytes, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget2_list) StartElem(index int) (elem vdl.Target, _ error) {
	target, err := vdl.ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
	return target, err
}
func (t *__VDLTarget2_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget2_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

func (x BlobAny) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.Baa) != 0 {
		return false
	}
	return true
}

func (x BlobAny) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobAny)(nil)).Elem()); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Baa) != 0 {
		if err := enc.NextField("Baa"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.Baa); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []*vom.RawBytes) error {
	if err := enc.StartValue(vdl.TypeOf((*[]*vom.RawBytes)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if x[i] == nil {
			if err := enc.NilValue(vdl.AnyType); err != nil {
				return err
			}
		} else {
			if err := x[i].VDLWrite(enc); err != nil {
				return err
			}
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobAny) VDLRead(dec vdl.Decoder) error {
	*x = BlobAny{}
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
		case "Info":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Baa":
			if err := __VDLReadAnon_list_2(dec, &x.Baa); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]*vom.RawBytes) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]*vom.RawBytes, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem *vom.RawBytes
		elem = new(vom.RawBytes)
		if err := elem.VDLRead(dec); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

type NonBlobSet struct {
	Info string
	S    map[string]struct{}
}

func (NonBlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.NonBlobSet"`
}) {
}

func (m *NonBlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Info == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Info"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.S) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("S"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("S")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			setTarget8, err := fieldTarget6.StartSet(tt.NonOptional().Field(1).Type, len(m.S))
			if err != nil {
				return err
			}
			for key10 := range m.S {
				keyTarget9, err := setTarget8.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget9.FromString(string(key10), tt.NonOptional().Field(1).Type.Key()); err != nil {
					return err
				}
				if err := setTarget8.FinishKey(keyTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishSet(setTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *NonBlobSet) MakeVDLTarget() vdl.Target {
	return &NonBlobSetTarget{Value: m}
}

type NonBlobSetTarget struct {
	Value      *NonBlobSet
	infoTarget vdl.StringTarget
	sTarget    __VDLTarget3_set
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *NonBlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*NonBlobSet)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *NonBlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "S":
		t.sTarget.Value = &t.Value.S
		target, err := &t.sTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *NonBlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *NonBlobSetTarget) ZeroField(name string) error {
	switch name {
	case "Info":
		t.Value.Info = ""
		return nil
	case "S":
		t.Value.S = map[string]struct{}(nil)
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *NonBlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string]struct{}
type __VDLTarget3_set struct {
	Value     *map[string]struct{}
	currKey   string
	keyTarget vdl.StringTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *__VDLTarget3_set) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {

	if ttWant := vdl.TypeOf((*map[string]struct{})(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string]struct{})
	return t, nil
}
func (t *__VDLTarget3_set) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget3_set) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *__VDLTarget3_set) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

func (x NonBlobSet) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.S) != 0 {
		return false
	}
	return true
}

func (x NonBlobSet) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*NonBlobSet)(nil)).Elem()); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.S) != 0 {
		if err := enc.NextField("S"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_set_3(enc, x.S); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_set_3(enc vdl.Encoder, x map[string]struct{}) error {
	if err := enc.StartValue(vdl.TypeOf((*map[string]struct{})(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(key); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *NonBlobSet) VDLRead(dec vdl.Decoder) error {
	*x = NonBlobSet{}
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
		case "Info":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "S":
			if err := __VDLReadAnon_set_3(dec, &x.S); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_set_3(dec vdl.Decoder, x *map[string]struct{}) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible set %T, from %v", *x, dec.Type())
	}
	var tmpMap map[string]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]struct{}, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key string
		{
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[string]struct{})
		}
		tmpMap[key] = struct{}{}
	}
}

type BlobOpt struct {
	Info string
	Bo   *BlobInfo
}

func (BlobOpt) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobOpt"`
}) {
}

func (m *BlobOpt) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Info == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Info"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Info), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Bo == (*BlobInfo)(nil))
	if var7 {
		if err := fieldsTarget1.ZeroField("Bo"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Bo")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Bo.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobOpt) MakeVDLTarget() vdl.Target {
	return &BlobOptTarget{Value: m}
}

type BlobOptTarget struct {
	Value      *BlobOpt
	infoTarget vdl.StringTarget
	boTarget   __VDLTarget4_optional
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobOptTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BlobOpt)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BlobOptTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bo":
		t.boTarget.Value = &t.Value.Bo
		target, err := &t.boTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *BlobOptTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobOptTarget) ZeroField(name string) error {
	switch name {
	case "Info":
		t.Value.Info = ""
		return nil
	case "Bo":
		t.Value.Bo = (*BlobInfo)(nil)
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *BlobOptTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Optional BlobInfo
type __VDLTarget4_optional struct {
	Value      **BlobInfo
	elemTarget BlobInfoTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *__VDLTarget4_optional) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if *t.Value == nil {
		*t.Value = &BlobInfo{}
	}
	t.elemTarget.Value = *t.Value
	target, err := &t.elemTarget, error(nil)
	if err != nil {
		return nil, err
	}
	return target.StartFields(tt)
}
func (t *__VDLTarget4_optional) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}
func (t *__VDLTarget4_optional) FromNil(tt *vdl.Type) error {
	*t.Value = (*BlobInfo)(nil)
	return nil
}

func (x BlobOpt) VDLIsZero() bool {
	return x == BlobOpt{}
}

func (x BlobOpt) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*BlobOpt)(nil)).Elem()); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Bo != nil {
		if err := enc.NextField("Bo"); err != nil {
			return err
		}
		enc.SetNextStartValueIsOptional()

		if err := x.Bo.VDLWrite(enc); err != nil {
			return err
		}

	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobOpt) VDLRead(dec vdl.Decoder) error {
	*x = BlobOpt{}
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
		case "Info":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Bo":
			if err := dec.StartValue(); err != nil {
				return err
			}
			if dec.IsNil() {
				if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(x.Bo), dec.Type()) {
					return fmt.Errorf("incompatible optional %T, from %v", x.Bo, dec.Type())
				}
				x.Bo = nil
				if err := dec.FinishValue(); err != nil {
					return err
				}
			} else {
				x.Bo = new(BlobInfo)
				dec.IgnoreNextStartValue()
				if err := x.Bo.VDLRead(dec); err != nil {
					return err
				}
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
	vdl.Register((*BlobInfo)(nil))
	vdl.Register((*BlobUnion)(nil))
	vdl.Register((*BlobSet)(nil))
	vdl.Register((*BlobAny)(nil))
	vdl.Register((*NonBlobSet)(nil))
	vdl.Register((*BlobOpt)(nil))

	return struct{}{}
}
