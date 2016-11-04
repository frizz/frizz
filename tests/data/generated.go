// info:{"Path":"kego.io/tests/data","Hash":6167257871666940584}
package data

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for alajs
type AlajsRule struct {
	*system.Object
	*system.Rule
}

func (v *AlajsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlajsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@alajs", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@alajs", system.J_OBJECT, nil
}

// Automatically created basic rule for alas
type AlasRule struct {
	*system.Object
	*system.Rule
}

func (v *AlasRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlasRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@alas", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@alas", system.J_OBJECT, nil
}

// Automatically created basic rule for alass
type AlassRule struct {
	*system.Object
	*system.Rule
}

func (v *AlassRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlassRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@alass", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@alass", system.J_OBJECT, nil
}

// Automatically created basic rule for aljb
type AljbRule struct {
	*system.Object
	*system.Rule
}

func (v *AljbRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AljbRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@aljb", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@aljb", system.J_OBJECT, nil
}

// Automatically created basic rule for aljn
type AljnRule struct {
	*system.Object
	*system.Rule
}

func (v *AljnRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AljnRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@aljn", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@aljn", system.J_OBJECT, nil
}

// Automatically created basic rule for aljs
type AljsRule struct {
	*system.Object
	*system.Rule
}

func (v *AljsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AljsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@aljs", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@aljs", system.J_OBJECT, nil
}

// Automatically created basic rule for almjs
type AlmjsRule struct {
	*system.Object
	*system.Rule
}

func (v *AlmjsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlmjsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@almjs", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@almjs", system.J_OBJECT, nil
}

// Automatically created basic rule for alms
type AlmsRule struct {
	*system.Object
	*system.Rule
}

func (v *AlmsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlmsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@alms", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@alms", system.J_OBJECT, nil
}

// Automatically created basic rule for almss
type AlmssRule struct {
	*system.Object
	*system.Rule
}

func (v *AlmssRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlmssRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@almss", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@almss", system.J_OBJECT, nil
}

// Automatically created basic rule for als
type AlsRule struct {
	*system.Object
	*system.Rule
}

func (v *AlsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@als", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@als", system.J_OBJECT, nil
}

// Automatically created basic rule for alss
type AlssRule struct {
	*system.Object
	*system.Rule
}

func (v *AlssRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *AlssRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@alss", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@alss", system.J_OBJECT, nil
}

// Automatically created basic rule for face
type FaceRule struct {
	*system.Object
	*system.Rule
}

func (v *FaceRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *FaceRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@face", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@face", system.J_OBJECT, nil
}

// Automatically created basic rule for facea
type FaceaRule struct {
	*system.Object
	*system.Rule
}

func (v *FaceaRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *FaceaRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@facea", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@facea", system.J_OBJECT, nil
}

// Automatically created basic rule for faceb
type FacebRule struct {
	*system.Object
	*system.Rule
}

func (v *FacebRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *FacebRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@faceb", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@faceb", system.J_OBJECT, nil
}

// Automatically created basic rule for multi
type MultiRule struct {
	*system.Object
	*system.Rule
}

func (v *MultiRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
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
func (v *MultiRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "@multi", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@multi", system.J_OBJECT, nil
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
		return nil, "kego.io/tests/data", "@simple", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "@simple", system.J_OBJECT, nil
}

type Alajs []string
type AlajsInterface interface {
	GetAlajs(ctx context.Context) Alajs
}

func (o Alajs) GetAlajs(ctx context.Context) Alajs {
	return o
}
func UnpackAlajsInterface(ctx context.Context, in system.Packed) (AlajsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "alajs")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlajsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlajsInterface", i)
		}
		return ob, nil
	case system.J_ARRAY:
		ob := new(Alajs)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlajsInterface.", in.Type())
	}
}
func (v *Alajs) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", in.Type())
	}
	ob0 := []string{}
	for i0 := range in.Array() {
		ob1, err := system.UnpackString(ctx, in.Array()[i0])
		if err != nil {
			return err
		}
		ob0 = append(ob0, ob1)
	}
	*v = ob0
	return nil
}
func (v Alajs) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "alajs", system.J_NULL, nil
	}
	ob0 := []interface{}{}
	for i0 := range v {
		ob1 := v[i0]
		ob0 = append(ob0, ob1)
	}
	return ob0, "kego.io/tests/data", "alajs", system.J_ARRAY, nil
}

type Alas []*Simple
type AlasInterface interface {
	GetAlas(ctx context.Context) Alas
}

func (o Alas) GetAlas(ctx context.Context) Alas {
	return o
}
func UnpackAlasInterface(ctx context.Context, in system.Packed) (AlasInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "alas")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlasInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlasInterface", i)
		}
		return ob, nil
	case system.J_ARRAY:
		ob := new(Alas)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlasInterface.", in.Type())
	}
}
func (v *Alas) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", in.Type())
	}
	ob0 := []*Simple{}
	for i0 := range in.Array() {
		ob1 := new(Simple)
		if err := ob1.Unpack(ctx, in.Array()[i0], false); err != nil {
			return err
		}
		ob0 = append(ob0, ob1)
	}
	*v = ob0
	return nil
}
func (v Alas) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "alas", system.J_NULL, nil
	}
	ob0 := []interface{}{}
	for i0 := range v {
		ob1, _, _, _, err := v[i0].Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		ob0 = append(ob0, ob1)
	}
	return ob0, "kego.io/tests/data", "alas", system.J_ARRAY, nil
}

type Alass []*system.String
type AlassInterface interface {
	GetAlass(ctx context.Context) Alass
}

func (o Alass) GetAlass(ctx context.Context) Alass {
	return o
}
func UnpackAlassInterface(ctx context.Context, in system.Packed) (AlassInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "alass")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlassInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlassInterface", i)
		}
		return ob, nil
	case system.J_ARRAY:
		ob := new(Alass)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlassInterface.", in.Type())
	}
}
func (v *Alass) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_ARRAY {
		return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", in.Type())
	}
	ob0 := []*system.String{}
	for i0 := range in.Array() {
		ob1 := new(system.String)
		if err := ob1.Unpack(ctx, in.Array()[i0], false); err != nil {
			return err
		}
		ob0 = append(ob0, ob1)
	}
	*v = ob0
	return nil
}
func (v Alass) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "alass", system.J_NULL, nil
	}
	ob0 := []interface{}{}
	for i0 := range v {
		ob1, _, _, _, err := v[i0].Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		ob0 = append(ob0, ob1)
	}
	return ob0, "kego.io/tests/data", "alass", system.J_ARRAY, nil
}

type Aljb bool
type AljbInterface interface {
	GetAljb(ctx context.Context) *Aljb
}

func (o *Aljb) GetAljb(ctx context.Context) *Aljb {
	return o
}
func UnpackAljbInterface(ctx context.Context, in system.Packed) (AljbInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "aljb")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AljbInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AljbInterface", i)
		}
		return ob, nil
	case system.J_BOOL:
		ob := new(Aljb)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AljbInterface.", in.Type())
	}
}
func (v *Aljb) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_BOOL {
		return fmt.Errorf("Invalid type %s while unpacking a bool.", in.Type())
	}
	*v = Aljb(in.Bool())
	return nil
}
func (v *Aljb) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "aljb", system.J_NULL, nil
	}
	if v != nil {
		return bool(*v), "kego.io/tests/data", "aljb", system.J_BOOL, nil
	}
	return nil, "kego.io/tests/data", "aljb", system.J_BOOL, nil
}

type Aljn float64
type AljnInterface interface {
	GetAljn(ctx context.Context) *Aljn
}

func (o *Aljn) GetAljn(ctx context.Context) *Aljn {
	return o
}
func UnpackAljnInterface(ctx context.Context, in system.Packed) (AljnInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "aljn")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AljnInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AljnInterface", i)
		}
		return ob, nil
	case system.J_NUMBER:
		ob := new(Aljn)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AljnInterface.", in.Type())
	}
}
func (v *Aljn) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_NUMBER {
		return fmt.Errorf("Invalid type %s while unpacking a number.", in.Type())
	}
	*v = Aljn(in.Number())
	return nil
}
func (v *Aljn) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "aljn", system.J_NULL, nil
	}
	if v != nil {
		return float64(*v), "kego.io/tests/data", "aljn", system.J_NUMBER, nil
	}
	return nil, "kego.io/tests/data", "aljn", system.J_NUMBER, nil
}

type Aljs string
type AljsInterface interface {
	GetAljs(ctx context.Context) *Aljs
}

func (o *Aljs) GetAljs(ctx context.Context) *Aljs {
	return o
}
func UnpackAljsInterface(ctx context.Context, in system.Packed) (AljsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "aljs")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AljsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AljsInterface", i)
		}
		return ob, nil
	case system.J_STRING:
		ob := new(Aljs)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AljsInterface.", in.Type())
	}
}
func (v *Aljs) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_STRING {
		return fmt.Errorf("Invalid type %s while unpacking a string.", in.Type())
	}
	*v = Aljs(in.String())
	return nil
}
func (v *Aljs) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "aljs", system.J_NULL, nil
	}
	if v != nil {
		return string(*v), "kego.io/tests/data", "aljs", system.J_STRING, nil
	}
	return nil, "kego.io/tests/data", "aljs", system.J_STRING, nil
}

type Almjs map[string]string
type AlmjsInterface interface {
	GetAlmjs(ctx context.Context) Almjs
}

func (o Almjs) GetAlmjs(ctx context.Context) Almjs {
	return o
}
func UnpackAlmjsInterface(ctx context.Context, in system.Packed) (AlmjsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "almjs")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlmjsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlmjsInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlmjsInterface.", in.Type())
	}
}
func (v *Almjs) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if iface {
		if in.Type() != system.J_MAP {
			return fmt.Errorf("Invalid type %s while unpacking a map.", in.Type())
		}
		in = in.Map()["value"]
	}
	if in.Type() != system.J_MAP {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_MAP {
		return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", in.Type())
	}
	ob0 := map[string]string{}
	for k0 := range in.Map() {
		ob1, err := system.UnpackString(ctx, in.Map()[k0])
		if err != nil {
			return err
		}
		ob0[k0] = ob1
	}
	*v = ob0
	return nil
}
func (v Almjs) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "almjs", system.J_NULL, nil
	}
	ob0 := map[string]interface{}{}
	for k0 := range v {
		ob1 := v[k0]
		ob0[k0] = ob1
	}
	return ob0, "kego.io/tests/data", "almjs", system.J_MAP, nil
}

type Alms map[string]*Simple
type AlmsInterface interface {
	GetAlms(ctx context.Context) Alms
}

func (o Alms) GetAlms(ctx context.Context) Alms {
	return o
}
func UnpackAlmsInterface(ctx context.Context, in system.Packed) (AlmsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "alms")
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
func (v *Alms) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if iface {
		if in.Type() != system.J_MAP {
			return fmt.Errorf("Invalid type %s while unpacking a map.", in.Type())
		}
		in = in.Map()["value"]
	}
	if in.Type() != system.J_MAP {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_MAP {
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
func (v Alms) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "alms", system.J_NULL, nil
	}
	ob0 := map[string]interface{}{}
	for k0 := range v {
		ob1, _, _, _, err := v[k0].Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		ob0[k0] = ob1
	}
	return ob0, "kego.io/tests/data", "alms", system.J_MAP, nil
}

type Almss map[string]*system.String
type AlmssInterface interface {
	GetAlmss(ctx context.Context) Almss
}

func (o Almss) GetAlmss(ctx context.Context) Almss {
	return o
}
func UnpackAlmssInterface(ctx context.Context, in system.Packed) (AlmssInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "almss")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlmssInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlmssInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlmssInterface.", in.Type())
	}
}
func (v *Almss) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if iface {
		if in.Type() != system.J_MAP {
			return fmt.Errorf("Invalid type %s while unpacking a map.", in.Type())
		}
		in = in.Map()["value"]
	}
	if in.Type() != system.J_MAP {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != system.J_MAP {
		return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", in.Type())
	}
	ob0 := map[string]*system.String{}
	for k0 := range in.Map() {
		ob1 := new(system.String)
		if err := ob1.Unpack(ctx, in.Map()[k0], false); err != nil {
			return err
		}
		ob0[k0] = ob1
	}
	*v = ob0
	return nil
}
func (v Almss) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "almss", system.J_NULL, nil
	}
	ob0 := map[string]interface{}{}
	for k0 := range v {
		ob1, _, _, _, err := v[k0].Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		ob0[k0] = ob1
	}
	return ob0, "kego.io/tests/data", "almss", system.J_MAP, nil
}

type Als Simple
type AlsInterface interface {
	GetAls(ctx context.Context) *Als
}

func (o *Als) GetAls(ctx context.Context) *Als {
	return o
}
func UnpackAlsInterface(ctx context.Context, in system.Packed) (AlsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "als")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlsInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlsInterface.", in.Type())
	}
}
func (v *Als) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["js"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Js = ob0
	}
	return nil
}
func (v *Als) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "als", system.J_NULL, nil
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
	if v.Js != "" {
		ob0 := v.Js
		m["js"] = ob0
	}
	return m, "kego.io/tests/data", "als", system.J_OBJECT, nil
}

type Alss system.String
type AlssInterface interface {
	GetAlss(ctx context.Context) *Alss
}

func (o *Alss) GetAlss(ctx context.Context) *Alss {
	return o
}
func UnpackAlssInterface(ctx context.Context, in system.Packed) (AlssInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "alss")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(AlssInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement AlssInterface", i)
		}
		return ob, nil
	case system.J_STRING:
		ob := new(Alss)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into AlssInterface.", in.Type())
	}
}
func (v *Alss) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if in.Type() == system.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != system.J_STRING {
		return fmt.Errorf("Invalid type %s while unpacking a string.", in.Type())
	}
	*v = Alss(in.String())
	return nil
}
func (v *Alss) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "alss", system.J_NULL, nil
	}
	if v != nil {
		return string(*v), "kego.io/tests/data", "alss", system.J_STRING, nil
	}
	return nil, "kego.io/tests/data", "alss", system.J_STRING, nil
}
func UnpackFace(ctx context.Context, in system.Packed) (Face, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "face")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(Face)
		if !ok {
			return nil, fmt.Errorf("%T does not implement Face", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into Face.", in.Type())
	}
}

type Facea struct {
	*system.Object
	A *system.String `json:"a"`
}
type FaceaInterface interface {
	GetFacea(ctx context.Context) *Facea
}

func (o *Facea) GetFacea(ctx context.Context) *Facea {
	return o
}
func UnpackFaceaInterface(ctx context.Context, in system.Packed) (FaceaInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "facea")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(FaceaInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement FaceaInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into FaceaInterface.", in.Type())
	}
}
func (v *Facea) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	return nil
}
func (v *Facea) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "facea", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "facea", system.J_OBJECT, nil
}

type Faceb struct {
	*system.Object
	B *system.String `json:"b"`
}
type FacebInterface interface {
	GetFaceb(ctx context.Context) *Faceb
}

func (o *Faceb) GetFaceb(ctx context.Context) *Faceb {
	return o
}
func UnpackFacebInterface(ctx context.Context, in system.Packed) (FacebInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "faceb")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(FacebInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement FacebInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into FacebInterface.", in.Type())
	}
}
func (v *Faceb) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
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
func (v *Faceb) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "faceb", system.J_NULL, nil
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
	return m, "kego.io/tests/data", "faceb", system.J_OBJECT, nil
}

// v2
type Multi struct {
	*system.Object
	Aalajs []Alajs                           `json:"aalajs"`
	Aalas  []Alas                            `json:"aalas"`
	Aalass []Alass                           `json:"aalass"`
	Aaljb  []*Aljb                           `json:"aaljb"`
	Aaljn  []*Aljn                           `json:"aaljn"`
	Aaljs  []*Aljs                           `json:"aaljs"`
	Aalmjs []Almjs                           `json:"aalmjs"`
	Aalms  []Alms                            `json:"aalms"`
	Aalmss []Almss                           `json:"aalmss"`
	Aals   []*Als                            `json:"aals"`
	Aalss  []*Alss                           `json:"aalss"`
	Ai     []Face                            `json:"ai"`
	Ajb    []bool                            `json:"ajb"`
	Ajn    []float64                         `json:"ajn"`
	Ajs    []string                          `json:"ajs"`
	Alajs  Alajs                             `json:"alajs"`
	Alas   Alas                              `json:"alas"`
	Alass  Alass                             `json:"alass"`
	Aljb   *Aljb                             `json:"aljb"`
	Aljbi  AljbInterface                     `json:"aljbi"`
	Aljn   *Aljn                             `json:"aljn"`
	Aljni  AljnInterface                     `json:"aljni"`
	Aljs   *Aljs                             `json:"aljs"`
	Aljsi  AljsInterface                     `json:"aljsi"`
	Almjs  Almjs                             `json:"almjs"`
	Alms   Alms                              `json:"alms"`
	Almss  Almss                             `json:"almss"`
	Als    *Als                              `json:"als"`
	Alss   *Alss                             `json:"alss"`
	Am     []*Multi                          `json:"am"`
	Anri   []system.StringInterface          `json:"anri"`
	Ari    []system.RuleInterface            `json:"ari"`
	Asb    []*system.Bool                    `json:"asb"`
	Asi    []*system.Int                     `json:"asi"`
	Asn    []*system.Number                  `json:"asn"`
	Asp    []*system.Package                 `json:"asp"`
	Asr    []*system.Reference               `json:"asr"`
	Ass    []*system.String                  `json:"ass"`
	Bri    system.BoolInterface              `json:"bri"`
	I      Face                              `json:"i"`
	Jb     bool                              `json:"jb"`
	Jn     float64                           `json:"jn"`
	Js     string                            `json:"js"`
	M      *Multi                            `json:"m"`
	Malajs map[string]Alajs                  `json:"malajs"`
	Malas  map[string]Alas                   `json:"malas"`
	Malass map[string]Alass                  `json:"malass"`
	Maljb  map[string]*Aljb                  `json:"maljb"`
	Maljn  map[string]*Aljn                  `json:"maljn"`
	Maljs  map[string]*Aljs                  `json:"maljs"`
	Malmjs map[string]Almjs                  `json:"malmjs"`
	Malms  map[string]Alms                   `json:"malms"`
	Malmss map[string]Almss                  `json:"malmss"`
	Mals   map[string]*Als                   `json:"mals"`
	Malss  map[string]*Alss                  `json:"malss"`
	Mi     map[string]Face                   `json:"mi"`
	Mjb    map[string]bool                   `json:"mjb"`
	Mjn    map[string]float64                `json:"mjn"`
	Mjs    map[string]string                 `json:"mjs"`
	Mm     map[string]*Multi                 `json:"mm"`
	Mnri   map[string]system.StringInterface `json:"mnri"`
	Mri    map[string]system.RuleInterface   `json:"mri"`
	Msb    map[string]*system.Bool           `json:"msb"`
	Msi    map[string]*system.Int            `json:"msi"`
	Msn    map[string]*system.Number         `json:"msn"`
	Msp    map[string]*system.Package        `json:"msp"`
	Msr    map[string]*system.Reference      `json:"msr"`
	Mss    map[string]*system.String         `json:"mss"`
	Nri    system.NumberInterface            `json:"nri"`
	Ri     system.RuleInterface              `json:"ri"`
	Sb     *system.Bool                      `json:"sb"`
	Si     *system.Int                       `json:"si"`
	Sn     *system.Number                    `json:"sn"`
	Sp     *system.Package                   `json:"sp"`
	Sr     *system.Reference                 `json:"sr"`
	Sri    system.StringInterface            `json:"sri"`
	Ss     *system.String                    `json:"ss"`
}
type MultiInterface interface {
	GetMulti(ctx context.Context) *Multi
}

func (o *Multi) GetMulti(ctx context.Context) *Multi {
	return o
}
func UnpackMultiInterface(ctx context.Context, in system.Packed) (MultiInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "multi")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(MultiInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement MultiInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into MultiInterface.", in.Type())
	}
}
func (v *Multi) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["i"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackFace(ctx, field)
		if err != nil {
			return err
		}
		v.I = ob0
	}
	if field, ok := in.Map()["anri"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []system.StringInterface{}
		for i0 := range field.Array() {
			ob1, err := system.UnpackStringInterface(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Anri = ob0
	}
	if field, ok := in.Map()["mjn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]float64{}
		for k0 := range field.Map() {
			ob1, err := system.UnpackNumber(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mjn = ob0
	}
	if field, ok := in.Map()["nri"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackNumberInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Nri = ob0
	}
	if field, ok := in.Map()["malms"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Alms{}
		for k0 := range field.Map() {
			ob1 := *new(Alms)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malms = ob0
	}
	if field, ok := in.Map()["msr"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.Reference{}
		for k0 := range field.Map() {
			ob1 := new(system.Reference)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Msr = ob0
	}
	if field, ok := in.Map()["aals"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Als{}
		for i0 := range field.Array() {
			ob1 := new(Als)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aals = ob0
	}
	if field, ok := in.Map()["sn"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Sn = ob0
	}
	if field, ok := in.Map()["aljbi"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackAljbInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Aljbi = ob0
	}
	if field, ok := in.Map()["jn"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackNumber(ctx, field)
		if err != nil {
			return err
		}
		v.Jn = ob0
	}
	if field, ok := in.Map()["malajs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Alajs{}
		for k0 := range field.Map() {
			ob1 := *new(Alajs)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malajs = ob0
	}
	if field, ok := in.Map()["am"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Multi{}
		for i0 := range field.Array() {
			ob1 := new(Multi)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Am = ob0
	}
	if field, ok := in.Map()["ass"]; ok && field.Type() != system.J_NULL {
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
		v.Ass = ob0
	}
	if field, ok := in.Map()["mnri"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]system.StringInterface{}
		for k0 := range field.Map() {
			ob1, err := system.UnpackStringInterface(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mnri = ob0
	}
	if field, ok := in.Map()["alms"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Alms)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Alms = ob0
	}
	if field, ok := in.Map()["alajs"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Alajs)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Alajs = ob0
	}
	if field, ok := in.Map()["aljsi"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackAljsInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Aljsi = ob0
	}
	if field, ok := in.Map()["mm"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Multi{}
		for k0 := range field.Map() {
			ob1 := new(Multi)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mm = ob0
	}
	if field, ok := in.Map()["js"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Js = ob0
	}
	if field, ok := in.Map()["aalass"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Alass{}
		for i0 := range field.Array() {
			ob1 := *new(Alass)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalass = ob0
	}
	if field, ok := in.Map()["si"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Si = ob0
	}
	if field, ok := in.Map()["asr"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.Reference{}
		for i0 := range field.Array() {
			ob1 := new(system.Reference)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Asr = ob0
	}
	if field, ok := in.Map()["maljn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Aljn{}
		for k0 := range field.Map() {
			ob1 := new(Aljn)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Maljn = ob0
	}
	if field, ok := in.Map()["ai"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Face{}
		for i0 := range field.Array() {
			ob1, err := UnpackFace(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Ai = ob0
	}
	if field, ok := in.Map()["jb"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Jb = ob0
	}
	if field, ok := in.Map()["msp"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.Package{}
		for k0 := range field.Map() {
			ob1 := new(system.Package)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Msp = ob0
	}
	if field, ok := in.Map()["sb"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Sb = ob0
	}
	if field, ok := in.Map()["m"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Multi)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.M = ob0
	}
	if field, ok := in.Map()["asi"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.Int{}
		for i0 := range field.Array() {
			ob1 := new(system.Int)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Asi = ob0
	}
	if field, ok := in.Map()["malmjs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Almjs{}
		for k0 := range field.Map() {
			ob1 := *new(Almjs)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malmjs = ob0
	}
	if field, ok := in.Map()["sp"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Package)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Sp = ob0
	}
	if field, ok := in.Map()["aljb"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Aljb)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Aljb = ob0
	}
	if field, ok := in.Map()["aalmss"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Almss{}
		for i0 := range field.Array() {
			ob1 := *new(Almss)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalmss = ob0
	}
	if field, ok := in.Map()["aaljs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Aljs{}
		for i0 := range field.Array() {
			ob1 := new(Aljs)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aaljs = ob0
	}
	if field, ok := in.Map()["almjs"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Almjs)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Almjs = ob0
	}
	if field, ok := in.Map()["ajs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []string{}
		for i0 := range field.Array() {
			ob1, err := system.UnpackString(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Ajs = ob0
	}
	if field, ok := in.Map()["als"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Als)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Als = ob0
	}
	if field, ok := in.Map()["aalajs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Alajs{}
		for i0 := range field.Array() {
			ob1 := *new(Alajs)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalajs = ob0
	}
	if field, ok := in.Map()["malas"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Alas{}
		for k0 := range field.Map() {
			ob1 := *new(Alas)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malas = ob0
	}
	if field, ok := in.Map()["mri"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]system.RuleInterface{}
		for k0 := range field.Map() {
			ob1, err := system.UnpackRuleInterface(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mri = ob0
	}
	if field, ok := in.Map()["aljn"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Aljn)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Aljn = ob0
	}
	if field, ok := in.Map()["ajn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []float64{}
		for i0 := range field.Array() {
			ob1, err := system.UnpackNumber(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Ajn = ob0
	}
	if field, ok := in.Map()["msi"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.Int{}
		for k0 := range field.Map() {
			ob1 := new(system.Int)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Msi = ob0
	}
	if field, ok := in.Map()["ajb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []bool{}
		for i0 := range field.Array() {
			ob1, err := system.UnpackBool(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Ajb = ob0
	}
	if field, ok := in.Map()["mals"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Als{}
		for k0 := range field.Map() {
			ob1 := new(Als)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mals = ob0
	}
	if field, ok := in.Map()["almss"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Almss)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Almss = ob0
	}
	if field, ok := in.Map()["aalss"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Alss{}
		for i0 := range field.Array() {
			ob1 := new(Alss)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalss = ob0
	}
	if field, ok := in.Map()["aaljb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Aljb{}
		for i0 := range field.Array() {
			ob1 := new(Aljb)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aaljb = ob0
	}
	if field, ok := in.Map()["asb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.Bool{}
		for i0 := range field.Array() {
			ob1 := new(system.Bool)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Asb = ob0
	}
	if field, ok := in.Map()["aljs"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Aljs)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Aljs = ob0
	}
	if field, ok := in.Map()["malass"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Alass{}
		for k0 := range field.Map() {
			ob1 := *new(Alass)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malass = ob0
	}
	if field, ok := in.Map()["asn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.Number{}
		for i0 := range field.Array() {
			ob1 := new(system.Number)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Asn = ob0
	}
	if field, ok := in.Map()["bri"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackBoolInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Bri = ob0
	}
	if field, ok := in.Map()["asp"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*system.Package{}
		for i0 := range field.Array() {
			ob1 := new(system.Package)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Asp = ob0
	}
	if field, ok := in.Map()["malss"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Alss{}
		for k0 := range field.Map() {
			ob1 := new(Alss)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malss = ob0
	}
	if field, ok := in.Map()["aalmjs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Almjs{}
		for i0 := range field.Array() {
			ob1 := *new(Almjs)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalmjs = ob0
	}
	if field, ok := in.Map()["aalas"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Alas{}
		for i0 := range field.Array() {
			ob1 := *new(Alas)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalas = ob0
	}
	if field, ok := in.Map()["malmss"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Almss{}
		for k0 := range field.Map() {
			ob1 := *new(Almss)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Malmss = ob0
	}
	if field, ok := in.Map()["aaljn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Aljn{}
		for i0 := range field.Array() {
			ob1 := new(Aljn)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aaljn = ob0
	}
	if field, ok := in.Map()["mi"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Face{}
		for k0 := range field.Map() {
			ob1, err := UnpackFace(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mi = ob0
	}
	if field, ok := in.Map()["mjb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]bool{}
		for k0 := range field.Map() {
			ob1, err := system.UnpackBool(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mjb = ob0
	}
	if field, ok := in.Map()["msn"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.Number{}
		for k0 := range field.Map() {
			ob1 := new(system.Number)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Msn = ob0
	}
	if field, ok := in.Map()["ari"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []system.RuleInterface{}
		for i0 := range field.Array() {
			ob1, err := system.UnpackRuleInterface(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Ari = ob0
	}
	if field, ok := in.Map()["alss"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Alss)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Alss = ob0
	}
	if field, ok := in.Map()["aljni"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackAljnInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Aljni = ob0
	}
	if field, ok := in.Map()["alas"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Alas)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Alas = ob0
	}
	if field, ok := in.Map()["sri"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackStringInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Sri = ob0
	}
	if field, ok := in.Map()["alass"]; ok && field.Type() != system.J_NULL {
		ob0 := *new(Alass)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Alass = ob0
	}
	if field, ok := in.Map()["sr"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Sr = ob0
	}
	if field, ok := in.Map()["ss"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Ss = ob0
	}
	if field, ok := in.Map()["aalms"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []Alms{}
		for i0 := range field.Array() {
			ob1 := *new(Alms)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Aalms = ob0
	}
	if field, ok := in.Map()["maljb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Aljb{}
		for k0 := range field.Map() {
			ob1 := new(Aljb)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Maljb = ob0
	}
	if field, ok := in.Map()["mss"]; ok && field.Type() != system.J_NULL {
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
		v.Mss = ob0
	}
	if field, ok := in.Map()["mjs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]string{}
		for k0 := range field.Map() {
			ob1, err := system.UnpackString(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Mjs = ob0
	}
	if field, ok := in.Map()["maljs"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Aljs{}
		for k0 := range field.Map() {
			ob1 := new(Aljs)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Maljs = ob0
	}
	if field, ok := in.Map()["ri"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Ri = ob0
	}
	if field, ok := in.Map()["msb"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*system.Bool{}
		for k0 := range field.Map() {
			ob1 := new(system.Bool)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Msb = ob0
	}
	return nil
}
func (v *Multi) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "multi", system.J_NULL, nil
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
	if v.Alss != nil {
		ob0, _, _, _, err := v.Alss.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["alss"] = ob0
	}
	if v.Aljni != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Aljni.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "aljn") {
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
		m["aljni"] = ob0
	}
	if v.Alas != nil {
		ob0, _, _, _, err := v.Alas.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["alas"] = ob0
	}
	if v.Msn != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Msn {
			ob1, _, _, _, err := v.Msn[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["msn"] = ob0
	}
	if v.Ari != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ari {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Ari[i0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0 = append(ob0, ob1)
		}
		m["ari"] = ob0
	}
	if v.Ss != nil {
		ob0, _, _, _, err := v.Ss.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["ss"] = ob0
	}
	if v.Aalms != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalms {
			ob1, _, _, _, err := v.Aalms[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalms"] = ob0
	}
	if v.Maljb != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Maljb {
			ob1, _, _, _, err := v.Maljb[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["maljb"] = ob0
	}
	if v.Sri != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Sri.(system.Repacker).Repack(ctx)
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
		m["sri"] = ob0
	}
	if v.Alass != nil {
		ob0, _, _, _, err := v.Alass.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["alass"] = ob0
	}
	if v.Sr != nil {
		ob0, _, _, _, err := v.Sr.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["sr"] = ob0
	}
	if v.Maljs != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Maljs {
			ob1, _, _, _, err := v.Maljs[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["maljs"] = ob0
	}
	if v.Ri != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Ri.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
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
		m["ri"] = ob0
	}
	if v.Msb != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Msb {
			ob1, _, _, _, err := v.Msb[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["msb"] = ob0
	}
	if v.Mss != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mss {
			ob1, _, _, _, err := v.Mss[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["mss"] = ob0
	}
	if v.Mjs != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mjs {
			ob1 := v.Mjs[k0]
			ob0[k0] = ob1
		}
		m["mjs"] = ob0
	}
	if v.I != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.I.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "face") {
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
		m["i"] = ob0
	}
	if v.Anri != nil {
		ob0 := []interface{}{}
		for i0 := range v.Anri {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Anri[i0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "string") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0 = append(ob0, ob1)
		}
		m["anri"] = ob0
	}
	if v.Mjn != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mjn {
			ob1 := v.Mjn[k0]
			ob0[k0] = ob1
		}
		m["mjn"] = ob0
	}
	if v.Msr != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Msr {
			ob1, _, _, _, err := v.Msr[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["msr"] = ob0
	}
	if v.Aals != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aals {
			ob1, _, _, _, err := v.Aals[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aals"] = ob0
	}
	if v.Sn != nil {
		ob0, _, _, _, err := v.Sn.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["sn"] = ob0
	}
	if v.Nri != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Nri.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "number") {
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
		m["nri"] = ob0
	}
	if v.Malms != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malms {
			ob1, _, _, _, err := v.Malms[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malms"] = ob0
	}
	if v.Ass != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ass {
			ob1, _, _, _, err := v.Ass[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["ass"] = ob0
	}
	if v.Mnri != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mnri {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Mnri[k0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "string") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0[k0] = ob1
		}
		m["mnri"] = ob0
	}
	if v.Alms != nil {
		ob0, _, _, _, err := v.Alms.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["alms"] = ob0
	}
	if v.Alajs != nil {
		ob0, _, _, _, err := v.Alajs.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["alajs"] = ob0
	}
	if v.Aljbi != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Aljbi.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "aljb") {
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
		m["aljbi"] = ob0
	}
	if v.Jn != 0.0 {
		ob0 := v.Jn
		m["jn"] = ob0
	}
	if v.Malajs != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malajs {
			ob1, _, _, _, err := v.Malajs[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malajs"] = ob0
	}
	if v.Am != nil {
		ob0 := []interface{}{}
		for i0 := range v.Am {
			ob1, _, _, _, err := v.Am[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["am"] = ob0
	}
	if v.Aalass != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalass {
			ob1, _, _, _, err := v.Aalass[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalass"] = ob0
	}
	if v.Si != nil {
		ob0, _, _, _, err := v.Si.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["si"] = ob0
	}
	if v.Asr != nil {
		ob0 := []interface{}{}
		for i0 := range v.Asr {
			ob1, _, _, _, err := v.Asr[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["asr"] = ob0
	}
	if v.Aljsi != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Aljsi.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "aljs") {
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
		m["aljsi"] = ob0
	}
	if v.Mm != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mm {
			ob1, _, _, _, err := v.Mm[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["mm"] = ob0
	}
	if v.Js != "" {
		ob0 := v.Js
		m["js"] = ob0
	}
	if v.Maljn != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Maljn {
			ob1, _, _, _, err := v.Maljn[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["maljn"] = ob0
	}
	if v.Msp != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Msp {
			ob1, _, _, _, err := v.Msp[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["msp"] = ob0
	}
	if v.Sb != nil {
		ob0, _, _, _, err := v.Sb.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["sb"] = ob0
	}
	if v.M != nil {
		ob0, _, _, _, err := v.M.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["m"] = ob0
	}
	if v.Ai != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ai {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Ai[i0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "face") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0 = append(ob0, ob1)
		}
		m["ai"] = ob0
	}
	if v.Jb != false {
		ob0 := v.Jb
		m["jb"] = ob0
	}
	if v.Malmjs != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malmjs {
			ob1, _, _, _, err := v.Malmjs[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malmjs"] = ob0
	}
	if v.Sp != nil {
		ob0, _, _, _, err := v.Sp.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["sp"] = ob0
	}
	if v.Aljb != nil {
		ob0, _, _, _, err := v.Aljb.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["aljb"] = ob0
	}
	if v.Asi != nil {
		ob0 := []interface{}{}
		for i0 := range v.Asi {
			ob1, _, _, _, err := v.Asi[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["asi"] = ob0
	}
	if v.Aaljs != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aaljs {
			ob1, _, _, _, err := v.Aaljs[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aaljs"] = ob0
	}
	if v.Almjs != nil {
		ob0, _, _, _, err := v.Almjs.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["almjs"] = ob0
	}
	if v.Ajs != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ajs {
			ob1 := v.Ajs[i0]
			ob0 = append(ob0, ob1)
		}
		m["ajs"] = ob0
	}
	if v.Aalmss != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalmss {
			ob1, _, _, _, err := v.Aalmss[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalmss"] = ob0
	}
	if v.Aljn != nil {
		ob0, _, _, _, err := v.Aljn.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["aljn"] = ob0
	}
	if v.Ajn != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ajn {
			ob1 := v.Ajn[i0]
			ob0 = append(ob0, ob1)
		}
		m["ajn"] = ob0
	}
	if v.Msi != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Msi {
			ob1, _, _, _, err := v.Msi[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["msi"] = ob0
	}
	if v.Ajb != nil {
		ob0 := []interface{}{}
		for i0 := range v.Ajb {
			ob1 := v.Ajb[i0]
			ob0 = append(ob0, ob1)
		}
		m["ajb"] = ob0
	}
	if v.Als != nil {
		ob0, _, _, _, err := v.Als.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["als"] = ob0
	}
	if v.Aalajs != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalajs {
			ob1, _, _, _, err := v.Aalajs[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalajs"] = ob0
	}
	if v.Malas != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malas {
			ob1, _, _, _, err := v.Malas[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malas"] = ob0
	}
	if v.Mri != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mri {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Mri[k0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0[k0] = ob1
		}
		m["mri"] = ob0
	}
	if v.Asb != nil {
		ob0 := []interface{}{}
		for i0 := range v.Asb {
			ob1, _, _, _, err := v.Asb[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["asb"] = ob0
	}
	if v.Aljs != nil {
		ob0, _, _, _, err := v.Aljs.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["aljs"] = ob0
	}
	if v.Malass != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malass {
			ob1, _, _, _, err := v.Malass[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malass"] = ob0
	}
	if v.Asn != nil {
		ob0 := []interface{}{}
		for i0 := range v.Asn {
			ob1, _, _, _, err := v.Asn[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["asn"] = ob0
	}
	if v.Mals != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mals {
			ob1, _, _, _, err := v.Mals[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["mals"] = ob0
	}
	if v.Almss != nil {
		ob0, _, _, _, err := v.Almss.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["almss"] = ob0
	}
	if v.Aalss != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalss {
			ob1, _, _, _, err := v.Aalss[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalss"] = ob0
	}
	if v.Aaljb != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aaljb {
			ob1, _, _, _, err := v.Aaljb[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aaljb"] = ob0
	}
	if v.Bri != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Bri.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "bool") {
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
		m["bri"] = ob0
	}
	if v.Asp != nil {
		ob0 := []interface{}{}
		for i0 := range v.Asp {
			ob1, _, _, _, err := v.Asp[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["asp"] = ob0
	}
	if v.Malss != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malss {
			ob1, _, _, _, err := v.Malss[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malss"] = ob0
	}
	if v.Aalas != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalas {
			ob1, _, _, _, err := v.Aalas[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalas"] = ob0
	}
	if v.Malmss != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Malmss {
			ob1, _, _, _, err := v.Malmss[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["malmss"] = ob0
	}
	if v.Aaljn != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aaljn {
			ob1, _, _, _, err := v.Aaljn[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aaljn"] = ob0
	}
	if v.Aalmjs != nil {
		ob0 := []interface{}{}
		for i0 := range v.Aalmjs {
			ob1, _, _, _, err := v.Aalmjs[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["aalmjs"] = ob0
	}
	if v.Mi != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mi {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Mi[k0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/tests/data", "face") {
				typRef := system.NewReference(pkg, name)
				typeVal, err := typRef.ValueContext(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1 = map[string]interface{}{
					"type":  typeVal,
					"value": ob1_value,
				}
			} else {
				ob1 = ob1_value
			}
			ob0[k0] = ob1
		}
		m["mi"] = ob0
	}
	if v.Mjb != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Mjb {
			ob1 := v.Mjb[k0]
			ob0[k0] = ob1
		}
		m["mjb"] = ob0
	}
	return m, "kego.io/tests/data", "multi", system.J_OBJECT, nil
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
func UnpackSimpleInterface(ctx context.Context, in system.Packed) (SimpleInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "kego.io/tests/data", "simple")
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
	if field, ok := in.Map()["js"]; ok && field.Type() != system.J_NULL {
		ob0, err := system.UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Js = ob0
	}
	return nil
}
func (v *Simple) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "kego.io/tests/data", "simple", system.J_NULL, nil
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
	if v.Js != "" {
		ob0 := v.Js
		m["js"] = ob0
	}
	return m, "kego.io/tests/data", "simple", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/tests/data")
	pkg.SetHash(6167257871666940584)
	pkg.Init("alajs", func() interface{} { return *new(Alajs) }, func() interface{} { return new(AlajsRule) }, func() reflect.Type { return reflect.TypeOf((*AlajsInterface)(nil)).Elem() })
	pkg.Init("alas", func() interface{} { return *new(Alas) }, func() interface{} { return new(AlasRule) }, func() reflect.Type { return reflect.TypeOf((*AlasInterface)(nil)).Elem() })
	pkg.Init("alass", func() interface{} { return *new(Alass) }, func() interface{} { return new(AlassRule) }, func() reflect.Type { return reflect.TypeOf((*AlassInterface)(nil)).Elem() })
	pkg.Init("aljb", func() interface{} { return new(Aljb) }, func() interface{} { return new(AljbRule) }, func() reflect.Type { return reflect.TypeOf((*AljbInterface)(nil)).Elem() })
	pkg.Init("aljn", func() interface{} { return new(Aljn) }, func() interface{} { return new(AljnRule) }, func() reflect.Type { return reflect.TypeOf((*AljnInterface)(nil)).Elem() })
	pkg.Init("aljs", func() interface{} { return new(Aljs) }, func() interface{} { return new(AljsRule) }, func() reflect.Type { return reflect.TypeOf((*AljsInterface)(nil)).Elem() })
	pkg.Init("almjs", func() interface{} { return *new(Almjs) }, func() interface{} { return new(AlmjsRule) }, func() reflect.Type { return reflect.TypeOf((*AlmjsInterface)(nil)).Elem() })
	pkg.Init("alms", func() interface{} { return *new(Alms) }, func() interface{} { return new(AlmsRule) }, func() reflect.Type { return reflect.TypeOf((*AlmsInterface)(nil)).Elem() })
	pkg.Init("almss", func() interface{} { return *new(Almss) }, func() interface{} { return new(AlmssRule) }, func() reflect.Type { return reflect.TypeOf((*AlmssInterface)(nil)).Elem() })
	pkg.Init("als", func() interface{} { return new(Als) }, func() interface{} { return new(AlsRule) }, func() reflect.Type { return reflect.TypeOf((*AlsInterface)(nil)).Elem() })
	pkg.Init("alss", func() interface{} { return new(Alss) }, func() interface{} { return new(AlssRule) }, func() reflect.Type { return reflect.TypeOf((*AlssInterface)(nil)).Elem() })
	pkg.Init("face", func() interface{} { return (*Face)(nil) }, func() interface{} { return new(FaceRule) }, func() reflect.Type { return reflect.TypeOf((*Face)(nil)).Elem() })
	pkg.Init("facea", func() interface{} { return new(Facea) }, func() interface{} { return new(FaceaRule) }, func() reflect.Type { return reflect.TypeOf((*FaceaInterface)(nil)).Elem() })
	pkg.Init("faceb", func() interface{} { return new(Faceb) }, func() interface{} { return new(FacebRule) }, func() reflect.Type { return reflect.TypeOf((*FacebInterface)(nil)).Elem() })
	pkg.Init("multi", func() interface{} { return new(Multi) }, func() interface{} { return new(MultiRule) }, func() reflect.Type { return reflect.TypeOf((*MultiInterface)(nil)).Elem() })
	pkg.Init("simple", func() interface{} { return new(Simple) }, func() interface{} { return new(SimpleRule) }, func() reflect.Type { return reflect.TypeOf((*SimpleInterface)(nil)).Elem() })
}
