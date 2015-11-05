package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/system/node"
)

type Item interface {
	Initialise(*dom.HTMLSpanElement)
}

type Async interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type Noder interface {
	Node() *node.Node
}
