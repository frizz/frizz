package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type BoolEditor struct {
	*Node
	*Editor
	original bool
}

var _ EditorInterface = (*BoolEditor)(nil)

func NewBoolEditor(n *Node) *BoolEditor {
	return &BoolEditor{Node: n, Editor: &Editor{}}
}

func (e *BoolEditor) Layout() Layout {
	return Inline
}

func (e *BoolEditor) Initialize(holder BranchInterface, layout Layout, path string, aliases map[string]string) error {

	e.Editor.Initialize(holder, layout, path, aliases)

	e.original = e.ValueBool

	cb := mdl.Checkbox(e.ValueBool, e.Node.Key)
	cb.Input.AddEventListener("change", true, func(ev dom.Event) {
		e.update(cb.Input.Checked)
		e.Send(e)
	})
	e.AppendChild(cb)

	return nil
}

func (e *BoolEditor) update(b bool) {
	e.Missing = false
	e.Null = false
	e.ValueBool = b
	e.Node.Value = b
}

func (e *BoolEditor) Dirty() bool {
	return e.ValueBool != e.original
}

func (e *BoolEditor) Value() interface{} {
	return e.Node.Value
}
