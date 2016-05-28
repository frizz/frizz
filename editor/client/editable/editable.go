package editable

import (
	"github.com/davelondon/vecty"
	"golang.org/x/net/context"
	"kego.io/system/node"
)

type Editable interface {
	GetEditorView(ctx context.Context, node *node.Node) vecty.Component
}
