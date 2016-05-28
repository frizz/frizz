package models

import "kego.io/system/node"

type EditorModel struct {
	Node *node.Node
}

func NewEditor(n *node.Node) *EditorModel {
	return &EditorModel{Node: n}
}

func AddEditorsRecursively(editors map[*node.Node]*EditorModel, n *node.Node) *EditorModel {
	e := NewEditor(n)
	editors[n] = e
	for _, c := range n.Map {
		AddEditorsRecursively(editors, c)
	}
	for _, c := range n.Array {
		AddEditorsRecursively(editors, c)
	}
	return e
}
