// info:{"Path":"kego.io/tests/data/alias","Hash":3681006963164671295}
package alias

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/packer"
	"kego.io/system"
)

// Automatically created basic rule for alms
type AlmsRule struct {
	*system.Object
	*system.Rule
}

func (v *AlmsRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
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
func (v *AlmsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "@alms", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/tests/data/alias", "@alms", nil
}

// Automatically created basic rule for main
type MainRule struct {
	*system.Object
	*system.Rule
}

func (v *MainRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
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
func (v *MainRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "@main", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/tests/data/alias", "@main", nil
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}

func (v *SimpleRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
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
func (v *SimpleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "@simple", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "kego.io/tests/data/alias", "@simple", nil
}

type Alms map[string]*Simple
type AlmsInterface interface {
	GetAlms(ctx context.Context) Alms
}

func (o Alms) GetAlms(ctx context.Context) Alms {
	return o
}
func UnpackAlmsInterface(ctx context.Context, in packer.Packed) (AlmsInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data/alias", "alms")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlmsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlmsInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlmsInterface.", in.Type())
	}
}
func (v *Alms) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if iface {
		if in.Type() != packer.J_MAP {
			return fmt.Errorf("Invalid type %s while unpacking a map.", in.Type())
		}
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_MAP {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != packer.J_MAP {
		return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", in.Type())
	}
	ob0 := map[string]*Simple{}
	for k0 := range in.Map() {
		ob1 := new(Simple)
		if err := ob1.Unpack(ctx, in.Map()[k0], false); err != nil {
			return err
		}
		ob0[k0] = ob1
	}
	*v = ob0
	return nil
}
func (v *Alms) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "alms", nil
	}
	ob0 := map[string]interface{}{}
	for k0 := range *v {
		ob1, _, _, err := (*v)[k0].Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		ob0[k0] = ob1
	}
	return ob0, "kego.io/tests/data/alias", "alms", nil
}

type Main struct {
	*system.Object
	A Alms `json:"a"`
}
type MainInterface interface {
	GetMain(ctx context.Context) *Main
}

func (o *Main) GetMain(ctx context.Context) *Main {
	return o
}
func UnpackMainInterface(ctx context.Context, in packer.Packed) (MainInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data/alias", "main")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(MainInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement MainInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into MainInterface.", in.Type())
	}
}
func (v *Main) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != packer.J_NULL {
		ob0 := *new(Alms)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	return nil
}
func (v *Main) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "main", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.A != nil {
		ob0, _, _, err := v.A.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["a"] = ob0
	}
	return m, "kego.io/tests/data/alias", "main", nil
}

type Simple struct {
	*system.Object
	Js string `json:"js"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}
func UnpackSimpleInterface(ctx context.Context, in packer.Packed) (SimpleInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data/alias", "simple")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(SimpleInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement SimpleInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into SimpleInterface.", in.Type())
	}
}
func (v *Simple) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["js"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Js = ob0
	}
	return nil
}
func (v *Simple) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/tests/data/alias", "simple", nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Js != "" {
		ob0 := v.Js
		m["js"] = ob0
	}
	return m, "kego.io/tests/data/alias", "simple", nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/tests/data/alias")
	pkg.SetHash(3681006963164671295)
	pkg.Init("alms", func() interface{} { return new(Alms) }, func() interface{} { return new(AlmsRule) }, func() reflect.Type { return reflect.TypeOf((*AlmsInterface)(nil)).Elem() })
	pkg.Init("main", func() interface{} { return new(Main) }, func() interface{} { return new(MainRule) }, func() reflect.Type { return reflect.TypeOf((*MainInterface)(nil)).Elem() })
	pkg.Init("simple", func() interface{} { return new(Simple) }, func() interface{} { return new(SimpleRule) }, func() reflect.Type { return reflect.TypeOf((*SimpleInterface)(nil)).Elem() })
}
