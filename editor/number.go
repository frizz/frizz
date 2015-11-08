package editor

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
)

type NumberEditor struct {
	*Node
	*Common
	Changes  chan float64
	original float64
}

func NewNumberEditor(n *Node) *NumberEditor {
	return &NumberEditor{Node: n, Common: &Common{}}
}

func (e *NumberEditor) Layout() Layout {
	return Inline
}

var _ Editor = (*NumberEditor)(nil)

func (e *NumberEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)
	e.Changes = make(chan float64, 1)

	e.original = e.ValueNumber

	tb := mdl.NewTextbox(strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
	e.AppendChild(tb)
	tb.Input.AddEventListener("input", true, func(ev dom.Event) {
		n, err := strconv.ParseFloat(tb.Input.Value, 64)
		if err != nil {
			// display a validation error
		}
		e.Missing = false
		e.Null = false
		e.ValueNumber = n
		e.MarkDirty(e.Dirty())
		select {
		case e.Changes <- e.ValueNumber:
		default:
		}
	})

	return nil
}

func (e *NumberEditor) Dirty() bool {
	return e.ValueNumber != e.original
}
