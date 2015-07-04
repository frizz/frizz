package system

import (
	reflect "reflect"

	json "kego.io/json"
)

// Automatically created basic rule for imports
type Imports_rule struct {
	*Base
	*RuleBase
}

// This is the most basic type.
type Type struct {
	*Base
	// Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.
	Rule *Type
	// Basic types don't have system:object added by default to the embedded types.
	Basic bool
	// Types which this should embed - system:object is always added unless base = true.
	Embed []Reference
	// Array of interface types that this type should support
	Is []Reference
	// This is the native json type that represents this type. If omitted, default is object.
	Native String `kego:"{\"default\":{\"value\":\"object\"}}"`
	// Is this type an interface?
	Interface bool
	// Exclude this type from the generated json?
	Exclude bool
	// Each field is listed with it's type
	Fields map[string]Rule
}

// This is the most basic type.
type Base struct {
	// Type of the object.
	Type Reference
	// All global objects should have an id.
	Id Reference
	// Description for the developer
	Description string
	// Extra validation rules for this object or descendants
	Rules []Rule
}

// Automatically created basic rule for ruleBase
type RuleBase_rule struct {
	*Base
	*RuleBase
}

// This is the native json array data type
type Array struct {
	*Base
}

// Restriction rules for maps
type Map_rule struct {
	*Base
	*RuleBase
	// This is a rule object, defining the type and restrictions on the value of the items.
	Items Rule
	// This is the minimum number of items alowed in the array
	MinItems Int
	// This is the maximum number of items alowed in the array
	MaxItems Int
}

// Restriction rules for references
type Reference_rule struct {
	*Base
	*RuleBase
	// Default value of this is missing or null
	Default Reference
}

// Restriction rules for integers
type Int_rule struct {
	*Base
	*RuleBase
	// Default value if this property is omitted
	Default Int
	// This restricts the number to be a multiple of the given number
	MultipleOf Int
	// This provides a lower bound for the restriction
	Minimum Int
	// This provides an upper bound for the restriction
	Maximum Int
}

// Restriction rules for arrays
type Array_rule struct {
	*Base
	*RuleBase
	// This is the minimum number of items allowed in the array
	MinItems Int
	// This is the maximum number of items allowed in the array
	MaxItems Int
	// If this is true, each item must be unique
	UniqueItems bool
	// This is a rule object, defining the type and restrictions on the value of the items
	Items Rule
}

// This is the native json object data type.
type Map struct {
	*Base
}

// Restriction rules for strings
type String_rule struct {
	*Base
	*RuleBase
	// The value of this string is restricted to one of the provided values
	Enum []string
	// The value must be longer or equal to the provided minimum length
	MinLength Int
	// The value must be shorter or equal to the provided maximum length
	MaxLength Int
	// This is a string that the value must match
	Equal String
	// This is a regex to match the value to
	Pattern String
	// This restricts the value to one of several built-in formats
	Format String
	// Default value of this is missing or null
	Default String
}

// Restriction rules for numbers
type Number_rule struct {
	*Base
	*RuleBase
	// This restricts the number to be a multiple of the given number
	MultipleOf Number
	// This provides a lower bound for the restriction
	Minimum Number
	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum bool
	// This provides an upper bound for the restriction
	Maximum Number
	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum bool
	// Default value if this property is omitted
	Default Number
}

// Automatically created basic rule for type
type Type_rule struct {
	*Base
	*RuleBase
}

// Automatically created basic rule for base
type Base_rule struct {
	*Base
	*RuleBase
}

// Restriction rules for bools
type Bool_rule struct {
	*Base
	*RuleBase
	// Default value if this is missing or null
	Default Bool
}

// Automatically created basic rule for rule
type Rule_rule struct {
	*Base
	*RuleBase
}

// All rules should embed this type.
type RuleBase struct {
	// If this rule is a field, this specifies that the field is optional
	Optional bool
	// Json selector defining what nodes this rule should be applied to.
	Selector string
}

// Lists imports used in this package.
type Imports struct {
	*Base
	// Map of import name to path.
	Imports map[string]String
}

func init() {
	json.RegisterType("kego.io/system", "@imports", reflect.TypeOf(&Imports_rule{}))
	json.RegisterType("kego.io/system", "type", reflect.TypeOf(&Type{}))
	json.RegisterType("kego.io/system", "base", reflect.TypeOf(&Base{}))
	json.RegisterType("kego.io/system", "@ruleBase", reflect.TypeOf(&RuleBase_rule{}))
	json.RegisterType("kego.io/system", "string", reflect.TypeOf(&String{}))
	json.RegisterType("kego.io/system", "array", reflect.TypeOf(&Array{}))
	json.RegisterType("kego.io/system", "int", reflect.TypeOf(&Int{}))
	json.RegisterType("kego.io/system", "@map", reflect.TypeOf(&Map_rule{}))
	json.RegisterType("kego.io/system", "number", reflect.TypeOf(&Number{}))
	json.RegisterType("kego.io/system", "@reference", reflect.TypeOf(&Reference_rule{}))
	json.RegisterType("kego.io/system", "@int", reflect.TypeOf(&Int_rule{}))
	json.RegisterType("kego.io/system", "reference", reflect.TypeOf(&Reference{}))
	json.RegisterType("kego.io/system", "@array", reflect.TypeOf(&Array_rule{}))
	json.RegisterType("kego.io/system", "map", reflect.TypeOf(&Map{}))
	json.RegisterType("kego.io/system", "@string", reflect.TypeOf(&String_rule{}))
	json.RegisterType("kego.io/system", "@number", reflect.TypeOf(&Number_rule{}))
	json.RegisterType("kego.io/system", "@type", reflect.TypeOf(&Type_rule{}))
	json.RegisterType("kego.io/system", "@base", reflect.TypeOf(&Base_rule{}))
	json.RegisterType("kego.io/system", "@bool", reflect.TypeOf(&Bool_rule{}))
	json.RegisterType("kego.io/system", "@rule", reflect.TypeOf(&Rule_rule{}))
	json.RegisterType("kego.io/system", "ruleBase", reflect.TypeOf(&RuleBase{}))
	json.RegisterType("kego.io/system", "imports", reflect.TypeOf(&Imports{}))
	json.RegisterType("kego.io/system", "bool", reflect.TypeOf(&Bool{}))
}
