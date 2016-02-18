package images

// ke: {"package": {"notest": true}}

import (
	"golang.org/x/net/context"
	"kego.io/editor"
	"kego.io/editor/client/mdl"
)

func (i *Photo) GetEditor(n *editor.Node) editor.EditorInterface {
	return &PhotoEditor{Photo: n.Value.(*Photo), Node: n, Editor: &editor.Editor{}}
}

var _ editor.Editable = (*Photo)(nil)

type PhotoEditor struct {
	*Photo
	*editor.Editor
	*editor.Node
	image *mdl.ImageStruct
	url   *editor.StringEditor
}

var _ editor.EditorInterface = (*PhotoEditor)(nil)

func (e *PhotoEditor) Layout() editor.Layout {
	return editor.Block
}

func (e *PhotoEditor) Initialize(ctx context.Context, holder editor.BranchInterface, layout editor.Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

	e.url = editor.NewStringEditor(e.Node.Map["url"].(*editor.Node))
	e.url.Initialize(ctx, holder, editor.Block, fail)
	e.Editors = append(e.Editors, e.url)
	e.AppendChild(e.url)

	e.image = mdl.Image(e.Url.Value())
	e.AppendChild(e.image)

	go func() {
		for se := range e.url.Listen().Ch {
			e.update(se.(*editor.StringEditor).ValueString)
			e.Send(e)
		}
	}()

	e.update(e.Url.Value())

	return nil
}

func (e *PhotoEditor) update(url string) {
	e.Url.Set(url)
	e.image.Src = url
	e.image.Visibility(url != "")
}

func (e *PhotoEditor) AddChildTreeEntry(child editor.EditorInterface) bool {
	return false
}

func (e *PhotoEditor) Focus() {
	e.url.Focus()
}

func (e *PhotoEditor) Value() interface{} {
	return e.Photo
}
