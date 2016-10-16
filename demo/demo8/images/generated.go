// info:{"Path":"kego.io/demo/demo8/images","Hash":3695972038873785155}
package images

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/packer"
	"kego.io/system"
)

type PhotoRule struct {
	*system.Object
	*system.Rule
	Big *system.Bool `json:"big"`
}

func (v *PhotoRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
	if field, ok := in.Map()["big"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Big = ob0
	}
	return nil
}
func (v *PhotoRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo8/images", "@photo", nil
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
	if v.Big != nil {
		ob0, _, _, err := v.Big.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["big"] = ob0
	}
	return m, "kego.io/demo/demo8/images", "@photo", nil
}

type Photo struct {
	*system.Object
	Height *system.Int    `json:"height"`
	Url    *system.String `json:"url"`
	Width  *system.Int    `json:"width"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func UnpackPhotoInterface(ctx context.Context, in packer.Packed) (PhotoInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo8/images", "photo")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PhotoInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PhotoInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PhotoInterface.", in.Type())
	}
}
func (v *Photo) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["url"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Url = ob0
	}
	if field, ok := in.Map()["width"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Width = ob0
	}
	if field, ok := in.Map()["height"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Height = ob0
	}
	return nil
}
func (v *Photo) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo8/images", "photo", nil
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
	if v.Url != nil {
		ob0, _, _, err := v.Url.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["url"] = ob0
	}
	if v.Width != nil {
		ob0, _, _, err := v.Width.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["width"] = ob0
	}
	if v.Height != nil {
		ob0, _, _, err := v.Height.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["height"] = ob0
	}
	return m, "kego.io/demo/demo8/images", "photo", nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo8/images")
	pkg.SetHash(3695972038873785155)
	pkg.Init("photo", func() interface{} { return new(Photo) }, func() interface{} { return new(PhotoRule) }, func() reflect.Type { return reflect.TypeOf((*PhotoInterface)(nil)).Elem() })
}
