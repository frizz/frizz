package images

import (
	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/mdl"
	"kego.io/system/node"
)

func (i *Icon) GetEditor(n *node.Node) editor.Editor {
	return &IconEditor{Icon: n.Value.(*Icon), node: n, Common: &editor.Common{}}
}

var _ editor.Editable = (*Icon)(nil)

type IconEditor struct {
	*Icon
	*editor.Common
	node    *node.Node
	image   *mdl.Image
	textbox *mdl.Textbox
}

var _ editor.Editor = (*IconEditor)(nil)

func (e *IconEditor) Initialize(panel *dom.HTMLDivElement, dirtyable editor.Dirtyable, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, dirtyable, path, aliases)

	e.image = mdl.NewImage(e.Url.Value())
	e.textbox = mdl.NewTextbox(e.Url.Value(), "url")

	e.Panel.AppendChild(e.image)
	e.Panel.AppendChild(e.textbox)

	e.textbox.AddEventListener("input", true, func(ev dom.Event) {
		e.Update()
	})

	e.Update()

	return nil
}

func (e *IconEditor) Update() {
	if e.textbox.Input.Value != "" {
		e.Url.Set(e.textbox.Input.Value)
		e.image.Src = e.Url.Value()
		e.image.Style().Set("display", "block")
	} else {
		e.Url.Set("")
		e.image.Src = ""
		e.image.Style().Set("display", "none")
	}
}
