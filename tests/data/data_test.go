package data_test

import (
	"testing"

	"context"

	"github.com/davelondon/ktest/require"
	"github.com/stretchr/testify/assert"
	"kego.io/ke"
	"kego.io/system/node"
	"kego.io/tests/data"
)

func TestUnpackInterface(t *testing.T) {
	ctx := ke.NewContext(context.Background(), "kego.io/tests/data", nil)

	test(
		t,
		ctx,
		`{
			"type": "aljb",
			"value": true
		}`,
		func(t *testing.T, v interface{}) {
			assert.Equal(t, true, bool(*v.(*data.Aljb)))
		},
	)

	test(
		t,
		ctx,
		`{
			"type": "multi",
			"sri": "a"
		}`,
		func(t *testing.T, v interface{}) {
			assert.Equal(t, "a", v.(*data.Multi).Sri.GetString(ctx).String())
		},
	)
}

func test(t *testing.T, ctx context.Context, json string, tester func(t *testing.T, v interface{})) {
	var v interface{}
	err := ke.Unmarshal(ctx, []byte(json), &v)
	require.NoError(t, err)
	tester(t, v)
	b, err := ke.Marshal(ctx, v)
	require.NoError(t, err)
	assert.JSONEq(t, json, string(b))
}

func TestData(t *testing.T) {
	_, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		require.Equal(t, "almsa", m.Alms["a"].Js)
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}
