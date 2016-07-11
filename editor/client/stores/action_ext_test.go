package stores

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"

	"kego.io/system/node"
	"kego.io/tests/data"
)

func TestAddMutation(t *testing.T) {

	cb, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var a, p *node.Node

		a = node.NewNode()
		p = n.Map["ajs"]
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil))
		assert.Equal(t, 5, len(n.Map["ajs"].Array))
		assert.Equal(t, 5, len(m.Ajs))
		require.NoError(t, mutateUndoAddNode(cb.Ctx(), p, "", 0))
		assert.Equal(t, 4, len(n.Map["ajs"].Array))
		assert.Equal(t, 4, len(m.Ajs))

		a = node.NewNode()
		p = n.Map["ass"]
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil))
		assert.Equal(t, 5, len(n.Map["ass"].Array))
		assert.Equal(t, 5, len(m.Ass))
		require.NoError(t, mutateUndoAddNode(cb.Ctx(), p, "", 0))
		assert.Equal(t, 4, len(n.Map["ass"].Array))
		assert.Equal(t, 4, len(m.Ass))

		a = node.NewNode()
		p = n.Map["am"]
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil))
		assert.Equal(t, 3, len(n.Map["am"].Array))
		assert.Equal(t, 3, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 3, len(m.Am))
		require.NoError(t, mutateUndoAddNode(cb.Ctx(), p, "", 0))
		assert.Equal(t, 2, len(n.Map["am"].Array))
		assert.Equal(t, 2, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 2, len(m.Am))
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}

func TestAddMutation2(t *testing.T) {
	cb, n := data.Setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var a, p *node.Node
		a = node.NewNode()
		p = n.Map["am"]
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil))
		assert.Equal(t, 3, len(n.Map["am"].Array))
		assert.Equal(t, 3, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 3, len(m.Am))
		assert.Equal(t, "", m.Am[0].Js)
		assert.Equal(t, "amjs0", m.Am[1].Js)
		assert.Equal(t, "amjs1", m.Am[2].Js)
		require.NoError(t, mutateUndoAddNode(cb.Ctx(), p, "", 0))
		assert.Equal(t, 2, len(n.Map["am"].Array))
		assert.Equal(t, 2, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 2, len(m.Am))
		assert.Equal(t, "amjs0", m.Am[0].Js)
		assert.Equal(t, "amjs1", m.Am[1].Js)
	}
	test(t, n.Map["m"], n.Value.(*data.Multi).M)
}

func TestDeleteMutation(t *testing.T) {

	cb, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var b, d, p *node.Node

		b = &node.Node{}
		d = n.Map["jb"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["jb"].ValueBool)
		assert.False(t, m.Jb)
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["jb"].ValueBool)
		assert.True(t, m.Jb)

		b = &node.Node{}
		d = n.Map["ss"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["ss"].Missing)
		assert.Equal(t, "", n.Map["ss"].ValueString)
		assert.Nil(t, m.Ss)
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["ss"].Missing)
		assert.Equal(t, "ss1", n.Map["ss"].ValueString)
		require.NotNil(t, m.Ss)
		assert.Equal(t, "ss1", m.Ss.Value())

		b = &node.Node{}
		d = n.Map["i"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["i"].Missing)
		assert.Nil(t, m.I)
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["i"].Missing)
		assert.Equal(t, "ia", n.Map["i"].Value.(*data.Facea).A.Value())
		require.NotNil(t, m.I)
		assert.Equal(t, "ia", m.I.Face())

		b = &node.Node{}
		d = n.Map["ass"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["ass"].Missing)
		assert.Equal(t, 0, len(m.Ass))
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["ass"].Missing)
		assert.Equal(t, 4, len(m.Ass))

		b = &node.Node{}
		d = n.Map["mss"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["mss"].Missing)
		assert.Nil(t, m.Mss)
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["mss"].Missing)
		require.NotNil(t, m.Mss)
		assert.Equal(t, 2, len(m.Mss))
		require.NotNil(t, m.Mss["a"])
		assert.Equal(t, "mssa", m.Mss["a"].Value())

		b = &node.Node{}
		d = n.Map["ass"].Array[0]
		p = n.Map["ass"]
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 3, len(n.Map["ass"].Array))
		assert.Equal(t, 3, len(m.Ass))
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 4, len(n.Map["ass"].Array))
		assert.Equal(t, 4, len(m.Ass))
		assert.Equal(t, "ass0", n.Map["ass"].Array[0].ValueString)
		require.NotNil(t, m.Ass[0])
		assert.Equal(t, "ass0", m.Ass[0].Value())

	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}

func TestDeleteMutation1(t *testing.T) {
	cb, n := data.Setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var b, d, p *node.Node
		b = &node.Node{}
		d = n.Map["am"].Array[0]
		p = n.Map["am"]
		assert.Equal(t, "amjs0", n.Map["am"].Array[0].Map["js"].ValueString)
		assert.Equal(t, "amjs1", n.Map["am"].Array[1].Map["js"].ValueString)
		assert.Equal(t, "amjs0", m.Am[0].Js)
		assert.Equal(t, "amjs1", m.Am[1].Js)
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 1, len(n.Map["am"].Array))
		assert.Equal(t, 1, len(m.Am))
		assert.Equal(t, "amjs1", n.Map["am"].Array[0].Map["js"].ValueString)
		assert.Equal(t, "amjs1", m.Am[0].Js)
		require.NoError(t, mutateUndoDeleteNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 2, len(n.Map["am"].Array))
		assert.Equal(t, 2, len(m.Am))
		assert.Equal(t, "amjs0", n.Map["am"].Array[0].Map["js"].ValueString)
		assert.Equal(t, "amjs1", n.Map["am"].Array[1].Map["js"].ValueString)
		assert.Equal(t, "amjs0", m.Am[0].Js)
		assert.Equal(t, "amjs1", m.Am[1].Js)
	}
	test(t, n.Map["m"], n.Value.(*data.Multi).M)
}
