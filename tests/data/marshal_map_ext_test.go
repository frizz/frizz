package data_test

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

func TestAlmjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"almjs": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Almjs, 2)
			assert.Equal(t, "b", v.(*data.Multi).Almjs["a"])
			assert.Equal(t, "d", v.(*data.Multi).Almjs["c"])
		}),
	)
}

func TestAlms(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"alms": {
				"a": {"type": "simple", "js": "b"},
				"c": {"type": "simple", "js": "d"}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Alms, 2)
			assert.Equal(t, "b", v.(*data.Multi).Alms["a"].Js)
			assert.Equal(t, "d", v.(*data.Multi).Alms["c"].Js)
		}),
	)
}

func TestAlmss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"almss": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Almss, 2)
			assert.Equal(t, "b", v.(*data.Multi).Almss["a"].Value())
			assert.Equal(t, "d", v.(*data.Multi).Almss["c"].Value())
		}),
	)
}

func TestMjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mjs": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mjs, 2)
			assert.Equal(t, "b", v.(*data.Multi).Mjs["a"])
			assert.Equal(t, "d", v.(*data.Multi).Mjs["c"])
		}),
	)
}

func TestMjn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mjn": {"a": 1.1, "b": 1.2}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mjn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Mjn["a"])
			assert.Equal(t, 1.2, v.(*data.Multi).Mjn["b"])
		}),
	)
}

func TestMjb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mjb": {"a": true, "b": false}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mjb, 2)
			assert.Equal(t, true, v.(*data.Multi).Mjb["a"])
			assert.Equal(t, false, v.(*data.Multi).Mjb["b"])
		}),
	)
}

func TestMss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mss": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mss, 2)
			assert.Equal(t, "b", v.(*data.Multi).Mss["a"].Value())
			assert.Equal(t, "d", v.(*data.Multi).Mss["c"].Value())
		}),
	)
}

func TestMsn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"msn": {"a": 1.1, "b": 1.2}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Msn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Msn["a"].Value())
			assert.Equal(t, 1.2, v.(*data.Multi).Msn["b"].Value())
		}),
	)
}

func TestMsb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"msb": {"a": true, "b": false}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Msb, 2)
			assert.Equal(t, true, v.(*data.Multi).Msb["a"].Value())
			assert.Equal(t, false, v.(*data.Multi).Msb["b"].Value())
		}),
	)
}

func TestMsr(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"msr": {"a": "a", "b": "system:b"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Msr, 2)
			assert.Equal(t, "kego.io/tests/data:a", v.(*data.Multi).Msr["a"].Value())
			assert.Equal(t, "kego.io/system:b", v.(*data.Multi).Msr["b"].Value())
		}),
	)
}

func TestMsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"msi": {"a": 2, "b": 3}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Msi, 2)
			assert.Equal(t, 2, v.(*data.Multi).Msi["a"].Value())
			assert.Equal(t, 3, v.(*data.Multi).Msi["b"].Value())
		}),
	)
}

func TestMsp(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"msp": {
				"a": {"type": "system:package", "recursive": true},
				"b": {"type": "system:package"}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Msp, 2)
			assert.Equal(t, true, v.(*data.Multi).Msp["a"].Recursive)
			assert.Equal(t, false, v.(*data.Multi).Msp["b"].Recursive)
		}),
	)
}

func TestAlmjsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Note: map alias types in interfaces are always marshalled to explicit
	// object notation.
	Run(t, ctx, `{
			"type": "multi",
			"almjsi": {"type": "almjs", "value": {"a": "b", "c": "d"}}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Len(t, v.(*data.Multi).Almjsi.GetAlmjs(ctx), 2)
			assert.Equal(t, "b", v.(*data.Multi).Almjsi.GetAlmjs(ctx)["a"])
			assert.Equal(t, "d", v.(*data.Multi).Almjsi.GetAlmjs(ctx)["c"])
		}),
	)
}

func TestAlmjsiError(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Note: map alias types in interfaces are always marshalled to explicit
	// object notation, so this throws an error.
	Run(t, ctx, `{
			"type": "multi",
			"almjsi": {"a": "b", "c": "d"}
		}`,
		UnmarshalError(func(t *testing.T, err error) {
			require.IsError(t, err, "SVXYHJWMOC")
			require.HasError(t, err, "RXEPCCGFKV")
		}),
	)
}

func TestMri(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mri": {
				"a": {
					"type": "system:@string",
					"max-length": 2
				},
				"b": {
					"type": "system:@bool",
					"default": true
				}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mri, 2)
			assert.Equal(t, 2, v.(*data.Multi).Mri["a"].(*system.StringRule).MaxLength.Value())
			assert.Equal(t, true, v.(*data.Multi).Mri["b"].(*system.BoolRule).Default.Value())
		}),
	)
}

func TestMnri(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mnri": {
				"a": "a",
				"b": {"type": "system:int", "value": 1},
				"c": {"type": "alajs", "value": ["a", "b", "c"]},
				"d": {"type": "almjs", "value": {"a": "b", "c": "d"}}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mnri, 4)
			assert.Equal(t, "a", v.(*data.Multi).Mnri["a"].GetString(ctx).Value())
			assert.Equal(t, "1", v.(*data.Multi).Mnri["b"].GetString(ctx).Value())
			assert.Equal(t, "abc", v.(*data.Multi).Mnri["c"].GetString(ctx).Value())
			assert.Equal(t, "abcd", v.(*data.Multi).Mnri["d"].GetString(ctx).Value())
		}),
	)
}

func TestMi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mi": {
				"a": {
					"type": "facea",
					"a": "a1"
				},
				"b": {
					"type": "faceb",
					"b": "b1"
				}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mi, 2)
			assert.Equal(t, "a1", v.(*data.Multi).Mi["a"].Face())
			assert.Equal(t, "b1", v.(*data.Multi).Mi["b"].Face())
		}),
	)
}

func TestMm(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mm": {
				"a": {
					"type": "multi",
					"js": "b"
				},
				"c": {
					"type": "multi",
					"js": "d"
				}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mm, 2)
			assert.Equal(t, "b", v.(*data.Multi).Mm["a"].Js)
			assert.Equal(t, "d", v.(*data.Multi).Mm["c"].Js)
		}),
	)
}

func TestMalajs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malajs": {"a": ["b", "c"], "d": ["e", "f"]}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malajs, 2)
			require.Len(t, v.(*data.Multi).Malajs["a"], 2)
			require.Len(t, v.(*data.Multi).Malajs["d"], 2)
			assert.Equal(t, "b", v.(*data.Multi).Malajs["a"][0])
			assert.Equal(t, "c", v.(*data.Multi).Malajs["a"][1])
			assert.Equal(t, "e", v.(*data.Multi).Malajs["d"][0])
			assert.Equal(t, "f", v.(*data.Multi).Malajs["d"][1])
		}),
	)
}

func TestMalas(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malas": {
				"a": [{"type": "simple", "js": "b"}, {"type": "simple", "js": "c"}],
				"d": [{"type": "simple", "js": "e"}, {"type": "simple", "js": "f"}]
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malas, 2)
			require.Len(t, v.(*data.Multi).Malas["a"], 2)
			require.Len(t, v.(*data.Multi).Malas["d"], 2)
			assert.Equal(t, "b", v.(*data.Multi).Malas["a"][0].Js)
			assert.Equal(t, "c", v.(*data.Multi).Malas["a"][1].Js)
			assert.Equal(t, "e", v.(*data.Multi).Malas["d"][0].Js)
			assert.Equal(t, "f", v.(*data.Multi).Malas["d"][1].Js)
		}),
	)
}

func TestMalass(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malass": {"a": ["b", "c"], "d": ["e", "f"]}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malass, 2)
			require.Len(t, v.(*data.Multi).Malass["a"], 2)
			require.Len(t, v.(*data.Multi).Malass["d"], 2)
			assert.Equal(t, "b", v.(*data.Multi).Malass["a"][0].Value())
			assert.Equal(t, "c", v.(*data.Multi).Malass["a"][1].Value())
			assert.Equal(t, "e", v.(*data.Multi).Malass["d"][0].Value())
			assert.Equal(t, "f", v.(*data.Multi).Malass["d"][1].Value())
		}),
	)
}

func TestMaljb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"maljb": {"a": true, "b": false}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Maljb, 2)
			assert.Equal(t, true, v.(*data.Multi).Maljb["a"].Value())
			assert.Equal(t, false, v.(*data.Multi).Maljb["b"].Value())
		}),
	)
}

func TestMaljn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"maljn": {"a": 1.1, "b": 1.2}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Maljn, 2)
			assert.Equal(t, 1.1, v.(*data.Multi).Maljn["a"].Value())
			assert.Equal(t, 1.2, v.(*data.Multi).Maljn["b"].Value())
		}),
	)
}

func TestMaljs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"maljs": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Maljs, 2)
			assert.Equal(t, "b", v.(*data.Multi).Maljs["a"].Value())
			assert.Equal(t, "d", v.(*data.Multi).Maljs["c"].Value())
		}),
	)
}

func TestMalmjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malmjs": {"a": {"b": "c", "d": "e"}, "f": {"g": "h", "i": "j"}}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malmjs, 2)
			require.Len(t, v.(*data.Multi).Malmjs["a"], 2)
			require.Len(t, v.(*data.Multi).Malmjs["f"], 2)
			assert.Equal(t, "c", v.(*data.Multi).Malmjs["a"]["b"])
			assert.Equal(t, "e", v.(*data.Multi).Malmjs["a"]["d"])
			assert.Equal(t, "h", v.(*data.Multi).Malmjs["f"]["g"])
			assert.Equal(t, "j", v.(*data.Multi).Malmjs["f"]["i"])
		}),
	)
}

func TestMalms(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malms": {
				"a": {
					"b": {"type":"simple", "js": "c"},
					"d": {"type":"simple", "js": "e"}
				},
				"f": {
					"g": {"type":"simple", "js": "h"},
					"i": {"type":"simple", "js": "j"}
				}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malms, 2)
			require.Len(t, v.(*data.Multi).Malms["a"], 2)
			require.Len(t, v.(*data.Multi).Malms["f"], 2)
			assert.Equal(t, "c", v.(*data.Multi).Malms["a"]["b"].Js)
			assert.Equal(t, "e", v.(*data.Multi).Malms["a"]["d"].Js)
			assert.Equal(t, "h", v.(*data.Multi).Malms["f"]["g"].Js)
			assert.Equal(t, "j", v.(*data.Multi).Malms["f"]["i"].Js)
		}),
	)
}

func TestMalmss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malmss": {"a": {"b": "c", "d": "e"}, "f": {"g": "h", "i": "j"}}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malmss, 2)
			require.Len(t, v.(*data.Multi).Malmss["a"], 2)
			require.Len(t, v.(*data.Multi).Malmss["f"], 2)
			assert.Equal(t, "c", v.(*data.Multi).Malmss["a"]["b"].Value())
			assert.Equal(t, "e", v.(*data.Multi).Malmss["a"]["d"].Value())
			assert.Equal(t, "h", v.(*data.Multi).Malmss["f"]["g"].Value())
			assert.Equal(t, "j", v.(*data.Multi).Malmss["f"]["i"].Value())
		}),
	)
}

func TestMalss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"malss": {"a": "b", "c": "d"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Malss, 2)
			assert.Equal(t, "b", v.(*data.Multi).Malss["a"].Value())
			assert.Equal(t, "d", v.(*data.Multi).Malss["c"].Value())
		}),
	)
}

func TestMals(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"mals": {
				"a": {"type": "als", "js": "b"},
				"c": {"type": "als", "js": "d"}
			}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			require.Len(t, v.(*data.Multi).Mals, 2)
			assert.Equal(t, "b", v.(*data.Multi).Mals["a"].Js)
			assert.Equal(t, "d", v.(*data.Multi).Mals["c"].Js)
		}),
	)
}
