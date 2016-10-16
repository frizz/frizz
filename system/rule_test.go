package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/tests"
)

func TestRuleWrapper_Kind(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty()
	at := &Type{
		Object: &Object{Id: NewReference("a.b/c", "foo")},
		Native: NewString("object"),
	}
	ar := &fooRuleStruct{Rule: &Rule{}}
	aw := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: ar,
		Struct:    ar.Rule,
		Parent:    at,
	}
	kind, alias := aw.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindStruct, kind)

	aw.Struct.Interface = true
	kind, alias = aw.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindInterface, kind)

	cr := &MapRule{Rule: &Rule{}, Items: &StringRule{Rule: &Rule{}}}
	aw.Interface = cr
	aw.Struct = cr.Rule
	aw.Parent.CustomKind = NewString(string(KindMap))
	kind, alias = aw.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindMap, kind)

	arr := &ArrayRule{Rule: &Rule{}, Items: &StringRule{Rule: &Rule{}}}
	aw.Interface = arr
	aw.Struct = arr.Rule
	aw.Parent.CustomKind = NewString(string(KindArray))
	kind, alias = aw.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindArray, kind)

	// DummyRule always implements CollectionRule, but we don't want to return
	// KindMap/KindArray unless GetItemsRule returns something.
	dr := &DummyRule{Rule: &Rule{}}
	aw.Interface = dr
	aw.Struct = dr.Rule
	aw.Parent.CustomKind = nil
	kind, alias = aw.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, KindStruct, kind)
}

func TestRuleWrapperHoldsDisplayType(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty()
	fooType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "foo")},
		Native: NewString("object"),
	}
	fooRule := &fooRuleStruct{Rule: &Rule{}}
	foo := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: fooRule,
		Struct:    fooRule.Rule,
		Parent:    fooType,
	}

	fooType.Id.Package = "d.e/f"
	_, err := foo.DisplayType()
	assert.IsError(t, err, "OPIFCOHGWI")

	fooType.Id.Package = "a.b/c"
	val, err := foo.DisplayType()
	require.NoError(t, err)
	assert.Equal(t, "foo", val)

	fooRule.Interface = true
	val, err = foo.DisplayType()
	require.NoError(t, err)
	assert.Equal(t, "foo*", val)

	fooRule.Interface = false
	fooType.Interface = true
	val, err = foo.DisplayType()
	require.NoError(t, err)
	assert.Equal(t, "foo*", val)

	bazType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "baz")},
		Native: NewString("object"),
	}
	cb.Stype("baz", bazType)
	barType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "foo")},
		Native: NewString("array"),
	}
	barRule := &barRuleStruct{
		Rule:  &Rule{},
		Items: &IntRule{Rule: &Rule{}, Object: &Object{Type: NewReference("a.b/c", "@baz")}},
	}
	bar := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: barRule,
		Struct:    barRule.Rule,
		Parent:    barType,
	}
	val, err = bar.DisplayType()
	require.NoError(t, err)
	assert.Equal(t, "[]baz", val)

	barType.Native = NewString("map")
	val, err = bar.DisplayType()
	require.NoError(t, err)
	assert.Equal(t, "map[]baz", val)

}

func TestRuleWrapperItemsRule1(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty()
	fooType := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}
	fooRule := &fooRuleStruct{Rule: &Rule{}}
	foo := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: fooRule,
		Struct:    fooRule.Rule,
		Parent:    fooType,
	}

	// String isn't a collection
	fooType.Native = NewString("string")
	_, err := foo.ItemsRule()
	assert.IsError(t, err, "VPAGXSTQHM")

	// FooRule isn't a collection rule
	fooType.Native = NewString("map")
	_, err = foo.ItemsRule()
	assert.IsError(t, err, "TNRVQVJIFH")

	barType := &Type{Object: &Object{Id: NewReference("a.b/c", "bar")}, Native: NewString("map")}
	barRule := &barRuleStruct{
		Object: &Object{},
		Rule:   &Rule{},
	}
	bar := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: barRule,
		Struct:    barRule.Rule,
		Parent:    barType,
	}
	// barRule has nil Items
	_, err = bar.ItemsRule()
	assert.IsError(t, err, "SUJLYBXPYS")

}

type fooRuleStruct struct {
	*Object
	*Rule
}
type fooStruct struct{}
type fooIface interface{}

type barRuleStruct struct {
	*Object
	*Rule
	Items RuleInterface
}

func (r *barRuleStruct) GetItemsRule() RuleInterface {
	return r.Items
}

type barStruct struct{}
type barIface interface{}

/*
func TestRuleGetReflectType(t *testing.T) {

	cb := tests.Context("a.b/c").Jempty()

	fooType := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}
	fooRule := &fooRuleStruct{Rule: &Rule{}}

	foo := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: fooRule,
		Struct:    fooRule.Rule,
		Parent:    fooType,
	}

	fooIfaceReflect := reflect.TypeOf((*fooIface)(nil)).Elem()
	fooStructReflect := reflect.TypeOf(fooStruct{})

	// Native types return the reflect.Type of the parent type, but it's not registered so we error
	fooType.Native = NewString("string")
	rt, err := foo.GetReflectType()
	assert.IsError(t, err, "DLAJJPJDPL")

	// Interface rule, but interface type isn't registered
	fooRule.Interface = true
	_, err = foo.GetReflectType()
	assert.IsError(t, err, "QGUVEUTXAN")

	// Add the interface and it works.
	cb.Jiface("foo", fooIfaceReflect)
	rt, err = foo.GetReflectType()
	require.NoError(t, err)
	assert.Equal(t, rt, fooIfaceReflect)

	// Reset interface to false
	fooRule.Interface = false

	cb.Jtype("foo", fooStructReflect)

	// Native types return the reflect.Type of the parent type
	fooType.Native = NewString("string")
	rt, err = foo.GetReflectType()
	require.NoError(t, err)
	assert.Equal(t, rt, fooStructReflect)

	fooType.Native = NewString("object")
	rt, err = foo.GetReflectType()
	require.NoError(t, err)
	assert.Equal(t, rt, fooStructReflect)

	barType := &Type{Object: &Object{Id: NewReference("a.b/c", "bar")}, Native: NewString("map")}
	barRule := &barRuleStruct{
		Object: &Object{},
		Rule:   &Rule{},
		Items:  &fooRuleStruct{Rule: &Rule{}, Object: &Object{Type: NewReference("a.b/c", "@foo")}},
	}

	cb.Sempty()

	bar := RuleWrapper{
		Ctx:       cb.Ctx(),
		Interface: barRule,
		Struct:    barRule.Rule,
		Parent:    barType,
	}

	// To get WrapRule to work, but GetReflectType to fail, we create a new contect without the foo
	// reflect types
	bar.Ctx = tests.Context("a.b/c").Jempty().Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "LMKEHHWHKL")

	// Finally we get this to work.
	bar.Ctx = cb.Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	require.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(map[string]fooStruct{}), rt)

}
*/
func TestRuleGetDefault(t *testing.T) {
	d := DummyRule{Default: "a"}
	assert.Equal(t, "a", d.GetDefault())
}

func TestDummyRule_GetItemsRule(t *testing.T) {
	i := &IntRule{}
	d := &DummyRule{Items: i}
	assert.Equal(t, i, d.GetItemsRule())
}

func TestWrapRule(t *testing.T) {

	type nonObjectStruct struct {
		*Rule
	}
	type ruleStruct struct {
		*Object
		*Rule
	}
	parentType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "a"), Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "@a"), Type: NewReference("kego.io/system", "type")},
	}

	ctx := tests.Context("a.b/c").Stype("a", parentType).Stype("@a", ruleType).Ctx()

	r := &ruleStruct{
		Object: &Object{Type: NewReference("a.b/c", "@a")},
	}
	w := WrapRule(ctx, r)
	assert.Equal(t, "a", w.Parent.Id.Name)

}

/*
func TestInitialiseAnonymousFields(t *testing.T) {

	type ruleStruct struct {
		*Object
		*Rule
	}

	ctx := tests.Context("a.b/c").Jtype("@b", reflect.TypeOf(&ruleStruct{})).Ctx()

	j := `{
		"type": "@b"
	}`
	var i interface{}
	err := Unmarshal(ctx, []byte(j), &i)
	require.NoError(t, err)
	rs, ok := i.(*ruleStruct)
	assert.True(t, ok)
	assert.NotNil(t, rs.Object)
	assert.NotNil(t, rs.Rule)

}
*/
