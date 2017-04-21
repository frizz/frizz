package system

import (
	"testing"

	"github.com/dave/ktest/assert"
)

func TestGoName(t *testing.T) {
	assert.Equal(t, "", GoName(""))
	assert.Equal(t, "Abc", GoName("abc"))
	assert.Equal(t, "AbcDef", GoName("abc-def"))
	assert.Equal(t, "AbcRule", GoName("@abc"))
	assert.Equal(t, "AbcDefRule", GoName("@abc-def"))

	assert.Equal(t, "", GoInterfaceName(""))
	assert.Equal(t, "AbcInterface", GoInterfaceName("abc"))
	assert.Equal(t, "AbcDefInterface", GoInterfaceName("abc-def"))
	assert.Equal(t, "AbcInterface", GoInterfaceName("@abc"))
	assert.Equal(t, "AbcDefInterface", GoInterfaceName("@abc-def"))
}

/*
func TestNoType(t *testing.T) {
	type C struct {
		*Object
		D string
	}
	type A struct {
		*Object
		B *C
	}

	ctx := tests.Context("frizz.io/system").
		Jtype("a", reflect.TypeOf(&A{})).
		Jtype("c", reflect.TypeOf(&C{})).Ctx()

	j := `{
		"type": "a",
		"b": {
			"d": "e"
		}
	}`

	var i interface{}
	err := Unmarshal(ctx, []byte(j), &i)
	require.NoError(t, err)
	a, ok := i.(*A)
	assert.True(t, ok)
	assert.NotNil(t, a.B.Object)
	assert.Equal(t, "frizz.io/system:c", a.B.Type.Value())

	j = `{
		"type": "a",
		"b": {
			"type": "f",
			"d": "e"
		}
	}`

	err = Unmarshal(ctx, []byte(j), &i)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "json: cannot unmarshal frizz.io/system:f into Go value of type system.C")

}


func TestNative(t *testing.T) {

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

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
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

	type Foo struct {
		StrHere    *String `frizz:"{\"default\":{\"type\":\"frizz.io/system:string\",\"value\":\"a\",\"path\":\"frizz.io/system\"}}"`
		NumHere    *Number `frizz:"{\"default\":{\"type\":\"frizz.io/system:number\",\"value\":2,\"path\":\"frizz.io/system\"}}"`
		BolHere    *Bool   `frizz:"{\"default\":{\"type\":\"frizz.io/system:bool\",\"value\":true,\"path\":\"frizz.io/system\"}}"`
		StrDefault *String `frizz:"{\"default\":{\"type\":\"frizz.io/system:string\",\"value\":\"b\",\"path\":\"frizz.io/system\"}}"`
		NumDefault *Number `frizz:"{\"default\":{\"type\":\"frizz.io/system:number\",\"value\":3,\"path\":\"frizz.io/system\"}}"`
		BolDefault *Bool   `frizz:"{\"default\":{\"type\":\"frizz.io/system:bool\",\"value\":true,\"path\":\"frizz.io/system\"}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
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

	type Foo struct {
		StrHere    *String `frizz:"{\"default\":{\"value\":\"a\"}}"`
		NumHere    *Number `frizz:"{\"default\":{\"value\":2}}"`
		BolHere    *Bool   `frizz:"{\"default\":{\"value\":true}}"`
		StrDefault *String `frizz:"{\"default\":{\"value\":\"b\"}}"`
		NumDefault *Number `frizz:"{\"default\":{\"value\":3}}"`
		BolDefault *Bool   `frizz:"{\"default\":{\"value\":true}}"`
	}

	data := `{
		"type": "foo",
		"strHere": "c",
		"numHere": 4,
		"bolHere": false
	}`

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
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

	// If we're generating the type structs, we probably don't have the
	// custom marshal function, so we add the type and context data in
	// the default tag. When the code actually executes, it will have
	// the custom unmarshal function, which will be called along with the
	// context data.

	type Foo struct {

		// The value is just a, so we should be picking up the package
		// frizz.io/b from the local package path
		Ref1 *Reference `frizz:"{\"default\":{\"type\":\"frizz.io/system:reference\",\"value\":\"a\",\"path\":\"frizz.io/b\"}}"`

		// The value is a:b, so a is the package alias for frizz.io/d
		// which we find in the package aliases, and b is the type.
		Ref2 *Reference `frizz:"{\"default\":{\"type\":\"frizz.io/system:reference\",\"value\":\"a:b\",\"path\":\"frizz.io/c\",\"aliases\":{\"a\":\"frizz.io/d\"}}}"`

		// The value is a full type with package path.
		Ref3 *Reference `frizz:"{\"default\":{\"type\":\"frizz.io/system:reference\",\"value\":\"a.b/c:d\",\"path\":\"frizz.io/d\",\"aliases\":{\"c\":\"a.b/c\"}}}"`
	}

	data := `{
		"type": "foo"
	}`

	ctx := tests.Context("frizz.io/system").Alias("c", "a.b/c").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref1)
	assert.Equal(t, "frizz.io/b:a", f.Ref1.Value())
	assert.NotNil(t, f.Ref2)
	assert.Equal(t, "frizz.io/d:b", f.Ref2.Value())
	assert.NotNil(t, f.Ref3)
	assert.Equal(t, "a.b/c:d", f.Ref3.Value())

}

func TestReferenceType(t *testing.T) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "typ"
	}`

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "frizz.io/system:typ", f.Ref.Value())
	assert.Equal(t, "frizz.io/system", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceEmpty(t *testing.T) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo"
	}`

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Nil(t, f.Ref)

}

func TestReferencePath(t *testing.T) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "frizz.io/pkg:typ"
	}`

	ctx := tests.Context("frizz.io/system").Alias("pkg", "frizz.io/pkg").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "frizz.io/pkg:typ", f.Ref.Value())
	assert.Equal(t, "frizz.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceImport(t *testing.T) {

	type Foo struct {
		Ref *Reference
	}

	data := `{
		"type": "foo",
		"ref": "pkg:typ"
	}`

	ctx := tests.Context("frizz.io/system").Alias("pkg", "frizz.io/pkg").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.Ref)
	assert.Equal(t, "frizz.io/pkg:typ", f.Ref.Value())
	assert.Equal(t, "frizz.io/pkg", f.Ref.Package)
	assert.Equal(t, "typ", f.Ref.Name)

}

func TestReferenceDefault(t *testing.T) {

	type Foo struct {
		RefHere    *Reference `frizz:"{\"default\":{\"type\":\"frizz.io/system:reference\",\"value\":\"frizz.io/pkga:typa\",\"path\":\"frizz.io/system\",\"aliases\":{\"pkga\":\"frizz.io/pkga\"}}}"`
		RefDefault *Reference `frizz:"{\"default\":{\"type\":\"frizz.io/system:reference\",\"value\":\"frizz.io/pkgb:typb\",\"path\":\"frizz.io/system\",\"aliases\":{\"pkgb\":\"frizz.io/pkgb\"}}}"`
	}

	data := `{
		"type": "foo",
		"refHere": "typc"
	}`

	ctx := tests.Context("frizz.io/system").Jtype("foo", reflect.TypeOf(&Foo{})).Ctx()

	var i interface{}
	err := Unmarshal(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*Foo)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.NotNil(t, f.RefHere)
	assert.NotNil(t, f.RefDefault)
	assert.Equal(t, "frizz.io/system:typc", f.RefHere.Value())
	assert.Equal(t, "frizz.io/system", f.RefHere.Package)
	assert.Equal(t, "typc", f.RefHere.Name)
	assert.Equal(t, "frizz.io/pkgb:typb", f.RefDefault.Value())
	assert.Equal(t, "frizz.io/pkgb", f.RefDefault.Package)
	assert.Equal(t, "typb", f.RefDefault.Name)

}
*/
