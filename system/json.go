package system

import (
	"fmt"

	"context"

	"frizz.io/context/jsonctx"
	"frizz.io/context/sysctx"
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

func RegisterJsonTypes(ctx context.Context) {

	scache := sysctx.FromContext(ctx)
	pcache := scache.Set("frizz.io/json")

	tr := NewReference("frizz.io/system", "type")

	makeRule := func(name string) *Type {
		return &Type{
			Object: &Object{
				Id:   NewReference("frizz.io/json", fmt.Sprint(jsonctx.RULE_PREFIX, name)),
				Type: tr},
			Interface: false,
			Embed:     []*Reference{NewReference("frizz.io/system", "rule")},
			Native:    NewString("object"),
			Fields:    map[string]RuleInterface{},
			Rule:      (*Type)(nil)}
	}

	makeType := func(name string) *Type {
		return &Type{
			Object: &Object{
				Id:   NewReference("frizz.io/json", name),
				Type: tr},
			Interface: false,
			Native:    NewString(name),
			Fields:    map[string]RuleInterface{},
			Rule:      makeRule(name)}
	}

	pcache.Types.Set("string", "", makeType("string"))
	pcache.Types.Set("number", "", makeType("number"))
	pcache.Types.Set("bool", "", makeType("bool"))
	pcache.Types.Set("@string", "", makeRule("string"))
	pcache.Types.Set("@number", "", makeRule("number"))
	pcache.Types.Set("@bool", "", makeRule("bool"))

}
