package system

import (
	"reflect"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
	"kego.io/process/tests/unpacker"
)

func TestRuleWrapperHoldsDisplayType(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty()
	fooType := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}
	fooRule := &fooRuleStruct{Rule: &Rule{}}
	foo := RuleWrapper{
		ctx:       cb.Ctx(),
		Interface: fooRule,
		Struct:    fooRule.Rule,
		Parent:    fooType,
	}

	fooType.Id.Package = "d.e/f"
	_, err := foo.HoldsDisplayType()
	assert.IsError(t, err, "OPIFCOHGWI")

	fooType.Id.Package = "a.b/c"
	val, err := foo.HoldsDisplayType()
	assert.NoError(t, err)
	assert.Equal(t, "foo", val)

	fooRule.Interface = true
	val, err = foo.HoldsDisplayType()
	assert.NoError(t, err)
	assert.Equal(t, "foo*", val)

	fooRule.Interface = false
	fooType.Interface = true
	val, err = foo.HoldsDisplayType()
	assert.NoError(t, err)
	assert.Equal(t, "foo*", val)

}

func TestRuleWrapperItemsRule1(t *testing.T) {
	cb := tests.Context("a.b/c").Jempty()
	fooType := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}
	fooRule := &fooRuleStruct{Rule: &Rule{}}
	foo := RuleWrapper{
		ctx:       cb.Ctx(),
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
		ctx:       cb.Ctx(),
		Interface: barRule,
		Struct:    barRule.Rule,
		Parent:    barType,
	}
	// barRule has nil Items
	_, err = bar.ItemsRule()
	assert.IsError(t, err, "SUJLYBXPYS")

	itemsRule := &fooRuleStruct{Rule: &Rule{}, Object: &Object{Type: NewReference("a.b/c", "@foo")}}

	// Foo system type is not registered, so WrapRule will fail
	bar.ctx = cb.Sempty().Ctx()
	barRule.Items = itemsRule
	_, err = bar.ItemsRule()
	assert.IsError(t, err, "SDSMCXSWOF")

	bar.ctx = cb.Stype("foo", fooType).Ctx()
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
		ctx:       cb.Ctx(),
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

	// Rule isn't a collection rule, so we error
	fooType.Native = NewString("map")
	rt, err = foo.GetReflectType()
	assert.IsError(t, err, "GSYSHQOWNH")

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
		ctx:       cb.Ctx(),
		Interface: barRule,
		Struct:    barRule.Rule,
		Parent:    barType,
	}

	// WrapRule can't find a.b/c:foo
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "EDEMPPVUNW")

	// To get WrapRule to work, but GetReflectType to fail, we create a new contect without the foo
	// reflect types
	bar.ctx = tests.Context("a.b/c").Jempty().Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "LMKEHHWHKL")

	// Finally we get this to work.
	bar.ctx = cb.Stype("foo", fooType).Ctx()
	rt, err = bar.GetReflectType()
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(map[string]fooStruct{}), rt)

	barType.Native = NewString("a")
	rt, err = bar.GetReflectType()
	assert.IsError(t, err, "VDEORSSUWA")

}

func TestRuleGetDefault(t *testing.T) {
	d := DummyRule{Default: "a"}
	assert.Equal(t, "a", d.GetDefault())
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
