package editor

import (
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/editor/client/mdl"
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

func (e *BoolEditor) Initialize(ctx context.Context, holder BranchInterface, layout Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

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
