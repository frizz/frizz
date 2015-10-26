package editor

import (
	"kego.io/json"
	"kego.io/system/node"
)

func Default(n *node.Node) Editor {
	switch n.JsonType {
	case json.J_STRING, json.J_BOOL, json.J_NUMBER:
		return &NodeValueEditor{Node: n, Common: &Common{}}
	case json.J_MAP:
		return &NodeMapEditor{Node: n, Common: &Common{}}
	case json.J_OBJECT:
		return &NodeObjectEditor{Node: n, Common: &Common{}}
	case json.J_ARRAY:
		return &NodeArrayEditor{Node: n, Common: &Common{}}
	}
	return nil
}
