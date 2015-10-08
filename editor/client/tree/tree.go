package tree // import "kego.io/editor/client/tree"

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/shared/connection"
)

type Tree struct {
	Root     *Branch
	selected *Branch
	Conn     *connection.Conn
	Fail     chan error
}

func New(parent dom.Element, conn *connection.Conn) *Tree {

	tree := &Tree{Conn: conn}
	tree.Root = &Branch{Tree: tree, root: true, open: true, item: &root{}}

	// We must tolerate passing in a nil dom element in order to run tests in pure go
	if parent == nil {
		tree.Root.Initialise()
		return tree
	}

	nodeDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	nodeDiv.SetAttribute("class", "node root")
	parent.AppendChild(nodeDiv)

	contentDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	contentDiv.SetAttribute("class", "content")
	nodeDiv.AppendChild(contentDiv)

	innerDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	innerDiv.SetAttribute("class", "children")
	nodeDiv.AppendChild(innerDiv)

	tree.Root.element = nodeDiv
	tree.Root.inner = innerDiv
	tree.Root.content = contentDiv

	tree.Root.Initialise()
	return tree
}

func (t *Tree) Branch(item Item) *Branch {
	return &Branch{item: item, Tree: t}
}
