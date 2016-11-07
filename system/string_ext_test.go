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

func TestAljb(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljb": true
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljb)
			assert.Equal(t, true, v.(*data.Multi).Aljb.Value())
		}),
	)
}

func TestAljbi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljbi": true
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljbi.GetAljb(ctx))
			assert.IsType(t, new(data.Aljb), v.(*data.Multi).Aljbi)
			assert.Equal(t, true, v.(*data.Multi).Aljbi.GetAljb(ctx).Value())
		}),
	)
}

func TestAljbi2(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// aljb2 is another alias bool type that implements the aljb default
	// interface.
	Run(t, ctx, `{
			"type": "multi",
			"aljbi": {"type": "aljb2", "value": true}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljbi.GetAljb(ctx))
			assert.IsType(t, new(data.Aljb2), v.(*data.Multi).Aljbi)
			assert.Equal(t, true, v.(*data.Multi).Aljbi.GetAljb(ctx).Value())
		}),
	)
}

func TestAljn(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljn": 1.1
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljn)
			assert.Equal(t, 1.1, v.(*data.Multi).Aljn.Value())
		}),
	)
}

func TestAljni(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljni": 1.1
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljni.GetAljn(ctx))
			assert.IsType(t, new(data.Aljn), v.(*data.Multi).Aljni)
			assert.Equal(t, 1.1, v.(*data.Multi).Aljni.GetAljn(ctx).Value())
		}),
	)
}

func TestAljni2(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Aljn2 is another alias number type that implements the Aljn default
	// interface.
	Run(t, ctx, `{
			"type": "multi",
			"aljni": {"type": "aljn2", "value": 1.1}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljni.GetAljn(ctx))
			assert.IsType(t, new(data.Aljn2), v.(*data.Multi).Aljni)
			assert.Equal(t, 1.1, v.(*data.Multi).Aljni.GetAljn(ctx).Value())
		}),
	)
}

func TestAljs(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljs": "a"
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljs)
			assert.Equal(t, "a", v.(*data.Multi).Aljs.Value())
		}),
	)
}

func TestAljsi(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"aljsi": "a"
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljsi.GetAljs(ctx))
			assert.IsType(t, new(data.Aljs), v.(*data.Multi).Aljsi)
			assert.Equal(t, "a", v.(*data.Multi).Aljsi.GetAljs(ctx).Value())
		}),
	)
}

func TestAljsi2(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Aljs2 is another alias string type that implements the Aljs default
	// interface.
	Run(t, ctx, `{
			"type": "multi",
			"aljsi": {"type": "aljs2", "value": "a"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Aljsi.GetAljs(ctx))
			assert.IsType(t, new(data.Aljs2), v.(*data.Multi).Aljsi)
			assert.Equal(t, "a", v.(*data.Multi).Aljsi.GetAljs(ctx).Value())
		}),
	)
}

func TestAls(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	Run(t, ctx, `{
			"type": "multi",
			"als": {"type": "als", "js": "a"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.NotNil(t, v.(*data.Multi).Als)
			assert.Equal(t, "a", v.(*data.Multi).Als.Js)
		}),
	)
}

func TestAlsError(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)
	// Type must be als when unpacking into als, even if it's an alias
	// of simple
	Run(t, ctx, `{
			"type": "multi",
			"als": {"type": "simple", "js": "a"}
		}`,
		UnmarshalError(func(t *testing.T, err error) {
			require.IsError(t, err, "SVXYHJWMOC")
			require.HasError(t, err, "BNHVLQPFKC")
		}),
	)
}

/*
	Als    *Als                              `json:"als"`
	Alss   *Alss                             `json:"alss"`
	Bri    system.BoolInterface              `json:"bri"`
	I      Face                              `json:"i"`
	Jb     bool                              `json:"jb"`
	Jn     float64                           `json:"jn"`
	Js     string                            `json:"js"`
	M      *Multi                            `json:"m"`
	Nri    system.NumberInterface            `json:"nri"`
	Ri     system.RuleInterface              `json:"ri"`
	Sb     *system.Bool                      `json:"sb"`
	Si     *system.Int                       `json:"si"`
	Sn     *system.Number                    `json:"sn"`
	Sp     *system.Package                   `json:"sp"`
	Sr     *system.Reference                 `json:"sr"`
	Sri    system.StringInterface            `json:"sri"`
	Ss     *system.String                    `json:"ss"`
*/
