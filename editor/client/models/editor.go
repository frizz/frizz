package models

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/editable"
	"kego.io/system/node"
)

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

func GetEditable(ctx context.Context, node *node.Node) editable.Editable {

	var e editable.Editable
	var ok bool

	if node == nil {
		e, _ = clientctx.FromContext(ctx).Get("")
		return e
	}

	// This is the recommended method of presenting an custom editor.
	ed, ok := node.Value.(editable.Editable)
	if ok {
		return ed
	}

	// Don't do this. Implement the Editable interface instead. We can't do this for system types
	// so we use this method instead.
	editors := clientctx.FromContext(ctx)
	e, ok = editors.Get(node.Type.Id.String())
	if ok {
		return e
	}

	if node.JsonType != "" {
		e, ok = editors.Get(string(node.JsonType))
		if ok {
			return e
		}
	}

	// DefaultEditor
	e, _ = editors.Get("")
	return e
}
