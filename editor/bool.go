package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/system/node"
)

type NodeBoolEditor struct {
	*node.Node
	*Common
	original bool
}

var _ Editor = (*NodeBoolEditor)(nil)

func (e *NodeBoolEditor) Initialize(panel *dom.HTMLDivElement, dirtyable Dirtyable, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, dirtyable, path, aliases)

	e.original = e.ValueBool

	cb := mdl.NewCheckbox(e.ValueBool, e.Node.Key)
	e.Panel.AppendChild(cb)
	cb.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.Missing = false
		e.Null = false
		e.ValueBool = cb.Input.Checked
		e.dirtyable.MarkDirty(e.Dirty())
	})

	return nil
}

func (e *NodeBoolEditor) Dirty() bool {
	return e.ValueBool != e.original
}
