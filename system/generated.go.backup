// info:{"Path":"kego.io/system","Hash":5319817629068713650}
package system

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"
	"reflect"

	"kego.io/context/jsonctx"
)

// Restriction rules for arrays
type ArrayRule struct {
	*Object
	*Rule
	// This is a rule object, defining the type and restrictions on the value of the items
	Items RuleInterface `json:"items"`
	// This is the maximum number of items allowed in the array
	MaxItems *Int `json:"max-items"`
	// This is the minimum number of items allowed in the array
	MinItems *Int `json:"min-items"`
	// If this is true, each item must be unique
	UniqueItems bool `json:"unique-items"`
}

func (v *ArrayRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Items = ob0
	}
	if field, ok := in.Map()["max-items"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxItems = ob0
	}
	if field, ok := in.Map()["min-items"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinItems = ob0
	}
	if field, ok := in.Map()["unique-items"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.UniqueItems = ob0
	}
	return nil
}
func (v *ArrayRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@array", J_NULL, nil
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
	if v.Items != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Items.(Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
			typRef := NewReference(pkg, name)
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
		m["items"] = ob0
	}
	if v.MaxItems != nil {
		ob0, _, _, _, err := v.MaxItems.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["max-items"] = ob0
	}
	if v.MinItems != nil {
		ob0, _, _, _, err := v.MinItems.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["min-items"] = ob0
	}
	if v.UniqueItems != false {
		ob0 := v.UniqueItems
		m["unique-items"] = ob0
	}
	return m, "kego.io/system", "@array", J_OBJECT, nil
}

// Restriction rules for bools
type BoolRule struct {
	*Object
	*Rule
	// Default value if this is missing or null
	Default *Bool `json:"default"`
}

func (v *BoolRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob0 := new(Bool)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob0
	}
	return nil
}
func (v *BoolRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@bool", J_NULL, nil
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
	if v.Default != nil {
		ob0, _, _, _, err := v.Default.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["default"] = ob0
	}
	return m, "kego.io/system", "@bool", J_OBJECT, nil
}

// Restriction rules for integers
type IntRule struct {
	*Object
	*Rule
	// Default value if this property is omitted
	Default *Int `json:"default"`
	// This provides an upper bound for the restriction
	Maximum *Int `json:"maximum"`
	// This provides a lower bound for the restriction
	Minimum *Int `json:"minimum"`
	// This restricts the number to be a multiple of the given number
	MultipleOf *Int `json:"multiple-of"`
}

func (v *IntRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob0
	}
	if field, ok := in.Map()["maximum"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Maximum = ob0
	}
	if field, ok := in.Map()["minimum"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Minimum = ob0
	}
	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MultipleOf = ob0
	}
	return nil
}
func (v *IntRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@int", J_NULL, nil
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
	if v.Default != nil {
		ob0, _, _, _, err := v.Default.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["default"] = ob0
	}
	if v.Maximum != nil {
		ob0, _, _, _, err := v.Maximum.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["maximum"] = ob0
	}
	if v.Minimum != nil {
		ob0, _, _, _, err := v.Minimum.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["minimum"] = ob0
	}
	if v.MultipleOf != nil {
		ob0, _, _, _, err := v.MultipleOf.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["multiple-of"] = ob0
	}
	return m, "kego.io/system", "@int", J_OBJECT, nil
}

// Restriction rules for maps
type MapRule struct {
	*Object
	*Rule
	// This is a rule object, defining the type and restrictions on the value of the items.
	Items RuleInterface `json:"items"`
	// Add a system:@string here to provide a restriction for the map keys
	Keys RuleInterface `json:"keys"`
	// This is the maximum number of items alowed in the array
	MaxItems *Int `json:"max-items"`
	// This is the minimum number of items alowed in the array
	MinItems *Int `json:"min-items"`
}

func (v *MapRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Items = ob0
	}
	if field, ok := in.Map()["keys"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Keys = ob0
	}
	if field, ok := in.Map()["max-items"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxItems = ob0
	}
	if field, ok := in.Map()["min-items"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinItems = ob0
	}
	return nil
}
func (v *MapRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@map", J_NULL, nil
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
	if v.Items != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Items.(Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
			typRef := NewReference(pkg, name)
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
		m["items"] = ob0
	}
	if v.Keys != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Keys.(Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
			typRef := NewReference(pkg, name)
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
		m["keys"] = ob0
	}
	if v.MaxItems != nil {
		ob0, _, _, _, err := v.MaxItems.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["max-items"] = ob0
	}
	if v.MinItems != nil {
		ob0, _, _, _, err := v.MinItems.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["min-items"] = ob0
	}
	return m, "kego.io/system", "@map", J_OBJECT, nil
}

// Restriction rules for numbers
type NumberRule struct {
	*Object
	*Rule
	// Default value if this property is omitted
	Default *Number `json:"default"`
	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum bool `json:"exclusive-maximum"`
	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum bool `json:"exclusive-minimum"`
	// This provides an upper bound for the restriction
	Maximum *Number `json:"maximum"`
	// This provides a lower bound for the restriction
	Minimum *Number `json:"minimum"`
	// This restricts the number to be a multiple of the given number
	MultipleOf *Number `json:"multiple-of"`
}

func (v *NumberRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob0 := new(Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob0
	}
	if field, ok := in.Map()["exclusive-maximum"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.ExclusiveMaximum = ob0
	}
	if field, ok := in.Map()["exclusive-minimum"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.ExclusiveMinimum = ob0
	}
	if field, ok := in.Map()["maximum"]; ok && field.Type() != J_NULL {
		ob0 := new(Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Maximum = ob0
	}
	if field, ok := in.Map()["minimum"]; ok && field.Type() != J_NULL {
		ob0 := new(Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Minimum = ob0
	}
	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != J_NULL {
		ob0 := new(Number)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MultipleOf = ob0
	}
	return nil
}
func (v *NumberRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@number", J_NULL, nil
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
	if v.Default != nil {
		ob0, _, _, _, err := v.Default.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["default"] = ob0
	}
	if v.ExclusiveMaximum != false {
		ob0 := v.ExclusiveMaximum
		m["exclusive-maximum"] = ob0
	}
	if v.ExclusiveMinimum != false {
		ob0 := v.ExclusiveMinimum
		m["exclusive-minimum"] = ob0
	}
	if v.Maximum != nil {
		ob0, _, _, _, err := v.Maximum.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["maximum"] = ob0
	}
	if v.Minimum != nil {
		ob0, _, _, _, err := v.Minimum.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["minimum"] = ob0
	}
	if v.MultipleOf != nil {
		ob0, _, _, _, err := v.MultipleOf.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["multiple-of"] = ob0
	}
	return m, "kego.io/system", "@number", J_OBJECT, nil
}

// Automatically created basic rule for object
type ObjectRule struct {
	*Object
	*Rule
}

func (v *ObjectRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *ObjectRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@object", J_NULL, nil
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
	return m, "kego.io/system", "@object", J_OBJECT, nil
}

// Automatically created basic rule for package
type PackageRule struct {
	*Object
	*Rule
}

func (v *PackageRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *PackageRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@package", J_NULL, nil
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
	return m, "kego.io/system", "@package", J_OBJECT, nil
}

// Restriction rules for references
type ReferenceRule struct {
	*Object
	*Rule
	// Default value of this is missing or null
	Default *Reference `json:"default"`
	// The value must match this regex
	Pattern *String `json:"pattern"`
	// The value must not match this regex
	PatternNot *String `json:"pattern-not"`
}

func (v *ReferenceRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob0 := new(Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob0
	}
	if field, ok := in.Map()["pattern"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Pattern = ob0
	}
	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.PatternNot = ob0
	}
	return nil
}
func (v *ReferenceRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@reference", J_NULL, nil
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
	if v.Default != nil {
		ob0, _, _, _, err := v.Default.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["default"] = ob0
	}
	if v.Pattern != nil {
		ob0, _, _, _, err := v.Pattern.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["pattern"] = ob0
	}
	if v.PatternNot != nil {
		ob0, _, _, _, err := v.PatternNot.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["pattern-not"] = ob0
	}
	return m, "kego.io/system", "@reference", J_OBJECT, nil
}

// Automatically created basic rule for rule
type RuleRule struct {
	*Object
	*Rule
}

func (v *RuleRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *RuleRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@rule", J_NULL, nil
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
	return m, "kego.io/system", "@rule", J_OBJECT, nil
}

// Restriction rules for strings
type StringRule struct {
	*Object
	*Rule
	// Default value of this is missing or null
	Default *String `json:"default"`
	// The value of this string is restricted to one of the provided values
	Enum []string `json:"enum"`
	// This is a string that the value must match
	Equal *String `json:"equal"`
	// This restricts the value to one of several built-in formats
	Format *String `json:"format"`
	// The editor should render as a multi-line textbox
	Long bool `json:"long"`
	// The value must be shorter or equal to the provided maximum length
	MaxLength *Int `json:"max-length"`
	// The value must be longer or equal to the provided minimum length
	MinLength *Int `json:"min-length"`
	// The value must match this regex
	Pattern *String `json:"pattern"`
	// The value must not match this regex
	PatternNot *String `json:"pattern-not"`
}

func (v *StringRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob0
	}
	if field, ok := in.Map()["enum"]; ok && field.Type() != J_NULL {
		if field.Type() != J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []string{}
		for i0 := range field.Array() {
			ob1, err := UnpackString(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Enum = ob0
	}
	if field, ok := in.Map()["equal"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Equal = ob0
	}
	if field, ok := in.Map()["format"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Format = ob0
	}
	if field, ok := in.Map()["long"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Long = ob0
	}
	if field, ok := in.Map()["max-length"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxLength = ob0
	}
	if field, ok := in.Map()["min-length"]; ok && field.Type() != J_NULL {
		ob0 := new(Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinLength = ob0
	}
	if field, ok := in.Map()["pattern"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Pattern = ob0
	}
	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.PatternNot = ob0
	}
	return nil
}
func (v *StringRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@string", J_NULL, nil
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
	if v.Default != nil {
		ob0, _, _, _, err := v.Default.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["default"] = ob0
	}
	if v.Enum != nil {
		ob0 := []interface{}{}
		for i0 := range v.Enum {
			ob1 := v.Enum[i0]
			ob0 = append(ob0, ob1)
		}
		m["enum"] = ob0
	}
	if v.Equal != nil {
		ob0, _, _, _, err := v.Equal.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["equal"] = ob0
	}
	if v.Format != nil {
		ob0, _, _, _, err := v.Format.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["format"] = ob0
	}
	if v.Long != false {
		ob0 := v.Long
		m["long"] = ob0
	}
	if v.MaxLength != nil {
		ob0, _, _, _, err := v.MaxLength.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["max-length"] = ob0
	}
	if v.MinLength != nil {
		ob0, _, _, _, err := v.MinLength.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["min-length"] = ob0
	}
	if v.Pattern != nil {
		ob0, _, _, _, err := v.Pattern.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["pattern"] = ob0
	}
	if v.PatternNot != nil {
		ob0, _, _, _, err := v.PatternNot.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["pattern-not"] = ob0
	}
	return m, "kego.io/system", "@string", J_OBJECT, nil
}

// Automatically created basic rule for tags
type TagsRule struct {
	*Object
	*Rule
}

func (v *TagsRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *TagsRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@tags", J_NULL, nil
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
	return m, "kego.io/system", "@tags", J_OBJECT, nil
}

// Automatically created basic rule for type
type TypeRule struct {
	*Object
	*Rule
}

func (v *TypeRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *TypeRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "@type", J_NULL, nil
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
	return m, "kego.io/system", "@type", J_OBJECT, nil
}

// This is the native json array data type
type Array struct {
	*Object
}

// This is the native json bool data type
type Bool bool

func (o *Bool) Value() bool {
	return bool(*o)
}

type BoolInterface interface {
	GetBool(ctx context.Context) *Bool
}

func (o *Bool) GetBool(ctx context.Context) *Bool {
	return o
}
func UnpackBoolInterface(ctx context.Context, in Packed) (BoolInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "bool")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(BoolInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement BoolInterface", i)
		}
		return ob, nil
	case J_BOOL:
		ob := new(Bool)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into BoolInterface.", in.Type())
	}
}
func (v *Bool) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if in.Type() == J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != J_BOOL {
		return fmt.Errorf("Invalid type %s while unpacking a bool.", in.Type())
	}
	*v = Bool(in.Bool())
	return nil
}
func (v *Bool) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "bool", J_NULL, nil
	}
	if v != nil {
		return bool(*v), "kego.io/system", "bool", J_BOOL, nil
	}
	return nil, "kego.io/system", "bool", J_BOOL, nil
}

type IntInterface interface {
	GetInt(ctx context.Context) *Int
}

func (o *Int) GetInt(ctx context.Context) *Int {
	return o
}
func UnpackIntInterface(ctx context.Context, in Packed) (IntInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "int")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(IntInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement IntInterface", i)
		}
		return ob, nil
	case J_NUMBER:
		ob := new(Int)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into IntInterface.", in.Type())
	}
}

// This is the native json object data type.
type Map struct {
	*Object
}

// This is the native json number data type
type Number float64

func (o *Number) Value() float64 {
	return float64(*o)
}

type NumberInterface interface {
	GetNumber(ctx context.Context) *Number
}

func (o *Number) GetNumber(ctx context.Context) *Number {
	return o
}
func UnpackNumberInterface(ctx context.Context, in Packed) (NumberInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "number")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(NumberInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement NumberInterface", i)
		}
		return ob, nil
	case J_NUMBER:
		ob := new(Number)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into NumberInterface.", in.Type())
	}
}
func (v *Number) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if in.Type() == J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != J_NUMBER {
		return fmt.Errorf("Invalid type %s while unpacking a number.", in.Type())
	}
	*v = Number(in.Number())
	return nil
}
func (v *Number) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "number", J_NULL, nil
	}
	if v != nil {
		return float64(*v), "kego.io/system", "number", J_NUMBER, nil
	}
	return nil, "kego.io/system", "number", J_NUMBER, nil
}

// This is the base type for the object interface. All ke objects have this type embedded.
type Object struct {
	// Description for the developer
	Description string `json:"description"`
	// All global objects should have an id.
	Id *Reference `json:"id"`
	// Extra validation rules for this object or descendants
	Rules []RuleInterface `json:"rules"`
	// Tags for general use
	Tags []string `json:"tags"`
	// Tags for general use
	TagsNew Tags `json:"tags-new"`
	// Type of the object.
	Type *Reference `json:"type"`
}
type ObjectInterface interface {
	GetObject(ctx context.Context) *Object
}

func (o *Object) GetObject(ctx context.Context) *Object {
	return o
}
func UnpackObjectInterface(ctx context.Context, in Packed) (ObjectInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "object")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ObjectInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ObjectInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into ObjectInterface.", in.Type())
	}
}
func (v *Object) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if field, ok := in.Map()["description"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Description = ob0
	}
	if field, ok := in.Map()["id"]; ok && field.Type() != J_NULL {
		ob0 := new(Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Id = ob0
	}
	if field, ok := in.Map()["rules"]; ok && field.Type() != J_NULL {
		if field.Type() != J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []RuleInterface{}
		for i0 := range field.Array() {
			ob1, err := UnpackRuleInterface(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Rules = ob0
	}
	if field, ok := in.Map()["tags"]; ok && field.Type() != J_NULL {
		if field.Type() != J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []string{}
		for i0 := range field.Array() {
			ob1, err := UnpackString(ctx, field.Array()[i0])
			if err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Tags = ob0
	}
	if field, ok := in.Map()["tags-new"]; ok && field.Type() != J_NULL {
		ob0 := *new(Tags)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.TagsNew = ob0
	}
	if field, ok := in.Map()["type"]; ok && field.Type() != J_NULL {
		ob0 := new(Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Type = ob0
	}
	return nil
}
func (v *Object) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "object", J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Description != "" {
		ob0 := v.Description
		m["description"] = ob0
	}
	if v.Id != nil {
		ob0, _, _, _, err := v.Id.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["id"] = ob0
	}
	if v.Rules != nil {
		ob0 := []interface{}{}
		for i0 := range v.Rules {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Rules[i0].(Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
				typRef := NewReference(pkg, name)
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
		m["rules"] = ob0
	}
	if v.Tags != nil {
		ob0 := []interface{}{}
		for i0 := range v.Tags {
			ob1 := v.Tags[i0]
			ob0 = append(ob0, ob1)
		}
		m["tags"] = ob0
	}
	if v.TagsNew != nil {
		ob0, _, _, _, err := v.TagsNew.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["tags-new"] = ob0
	}
	if v.Type != nil {
		ob0, _, _, _, err := v.Type.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["type"] = ob0
	}
	return m, "kego.io/system", "object", J_OBJECT, nil
}

// Package info - forms the root node of the package
type Package struct {
	*Object
	// Map of import aliases used in this package: key = alias, value = package path.
	Aliases map[string]string `json:"aliases"`
	// Should we scan subdirectories for data files?
	Recursive bool `json:"recursive"`
}
type PackageInterface interface {
	GetPackage(ctx context.Context) *Package
}

func (o *Package) GetPackage(ctx context.Context) *Package {
	return o
}
func UnpackPackageInterface(ctx context.Context, in Packed) (PackageInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "package")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PackageInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PackageInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PackageInterface.", in.Type())
	}
}
func (v *Package) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["aliases"]; ok && field.Type() != J_NULL {
		if field.Type() != J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]string{}
		for k0 := range field.Map() {
			ob1, err := UnpackString(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Aliases = ob0
	}
	if field, ok := in.Map()["recursive"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Recursive = ob0
	}
	return nil
}
func (v *Package) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "package", J_NULL, nil
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
	if v.Aliases != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Aliases {
			ob1 := v.Aliases[k0]
			ob0[k0] = ob1
		}
		m["aliases"] = ob0
	}
	if v.Recursive != false {
		ob0 := v.Recursive
		m["recursive"] = ob0
	}
	return m, "kego.io/system", "package", J_OBJECT, nil
}

type ReferenceInterface interface {
	GetReference(ctx context.Context) *Reference
}

func (o *Reference) GetReference(ctx context.Context) *Reference {
	return o
}
func UnpackReferenceInterface(ctx context.Context, in Packed) (ReferenceInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "reference")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ReferenceInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ReferenceInterface", i)
		}
		return ob, nil
	case J_STRING:
		ob := new(Reference)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into ReferenceInterface.", in.Type())
	}
}

// All rules will have this embedded in them.
type Rule struct {
	// Use the single method getter interface for this type
	Interface bool `json:"interface"`
	// If this rule is a field, this specifies that the field is optional
	Optional bool `json:"optional"`
	// Json selector defining what nodes this rule should be applied to.
	Selector string `json:"selector"`
}
type RuleInterface interface {
	GetRule(ctx context.Context) *Rule
}

func (o *Rule) GetRule(ctx context.Context) *Rule {
	return o
}
func UnpackRuleInterface(ctx context.Context, in Packed) (RuleInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "rule")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(RuleInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement RuleInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into RuleInterface.", in.Type())
	}
}
func (v *Rule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if field, ok := in.Map()["interface"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Interface = ob0
	}
	if field, ok := in.Map()["optional"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Optional = ob0
	}
	if field, ok := in.Map()["selector"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Selector = ob0
	}
	return nil
}
func (v *Rule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "rule", J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Interface != false {
		ob0 := v.Interface
		m["interface"] = ob0
	}
	if v.Optional != false {
		ob0 := v.Optional
		m["optional"] = ob0
	}
	if v.Selector != "" {
		ob0 := v.Selector
		m["selector"] = ob0
	}
	return m, "kego.io/system", "rule", J_OBJECT, nil
}

// This is the native json string data type
type String string

func (o *String) Value() string {
	return string(*o)
}

type StringInterface interface {
	GetString(ctx context.Context) *String
}

func (o *String) GetString(ctx context.Context) *String {
	return o
}
func UnpackStringInterface(ctx context.Context, in Packed) (StringInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "string")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(StringInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement StringInterface", i)
		}
		return ob, nil
	case J_STRING:
		ob := new(String)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into StringInterface.", in.Type())
	}
}
func (v *String) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if in.Type() == J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != J_STRING {
		return fmt.Errorf("Invalid type %s while unpacking a string.", in.Type())
	}
	*v = String(in.String())
	return nil
}
func (v *String) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "string", J_NULL, nil
	}
	if v != nil {
		return string(*v), "kego.io/system", "string", J_STRING, nil
	}
	return nil, "kego.io/system", "string", J_STRING, nil
}

// Tag cloud.
type Tags []string
type TagsInterface interface {
	GetTags(ctx context.Context) Tags
}

func (o Tags) GetTags(ctx context.Context) Tags {
	return o
}
func UnpackTagsInterface(ctx context.Context, in Packed) (TagsInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "tags")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(TagsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement TagsInterface", i)
		}
		return ob, nil
	case J_ARRAY:
		ob := new(Tags)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into TagsInterface.", in.Type())
	}
}
func (v *Tags) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if in.Type() == J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != J_ARRAY {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	if in.Type() != J_ARRAY {
		return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", in.Type())
	}
	ob0 := []string{}
	for i0 := range in.Array() {
		ob1, err := UnpackString(ctx, in.Array()[i0])
		if err != nil {
			return err
		}
		ob0 = append(ob0, ob1)
	}
	*v = ob0
	return nil
}
func (v Tags) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "tags", J_NULL, nil
	}
	ob0 := []interface{}{}
	for i0 := range v {
		ob1 := v[i0]
		ob0 = append(ob0, ob1)
	}
	return ob0, "kego.io/system", "tags", J_ARRAY, nil
}

// This is the most basic type.
type Type struct {
	*Object
	// If this type is an alias of another type, specify the underlying type here
	Alias RuleInterface `json:"alias"`
	// Basic types don't have system:object added by default to the embedded types.
	Basic bool `json:"basic"`
	// Custom types are not emitted into the generated source
	Custom bool `json:"custom"`
	// The kind of the type if custom is specified - must be value, struct, collection or interface
	CustomKind *String `json:"custom-kind"`
	// Types which this should embed - system:object is always added unless basic = true.
	Embed []*Reference `json:"embed"`
	// Each field is listed with it's type
	Fields map[string]RuleInterface `json:"fields"`
	// Is this type an interface?
	Interface bool `json:"interface"`
	// This is the native json type that represents this type. If omitted, default is object.
	Native *String `json:"native"`
	// Type that defines restriction rules for this type.
	Rule *Type `json:"rule"`
}
type TypeInterface interface {
	GetType(ctx context.Context) *Type
}

func (o *Type) GetType(ctx context.Context) *Type {
	return o
}
func UnpackTypeInterface(ctx context.Context, in Packed) (TypeInterface, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "kego.io/system", "type")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(TypeInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement TypeInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into TypeInterface.", in.Type())
	}
}
func (v *Type) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["alias"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Alias = ob0
	}
	if field, ok := in.Map()["basic"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Basic = ob0
	}
	if field, ok := in.Map()["custom"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Custom = ob0
	}
	if field, ok := in.Map()["custom-kind"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.CustomKind = ob0
	}
	if field, ok := in.Map()["embed"]; ok && field.Type() != J_NULL {
		if field.Type() != J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		ob0 := []*Reference{}
		for i0 := range field.Array() {
			ob1 := new(Reference)
			if err := ob1.Unpack(ctx, field.Array()[i0], false); err != nil {
				return err
			}
			ob0 = append(ob0, ob1)
		}
		v.Embed = ob0
	}
	if field, ok := in.Map()["fields"]; ok && field.Type() != J_NULL {
		if field.Type() != J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		ob0 := map[string]RuleInterface{}
		for k0 := range field.Map() {
			ob1, err := UnpackRuleInterface(ctx, field.Map()[k0])
			if err != nil {
				return err
			}
			ob0[k0] = ob1
		}
		v.Fields = ob0
	}
	if field, ok := in.Map()["interface"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Interface = ob0
	}
	if field, ok := in.Map()["native"]; ok && field.Type() != J_NULL {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Native = ob0
	} else {
		ob0 := new(String)
		if err := ob0.Unpack(ctx, Pack("object"), false); err != nil {
			return err
		}
		v.Native = ob0
	}
	if field, ok := in.Map()["rule"]; ok && field.Type() != J_NULL {
		ob0 := new(Type)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Rule = ob0
	}
	return nil
}
func (v *Type) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if v == nil {
		return nil, "kego.io/system", "type", J_NULL, nil
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
	if v.Alias != nil {
		var ob0 interface{}
		ob0_value, pkg, name, typ, err := v.Alias.(Repacker).Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
			typRef := NewReference(pkg, name)
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
		m["alias"] = ob0
	}
	if v.Basic != false {
		ob0 := v.Basic
		m["basic"] = ob0
	}
	if v.Custom != false {
		ob0 := v.Custom
		m["custom"] = ob0
	}
	if v.CustomKind != nil {
		ob0, _, _, _, err := v.CustomKind.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["custom-kind"] = ob0
	}
	if v.Embed != nil {
		ob0 := []interface{}{}
		for i0 := range v.Embed {
			ob1, _, _, _, err := v.Embed[i0].Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			ob0 = append(ob0, ob1)
		}
		m["embed"] = ob0
	}
	if v.Fields != nil {
		ob0 := map[string]interface{}{}
		for k0 := range v.Fields {
			var ob1 interface{}
			ob1_value, pkg, name, typ, err := v.Fields[k0].(Repacker).Repack(ctx)
			if err != nil {
				return nil, "", "", "", err
			}
			if ShouldUseExplicitTypeNotation(pkg, name, typ, "kego.io/system", "rule") {
				typRef := NewReference(pkg, name)
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
		m["fields"] = ob0
	}
	if v.Interface != false {
		ob0 := v.Interface
		m["interface"] = ob0
	}
	if v.Native != nil {
		ob0, _, _, _, err := v.Native.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["native"] = ob0
	}
	if v.Rule != nil {
		ob0, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["rule"] = ob0
	}
	return m, "kego.io/system", "type", J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/system")
	pkg.SetHash(5319817629068713650)
	pkg.Init("array", nil, func() interface{} { return new(ArrayRule) }, nil)
	pkg.Init("bool", func() interface{} { return new(Bool) }, func() interface{} { return new(BoolRule) }, func() reflect.Type { return reflect.TypeOf((*BoolInterface)(nil)).Elem() })
	pkg.Init("int", func() interface{} { return new(Int) }, func() interface{} { return new(IntRule) }, func() reflect.Type { return reflect.TypeOf((*IntInterface)(nil)).Elem() })
	pkg.Init("map", nil, func() interface{} { return new(MapRule) }, nil)
	pkg.Init("number", func() interface{} { return new(Number) }, func() interface{} { return new(NumberRule) }, func() reflect.Type { return reflect.TypeOf((*NumberInterface)(nil)).Elem() })
	pkg.Init("object", func() interface{} { return new(Object) }, func() interface{} { return new(ObjectRule) }, func() reflect.Type { return reflect.TypeOf((*ObjectInterface)(nil)).Elem() })
	pkg.Init("package", func() interface{} { return new(Package) }, func() interface{} { return new(PackageRule) }, func() reflect.Type { return reflect.TypeOf((*PackageInterface)(nil)).Elem() })
	pkg.Init("reference", func() interface{} { return new(Reference) }, func() interface{} { return new(ReferenceRule) }, func() reflect.Type { return reflect.TypeOf((*ReferenceInterface)(nil)).Elem() })
	pkg.Init("rule", func() interface{} { return new(Rule) }, func() interface{} { return new(RuleRule) }, func() reflect.Type { return reflect.TypeOf((*RuleInterface)(nil)).Elem() })
	pkg.Init("string", func() interface{} { return new(String) }, func() interface{} { return new(StringRule) }, func() reflect.Type { return reflect.TypeOf((*StringInterface)(nil)).Elem() })
	pkg.Init("tags", func() interface{} { return *new(Tags) }, func() interface{} { return new(TagsRule) }, func() reflect.Type { return reflect.TypeOf((*TagsInterface)(nil)).Elem() })
	pkg.Init("type", func() interface{} { return new(Type) }, func() interface{} { return new(TypeRule) }, func() reflect.Type { return reflect.TypeOf((*TypeInterface)(nil)).Elem() })
}
