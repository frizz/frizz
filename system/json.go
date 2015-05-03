package system

import (
	"reflect"

	"fmt"

	"kego.io/json"
)

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
	json.RegisterType("kego.io/json:number", reflect.TypeOf(1.1))
	json.RegisterType("kego.io/json:string", reflect.TypeOf(""))
	json.RegisterType("kego.io/json:bool", reflect.TypeOf(true))
	json.RegisterType("kego.io/json:@number", reflect.TypeOf(&JsonNumber_rule{}))
	json.RegisterType("kego.io/json:@string", reflect.TypeOf(&JsonString_rule{}))
	json.RegisterType("kego.io/json:@bool", reflect.TypeOf(&JsonBool_rule{}))

	c := Context{Imports: map[string]String{}, Package: String{Value: "kego.io/json", Exists: true}}
	t := Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}
	o := Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}
	rule := func(name string) *Type {
		return &Type{
			Object: &Object{
				Context: &c,
				Id:      String{Value: fmt.Sprintf("@%s", name), Exists: true},
				Type:    t},
			Extends:    o,
			Interface:  Bool{Value: false, Exists: true},
			Is:         []Reference{Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}},
			Native:     String{Value: "object", Exists: true},
			Properties: map[string]*Property{},
			Rule:       (*Type)(nil)}
	}
	minimal := func(name string) *Type {
		return &Type{
			Object: &Object{
				Context: &c,
				Id:      String{Value: name, Exists: true},
				Type:    t},
			Extends:    o,
			Interface:  Bool{Value: false, Exists: true},
			Is:         []Reference(nil),
			Native:     String{Value: "bool", Exists: true},
			Properties: map[string]*Property(nil),
			Rule:       rule(name)}
	}

	RegisterType("kego.io/json:string", minimal("string"))
	RegisterType("kego.io/json:number", minimal("number"))
	RegisterType("kego.io/json:bool", minimal("bool"))
	RegisterType("kego.io/json:@string", rule("string"))
	RegisterType("kego.io/json:@number", rule("number"))
	RegisterType("kego.io/json:@bool", rule("bool"))

}
