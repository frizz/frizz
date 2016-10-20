package stores

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"

	"kego.io/system"
	"kego.io/system/node"
	"kego.io/tests/data"
)

func TestAddMutationRoot(t *testing.T) {
	cb, _ := data.Setup(t)
	a := node.NewNode()
	ty, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/tests/data", "multi")
	require.True(t, ok)
	require.NoError(t, mutateAddNode(cb.Ctx(), a, nil, "", 2, ty, "z"))
	require.NotNil(t, a.Value.(*data.Multi).Id)
	require.Equal(t, "z", a.Value.(*data.Multi).Id.Name)
}

func TestAddMutationRedo(t *testing.T) {
	cb, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {

		var a, p, b, a1, p1, b1 *node.Node
		a = node.NewNode()
		p = n.Map["am"]
		b = node.NewNode()
		ty, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/tests/data", "multi")
		require.True(t, ok)
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 2, ty, ""))
		require.Equal(t, `[{"js":"amjs0","type":"multi"},{"js":"amjs1","type":"multi"},{"type":"multi"}]`, p.Print(cb.Ctx()))

		a1 = n.Map["am"].Array[2].Map["m"]
		p1 = n.Map["am"].Array[2]
		b1 = node.NewNode()
		require.NoError(t, mutateAddNode(cb.Ctx(), a1, p1, "m", -1, ty, ""))
		require.Equal(t, 3, len(n.Map["am"].Array))
		require.Equal(t, 3, len(m.Am))
		require.False(t, n.Map["am"].Array[2].Map["m"].Missing)
		require.Equal(t, `[{"js":"amjs0","type":"multi"},{"js":"amjs1","type":"multi"},{"m":{"type":"multi"},"type":"multi"}]`, p.Print(cb.Ctx()))

		require.NoError(t, mutateDeleteNode(cb.Ctx(), a, p, b))
		require.NoError(t, mutateDeleteNode(cb.Ctx(), a1, p1, b1))
		require.Equal(t, 2, len(n.Map["am"].Array))
		require.Equal(t, 2, len(m.Am))
		require.NoError(t, mutateRestoreNode(cb.Ctx(), a, p, b))
		require.NoError(t, mutateRestoreNode(cb.Ctx(), a1, p1, b1))
		require.Equal(t, 3, len(n.Map["am"].Array))
		require.Equal(t, 3, len(m.Am))
		require.NotNil(t, n.Map["am"].Array[2])
		require.NotNil(t, m.Am[2])
	}

	test(t, n.Map["m"], n.Value.(*data.Multi).M)
}

func TestAddMutation(t *testing.T) {

	cb, n := data.Setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var a, p, b *node.Node

		a = node.NewNode()
		p = n.Map["ajs"]
		b = node.NewNode()
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil, ""))
		assert.Equal(t, 5, len(n.Map["ajs"].Array))
		assert.Equal(t, 5, len(m.Ajs))

		require.NoError(t, mutateDeleteNode(cb.Ctx(), a, p, b))
		assert.Equal(t, 4, len(n.Map["ajs"].Array))
		assert.Equal(t, 4, len(m.Ajs))

		a = node.NewNode()
		p = n.Map["ass"]
		b = node.NewNode()
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil, ""))
		assert.Equal(t, 5, len(n.Map["ass"].Array))
		assert.Equal(t, 5, len(m.Ass))
		require.NoError(t, mutateDeleteNode(cb.Ctx(), a, p, b))
		assert.Equal(t, 4, len(n.Map["ass"].Array))
		assert.Equal(t, 4, len(m.Ass))

		a = node.NewNode()
		p = n.Map["am"]
		b = node.NewNode()
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil, ""))
		assert.Equal(t, 3, len(n.Map["am"].Array))
		assert.Equal(t, 3, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 3, len(m.Am))
		require.NoError(t, mutateDeleteNode(cb.Ctx(), a, p, b))
		assert.Equal(t, 2, len(n.Map["am"].Array))
		assert.Equal(t, 2, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 2, len(m.Am))
	}

	data.Run(t, n, n.Value.(*data.Multi), test)
}

func TestAddMutation2(t *testing.T) {
	cb, n := data.Setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		var a, p, b *node.Node
		a = node.NewNode()
		p = n.Map["am"]
		b = node.NewNode()
		require.NoError(t, mutateAddNode(cb.Ctx(), a, p, "", 0, nil, ""))
		assert.Equal(t, 3, len(n.Map["am"].Array))
		assert.Equal(t, 3, len(n.Map["am"].Value.([]*data.Multi)))
		assert.Equal(t, 3, len(m.Am))
		assert.Equal(t, "", m.Am[0].Js)
		assert.Equal(t, "amjs0", m.Am[1].Js)
		assert.Equal(t, "amjs1", m.Am[2].Js)
		require.NoError(t, mutateDeleteNode(cb.Ctx(), a, p, b))
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

		b = node.NewNode()
		d = n.Map["jb"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["jb"].ValueBool)
		assert.False(t, m.Jb)
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["jb"].ValueBool)
		assert.True(t, m.Jb)

		b = node.NewNode()
		d = n.Map["ss"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["ss"].Missing)
		assert.Equal(t, "", n.Map["ss"].ValueString)
		assert.Nil(t, m.Ss)
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["ss"].Missing)
		assert.Equal(t, "ss1", n.Map["ss"].ValueString)
		require.NotNil(t, m.Ss)
		assert.Equal(t, "ss1", m.Ss.Value())

		b = node.NewNode()
		d = n.Map["i"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["i"].Missing)
		assert.Nil(t, m.I)
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["i"].Missing)
		assert.Equal(t, "ia", n.Map["i"].Value.(*data.Facea).A.Value())
		require.NotNil(t, m.I)
		assert.Equal(t, "ia", m.I.Face())

		b = node.NewNode()
		d = n.Map["ass"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["ass"].Missing)
		assert.Equal(t, 0, len(m.Ass))
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["ass"].Missing)
		assert.Equal(t, 4, len(m.Ass))

		b = node.NewNode()
		d = n.Map["mss"]
		p = n
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.True(t, n.Map["mss"].Missing)
		assert.Nil(t, m.Mss)
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.False(t, n.Map["mss"].Missing)
		require.NotNil(t, m.Mss)
		assert.Equal(t, 2, len(m.Mss))
		require.NotNil(t, m.Mss["a"])
		assert.Equal(t, "mssa", m.Mss["a"].Value())

		b = node.NewNode()
		d = n.Map["ass"].Array[0]
		p = n.Map["ass"]
		require.NoError(t, mutateDeleteNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 3, len(n.Map["ass"].Array))
		assert.Equal(t, 3, len(m.Ass))
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
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
		b = node.NewNode()
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
		require.NoError(t, mutateRestoreNode(cb.Ctx(), d, p, b))
		assert.Equal(t, 2, len(n.Map["am"].Array))
		assert.Equal(t, 2, len(m.Am))
		assert.Equal(t, "amjs0", n.Map["am"].Array[0].Map["js"].ValueString)
		assert.Equal(t, "amjs1", n.Map["am"].Array[1].Map["js"].ValueString)
		assert.Equal(t, "amjs0", m.Am[0].Js)
		assert.Equal(t, "amjs1", m.Am[1].Js)
	}
	test(t, n.Map["m"], n.Value.(*data.Multi).M)
}
