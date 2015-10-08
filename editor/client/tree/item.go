package tree

import "honnef.co/go/js/dom"

type Item interface {
	Initialise(*dom.HTMLDivElement)
}

type AsyncItem interface {
	LoadContent() chan bool
	ContentLoaded() bool
}
