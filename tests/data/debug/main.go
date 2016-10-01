package main

import (
	"testing"

	"github.com/davelondon/ktest/require"
	"kego.io/system/node"
	"kego.io/tests/data"
)

func main() {
	t := &testing.T{}
	_, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		require.Equal(t, "aljs", *m.Aljs)
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}
