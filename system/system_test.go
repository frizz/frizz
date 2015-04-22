package system

import (
	"reflect"
	"testing"

	"github.com/kego/json"
	"github.com/stretchr/testify/assert"
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

func TestDecodeSimple(t *testing.T) {

	test := func(data string) {

		type Foo struct {
			Foo string
			Bar float64
			Baz bool
		}

		json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

		context := &json.Context{
			PackageName: "system",
			PackagePath: "github.com/kego/system",
			Imports:     map[string]string{},
		}

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, context)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Foo, "a")
		assert.Equal(t, f.Bar, 2.0)
		assert.Equal(t, f.Baz, true)

		json.UnregisterType("github.com/kego/system:foo")

	}

	// Standard
	test(`{
		"type": "foo",
		"foo": "a",
		"bar": 2,
		"baz": true
	}`)

	// Type not first
	test(`{
		"foo": "a",
		"bar": 2,
		"type": "foo",
		"baz": true
	}`)

	// Extra attributes
	test(`{
		"foo": "a",
		"qux": "extra value",
		"bar": 2,
		"type": "foo",
		"baz": true
	}`)

}

func TestDecodeDefaults(t *testing.T) {

	test := func(data string, strExpected string, numExpected float64, boolExpected bool) {

		type Foo struct {
			Foo string  `kego:"{\"default\": \"a\"}"`
			Bar float64 `kego:"{\"default\": 2}"`
			Baz bool    `kego:"{\"default\": true}"`
		}

		json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

		context := &json.Context{
			PackageName: "system",
			PackagePath: "github.com/kego/system",
			Imports:     map[string]string{},
		}

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, context)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Foo, strExpected)
		assert.Equal(t, f.Bar, numExpected)
		assert.Equal(t, f.Baz, boolExpected)

		json.UnregisterType("github.com/kego/system:foo")

	}

	test(`{
		"type": "foo",
		"foo": "b",
		"bar": 3,
		"baz": false
	}`, "b", 3.0, false)

	test(`{
		"type": "foo",
		"bar": 3,
		"baz": false
	}`, "a", 3.0, false)

	test(`{
		"type": "foo",
		"baz": false
	}`, "a", 2.0, false)

	test(`{
		"type": "foo"
	}`, "a", 2.0, true)

}

func TestDecodeCollections(t *testing.T) {

	type Foo struct {
		StringMap   map[string]string
		NumberMap   map[string]float64
		BoolMap     map[string]bool
		StringArray []string
		NumberArray []float64
		BoolArray   []bool
	}

	data := []byte(`{
		"type": "foo",
		"stringMap": {"a": "aa", "b": "bb", "c": "cc"},
		"numberMap": {"d": 1, "e": 1.5, "f": 2},
		"boolMap": {"g": true, "h": false, "i": true},
		"stringArray": ["a", "b", "c"],
		"numberArray": [1, 1.5, 2],
		"boolArray": [true, false, true]
	}`)

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	context := &json.Context{
		PackageName: "system",
		PackagePath: "github.com/kego/system",
		Imports:     map[string]string{},
	}

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, context)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.StringMap["a"], "aa")
	assert.Equal(t, f.StringMap["b"], "bb")
	assert.Equal(t, f.StringMap["c"], "cc")
	assert.Equal(t, f.NumberMap["d"], 1.0)
	assert.Equal(t, f.NumberMap["e"], 1.5)
	assert.Equal(t, f.NumberMap["f"], 2.0)
	assert.Equal(t, f.BoolMap["g"], true)
	assert.Equal(t, f.BoolMap["h"], false)
	assert.Equal(t, f.BoolMap["i"], true)
	assert.Equal(t, f.StringArray, []string{"a", "b", "c"})
	assert.Equal(t, f.NumberArray, []float64{1.0, 1.5, 2.0})
	assert.Equal(t, f.BoolArray, []bool{true, false, true})

	json.UnregisterType("github.com/kego/system:foo")

}

func TestDecodeEmbed(t *testing.T) {

	test := func(data string) {

		type Bar struct {
			String string
			Number float64
			Bool   bool
		}

		type Foo struct {
			Embed Bar
		}

		json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
		json.RegisterType("github.com/kego/system:bar", reflect.TypeOf(&Bar{}))

		context := &json.Context{
			PackageName: "system",
			PackagePath: "github.com/kego/system",
			Imports:     map[string]string{},
		}

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, context)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Embed.String, "a")
		assert.Equal(t, f.Embed.Number, 2.0)
		assert.Equal(t, f.Embed.Bool, true)

		json.UnregisterType("github.com/kego/system:foo")
		json.UnregisterType("github.com/kego/system:bar")
	}

	// Standard
	test(`{
		"type": "foo",
		"embed": {
			"type": "bar",
			"string": "a",
			"number": 2,
			"bool": true
		}
	}`)

	// Type not first
	test(`{
		"embed": {
			"string": "a",
			"number": 2,
			"type": "bar",
			"bool": true
		},
		"type": "foo"
	}`)

}

func TestDecodeEmbedCollections(t *testing.T) {

	type Bar struct {
		String string
	}

	type Foo struct {
		MapEmbed   map[string]Bar
		ArrayEmbed []Bar
	}

	data := []byte(`{
		"type": "foo",
		"mapEmbed": {
			"a": {
				"type": "bar",
				"string": "a"
			},
			"b": {
				"type": "bar",
				"string": "b"
			}
		},
		"arrayEmbed": [
			{
				"type": "bar",
				"string": "c"
			},
			{
				"type": "bar",
				"string": "d"
			}
		]
	}`)

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:bar", reflect.TypeOf(&Bar{}))

	context := &json.Context{
		PackageName: "system",
		PackagePath: "github.com/kego/system",
		Imports:     map[string]string{},
	}

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, context)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.MapEmbed["a"].String, "a")
	assert.Equal(t, f.MapEmbed["b"].String, "b")
	assert.Equal(t, f.ArrayEmbed[0].String, "c")
	assert.Equal(t, f.ArrayEmbed[1].String, "d")

	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:bar")

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
