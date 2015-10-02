package system

import (
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestRuleTypes(t *testing.T) {

	type nonRuleStruct struct {
		*Rule_base
	}
	type ruleStruct struct {
		*Object_base
		*Rule_base
	}
	parentType := &Type{
		Object_base: &Object_base{Id: NewReference("a.b/c", "a"), Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object_base: &Object_base{Id: NewReference("a.b/c", "@a"), Type: NewReference("kego.io/system", "type")},
	}
	Register("a.b/c", "a", parentType, 0)
	Register("a.b/c", "@a", ruleType, 0)
	defer Unregister("a.b/c", "a")
	defer Unregister("a.b/c", "@a")

	r := &ruleStruct{
		Object_base: &Object_base{Type: NewReference("a.b/c", "@a")},
	}
	rt, pt, err := ruleTypes(r)
	assert.NoError(t, err)
	assert.Equal(t, "a", pt.Id.Name)
	assert.Equal(t, "@a", rt.Id.Name)

	r1 := nonRuleStruct{}
	rt, pt, err = ruleTypes(r1)
	// A non Object rule will cause ruleTypeReference to return an error
	assert.IsError(t, err, "BNEKIFYDDL")

	r = &ruleStruct{
		Object_base: &Object_base{Type: NewReference("a.b/c", "unregistered")},
	}
	rt, pt, err = ruleTypes(r)
	// An unregistered type will cause ruleReference.GetType to return an error
	assert.IsError(t, err, "PFGWISOHRR")

	r = &ruleStruct{
		Object_base: &Object_base{Type: NewReference("a.b/c", "a")},
	}
	rt, pt, err = ruleTypes(r)
	// A rule with a non rule type will cause ruleReference.RuleToParentType to error
	assert.IsError(t, err, "NXRCPQMUIE")

	Register("a.b/c", "@b", ruleType, 0)
	defer Unregister("a.b/c", "@b")
	r = &ruleStruct{
		Object_base: &Object_base{Type: NewReference("a.b/c", "@b")},
	}
	rt, pt, err = ruleTypes(r)
	// An rule type with an unregistered parent type typeReference.GetType to return an error
	assert.IsError(t, err, "KYCTDXKFYR")

}

func TestInitialiseAnonymousFields(t *testing.T) {
	testInitialiseAnonymousFields(t, unmarshalFunc)
	testInitialiseAnonymousFields(t, unpackFunc)
}
func testInitialiseAnonymousFields(t *testing.T, unpacker unpackerFunc) {

	type ruleStruct struct {
		*Object_base
		*Rule_base
	}

	json.Register("a.b/c", "@b", reflect.TypeOf(&ruleStruct{}), 0)
	defer json.Unregister("a.b/c", "@b")
	j := `{
		"type": "@b"
	}`
	var i interface{}
	err := unpacker([]byte(j), &i, "a.b/c", map[string]string{})
	assert.NoError(t, err)
	rs, ok := i.(*ruleStruct)
	assert.True(t, ok)
	assert.NotNil(t, rs.Object_base)
	assert.NotNil(t, rs.Rule_base)

}

func TestRuleTypeReference(t *testing.T) {

	type ruleStruct struct {
		*Object_base
		*Rule_base
	}
	rs := &ruleStruct{
		Object_base: &Object_base{Type: NewReference("a.b/c", "@a")},
	}
	r, err := ruleTypeReference(rs)
	assert.NoError(t, err)
	assert.Equal(t, "a.b/c:@a", r.Value())

	/*
		ri := map[string]interface{}{}
		r, err = ruleTypeReference(ri, "", map[string]string{})
		assert.IsError(t, err, "OLHOVKXEXN")

		ri = map[string]interface{}{
			"type": 1, //not a string
		}
		r, err = ruleTypeReference(ri, "", map[string]string{})
		assert.IsError(t, err, "IILEXGQDXL")

		ri = map[string]interface{}{
			"type":     "a:b", // package will not be registered so UnmarshalJSON will error
			"_path":    "a.b/c",
			"_imports": map[string]string{},
		}
		r, err = ruleTypeReference(ri, "", map[string]string{})
		assert.IsError(t, err, "QBTHPRVBWN")

		ri = map[string]interface{}{
			"type": "a.b/c:@a",
		}
		r, err = ruleTypeReference(rs, "", map[string]string{})
		assert.NoError(t, err)
		assert.Equal(t, "a.b/c:@a", r.Value)
	*/

	rsp := ruleStruct{}
	r, err = ruleTypeReference(rsp)
	// rsp has no base, so ruleTypeReference will return a zero base
	assert.NoError(t, err)
	assert.False(t, r.Exists)

	/*
		type structWithoutType struct{}
		rwt := &structWithoutType{}
		r, err = ruleTypeReference(rwt, "", map[string]string{})
		assert.IsError(t, err, "NXYRAJITEV")

		type structWithIntType struct {
			Type int
		}
		rwi := &structWithIntType{
			Type: 1,
		}
		r, err = ruleTypeReference(rwi, "", map[string]string{})
		assert.IsError(t, err, "FHUPSRTRFE")
	*/
}

func TestRuleHolderItemsRule(t *testing.T) {
	type parentStruct struct {
		*Object_base
	}
	type ruleStruct struct {
		*Object_base
		*Rule_base
	}
	parentType := &Type{
		Object_base: &Object_base{Id: NewReference("a.b/c", "a"), Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object_base: &Object_base{Id: NewReference("a.b/c", "@a"), Type: NewReference("kego.io/system", "type")},
	}
	json.Register("a.b/c", "a", reflect.TypeOf(&parentStruct{}), 0)
	json.Register("a.b/c", "@a", reflect.TypeOf(&ruleStruct{}), 0)
	Register("a.b/c", "a", parentType, 0)
	Register("a.b/c", "@a", ruleType, 0)
	defer json.Unregister("a.b/c", "a")
	defer json.Unregister("a.b/c", "@a")
	defer Unregister("a.b/c", "a")
	defer Unregister("a.b/c", "@a")

	rh := &RuleHolder{
		Rule:       &ruleStruct{},
		RuleType:   ruleType,
		ParentType: parentType,
	}
	_, err := rh.ItemsRule()
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
