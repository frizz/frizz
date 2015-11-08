package editor

import (
	"kego.io/json"
	"kego.io/system/node"
)

type Node struct {
	*node.Node
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
}
