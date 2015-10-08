package items

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/tree"
)

type root struct{}

func Root() *root {
	return &root{}
}

var _ tree.Item = (*root)(nil)

func (r *root) Initialise(div *dom.HTMLDivElement) {}
