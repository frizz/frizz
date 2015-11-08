package units

import "kego.io/editor"

func (i *Rectangle) GetEditor(n *editor.Node) editor.Editor {
	return &RectangleEditor{Rectangle: n.Value.(*Rectangle), Node: n, Common: &editor.Common{}}
}

var _ editor.Editable = (*Rectangle)(nil)

type RectangleEditor struct {
	*Rectangle
	*editor.Common
	*editor.Node
	Changes chan *Rectangle
	height  *editor.NumberEditor
	width   *editor.NumberEditor
}

var _ editor.Editor = (*RectangleEditor)(nil)

func (e *RectangleEditor) Layout() editor.Layout {
	return editor.Inline
}

func (e *RectangleEditor) Initialize(holder editor.Holder, layout editor.Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)
	e.Changes = make(chan *Rectangle, 1)

	e.height = editor.NewNumberEditor(&editor.Node{e.Node.Fields["height"]})
	e.height.Initialize(holder, editor.Inline, path, aliases)
	e.height.Style().Set("width", "50%")

	e.width = editor.NewNumberEditor(&editor.Node{e.Node.Fields["width"]})
	e.width.Initialize(holder, editor.Inline, path, aliases)
	e.width.Style().Set("width", "50%")

	e.AppendChild(e.height)
	e.AppendChild(e.width)

	go func() {
		for {
			select {
			case height := <-e.height.Changes:
				e.Height.Set(int(height))
				e.update()
			case width := <-e.width.Changes:
				e.Width.Set(int(width))
				e.update()
			}
		}
	}()

	return nil
}

func (e *RectangleEditor) update() {
	select {
	case e.Changes <- e.Rectangle:
	default:
	}
}
