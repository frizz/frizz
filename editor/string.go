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
	textbox  *mdl.TextboxStruct
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

	e.textbox = mdl.Textbox(e.ValueString, e.Node.Key)
	e.textbox.Style().Set("width", "100%")
	e.AppendChild(e.textbox)
	e.textbox.Input.AddEventListener("input", true, func(ev dom.Event) {
		e.update(e.textbox.Input.Value)
		e.Notify(e)
	})

	return nil
}

func (e *StringEditor) update(s string) {
	e.Missing = false
	e.Null = false
	e.ValueString = s
	e.Node.Value = s
}

func (e *StringEditor) Dirty() bool {
	return e.ValueString != e.original
}

func (e *StringEditor) Focus() {
	e.textbox.Input.Focus()
}

func (e *StringEditor) Value() interface{} {
	return e.Node.Value
}
