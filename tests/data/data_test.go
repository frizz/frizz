package data_test

import (
	"testing"

	"github.com/davelondon/ktest/require"
	"kego.io/system/node"
	"kego.io/tests/data"
)

func TestData(t *testing.T) {
	_, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		require.Equal(t, "almsa", m.Alms["a"].Js)
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}