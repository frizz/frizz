package tree

import "honnef.co/go/js/dom"

type Item interface {
	Initialise(dom.Element)
	Open(node *Node, success func(), fail func(error))
	AsyncOpen() bool
	ContentLoaded() bool
}

type root struct {
}

func (r *root) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	label.SetTextContent("root")
	div.AppendChild(label)
}

func (r *root) Open(*Node, func(), func(error)) {
	// Root node can't be opened
}

func (r *root) AsyncOpen() bool {
	return false
}
func (r *root) ContentLoaded() bool {
	return true
}
