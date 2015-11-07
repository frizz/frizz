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

	return Default(n)
}

func Default(n *Node) Editor {
	switch n.JsonType {
	case json.J_STRING:
		return &NodeStringEditor{Node: n, Common: &Common{}}
	case json.J_BOOL:
		return &NodeBoolEditor{Node: n, Common: &Common{}}
	case json.J_NUMBER:
		return &NodeNumberEditor{Node: n, Common: &Common{}}
	case json.J_MAP:
		return &NodeMapEditor{Node: n, Common: &Common{}}
	case json.J_OBJECT:
		return &NodeObjectEditor{Node: n, Common: &Common{}}
	case json.J_ARRAY:
		return &NodeArrayEditor{Node: n, Common: &Common{}}
	}
	return nil
}
