// info:{"Path":"kego.io/process/validate/tests","Hash":17573435268142783549}
package tests

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for a
type ARule struct {
	*system.Object
	*system.Rule
}

func (v *ARule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@a"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *ARule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@a", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/process/validate/tests", "@a", system.J_OBJECT, nil
}

// Automatically created basic rule for b
type BRule struct {
	*system.Object
	*system.Rule
}

func (v *BRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@b"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *BRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@b", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/process/validate/tests", "@b", system.J_OBJECT, nil
}

type CRule struct {
	*system.Object
	*system.Rule
	Fail *system.Bool `json:"fail"`
}

func (v *CRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@c"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["fail"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Fail = ob0
	}
	return nil
}
func (v *CRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@c", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Fail != nil {
		ob0, _, _, _, err := v.Fail.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["fail"] = ob0
	}
	return m, "kego.io/process/validate/tests", "@c", system.J_OBJECT, nil
}

// Automatically created basic rule for d
type DRule struct {
	*system.Object
	*system.Rule
}

func (v *DRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@d"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *DRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@d", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/process/validate/tests", "@d", system.J_OBJECT, nil
}

// Automatically created basic rule for e
type ERule struct {
	*system.Object
	*system.Rule
}

func (v *ERule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@e"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *ERule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@e", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/process/validate/tests", "@e", system.J_OBJECT, nil
}

// Automatically created basic rule for f
type FRule struct {
	*system.Object
	*system.Rule
}

func (v *FRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "@f"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *FRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "@f", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/process/validate/tests", "@f", system.J_OBJECT, nil
}

// A is a simple type containing a string B
type A struct {
	*system.Object
	B *system.String `json:"b"`
}
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}
func UnpackAInterface(ctx context.Context, in system.Packed) (AInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "a")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AInterface.", in.Type())
	}
}
func (v *A) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "a"); err != nil {
		return err
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.B = ob0
	}
	return nil
}
func (v *A) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "a", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.B != nil {
		ob0, _, _, _, err := v.B.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["b"] = ob0
	}
	return m, "kego.io/process/validate/tests", "a", system.J_OBJECT, nil
}

// A is a type containing a string interface B
type B struct {
	*system.Object
	C system.StringInterface `json:"c"`
}
type BInterface interface {
	GetB(ctx context.Context) *B
}

func (o *B) GetB(ctx context.Context) *B {
	return o
}
func UnpackBInterface(ctx context.Context, in system.Packed) (BInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "b")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(BInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement BInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into BInterface.", in.Type())
	}
}
func (v *B) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "b"); err != nil {
		return err
	}
	if field, ok := in.Map()["c"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.C = ob0
	}
	return nil
}
func (v *B) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "b", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.C != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.C.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "string") {
			typRef := system.NewReference(pkg, name)
			typeVal, err := typRef.ValueContext(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = map[string]interface{}{
				"type":  typeVal,
				"value": ob0_value,
			}
		} else {
			ob0 = ob0_value
		}
		m["c"] = ob0
	}
	return m, "kego.io/process/validate/tests", "b", system.J_OBJECT, nil
}
func UnpackC(ctx context.Context, in system.Packed) (C, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "c")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(C)
		if !ok {
			return nil, fmt.Errorf("%T does not implement C", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into C.", in.Type())
	}
}

// D is a type containing a field of type C
type D struct {
	*system.Object
	A C `json:"a"`
}
type DInterface interface {
	GetD(ctx context.Context) *D
}

func (o *D) GetD(ctx context.Context) *D {
	return o
}
func UnpackDInterface(ctx context.Context, in system.Packed) (DInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "d")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(DInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement DInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into DInterface.", in.Type())
	}
}
func (v *D) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "d"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackC(ctx, field)
		if err != nil {
			return err
		}
		v.A = ob0
	}
	return nil
}
func (v *D) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "d", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.A != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.A.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/process/validate/tests", "c") {
			typRef := system.NewReference(pkg, name)
			typeVal, err := typRef.ValueContext(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = map[string]interface{}{
				"type":  typeVal,
				"value": ob0_value,
			}
		} else {
			ob0 = ob0_value
		}
		m["a"] = ob0
	}
	return m, "kego.io/process/validate/tests", "d", system.J_OBJECT, nil
}

// E is a type containing an array of strings, and a map of strings
type E struct {
	*system.Object
	A []*system.String          `json:"a"`
	B map[string]*system.String `json:"b"`
}
type EInterface interface {
	GetE(ctx context.Context) *E
}

func (o *E) GetE(ctx context.Context) *E {
	return o
}
func UnpackEInterface(ctx context.Context, in system.Packed) (EInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "e")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(EInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement EInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into EInterface.", in.Type())
	}
}
func (v *E) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "e"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.String{}
		for i0 := range field.Array() {
			ob1 := new(system.String)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.A = ob0
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.String{}
		for k0 := range field.Map() {
			ob1 := new(system.String)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.B = ob0
	}
	return nil
}
func (v *E) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "e", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.A != nil {
		ob0 := []interface{}{}
		for i0 := range v.A {
			ob1, _, _, _, err := v.A[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["a"] = ob0
	}
	if v.B != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.B {
			ob1, _, _, _, err := v.B[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["b"] = ob0
	}
	return m, "kego.io/process/validate/tests", "e", system.J_OBJECT, nil
}

// F is a type with an extra rule attached to the field
type F struct {
	*system.Object
	A *A             `json:"a"`
	B []*A           `json:"b"`
	C map[string]*A  `json:"c"`
	D *system.String `json:"d"`
}
type FInterface interface {
	GetF(ctx context.Context) *F
}

func (o *F) GetF(ctx context.Context) *F {
	return o
}
func UnpackFInterface(ctx context.Context, in system.Packed) (FInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/process/validate/tests", "f")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(FInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement FInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into FInterface.", in.Type())
	}
}
func (v *F) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/process/validate/tests", "f"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0 := new(A)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*A{}
		for i0 := range field.Array() {
			ob1 := new(A)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.B = ob0
	}
	if field, ok := in.Map()["c"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*A{}
		for k0 := range field.Map() {
			ob1 := new(A)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.C = ob0
	}
	if field, ok := in.Map()["d"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.D = ob0
	}
	return nil
}
func (v *F) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/process/validate/tests", "f", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.A != nil {
		ob0, _, _, _, err := v.A.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["a"] = ob0
	}
	if v.B != nil {
		ob0 := []interface{}{}
		for i0 := range v.B {
			ob1, _, _, _, err := v.B[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["b"] = ob0
	}
	if v.C != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.C {
			ob1, _, _, _, err := v.C[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["c"] = ob0
	}
	if v.D != nil {
		ob0, _, _, _, err := v.D.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["d"] = ob0
	}
	return m, "kego.io/process/validate/tests", "f", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/process/validate/tests")
	pkg.SetHash(17573435268142783549)
	pkg.Init(
		"a",
		func() interface{} { return new(A) },
		nil,
		func() interface{} { return new(ARule) },
		func() reflect.Type { return reflect.TypeOf((*AInterface)(nil)).Elem() },
	)

	pkg.Init(
		"b",
		func() interface{} { return new(B) },
		nil,
		func() interface{} { return new(BRule) },
		func() reflect.Type { return reflect.TypeOf((*BInterface)(nil)).Elem() },
	)

	pkg.Init(
		"c",
		func() interface{} { return (*C)(nil) },
		nil,
		func() interface{} { return new(CRule) },
		func() reflect.Type { return reflect.TypeOf((*C)(nil)).Elem() },
	)

	pkg.Init(
		"d",
		func() interface{} { return new(D) },
		nil,
		func() interface{} { return new(DRule) },
		func() reflect.Type { return reflect.TypeOf((*DInterface)(nil)).Elem() },
	)

	pkg.Init(
		"e",
		func() interface{} { return new(E) },
		nil,
		func() interface{} { return new(ERule) },
		func() reflect.Type { return reflect.TypeOf((*EInterface)(nil)).Elem() },
	)

	pkg.Init(
		"f",
		func() interface{} { return new(F) },
		nil,
		func() interface{} { return new(FRule) },
		func() reflect.Type { return reflect.TypeOf((*FInterface)(nil)).Elem() },
	)

}
