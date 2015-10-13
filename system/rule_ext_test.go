package system_test

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/system"
	_ "kego.io/system/types"
)

func TestReflectType(t *testing.T) {
	checkReflectType(t, "kego.io/system", "type", "basic", "bool")
	checkReflectType(t, "kego.io/system", "type", "embed", "[]system.Reference")
	checkReflectType(t, "kego.io/system", "type", "is", "[]system.Reference")
	checkReflectType(t, "kego.io/system", "type", "native", "system.String")
	checkReflectType(t, "kego.io/system", "type", "interface", "bool")
	checkReflectType(t, "kego.io/system", "type", "fields", "map[string]system.RuleInterface")
	checkReflectType(t, "kego.io/system", "type", "rule", "*system.Type")
}

func checkReflectType(t *testing.T, path string, name string, field string, output string) {
	h, ok := system.GetGlobal(path, name)
	assert.True(t, ok)
	ty, ok := h.Object.(*system.Type)
	assert.True(t, ok)
	r, ok := ty.Fields[field]
	assert.True(t, ok)
	rh, err := system.NewRuleHolder(r)
	assert.NoError(t, err)
	rt, err := rh.GetReflectType()
	assert.NoError(t, err)
	assert.Equal(t, output, rt.String())
}
