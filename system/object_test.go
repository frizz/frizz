package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/process/tests"
)

func TestObjectGetType(t *testing.T) {
	ty := &Type{Object: &Object{Id: NewReference("a.b/c", "foo")}}

	o := &Object{Type: NewReference("a.b/c", "t")}

	gt, ok := o.Type.GetType(tests.Context("a.b/c").Stype("t", ty).Ctx())
	assert.True(t, ok)
	assert.Equal(t, "a.b/c:foo", gt.Id.Value())

}

func TestRulesApplyToObjects(t *testing.T) {

	//_, isRule := object.(RuleInterface)
	//_, isType := object.(*Type)
	//_, isObject := object.(ObjectInterface)
	//return !isRule && !isType && isObject

	r := &StringRule{}
	assert.False(t, RulesApplyToObjects(r)) // isRule = true

	ty := &Type{}
	assert.False(t, RulesApplyToObjects(ty)) // isType = true

	ruleBase := &Rule{}
	assert.False(t, RulesApplyToObjects(ruleBase)) // isObject = false

	s := NewString("")
	assert.False(t, RulesApplyToObjects(s)) // isObject = false

	p := &Package{}
	assert.True(t, RulesApplyToObjects(p))
}
