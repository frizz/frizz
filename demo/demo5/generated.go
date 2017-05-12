// info:{"Path":"frizz.io/demo/demo5","Hash":15036655152894811909}
package demo5

import (
	context "context"
	fmt "fmt"
	reflect "reflect"

	jsonctx "frizz.io/context/jsonctx"
	translation "frizz.io/demo/demo5/translation"
	system "frizz.io/system"
)

// notest

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}

func (v *PageRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo5", "@page"); err != nil {
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
func (v *PageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo5", "@page", system.J_NULL, nil
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
	return m, "frizz.io/demo/demo5", "@page", system.J_OBJECT, nil
}

type Page struct {
	*system.Object
	Body  translation.Localized
	Title translation.Localized
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func UnpackPageInterface(ctx context.Context, in system.Packed) (PageInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/demo/demo5", "page")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PageInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PageInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PageInterface.", in.Type())
	}
}
func (v *Page) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo5", "page"); err != nil {
		return err
	}
	if field, ok := in.Map()["body"]; ok && field.Type() != system.J_NULL {
		ob0, err := translation.UnpackLocalized(ctx, field)
		if err != nil {
			return err
		}
		v.Body = ob0
	}
	if field, ok := in.Map()["title"]; ok && field.Type() != system.J_NULL {
		ob0, err := translation.UnpackLocalized(ctx, field)
		if err != nil {
			return err
		}
		v.Title = ob0
	}
	return nil
}
func (v *Page) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo5", "page", system.J_NULL, nil
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
	if v.Body != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Body.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "frizz.io/demo/demo5/translation", "localized") {
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
		m["body"] = ob0
	}
	if v.Title != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Title.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "frizz.io/demo/demo5/translation", "localized") {
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
		m["title"] = ob0
	}
	return m, "frizz.io/demo/demo5", "page", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("frizz.io/demo/demo5")
	pkg.SetHash(uint64(0xd0acee31a6579f05))
	pkg.Init("page", func() interface{} {
		return new(Page)
	}, nil, func() interface{} {
		return new(PageRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PageInterface)(nil)).Elem()
	})
}
