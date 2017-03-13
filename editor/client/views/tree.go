package views // import "kego.io/editor/client/views"

import (
	"context"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type TreeView struct {
	*View

	Root     *models.BranchModel
	Selected *models.BranchModel
}

func NewTreeView(ctx context.Context) *TreeView {
	v := &TreeView{}
	v.View = New(ctx, v)
	v.Watch(v.App.Branches.Root(),
		stores.BranchInitialStateLoaded,
		stores.ViewChanged,
	)
	return v
}

func (v *TreeView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.Root = v.App.Branches.Root()
	vecty.Rerender(v)
}

func (v *TreeView) Render() *vecty.HTML {
	if v.Root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content content-tree"),
		NewBranchView(v.Ctx, v.Root),
	)
}
