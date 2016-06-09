package models

import (
	"golang.org/x/net/context"
	"kego.io/system"
	"kego.io/system/node"
)

type AddPopModel struct {
	ctx     context.Context
	Visible bool
	Parent  *node.Node
	Node    *node.Node
	Types   []*system.Type
	Success bool
}
