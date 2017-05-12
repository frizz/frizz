// info:{"Path":"frizz.io/demo/demo7/images","Hash":8088901034031058716}
package images

import (
	context "context"
	fmt "fmt"
	reflect "reflect"

	jsonctx "frizz.io/context/jsonctx"
	system "frizz.io/system"
)

// notest

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
}

func (v *PhotoRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo7/images", "@photo"); err != nil {
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
func (v *PhotoRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo7/images", "@photo", system.J_NULL, nil
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
	return m, "frizz.io/demo/demo7/images", "@photo", system.J_OBJECT, nil
}

type Photo struct {
	*system.Object
	Height *system.Int
	Url    *system.String
	Width  *system.Int
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func UnpackPhotoInterface(ctx context.Context, in system.Packed) (PhotoInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/demo/demo7/images", "photo")
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
func (v *Photo) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo7/images", "photo"); err != nil {
		return err
	}
	if field, ok := in.Map()["height"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Height = ob0
	}
	if field, ok := in.Map()["url"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Url = ob0
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
func (v *Photo) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo7/images", "photo", system.J_NULL, nil
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
	if v.Url != nil {
		ob0, _, _, _, err := v.Url.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["url"] = ob0
	}
	if v.Width != nil {
		ob0, _, _, _, err := v.Width.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["width"] = ob0
	}
	return m, "frizz.io/demo/demo7/images", "photo", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("frizz.io/demo/demo7/images")
	pkg.SetHash(uint64(0x70418ca234b60f1c))
	pkg.Init("photo", func() interface{} {
		return new(Photo)
	}, nil, func() interface{} {
		return new(PhotoRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PhotoInterface)(nil)).Elem()
	})
}
