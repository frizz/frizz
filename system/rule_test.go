package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
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
	Register("a.b/c", "a", parentType, 0)
	Register("a.b/c", "@a", ruleType, 0)
	defer Unregister("a.b/c", "a")
	defer Unregister("a.b/c", "@a")

	r := &ruleStruct{
		Object: &Object{Type: NewReference("a.b/c", "@a")},
	}
	w, err := WrapRule(r)
	assert.NoError(t, err)
	assert.Equal(t, "a", w.Type.Id.Name)

	r1 := nonObjectStruct{}
	w, err = WrapRule(r1)
	// A non Object rule will cause an error
	assert.IsError(t, err, "VKFNPJDNVB")

	r = &ruleStruct{
		Object: &Object{Type: NewReference("a.b/c", "unregistered")},
	}
	w, err = WrapRule(r)
	// An unregistered type will cause WrapRule to return an error
	assert.IsError(t, err, "KYCTDXKFYR")

}

func TestInitialiseAnonymousFields(t *testing.T) {
	testInitialiseAnonymousFields(t, unmarshalFunc)
	testInitialiseAnonymousFields(t, unpackFunc)

	// needs types
	//testInitialiseAnonymousFields(t, repackFunc)
}
func testInitialiseAnonymousFields(t *testing.T, unpacker unpackerFunc) {

	type ruleStruct struct {
		*Object
		*Rule
	}

	json.Register("a.b/c", "@b", reflect.TypeOf(&ruleStruct{}), nil, 0)
	defer json.Unregister("a.b/c", "@b")
	j := `{
		"type": "@b"
	}`
	var i interface{}
	err := unpacker([]byte(j), &i, "a.b/c", map[string]string{})
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
	}
	ruleType := &Type{
		Object: &Object{Id: NewReference("a.b/c", "@a"), Type: NewReference("kego.io/system", "type")},
	}
	json.Register("a.b/c", "a", reflect.TypeOf(&parentStruct{}), nil, 0)
	json.Register("a.b/c", "@a", reflect.TypeOf(&ruleStruct{}), nil, 0)
	Register("a.b/c", "a", parentType, 0)
	Register("a.b/c", "@a", ruleType, 0)
	defer json.Unregister("a.b/c", "a")
	defer json.Unregister("a.b/c", "@a")
	defer Unregister("a.b/c", "a")
	defer Unregister("a.b/c", "@a")

	w, err := WrapRule(&ruleStruct{Object: &Object{Type: NewReference("a.b/c", "a")}})
	assert.NoError(t, err)
	_, err = w.ItemsRule()
	assert.IsError(t, err, "VPAGXSTQHM")

	parentType.Native = NewString("array")
	/*
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
func TestRuleFieldByReflection(t *testing.T) {

	_, _, _, err := RuleFieldByReflection(struct{}{}, "")
	// Must be a pointer type
	assert.IsError(t, err, "QOYMWPXWUO")

	i := 1
	_, _, _, err = RuleFieldByReflection(&i, "")
	// Must point to a struct type
	assert.IsError(t, err, "IGOUOBGXAN")

	_, _, ok, err := RuleFieldByReflection(&struct{}{}, "A")
	// Must have field called a, or it returns false
	assert.NoError(t, err)
	assert.False(t, ok)

	_, _, ok, err = RuleFieldByReflection(&struct{ A *string }{A: nil}, "A")
	// If field is a pointer, and nil, returns false
	assert.NoError(t, err)
	assert.False(t, ok)

	_, _, ok, err = RuleFieldByReflection(&struct{ A string }{A: ""}, "A")
	// If field is not a pointer, and equal to the zero value for that type, returns false
	assert.NoError(t, err)
	assert.False(t, ok)

	_, _, ok, err = RuleFieldByReflection(&struct {
		A struct {
			B int
			C string
		}
	}{A: struct {
		B int
		C string
	}{B: 0, C: ""}}, "A")
	// If field is not a pointer, and equal to the zero value for that type, returns false
	assert.NoError(t, err)
	assert.False(t, ok)

	fi, fpi, ok, err := RuleFieldByReflection(&struct{ A string }{A: "b"}, "A")
	assert.NoError(t, err)
	assert.True(t, ok)
	f, ok := fi.(string)
	assert.True(t, ok)
	assert.Equal(t, "b", f)
	fp, ok := fpi.(*string)
	assert.True(t, ok)
	assert.Equal(t, "b", *fp)

	b := "b"
	fi, fpi, ok, err = RuleFieldByReflection(&struct{ A *string }{A: &b}, "A")
	assert.NoError(t, err)
	assert.True(t, ok)
	fp, ok = fi.(*string)
	assert.True(t, ok)
	assert.Equal(t, "b", *fp)
	fp, ok = fpi.(*string)
	assert.True(t, ok)
	assert.Equal(t, "b", *fp)

}
