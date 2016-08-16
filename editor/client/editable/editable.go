package editable

import (
	"context"

	"github.com/davelondon/vecty"
	"kego.io/system"
	"kego.io/system/node"
)

type Format string

const (
	Branch Format = "branch"
	Block  Format = "block"
	Inline Format = "inline"
)

type Editable interface {
	EditorView(ctx context.Context, node *node.Node, format Format) vecty.Component
	Format(rule *system.RuleWrapper) Format
}
