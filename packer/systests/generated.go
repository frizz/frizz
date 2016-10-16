// info:{"Path":"kego.io/packer/systests","Hash":16375361147145841280}
package systests

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/packer"
	"kego.io/system"
)

// Automatically created basic rule for a
type ARule struct {
	*system.Object
	*system.Rule
}

func (v *ARule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
func (v *ARule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/packer/systests", "@a", nil
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
	return m, "kego.io/packer/systests", "@a", nil
}

type A struct {
	*system.Object
	A system.StringInterface `json:"a"`
	B system.NumberInterface `json:"b"`
	C system.BoolInterface   `json:"c"`
	D system.StringInterface `json:"d"`
}
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}
func UnpackAInterface(ctx context.Context, in packer.Packed) (AInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/packer/systests", "a")
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
func (v *A) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.A = ob0
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackNumberInterface(ctx, field)
		if err != nil {
			return err
		}
		v.B = ob0
	}
	if field, ok := in.Map()["c"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackBoolInterface(ctx, field)
		if err != nil {
			return err
		}
		v.C = ob0
	}
	if field, ok := in.Map()["d"]; ok && field.Type() != packer.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.D = ob0
	}
	return nil
}
func (v *A) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/packer/systests", "a", nil
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
	return m, "kego.io/packer/systests", "a", nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/packer/systests")
	pkg.SetHash(16375361147145841280)
	pkg.Init("a", func() interface{} { return new(A) }, func() interface{} { return new(ARule) }, func() reflect.Type { return reflect.TypeOf((*AInterface)(nil)).Elem() })
}
