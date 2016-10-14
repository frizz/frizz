package system_test

import (
	"testing"

	"context"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/sysctx"
	"kego.io/process"
	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/tests"
)

func TestRuleWrapper_InnerType(t *testing.T) {
	cb := tests.Context("kego.io/system").Jauto().Sauto(parser.Parse)
	r := system.WrapRule(cb.Ctx(), &system.MapRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@map")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})
	inner := r.InnerType(cb.Ctx())
	assert.Equal(t, "kego.io/system:string", inner.Id.String())

	kind, alias := r.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, kind, system.KindMap)

	assert.False(t, r.PassedAsPointer(cb.Ctx()))

	r = system.WrapRule(cb.Ctx(), &system.DummyRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@int")},
		Rule:   &system.Rule{},
	})
	assert.Equal(t, "kego.io/system:int", r.InnerType(cb.Ctx()).Id.String())

	kind, alias = r.Kind(cb.Ctx())
	assert.True(t, alias)
	assert.Equal(t, kind, system.KindValue)

	assert.True(t, r.PassedAsPointer(cb.Ctx()))

	r = system.WrapRule(cb.Ctx(), &system.DummyRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@int")},
		Rule: &system.Rule{
			Interface: true,
		},
	})
	assert.Equal(t, "kego.io/system:int", r.InnerType(cb.Ctx()).Id.String())

	kind, alias = r.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, kind, system.KindInterface)

	assert.False(t, r.PassedAsPointer(cb.Ctx()))

	r = system.WrapRule(cb.Ctx(), &system.DummyRule{
		Object: &system.Object{Type: system.NewReference("kego.io/json", "@string")},
		Rule:   &system.Rule{},
	})
	assert.Equal(t, "kego.io/json:string", r.InnerType(cb.Ctx()).Id.String())

	assert.False(t, r.PassedAsPointer(cb.Ctx()))

	r = system.WrapRule(cb.Ctx(), &system.DummyRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@package")},
		Rule:   &system.Rule{},
	})
	assert.Equal(t, "kego.io/system:package", r.InnerType(cb.Ctx()).Id.String())

	kind, alias = r.Kind(cb.Ctx())
	assert.False(t, alias)
	assert.Equal(t, kind, system.KindStruct)

	assert.True(t, r.PassedAsPointer(cb.Ctx()))
}

func TestRuleWrapper_ZeroValue(t *testing.T) {
	cb := tests.Context("kego.io/system").Jauto().Sauto(parser.Parse)
	r := system.WrapRule(cb.Ctx(), &system.MapRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@map")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})

	v, err := r.ZeroValue(true)
	require.NoError(t, err)
	assert.IsType(t, map[string]*system.String{}, v.Interface())
	assert.Nil(t, v.Interface())

	v, err = r.ZeroValue(false)
	require.NoError(t, err)
	assert.IsType(t, map[string]*system.String{}, v.Interface())
	assert.NotNil(t, v.Interface())
	vv := v.Interface().(map[string]*system.String)
	vv["a"] = system.NewString("")

	r = system.WrapRule(cb.Ctx(), &system.MapRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@array")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})

	v, err = r.ZeroValue(true)
	require.NoError(t, err)
	assert.IsType(t, []*system.String{}, v.Interface())
	assert.Nil(t, v.Interface())

	v, err = r.ZeroValue(false)
	require.NoError(t, err)
	assert.IsType(t, []*system.String{}, v.Interface())
	assert.NotNil(t, v.Interface())
	va := v.Interface().([]*system.String)
	va = append(va, system.NewString(""))
}

func TestReflectType(t *testing.T) {
	ctx, _, err := process.Initialise(context.Background(), process.Options{Path: "kego.io/system"})
	require.NoError(t, err)
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
	ty, ok := typ.Type.(*system.Type)
	assert.True(t, ok)
	r, ok := ty.Fields[field]
	assert.True(t, ok)
	rh := system.WrapRule(ctx, r)
	rt, err := rh.GetReflectType()
	require.NoError(t, err)
	assert.Equal(t, output, rt.String())
}
