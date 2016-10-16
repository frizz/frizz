// info:{"Path":"kego.io/demo/common/images","Hash":7237742861799384387}
package images

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/demo/common/units"
	"kego.io/packer"
	"kego.io/system"
)

// Automatically created basic rule for icon
type IconRule struct {
	*system.Object
	*system.Rule
}

func (v *IconRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
func (v *IconRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/images", "@icon", nil
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
	return m, "kego.io/demo/common/images", "@icon", nil
}

// Restriction rules for images
type ImageRule struct {
	*system.Object
	*system.Rule
	Secure *system.Bool `json:"secure"`
}

func (v *ImageRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
	if field, ok := in.Map()["secure"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Secure = ob0
	}
	return nil
}
func (v *ImageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/images", "@image", nil
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
	if v.Secure != nil {
		ob0, _, _, err := v.Secure.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["secure"] = ob0
	}
	return m, "kego.io/demo/common/images", "@image", nil
}

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
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
	return nil
}
func (v *PhotoRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/images", "@photo", nil
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
	return m, "kego.io/demo/common/images", "@photo", nil
}

// This is a type of image, which just contains the url of the image
type Icon struct {
	*system.Object
	Url *system.String `json:"url"`
}
type IconInterface interface {
	GetIcon(ctx context.Context) *Icon
}

func (o *Icon) GetIcon(ctx context.Context) *Icon {
	return o
}
func UnpackIconInterface(ctx context.Context, in packer.Packed) (IconInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/images", "icon")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(IconInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement IconInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into IconInterface.", in.Type())
	}
}
func (v *Icon) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
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
	return nil
}
func (v *Icon) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/images", "icon", nil
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
	return m, "kego.io/demo/common/images", "icon", nil
}
func UnpackImage(ctx context.Context, in packer.Packed) (Image, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/images", "image")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(Image)
		if !ok {
			return nil, fmt.Errorf("%T does not implement Image", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into Image.", in.Type())
	}
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Object
	// The path for the url - e.g. /foo/bar.jpg
	Path *system.String `json:"path"`
	// The protocol for the url - e.g. http or https
	Protocol *system.String `json:"protocol"`
	// The server for the url - e.g. www.google.com
	Server *system.String   `json:"server"`
	Size   *units.Rectangle `json:"size"`
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
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/images", "photo")
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
	if field, ok := in.Map()["server"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Server = ob0
	}
	if field, ok := in.Map()["size"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(units.Rectangle)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Size = ob0
	}
	if field, ok := in.Map()["path"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Path = ob0
	}
	if field, ok := in.Map()["protocol"]; ok && field.Type() != packer.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Protocol = ob0
	} else {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, packer.Pack("http"), false); err != nil {
			return err
		}
		v.Protocol = ob0
	}
	return nil
}
func (v *Photo) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/images", "photo", nil
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
	if v.Server != nil {
		ob0, _, _, err := v.Server.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["server"] = ob0
	}
	if v.Size != nil {
		ob0, _, _, err := v.Size.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["size"] = ob0
	}
	if v.Path != nil {
		ob0, _, _, err := v.Path.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["path"] = ob0
	}
	if v.Protocol != nil {
		ob0, _, _, err := v.Protocol.Repack(ctx)
		if err != nil {
			return nil, "", "", err
		}
		m["protocol"] = ob0
	}
	return m, "kego.io/demo/common/images", "photo", nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/common/images")
	pkg.SetHash(7237742861799384387)
	pkg.Init("icon", func() interface{} { return new(Icon) }, func() interface{} { return new(IconRule) }, func() reflect.Type { return reflect.TypeOf((*IconInterface)(nil)).Elem() })
	pkg.Init("image", func() interface{} { return new(Image) }, func() interface{} { return new(ImageRule) }, func() reflect.Type { return reflect.TypeOf((*Image)(nil)).Elem() })
	pkg.Init("photo", func() interface{} { return new(Photo) }, func() interface{} { return new(PhotoRule) }, func() reflect.Type { return reflect.TypeOf((*PhotoInterface)(nil)).Elem() })
}
