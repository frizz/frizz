package system_test

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"golang.org/x/net/context"
	"kego.io/context/sysctx"
	"kego.io/process"
	"kego.io/process/parser"
	"kego.io/process/tests"
	"kego.io/system"
)

func TestRuleWrapper_ZeroValue(t *testing.T) {
	cb := tests.Context("kego.io/system").Jauto().Sauto(parser.Parse)
	r, err := system.WrapRule(cb.Ctx(), &system.MapRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@map")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})
	require.NoError(t, err)
	v, err := r.ZeroValue()
	require.NoError(t, err)
	assert.IsType(t, map[string]*system.String{}, v)
}

func TestReflectType(t *testing.T) {
	ctx, _, err := process.Initialise(context.Background(), process.Options{Path: "kego.io/system"})
	assert.NoError(t, err)
	checkReflectType(ctx, t, "kego.io/system", "type", "basic", "bool")
	checkReflectType(ctx, t, "kego.io/system", "type", "embed", "[]*system.Reference")
	checkReflectType(ctx, t, "kego.io/system", "type", "native", "*system.String")
	checkReflectType(ctx, t, "kego.io/system", "type", "interface", "bool")
	checkReflectType(ctx, t, "kego.io/system", "type", "fields", "map[string]system.RuleInterface")
	checkReflectType(ctx, t, "kego.io/system", "type", "rule", "*system.Type")
}

func checkReflectType(ctx context.Context, t *testing.T, path string, name string, field string, output string) {
	scache := sysctx.FromContext(ctx)
	p, ok := scache.Get(path)
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
