package stores_test

import (
	"testing"

	"github.com/davelondon/ktest/assert"

	"kego.io/editor/client/stores"
	"kego.io/system/node"
	"kego.io/tests/data"
)

func TestDeleteMutation(t *testing.T) {

	cb, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {

		mu := stores.Delete{
			Parent: n,
			Node:   n.Map["jb"],
		}
		assert.NoError(t, mu.NodeDo(cb.Ctx()))
		assert.False(t, n.Map["jb"].ValueBool)
		assert.False(t, m.Jb)
		assert.NoError(t, mu.NodeUndo(cb.Ctx()))
		assert.True(t, n.Map["jb"].ValueBool)
		assert.True(t, m.Jb)

		mu = stores.Delete{
			Parent: n,
			Node:   n.Map["ss"],
		}
		assert.NoError(t, mu.NodeDo(cb.Ctx()))
		assert.True(t, n.Map["ss"].Missing)
		assert.Equal(t, "", n.Map["ss"].ValueString)
		assert.Nil(t, m.Ss)
		assert.NoError(t, mu.NodeUndo(cb.Ctx()))
		assert.False(t, n.Map["ss"].Missing)
		assert.Equal(t, "ss1", n.Map["ss"].ValueString)
		assert.NotNil(t, m.Ss)
		assert.Equal(t, "ss1", m.Ss.Value())

		mu = stores.Delete{
			Parent: n,
			Node:   n.Map["i"],
		}
		assert.NoError(t, mu.NodeDo(cb.Ctx()))
		assert.True(t, n.Map["i"].Missing)
		assert.Nil(t, m.I)
		assert.NoError(t, mu.NodeUndo(cb.Ctx()))
		assert.False(t, n.Map["i"].Missing)
		assert.Equal(t, "ia", n.Map["i"].Value.(*data.Facea).A.Value())
		assert.NotNil(t, m.I)
		assert.Equal(t, "ia", m.I.Face())

		mu = stores.Delete{
			Parent: n,
			Node:   n.Map["ass"],
		}
		assert.NoError(t, mu.NodeDo(cb.Ctx()))
		assert.True(t, n.Map["ass"].Missing)
		assert.Equal(t, 0, len(m.Ass))
		assert.NoError(t, mu.NodeUndo(cb.Ctx()))
		assert.False(t, n.Map["ass"].Missing)
		assert.Equal(t, 4, len(m.Ass))

		mu = stores.Delete{
			Parent: n,
			Node:   n.Map["mss"],
		}
		assert.NoError(t, mu.NodeDo(cb.Ctx()))
		assert.True(t, n.Map["mss"].Missing)
		assert.Nil(t, m.Mss)
		assert.NoError(t, mu.NodeUndo(cb.Ctx()))
		assert.False(t, n.Map["mss"].Missing)
		assert.NotNil(t, m.Mss)
		assert.Equal(t, 2, len(m.Mss))
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}
