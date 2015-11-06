package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
)

type Branch interface {
	Append(child Branch)
	Select(fromKeyboard bool)
	Unselect()
	Open()
	Close()
	Element() *dom.HTMLDivElement
	Tree() *Tree
	Root() bool
	Parent() Branch
	Child(i int) Branch
	Children() []Branch
	IsOpen() bool
	CanOpen() bool
	Sibling(i int) Branch
	Siblings() []Branch
	Index() int
	LastVisible() Branch
	PrevVisible() Branch
	NextVisible(includeChildren bool) Branch
	Update()
	IsVisible() bool
	SetDirtyIconState()
	Level() int
	SetParent(parent Branch)
	SetIndex(index int)
	DirtyDescendant(state bool, descendant Branch)
}

type Async interface {
	LoadContent() chan bool
	ContentLoaded() bool
}

type Editable interface {
	Editor() editor.Editor
}
