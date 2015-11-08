package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type BoolEditor struct {
	*Node
	*Common
	Changes  chan bool
	original bool
}

func NewBoolEditor(n *Node) *BoolEditor {
	return &BoolEditor{Node: n, Common: &Common{}}
}

func (e *BoolEditor) Layout() Layout {
	return Inline
}

var _ Editor = (*BoolEditor)(nil)

func (e *BoolEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)
	e.Changes = make(chan bool, 1)

	e.original = e.ValueBool

	cb := mdl.NewCheckbox(e.ValueBool, e.Node.Key)
	cb.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.Missing = false
		e.Null = false
		e.ValueBool = cb.Input.Checked
		e.MarkDirty(e.Dirty())
		select {
		case e.Changes <- e.ValueBool:
		default:
		}
	})
	e.AppendChild(cb)

	return nil
}

func (e *BoolEditor) Dirty() bool {
	return e.ValueBool != e.original
}
