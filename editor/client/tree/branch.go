package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
)

var _ BranchInterface = (*Branch)(nil)
var _ editor.BranchInterface = (*Branch)(nil)

type BranchInterface interface {
	Append(child BranchInterface)
	ListenForEditorChanges(changes <-chan interface{})
	Selectable
	KeyboardActions
	branch() *Branch
}

type Editable interface {
	Editor() editor.EditorInterface
}

type Branch struct {
	tree *Tree
	self BranchInterface

	parent   *Branch
	children []*Branch
	index    int
	opened   bool

	element    *dom.HTMLDivElement
	opener     *dom.HTMLAnchorElement
	inner      *dom.HTMLDivElement
	content    *dom.HTMLDivElement
	label      *dom.HTMLSpanElement
	badge      *dom.HTMLSpanElement
	selected   bool
	editor     editor.EditorInterface
	dirtyChild map[editor.EditorInterface]bool // editors of descendant branched that have changes
	dirtySelf  map[editor.EditorInterface]bool // editors of this branch that have changes
}

func NewBranch(self BranchInterface, parent BranchInterface) *Branch {
	b := &Branch{tree: parent.branch().tree, parent: parent.branch(), self: self}
	b.init()
	return b
}

func (b *Branch) Append(c BranchInterface) {
	child := c.branch()
	if b.inner != nil {
		b.inner.AppendChild(child.element)
	}
	child.parent = b
	child.index = len(b.children)
	b.children = append(b.children, child)
	b.updateOpenerIcon()
	child.updateOpenerIcon()
}

func (b *Branch) branch() *Branch {
	return b
}

func (b *Branch) setLabel(text string) {
	b.label.SetTextContent(text)
}
