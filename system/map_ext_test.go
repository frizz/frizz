package system_test

import (
	"context"
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/ke"
	"kego.io/tests/data"
	. "kego.io/tests/marshal"
)

func TestMapAlmjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapAlms(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapAlmss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMjs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMjn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMjb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMss(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMsn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMsb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMsr(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

func TestMapMsp(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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

/*
func TestMapMri(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Test(t, ctx, `{
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
*/

/*
   "mnri": {
     "type": "system:@map",
     "items": {
       "type": "system:@string",
       "interface": true
     }
   },
   "mi": {
     "type": "system:@map",
     "items": {
       "type": "@face"
     }
   },
   "mm": {
     "type": "system:@map",
     "items": {
       "type": "@multi"
     }
   },
   "malajs": {
     "type": "system:@map",
     "items": {
       "type": "@alajs"
     }
   },
   "malas": {
     "type": "system:@map",
     "items": {
       "type": "@alas"
     }
   },
   "malass": {
     "type": "system:@map",
     "items": {
       "type": "@alass"
     }
   },
   "maljb": {
     "type": "system:@map",
     "items": {
       "type": "@aljb"
     }
   },
   "maljn": {
     "type": "system:@map",
     "items": {
       "type": "@aljn"
     }
   },
   "maljs": {
     "type": "system:@map",
     "items": {
       "type": "@aljs"
     }
   },
   "malmjs": {
     "type": "system:@map",
     "items": {
       "type": "@almjs"
     }
   },
   "malms": {
     "type": "system:@map",
     "items": {
       "type": "@alms"
     }
   },
   "malmss": {
     "type": "system:@map",
     "items": {
       "type": "@almss"
     }
   },
   "malss": {
     "type": "system:@map",
     "items": {
       "type": "@alss"
     }
   },
   "mals": {
     "type": "system:@map",
     "items": {
       "type": "@als"
     }
   }
*/
