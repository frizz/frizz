package system_test

import (
	"context"
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/ke"
	"kego.io/system"
	"kego.io/tests/data"
	. "kego.io/tests/marshal"
)

func TestAlajs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"alajs": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Alajs, 2)
			assert.Equal(t, "a", v.(*data.Multi).Alajs[0])
			assert.Equal(t, "b", v.(*data.Multi).Alajs[1])
		}),
	)
}

func TestAlas(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"alas": [
				{"type":"simple", "js": "a"},
				{"type":"simple", "js": "b"}
			]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Alas, 2)
			assert.Equal(t, "a", v.(*data.Multi).Alas[0].Js)
			assert.Equal(t, "b", v.(*data.Multi).Alas[1].Js)
		}),
	)
}

func TestAlass(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"alass": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Alass, 2)
			assert.Equal(t, "a", v.(*data.Multi).Alass[0].String())
			assert.Equal(t, "b", v.(*data.Multi).Alass[1].String())
		}),
	)
}

func TestAjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ajs": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ajs, 2)
			assert.Equal(t, "a", v.(*data.Multi).Ajs[0])
			assert.Equal(t, "b", v.(*data.Multi).Ajs[1])
		}),
	)
}

func TestAjn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ajn": [1.1, 1.2]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ajn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Ajn[0])
			assert.Equal(t, 1.2, v.(*data.Multi).Ajn[1])
		}),
	)
}

func TestAjb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ajb": [true, false]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ajb, 2)
			assert.Equal(t, true, v.(*data.Multi).Ajb[0])
			assert.Equal(t, false, v.(*data.Multi).Ajb[1])
		}),
	)
}

func TestAss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ass": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ass, 2)
			assert.Equal(t, "a", v.(*data.Multi).Ass[0].String())
			assert.Equal(t, "b", v.(*data.Multi).Ass[1].String())
		}),
	)
}

func TestAsn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"asn": [1.1, 1.2]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Asn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Asn[0].Value())
			assert.Equal(t, 1.2, v.(*data.Multi).Asn[1].Value())
		}),
	)
}

func TestAsb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"asb": [true, false]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Asb, 2)
			assert.Equal(t, true, v.(*data.Multi).Asb[0].Value())
			assert.Equal(t, false, v.(*data.Multi).Asb[1].Value())
		}),
	)
}

func TestAsr(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"asr": ["a", "system:b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Asr, 2)
			assert.Equal(t, "kego.io/tests/data:a", v.(*data.Multi).Asr[0].Value())
			assert.Equal(t, "kego.io/system:b", v.(*data.Multi).Asr[1].Value())
		}),
	)
}

func TestAsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"asi": [2, 3]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Asi, 2)
			assert.Equal(t, 2, v.(*data.Multi).Asi[0].Value())
			assert.Equal(t, 3, v.(*data.Multi).Asi[1].Value())
		}),
	)
}

func TestAsp(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"asp": [{
				"type": "system:package",
				"recursive": true
			},{
				"type": "system:package"
			}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Asp, 2)
			assert.Equal(t, true, v.(*data.Multi).Asp[0].Recursive)
			assert.Equal(t, false, v.(*data.Multi).Asp[1].Recursive)
		}),
	)
}

func TestAri(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ari": [{
				"type": "system:@string",
				"max-length": 2
			},{
				"type": "system:@bool",
				"default": true
			}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ari, 2)
			assert.Equal(t, 2, v.(*data.Multi).Ari[0].(*system.StringRule).MaxLength.Value())
			assert.Equal(t, true, v.(*data.Multi).Ari[1].(*system.BoolRule).Default.Value())
		}),
	)
}

func TestAlajsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Note: array alias types in an interface can be expressed without
	// explicit type notation.
	Run(t, ctx, `{
			"type": "multi",
			"alajsi": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Len(t, v.(*data.Multi).Alajsi.GetAlajs(ctx), 2)
			assert.Equal(t, "a", v.(*data.Multi).Alajsi.GetAlajs(ctx)[0])
			assert.Equal(t, "b", v.(*data.Multi).Alajsi.GetAlajs(ctx)[1])
		}),
	)
}

func TestAlajsiExplicit(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Note: array alias types in an interface can be expressed without
	// explicit type notation. Explicit type notation is permitted, but it
	// will marshal back to the bare form.
	Run(t, ctx, `{
			"type": "multi",
			"alajsi": {"type":"alajs", "value": ["a", "b"]}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Len(t, v.(*data.Multi).Alajsi.GetAlajs(ctx), 2)
			assert.Equal(t, "a", v.(*data.Multi).Alajsi.GetAlajs(ctx)[0])
			assert.Equal(t, "b", v.(*data.Multi).Alajsi.GetAlajs(ctx)[1])
		}),
		MarshalledString(`{
			"type": "multi",
			"alajsi": ["a", "b"]
		}`),
	)
}

func TestAnri(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"anri": [
				"a",
				{"type": "system:int", "value": 1},
				{"type": "alajs", "value": ["a", "b", "c"]},
				{"type": "almjs", "value": {"a": "b", "c": "d"}}
			]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Anri, 4)
			assert.Equal(t, "a", v.(*data.Multi).Anri[0].GetString(ctx).Value())
			assert.Equal(t, "1", v.(*data.Multi).Anri[1].GetString(ctx).Value())
			assert.Equal(t, "abc", v.(*data.Multi).Anri[2].GetString(ctx).Value())
			assert.Equal(t, "abcd", v.(*data.Multi).Anri[3].GetString(ctx).Value())
		}),
	)
}

func TestAi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"ai": [{
				"type": "facea",
				"a": "a1"
			},{
				"type": "faceb",
				"b": "b1"
			}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Ai, 2)
			assert.Equal(t, "a1", v.(*data.Multi).Ai[0].Face())
			assert.Equal(t, "b1", v.(*data.Multi).Ai[1].Face())
		}),
	)
}

func TestAm(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"am": [{
				"type": "multi",
				"js": "a"
			},{
				"type": "multi",
				"js": "b"
			}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Am, 2)
			assert.Equal(t, "a", v.(*data.Multi).Am[0].Js)
			assert.Equal(t, "b", v.(*data.Multi).Am[1].Js)
		}),
	)
}

func TestAalajs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalajs": [["a", "b"],["c", "d"]]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalajs, 2)
			require.Len(t, v.(*data.Multi).Aalajs[0], 2)
			require.Len(t, v.(*data.Multi).Aalajs[1], 2)
			assert.Equal(t, "a", v.(*data.Multi).Aalajs[0][0])
			assert.Equal(t, "b", v.(*data.Multi).Aalajs[0][1])
			assert.Equal(t, "c", v.(*data.Multi).Aalajs[1][0])
			assert.Equal(t, "d", v.(*data.Multi).Aalajs[1][1])
		}),
	)
}

func TestAalas(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalas": [
				[{"type": "simple", "js": "a"}, {"type": "simple", "js": "b"}],
				[{"type": "simple", "js": "c"}, {"type": "simple", "js": "d"}]
			]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalas, 2)
			require.Len(t, v.(*data.Multi).Aalas[0], 2)
			require.Len(t, v.(*data.Multi).Aalas[1], 2)
			assert.Equal(t, "a", v.(*data.Multi).Aalas[0][0].Js)
			assert.Equal(t, "b", v.(*data.Multi).Aalas[0][1].Js)
			assert.Equal(t, "c", v.(*data.Multi).Aalas[1][0].Js)
			assert.Equal(t, "d", v.(*data.Multi).Aalas[1][1].Js)
		}),
	)
}

func TestAalass(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalass": [["a", "b"],["c", "d"]]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalass, 2)
			require.Len(t, v.(*data.Multi).Aalass[0], 2)
			require.Len(t, v.(*data.Multi).Aalass[1], 2)
			assert.Equal(t, "a", v.(*data.Multi).Aalass[0][0].Value())
			assert.Equal(t, "b", v.(*data.Multi).Aalass[0][1].Value())
			assert.Equal(t, "c", v.(*data.Multi).Aalass[1][0].Value())
			assert.Equal(t, "d", v.(*data.Multi).Aalass[1][1].Value())
		}),
	)
}

func TestAaljb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aaljb": [true, false]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aaljb, 2)
			assert.Equal(t, true, v.(*data.Multi).Aaljb[0].Value())
			assert.Equal(t, false, v.(*data.Multi).Aaljb[1].Value())
		}),
	)
}

func TestAaljn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aaljn": [1.1, 1.2]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aaljn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Aaljn[0].Value())
			assert.Equal(t, 1.2, v.(*data.Multi).Aaljn[1].Value())
		}),
	)
}

func TestAaljs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aaljs": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aaljs, 2)
			assert.Equal(t, "a", v.(*data.Multi).Aaljs[0].Value())
			assert.Equal(t, "b", v.(*data.Multi).Aaljs[1].Value())
		}),
	)
}

func TestAalmjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalmjs": [{"a": "b", "c": "d"}, {"e": "f", "g": "h"}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalmjs, 2)
			require.Len(t, v.(*data.Multi).Aalmjs[0], 2)
			require.Len(t, v.(*data.Multi).Aalmjs[1], 2)
			assert.Equal(t, "b", v.(*data.Multi).Aalmjs[0]["a"])
			assert.Equal(t, "d", v.(*data.Multi).Aalmjs[0]["c"])
			assert.Equal(t, "f", v.(*data.Multi).Aalmjs[1]["e"])
			assert.Equal(t, "h", v.(*data.Multi).Aalmjs[1]["g"])
		}),
	)
}

func TestAalms(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalms": [{
				"a": {"type":"simple", "js": "b"},
				"c": {"type":"simple", "js": "d"}
			}, {
				"e": {"type":"simple", "js": "f"},
				"g": {"type":"simple", "js": "h"}
			}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalms, 2)
			require.Len(t, v.(*data.Multi).Aalms[0], 2)
			require.Len(t, v.(*data.Multi).Aalms[1], 2)
			assert.Equal(t, "b", v.(*data.Multi).Aalms[0]["a"].Js)
			assert.Equal(t, "d", v.(*data.Multi).Aalms[0]["c"].Js)
			assert.Equal(t, "f", v.(*data.Multi).Aalms[1]["e"].Js)
			assert.Equal(t, "h", v.(*data.Multi).Aalms[1]["g"].Js)
		}),
	)
}

func TestAalmss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalmss": [{"a": "b", "c": "d"}, {"e": "f", "g": "h"}]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalmss, 2)
			require.Len(t, v.(*data.Multi).Aalmss[0], 2)
			require.Len(t, v.(*data.Multi).Aalmss[1], 2)
			assert.Equal(t, "b", v.(*data.Multi).Aalmss[0]["a"].Value())
			assert.Equal(t, "d", v.(*data.Multi).Aalmss[0]["c"].Value())
			assert.Equal(t, "f", v.(*data.Multi).Aalmss[1]["e"].Value())
			assert.Equal(t, "h", v.(*data.Multi).Aalmss[1]["g"].Value())
		}),
	)
}

func TestAalss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aalss": ["a", "b"]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aalss, 2)
			assert.Equal(t, "a", v.(*data.Multi).Aalss[0].Value())
			assert.Equal(t, "b", v.(*data.Multi).Aalss[1].Value())
		}),
	)
}

func TestAals(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aals": [
				{"type": "simple", "js": "a"},
				{"type": "simple", "js": "b"}
			]
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Aals, 2)
			assert.Equal(t, "a", v.(*data.Multi).Aals[0].Js)
			assert.Equal(t, "b", v.(*data.Multi).Aals[1].Js)
		}),
	)
}
