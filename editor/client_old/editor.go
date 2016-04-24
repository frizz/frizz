package editor // import "kego.io/editor/client_old"

import (
	"github.com/tjgq/broadcast"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
)

type Layout int

const (
	Inline Layout = iota
	Block
	Page
)

type Editable interface {
	GetEditor(n *Node) EditorInterface
}

type BranchInterface interface {
	Select(fromKeyboard bool)
	ListenForEditorChanges(changes <-chan interface{})
	AddFieldNode(node *Node)
}

type EditorInterface interface {
	dom.Node
	Initialize(ctx context.Context, branch BranchInterface, layout Layout, fail chan error) error
	Show()
	Hide()
	Layout() Layout
	AddChildTreeEntry(child EditorInterface) bool
	Dirty() bool
	Select()
	Focus()
	Value() interface{}
	Send(v interface{})
	Listen() *broadcast.Listener
}

type Editor struct {
	*dom.HTMLDivElement
	fail    chan error
	Editors []EditorInterface
	layout  Layout
	branch  BranchInterface
	changes *broadcast.Broadcaster
	ctx     context.Context
}

func (c *Editor) Send(v interface{}) {
	c.changes.Send(v)
}

func (c *Editor) Listen() *broadcast.Listener {
	return c.changes.Listen()
}

func (c *Editor) Initialize(ctx context.Context, branch BranchInterface, layout Layout, fail chan error) error {

	c.HTMLDivElement = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	if layout == Page {
		c.Style().Set("display", "none")
		c.Class().SetString("mdl-color--white mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-grid")
	}
	if layout == Inline {
		c.Style().Set("display", "inline-block")
	}

	c.changes = broadcast.New(0)

	branch.ListenForEditorChanges(c.changes.Listen().Ch)

	c.ctx = ctx
	c.branch = branch
	c.layout = layout
	c.fail = fail
	return nil

}

func (c *Editor) Show() {
	if c.layout == Inline {
		c.Style().Set("display", "inline-block")
	} else {
		c.Style().Set("display", "block")
	}
}

func (c *Editor) Hide() {
	c.Style().Set("display", "none")
}

func (c *Editor) AddChildTreeEntry(child EditorInterface) bool {
	return true
}

func (c *Editor) Dirty() bool {
	for _, e := range c.Editors {
		if e.Dirty() {
			return true
		}
	}
	return false
}

func (c *Editor) Select() {
	c.branch.Select(false)
}

func (c *Editor) Focus() {}
