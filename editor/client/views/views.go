package views

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type View struct {
	*flux.View
	App *stores.App
}

func New(ctx context.Context, self flux.ViewInterface) *View {
	app := stores.FromContext(ctx)
	return &View{
		View: flux.NewView(ctx, self, app),
		App:  app,
	}
}
