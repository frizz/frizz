package system

import (
	"reflect"

	"kego.io/json"
)

//***********************************************************
//*** @array ***
//***********************************************************

// Restriction rules for arrays
type Array_rule struct {
	*Selector

	// This is a rule object, defining the type and restrictions on the value of the items
	Items Rule

	// This is the maximum number of items alowed in the array
	MaxItems Int

	// This is the minimum number of items alowed in the array
	MinItems Int

	// If this is true, each item must be unique
	UniqueItems Bool `kego:"{\"default\":{\"value\":false}}"`
}

//***********************************************************
//*** @bool ***
//***********************************************************

// Restriction rules for bools
type Bool_rule struct {
	*Selector

	// Default value of this is missing or null
	Default Bool
}

//***********************************************************
//*** @context ***
//***********************************************************

// Automatically created basic rule for context
type Context_rule struct {
	*Selector
}

//***********************************************************
//*** @int ***
//***********************************************************

// Restriction rules for integers
type Int_rule struct {
	*Selector

	// Default value if this property is omitted
	Default Int

	// This provides an upper bound for the restriction
	Maximum Int

	// This provides a lower bound for the restriction
	Minimum Int

	// This restricts the number to be a multiple of the given number
	MultipleOf Int
}

//***********************************************************
//*** @map ***
//***********************************************************

// Restriction rules for maps
type Map_rule struct {
	*Selector

	// This is a rule object, defining the type and restrictions on the value of the items.
	Items Rule

	// This is the maximum number of items alowed in the array
	MaxItems Int

	// This is the minimum number of items alowed in the array
	MinItems Int
}

//***********************************************************
//*** @number ***
//***********************************************************

// Restriction rules for numbers
type Number_rule struct {
	*Selector

	// Default value if this property is omitted
	Default Number

	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum Bool `kego:"{\"default\":{\"value\":false}}"`

	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum Bool `kego:"{\"default\":{\"value\":false}}"`

	// This provides an upper bound for the restriction
	Maximum Number

	// This provides a lower bound for the restriction
	Minimum Number

	// This restricts the number to be a multiple of the given number
	MultipleOf Number
}

//***********************************************************
//*** @object ***
//***********************************************************

// Automatically created basic rule for object
type Object_rule struct {
	*Selector
}

//***********************************************************
//*** @property ***
//***********************************************************

// Automatically created basic rule for property
type Property_rule struct {
	*Selector
}

//***********************************************************
//*** @reference ***
//***********************************************************

// Restriction rules for references
type Reference_rule struct {
	*Selector

	// Default value of this is missing or null
	Default Reference
}

//***********************************************************
//*** @rule ***
//***********************************************************

// Automatically created basic rule for rule
type Rule_rule struct {
	*Selector
}

//***********************************************************
//*** @selector ***
//***********************************************************

// Automatically created basic rule for selector
type Selector_rule struct {
	*Selector
}

//***********************************************************
//*** @string ***
//***********************************************************

// Restriction rules for strings
type String_rule struct {
	*Selector

	// Default value of this is missing or null
	Default String

	// The value of this string is restricted to one of the provided values
	Enum []String

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

//***********************************************************
//*** @type ***
//***********************************************************

// Automatically created basic rule for type
type Type_rule struct {
	*Selector
}

//***********************************************************
//*** array ***
//***********************************************************

// This is the native json array data type
type Array struct {
	*Object
}

//***********************************************************
//*** bool ***
//***********************************************************

//***********************************************************
//*** context ***
//***********************************************************

//***********************************************************
//*** int ***
//***********************************************************

//***********************************************************
//*** map ***
//***********************************************************

// This is the native json object data type.
type Map struct {
	*Object
}

//***********************************************************
//*** number ***
//***********************************************************

//***********************************************************
//*** object ***
//***********************************************************

// This is the most basic type.
type Object struct {

	// Unmarshaling context. This should not be in the json - it's added by the unmarshaler.
	Context *Context

	// Description for the developer
	Description string

	// All global objects should have an id.
	Id string

	// Extra validation rules for this object or descendants
	Rules []Rule

	// Type of the object.
	Type Reference
}

//***********************************************************
//*** property ***
//***********************************************************

// This is a property of a type
type Property struct {
	*Object

	// This specifies that the field is the default value for a rule
	Defaulter bool

	// This is a rule object, defining the type and restrictions on the value of the this property
	Item Rule

	// This specifies that the field is optional
	Optional bool
}

//***********************************************************
//*** reference ***
//***********************************************************

//***********************************************************
//*** rule ***
//***********************************************************

//***********************************************************
//*** selector ***
//***********************************************************

// All rules should extend this type.
type Selector struct {
	*Object

	// Json selector defining what nodes this rule should be applied to.
	Selector String `kego:"{\"default\":{\"value\":\":root\"}}"`
}

//***********************************************************
//*** string ***
//***********************************************************

//***********************************************************
//*** type ***
//***********************************************************

// This is the most basic type.
type Type struct {
	*Object

	// Exclude this type from the generated json?
	Exclude bool

	// Type which this should extend
	Extends Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"kego.io/system:object\",\"path\":\"kego.io/system\"}}"`

	// Is this type an interface?
	Interface bool

	// Array of interface types that this type should support
	Is []Reference

	// This is the native json type that represents this type. If omitted, default is object.
	Native String `kego:"{\"default\":{\"value\":\"object\"}}"`

	// Each field is listed with it's type
	Properties map[string]*Property

	// Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.
	Rule *Type
}

func init() {

	json.RegisterType("kego.io/system:@array", reflect.TypeOf(&Array_rule{}))

	json.RegisterType("kego.io/system:@bool", reflect.TypeOf(&Bool_rule{}))

	json.RegisterType("kego.io/system:@context", reflect.TypeOf(&Context_rule{}))

	json.RegisterType("kego.io/system:@int", reflect.TypeOf(&Int_rule{}))

	json.RegisterType("kego.io/system:@map", reflect.TypeOf(&Map_rule{}))

	json.RegisterType("kego.io/system:@number", reflect.TypeOf(&Number_rule{}))

	json.RegisterType("kego.io/system:@object", reflect.TypeOf(&Object_rule{}))

	json.RegisterType("kego.io/system:@property", reflect.TypeOf(&Property_rule{}))

	json.RegisterType("kego.io/system:@reference", reflect.TypeOf(&Reference_rule{}))

	json.RegisterType("kego.io/system:@rule", reflect.TypeOf(&Rule_rule{}))

	json.RegisterType("kego.io/system:@selector", reflect.TypeOf(&Selector_rule{}))

	json.RegisterType("kego.io/system:@string", reflect.TypeOf(&String_rule{}))

	json.RegisterType("kego.io/system:@type", reflect.TypeOf(&Type_rule{}))

	json.RegisterType("kego.io/system:array", reflect.TypeOf(&Array{}))

	json.RegisterType("kego.io/system:bool", reflect.TypeOf(&Bool{}))

	json.RegisterType("kego.io/system:context", reflect.TypeOf(&Context{}))

	json.RegisterType("kego.io/system:int", reflect.TypeOf(&Int{}))

	json.RegisterType("kego.io/system:map", reflect.TypeOf(&Map{}))

	json.RegisterType("kego.io/system:number", reflect.TypeOf(&Number{}))

	json.RegisterType("kego.io/system:object", reflect.TypeOf(&Object{}))

	json.RegisterType("kego.io/system:property", reflect.TypeOf(&Property{}))

	json.RegisterType("kego.io/system:reference", reflect.TypeOf(&Reference{}))

	json.RegisterType("kego.io/system:selector", reflect.TypeOf(&Selector{}))

	json.RegisterType("kego.io/system:string", reflect.TypeOf(&String{}))

	json.RegisterType("kego.io/system:type", reflect.TypeOf(&Type{}))

}
