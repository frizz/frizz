package system

import (
	"reflect"

	"kego.io/json"
)

// Restriction rules for arrays
type ArrayRule struct {
	*Object
	*Rule
	// This is a rule object, defining the type and restrictions on the value of the items
	Items RuleInterface `json:"items"`
	// This is the maximum number of items allowed in the array
	MaxItems *Int `json:"maxItems"`
	// This is the minimum number of items allowed in the array
	MinItems *Int `json:"minItems"`
	// If this is true, each item must be unique
	UniqueItems bool `json:"uniqueItems"`
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
	MultipleOf *Int `json:"multipleOf"`
}

// Restriction rules for maps
type MapRule struct {
	*Object
	*Rule
	// This is a rule object, defining the type and restrictions on the value of the items.
	Items RuleInterface `json:"items"`
	// This is the maximum number of items alowed in the array
	MaxItems *Int `json:"maxItems"`
	// This is the minimum number of items alowed in the array
	MinItems *Int `json:"minItems"`
}

// Restriction rules for numbers
type NumberRule struct {
	*Object
	*Rule
	// Default value if this property is omitted
	Default *Number `json:"default"`
	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum bool `json:"exclusiveMaximum"`
	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum bool `json:"exclusiveMinimum"`
	// This provides an upper bound for the restriction
	Maximum *Number `json:"maximum"`
	// This provides a lower bound for the restriction
	Minimum *Number `json:"minimum"`
	// This restricts the number to be a multiple of the given number
	MultipleOf *Number `json:"multipleOf"`
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
	// The value must be shorter or equal to the provided maximum length
	MaxLength *Int `json:"maxLength"`
	// The value must be longer or equal to the provided minimum length
	MinLength *Int `json:"minLength"`
	// This is a regex to match the value to
	Pattern *String `json:"pattern"`
}

// Automatically created basic rule for type
type TypeRule struct {
	*Object
	*Rule
}
type BoolInterface interface {
	GetBool() *Bool
}

func (o *Bool) GetBool() *Bool {
	return o
}

type IntInterface interface {
	GetInt() *Int
}

func (o *Int) GetInt() *Int {
	return o
}

type NumberInterface interface {
	GetNumber() *Number
}

func (o *Number) GetNumber() *Number {
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
	// Type of the object.
	Type *Reference `json:"type"`
}
type ObjectInterface interface {
	GetObject() *Object
}

func (o *Object) GetObject() *Object {
	return o
}

// Package info - forms the root node of the package
type Package struct {
	*Object
	// Map of import aliases used in this package: key = package path, value = alias.
	Aliases map[string]string `json:"aliases"`
}
type PackageInterface interface {
	GetPackage() *Package
}

func (o *Package) GetPackage() *Package {
	return o
}

type ReferenceInterface interface {
	GetReference() *Reference
}

func (o *Reference) GetReference() *Reference {
	return o
}

// All rules will have this embedded in them.
type Rule struct {
	// Special field that should be excluded when marshaling - e.g. package.global
	Exclude bool `json:"exclude"`
	// Use the single method getter interface for this type
	Interface bool `json:"interface"`
	// If this rule is a field, this specifies that the field is optional
	Optional bool `json:"optional"`
	// Json selector defining what nodes this rule should be applied to.
	Selector string `json:"selector"`
}
type RuleInterface interface {
	GetRule() *Rule
}

func (o *Rule) GetRule() *Rule {
	return o
}

type StringInterface interface {
	GetString() *String
}

func (o *String) GetString() *String {
	return o
}

// This is the most basic type.
type Type struct {
	*Object
	// Basic types don't have system:object added by default to the embedded types.
	Basic bool `json:"basic"`
	// Types which this should embed - system:object is always added unless basic = true.
	Embed []*Reference `json:"embed"`
	// Each field is listed with it's type
	Fields map[string]RuleInterface `json:"fields"`
	// Is this type an interface?
	Interface bool `json:"interface"`
	// Array of interface types that this type should support
	Is []*Reference `json:"is"`
	// This is the native json type that represents this type. If omitted, default is object.
	Native *String `kego:"{\"default\":{\"value\":\"object\"}}" json:"native"`
	// Type that defines restriction rules for this type.
	Rule *Type `json:"rule"`
}
type TypeInterface interface {
	GetType() *Type
}

func (o *Type) GetType() *Type {
	return o
}
func init() {
	json.Register("kego.io/system", "@array", reflect.TypeOf((*ArrayRule)(nil)), nil, 7974503418364850785)
	json.Register("kego.io/system", "@bool", reflect.TypeOf((*BoolRule)(nil)), nil, 11092538838457238262)
	json.Register("kego.io/system", "@int", reflect.TypeOf((*IntRule)(nil)), nil, 14978258480540711733)
	json.Register("kego.io/system", "@map", reflect.TypeOf((*MapRule)(nil)), nil, 9348274211988509508)
	json.Register("kego.io/system", "@number", reflect.TypeOf((*NumberRule)(nil)), nil, 4117207956434057366)
	json.Register("kego.io/system", "@object", reflect.TypeOf((*ObjectRule)(nil)), nil, 9849996233444553044)
	json.Register("kego.io/system", "@package", reflect.TypeOf((*PackageRule)(nil)), nil, 5866188538977148683)
	json.Register("kego.io/system", "@reference", reflect.TypeOf((*ReferenceRule)(nil)), nil, 17067036908016315687)
	json.Register("kego.io/system", "@rule", reflect.TypeOf((*RuleRule)(nil)), nil, 12761314134647754082)
	json.Register("kego.io/system", "@string", reflect.TypeOf((*StringRule)(nil)), nil, 12589443105519283292)
	json.Register("kego.io/system", "@type", reflect.TypeOf((*TypeRule)(nil)), nil, 10828715750116313960)
	json.Register("kego.io/system", "bool", reflect.TypeOf((*Bool)(nil)), reflect.TypeOf((*BoolInterface)(nil)).Elem(), 11092538838457238262)
	json.Register("kego.io/system", "int", reflect.TypeOf((*Int)(nil)), reflect.TypeOf((*IntInterface)(nil)).Elem(), 14978258480540711733)
	json.Register("kego.io/system", "number", reflect.TypeOf((*Number)(nil)), reflect.TypeOf((*NumberInterface)(nil)).Elem(), 4117207956434057366)
	json.Register("kego.io/system", "object", reflect.TypeOf((*Object)(nil)), reflect.TypeOf((*ObjectInterface)(nil)).Elem(), 9849996233444553044)
	json.Register("kego.io/system", "package", reflect.TypeOf((*Package)(nil)), reflect.TypeOf((*PackageInterface)(nil)).Elem(), 5866188538977148683)
	json.Register("kego.io/system", "reference", reflect.TypeOf((*Reference)(nil)), reflect.TypeOf((*ReferenceInterface)(nil)).Elem(), 17067036908016315687)
	json.Register("kego.io/system", "rule", reflect.TypeOf((*Rule)(nil)), reflect.TypeOf((*RuleInterface)(nil)).Elem(), 12761314134647754082)
	json.Register("kego.io/system", "string", reflect.TypeOf((*String)(nil)), reflect.TypeOf((*StringInterface)(nil)).Elem(), 12589443105519283292)
	json.Register("kego.io/system", "type", reflect.TypeOf((*Type)(nil)), reflect.TypeOf((*TypeInterface)(nil)).Elem(), 10828715750116313960)
}
