package views

import (
	"context"

	"frizz.io/edit/editor/stores"
	"frizz.io/flux"
)

type View struct {
	*flux.View
	App *stores.App
}

func New(ctx context.Context, app *stores.App, self flux.ViewInterface) *View {
	return &View{
		View: flux.NewView(ctx, self, app),
		App:  app,
	}
}
