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

type JsonNumber_rule struct {
	*Base
	*RuleBase
}
type JsonString_rule struct {
	*Base
	*RuleBase
}
type JsonBool_rule struct {
	*Base
	*RuleBase
}

func init() {
	json.RegisterType("kego.io/json", "number", reflect.TypeOf(float64(1.1)), 0)
	json.RegisterType("kego.io/json", "string", reflect.TypeOf(string("")), 0)
	json.RegisterType("kego.io/json", "bool", reflect.TypeOf(bool(true)), 0)
	json.RegisterType("kego.io/json", "@number", reflect.TypeOf(&JsonNumber_rule{}), 0)
	json.RegisterType("kego.io/json", "@string", reflect.TypeOf(&JsonString_rule{}), 0)
	json.RegisterType("kego.io/json", "@bool", reflect.TypeOf(&JsonBool_rule{}), 0)

	tr := NewReference("kego.io/system", "type")
	or := NewReference("kego.io/system", "object")

	makeRule := func(name string) *Type {
		return &Type{
			Base: &Base{
				Id:   NewReference("kego.io/json", fmt.Sprint("@", name)),
				Type: tr},
			Embed:     []Reference{or},
			Interface: false,
			Is:        []Reference{NewReference("kego.io/system", "rule")},
			Native:    NewString("object"),
			Fields:    map[string]Rule{},
			Rule:      (*Type)(nil)}
	}

	makeType := func(name string) *Type {
		return &Type{
			Base: &Base{
				Id:   NewReference("kego.io/json", name),
				Type: tr},
			Embed:     []Reference{or},
			Interface: false,
			Is:        []Reference(nil),
			Native:    NewString(name),
			Fields:    map[string]Rule(nil),
			Rule:      makeRule(name)}
	}

	RegisterType("kego.io/json", "string", makeType("string"), 0)
	RegisterType("kego.io/json", "number", makeType("number"), 0)
	RegisterType("kego.io/json", "bool", makeType("bool"), 0)
	RegisterType("kego.io/json", "@string", makeRule("string"), 0)
	RegisterType("kego.io/json", "@number", makeRule("number"), 0)
	RegisterType("kego.io/json", "@bool", makeRule("bool"), 0)

}
