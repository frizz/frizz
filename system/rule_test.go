package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/uerr"
)

func TestRuleTypes(t *testing.T) {

	type ruleStruct struct {
		*Object
	}
	parentType := &Type{
		Object: &Object{Id: "a", Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object: &Object{Id: "@a", Type: NewReference("kego.io/system", "type")},
	}
	RegisterType("a.b/c:a", parentType)
	RegisterType("a.b/c:@a", ruleType)
	defer UnregisterType("a.b/c:a")
	defer UnregisterType("a.b/c:@a")

	r := &ruleStruct{
		&Object{Type: NewReference("a.b/c", "@a")},
	}
	rt, pt, err := ruleTypes(r, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "a", pt.Id)
	assert.Equal(t, "@a", rt.Id)

	r1 := ruleStruct{}
	rt, pt, err = ruleTypes(r1, "", map[string]string{})
	// A non pointer rule will cause ruleTypeReference to return an error
	uerr.Assert(t, err, "BNEKIFYDDL")

	r = &ruleStruct{
		&Object{Type: NewReference("a.b/c", "unregistered")},
	}
	rt, pt, err = ruleTypes(r, "", map[string]string{})
	// An unregistered type will cause ruleReference.GetType to return an error
	uerr.Assert(t, err, "PFGWISOHRR")

	r = &ruleStruct{
		&Object{Type: NewReference("a.b/c", "a")},
	}
	rt, pt, err = ruleTypes(r, "", map[string]string{})
	// A rule with a non rule type will cause ruleReference.RuleToParentType to error
	uerr.Assert(t, err, "NXRCPQMUIE")

	RegisterType("a.b/c:@b", ruleType)
	r = &ruleStruct{
		&Object{Type: NewReference("a.b/c", "@b")},
	}
	rt, pt, err = ruleTypes(r, "", map[string]string{})
	// An rule type with an unregistered parent type typeReference.GetType to return an error
	uerr.Assert(t, err, "KYCTDXKFYR")

}
