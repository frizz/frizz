package system

import (
	"reflect"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

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
	testInitialiseAnonymousFields(t, unmarshalFunc)
	testInitialiseAnonymousFields(t, unpackFunc)
}
func testInitialiseAnonymousFields(t *testing.T, unpacker unpackerFunc) {

	type ruleStruct struct {
		*Object
		*Rule
	}

	ctx := tests.Context("a.b/c").Jtype("@b", reflect.TypeOf(&ruleStruct{})).Ctx()

	j := `{
		"type": "@b"
	}`
	var i interface{}
	err := unpacker(ctx, []byte(j), &i)
	assert.NoError(t, err)
	rs, ok := i.(*ruleStruct)
	assert.True(t, ok)
	assert.NotNil(t, rs.Object)
	assert.NotNil(t, rs.Rule)

}

func TestRuleWrapperItemsRule(t *testing.T) {
	type parentStruct struct {
		*Object
	}
	type ruleStruct struct {
		*Object
		*Rule
	}
	parentType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "a"), Type: NewReference("kego.io/system", "type")},
		Native: NewString("object"),
	}
	ruleType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "@a"), Type: NewReference("kego.io/system", "type")},
		Native: NewString("object"),
	}

	ctx := tests.Context("a.b/c").
		Jtype("a", reflect.TypeOf(&parentStruct{})).
		Jtype("@a", reflect.TypeOf(&ruleStruct{})).
		Stype("a", parentType).
		Stype("@a", ruleType).Ctx()

	w, err := WrapRule(ctx, &ruleStruct{
		Object: &Object{Type: NewReference("a.b/c", "a")},
		Rule:   &Rule{},
	})
	assert.NoError(t, err)
	_, err = w.ItemsRule()
	assert.IsError(t, err, "VPAGXSTQHM")

	/*
		parentType.Native = NewString("array")

			rh.Rule = "a"
			_, err = rh.ItemsRule()
			// rh.rule must be a pointer or RuleFieldByReflection will error
			assert.IsError(t, err, "LIDXIQYGJD")

			rh.Rule = &struct{}{}
			_, err = rh.ItemsRule()
			// rh.rule needs an Items field
			assert.IsError(t, err, "VYTHGJTSNJ")

			rh.Rule = &struct{ Items int }{Items: 1}
			_, err = rh.ItemsRule()
			// Items must be a rule or NewRuleHolder will error
			assert.IsError(t, err, "FGYMQPNBQJ")
	*/
}
