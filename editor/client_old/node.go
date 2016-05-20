package editor

import (
	"github.com/davelondon/kerr"
	"github.com/tjgq/broadcast"
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/system/node"
)

type Node struct {
	*node.Node
	editor  EditorInterface
	changes *broadcast.Broadcaster
}

func NewEditorNode() *Node {
	n := &Node{Node: &node.Node{}}
	n.Self = n
	n.changes = broadcast.New(0)
	return n
}

func UnmarshalNode(ctx context.Context, data []byte) (*Node, error) {
	n := NewEditorNode()
	if err := ke.UnmarshalUntyped(ctx, data, n); err != nil {
		return nil, kerr.Wrap("SIWBWHJDJD", err)
	}
	return n, nil
}

func (n *Node) NewChild() node.NodeInterface {
	return NewEditorNode()
}

var _ node.NodeInterface = (*Node)(nil)

func (n *Node) Editor() EditorInterface {

	if n.Null {
		return nil
	}

	if n.editor != nil {
		return n.editor
	}

	if ed, ok := n.Value.(Editable); ok {
		n.editor = ed.GetEditor(n)
		return n.editor
	}

	if factory := Get(*n.Type.Id); factory != nil {
		n.editor = factory(n)
		return n.editor
	}

	switch n.JsonType {
	case json.J_STRING:
		n.editor = NewStringEditor(n)
	case json.J_BOOL:
		n.editor = NewBoolEditor(n)
	case json.J_NUMBER:
		n.editor = NewNumberEditor(n)
	case json.J_MAP:
		n.editor = NewMapEditor(n)
	case json.J_OBJECT:
		n.editor = NewObjectEditor(n)
	case json.J_ARRAY:
		n.editor = NewArrayEditor(n)
	}
	return n.editor

}
