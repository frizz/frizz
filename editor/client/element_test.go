package client

import (
	"reflect"
	"testing"

	"kego.io/editor/tree"
	"kego.io/kerr/assert"
	"kego.io/system"
	_ "kego.io/system/types"
)

func TestTree(t *testing.T) {
	root := tree.New(nil)
	h, ok := system.GetGlobal("kego.io/system", "type")
	assert.True(t, ok)
	ty, ok := h.Object.(*system.Type)
	assert.True(t, ok)
	typ, ok := ty.Type.GetType()
	rule := system.NewMinimalRuleHolder(typ)
	child := &element{data: ty, rule: rule, value: reflect.ValueOf(ty), index: -1, name: "type"}
	err := addElement(child, root)
	assert.NoError(t, err)
	err = root.Each(func(n *tree.Node) error {
		return nil
	})
}
