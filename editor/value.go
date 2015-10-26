package editor

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/json"
	"kego.io/system/node"
)

type NodeValueEditor struct {
	*node.Node
	*Common
}

var _ Editor = (*NodeValueEditor)(nil)

func (e *NodeValueEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.Panel = panel
	e.Path = path
	e.Aliases = aliases

	switch e.JsonType {
	case json.J_STRING:
		tb := mdl.NewTextbox(e.ValueString, e.Node.Key)
		e.Panel.AppendChild(tb)
		tb.Input.AddEventListener("change", true, func(ev dom.Event) {
			e.Missing = false
			e.Null = false
			e.ValueString = tb.Input.Value
		})
	case json.J_NUMBER:
		tb := mdl.NewTextbox(strconv.FormatFloat(e.ValueNumber, 'f', -1, 64), e.Node.Key)
		e.Panel.AppendChild(tb)
		tb.Input.AddEventListener("change", true, func(ev dom.Event) {
			n, err := strconv.ParseFloat(tb.Input.Value, 64)
			if err != nil {
				// display a validation error
			}
			e.Missing = false
			e.Null = false
			e.ValueNumber = n
		})
	case json.J_BOOL:
		cb := mdl.NewCheckbox(e.ValueBool, e.Node.Key)
		e.Panel.AppendChild(cb)
		cb.Input.AddEventListener("change", true, func(ev dom.Event) {
			e.Missing = false
			e.Null = false
			e.ValueBool = cb.Input.Checked
		})
	}

	e.Initialized = true
	return nil
}
