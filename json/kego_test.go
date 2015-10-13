package json

import (
	"fmt"
	"reflect"
	"testing"

	"kego.io/kerr/assert"
)

type unpackerFunc func([]byte, *interface{}, string, map[string]string) error

func unmarshalFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	return Unmarshal(data, i, path, aliases)
}
func unpackFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	var j interface{}
	if err := UnmarshalPlain(data, &j); err != nil {
		return err
	}
	return Unpack(NewJsonUnpacker(j), i, path, aliases)
}

func TestDecodeSimple(t *testing.T) {
	testDecodeSimple(t, unmarshalFunc)
	testDecodeSimple(t, unpackFunc)
}
func testDecodeSimple(t *testing.T, unpacker unpackerFunc) {

	test := func(data string) {

		type Foo struct {
			Foo string
			Bar float64
			Baz bool
		}

		Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)

		// Clean up for the tests - don't normally need to unregister types
		defer Unregister("kego.io/json", "foo")

		var i interface{}
		err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testDecodeDefaults(t, unmarshalFunc)
	testDecodeDefaults(t, unpackFunc)
}
func testDecodeDefaults(t *testing.T, unpacker unpackerFunc) {

	test := func(data string, strExpected string, numExpected float64, boolExpected bool) {

		type Foo struct {
			Foo string  `kego:"{\"default\":{\"value\":\"a\"}}"`
			Bar float64 `kego:"{\"default\":{\"value\":2}}"`
			Baz bool    `kego:"{\"default\":{\"value\":true}}"`
		}

		Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)

		// Clean up for the tests - don't normally need to unregister types
		defer Unregister("kego.io/json", "foo")

		var i interface{}
		err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testDecodeCollections(t, unmarshalFunc)
	testDecodeCollections(t, unpackFunc)
}
func testDecodeCollections(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testDecodeEmbed(t, unmarshalFunc)
	testDecodeEmbed(t, unpackFunc)
}
func testDecodeEmbed(t *testing.T, unpacker unpackerFunc) {

	test := func(data string) {

		type Bar struct {
			String string
			Number float64
			Bool   bool
		}

		type Foo struct {
			Embed Bar
		}

		Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
		Register("kego.io/json", "bar", reflect.TypeOf(&Bar{}), nil, 0)

		// Clean up for the tests - don't normally need to unregister types
		defer Unregister("kego.io/json", "foo")
		defer Unregister("kego.io/json", "bar")

		var i interface{}
		err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testDecodeEmbedCollections(t, unmarshalFunc)
	testDecodeEmbedCollections(t, unpackFunc)
}
func testDecodeEmbedCollections(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "bar", reflect.TypeOf(&Bar{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "bar")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testDecodeComposition(t, unmarshalFunc)
	testDecodeComposition(t, unpackFunc)
}
func testDecodeComposition(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "base", reflect.TypeOf(&Base{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "base")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testInterface(t, unmarshalFunc)
	testInterface(t, unpackFunc)
}
func testInterface(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "http://www.photos.com/a.jpg", f.Img.Url())

}

func TestNilInterface(t *testing.T) {
	testNilInterface(t, unmarshalFunc)
	testNilInterface(t, unpackFunc)
}
func testNilInterface(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	p, ok := f.Iface.(*Photo)
	assert.True(t, ok, "Type %T not correct", f.Iface)
	assert.Equal(t, "http://www.photos.com/a.jpg", p.Url())

}

func TestInterfaceCollections(t *testing.T) {
	testInterfaceCollections(t, unmarshalFunc)
	testInterfaceCollections(t, unpackFunc)
}
func testInterfaceCollections(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
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
	testInterfaceCollectionsComplex(t, unmarshalFunc)
	testInterfaceCollectionsComplex(t, unpackFunc)
}
func testInterfaceCollectionsComplex(t *testing.T, unpacker unpackerFunc) {

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

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "http://www.photos.com/c.jpg", f.ImageMap["a"][0].Url())
	assert.Equal(t, "http://www.diagrams.com/d.jpg", f.ImageMap["a"][1].Url())
	assert.Equal(t, "http://www.photos.com/e.jpg", f.ImageMap["b"][0].Url())
	assert.Equal(t, "http://www.diagrams.com/f.jpg", f.ImageMap["b"][1].Url())

}

func TestDummyInterfaceNotFound(t *testing.T) {
	testDummyInterfaceNotFound(t, unmarshalFunc)
	testDummyInterfaceNotFound(t, unpackFunc)
}
func testDummyInterfaceNotFound(t *testing.T, unpacker unpackerFunc) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		Img Image
	}

	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	data := `{
		"type": "foo",
		"img": {
			"type": "bar",
			"id": "a"
		}
	}`
	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown type kego.io/json:bar")
	assert.True(t, i.(*Foo).Img == nil)

	data = `{
			"type": "foo",
			"img": {
				"type": "foo:bar",
				"id": "a"
			}
		}`
	err = unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown package foo")

	data = `{
			"type": "foo",
			"img": {
				"type": "a.b/c:bar",
				"id": "a"
			}
		}`
	err = unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown package a.b/c")

	err = unpacker([]byte(data), &i, "kego.io/json", map[string]string{"a.b/c": "d"})
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown type a.b/c:bar")
}

type dummyImage struct{}

func (d *dummyImage) Url() string {
	return ""
}

func TestDummyInterface(t *testing.T) {
	testDummyInterface(t, unmarshalFunc)
	testDummyInterface(t, unpackFunc)
}
func testDummyInterface(t *testing.T, unpacker unpackerFunc) {

	type Image interface {
		Url() string
	}

	type Foo struct {
		Img Image
	}

	RegisterInterface(reflect.TypeOf((*Image)(nil)).Elem(), reflect.TypeOf(&dummyImage{}))
	Register("kego.io/json", "foo", reflect.TypeOf(&Foo{}), nil, 0)
	Register("kego.io/json", "photo", reflect.TypeOf(&Photo{}), nil, 0)
	Register("kego.io/json", "diagram", reflect.TypeOf(&Diagram{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer UnregisterInterface(reflect.TypeOf((*Image)(nil)).Elem())
	defer Unregister("kego.io/json", "foo")
	defer Unregister("kego.io/json", "photo")
	defer Unregister("kego.io/json", "diagram")

	data := `{
		"type": "foo",
		"img": {
			"type": "bar",
			"id": "a"
		}
	}`
	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/json", map[string]string{})
	assert.NoError(t, err)
}
