// info:{"Path":"kego.io/system","Hash":13795700858587685492}
package system

// ke: {"file": {"notest": true}}

import (
	"context"
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

// Restriction rules for bools
type BoolRule struct {
	*Object
	*Rule
	// Default value if this is missing or null
	Default *Bool `json:"default"`
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

// Automatically created basic rule for object
type ObjectRule struct {
	*Object
	*Rule
}

// Automatically created basic rule for package
type PackageRule struct {
	*Object
	*Rule
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

// Automatically created basic rule for rule
type RuleRule struct {
	*Object
	*Rule
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

// Automatically created basic rule for tags
type TagsRule struct {
	*Object
	*Rule
}

// Automatically created basic rule for type
type TypeRule struct {
	*Object
	*Rule
}

// This is the native json array data type
type Array struct {
	*Object
}
type ArrayInterface interface {
	GetArray(ctx context.Context) *Array
}

func (o *Array) GetArray(ctx context.Context) *Array {
	return o
}

// This is the native json bool data type
type Bool bool
type BoolInterface interface {
	GetBool(ctx context.Context) *Bool
}

func (o *Bool) GetBool(ctx context.Context) *Bool {
	return o
}

type IntInterface interface {
	GetInt(ctx context.Context) *Int
}

func (o *Int) GetInt(ctx context.Context) *Int {
	return o
}

// This is the native json object data type.
type Map struct {
	*Object
}
type MapInterface interface {
	GetMap(ctx context.Context) *Map
}

func (o *Map) GetMap(ctx context.Context) *Map {
	return o
}

// This is the native json number data type
type Number float64
type NumberInterface interface {
	GetNumber(ctx context.Context) *Number
}

func (o *Number) GetNumber(ctx context.Context) *Number {
	return o
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
	Tags *Tags `json:"tags"`
	// Type of the object.
	Type *Reference `json:"type"`
}
type ObjectInterface interface {
	GetObject(ctx context.Context) *Object
}

func (o *Object) GetObject(ctx context.Context) *Object {
	return o
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

type ReferenceInterface interface {
	GetReference(ctx context.Context) *Reference
}

func (o *Reference) GetReference(ctx context.Context) *Reference {
	return o
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

// This is the native json string data type
type String string
type StringInterface interface {
	GetString(ctx context.Context) *String
}

func (o *String) GetString(ctx context.Context) *String {
	return o
}

// Tag cloud.
type Tags []string
type TagsInterface interface {
	GetTags(ctx context.Context) *Tags
}

func (o *Tags) GetTags(ctx context.Context) *Tags {
	return o
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
	// Types which this should embed - system:object is always added unless basic = true.
	Embed []*Reference `json:"embed"`
	// Each field is listed with it's type
	Fields map[string]RuleInterface `json:"fields"`
	// Is this type an interface?
	Interface bool `json:"interface"`
	// This is the native json type that represents this type. If omitted, default is object.
	Native *String `kego:"{\"default\":{\"value\":\"object\"}}" json:"native"`
	// Type that defines restriction rules for this type.
	Rule *Type `json:"rule"`
}
type TypeInterface interface {
	GetType(ctx context.Context) *Type
}

func (o *Type) GetType(ctx context.Context) *Type {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/system", 13795700858587685492)
	pkg.InitType("array", nil, reflect.TypeOf((*ArrayRule)(nil)), nil)
	pkg.InitType("bool", reflect.TypeOf((*Bool)(nil)), reflect.TypeOf((*BoolRule)(nil)), reflect.TypeOf((*BoolInterface)(nil)).Elem())
	pkg.InitType("int", reflect.TypeOf((*Int)(nil)), reflect.TypeOf((*IntRule)(nil)), reflect.TypeOf((*IntInterface)(nil)).Elem())
	pkg.InitType("map", nil, reflect.TypeOf((*MapRule)(nil)), nil)
	pkg.InitType("number", reflect.TypeOf((*Number)(nil)), reflect.TypeOf((*NumberRule)(nil)), reflect.TypeOf((*NumberInterface)(nil)).Elem())
	pkg.InitType("object", reflect.TypeOf((*Object)(nil)), reflect.TypeOf((*ObjectRule)(nil)), reflect.TypeOf((*ObjectInterface)(nil)).Elem())
	pkg.InitType("package", reflect.TypeOf((*Package)(nil)), reflect.TypeOf((*PackageRule)(nil)), reflect.TypeOf((*PackageInterface)(nil)).Elem())
	pkg.InitType("reference", reflect.TypeOf((*Reference)(nil)), reflect.TypeOf((*ReferenceRule)(nil)), reflect.TypeOf((*ReferenceInterface)(nil)).Elem())
	pkg.InitType("rule", reflect.TypeOf((*Rule)(nil)), reflect.TypeOf((*RuleRule)(nil)), reflect.TypeOf((*RuleInterface)(nil)).Elem())
	pkg.InitType("string", reflect.TypeOf((*String)(nil)), reflect.TypeOf((*StringRule)(nil)), reflect.TypeOf((*StringInterface)(nil)).Elem())
	pkg.InitType("tags", reflect.TypeOf((*Tags)(nil)), reflect.TypeOf((*TagsRule)(nil)), reflect.TypeOf((*TagsInterface)(nil)).Elem())
	pkg.InitType("type", reflect.TypeOf((*Type)(nil)), reflect.TypeOf((*TypeRule)(nil)), reflect.TypeOf((*TypeInterface)(nil)).Elem())
}
