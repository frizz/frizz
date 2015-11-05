package editor

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/system/node"
)

type NodeNumberEditor struct {
	*node.Node
	*Common
	original float64
}

var _ Editor = (*NodeNumberEditor)(nil)

func (e *NodeNumberEditor) Initialize(panel *dom.HTMLDivElement, dirtyable Dirtyable, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, dirtyable, path, aliases)

	e.original = e.ValueNumber

	tb := mdl.NewTextbox(strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
	e.Panel.AppendChild(tb)
	tb.Input.AddEventListener("input", true, func(ev dom.Event) {
		n, err := strconv.ParseFloat(tb.Input.Value, 64)
		if err != nil {
			// display a validation error
		}
		e.Missing = false
		e.Null = false
		e.ValueNumber = n
		e.dirtyable.MarkDirty(e.Dirty())
	})

	return nil
}

func (e *NodeNumberEditor) Dirty() bool {
	return e.ValueNumber != e.original
}
