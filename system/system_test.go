package system

import (
	"reflect"
	"testing"

	"github.com/kego/json"
	"github.com/stretchr/testify/assert"
)

var defaultSystemContext = &json.Context{
	PackageName: "system",
	PackagePath: "github.com/kego/system",
	Imports:     map[string]string{},
}

func TestNative(t *testing.T) {

	type Foo struct {
		StrHere  String
		StrEmpty String
		NumHere  Number
		NumEmpty Number
		BolHere  Bool
		BolEmpty Bool
	}

	data := `{
		"type": "foo",
		"strHere": "a",
		"numHere": 2,
		"bolHere": true
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
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
	assert.Equal(t, f.StrHere.Value, "a")
	assert.Equal(t, f.NumHere.Value, 2.0)
	assert.Equal(t, f.BolHere.Value, true)

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

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

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
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
	json.UnregisterType("github.com/kego/system:foo")

}

func TestReferenceType(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "typ"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "github.com/kego/system:typ")
	assert.Equal(t, f.Ref.Package, "github.com/kego/system")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

}

func TestReferenceEmpty(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.False(t, f.Ref.Exists)

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

}

func TestReferencePath(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "github.com/kego/pkg:typ"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "github.com/kego/pkg:typ")
	assert.Equal(t, f.Ref.Package, "github.com/kego/pkg")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

}

func TestReferenceImport(t *testing.T) {

	type Foo struct {
		Ref Reference
	}

	data := `{
		"type": "foo",
		"ref": "pkg:typ"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	contextWithImport := &json.Context{
		PackageName: "system",
		PackagePath: "github.com/kego/system",
		Imports: map[string]string{
			"pkg": "github.com/kego/pkg",
		},
	}

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, contextWithImport)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Ref.Exists)
	assert.Equal(t, f.Ref.Value, "github.com/kego/pkg:typ")
	assert.Equal(t, f.Ref.Package, "github.com/kego/pkg")
	assert.Equal(t, f.Ref.Type, "typ")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

}

func TestReferenceDefault(t *testing.T) {

	type Foo struct {
		RefHere    Reference `kego:"{\"default\": \"github.com/kego/pkga:typa\"}"`
		RefDefault Reference `kego:"{\"default\": \"github.com/kego/pkgb:typb\"}"`
	}

	data := `{
		"type": "foo",
		"refHere": "typc"
	}`

	json.RegisterType("github.com/kego/system:foo", reflect.TypeOf(&Foo{}))

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, defaultSystemContext)
	assert.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.RefHere.Exists)
	assert.True(t, f.RefDefault.Exists)
	assert.Equal(t, f.RefHere.Value, "github.com/kego/system:typc")
	assert.Equal(t, f.RefHere.Package, "github.com/kego/system")
	assert.Equal(t, f.RefHere.Type, "typc")
	assert.Equal(t, f.RefDefault.Value, "github.com/kego/pkgb:typb")
	assert.Equal(t, f.RefDefault.Package, "github.com/kego/pkgb")
	assert.Equal(t, f.RefDefault.Type, "typb")

	// Clean up for the tests - don't normally need to unregister types
	json.UnregisterType("github.com/kego/system:foo")

}
