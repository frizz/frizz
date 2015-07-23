package system

import (
	"reflect"

	"kego.io/json"
)

// Automatically created basic rule for aliases
type Aliases_rule struct {
	*Base
	*RuleBase
}

// Restriction rules for arrays
type Array_rule struct {
	*Base
	*RuleBase
	// This is a rule object, defining the type and restrictions on the value of the items
	Items Rule
	// This is the maximum number of items allowed in the array
	MaxItems Int
	// This is the minimum number of items allowed in the array
	MinItems Int
	// If this is true, each item must be unique
	UniqueItems bool
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

// Automatically created basic rule for imports
type Imports_rule struct {
	*Base
	*RuleBase
}

// Restriction rules for integers
type Int_rule struct {
	*Base
	*RuleBase
	// Default value if this property is omitted
	Default Int
	// This provides an upper bound for the restriction
	Maximum Int
	// This provides a lower bound for the restriction
	Minimum Int
	// This restricts the number to be a multiple of the given number
	MultipleOf Int
}

// Restriction rules for maps
type Map_rule struct {
	*Base
	*RuleBase
	// This is a rule object, defining the type and restrictions on the value of the items.
	Items Rule
	// This is the maximum number of items alowed in the array
	MaxItems Int
	// This is the minimum number of items alowed in the array
	MinItems Int
}

// Restriction rules for numbers
type Number_rule struct {
	*Base
	*RuleBase
	// Default value if this property is omitted
	Default Number
	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum bool
	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum bool
	// This provides an upper bound for the restriction
	Maximum Number
	// This provides a lower bound for the restriction
	Minimum Number
	// This restricts the number to be a multiple of the given number
	MultipleOf Number
}

// Restriction rules for references
type Reference_rule struct {
	*Base
	*RuleBase
	// Default value of this is missing or null
	Default Reference
}

// Automatically created basic rule for rule
type Rule_rule struct {
	*Base
	*RuleBase
}

// Automatically created basic rule for ruleBase
type RuleBase_rule struct {
	*Base
	*RuleBase
}

// Restriction rules for strings
type String_rule struct {
	*Base
	*RuleBase
	// Default value of this is missing or null
	Default String
	// The value of this string is restricted to one of the provided values
	Enum []string
	// This is a string that the value must match
	Equal String
	// This restricts the value to one of several built-in formats
	Format String
	// The value must be shorter or equal to the provided maximum length
	MaxLength Int
	// The value must be longer or equal to the provided minimum length
	MinLength Int
	// This is a regex to match the value to
	Pattern String
}

// Automatically created basic rule for type
type Type_rule struct {
	*Base
	*RuleBase
}

// Lists aliases used in this package.
type Aliases struct {
	*Base
	// Map of path to alias.
	Aliases map[string]string
}

// This is the native json array data type
type Array struct {
	*Base
}

// This is the most basic type.
type Base struct {
	// Description for the developer
	Description string
	// All global objects should have an id.
	Id Reference
	// Extra validation rules for this object or descendants
	Rules []Rule
	// Type of the object.
	Type Reference
}

// Lists imports used in this package.
type Imports struct {
	*Base
	// Map of path to alias.
	Imports map[string]string
}

// This is the native json object data type.
type Map struct {
	*Base
}

// All rules should embed this type.
type RuleBase struct {
	// If this rule is a field, this specifies that the field is optional
	Optional bool
	// Json selector defining what nodes this rule should be applied to.
	Selector string
}

// This is the most basic type.
type Type struct {
	*Base
	// Basic types don't have system:object added by default to the embedded types.
	Basic bool
	// Types which this should embed - system:object is always added unless base = true.
	Embed []Reference
	// Exclude this type from the generated json?
	Exclude bool
	// Each field is listed with it's type
	Fields map[string]Rule
	// Is this type an interface?
	Interface bool
	// Array of interface types that this type should support
	Is []Reference
	// This is the native json type that represents this type. If omitted, default is object.
	Native String `kego:"{\"default\":{\"value\":\"object\"}}"`
	// Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.
	Rule *Type
}

func init() {
	json.RegisterType("kego.io/system", "@aliases", reflect.TypeOf(&Aliases_rule{}), 0xa201259ad19e56d5)
	json.RegisterType("kego.io/system", "@array", reflect.TypeOf(&Array_rule{}), 0x9396ab3adc9b26c2)
	json.RegisterType("kego.io/system", "@base", reflect.TypeOf(&Base_rule{}), 0x55165bf6126129f)
	json.RegisterType("kego.io/system", "@bool", reflect.TypeOf(&Bool_rule{}), 0x18e72064215d8f71)
	json.RegisterType("kego.io/system", "@imports", reflect.TypeOf(&Imports_rule{}), 0x9bb5dc657899f549)
	json.RegisterType("kego.io/system", "@int", reflect.TypeOf(&Int_rule{}), 0xe7b98d2f66bb5b1)
	json.RegisterType("kego.io/system", "@map", reflect.TypeOf(&Map_rule{}), 0x68d11e9fab3ed4bb)
	json.RegisterType("kego.io/system", "@number", reflect.TypeOf(&Number_rule{}), 0x8c2b0d23320b8aeb)
	json.RegisterType("kego.io/system", "@reference", reflect.TypeOf(&Reference_rule{}), 0x2a5ddf0f63b58876)
	json.RegisterType("kego.io/system", "@rule", reflect.TypeOf(&Rule_rule{}), 0x50ad825a1446eccd)
	json.RegisterType("kego.io/system", "@ruleBase", reflect.TypeOf(&RuleBase_rule{}), 0xab213a8c8218179a)
	json.RegisterType("kego.io/system", "@string", reflect.TypeOf(&String_rule{}), 0xae4613a6af2bb7d6)
	json.RegisterType("kego.io/system", "@type", reflect.TypeOf(&Type_rule{}), 0xc1bba2c7da9518b7)
	json.RegisterType("kego.io/system", "aliases", reflect.TypeOf(&Aliases{}), 0xa201259ad19e56d5)
	json.RegisterType("kego.io/system", "array", reflect.TypeOf(&Array{}), 0x9396ab3adc9b26c2)
	json.RegisterType("kego.io/system", "base", reflect.TypeOf(&Base{}), 0x55165bf6126129f)
	json.RegisterType("kego.io/system", "bool", reflect.TypeOf(&Bool{}), 0x18e72064215d8f71)
	json.RegisterType("kego.io/system", "imports", reflect.TypeOf(&Imports{}), 0x9bb5dc657899f549)
	json.RegisterType("kego.io/system", "int", reflect.TypeOf(&Int{}), 0xe7b98d2f66bb5b1)
	json.RegisterType("kego.io/system", "map", reflect.TypeOf(&Map{}), 0x68d11e9fab3ed4bb)
	json.RegisterType("kego.io/system", "number", reflect.TypeOf(&Number{}), 0x8c2b0d23320b8aeb)
	json.RegisterType("kego.io/system", "reference", reflect.TypeOf(&Reference{}), 0x2a5ddf0f63b58876)
	json.RegisterType("kego.io/system", "ruleBase", reflect.TypeOf(&RuleBase{}), 0xab213a8c8218179a)
	json.RegisterType("kego.io/system", "string", reflect.TypeOf(&String{}), 0xae4613a6af2bb7d6)
	json.RegisterType("kego.io/system", "type", reflect.TypeOf(&Type{}), 0xc1bba2c7da9518b7)
}