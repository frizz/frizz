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
	*Object
}
type JsonString_rule struct {
	*Object
}
type JsonBool_rule struct {
	*Object
}

func init() {
	json.RegisterType("kego.io/json:number", reflect.TypeOf(float64(1.1)))
	json.RegisterType("kego.io/json:string", reflect.TypeOf(string("")))
	json.RegisterType("kego.io/json:bool", reflect.TypeOf(bool(true)))
	json.RegisterType("kego.io/json:@number", reflect.TypeOf(&JsonNumber_rule{}))
	json.RegisterType("kego.io/json:@string", reflect.TypeOf(&JsonString_rule{}))
	json.RegisterType("kego.io/json:@bool", reflect.TypeOf(&JsonBool_rule{}))

	c := Context{Imports: map[string]string{}, Package: "kego.io/json"}
	tr := NewReference("kego.io/system", "type")
	or := NewReference("kego.io/system", "object")

	makeRule := func(name string) *Type {
		return &Type{
			Object: &Object{
				Context: &c,
				Id:      fmt.Sprintf("@%s", name),
				Type:    tr},
			Extends:    or,
			Interface:  false,
			Is:         []Reference{NewReference("kego.io/system", "rule")},
			Native:     NewString("object"),
			Properties: map[string]*Property{},
			Rule:       (*Type)(nil)}
	}

	makeType := func(name string) *Type {
		return &Type{
			Object: &Object{
				Context: &c,
				Id:      name,
				Type:    tr},
			Extends:    or,
			Interface:  false,
			Is:         []Reference(nil),
			Native:     NewString(name),
			Properties: map[string]*Property(nil),
			Rule:       makeRule(name)}
	}

	RegisterType("kego.io/json:string", makeType("string"))
	RegisterType("kego.io/json:number", makeType("number"))
	RegisterType("kego.io/json:bool", makeType("bool"))
	RegisterType("kego.io/json:@string", makeRule("string"))
	RegisterType("kego.io/json:@number", makeRule("number"))
	RegisterType("kego.io/json:@bool", makeRule("bool"))

}
