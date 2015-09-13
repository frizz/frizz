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
	o, _, ok := system.GetType("kego.io/system", "type")
	assert.True(t, ok)
	ty, ok := o.Type.GetType()
	assert.True(t, ok)
	rule := system.NewMinimalRuleHolder(ty)
	child := &element{data: o, rule: rule, value: reflect.ValueOf(o), index: -1, name: "type"}
	err := addElement(child, root)
	assert.NoError(t, err)
	err = root.Each(func(n *tree.Node) error {
		return nil
	})
}
