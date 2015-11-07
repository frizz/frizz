package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type NodeStringEditor struct {
	*Node
	*Common
	original string
}

func (e *NodeStringEditor) Layout() Layout {
	return Inline
}

var _ Editor = (*NodeStringEditor)(nil)

func (e *NodeStringEditor) Initialize(panel *dom.HTMLDivElement, holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, holder, layout, path, aliases)

	e.original = e.ValueString

	tb := mdl.NewTextbox(e.ValueString, e.Node.Key)
	tb.Style().Set("width", "100%")
	e.Panel.AppendChild(tb)
	tb.Input.AddEventListener("input", true, func(ev dom.Event) {
		e.Missing = false
		e.Null = false
		e.ValueString = tb.Input.Value
		e.holder.MarkDirty(e.Dirty())
	})

	return nil
}

func (e *NodeStringEditor) Dirty() bool {
	return e.ValueString != e.original
}
