package system

import (
	"reflect"
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/tests"
	"kego.io/tests/unpacker"
)

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
	assert.NoError(t, err)
	assert.Equal(t, "foo", val)

	fooRule.Interface = true
	val, err = foo.DisplayType()
	assert.NoError(t, err)
	assert.Equal(t, "foo*", val)

	fooRule.Interface = false
	fooType.Interface = true
	val, err = foo.DisplayType()
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	assert.Equal(t, "[]baz", val)

	barType.Native = NewString("map")
	val, err = bar.DisplayType()
	assert.NoError(t, err)
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

	itemsRule := &fooRuleStruct{Rule: &Rule{}, Object: &Object{Type: NewReference("a.b/c", "@foo")}}

	// Foo system type is not registered, so WrapRule will fail
	bar.Ctx = cb.Sempty().Ctx()
	barRule.Items = itemsRule
	_, err = bar.ItemsRule()
	assert.IsError(t, err, "SDSMCXSWOF")

	bar.Ctx = cb.Stype("foo", fooType).Ctx()
	rw, err := bar.ItemsRule()
	assert.NoError(t, err)
	assert.Equal(t, rw.Interface, itemsRule)

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
	assert.NoError(t, err)
	assert.Equal(t, rt, fooIfaceReflect)

	// Reset interface to false
	fooRule.Interface = false

	cb.Jtype("foo", fooStructReflect)

	// Native types return the reflect.Type of the parent type
	fooType.Native = NewString("string")
	rt, err = foo.GetReflectType()
	assert.NoError(t, err)
	assert.Equal(t, rt, fooStructReflect)

	fooType.Native = NewString("object")
	rt, err = foo.GetReflectType()
	assert.NoError(t, err)
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

	// WrapRule can't find a.b/c:foo
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "EDEMPPVUNW")

	// To get WrapRule to work, but GetReflectType to fail, we create a new contect without the foo
	// reflect types
	bar.Ctx = tests.Context("a.b/c").Jempty().Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "LMKEHHWHKL")

	// Finally we get this to work.
	bar.Ctx = cb.Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(map[string]fooStruct{}), rt)

}

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
	w, err := WrapRule(ctx, r)
	assert.NoError(t, err)
	assert.Equal(t, "a", w.Parent.Id.Name)

	r1 := nonObjectStruct{}
	w, err = WrapRule(ctx, r1)
	// A non Object rule will cause an error
	assert.IsError(t, err, "VKFNPJDNVB")

	r = &ruleStruct{
		Object: &Object{Type: NewReference("a.b/c", "unregistered")},
	}
	w, err = WrapRule(ctx, r)
	// An unregistered type will cause WrapRule to return an error
	assert.IsError(t, err, "KYCTDXKFYR")

}

func TestInitialiseAnonymousFields(t *testing.T) {
	testInitialiseAnonymousFields(t, unpacker.Unmarshal)
	testInitialiseAnonymousFields(t, unpacker.Unpack)
	testInitialiseAnonymousFields(t, unpacker.Decode)
}
func testInitialiseAnonymousFields(t *testing.T, up unpacker.Interface) {

	type ruleStruct struct {
		*Object
		*Rule
	}

	ctx := tests.Context("a.b/c").Jtype("@b", reflect.TypeOf(&ruleStruct{})).Ctx()

	j := `{
		"type": "@b"
	}`
	var i interface{}
	err := up.Process(ctx, []byte(j), &i)
	assert.NoError(t, err)
	rs, ok := i.(*ruleStruct)
	assert.True(t, ok)
	assert.NotNil(t, rs.Object)
	assert.NotNil(t, rs.Rule)

}
