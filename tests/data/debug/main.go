package main

import (
	"testing"

	"frizz.io/system/node"
	"frizz.io/tests/data"
	"github.com/dave/ktest/require"
)

func main() {
	t := &testing.T{}
	_, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		require.Equal(t, "aljs", *m.Aljs)
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}
