// info:{"Path":"kego.io/system","Hash":5319817629068713650}
package system

// ke: {"file": {"notest": true}}

import (
	"context"
	"fmt"

	"kego.io/context/jsonctx"
	"kego.io/packer"
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

func (v *ArrayRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["unique-items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.UniqueItems = ob
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Items = ob
	}
	if field, ok := in.Map()["min-items"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinItems = ob
	}
	if field, ok := in.Map()["max-items"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxItems = ob
	}
	return nil
}

// Restriction rules for bools
type BoolRule struct {
	*Object
	*Rule
	// Default value if this is missing or null
	Default *Bool `json:"default"`
}

func (v *BoolRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		ob := new(Bool)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob
	}
	return nil
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

func (v *IntRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob
	}
	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MultipleOf = ob
	}
	if field, ok := in.Map()["minimum"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Minimum = ob
	}
	if field, ok := in.Map()["maximum"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Maximum = ob
	}
	return nil
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

func (v *MapRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["keys"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Keys = ob
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Items = ob
	}
	if field, ok := in.Map()["min-items"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinItems = ob
	}
	if field, ok := in.Map()["max-items"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxItems = ob
	}
	return nil
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

func (v *NumberRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		ob := new(Number)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob
	}
	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != packer.J_NULL {
		ob := new(Number)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MultipleOf = ob
	}
	if field, ok := in.Map()["minimum"]; ok && field.Type() != packer.J_NULL {
		ob := new(Number)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Minimum = ob
	}
	if field, ok := in.Map()["exclusive-minimum"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.ExclusiveMinimum = ob
	}
	if field, ok := in.Map()["maximum"]; ok && field.Type() != packer.J_NULL {
		ob := new(Number)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Maximum = ob
	}
	if field, ok := in.Map()["exclusive-maximum"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.ExclusiveMaximum = ob
	}
	return nil
}

// Automatically created basic rule for object
type ObjectRule struct {
	*Object
	*Rule
}

func (v *ObjectRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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

// Automatically created basic rule for package
type PackageRule struct {
	*Object
	*Rule
}

func (v *PackageRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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

func (v *ReferenceRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		ob := new(Reference)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob
	}
	if field, ok := in.Map()["pattern"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Pattern = ob
	}
	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.PatternNot = ob
	}
	return nil
}

// Automatically created basic rule for rule
type RuleRule struct {
	*Object
	*Rule
}

func (v *RuleRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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

func (v *StringRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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
	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Default = ob
	}
	if field, ok := in.Map()["pattern"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Pattern = ob
	}
	if field, ok := in.Map()["format"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Format = ob
	}
	if field, ok := in.Map()["min-length"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MinLength = ob
	}
	if field, ok := in.Map()["equal"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Equal = ob
	}
	if field, ok := in.Map()["enum"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		for _, field := range field.Array() {
			ob, err := UnpackString(ctx, field)
			if err != nil {
				return err
			}
			v.Enum = append(v.Enum, ob)
		}
	}
	if field, ok := in.Map()["max-length"]; ok && field.Type() != packer.J_NULL {
		ob := new(Int)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.MaxLength = ob
	}
	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.PatternNot = ob
	}
	if field, ok := in.Map()["long"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Long = ob
	}
	return nil
}

// Automatically created basic rule for tags
type TagsRule struct {
	*Object
	*Rule
}

func (v *TagsRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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

// Automatically created basic rule for type
type TypeRule struct {
	*Object
	*Rule
}

func (v *TypeRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
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

// This is the native json array data type
type Array struct {
	*Object
}
type ArrayInterface interface {
	GetArray(ctx context.Context) Array
}

func (o Array) GetArray(ctx context.Context) Array {
	return o
}
func UnpackArrayInterface(ctx context.Context, in packer.Packed) (ArrayInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ArrayInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ArrayInterface", i)
		}
		return ob, nil
	case packer.J_ARRAY:
		ob := new(Array)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into ArrayInterface.", in.Type())
	}
}

// This is the native json bool data type
type Bool bool
type BoolInterface interface {
	GetBool(ctx context.Context) *Bool
}

func (o *Bool) GetBool(ctx context.Context) *Bool {
	return o
}
func UnpackBoolInterface(ctx context.Context, in packer.Packed) (BoolInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(BoolInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement BoolInterface", i)
		}
		return ob, nil
	case packer.J_BOOL:
		ob := new(Bool)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into BoolInterface.", in.Type())
	}
}
func (v *Bool) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_BOOL {
		return fmt.Errorf("Invalid type %s while unpacking a bool.", in.Type())
	}
	*v = Bool(in.Bool())
	return nil
}

type IntInterface interface {
	GetInt(ctx context.Context) *Int
}

func (o *Int) GetInt(ctx context.Context) *Int {
	return o
}
func UnpackIntInterface(ctx context.Context, in packer.Packed) (IntInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(IntInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement IntInterface", i)
		}
		return ob, nil
	case packer.J_NUMBER:
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
type MapInterface interface {
	GetMap(ctx context.Context) Map
}

func (o Map) GetMap(ctx context.Context) Map {
	return o
}
func UnpackMapInterface(ctx context.Context, in packer.Packed) (MapInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(MapInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement MapInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into MapInterface.", in.Type())
	}
}

// This is the native json number data type
type Number float64
type NumberInterface interface {
	GetNumber(ctx context.Context) *Number
}

func (o *Number) GetNumber(ctx context.Context) *Number {
	return o
}
func UnpackNumberInterface(ctx context.Context, in packer.Packed) (NumberInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(NumberInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement NumberInterface", i)
		}
		return ob, nil
	case packer.J_NUMBER:
		ob := new(Number)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into NumberInterface.", in.Type())
	}
}
func (v *Number) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_NUMBER {
		return fmt.Errorf("Invalid type %s while unpacking a number.", in.Type())
	}
	*v = Number(in.Number())
	return nil
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
func UnpackObjectInterface(ctx context.Context, in packer.Packed) (ObjectInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
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
func (v *Object) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if field, ok := in.Map()["description"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Description = ob
	}
	if field, ok := in.Map()["rules"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		for _, field := range field.Array() {
			ob, err := UnpackRuleInterface(ctx, field)
			if err != nil {
				return err
			}
			v.Rules = append(v.Rules, ob)
		}
	}
	if field, ok := in.Map()["type"]; ok && field.Type() != packer.J_NULL {
		ob := new(Reference)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Type = ob
	}
	if field, ok := in.Map()["id"]; ok && field.Type() != packer.J_NULL {
		ob := new(Reference)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Id = ob
	}
	if field, ok := in.Map()["tags-new"]; ok && field.Type() != packer.J_NULL {
		ob := *new(Tags)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.TagsNew = ob
	}
	if field, ok := in.Map()["tags"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		for _, field := range field.Array() {
			ob, err := UnpackString(ctx, field)
			if err != nil {
				return err
			}
			v.Tags = append(v.Tags, ob)
		}
	}
	return nil
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
func UnpackPackageInterface(ctx context.Context, in packer.Packed) (PackageInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
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
func (v *Package) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["aliases"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		for key, field := range field.Map() {
			ob, err := UnpackString(ctx, field)
			if err != nil {
				return err
			}
			if v.Aliases == nil {
				v.Aliases = make(map[string]string)
			}
			v.Aliases[key] = ob
		}
	}
	if field, ok := in.Map()["recursive"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Recursive = ob
	}
	return nil
}

type ReferenceInterface interface {
	GetReference(ctx context.Context) *Reference
}

func (o *Reference) GetReference(ctx context.Context) *Reference {
	return o
}
func UnpackReferenceInterface(ctx context.Context, in packer.Packed) (ReferenceInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(ReferenceInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement ReferenceInterface", i)
		}
		return ob, nil
	case packer.J_STRING:
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
func UnpackRuleInterface(ctx context.Context, in packer.Packed) (RuleInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
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
func (v *Rule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if field, ok := in.Map()["interface"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Interface = ob
	}
	if field, ok := in.Map()["selector"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Selector = ob
	}
	if field, ok := in.Map()["optional"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Optional = ob
	}
	return nil
}

// This is the native json string data type
type String string
type StringInterface interface {
	GetString(ctx context.Context) *String
}

func (o *String) GetString(ctx context.Context) *String {
	return o
}
func UnpackStringInterface(ctx context.Context, in packer.Packed) (StringInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(StringInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement StringInterface", i)
		}
		return ob, nil
	case packer.J_STRING:
		ob := new(String)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into StringInterface.", in.Type())
	}
}
func (v *String) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_STRING {
		return fmt.Errorf("Invalid type %s while unpacking a string.", in.Type())
	}
	*v = String(in.String())
	return nil
}

// Tag cloud.
type Tags []string
type TagsInterface interface {
	GetTags(ctx context.Context) Tags
}

func (o Tags) GetTags(ctx context.Context) Tags {
	return o
}
func UnpackTagsInterface(ctx context.Context, in packer.Packed) (TagsInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, err
		}
		ob, ok := i.(TagsInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement TagsInterface", i)
		}
		return ob, nil
	case packer.J_ARRAY:
		ob := new(Tags)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, err
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into TagsInterface.", in.Type())
	}
}
func (v *Tags) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_ARRAY {
		return fmt.Errorf("Invalid type %s while unpacking an array.", in.Type())
	}
	for _, field := range in.Array() {
		ob, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		*v = append(*v, ob)
	}
	return nil
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
func UnpackTypeInterface(ctx context.Context, in packer.Packed) (TypeInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
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
func (v *Type) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["rule"]; ok && field.Type() != packer.J_NULL {
		ob := new(Type)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Rule = ob
	}
	if field, ok := in.Map()["embed"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return fmt.Errorf("Unsupported json type %s found while unpacking into an array.", field.Type())
		}
		for _, field := range field.Array() {
			ob := new(Reference)
			if err := ob.Unpack(ctx, field, false); err != nil {
				return err
			}
			v.Embed = append(v.Embed, ob)
		}
	}
	if field, ok := in.Map()["custom-kind"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.CustomKind = ob
	}
	if field, ok := in.Map()["interface"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Interface = ob
	}
	if field, ok := in.Map()["fields"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_MAP {
			return fmt.Errorf("Unsupported json type %s found while unpacking into a map.", field.Type())
		}
		for key, field := range field.Map() {
			ob, err := UnpackRuleInterface(ctx, field)
			if err != nil {
				return err
			}
			if v.Fields == nil {
				v.Fields = make(map[string]RuleInterface)
			}
			v.Fields[key] = ob
		}
	}
	if field, ok := in.Map()["native"]; ok && field.Type() != packer.J_NULL {
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Native = ob
	} else {
		field = packer.Pack("object")
		ob := new(String)
		if err := ob.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Native = ob
	}
	if field, ok := in.Map()["alias"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Alias = ob
	}
	if field, ok := in.Map()["basic"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Basic = ob
	}
	if field, ok := in.Map()["custom"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Custom = ob
	}
	return nil
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/system", 5319817629068713650)
	pkg.InitNew("array", nil, func() interface{} { return new(ArrayRule) })
	pkg.InitNew("bool", func() interface{} { return new(Bool) }, func() interface{} { return new(BoolRule) })
	pkg.InitNew("int", func() interface{} { return new(Int) }, func() interface{} { return new(IntRule) })
	pkg.InitNew("map", nil, func() interface{} { return new(MapRule) })
	pkg.InitNew("number", func() interface{} { return new(Number) }, func() interface{} { return new(NumberRule) })
	pkg.InitNew("object", func() interface{} { return new(Object) }, func() interface{} { return new(ObjectRule) })
	pkg.InitNew("package", func() interface{} { return new(Package) }, func() interface{} { return new(PackageRule) })
	pkg.InitNew("reference", func() interface{} { return new(Reference) }, func() interface{} { return new(ReferenceRule) })
	pkg.InitNew("rule", func() interface{} { return new(Rule) }, func() interface{} { return new(RuleRule) })
	pkg.InitNew("string", func() interface{} { return new(String) }, func() interface{} { return new(StringRule) })
	pkg.InitNew("tags", func() interface{} { return new(Tags) }, func() interface{} { return new(TagsRule) })
	pkg.InitNew("type", func() interface{} { return new(Type) }, func() interface{} { return new(TypeRule) })
}
