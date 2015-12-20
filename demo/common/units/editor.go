package units

import (
	"golang.org/x/net/context"
	"kego.io/editor"
)

func (i *Rectangle) GetEditor(n *editor.Node) editor.EditorInterface {
	return &RectangleEditor{Rectangle: n.Value.(*Rectangle), Node: n, Editor: &editor.Editor{}}
}

var _ editor.Editable = (*Rectangle)(nil)

type RectangleEditor struct {
	*Rectangle
	*editor.Editor
	*editor.Node
	height *editor.NumberEditor
	width  *editor.NumberEditor
}

var _ editor.EditorInterface = (*RectangleEditor)(nil)

func (e *RectangleEditor) Layout() editor.Layout {
	return editor.Inline
}

func (e *RectangleEditor) Initialize(ctx context.Context, holder editor.BranchInterface, layout editor.Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

	e.height = editor.NewNumberEditor(e.Node.Map["height"].(*editor.Node))
	e.height.Initialize(ctx, holder, editor.Inline, fail)
	e.height.Style().Set("width", "50%")

	e.width = editor.NewNumberEditor(e.Node.Map["width"].(*editor.Node))
	e.width.Initialize(ctx, holder, editor.Inline, fail)
	e.width.Style().Set("width", "50%")

	e.Editors = append(e.Editors, e.height)
	e.Editors = append(e.Editors, e.width)

	e.AppendChild(e.height)
	e.AppendChild(e.width)

	go func() {
		width := e.width.Listen().Ch
		height := e.height.Listen().Ch
		for {
			select {
			case ne := <-width:
				e.Width.Set(int(ne.(*editor.NumberEditor).ValueNumber))
			case ne := <-height:
				e.Height.Set(int(ne.(*editor.NumberEditor).ValueNumber))
			}
			e.Send(e)
		}
	}()

	return nil
}

func (e *RectangleEditor) Focus() {
	e.height.Focus()
}

func (e *RectangleEditor) Value() interface{} {
	return e.Rectangle
}
