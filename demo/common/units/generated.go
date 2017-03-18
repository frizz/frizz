// info:{"Path":"kego.io/demo/common/units","Hash":5562844427559446456}
package units

import (
	context "context"
	fmt "fmt"
	reflect "reflect"

	jsonctx "kego.io/context/jsonctx"
	system "kego.io/system"
)

// notest

// Automatically created basic rule for rectangle
type RectangleRule struct {
	*system.Object
	*system.Rule
}

func (v *RectangleRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/demo/common/units", "@rectangle"); err != nil {
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
func (v *RectangleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/units", "@rectangle", system.J_NULL, nil
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
	return m, "kego.io/demo/common/units", "@rectangle", system.J_OBJECT, nil
}

type Rectangle struct {
	*system.Object
	Height *system.Int `json:"height"`
	Width  *system.Int `json:"width"`
}
type RectangleInterface interface {
	GetRectangle(ctx context.Context) *Rectangle
}

func (o *Rectangle) GetRectangle(ctx context.Context) *Rectangle {
	return o
}
func UnpackRectangleInterface(ctx context.Context, in system.Packed) (RectangleInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/units", "rectangle")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(RectangleInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement RectangleInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into RectangleInterface.", in.Type())
	}
}
func (v *Rectangle) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("kego.io/demo/common/units", "rectangle"); err != nil {
		return err
	}
	if field, ok := in.Map()["height"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Height = ob0
	}
	if field, ok := in.Map()["width"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Width = ob0
	}
	return nil
}
func (v *Rectangle) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/units", "rectangle", system.J_NULL, nil
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
	if v.Height != nil {
		ob0, _, _, _, err := v.Height.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["height"] = ob0
	}
	if v.Width != nil {
		ob0, _, _, _, err := v.Width.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["width"] = ob0
	}
	return m, "kego.io/demo/common/units", "rectangle", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/common/units")
	pkg.SetHash(0x4d3331889d7127b8)
	pkg.Init("rectangle", func() interface{} {
		return new(Rectangle)
	}, nil, func() interface{} {
		return new(RectangleRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*RectangleInterface)(nil)).Elem()
	})
}
