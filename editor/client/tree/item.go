package tree

import "kego.io/editor"

type Item interface {
}

type Async interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type Editable interface {
	Editor() editor.Editor
}
