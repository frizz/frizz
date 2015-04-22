package system

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kego/json"
	"github.com/stretchr/testify/assert"
)

var defaultJsonContext = &json.Context{
	PackageName: "system",
	PackagePath: "github.com/kego/system",
	Imports:     map[string]string{},
}

func TestDecodeSimple(t *testing.T) {

	test := func(data string) {

		type Foo struct {
			Foo string
			Bar float64
			Baz bool
		}

		json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Foo, "a")
		assert.Equal(t, f.Bar, 2.0)
		assert.Equal(t, f.Baz, true)

		// Clean up for the tests - don't normally need to unregister types
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

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Foo, strExpected)
		assert.Equal(t, f.Bar, numExpected)
		assert.Equal(t, f.Baz, boolExpected)

		// Clean up for the tests - don't normally need to unregister types
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

	data := `{
		"type": "foo",
		"stringMap": {"a": "aa", "b": "bb", "c": "cc"},
		"numberMap": {"d": 1, "e": 1.5, "f": 2},
		"boolMap": {"g": true, "h": false, "i": true},
		"stringArray": ["a", "b", "c"],
		"numberArray": [1, 1.5, 2],
		"boolArray": [true, false, true]
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
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

	// Clean up for the tests - don't normally need to unregister types
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

		var i interface{}
		err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, f.Embed.String, "a")
		assert.Equal(t, f.Embed.Number, 2.0)
		assert.Equal(t, f.Embed.Bool, true)

		// Clean up for the tests - don't normally need to unregister types
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

	data := `{
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
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:bar", reflect.TypeOf(&Bar{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.MapEmbed["a"].String, "a")
	assert.Equal(t, f.MapEmbed["b"].String, "b")
	assert.Equal(t, f.ArrayEmbed[0].String, "c")
	assert.Equal(t, f.ArrayEmbed[1].String, "d")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:bar")

}

func TestDecodeComposition(t *testing.T) {

	type Base struct {
		BaseString string
	}

	type Foo struct {
		Base
		FooString string
	}

	data := `{
		"type": "foo",
		"baseString": "a",
		"fooString": "b"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:base", reflect.TypeOf(&Base{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.BaseString, "a")
	assert.Equal(t, f.FooString, "b")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:base")

}

type Photo struct {
	Id string
}
type Diagram struct {
	Key string
}

func (p *Photo) Url() string {
	return fmt.Sprintf("http://www.photos.com/%s.jpg", p.Id)
}
func (d *Diagram) Url() string {
	return fmt.Sprintf("http://www.diagrams.com/%s.jpg", d.Key)
}

func TestInterface(t *testing.T) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		Img Image
	}

	data := `{
		"type": "foo",
		"img": {
			"type": "photo",
			"id": "a"
		}
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:photo", reflect.TypeOf(&Photo{}))
	json.RegisterType("github.com/kego/system:diagram", reflect.TypeOf(&Diagram{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.Img.Url(), "http://www.photos.com/a.jpg")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:photo")
	json.UnregisterType("github.com/kego/system:diagram")

}

func TestNilInterface(t *testing.T) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		Iface interface{}
	}

	data := `{
		"type": "foo",
		"iface": {
			"type": "photo",
			"id": "a"
		}
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:photo", reflect.TypeOf(&Photo{}))
	json.RegisterType("github.com/kego/system:diagram", reflect.TypeOf(&Diagram{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	p, ok := f.Iface.(*Photo)
	assert.True(t, ok, "Type %T not correct", f.Iface)
	assert.Equal(t, p.Url(), "http://www.photos.com/a.jpg")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:photo")
	json.UnregisterType("github.com/kego/system:diagram")

}

func TestInterfaceCollections(t *testing.T) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		ImageArray []Image
		ImageMap   map[string]Image
	}

	data := `{
		"type": "foo",
		"imageArray": [
			{
				"type": "photo",
				"id": "a"
			},
			{
				"type": "diagram",
				"key": "b"
			}
		],
		"imageMap": {
			"c": {
				"type": "photo",
				"id": "c"
			},
			"d": {
				"type": "diagram",
				"key": "d"
			}
		}
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:photo", reflect.TypeOf(&Photo{}))
	json.RegisterType("github.com/kego/system:diagram", reflect.TypeOf(&Diagram{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.ImageArray[0].Url(), "http://www.photos.com/a.jpg")
	assert.Equal(t, f.ImageArray[1].Url(), "http://www.diagrams.com/b.jpg")
	assert.Equal(t, f.ImageMap["c"].Url(), "http://www.photos.com/c.jpg")
	assert.Equal(t, f.ImageMap["d"].Url(), "http://www.diagrams.com/d.jpg")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:photo")
	json.UnregisterType("github.com/kego/system:diagram")

}

func TestInterfaceCollectionsComplex(t *testing.T) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		ImageMap map[string][]Image
	}

	data := `{
		"type": "foo",
		"imageMap": {
			"a": [
				{
					"type": "photo",
					"id": "c"
				},
				{
					"type": "diagram",
					"key": "d"
				}
			],
			"b": [
				{
					"type": "photo",
					"id": "e"
				},
				{
					"type": "diagram",
					"key": "f"
				}
			]
		}
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))
	json.RegisterType("github.com/kego/system:photo", reflect.TypeOf(&Photo{}))
	json.RegisterType("github.com/kego/system:diagram", reflect.TypeOf(&Diagram{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultJsonContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.ImageMap["a"][0].Url(), "http://www.photos.com/c.jpg")
	assert.Equal(t, f.ImageMap["a"][1].Url(), "http://www.diagrams.com/d.jpg")
	assert.Equal(t, f.ImageMap["b"][0].Url(), "http://www.photos.com/e.jpg")
	assert.Equal(t, f.ImageMap["b"][1].Url(), "http://www.diagrams.com/f.jpg")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")
	json.UnregisterType("github.com/kego/system:photo")
	json.UnregisterType("github.com/kego/system:diagram")

}
