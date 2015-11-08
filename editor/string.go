package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type StringEditor struct {
	*Node
	*Common
	Changes  chan string
	original string
}

func NewStringEditor(n *Node) *StringEditor {
	return &StringEditor{Node: n, Common: &Common{}}
}

func (e *StringEditor) Layout() Layout {
	return Inline
}

var _ Editor = (*StringEditor)(nil)

func (e *StringEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)
	e.Changes = make(chan string, 1)

	e.original = e.ValueString

	tb := mdl.NewTextbox(e.ValueString, e.Node.Key)
	tb.Style().Set("width", "100%")
	e.AppendChild(tb)
	tb.Input.AddEventListener("input", true, func(ev dom.Event) {
		e.update(tb.Input.Value)
		e.notify()
	})

	return nil
}

func (e *StringEditor) update(s string) {
	e.Missing = false
	e.Null = false
	e.ValueString = s
	e.MarkDirty(e.Dirty())
}

func (e *StringEditor) notify() {
	select {
	case e.Changes <- e.ValueString:
	default:
	}
}

func (e *StringEditor) Dirty() bool {
	return e.ValueString != e.original
}
