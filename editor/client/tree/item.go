package tree

import "honnef.co/go/js/dom"

type Item interface {
	Initialise(dom.Element)
}

type AsyncItem interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type root struct {
}

func (r *root) Initialise(div dom.Element) {}
