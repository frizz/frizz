package editor

import (
	"kego.io/json"
	"kego.io/system/node"
)

type Node struct {
	*node.Node
	editor Editor
	Parent *Node
	Array  []*Node
	Map    map[string]*Node
}

func NewNode(n *node.Node, parent *Node) *Node {
	n1 := &Node{Node: n}
	n1.Parent = parent
	for _, child := range n.Array {
		n1.Array = append(n1.Array, NewNode(child, n1))
	}
	n1.Map = map[string]*Node{}
	for name, child := range n.Map {
		n1.Map[name] = NewNode(child, n1)
	}
	return n1
}

func (n *Node) Editor() Editor {

	if ed, ok := n.Value.(Editable); ok {
		return ed.GetEditor(n)
	}

	if factory := Get(*n.Type.Id); factory != nil {
		return factory(n)
	}

	switch n.JsonType {
	case json.J_STRING:
		return NewStringEditor(n)
	case json.J_BOOL:
		return NewBoolEditor(n)
	case json.J_NUMBER:
		return NewNumberEditor(n)
	case json.J_MAP:
		return NewMapEditor(n)
	case json.J_OBJECT:
		return NewObjectEditor(n)
	case json.J_ARRAY:
		return NewArrayEditor(n)
	}
	return nil

	/*
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
	*/
}
