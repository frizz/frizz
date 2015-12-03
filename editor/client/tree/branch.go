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

	loading  bool
	element  *dom.HTMLDivElement
	opener   *dom.HTMLAnchorElement
	inner    *dom.HTMLDivElement
	content  *dom.HTMLDivElement
	label    *dom.HTMLSpanElement
	badge    *dom.HTMLSpanElement
	selected bool
	editor   editor.EditorInterface
	dirty    map[editor.EditorInterface]bool // descendant editors that have changes
}

func NewBranch(self BranchInterface, parent BranchInterface) *Branch {
	b := &Branch{tree: parent.branch().tree, parent: parent.branch(), self: self}
	b.init()
	return b
}

func (b *Branch) Append(child BranchInterface) {
	b.append(child.branch())
}

func (b *Branch) branch() *Branch {
	return b
}

func (b *Branch) append(child *Branch) {
	if b.inner != nil {
		b.inner.AppendChild(child.element)
	}
	b.children = append(b.children, child)
	b.updateVisibleDescendants(true, true)
}

// updateVisibleDescendants works recursively for all open children. It calls updateParent and
// updateOpenerIcon for each descendant.
func (b *Branch) updateVisibleDescendants(updateOpenerIcon bool, updateParent bool) {
	if updateOpenerIcon {
		b.updateOpenerIcon()
	}
	if b.opened && len(b.children) > 0 {
		for index, child := range b.children {
			if updateParent {
				child.updateParent(index, b)
			}
			if updateOpenerIcon {
				child.updateOpenerIcon()
			}
		}
	}
}

// updateChildren assumes the children array is the only source of truth, and updates index and
// parent properties of the children. It works recursively for all open children.
func (b *Branch) updateParent(index int, parent *Branch) {
	b.parent = parent
	b.index = index
}

func (b *Branch) setLabel(text string) {
	b.label.SetTextContent(text)
}
