package system

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
)

type NodeObjectEditor struct {
	*Node
	*editorCommon
}

var _ Editor = (*NodeObjectEditor)(nil)

func (e *NodeObjectEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	table := mdl.Table()

	names := table.Column("name")
	origins := table.Column("origin")
	holds := table.Column("holds")
	values := table.Column("value")

	for name, field := range e.Fields {

		names.Cell(name)

		origin, err := field.Origin.ValueContext(e.path, e.aliases)
		if err != nil {
			return kerr.New("ACQLJXWYQX", err, "ValueContext")
		}
		origins.Cell(origin)

		hold, err := field.Rule.HoldsDisplayType(e.path, e.aliases)
		if err != nil {
			return kerr.New("OYMARPFDGA", err, "ValueContext")
		}
		holds.Cell(hold)

		if field.Missing || field.Null {
			values.Cell("")
		} else {
			value, err := field.Type.Id.ValueContext(e.path, e.aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			values.Cell(value)
		}

	}
	e.panel.AppendChild(table.Build())

	e.initialized = true
	return nil
}

func (e *NodeObjectEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
