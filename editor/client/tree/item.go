package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
)

type Item interface {
	Initialise(*dom.HTMLSpanElement)
}

type Async interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type Editable interface {
	Editor() editor.Editor
}
