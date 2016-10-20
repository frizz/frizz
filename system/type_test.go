package system

import (
	"testing"

	"reflect"

	"context"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/tests"
)

func TestType_Kind(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty().Sempty()
	ty := &Type{
		Object: &Object{Id: NewReference("a.b/c", "foo")},
		Native: NewString("object"),
	}
	kind, alias := ty.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindStruct, kind)

	ty.CustomKind = NewString(string(KindValue))
	kind, alias = ty.Kind(cb.Ctx())
	assert.True(t, alias)
	assert.Equal(t, KindValue, kind)

	ty.Interface = true
	kind, alias = ty.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindInterface, kind)

	cb.StypePath("kego.io/json", "string", &Type{
		Object:     &Object{Id: NewReference("a.b/c", "bar")},
		Native:     NewString("string"),
		CustomKind: NewString("value"),
	})
	ty.Alias = JsonStringRule{Rule: &Rule{}, Object: &Object{Type: NewReference("kego.io/json", "@string")}}
	kind, alias = ty.Kind(cb.Ctx())
	assert.True(t, alias)
	assert.Equal(t, KindValue, kind)

	ty.Fields = map[string]RuleInterface{}
	ty.Fields["a"] = &JsonStringRule{}
	kind, alias = ty.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindStruct, kind)

	ty.Id.Package = "kego.io/json"
	kind, alias = ty.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindValue, kind)
}

func TestFieldOrigins(t *testing.T) {

	ty := &Type{
		Object: &Object{Id: NewReference("a.b/c", "d")},
		Embed: []*Reference{
			NewReference("e.f/g", "h"),
			NewReference("i.j/k", "l"),
		}}
	o := ty.FieldOrigins()
	assert.Equal(t, 4, len(o))
	assert.Equal(t, *NewReference("a.b/c", "d"), *o[0])
	assert.Equal(t, *NewReference("e.f/g", "h"), *o[1])
	assert.Equal(t, *NewReference("i.j/k", "l"), *o[2])
	assert.Equal(t, *NewReference("kego.io/system", "object"), *o[3])

	ty.Basic = true
	o = ty.FieldOrigins()
	assert.Equal(t, 3, len(o))
	assert.Equal(t, *NewReference("a.b/c", "d"), *o[0])
	assert.Equal(t, *NewReference("e.f/g", "h"), *o[1])
	assert.Equal(t, *NewReference("i.j/k", "l"), *o[2])

}

func TestTypeSortedFields(t *testing.T) {
	ty := &Type{
		Object: &Object{Id: NewReference("a.b/c", "d")},
		Fields: map[string]RuleInterface{
			"c": &StringRule{},
			"b": &StringRule{},
			"a": &StringRule{},
		},
	}
	f := ty.SortedFields()
	assert.Equal(t, "a", f[0].Name)
	assert.Equal(t, "b", f[1].Name)
	assert.Equal(t, "c", f[2].Name)
}

func TestNativeValueGolangType(t *testing.T) {
	test := func(expected string, in string) {
		ty := &Type{Native: NewString(in)}
		v, err := ty.NativeValueGolangType()
		require.NoError(t, err)
		assert.Equal(t, expected, v)
	}
	test("float64", "number")
	test("string", "string")
	test("bool", "bool")
	testErr := func(in string) {
		ty := &Type{Native: NewString(in)}
		_, err := ty.NativeValueGolangType()
		assert.IsError(t, err, "TXQIDRBJRH")
	}
	testErr("map")
	testErr("object")
	testErr("array")
	testErr("null")
	testErr("")
	testErr("foo")
}

func TestIsNativeObject(t *testing.T) {
	test := func(expected bool, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.IsNativeObject(), in)
	}
	test(false, "number")
	test(false, "string")
	test(false, "bool")
	test(false, "map")
	test(true, "object")
	test(false, "array")
	test(true, "null")
	test(true, "")
	test(true, "foo")
}

func TestIsNativeArray(t *testing.T) {
	test := func(expected bool, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.IsNativeArray(), in)
	}
	test(false, "number")
	test(false, "string")
	test(false, "bool")
	test(false, "map")
	test(false, "object")
	test(true, "array")
	test(false, "null")
	test(false, "")
	test(false, "foo")
}

func TestIsNativeMap(t *testing.T) {
	test := func(expected bool, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.IsNativeMap(), in)
	}
	test(false, "number")
	test(false, "string")
	test(false, "bool")
	test(true, "map")
	test(false, "object")
	test(false, "array")
	test(false, "null")
	test(false, "")
	test(false, "foo")
}

func TestIsNativeCollection(t *testing.T) {
	test := func(expected bool, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.IsNativeCollection(), in)
	}
	test(false, "number")
	test(false, "string")
	test(false, "bool")
	test(true, "map")
	test(false, "object")
	test(true, "array")
	test(false, "null")
	test(false, "")
	test(false, "foo")
}

func TestIsNativeValue(t *testing.T) {
	test := func(expected bool, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.IsNativeValue(), in)
	}
	test(true, "number")
	test(true, "string")
	test(true, "bool")
	test(false, "map")
	test(false, "object")
	test(false, "array")
	test(false, "null")
	test(false, "")
	test(false, "foo")
}

func TestIsJsonValue(t *testing.T) {
	test := func(expected bool, in *Reference, na JsonType) {
		ty := &Type{Object: &Object{Id: in}, Native: NewString(string(na))}
		assert.Equal(t, expected, ty.IsJsonValue())
	}
	test(false, NewReference("a.b/c", "d"), J_STRING)
	test(false, NewReference("kego.io/json", "a"), J_OBJECT)
	test(false, NewReference("kego.io/json", "a"), J_NULL)
	test(false, NewReference("kego.io/json", "a"), J_MAP)
	test(true, NewReference("kego.io/json", "a"), J_STRING)
}

func TestNativeJsonType(t *testing.T) {
	cb := tests.New()
	test := func(expected JsonType, in string) {
		ty := &Type{Native: NewString(in)}
		assert.Equal(t, expected, ty.NativeJsonType(cb.Ctx()), in)
	}
	test(J_NUMBER, "number")
	test(J_STRING, "string")
	test(J_BOOL, "bool")
	test(J_MAP, "map")
	test(J_OBJECT, "object")
	test(J_ARRAY, "array")
	test(J_NULL, "null")
	test(J_NULL, "")
	test(J_NULL, "foo")
}

func TestGetTypeFromCache(t *testing.T) {
	d := &Type{Object: &Object{Id: NewReference("a.b/c", "d")}}
	cb := tests.Context("a.b/c").Stype("d", d)
	ty, ok := GetTypeFromCache(cb.Ctx(), "a.b/c", "d")
	assert.True(t, ok)
	assert.Equal(t, d, ty)

	ty, ok = GetTypeFromCache(cb.Ctx(), "a.b/c", "e")
	assert.False(t, ok)

	ty, ok = GetTypeFromCache(cb.Ctx(), "f", "g")
	assert.False(t, ok)
}

type iFoo interface {
	foo()
}
type iBar interface {
	bar()
}
type tFoo struct {
	*Object
}

func (*tFoo) foo() {

}

type tFooBar struct{}

func (*tFooBar) foo() {

}
func (*tFooBar) bar() {

}

func Test(t *testing.T) {

}

/*
func TestGetAllTypesThatImplementInterface(t *testing.T) {

	ibar := &Type{Object: &Object{Id: NewReference("a.b/c", "ibar")}, Interface: true}
	tfoo := &Type{Object: &Object{Id: NewReference("a.b/c", "tfoo")}, Native: NewString("object")}
	tfoobar := &Type{Object: &Object{Id: NewReference("a.b/c", "tfoobar")}, Native: NewString("object")}

	cb := tests.
		Context("a.b/c").
		Jtype("ibar", reflect.TypeOf((*iBar)(nil)).Elem()).Stype("ibar", ibar).
		JtypeIface("tfoo", reflect.TypeOf(&tFoo{}), reflect.TypeOf((*iFoo)(nil)).Elem()).Stype("tfoo", tfoo).
		Jtype("tfoobar", reflect.TypeOf(&tFooBar{})).Stype("tfoobar", tfoobar)

	types := GetAllTypesThatImplementInterface(cb.Ctx(), ibar)
	assert.Equal(t, 1, len(types))
	assert.Equal(t, "tfoobar", types[0].Id.Name)

	types = GetAllTypesThatImplementInterface(cb.Ctx(), tfoo)
	assert.Equal(t, 2, len(types))
	assert.Equal(t, "tfoo", types[0].Id.Name)
	assert.Equal(t, "tfoobar", types[1].Id.Name)

	types = GetAllTypesThatImplementInterface(cb.Ctx(), &Type{Object: &Object{Id: NewReference("a.b/c", "d")}})
	assert.Nil(t, types)

	types = GetAllTypesThatImplementInterface(cb.Ctx(), &Type{Object: &Object{Id: NewReference("a.b/c", "d")}, Interface: true})
	assert.Nil(t, types)

	rw := &RuleWrapper{Ctx: cb.Ctx(), Interface: nil, Parent: tfoo, Struct: &Rule{Interface: false}}
	types = rw.PermittedTypes()
	assert.Equal(t, 1, len(types))
	assert.Equal(t, "tfoo", types[0].Id.Name)

	rw = &RuleWrapper{Ctx: cb.Ctx(), Interface: nil, Parent: tfoo, Struct: &Rule{Interface: true}}
	types = rw.PermittedTypes()
	assert.Equal(t, 2, len(types))
	assert.Equal(t, "tfoo", types[0].Id.Name)
	assert.Equal(t, "tfoobar", types[1].Id.Name)
}
*/

/*
type tInt int

func TestZeroValue(t *testing.T) {

	ty := &Type{Native: NewString("map")}
	_, err := ty.ZeroValue(context.Background(), true)
	assert.IsError(t, err, "PGUHCGBJWE")

	tfoo := &Type{Object: &Object{Id: NewReference("a.b/c", "tfoo")}, Native: NewString("object")}
	cb := tests.
		Context("a.b/c").
		Jtype("tfoo", reflect.TypeOf(&tFoo{})).Stype("tfoo", tfoo)

	i, err := tfoo.ZeroValue(cb.Ctx(), true)
	require.NoError(t, err)
	tfn, ok := i.Interface().(*tFoo)
	assert.True(t, ok)
	assert.Nil(t, tfn)

	i, err = tfoo.ZeroValue(cb.Ctx(), false)
	require.NoError(t, err)
	tf, ok := i.Interface().(*tFoo)
	assert.True(t, ok)
	assert.NotNil(t, tf)
	assert.NotNil(t, tf.Object)

	tnil := &Type{Object: &Object{Id: NewReference("a.b/c", "d")}, Native: NewString("object")}
	i, err = tnil.ZeroValue(cb.Ctx(), true)
	assert.IsError(t, err, "RSWTEOTNBD")

	tint := &Type{Object: &Object{Id: NewReference("a.b/c", "tint")}, Native: NewString("number")}
	cb.Jtype("tint", reflect.TypeOf((*tInt)(nil))).Stype("tint", tint)

	i, err = tint.ZeroValue(cb.Ctx(), true)
	require.NoError(t, err)
	tin, ok := i.Interface().(*tInt)
	assert.True(t, ok)
	assert.Nil(t, tin)

	i, err = tint.ZeroValue(cb.Ctx(), false)
	require.NoError(t, err)
	ti, ok := i.Interface().(*tInt)
	assert.True(t, ok)
	assert.Equal(t, tInt(0), *ti)

	tjstr := &Type{Object: &Object{Id: NewReference("a.b/c", "tjstr")}, Native: NewString("string")}
	cb.Jtype("tjstr", reflect.TypeOf("")).Stype("tjstr", tjstr)

	i, err = tjstr.ZeroValue(cb.Ctx(), true)
	require.NoError(t, err)
	tjs, ok := i.Interface().(string)
	assert.True(t, ok)
	assert.Equal(t, "", tjs)

	i, err = tjstr.ZeroValue(cb.Ctx(), false)
	require.NoError(t, err)
	tjs, ok = i.Interface().(string)
	assert.True(t, ok)
	assert.Equal(t, "", tjs)

}
*/

/*
func TestTypeImplements(t *testing.T) {
	tfoo := &Type{Object: &Object{Id: NewReference("a.b/c", "tfoo")}, Native: NewString("object")}
	cb := tests.
		Context("a.b/c").
		Jtype("tfoo", reflect.TypeOf(&tFoo{})).Stype("tfoo", tfoo)

	assert.False(t, tfoo.Implements(cb.Ctx(), reflect.TypeOf((*iBar)(nil)).Elem()))
	assert.True(t, tfoo.Implements(cb.Ctx(), reflect.TypeOf((*iFoo)(nil)).Elem()))
	tnil := &Type{Object: &Object{Id: NewReference("a.b/c", "tnil")}, Native: NewString("object")}
	assert.False(t, tnil.Implements(cb.Ctx(), reflect.TypeOf((*iFoo)(nil)).Elem()))
}
*/

func TestNativeGoType(t *testing.T) {
	n, err := nativeGoType("string")
	require.NoError(t, err)
	assert.Equal(t, "string", n)

	n, err = nativeGoType("number")
	require.NoError(t, err)
	assert.Equal(t, "float64", n)

	n, err = nativeGoType("bool")
	require.NoError(t, err)
	assert.Equal(t, "bool", n)

	_, err = nativeGoType("a")
	assert.IsError(t, err, "TXQIDRBJRH")
}

func TestTypeIsNativeType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n := y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("array")}
	n = y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("object")}
	n = y.IsNativeObject()
	assert.True(t, n)

	// Will also return true for any other value
	y = &Type{Native: NewString("a")}
	n = y.IsNativeObject()
	assert.True(t, n)
}

func TestTypeNativeValueGolangType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n, err := y.NativeValueGolangType()
	require.NoError(t, err)
	assert.Equal(t, "bool", n)
}

func TestTypeGoName(t *testing.T) {
	y := &Type{Object: &Object{Id: NewReference("a.b/c", "aa")}}
	n := y.GoName()
	assert.Equal(t, "Aa", n)
}

func TestType_Implements(t *testing.T) {
	ty := Type{Native: NewString("map")}
	assert.False(t, ty.Implements(context.Background(), (reflect.Type)(nil)))
}
