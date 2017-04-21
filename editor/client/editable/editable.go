package editable

import (
	"context"

	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/vecty"
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

type EditsExtraEmbeddedTypes interface {
	ExtraEmbeddedTypes() []*system.Reference
}
