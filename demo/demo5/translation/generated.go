// info:{"Path":"kego.io/demo/demo5/translation","Hash":13251613636822643043}
package translation

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for localized
type LocalizedRule struct {
	*system.Object
	*system.Rule
}

func (v *LocalizedRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
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
func (v *LocalizedRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo5/translation", "@localized", system.J_NULL, nil
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
	return m, "kego.io/demo/demo5/translation", "@localized", system.J_OBJECT, nil
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}

func (v *SimpleRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
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
func (v *SimpleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo5/translation", "@simple", system.J_NULL, nil
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
	return m, "kego.io/demo/demo5/translation", "@simple", system.J_OBJECT, nil
}

// Automatically created basic rule for smartling
type SmartlingRule struct {
	*system.Object
	*system.Rule
}

func (v *SmartlingRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
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
func (v *SmartlingRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo5/translation", "@smartling", system.J_NULL, nil
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
	return m, "kego.io/demo/demo5/translation", "@smartling", system.J_OBJECT, nil
}
func UnpackLocalized(ctx context.Context, in system.Packed) (Localized, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo5/translation", "localized")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(Localized)
		if !ok {
			return nil, fmt.Errorf("%T does not implement Localized", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into Localized.", in.Type())
	}
}

type Simple struct {
	*system.Object
	Text *system.String `json:"text"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}
func UnpackSimpleInterface(ctx context.Context, in system.Packed) (SimpleInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo5/translation", "simple")
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
func (v *Simple) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["text"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Text = ob0
	}
	return nil
}
func (v *Simple) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo5/translation", "simple", system.J_NULL, nil
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
	if v.Text != nil {
		ob0, _, _, _, err := v.Text.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["text"] = ob0
	}
	return m, "kego.io/demo/demo5/translation", "simple", system.J_OBJECT, nil
}

type Smartling struct {
	*system.Object
	English      *system.String            `json:"english"`
	Translations map[string]*system.String `json:"translations"`
}
type SmartlingInterface interface {
	GetSmartling(ctx context.Context) *Smartling
}

func (o *Smartling) GetSmartling(ctx context.Context) *Smartling {
	return o
}
func UnpackSmartlingInterface(ctx context.Context, in system.Packed) (SmartlingInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/demo5/translation", "smartling")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(SmartlingInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement SmartlingInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into SmartlingInterface.", in.Type())
	}
}
func (v *Smartling) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["english"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.English = ob0
	}
	if field, ok := in.Map()["translations"]; ok && field.Type() != system.J_NULL {
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
		v.Translations = ob0
	}
	return nil
}
func (v *Smartling) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/demo5/translation", "smartling", system.J_NULL, nil
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
	if v.English != nil {
		ob0, _, _, _, err := v.English.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["english"] = ob0
	}
	return m, "kego.io/demo/demo5/translation", "smartling", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo5/translation")
	pkg.SetHash(13251613636822643043)
	pkg.Init("localized", func() interface{} { return (*Localized)(nil) }, func() interface{} { return new(LocalizedRule) }, func() reflect.Type { return reflect.TypeOf((*Localized)(nil)).Elem() })
	pkg.Init("simple", func() interface{} { return new(Simple) }, func() interface{} { return new(SimpleRule) }, func() reflect.Type { return reflect.TypeOf((*SimpleInterface)(nil)).Elem() })
	pkg.Init("smartling", func() interface{} { return new(Smartling) }, func() interface{} { return new(SmartlingRule) }, func() reflect.Type { return reflect.TypeOf((*SmartlingInterface)(nil)).Elem() })
}
