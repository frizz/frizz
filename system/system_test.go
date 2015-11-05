package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestNoType(t *testing.T) {
	testNoType(t, unmarshalFunc)
	testNoType(t, unpackFunc)
}
func testNoType(t *testing.T, unpacker unpackerFunc) {
	type C struct {
		*Object
		D string
	}
	type A struct {
		*Object
		B *C
	}

	json.Register("kego.io/system", "a", reflect.TypeOf(&A{}), nil, 0)
	json.Register("kego.io/system", "c", reflect.TypeOf(&C{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "a")
	defer json.Unregister("kego.io/system", "c")

	j := `{
		"type": "a",
		"b": {
			"d": "e"
		}
	}`

	var i interface{}
	err := unpacker([]byte(j), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok)
	assert.NotNil(t, a.B.Object)
	assert.Equal(t, "kego.io/system:c", a.B.Type.Value())

	j = `{
		"type": "a",
		"b": {
			"type": "f",
			"d": "e"
		}
	}`

	err = unpacker([]byte(j), &i, "kego.io/system", map[string]string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "json: cannot unmarshal kego.io/system:f into Go value of type system.C")

}

func TestNative(t *testing.T) {
	testNative(t, unmarshalFunc)
	testNative(t, unpackFunc)
}
func testNative(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		StrHere  *String
		StrEmpty *String
		StrNull  *String
		NumHere  *Number
		NumEmpty *Number
		NumNull  *Number
		BolHere  *Bool
		BolEmpty *Bool
		BolNull  *Bool
	}

	data := `{
		"type": "foo",
		"strHere": "a",
		"strNull": null,
		"numHere": 2,
		"numNull": null,
		"bolHere": true,
		"bolNull": null
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.StrHere)
	assert.NotNil(t, f.NumHere)
	assert.NotNil(t, f.BolHere)
	assert.Nil(t, f.StrEmpty)
	assert.Nil(t, f.NumEmpty)
	assert.Nil(t, f.BolEmpty)
	assert.Nil(t, f.StrNull)
	assert.Nil(t, f.NumNull)
	assert.Nil(t, f.BolNull)
	assert.Equal(t, "a", f.StrHere.Value())
	assert.Equal(t, 2.0, f.NumHere.Value())
	assert.Equal(t, true, f.BolHere.Value())

}

func TestNativeDefaults(t *testing.T) {
	testNativeDefaults(t, unmarshalFunc)
	testNativeDefaults(t, unpackFunc)
}
func testNativeDefaults(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		StrHere    *String `kego:"{\"default\":{\"type\":\"kego.io/system:string\",\"value\":\"a\",\"path\":\"kego.io/system\"}}"`
		NumHere    *Number `kego:"{\"default\":{\"type\":\"kego.io/system:number\",\"value\":2,\"path\":\"kego.io/system\"}}"`
		BolHere    *Bool   `kego:"{\"default\":{\"type\":\"kego.io/system:bool\",\"value\":true,\"path\":\"kego.io/system\"}}"`
		StrDefault *String `kego:"{\"default\":{\"type\":\"kego.io/system:string\",\"value\":\"b\",\"path\":\"kego.io/system\"}}"`
		NumDefault *Number `kego:"{\"default\":{\"type\":\"kego.io/system:number\",\"value\":3,\"path\":\"kego.io/system\"}}"`
		BolDefault *Bool   `kego:"{\"default\":{\"type\":\"kego.io/system:bool\",\"value\":true,\"path\":\"kego.io/system\"}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.StrHere)
	assert.NotNil(t, f.NumHere)
	assert.NotNil(t, f.BolHere)
	assert.NotNil(t, f.StrDefault)
	assert.NotNil(t, f.NumDefault)
	assert.NotNil(t, f.BolDefault)
	assert.Equal(t, "c", f.StrHere.Value())
	assert.Equal(t, 4.0, f.NumHere.Value())
	assert.Equal(t, false, f.BolHere.Value())
	assert.Equal(t, "b", f.StrDefault.Value())
	assert.Equal(t, 3.0, f.NumDefault.Value())
	assert.Equal(t, true, f.BolDefault.Value())

}

func TestNativeDefaultsShort(t *testing.T) {
	testNativeDefaultsShort(t, unmarshalFunc)
	testNativeDefaultsShort(t, unpackFunc)
}
func testNativeDefaultsShort(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		StrHere    *String `kego:"{\"default\":{\"value\":\"a\"}}"`
		NumHere    *Number `kego:"{\"default\":{\"value\":2}}"`
		BolHere    *Bool   `kego:"{\"default\":{\"value\":true}}"`
		StrDefault *String `kego:"{\"default\":{\"value\":\"b\"}}"`
		NumDefault *Number `kego:"{\"default\":{\"value\":3}}"`
		BolDefault *Bool   `kego:"{\"default\":{\"value\":true}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.StrHere)
	assert.NotNil(t, f.NumHere)
	assert.NotNil(t, f.BolHere)
	assert.NotNil(t, f.StrDefault)
	assert.NotNil(t, f.NumDefault)
	assert.NotNil(t, f.BolDefault)
	assert.Equal(t, "c", f.StrHere.Value())
	assert.Equal(t, 4.0, f.NumHere.Value())
	assert.Equal(t, false, f.BolHere.Value())
	assert.Equal(t, "b", f.StrDefault.Value())
	assert.Equal(t, 3.0, f.NumDefault.Value())
	assert.Equal(t, true, f.BolDefault.Value())

}

func TestDefaultCustomUnmarshal(t *testing.T) {
	testDefaultCustomUnmarshal(t, unmarshalFunc)
	testDefaultCustomUnmarshal(t, unpackFunc)
}
func testDefaultCustomUnmarshal(t *testing.T, unpacker unpackerFunc) {

	// If we're generating the type structs, we probably don't have the
	// custom marshal function, so we add the type and context data in
	// the default tag. When the code actually executes, it will have
	// the custom unmarshal function, which will be called along with the
	// context data.

	type Foo struct {

		// The value is just a, so we should be picking up the package
		// kego.io/b from the local package path
		Ref1 *Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a\",\"path\":\"kego.io/b\"}}"`

		// The value is a:b, so a is the package alias for kego.io/d
		// which we find in the package aliases, and b is the type.
		Ref2 *Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a:b\",\"path\":\"kego.io/c\",\"aliases\":{\"kego.io/d\":\"a\"}}}"`

		// The value is a full type with package path.
		Ref3 *Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a.b/c:d\",\"path\":\"kego.io/d\",\"aliases\":{\"a.b/c\":\"c\"}}}"`
	}

	data := `{
		"type": "foo"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{"a.b/c": "c"})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref1)
	assert.Equal(t, "kego.io/b:a", f.Ref1.Value())
	assert.NotNil(t, f.Ref2)
	assert.Equal(t, "kego.io/d:b", f.Ref2.Value())
	assert.NotNil(t, f.Ref3)
	assert.Equal(t, "a.b/c:d", f.Ref3.Value())

}

func TestReferenceType(t *testing.T) {
	testReferenceType(t, unmarshalFunc)
	testReferenceType(t, unpackFunc)
}
func testReferenceType(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "typ"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "kego.io/system:typ", f.Ref.Value())
	assert.Equal(t, "kego.io/system", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceEmpty(t *testing.T) {
	testReferenceEmpty(t, unmarshalFunc)
	testReferenceEmpty(t, unpackFunc)
}
func testReferenceEmpty(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Nil(t, f.Ref)

}

func TestReferencePath(t *testing.T) {
	testReferencePath(t, unmarshalFunc)
	testReferencePath(t, unpackFunc)
}
func testReferencePath(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "kego.io/pkg:typ"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{"kego.io/pkg": "pkg"})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "kego.io/pkg:typ", f.Ref.Value())
	assert.Equal(t, "kego.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceImport(t *testing.T) {
	testReferenceImport(t, unmarshalFunc)
	testReferenceImport(t, unpackFunc)
}
func testReferenceImport(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "pkg:typ"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{"kego.io/pkg": "pkg"})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "kego.io/pkg:typ", f.Ref.Value())
	assert.Equal(t, "kego.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceDefault(t *testing.T) {
	testReferenceDefault(t, unmarshalFunc)
	testReferenceDefault(t, unpackFunc)
}
func testReferenceDefault(t *testing.T, unpacker unpackerFunc) {

	type Foo struct {
		RefHere    *Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"kego.io/pkga:typa\",\"path\":\"kego.io/system\",\"aliases\":{\"kego.io/pkga\":\"pkga\"}}}"`
		RefDefault *Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"kego.io/pkgb:typb\",\"path\":\"kego.io/system\",\"aliases\":{\"kego.io/pkgb\":\"pkgb\"}}}"`
	}

	data := `{
		"type": "foo",
		"refHere": "typc"
	}`

	json.Register("kego.io/system", "foo", reflect.TypeOf(&Foo{}), nil, 0)

	// Clean up for the tests - don't normally need to unregister types
	defer json.Unregister("kego.io/system", "foo")

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.RefHere)
	assert.NotNil(t, f.RefDefault)
	assert.Equal(t, "kego.io/system:typc", f.RefHere.Value())
	assert.Equal(t, "kego.io/system", f.RefHere.Package)
	assert.Equal(t, "typc", f.RefHere.Name)
	assert.Equal(t, "kego.io/pkgb:typb", f.RefDefault.Value())
	assert.Equal(t, "kego.io/pkgb", f.RefDefault.Package)
	assert.Equal(t, "typb", f.RefDefault.Name)

}
