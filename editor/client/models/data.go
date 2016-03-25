package models

import "kego.io/system/node"

type DataModel struct {
	Name string
	File string
	Node *node.Node
}
