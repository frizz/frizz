package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestNative(t *testing.T) {

	type Foo struct {
		StrHere  String
		StrEmpty String
		StrNull  String
		NumHere  Number
		NumEmpty Number
		NumNull  Number
		BolHere  Bool
		BolEmpty Bool
		BolNull  Bool
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

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.StrHere.Exists)
	assert.True(t, f.NumHere.Exists)
	assert.True(t, f.BolHere.Exists)
	assert.False(t, f.StrEmpty.Exists)
	assert.False(t, f.NumEmpty.Exists)
	assert.False(t, f.BolEmpty.Exists)
	assert.False(t, f.StrNull.Exists)
	assert.False(t, f.NumNull.Exists)
	assert.False(t, f.BolNull.Exists)
	assert.Equal(t, "a", f.StrHere.Value)
	assert.Equal(t, 2.0, f.NumHere.Value)
	assert.Equal(t, true, f.BolHere.Value)

}

func TestNativeDefaults(t *testing.T) {

	type Foo struct {
		StrHere    String `kego:"{\"default\":{\"type\":\"kego.io/system:string\",\"value\":\"a\",\"path\":\"kego.io/system\"}}"`
		NumHere    Number `kego:"{\"default\":{\"type\":\"kego.io/system:number\",\"value\":2,\"path\":\"kego.io/system\"}}"`
		BolHere    Bool   `kego:"{\"default\":{\"type\":\"kego.io/system:bool\",\"value\":true,\"path\":\"kego.io/system\"}}"`
		StrDefault String `kego:"{\"default\":{\"type\":\"kego.io/system:string\",\"value\":\"b\",\"path\":\"kego.io/system\"}}"`
		NumDefault Number `kego:"{\"default\":{\"type\":\"kego.io/system:number\",\"value\":3,\"path\":\"kego.io/system\"}}"`
		BolDefault Bool   `kego:"{\"default\":{\"type\":\"kego.io/system:bool\",\"value\":true,\"path\":\"kego.io/system\"}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.StrHere.Exists)
	assert.True(t, f.NumHere.Exists)
	assert.True(t, f.BolHere.Exists)
	assert.True(t, f.StrDefault.Exists)
	assert.True(t, f.NumDefault.Exists)
	assert.True(t, f.BolDefault.Exists)
	assert.Equal(t, "c", f.StrHere.Value)
	assert.Equal(t, 4.0, f.NumHere.Value)
	assert.Equal(t, false, f.BolHere.Value)
	assert.Equal(t, "b", f.StrDefault.Value)
	assert.Equal(t, 3.0, f.NumDefault.Value)
	assert.Equal(t, true, f.BolDefault.Value)

}

func TestNativeDefaultsShort(t *testing.T) {

	type Foo struct {
		StrHere    String `kego:"{\"default\":{\"value\":\"a\"}}"`
		NumHere    Number `kego:"{\"default\":{\"value\":2}}"`
		BolHere    Bool   `kego:"{\"default\":{\"value\":true}}"`
		StrDefault String `kego:"{\"default\":{\"value\":\"b\"}}"`
		NumDefault Number `kego:"{\"default\":{\"value\":3}}"`
		BolDefault Bool   `kego:"{\"default\":{\"value\":true}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.StrHere.Exists)
	assert.True(t, f.NumHere.Exists)
	assert.True(t, f.BolHere.Exists)
	assert.True(t, f.StrDefault.Exists)
	assert.True(t, f.NumDefault.Exists)
	assert.True(t, f.BolDefault.Exists)
	assert.Equal(t, "c", f.StrHere.Value)
	assert.Equal(t, 4.0, f.NumHere.Value)
	assert.Equal(t, false, f.BolHere.Value)
	assert.Equal(t, "b", f.StrDefault.Value)
	assert.Equal(t, 3.0, f.NumDefault.Value)
	assert.Equal(t, true, f.BolDefault.Value)

}

func TestDefaultCustomUnmarshal(t *testing.T) {

	// If we're generating the type structs, we probably don't have the
	// custom marshal function, so we add the type and context data in
	// the default tag. When the code actually executes, it will have
	// the custom unmarshal function, which will be called along with the
	// context data.

	type Foo struct {

		// The value is just a, so we should be picking up the package
		// kego.io/b from the local package path
		Ref1 Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a\",\"path\":\"kego.io/b\",\"imports\":{\"json\":\"kego.io/json\",\"system\":\"kego.io/system\"}}}"`

		// The value is a:b, so a is the package alias for kego.io/d
		// which we find in the package imports, and b is the type.
		Ref2 Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a:b\",\"path\":\"kego.io/c\",\"imports\":{\"json\":\"kego.io/json\",\"a\":\"kego.io/d\",\"system\":\"kego.io/system\"}}}"`

		// The value is a full type with package path.
		Ref3 Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"a.b/c:d\",\"path\":\"kego.io/d\",\"imports\":{\"json\":\"kego.io/json\",\"system\":\"kego.io/system\"}}}"`
	}

	data := `{
		"type": "foo"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref1.Exists)
	assert.Equal(t, "kego.io/b:a", f.Ref1.Value)
	assert.True(t, f.Ref2.Exists)
	assert.Equal(t, "kego.io/d:b", f.Ref2.Value)
	assert.True(t, f.Ref3.Exists)
	assert.Equal(t, "a.b/c:d", f.Ref3.Value)

}

func TestReferenceType(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "typ"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, "kego.io/system:typ", f.Ref.Value)
	assert.Equal(t, "kego.io/system", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Type)

}

func TestReferenceEmpty(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.False(t, f.Ref.Exists)

}

func TestReferencePath(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "kego.io/pkg:typ"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, "kego.io/pkg:typ", f.Ref.Value)
	assert.Equal(t, "kego.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Type)

}

func TestReferenceImport(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "pkg:typ"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{"pkg": "kego.io/pkg"})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, "kego.io/pkg:typ", f.Ref.Value)
	assert.Equal(t, "kego.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Type)

}

func TestReferenceDefault(t *testing.T) {

	type Foo struct {
		RefHere    Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"kego.io/pkga:typa\",\"path\":\"kego.io/system\",\"imports\":{\"json\":\"kego.io/json\",\"system\":\"kego.io/system\"}}}"`
		RefDefault Reference `kego:"{\"default\":{\"type\":\"kego.io/system:reference\",\"value\":\"kego.io/pkgb:typb\",\"path\":\"kego.io/system\",\"imports\":{\"json\":\"kego.io/json\",\"system\":\"kego.io/system\"}}}"`
	}

	data := `{
		"type": "foo",
		"refHere": "typc"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.RefHere.Exists)
	assert.True(t, f.RefDefault.Exists)
	assert.Equal(t, "kego.io/system:typc", f.RefHere.Value)
	assert.Equal(t, "kego.io/system", f.RefHere.Package)
	assert.Equal(t, "typc", f.RefHere.Type)
	assert.Equal(t, "kego.io/pkgb:typb", f.RefDefault.Value)
	assert.Equal(t, "kego.io/pkgb", f.RefDefault.Package)
	assert.Equal(t, "typb", f.RefDefault.Type)

}

func TestContext(t *testing.T) {

	type Foo struct {
		*Base
		Bar string
	}

	data := `{
		"type": "foo",
		"bar": "a"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	// Clean up for the tests - don't normally need to unregister types
	defer json.UnregisterType("kego.io/system:foo")

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{"d": "e.f/g"})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "a", f.Bar)
	assert.Equal(t, "kego.io/system", f.Context.Package)
	assert.Equal(t, "e.f/g", f.Context.Imports["d"])

}
