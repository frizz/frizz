package system

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/json"
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

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
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
	assert.Equal(t, f.StrHere.Value, "a")
	assert.Equal(t, f.NumHere.Value, 2.0)
	assert.Equal(t, f.BolHere.Value, true)

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

}

func TestNativeDefaults(t *testing.T) {

	type Foo struct {
		StrHere    String `kego:"{\"default\": \"a\"}"`
		NumHere    Number `kego:"{\"default\": 2}"`
		BolHere    Bool   `kego:"{\"default\": true}"`
		StrDefault String `kego:"{\"default\": \"b\"}"`
		NumDefault Number `kego:"{\"default\": 3}"`
		BolDefault Bool   `kego:"{\"default\": true}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
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
	assert.Equal(t, f.StrHere.Value, "c")
	assert.Equal(t, f.NumHere.Value, 4.0)
	assert.Equal(t, f.BolHere.Value, false)
	assert.Equal(t, f.StrDefault.Value, "b")
	assert.Equal(t, f.NumDefault.Value, 3.0)
	assert.Equal(t, f.BolDefault.Value, true)

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

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

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "kego.io/system:typ")
	assert.Equal(t, f.Ref.Package, "kego.io/system")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

}

func TestReferenceEmpty(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.False(t, f.Ref.Exists)

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

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

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "kego.io/pkg:typ")
	assert.Equal(t, f.Ref.Package, "kego.io/pkg")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

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

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{"pkg": "kego.io/pkg"})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "kego.io/pkg:typ")
	assert.Equal(t, f.Ref.Package, "kego.io/pkg")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

}

func TestReferenceDefault(t *testing.T) {

	type Foo struct {
		RefHere    Reference `kego:"{\"default\": \"kego.io/pkga:typa\"}"`
		RefDefault Reference `kego:"{\"default\": \"kego.io/pkgb:typb\"}"`
	}

	data := `{
		"type": "foo",
		"refHere": "typc"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.RefHere.Exists)
	assert.True(t, f.RefDefault.Exists)
	assert.Equal(t, f.RefHere.Value, "kego.io/system:typc")
	assert.Equal(t, f.RefHere.Package, "kego.io/system")
	assert.Equal(t, f.RefHere.Type, "typc")
	assert.Equal(t, f.RefDefault.Value, "kego.io/pkgb:typb")
	assert.Equal(t, f.RefDefault.Package, "kego.io/pkgb")
	assert.Equal(t, f.RefDefault.Type, "typb")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

}

func TestContext(t *testing.T) {

	type Foo struct {
		*Object
		Bar string
	}

	data := `{
		"type": "foo",
		"bar": "a"
	}`

	json.RegisterType("kego.io/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{"d": "e.f/g"})
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, f.Bar, "a")
	assert.Equal(t, f.Context.Package.Value, "kego.io/system")
	assert.Equal(t, f.Context.Imports["d"].Value, "e.f/g")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("kego.io/system:foo")

}
