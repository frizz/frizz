package system

import (
	"reflect"

	"kego.io/json"
)

// This is the base type for the object interface. All ke objects have this type embedded.
type Object_base struct {
	// Description for the developer
	Description string `json:"description"`
	// All global objects should have an id.
	Id Reference `json:"id"`
	// Extra validation rules for this object or descendants
	Rules []Rule `json:"rules"`
	// Type of the object.
	Type Reference `json:"type"`
}

// All rules will have this embedded in them.
type Rule_base struct {
	// If this rule is a field, this specifies that the field is optional
	Optional bool `json:"optional"`
	// Json selector defining what nodes this rule should be applied to.
	Selector string `json:"selector"`
}

// Restriction rules for arrays
type Array_rule struct {
	*Object_base
	*Rule_base
	// This is a rule object, defining the type and restrictions on the value of the items
	Items Rule `json:"items"`
	// This is the maximum number of items allowed in the array
	MaxItems Int `json:"maxItems"`
	// This is the minimum number of items allowed in the array
	MinItems Int `json:"minItems"`
	// If this is true, each item must be unique
	UniqueItems bool `json:"uniqueItems"`
}

// Restriction rules for bools
type Bool_rule struct {
	*Object_base
	*Rule_base
	// Default value if this is missing or null
	Default Bool `json:"default"`
}

// Restriction rules for integers
type Int_rule struct {
	*Object_base
	*Rule_base
	// Default value if this property is omitted
	Default Int `json:"default"`
	// This provides an upper bound for the restriction
	Maximum Int `json:"maximum"`
	// This provides a lower bound for the restriction
	Minimum Int `json:"minimum"`
	// This restricts the number to be a multiple of the given number
	MultipleOf Int `json:"multipleOf"`
}

// Restriction rules for maps
type Map_rule struct {
	*Object_base
	*Rule_base
	// This is a rule object, defining the type and restrictions on the value of the items.
	Items Rule `json:"items"`
	// This is the maximum number of items alowed in the array
	MaxItems Int `json:"maxItems"`
	// This is the minimum number of items alowed in the array
	MinItems Int `json:"minItems"`
}

// Restriction rules for numbers
type Number_rule struct {
	*Object_base
	*Rule_base
	// Default value if this property is omitted
	Default Number `json:"default"`
	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum bool `json:"exclusiveMaximum"`
	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum bool `json:"exclusiveMinimum"`
	// This provides an upper bound for the restriction
	Maximum Number `json:"maximum"`
	// This provides a lower bound for the restriction
	Minimum Number `json:"minimum"`
	// This restricts the number to be a multiple of the given number
	MultipleOf Number `json:"multipleOf"`
}

// Automatically created basic rule for object
type Object_rule struct {
	*Object_base
	*Rule_base
}

// Automatically created basic rule for package
type Package_rule struct {
	*Object_base
	*Rule_base
}

// Restriction rules for references
type Reference_rule struct {
	*Object_base
	*Rule_base
	// Default value of this is missing or null
	Default Reference `json:"default"`
}

// Automatically created basic rule for rule
type Rule_rule struct {
	*Object_base
	*Rule_base
}

// Restriction rules for strings
type String_rule struct {
	*Object_base
	*Rule_base
	// Default value of this is missing or null
	Default String `json:"default"`
	// The value of this string is restricted to one of the provided values
	Enum []string `json:"enum"`
	// This is a string that the value must match
	Equal String `json:"equal"`
	// This restricts the value to one of several built-in formats
	Format String `json:"format"`
	// The value must be shorter or equal to the provided maximum length
	MaxLength Int `json:"maxLength"`
	// The value must be longer or equal to the provided minimum length
	MinLength Int `json:"minLength"`
	// This is a regex to match the value to
	Pattern String `json:"pattern"`
}

// Automatically created basic rule for type
type Type_rule struct {
	*Object_base
	*Rule_base
}

// This is the native json array data type
type Array struct {
	*Object_base
}

// This is the native json object data type.
type Map struct {
	*Object_base
}

// Package info - forms the root node of the package
type Package struct {
	*Object_base
	// The global map is populated with all the global objects in the package
	Global map[string]Object `json:"global"`
	// Map of import aliases used in this package: key = package path, value = alias.
	Import map[string]string `json:"import"`
}

// This is the most basic type.
type Type struct {
	*Object_base
	// Optionally this is a type which is embedded in each object that implements this interface.
	Base *Type `json:"base"`
	// Basic types don't have system:object added by default to the embedded types.
	Basic bool `json:"basic"`
	// Types which this should embed - system:object is always added unless basic = true.
	Embed []Reference `json:"embed"`
	// Exclude this type from the generated json?
	Exclude bool `json:"exclude"`
	// Each field is listed with it's type
	Fields map[string]Rule `json:"fields"`
	// Is this type an interface?
	Interface bool `json:"interface"`
	// Array of interface types that this type should support
	Is []Reference `json:"is"`
	// This is the native json type that represents this type. If omitted, default is object.
	Native String `kego:"{\"default\":{\"value\":\"object\"}}" json:"native"`
	// Type that defines restriction rules for this type.
	Rule *Type `json:"rule"`
}

func init() {
	json.Register("kego.io/system", "$object", reflect.TypeOf(&Object_base{}), 0xa80f42e03150e43e)
	json.Register("kego.io/system", "$rule", reflect.TypeOf(&Rule_base{}), 0x61f37939f1737cbf)
	json.Register("kego.io/system", "@array", reflect.TypeOf(&Array_rule{}), 0xf6f09a20ac87e96f)
	json.Register("kego.io/system", "@bool", reflect.TypeOf(&Bool_rule{}), 0x849d95e096ea903a)
	json.Register("kego.io/system", "@int", reflect.TypeOf(&Int_rule{}), 0x9f6cffbf2042e2aa)
	json.Register("kego.io/system", "@map", reflect.TypeOf(&Map_rule{}), 0xb95cdb828f1494cc)
	json.Register("kego.io/system", "@number", reflect.TypeOf(&Number_rule{}), 0x14f986941edf71e9)
	json.Register("kego.io/system", "@object", reflect.TypeOf(&Object_rule{}), 0xa80f42e03150e43e)
	json.Register("kego.io/system", "@package", reflect.TypeOf(&Package_rule{}), 0x45625b24c593f855)
	json.Register("kego.io/system", "@reference", reflect.TypeOf(&Reference_rule{}), 0x67e9d97dde75d10f)
	json.Register("kego.io/system", "@rule", reflect.TypeOf(&Rule_rule{}), 0x61f37939f1737cbf)
	json.Register("kego.io/system", "@string", reflect.TypeOf(&String_rule{}), 0xe1e0d90cd0a489ca)
	json.Register("kego.io/system", "@type", reflect.TypeOf(&Type_rule{}), 0xe6e176a06eb36092)
	json.Register("kego.io/system", "array", reflect.TypeOf(Array{}), 0xf6f09a20ac87e96f)
	json.Register("kego.io/system", "bool", reflect.TypeOf(Bool{}), 0x849d95e096ea903a)
	json.Register("kego.io/system", "int", reflect.TypeOf(Int{}), 0x9f6cffbf2042e2aa)
	json.Register("kego.io/system", "map", reflect.TypeOf(Map{}), 0xb95cdb828f1494cc)
	json.Register("kego.io/system", "number", reflect.TypeOf(Number{}), 0x14f986941edf71e9)
	json.Register("kego.io/system", "object", reflect.TypeOf((*Object)(nil)).Elem(), 0xa80f42e03150e43e)
	json.Register("kego.io/system", "package", reflect.TypeOf(&Package{}), 0x45625b24c593f855)
	json.Register("kego.io/system", "reference", reflect.TypeOf(Reference{}), 0x67e9d97dde75d10f)
	json.Register("kego.io/system", "rule", reflect.TypeOf((*Rule)(nil)).Elem(), 0x61f37939f1737cbf)
	json.Register("kego.io/system", "string", reflect.TypeOf(String{}), 0xe1e0d90cd0a489ca)
	json.Register("kego.io/system", "type", reflect.TypeOf(&Type{}), 0xe6e176a06eb36092)
}
