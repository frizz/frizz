package images

import (
	"kego.io/editor"
	"kego.io/editor/mdl"
)

func (i *Icon) GetEditor(n *editor.Node) editor.Editor {
	return &IconEditor{Icon: n.Value.(*Icon), Node: n, Common: &editor.Common{}}
}

var _ editor.Editable = (*Icon)(nil)

type IconEditor struct {
	*Icon
	*editor.Common
	*editor.Node
	Changes chan *Icon
	image   *mdl.Image
	url     *editor.StringEditor
}

var _ editor.Editor = (*IconEditor)(nil)

func (e *IconEditor) Layout() editor.Layout {
	return editor.Block
}

func (e *IconEditor) Initialize(holder editor.Holder, layout editor.Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)
	e.Changes = make(chan *Icon, 1)

	e.image = mdl.NewImage(e.Url.Value())
	e.AppendChild(e.image)

	e.url = editor.NewStringEditor(e.Node.Map["url"])
	e.url.Initialize(holder, editor.Block, path, aliases)
	e.Editors = append(e.Editors, e.url)
	e.AppendChild(e.url)

	go func() {
		for {
			e.update(<-e.url.Changes)
			e.notify()
		}
	}()

	e.update(e.Url.Value())

	return nil
}

func (e *IconEditor) update(url string) {
	e.Url.Set(url)
	e.image.Src = url
	e.image.Visibility(url != "")
	e.MarkDirty(e.Dirty())
}

func (e *IconEditor) notify() {
	select {
	case e.Changes <- e.Icon:
	default:
	}
}

func (e *IconEditor) AddChildTreeEntry(child editor.Editor) bool {
	return false
}

func (e *IconEditor) Focus() {
	e.url.Focus()
}
