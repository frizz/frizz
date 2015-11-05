package system

import (
	"reflect"

	"fmt"

	"kego.io/json"
)

/*
	This file simulates the json package having minimal native types for number,
	string and bool types. They can't actually be in the json package because it
	would cause an import loop to be able to use the system types (object,
	property, reference etc). This is a special case to support native json types,
	and a kluge like this shouldn't be needed in normal usage.

	These types (json:string, json:bool and json:number) are not recommended for
	high-level usage, because the lack the ability to tell if the json attribute
	is null or missing. In that case they just get the golang zero value ("", 0
	or false). Additionally their rule types don't have any of the validation
	properties. Best to use the native system types as much as possible:
	system:string, system:bool and system:number.
*/

type JsonNumberRule struct {
	*Object
	*Rule
}
type JsonStringRule struct {
	*Object
	*Rule
}
type JsonBoolRule struct {
	*Object
	*Rule
}

type JsonStringInterface interface {
	String() string
}
type JsonNumberInterface interface {
	Number() float64
}
type JsonBoolInterface interface {
	Bool() bool
}

func init() {
	json.Register("kego.io/json", "number", reflect.TypeOf(float64(1.1)), reflect.TypeOf((*JsonNumberInterface)(nil)).Elem(), 0)
	json.Register("kego.io/json", "string", reflect.TypeOf(string("")), reflect.TypeOf((*JsonStringInterface)(nil)).Elem(), 0)
	json.Register("kego.io/json", "bool", reflect.TypeOf(bool(true)), reflect.TypeOf((*JsonBoolInterface)(nil)).Elem(), 0)
	json.Register("kego.io/json", "@number", reflect.TypeOf(&JsonNumberRule{}), nil, 0)
	json.Register("kego.io/json", "@string", reflect.TypeOf(&JsonStringRule{}), nil, 0)
	json.Register("kego.io/json", "@bool", reflect.TypeOf(&JsonBoolRule{}), nil, 0)

	tr := NewReference("kego.io/system", "type")

	makeRule := func(name string) *Type {
		return &Type{
			Object: &Object{
				Id:   NewReference("kego.io/json", fmt.Sprint(RULE_PREFIX, name)),
				Type: tr},
			Interface: false,
			Embed:     []*Reference{NewReference("kego.io/system", "rule")},
			Native:    NewString("object"),
			Fields:    map[string]RuleInterface{},
			Rule:      (*Type)(nil)}
	}

	makeType := func(name string) *Type {
		return &Type{
			Object: &Object{
				Id:   NewReference("kego.io/json", name),
				Type: tr},
			Interface: false,
			Is:        []*Reference(nil),
			Native:    NewString(name),
			Fields:    map[string]RuleInterface{},
			Rule:      makeRule(name)}
	}

	Register("kego.io/json", "string", makeType("string"), 0)
	Register("kego.io/json", "number", makeType("number"), 0)
	Register("kego.io/json", "bool", makeType("bool"), 0)
	Register("kego.io/json", "@string", makeRule("string"), 0)
	Register("kego.io/json", "@number", makeRule("number"), 0)
	Register("kego.io/json", "@bool", makeRule("bool"), 0)

}
