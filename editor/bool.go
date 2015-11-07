package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type NodeBoolEditor struct {
	*Node
	*Common
	original bool
}

func (e *NodeBoolEditor) Layout() Layout {
	return Inline
}

var _ Editor = (*NodeBoolEditor)(nil)

func (e *NodeBoolEditor) Initialize(panel *dom.HTMLDivElement, holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, holder, layout, path, aliases)

	e.original = e.ValueBool

	cb := mdl.NewCheckbox(e.ValueBool, e.Node.Key)
	e.Panel.AppendChild(cb)
	cb.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.Missing = false
		e.Null = false
		e.ValueBool = cb.Input.Checked
		e.holder.MarkDirty(e.Dirty())
	})

	return nil
}

func (e *NodeBoolEditor) Dirty() bool {
	return e.ValueBool != e.original
}
