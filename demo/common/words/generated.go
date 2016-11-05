// info:{"Path":"kego.io/demo/common/words","Hash":15839668451341961644}
package words

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for localizer
type LocalizerRule struct {
	*system.Object
	*system.Rule
}

func (v *LocalizerRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *LocalizerRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/words", "@localizer", system.J_NULL, nil
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
	return m, "kego.io/demo/common/words", "@localizer", system.J_OBJECT, nil
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
		return nil, "kego.io/demo/common/words", "@simple", system.J_NULL, nil
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
	return m, "kego.io/demo/common/words", "@simple", system.J_OBJECT, nil
}

// Automatically created basic rule for translation
type TranslationRule struct {
	*system.Object
	*system.Rule
}

func (v *TranslationRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *TranslationRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/words", "@translation", system.J_NULL, nil
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
	return m, "kego.io/demo/common/words", "@translation", system.J_OBJECT, nil
}
func UnpackLocalizer(ctx context.Context, in system.Packed) (Localizer, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/words", "localizer")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(Localizer)
		if !ok {
			return nil, fmt.Errorf("%T does not implement Localizer", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into Localizer.", in.Type())
	}
}

type Simple struct {
	*system.Object
	String *system.String `json:"string"`
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
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/words", "simple")
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
	if field, ok := in.Map()["string"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.String = ob0
	}
	return nil
}
func (v *Simple) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/words", "simple", system.J_NULL, nil
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
	if v.String != nil {
		ob0, _, _, _, err := v.String.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["string"] = ob0
	}
	return m, "kego.io/demo/common/words", "simple", system.J_OBJECT, nil
}

// This represents a translated string
type Translation struct {
	*system.Object
	// The original English string
	English *system.String `json:"english"`
	// The translated strings
	Translations map[string]*system.String `json:"translations"`
}
type TranslationInterface interface {
	GetTranslation(ctx context.Context) *Translation
}

func (o *Translation) GetTranslation(ctx context.Context) *Translation {
	return o
}
func UnpackTranslationInterface(ctx context.Context, in system.Packed) (TranslationInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/demo/common/words", "translation")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(TranslationInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement TranslationInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into TranslationInterface.", in.Type())
	}
}
func (v *Translation) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *Translation) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/demo/common/words", "translation", system.J_NULL, nil
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
	if v.Translations != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Translations {
			ob1, _, _, _, err := v.Translations[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["translations"] = ob0
	}
	return m, "kego.io/demo/common/words", "translation", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/common/words")
	pkg.SetHash(15839668451341961644)
	pkg.Init("localizer", func() interface{} { return (*Localizer)(nil) }, func() interface{} { return new(LocalizerRule) }, func() reflect.Type { return reflect.TypeOf((*Localizer)(nil)).Elem() })
	pkg.Init("simple", func() interface{} { return new(Simple) }, func() interface{} { return new(SimpleRule) }, func() reflect.Type { return reflect.TypeOf((*SimpleInterface)(nil)).Elem() })
	pkg.Init("translation", func() interface{} { return new(Translation) }, func() interface{} { return new(TranslationRule) }, func() reflect.Type { return reflect.TypeOf((*TranslationInterface)(nil)).Elem() })
}
