package actions

import (
	"kego.io/editor/shared"
	"kego.io/system/node"
)

type NewMessage struct {
	Message string
}

type InitialState struct {
	Info *shared.Info
}

type ToggleNode struct {
	Node *node.Node
}

type SelectNode struct {
	Node *node.Node
}
