package system

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kego/json"
)

// This is the most basic type.
type Test struct {
	Foo  map[string]string `kego:"{\"default\": {\"foo\": \"bar\"}}"`
	Test TestEmbedded      `kego:"{\"default\": {\"type\": \"testEmbedded\", \"foo\": \"foo2\", \"bar\": \"foo2\"}}"`
}
type TestEmbedded struct {
	Foo string
	Bar String
}

func init() {
	json.RegisterType("github.com/kego/system:test", reflect.TypeOf(&Test{}))
	json.RegisterType("github.com/kego/system:testEmbedded", reflect.TypeOf(&TestEmbedded{}))
}

func TestDecodeMap(t *testing.T) {

	data := []byte(`{
		"type": "test",
		"foo": {"foo": "qux"}
	}`)
	ctx := &json.Context{
		PackageName: "system",
		PackagePath: "github.com/kego/system",
		Imports:     map[string]string{},
	}
	var v interface{}
	err := json.UnmarshalTyped(data, &v, ctx)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("%#v\n", v)
	typ, ok := v.(*Test)
	if !ok {
		t.Errorf("Wrong type %T - expecting *system.Test", v)
		return
	}
	if typ == nil {
		t.Errorf("Typ is nil")
		return
	}
	if typ.Test.Foo != "foo2" {
		t.Errorf("t.Test.Foo has wrong value %s. Expected foo2\n", typ.Test.Foo)
		return
	}
	if typ.Foo["foo"] != "qux" {
		t.Errorf("t.Foo[foo] has wrong value %s. Expected bar\n", typ.Foo["bar"])
		return
	}

}

func TestDecode(t *testing.T) {

	data := []byte(`{
		"extends": "foo",
		"description": "This is the native json bool data type",
		"type": "type",
		"id": "bool",
		"native": "bool",
		"is": ["foo", "bar"],
		"foo": {
			"bar": "barValue",
			"quz": "quzValue",
			"type": "typeValue"
		},
		"rule": {
			"description": "Restriction rules for bools",
			"type": "type",
			"id": "@bool",
			"is": ["rule"],
			"properties": {
				"default": {
					"description": "Default value of this is missing or null",
					"type": "property",
					"optional": true,
					"item": {
						"type": "@bool"
					}
				}
			}
		}
	}`)
	ctx := &json.Context{
		PackageName: "system",
		PackagePath: "github.com/kego/system",
		Imports:     map[string]string{},
	}
	var v interface{}
	err := json.UnmarshalTyped(data, &v, ctx)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("%#v\n", v)
	typ, ok := v.(*Type)
	if !ok {
		t.Errorf("Wrong type %T - expecting *system.Type", v)
		return
	}
	if typ == nil {
		t.Errorf("Typ is nil")
		return
	}
	if typ.Foo["bar"] != "barValue" {
		t.Errorf("t.Foo[bar] has wrong value %s. Expected barValue\n", typ.Foo["bar"])
		return
	}

}
