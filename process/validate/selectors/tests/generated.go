// info:{"Path":"frizz.io/process/validate/selectors/tests","Hash":958437053934508214}
package tests

import (
	context "context"
	fmt "fmt"
	reflect "reflect"

	jsonctx "frizz.io/context/jsonctx"
	system "frizz.io/system"
)

// notest

// Automatically created basic rule for basic
type BasicRule struct {
	*system.Object
	*system.Rule
}

func (v *BasicRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@basic"); err != nil {
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
func (v *BasicRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@basic", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@basic", system.J_OBJECT, nil
}

// Automatically created basic rule for c
type CRule struct {
	*system.Object
	*system.Rule
}

func (v *CRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@c"); err != nil {
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
func (v *CRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@c", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@c", system.J_OBJECT, nil
}

// Automatically created basic rule for collision
type CollisionRule struct {
	*system.Object
	*system.Rule
}

func (v *CollisionRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@collision"); err != nil {
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
func (v *CollisionRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@collision", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@collision", system.J_OBJECT, nil
}

// Automatically created basic rule for diagram
type DiagramRule struct {
	*system.Object
	*system.Rule
}

func (v *DiagramRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@diagram"); err != nil {
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
func (v *DiagramRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@diagram", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@diagram", system.J_OBJECT, nil
}

// Automatically created basic rule for empty
type EmptyRule struct {
	*system.Object
	*system.Rule
}

func (v *EmptyRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@empty"); err != nil {
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
func (v *EmptyRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@empty", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@empty", system.J_OBJECT, nil
}

// Automatically created basic rule for emptyItem
type EmptyItemRule struct {
	*system.Object
	*system.Rule
}

func (v *EmptyItemRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@emptyItem"); err != nil {
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
func (v *EmptyItemRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@emptyItem", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@emptyItem", system.J_OBJECT, nil
}

// Automatically created basic rule for expr
type ExprRule struct {
	*system.Object
	*system.Rule
}

func (v *ExprRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@expr"); err != nil {
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
func (v *ExprRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@expr", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@expr", system.J_OBJECT, nil
}

// Automatically created basic rule for gallery
type GalleryRule struct {
	*system.Object
	*system.Rule
}

func (v *GalleryRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@gallery"); err != nil {
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
func (v *GalleryRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@gallery", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@gallery", system.J_OBJECT, nil
}

// Automatically created basic rule for image
type ImageRule struct {
	*system.Object
	*system.Rule
}

func (v *ImageRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@image"); err != nil {
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
func (v *ImageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@image", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@image", system.J_OBJECT, nil
}

// Automatically created basic rule for instance
type InstanceRule struct {
	*system.Object
	*system.Rule
}

func (v *InstanceRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@instance"); err != nil {
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
func (v *InstanceRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@instance", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@instance", system.J_OBJECT, nil
}

// Automatically created basic rule for instanceItem
type InstanceItemRule struct {
	*system.Object
	*system.Rule
}

func (v *InstanceItemRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@instanceItem"); err != nil {
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
func (v *InstanceItemRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@instanceItem", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@instanceItem", system.J_OBJECT, nil
}

// Automatically created basic rule for kid
type KidRule struct {
	*system.Object
	*system.Rule
}

func (v *KidRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@kid"); err != nil {
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
func (v *KidRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@kid", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@kid", system.J_OBJECT, nil
}

// Automatically created basic rule for people
type PeopleRule struct {
	*system.Object
	*system.Rule
}

func (v *PeopleRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@people"); err != nil {
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
func (v *PeopleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@people", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@people", system.J_OBJECT, nil
}

// Automatically created basic rule for person
type PersonRule struct {
	*system.Object
	*system.Rule
}

func (v *PersonRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@person"); err != nil {
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
func (v *PersonRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@person", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@person", system.J_OBJECT, nil
}

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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@photo"); err != nil {
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
		return nil, "frizz.io/process/validate/selectors/tests", "@photo", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@photo", system.J_OBJECT, nil
}

// Automatically created basic rule for polykids
type PolykidsRule struct {
	*system.Object
	*system.Rule
}

func (v *PolykidsRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@polykids"); err != nil {
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
func (v *PolykidsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@polykids", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@polykids", system.J_OBJECT, nil
}

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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@rectangle"); err != nil {
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
		return nil, "frizz.io/process/validate/selectors/tests", "@rectangle", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@rectangle", system.J_OBJECT, nil
}

// Automatically created basic rule for rightscale
type RightscaleRule struct {
	*system.Object
	*system.Rule
}

func (v *RightscaleRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@rightscale"); err != nil {
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
func (v *RightscaleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@rightscale", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@rightscale", system.J_OBJECT, nil
}

// Automatically created basic rule for rightscaleLink
type RightscaleLinkRule struct {
	*system.Object
	*system.Rule
}

func (v *RightscaleLinkRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@rightscaleLink"); err != nil {
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
func (v *RightscaleLinkRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@rightscaleLink", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@rightscaleLink", system.J_OBJECT, nil
}

// Automatically created basic rule for rightscaleList
type RightscaleListRule struct {
	*system.Object
	*system.Rule
}

func (v *RightscaleListRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@rightscaleList"); err != nil {
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
func (v *RightscaleListRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@rightscaleList", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@rightscaleList", system.J_OBJECT, nil
}

// Automatically created basic rule for sibling
type SiblingRule struct {
	*system.Object
	*system.Rule
}

func (v *SiblingRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@sibling"); err != nil {
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
func (v *SiblingRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@sibling", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@sibling", system.J_OBJECT, nil
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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@simple"); err != nil {
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
		return nil, "frizz.io/process/validate/selectors/tests", "@simple", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@simple", system.J_OBJECT, nil
}

// Automatically created basic rule for simpleArray
type SimpleArrayRule struct {
	*system.Object
	*system.Rule
}

func (v *SimpleArrayRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@simpleArray"); err != nil {
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
func (v *SimpleArrayRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@simpleArray", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@simpleArray", system.J_OBJECT, nil
}

// Automatically created basic rule for simpleItem
type SimpleItemRule struct {
	*system.Object
	*system.Rule
}

func (v *SimpleItemRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@simpleItem"); err != nil {
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
func (v *SimpleItemRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@simpleItem", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@simpleItem", system.J_OBJECT, nil
}

// Automatically created basic rule for typed
type TypedRule struct {
	*system.Object
	*system.Rule
}

func (v *TypedRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "@typed"); err != nil {
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
func (v *TypedRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "@typed", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "@typed", system.J_OBJECT, nil
}

type Basic struct {
	*system.Object
	DrinkPreference   []*system.String            `json:"drinkPreference"`
	FavoriteColor     *system.String              `json:"favoriteColor"`
	LanguagesSpoken   []map[string]*system.String `json:"languagesSpoken"`
	Name              map[string]*system.String   `json:"name"`
	SeatingPreference []*system.String            `json:"seatingPreference"`
	Weight            *system.Number              `json:"weight"`
}
type BasicInterface interface {
	GetBasic(ctx context.Context) *Basic
}

func (o *Basic) GetBasic(ctx context.Context) *Basic {
	return o
}
func UnpackBasicInterface(ctx context.Context, in system.Packed) (BasicInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "basic")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(BasicInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement BasicInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into BasicInterface.", in.Type())
	}
}
func (v *Basic) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "basic"); err != nil {
		return err
	}
	if field, ok := in.Map()["drinkPreference"]; ok && field.Type() != system.J_NULL {
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
		v.DrinkPreference = ob0
	}
	if field, ok := in.Map()["favoriteColor"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.FavoriteColor = ob0
	}
	if field, ok := in.Map()["languagesSpoken"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []map[string]*system.String{}
		for i0 := range field.Array() {
			if field.Array()[i0].Type() != system.J_MAP {
				return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Array()[i0].Type())
			}
			ob1 := map[string]*system.String{}
			for k1 := range field.Array()[i0].Map() {
				ob2 := new(system.String)
				if err := ob2.Unpack(ctx, field.Array()[i0].Map()[k1], false); err != nil {
					return err
				}
				ob1[k1] = ob2
			}
			ob0 = append(ob0, ob1)
		}
		v.LanguagesSpoken = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
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
		v.Name = ob0
	}
	if field, ok := in.Map()["seatingPreference"]; ok && field.Type() != system.J_NULL {
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
		v.SeatingPreference = ob0
	}
	if field, ok := in.Map()["weight"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Weight = ob0
	}
	return nil
}
func (v *Basic) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "basic", system.J_NULL, nil
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
	if v.DrinkPreference != nil {
		ob0 := []interface{}{}
		for i0 := range v.DrinkPreference {
			ob1, _, _, _, err := v.DrinkPreference[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["drinkPreference"] = ob0
	}
	if v.FavoriteColor != nil {
		ob0, _, _, _, err := v.FavoriteColor.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["favoriteColor"] = ob0
	}
	if v.LanguagesSpoken != nil {
		ob0 := []interface{}{}
		for i0 := range v.LanguagesSpoken {
			ob1 := map[string]interface{}{}
			for k1 := range v.LanguagesSpoken[i0] {
				ob2, _, _, _, err := v.LanguagesSpoken[i0][k1].Repack(ctx)
				if err != nil {
					return nil, "", "", "", err
				}
				ob1[k1] = ob2
			}
			ob0 = append(ob0, ob1)
		}
		m["languagesSpoken"] = ob0
	}
	if v.Name != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Name {
			ob1, _, _, _, err := v.Name[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["name"] = ob0
	}
	if v.SeatingPreference != nil {
		ob0 := []interface{}{}
		for i0 := range v.SeatingPreference {
			ob1, _, _, _, err := v.SeatingPreference[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["seatingPreference"] = ob0
	}
	if v.Weight != nil {
		ob0, _, _, _, err := v.Weight.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["weight"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "basic", system.J_OBJECT, nil
}

type C struct {
	*system.Object
	A *system.Number            `json:"a"`
	B *system.Number            `json:"b"`
	C map[string]*system.Number `json:"c"`
}
type CInterface interface {
	GetC(ctx context.Context) *C
}

func (o *C) GetC(ctx context.Context) *C {
	return o
}
func UnpackCInterface(ctx context.Context, in system.Packed) (CInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "c")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(CInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement CInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into CInterface.", in.Type())
	}
}
func (v *C) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "c"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.B = ob0
	}
	if field, ok := in.Map()["c"]; ok && field.Type() != system.J_NULL {
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
		v.C = ob0
	}
	return nil
}
func (v *C) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "c", system.J_NULL, nil
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
	if v.B != nil {
		ob0, _, _, _, err := v.B.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["b"] = ob0
	}
	if v.C != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.C {
			ob1, _, _, _, err := v.C[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["c"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "c", system.J_OBJECT, nil
}

type Collision struct {
	*system.Object
	Number map[string]*system.String `json:"number"`
}
type CollisionInterface interface {
	GetCollision(ctx context.Context) *Collision
}

func (o *Collision) GetCollision(ctx context.Context) *Collision {
	return o
}
func UnpackCollisionInterface(ctx context.Context, in system.Packed) (CollisionInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "collision")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(CollisionInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement CollisionInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into CollisionInterface.", in.Type())
	}
}
func (v *Collision) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "collision"); err != nil {
		return err
	}
	if field, ok := in.Map()["number"]; ok && field.Type() != system.J_NULL {
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
		v.Number = ob0
	}
	return nil
}
func (v *Collision) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "collision", system.J_NULL, nil
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
	if v.Number != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Number {
			ob1, _, _, _, err := v.Number[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["number"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "collision", system.J_OBJECT, nil
}

type Diagram struct {
	*system.Object
	Url *system.String `json:"url"`
}
type DiagramInterface interface {
	GetDiagram(ctx context.Context) *Diagram
}

func (o *Diagram) GetDiagram(ctx context.Context) *Diagram {
	return o
}
func UnpackDiagramInterface(ctx context.Context, in system.Packed) (DiagramInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "diagram")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(DiagramInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement DiagramInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into DiagramInterface.", in.Type())
	}
}
func (v *Diagram) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "diagram"); err != nil {
		return err
	}
	if field, ok := in.Map()["url"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Url = ob0
	}
	return nil
}
func (v *Diagram) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "diagram", system.J_NULL, nil
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
	if v.Url != nil {
		ob0, _, _, _, err := v.Url.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["url"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "diagram", system.J_OBJECT, nil
}

type Empty struct {
	*system.Object
	Items []*EmptyItem `json:"items"`
}
type EmptyInterface interface {
	GetEmpty(ctx context.Context) *Empty
}

func (o *Empty) GetEmpty(ctx context.Context) *Empty {
	return o
}
func UnpackEmptyInterface(ctx context.Context, in system.Packed) (EmptyInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "empty")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(EmptyInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement EmptyInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into EmptyInterface.", in.Type())
	}
}
func (v *Empty) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "empty"); err != nil {
		return err
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*EmptyItem{}
		for i0 := range field.Array() {
			ob1 := new(EmptyItem)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Items = ob0
	}
	return nil
}
func (v *Empty) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "empty", system.J_NULL, nil
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
	if v.Items != nil {
		ob0 := []interface{}{}
		for i0 := range v.Items {
			ob1, _, _, _, err := v.Items[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["items"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "empty", system.J_OBJECT, nil
}

type EmptyItem struct {
	*system.Object
	Arr  []*system.String `json:"arr"`
	Name *system.String   `json:"name"`
}
type EmptyItemInterface interface {
	GetEmptyItem(ctx context.Context) *EmptyItem
}

func (o *EmptyItem) GetEmptyItem(ctx context.Context) *EmptyItem {
	return o
}
func UnpackEmptyItemInterface(ctx context.Context, in system.Packed) (EmptyItemInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "emptyItem")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(EmptyItemInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement EmptyItemInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into EmptyItemInterface.", in.Type())
	}
}
func (v *EmptyItem) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "emptyItem"); err != nil {
		return err
	}
	if field, ok := in.Map()["arr"]; ok && field.Type() != system.J_NULL {
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
		v.Arr = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *EmptyItem) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "emptyItem", system.J_NULL, nil
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
	if v.Arr != nil {
		ob0 := []interface{}{}
		for i0 := range v.Arr {
			ob1, _, _, _, err := v.Arr[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["arr"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "emptyItem", system.J_OBJECT, nil
}

type Expr struct {
	*system.Object
	False   *system.Bool   `json:"false"`
	Float   *system.Number `json:"float"`
	Int     *system.Number `json:"int"`
	Null    *system.String `json:"null"`
	String  *system.String `json:"string"`
	String2 *system.String `json:"string2"`
	True    *system.Bool   `json:"true"`
}
type ExprInterface interface {
	GetExpr(ctx context.Context) *Expr
}

func (o *Expr) GetExpr(ctx context.Context) *Expr {
	return o
}
func UnpackExprInterface(ctx context.Context, in system.Packed) (ExprInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "expr")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ExprInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ExprInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into ExprInterface.", in.Type())
	}
}
func (v *Expr) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "expr"); err != nil {
		return err
	}
	if field, ok := in.Map()["false"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.False = ob0
	}
	if field, ok := in.Map()["float"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Float = ob0
	}
	if field, ok := in.Map()["int"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Int = ob0
	}
	if field, ok := in.Map()["null"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Null = ob0
	}
	if field, ok := in.Map()["string"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.String = ob0
	}
	if field, ok := in.Map()["string2"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.String2 = ob0
	}
	if field, ok := in.Map()["true"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.True = ob0
	}
	return nil
}
func (v *Expr) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "expr", system.J_NULL, nil
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
	if v.False != nil {
		ob0, _, _, _, err := v.False.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["false"] = ob0
	}
	if v.Float != nil {
		ob0, _, _, _, err := v.Float.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["float"] = ob0
	}
	if v.Int != nil {
		ob0, _, _, _, err := v.Int.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["int"] = ob0
	}
	if v.Null != nil {
		ob0, _, _, _, err := v.Null.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["null"] = ob0
	}
	if v.String != nil {
		ob0, _, _, _, err := v.String.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["string"] = ob0
	}
	if v.String2 != nil {
		ob0, _, _, _, err := v.String2.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["string2"] = ob0
	}
	if v.True != nil {
		ob0, _, _, _, err := v.True.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["true"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "expr", system.J_OBJECT, nil
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Object
	Images map[string]Image `json:"images"`
}
type GalleryInterface interface {
	GetGallery(ctx context.Context) *Gallery
}

func (o *Gallery) GetGallery(ctx context.Context) *Gallery {
	return o
}
func UnpackGalleryInterface(ctx context.Context, in system.Packed) (GalleryInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "gallery")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(GalleryInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement GalleryInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into GalleryInterface.", in.Type())
	}
}
func (v *Gallery) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "gallery"); err != nil {
		return err
	}
	if field, ok := in.Map()["images"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]Image{}
		for k0 := range field.Map() {
			ob1, err := UnpackImage(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Images = ob0
	}
	return nil
}
func (v *Gallery) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "gallery", system.J_NULL, nil
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
	if v.Images != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Images {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Images[k0].(system.Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "frizz.io/process/validate/selectors/tests", "image") {
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
		m["images"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "gallery", system.J_OBJECT, nil
}
func UnpackImage(ctx context.Context, in system.Packed) (Image, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "image")
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

type Instance struct {
	*system.Object
	Child        map[string]*InstanceItem `json:"child"`
	Cloud_type   *system.String           `json:"cloud_type"`
	Display_name *system.String           `json:"display_name"`
	Links        []*RightscaleLink        `json:"links"`
	Name         *system.String           `json:"name"`
}
type InstanceInterface interface {
	GetInstance(ctx context.Context) *Instance
}

func (o *Instance) GetInstance(ctx context.Context) *Instance {
	return o
}
func UnpackInstanceInterface(ctx context.Context, in system.Packed) (InstanceInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "instance")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(InstanceInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement InstanceInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into InstanceInterface.", in.Type())
	}
}
func (v *Instance) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "instance"); err != nil {
		return err
	}
	if field, ok := in.Map()["child"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*InstanceItem{}
		for k0 := range field.Map() {
			ob1 := new(InstanceItem)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Child = ob0
	}
	if field, ok := in.Map()["cloud_type"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Cloud_type = ob0
	}
	if field, ok := in.Map()["display_name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Display_name = ob0
	}
	if field, ok := in.Map()["links"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*RightscaleLink{}
		for i0 := range field.Array() {
			ob1 := new(RightscaleLink)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Links = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *Instance) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "instance", system.J_NULL, nil
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
	if v.Child != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Child {
			ob1, _, _, _, err := v.Child[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["child"] = ob0
	}
	if v.Cloud_type != nil {
		ob0, _, _, _, err := v.Cloud_type.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["cloud_type"] = ob0
	}
	if v.Display_name != nil {
		ob0, _, _, _, err := v.Display_name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["display_name"] = ob0
	}
	if v.Links != nil {
		ob0 := []interface{}{}
		for i0 := range v.Links {
			ob1, _, _, _, err := v.Links[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["links"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "instance", system.J_OBJECT, nil
}

type InstanceItem struct {
	*system.Object
	Name *system.String `json:"name"`
}
type InstanceItemInterface interface {
	GetInstanceItem(ctx context.Context) *InstanceItem
}

func (o *InstanceItem) GetInstanceItem(ctx context.Context) *InstanceItem {
	return o
}
func UnpackInstanceItemInterface(ctx context.Context, in system.Packed) (InstanceItemInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "instanceItem")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(InstanceItemInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement InstanceItemInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into InstanceItemInterface.", in.Type())
	}
}
func (v *InstanceItem) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "instanceItem"); err != nil {
		return err
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *InstanceItem) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "instanceItem", system.J_NULL, nil
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
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "instanceItem", system.J_OBJECT, nil
}

type Kid struct {
	*system.Object
	Language  *system.String `json:"language"`
	Level     *system.String `json:"level"`
	Preferred *system.Bool   `json:"preferred"`
}
type KidInterface interface {
	GetKid(ctx context.Context) *Kid
}

func (o *Kid) GetKid(ctx context.Context) *Kid {
	return o
}
func UnpackKidInterface(ctx context.Context, in system.Packed) (KidInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "kid")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(KidInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement KidInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into KidInterface.", in.Type())
	}
}
func (v *Kid) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "kid"); err != nil {
		return err
	}
	if field, ok := in.Map()["language"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Language = ob0
	}
	if field, ok := in.Map()["level"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Level = ob0
	}
	if field, ok := in.Map()["preferred"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Preferred = ob0
	}
	return nil
}
func (v *Kid) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "kid", system.J_NULL, nil
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
	if v.Language != nil {
		ob0, _, _, _, err := v.Language.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["language"] = ob0
	}
	if v.Level != nil {
		ob0, _, _, _, err := v.Level.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["level"] = ob0
	}
	if v.Preferred != nil {
		ob0, _, _, _, err := v.Preferred.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["preferred"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "kid", system.J_OBJECT, nil
}

type People struct {
	*system.Object
	Others []*Person `json:"others"`
	People []*Person `json:"people"`
}
type PeopleInterface interface {
	GetPeople(ctx context.Context) *People
}

func (o *People) GetPeople(ctx context.Context) *People {
	return o
}
func UnpackPeopleInterface(ctx context.Context, in system.Packed) (PeopleInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "people")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PeopleInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PeopleInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PeopleInterface.", in.Type())
	}
}
func (v *People) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "people"); err != nil {
		return err
	}
	if field, ok := in.Map()["others"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Person{}
		for i0 := range field.Array() {
			ob1 := new(Person)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Others = ob0
	}
	if field, ok := in.Map()["people"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Person{}
		for i0 := range field.Array() {
			ob1 := new(Person)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.People = ob0
	}
	return nil
}
func (v *People) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "people", system.J_NULL, nil
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
	if v.Others != nil {
		ob0 := []interface{}{}
		for i0 := range v.Others {
			ob1, _, _, _, err := v.Others[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["others"] = ob0
	}
	if v.People != nil {
		ob0 := []interface{}{}
		for i0 := range v.People {
			ob1, _, _, _, err := v.People[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["people"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "people", system.J_OBJECT, nil
}

type Person struct {
	*system.Object
	Age    *system.Int    `json:"age"`
	Colour *system.String `json:"colour"`
	Hat    *system.Bool   `json:"hat"`
	Name   *system.String `json:"name"`
}
type PersonInterface interface {
	GetPerson(ctx context.Context) *Person
}

func (o *Person) GetPerson(ctx context.Context) *Person {
	return o
}
func UnpackPersonInterface(ctx context.Context, in system.Packed) (PersonInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "person")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PersonInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PersonInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PersonInterface.", in.Type())
	}
}
func (v *Person) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "person"); err != nil {
		return err
	}
	if field, ok := in.Map()["age"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Age = ob0
	}
	if field, ok := in.Map()["colour"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Colour = ob0
	}
	if field, ok := in.Map()["hat"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Hat = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *Person) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "person", system.J_NULL, nil
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
	if v.Age != nil {
		ob0, _, _, _, err := v.Age.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["age"] = ob0
	}
	if v.Colour != nil {
		ob0, _, _, _, err := v.Colour.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["colour"] = ob0
	}
	if v.Hat != nil {
		ob0, _, _, _, err := v.Hat.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["hat"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "person", system.J_OBJECT, nil
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Object
	// The path for the url - e.g. /foo/bar.jpg
	Path *system.String `json:"path"`
	// The protocol for the url - e.g. http or https
	Protocol *system.String `json:"protocol"`
	// The server for the url - e.g. www.google.com
	Server *system.String `json:"server"`
	Size   *Rectangle     `json:"size"`
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
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "photo")
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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "photo"); err != nil {
		return err
	}
	if field, ok := in.Map()["path"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Path = ob0
	}
	if field, ok := in.Map()["protocol"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Protocol = ob0
	} else {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, system.MustPackString("\"http\""), false); err != nil {
			return err
		}
		v.Protocol = ob0
	}
	if field, ok := in.Map()["server"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Server = ob0
	}
	if field, ok := in.Map()["size"]; ok && field.Type() != system.J_NULL {
		ob0 := new(Rectangle)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Size = ob0
	}
	return nil
}
func (v *Photo) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "photo", system.J_NULL, nil
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
	if v.Path != nil {
		ob0, _, _, _, err := v.Path.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["path"] = ob0
	}
	if v.Protocol != nil {
		ob0, _, _, _, err := v.Protocol.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["protocol"] = ob0
	}
	if v.Server != nil {
		ob0, _, _, _, err := v.Server.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["server"] = ob0
	}
	if v.Size != nil {
		ob0, _, _, _, err := v.Size.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["size"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "photo", system.J_OBJECT, nil
}

type Polykids struct {
	*system.Object
	A []*Kid `json:"a"`
}
type PolykidsInterface interface {
	GetPolykids(ctx context.Context) *Polykids
}

func (o *Polykids) GetPolykids(ctx context.Context) *Polykids {
	return o
}
func UnpackPolykidsInterface(ctx context.Context, in system.Packed) (PolykidsInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "polykids")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PolykidsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PolykidsInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PolykidsInterface.", in.Type())
	}
}
func (v *Polykids) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "polykids"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Kid{}
		for i0 := range field.Array() {
			ob1 := new(Kid)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.A = ob0
	}
	return nil
}
func (v *Polykids) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "polykids", system.J_NULL, nil
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
		ob0 := []interface{}{}
		for i0 := range v.A {
			ob1, _, _, _, err := v.A[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["a"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "polykids", system.J_OBJECT, nil
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
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "rectangle")
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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "rectangle"); err != nil {
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
		return nil, "frizz.io/process/validate/selectors/tests", "rectangle", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "rectangle", system.J_OBJECT, nil
}

type Rightscale struct {
	*system.Object
	Child        map[string]*InstanceItem `json:"child"`
	Cloud_type   *system.String           `json:"cloud_type"`
	Display_name *system.String           `json:"display_name"`
	Links        []*RightscaleLink        `json:"links"`
	Name         *system.String           `json:"name"`
}
type RightscaleInterface interface {
	GetRightscale(ctx context.Context) *Rightscale
}

func (o *Rightscale) GetRightscale(ctx context.Context) *Rightscale {
	return o
}
func UnpackRightscaleInterface(ctx context.Context, in system.Packed) (RightscaleInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "rightscale")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(RightscaleInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement RightscaleInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into RightscaleInterface.", in.Type())
	}
}
func (v *Rightscale) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "rightscale"); err != nil {
		return err
	}
	if field, ok := in.Map()["child"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*InstanceItem{}
		for k0 := range field.Map() {
			ob1 := new(InstanceItem)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Child = ob0
	}
	if field, ok := in.Map()["cloud_type"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Cloud_type = ob0
	}
	if field, ok := in.Map()["display_name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Display_name = ob0
	}
	if field, ok := in.Map()["links"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*RightscaleLink{}
		for i0 := range field.Array() {
			ob1 := new(RightscaleLink)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Links = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *Rightscale) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "rightscale", system.J_NULL, nil
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
	if v.Child != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Child {
			ob1, _, _, _, err := v.Child[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["child"] = ob0
	}
	if v.Cloud_type != nil {
		ob0, _, _, _, err := v.Cloud_type.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["cloud_type"] = ob0
	}
	if v.Display_name != nil {
		ob0, _, _, _, err := v.Display_name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["display_name"] = ob0
	}
	if v.Links != nil {
		ob0 := []interface{}{}
		for i0 := range v.Links {
			ob1, _, _, _, err := v.Links[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["links"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "rightscale", system.J_OBJECT, nil
}

type RightscaleLink struct {
	*system.Object
	Href *system.String `json:"href"`
	Rel  *system.String `json:"rel"`
}
type RightscaleLinkInterface interface {
	GetRightscaleLink(ctx context.Context) *RightscaleLink
}

func (o *RightscaleLink) GetRightscaleLink(ctx context.Context) *RightscaleLink {
	return o
}
func UnpackRightscaleLinkInterface(ctx context.Context, in system.Packed) (RightscaleLinkInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "rightscaleLink")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(RightscaleLinkInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement RightscaleLinkInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into RightscaleLinkInterface.", in.Type())
	}
}
func (v *RightscaleLink) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "rightscaleLink"); err != nil {
		return err
	}
	if field, ok := in.Map()["href"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Href = ob0
	}
	if field, ok := in.Map()["rel"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Rel = ob0
	}
	return nil
}
func (v *RightscaleLink) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "rightscaleLink", system.J_NULL, nil
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
	if v.Href != nil {
		ob0, _, _, _, err := v.Href.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["href"] = ob0
	}
	if v.Rel != nil {
		ob0, _, _, _, err := v.Rel.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["rel"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "rightscaleLink", system.J_OBJECT, nil
}

type RightscaleList struct {
	*system.Object
	Foo []*Rightscale `json:"foo"`
}
type RightscaleListInterface interface {
	GetRightscaleList(ctx context.Context) *RightscaleList
}

func (o *RightscaleList) GetRightscaleList(ctx context.Context) *RightscaleList {
	return o
}
func UnpackRightscaleListInterface(ctx context.Context, in system.Packed) (RightscaleListInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "rightscaleList")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(RightscaleListInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement RightscaleListInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into RightscaleListInterface.", in.Type())
	}
}
func (v *RightscaleList) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "rightscaleList"); err != nil {
		return err
	}
	if field, ok := in.Map()["foo"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Rightscale{}
		for i0 := range field.Array() {
			ob1 := new(Rightscale)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Foo = ob0
	}
	return nil
}
func (v *RightscaleList) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "rightscaleList", system.J_NULL, nil
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
	if v.Foo != nil {
		ob0 := []interface{}{}
		for i0 := range v.Foo {
			ob1, _, _, _, err := v.Foo[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["foo"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "rightscaleList", system.J_OBJECT, nil
}

type Sibling struct {
	*system.Object
	A *system.Number            `json:"a"`
	B *system.Number            `json:"b"`
	C *C                        `json:"c"`
	D map[string]*system.Number `json:"d"`
	E map[string]*system.Number `json:"e"`
}
type SiblingInterface interface {
	GetSibling(ctx context.Context) *Sibling
}

func (o *Sibling) GetSibling(ctx context.Context) *Sibling {
	return o
}
func UnpackSiblingInterface(ctx context.Context, in system.Packed) (SiblingInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "sibling")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(SiblingInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement SiblingInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into SiblingInterface.", in.Type())
	}
}
func (v *Sibling) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "sibling"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	if field, ok := in.Map()["b"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.B = ob0
	}
	if field, ok := in.Map()["c"]; ok && field.Type() != system.J_NULL {
		ob0 := new(C)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.C = ob0
	}
	if field, ok := in.Map()["d"]; ok && field.Type() != system.J_NULL {
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
		v.D = ob0
	}
	if field, ok := in.Map()["e"]; ok && field.Type() != system.J_NULL {
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
		v.E = ob0
	}
	return nil
}
func (v *Sibling) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "sibling", system.J_NULL, nil
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
	if v.B != nil {
		ob0, _, _, _, err := v.B.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["b"] = ob0
	}
	if v.C != nil {
		ob0, _, _, _, err := v.C.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["c"] = ob0
	}
	if v.D != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.D {
			ob1, _, _, _, err := v.D[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["d"] = ob0
	}
	if v.E != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.E {
			ob1, _, _, _, err := v.E[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["e"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "sibling", system.J_OBJECT, nil
}

type Simple struct {
	*system.Object
	A *SimpleItem `json:"a"`
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
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "simple")
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
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "simple"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		ob0 := new(SimpleItem)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.A = ob0
	}
	return nil
}
func (v *Simple) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "simple", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "simple", system.J_OBJECT, nil
}

type SimpleArray struct {
	*system.Object
	A []*SimpleItem `json:"a"`
}
type SimpleArrayInterface interface {
	GetSimpleArray(ctx context.Context) *SimpleArray
}

func (o *SimpleArray) GetSimpleArray(ctx context.Context) *SimpleArray {
	return o
}
func UnpackSimpleArrayInterface(ctx context.Context, in system.Packed) (SimpleArrayInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "simpleArray")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(SimpleArrayInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement SimpleArrayInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into SimpleArrayInterface.", in.Type())
	}
}
func (v *SimpleArray) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "simpleArray"); err != nil {
		return err
	}
	if field, ok := in.Map()["a"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*SimpleItem{}
		for i0 := range field.Array() {
			ob1 := new(SimpleItem)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.A = ob0
	}
	return nil
}
func (v *SimpleArray) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "simpleArray", system.J_NULL, nil
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
		ob0 := []interface{}{}
		for i0 := range v.A {
			ob1, _, _, _, err := v.A[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["a"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "simpleArray", system.J_OBJECT, nil
}

type SimpleItem struct {
	*system.Object
	B *system.String `json:"b"`
}
type SimpleItemInterface interface {
	GetSimpleItem(ctx context.Context) *SimpleItem
}

func (o *SimpleItem) GetSimpleItem(ctx context.Context) *SimpleItem {
	return o
}
func UnpackSimpleItemInterface(ctx context.Context, in system.Packed) (SimpleItemInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "simpleItem")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(SimpleItemInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement SimpleItemInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into SimpleItemInterface.", in.Type())
	}
}
func (v *SimpleItem) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "simpleItem"); err != nil {
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
func (v *SimpleItem) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "simpleItem", system.J_NULL, nil
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
	return m, "frizz.io/process/validate/selectors/tests", "simpleItem", system.J_OBJECT, nil
}

type Typed struct {
	*system.Object
	Avatar          Image                     `json:"avatar"`
	DrinkPreference []*system.String          `json:"drinkPreference"`
	FavoriteColor   *system.String            `json:"favoriteColor"`
	Kids            map[string]*Kid           `json:"kids"`
	Name            map[string]*system.String `json:"name"`
	Weight          *system.Number            `json:"weight"`
}
type TypedInterface interface {
	GetTyped(ctx context.Context) *Typed
}

func (o *Typed) GetTyped(ctx context.Context) *Typed {
	return o
}
func UnpackTypedInterface(ctx context.Context, in system.Packed) (TypedInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/process/validate/selectors/tests", "typed")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(TypedInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement TypedInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into TypedInterface.", in.Type())
	}
}
func (v *Typed) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/process/validate/selectors/tests", "typed"); err != nil {
		return err
	}
	if field, ok := in.Map()["avatar"]; ok && field.Type() != system.J_NULL {
		ob0, err := UnpackImage(ctx, field)
		if err != nil {
			return err
		}
		v.Avatar = ob0
	}
	if field, ok := in.Map()["drinkPreference"]; ok && field.Type() != system.J_NULL {
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
		v.DrinkPreference = ob0
	}
	if field, ok := in.Map()["favoriteColor"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.FavoriteColor = ob0
	}
	if field, ok := in.Map()["kids"]; ok && field.Type() != system.J_NULL {
		if field.Type() != system.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]*Kid{}
		for k0 := range field.Map() {
			ob1 := new(Kid)
			if err := ob1.Unpack(ctx, field.Map()[k0], false); err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Kids = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
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
		v.Name = ob0
	}
	if field, ok := in.Map()["weight"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Weight = ob0
	}
	return nil
}
func (v *Typed) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/process/validate/selectors/tests", "typed", system.J_NULL, nil
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
	if v.Avatar != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Avatar.(system.Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if system.ShouldUseExplicitTypeNotation(pkg, name, typ, "frizz.io/process/validate/selectors/tests", "image") {
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
		m["avatar"] = ob0
	}
	if v.DrinkPreference != nil {
		ob0 := []interface{}{}
		for i0 := range v.DrinkPreference {
			ob1, _, _, _, err := v.DrinkPreference[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["drinkPreference"] = ob0
	}
	if v.FavoriteColor != nil {
		ob0, _, _, _, err := v.FavoriteColor.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["favoriteColor"] = ob0
	}
	if v.Kids != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Kids {
			ob1, _, _, _, err := v.Kids[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["kids"] = ob0
	}
	if v.Name != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Name {
			ob1, _, _, _, err := v.Name[k0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0[k0] = ob1
		}
		m["name"] = ob0
	}
	if v.Weight != nil {
		ob0, _, _, _, err := v.Weight.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["weight"] = ob0
	}
	return m, "frizz.io/process/validate/selectors/tests", "typed", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("frizz.io/process/validate/selectors/tests")
	pkg.SetHash(uint64(0xd4d0d6c2b5698b6))
	pkg.Init("basic", func() interface{} {
		return new(Basic)
	}, nil, func() interface{} {
		return new(BasicRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*BasicInterface)(nil)).Elem()
	})
	pkg.Init("c", func() interface{} {
		return new(C)
	}, nil, func() interface{} {
		return new(CRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*CInterface)(nil)).Elem()
	})
	pkg.Init("collision", func() interface{} {
		return new(Collision)
	}, nil, func() interface{} {
		return new(CollisionRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*CollisionInterface)(nil)).Elem()
	})
	pkg.Init("diagram", func() interface{} {
		return new(Diagram)
	}, nil, func() interface{} {
		return new(DiagramRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*DiagramInterface)(nil)).Elem()
	})
	pkg.Init("empty", func() interface{} {
		return new(Empty)
	}, nil, func() interface{} {
		return new(EmptyRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*EmptyInterface)(nil)).Elem()
	})
	pkg.Init("emptyItem", func() interface{} {
		return new(EmptyItem)
	}, nil, func() interface{} {
		return new(EmptyItemRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*EmptyItemInterface)(nil)).Elem()
	})
	pkg.Init("expr", func() interface{} {
		return new(Expr)
	}, nil, func() interface{} {
		return new(ExprRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*ExprInterface)(nil)).Elem()
	})
	pkg.Init("gallery", func() interface{} {
		return new(Gallery)
	}, nil, func() interface{} {
		return new(GalleryRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*GalleryInterface)(nil)).Elem()
	})
	pkg.Init("image", func() interface{} {
		return (*Image)(nil)
	}, nil, func() interface{} {
		return new(ImageRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*Image)(nil)).Elem()
	})
	pkg.Init("instance", func() interface{} {
		return new(Instance)
	}, nil, func() interface{} {
		return new(InstanceRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*InstanceInterface)(nil)).Elem()
	})
	pkg.Init("instanceItem", func() interface{} {
		return new(InstanceItem)
	}, nil, func() interface{} {
		return new(InstanceItemRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*InstanceItemInterface)(nil)).Elem()
	})
	pkg.Init("kid", func() interface{} {
		return new(Kid)
	}, nil, func() interface{} {
		return new(KidRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*KidInterface)(nil)).Elem()
	})
	pkg.Init("people", func() interface{} {
		return new(People)
	}, nil, func() interface{} {
		return new(PeopleRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PeopleInterface)(nil)).Elem()
	})
	pkg.Init("person", func() interface{} {
		return new(Person)
	}, nil, func() interface{} {
		return new(PersonRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PersonInterface)(nil)).Elem()
	})
	pkg.Init("photo", func() interface{} {
		return new(Photo)
	}, nil, func() interface{} {
		return new(PhotoRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PhotoInterface)(nil)).Elem()
	})
	pkg.Init("polykids", func() interface{} {
		return new(Polykids)
	}, nil, func() interface{} {
		return new(PolykidsRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PolykidsInterface)(nil)).Elem()
	})
	pkg.Init("rectangle", func() interface{} {
		return new(Rectangle)
	}, nil, func() interface{} {
		return new(RectangleRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*RectangleInterface)(nil)).Elem()
	})
	pkg.Init("rightscale", func() interface{} {
		return new(Rightscale)
	}, nil, func() interface{} {
		return new(RightscaleRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*RightscaleInterface)(nil)).Elem()
	})
	pkg.Init("rightscaleLink", func() interface{} {
		return new(RightscaleLink)
	}, nil, func() interface{} {
		return new(RightscaleLinkRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*RightscaleLinkInterface)(nil)).Elem()
	})
	pkg.Init("rightscaleList", func() interface{} {
		return new(RightscaleList)
	}, nil, func() interface{} {
		return new(RightscaleListRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*RightscaleListInterface)(nil)).Elem()
	})
	pkg.Init("sibling", func() interface{} {
		return new(Sibling)
	}, nil, func() interface{} {
		return new(SiblingRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*SiblingInterface)(nil)).Elem()
	})
	pkg.Init("simple", func() interface{} {
		return new(Simple)
	}, nil, func() interface{} {
		return new(SimpleRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*SimpleInterface)(nil)).Elem()
	})
	pkg.Init("simpleArray", func() interface{} {
		return new(SimpleArray)
	}, nil, func() interface{} {
		return new(SimpleArrayRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*SimpleArrayInterface)(nil)).Elem()
	})
	pkg.Init("simpleItem", func() interface{} {
		return new(SimpleItem)
	}, nil, func() interface{} {
		return new(SimpleItemRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*SimpleItemInterface)(nil)).Elem()
	})
	pkg.Init("typed", func() interface{} {
		return new(Typed)
	}, nil, func() interface{} {
		return new(TypedRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*TypedInterface)(nil)).Elem()
	})
}
