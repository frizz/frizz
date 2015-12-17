package editor

import (
	"github.com/tjgq/broadcast"
	"kego.io/json"
	"kego.io/system/node"
)

type Node struct {
	*node.Node
	editor  EditorInterface
	Parent  *Node
	Array   []*Node
	Map     map[string]*Node
	changes *broadcast.Broadcaster
}

func (n *Node) UpdateFromInnerNode() {
	n.Array = []*Node{}
	for _, child := range n.Node.Array {
		n.Array = append(n.Array, NewNode(child, n))
	}
	n.Map = map[string]*Node{}
	for name, child := range n.Node.Map {
		n.Map[name] = NewNode(child, n)
	}
}

func NewNode(inner *node.Node, parent *Node) *Node {
	n := &Node{Node: inner}
	n.Parent = parent
	n.UpdateFromInnerNode()
	n.changes = broadcast.New(0)
	return n
}

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
