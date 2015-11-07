package editors // import "kego.io/system/editors"

/*
func init() {
	// Will leave this disabled for now... I think it's not really
	// needed. "Exists" checkbox shouldn't be in the editor UI.

	//editor.Register(system.NewReference("kego.io/system", "string"), editor.Factory(func(n *node.Node) editor.Editor {
	//return &StringEditor{String: n.Value.(system.String), node: n, Common: &editor.Common{}}
	//}))
}

var _ editor.Editor = (*StringEditor)(nil)

type StringEditor struct {
	*system.String
	*editor.Common
	node    *node.Node
	textbox *mdl.Textbox
}

func (e *StringEditor) Initialize(panel *dom.HTMLDivElement, holder editor.Holder, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, holder, path, aliases)

	e.textbox = mdl.NewTextbox(e.Value(), e.node.Key)
	e.Panel.AppendChild(e.textbox)
	e.textbox.AddEventListener("change", true, func(ev dom.Event) {
		e.node.ValueString = e.textbox.Input.Value
	})

	return nil
}
*/
