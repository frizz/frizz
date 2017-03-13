package data_test

import (
	"testing"

	"context"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"kego.io/ke"
	"kego.io/system/node"
	"kego.io/tests/data"
	. "kego.io/tests/marshal"
)

func TestUnpackInterface(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)

	Run(t, ctx, `{
			"type": "aljb",
			"value": true
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Equal(t, true, bool(*v.(*data.Aljb)))
		}),
	)

	Run(t, ctx, `meh`,
		UnmarshalError(func(t *testing.T, err error) {
			require.IsError(t, err, "SVXYHJWMOC")
			require.HasError(t, err, "PDTPGAYXRX")
		}),
	)

	Run(t, ctx, `{
			"type": "multi",
			"sri": "a"
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Equal(t, "a", v.(*data.Multi).Sri.GetString(ctx).String())
		}),
	)

	Run(t, ctx, `{
			"type": "multi",
			"sri": {"type": "system:string", "value": "a"}
		}`,
		TestValue(func(t *testing.T, v interface{}) {
			assert.Equal(t, "a", v.(*data.Multi).Sri.GetString(ctx).String())
		}),
		MarshalledString(`{
			"type": "multi",
			"sri": "a"
		}`),
	)
}

func TestData(t *testing.T) {
	_, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		require.Equal(t, "almsa", m.Alms["a"].Js)
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}
