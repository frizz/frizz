package json

import (
	"fmt"
	"reflect"
	"testing"

	"kego.io/kerr/assert"
)

func TestDecodeSimple(t *testing.T) {

	test := func(data string) {

		type Foo struct {
			Foo string
			Bar float64
			Baz bool
		}

		RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))

		// Clean up for the tests - don't normally need to unregister types
		defer UnregisterType("kego.io/json:foo")

		var i interface{}
		unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
		assert.False(t, unknown)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, "a", f.Foo)
		assert.Equal(t, 2.0, f.Bar)
		assert.Equal(t, true, f.Baz)

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
			Foo string  `kego:"{\"default\":{\"value\":\"a\"}}"`
			Bar float64 `kego:"{\"default\":{\"value\":2}}"`
			Baz bool    `kego:"{\"default\":{\"value\":true}}"`
		}

		RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))

		// Clean up for the tests - don't normally need to unregister types
		defer UnregisterType("kego.io/json:foo")

		var i interface{}
		unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
		assert.False(t, unknown)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, strExpected, f.Foo)
		assert.Equal(t, numExpected, f.Bar)
		assert.Equal(t, boolExpected, f.Baz)

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "aa", f.StringMap["a"])
	assert.Equal(t, "bb", f.StringMap["b"])
	assert.Equal(t, "cc", f.StringMap["c"])
	assert.Equal(t, 1.0, f.NumberMap["d"])
	assert.Equal(t, 1.5, f.NumberMap["e"])
	assert.Equal(t, 2.0, f.NumberMap["f"])
	assert.Equal(t, true, f.BoolMap["g"])
	assert.Equal(t, false, f.BoolMap["h"])
	assert.Equal(t, true, f.BoolMap["i"])
	assert.Equal(t, []string{"a", "b", "c"}, f.StringArray)
	assert.Equal(t, []float64{1.0, 1.5, 2.0}, f.NumberArray)
	assert.Equal(t, []bool{true, false, true}, f.BoolArray)

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

		RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
		RegisterType("kego.io/json:bar", reflect.TypeOf(&Bar{}))

		// Clean up for the tests - don't normally need to unregister types
		defer UnregisterType("kego.io/json:foo")
		defer UnregisterType("kego.io/json:bar")

		var i interface{}
		unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
		assert.False(t, unknown)
		assert.NoError(t, err)
		f, ok := i.(*Foo)
		assert.True(t, ok, "Type %T not correct", i)
		assert.NotNil(t, f)
		assert.Equal(t, "a", f.Embed.String)
		assert.Equal(t, 2.0, f.Embed.Number)
		assert.Equal(t, true, f.Embed.Bool)

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:bar", reflect.TypeOf(&Bar{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:bar")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "a", f.MapEmbed["a"].String)
	assert.Equal(t, "b", f.MapEmbed["b"].String)
	assert.Equal(t, "c", f.ArrayEmbed[0].String)
	assert.Equal(t, "d", f.ArrayEmbed[1].String)

}

func TestDecodeComposition(t *testing.T) {

	type Base struct {
		BaseString string
	}

	type Foo struct {
		*Base
		FooString string
	}

	data := `{
		"type": "foo",
		"baseString": "a",
		"fooString": "b"
	}`

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:base", reflect.TypeOf(&Base{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:base")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "a", f.BaseString)
	assert.Equal(t, "b", f.FooString)

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:photo", reflect.TypeOf(&Photo{}))
	RegisterType("kego.io/json:diagram", reflect.TypeOf(&Diagram{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:photo")
	defer UnregisterType("kego.io/json:diagram")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "http://www.photos.com/a.jpg", f.Img.Url())

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:photo", reflect.TypeOf(&Photo{}))
	RegisterType("kego.io/json:diagram", reflect.TypeOf(&Diagram{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:photo")
	defer UnregisterType("kego.io/json:diagram")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	p, ok := f.Iface.(*Photo)
	assert.True(t, ok, "Type %T not correct", f.Iface)
	assert.Equal(t, "http://www.photos.com/a.jpg", p.Url())

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:photo", reflect.TypeOf(&Photo{}))
	RegisterType("kego.io/json:diagram", reflect.TypeOf(&Diagram{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:photo")
	defer UnregisterType("kego.io/json:diagram")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "http://www.photos.com/a.jpg", f.ImageArray[0].Url())
	assert.Equal(t, "http://www.diagrams.com/b.jpg", f.ImageArray[1].Url())
	assert.Equal(t, "http://www.photos.com/c.jpg", f.ImageMap["c"].Url())
	assert.Equal(t, "http://www.diagrams.com/d.jpg", f.ImageMap["d"].Url())

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

	RegisterType("kego.io/json:foo", reflect.TypeOf(&Foo{}))
	RegisterType("kego.io/json:photo", reflect.TypeOf(&Photo{}))
	RegisterType("kego.io/json:diagram", reflect.TypeOf(&Diagram{}))

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterType("kego.io/json:foo")
	defer UnregisterType("kego.io/json:photo")
	defer UnregisterType("kego.io/json:diagram")

	var i interface{}
	unknown, err := Unmarshal([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "http://www.photos.com/c.jpg", f.ImageMap["a"][0].Url())
	assert.Equal(t, "http://www.diagrams.com/d.jpg", f.ImageMap["a"][1].Url())
	assert.Equal(t, "http://www.photos.com/e.jpg", f.ImageMap["b"][0].Url())
	assert.Equal(t, "http://www.diagrams.com/f.jpg", f.ImageMap["b"][1].Url())

}
