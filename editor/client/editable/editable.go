package editable

import (
	"github.com/davelondon/vecty"
	"golang.org/x/net/context"
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
