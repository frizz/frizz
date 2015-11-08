package units

import (
	"strconv"

	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/mdl"
)

func (i *Rectangle) GetEditor(n *editor.Node) editor.Editor {
	return &RectangleEditor{Rectangle: n.Value.(*Rectangle), Node: n, Common: &editor.Common{}}
}

var _ editor.Editable = (*Rectangle)(nil)

type RectangleEditor struct {
	*Rectangle
	*editor.Common
	*editor.Node
	height *mdl.Textbox
	width  *mdl.Textbox
}

var _ editor.Editor = (*RectangleEditor)(nil)

func (e *RectangleEditor) Layout() editor.Layout {
	return editor.Inline
}

func (e *RectangleEditor) Initialize(panel *dom.HTMLDivElement, holder editor.Holder, layout editor.Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, holder, layout, path, aliases)

	e.height = mdl.NewTextbox(e.Rectangle.Height.String(), "height")
	e.width = mdl.NewTextbox(e.Rectangle.Width.String(), "width")
	e.width.Style().Set("width", "50%")
	e.height.Style().Set("width", "50%")

	e.Panel.AppendChild(e.height)
	e.Panel.AppendChild(e.width)

	e.height.AddEventListener("input", true, func(ev dom.Event) {
		e.Update()
	})

	e.width.AddEventListener("input", true, func(ev dom.Event) {
		e.Update()
	})

	e.Update()

	return nil
}

func (e *RectangleEditor) Update() {
	height, _ := strconv.Atoi(e.height.Input.Value)
	e.Height.Set(height)
	width, _ := strconv.Atoi(e.width.Input.Value)
	e.Width.Set(width)
}
