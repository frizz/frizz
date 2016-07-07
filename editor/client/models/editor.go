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

func GetEditable(ctx context.Context, node *node.Node) editable.Editable {

	if node != nil {
		// This is the recommended method of presenting an custom editor.
		if ed, ok := node.Value.(editable.Editable); ok {
			return ed
		}
	}

	editors := clientctx.FromContext(ctx)

	if node == nil {
		e, _ := editors.Get("")
		return e
	}

	// Don't do this. Implement the Editable interface instead. We can't do this
	// for system types so we use this method instead.
	if e, ok := editors.Get(node.Type.Id.String()); ok {
		return e
	}

	if node.JsonType != "" {
		if e, ok := editors.Get(string(node.JsonType)); ok {
			return e
		}
	}

	// DefaultEditor
	e, _ := editors.Get("")
	return e
}
