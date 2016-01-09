package system_test

import (
	"testing"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/system"
)

func TestReflectType(t *testing.T) {
	ctx, _, err := process.Initialise(process.FromDefaults{Path: "kego.io/system"})
	assert.NoError(t, err)
	checkReflectType(ctx, t, "kego.io/system", "type", "basic", "bool")
	checkReflectType(ctx, t, "kego.io/system", "type", "embed", "[]*system.Reference")
	checkReflectType(ctx, t, "kego.io/system", "type", "native", "*system.String")
	checkReflectType(ctx, t, "kego.io/system", "type", "interface", "bool")
	checkReflectType(ctx, t, "kego.io/system", "type", "fields", "map[string]system.RuleInterface")
	checkReflectType(ctx, t, "kego.io/system", "type", "rule", "*system.Type")
}

func checkReflectType(ctx context.Context, t *testing.T, path string, name string, field string, output string) {
	cache := cachectx.FromContext(ctx)
	p, ok := cache.Get(path)
	assert.True(t, ok)
	typ, ok := p.Types.Get(name)
	assert.True(t, ok)
	ty, ok := typ.(*system.Type)
	assert.True(t, ok)
	r, ok := ty.Fields[field]
	assert.True(t, ok)
	rh, err := system.WrapRule(ctx, r)
	assert.NoError(t, err)
	rt, err := rh.GetReflectType()
	assert.NoError(t, err)
	assert.Equal(t, output, rt.String())
}
