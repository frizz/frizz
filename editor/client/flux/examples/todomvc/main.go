package main

import (
	"github.com/gopherjs/vecty"
	"golang.org/x/net/context"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/flux/examples/todomvc/stores"
	"kego.io/editor/client/flux/examples/todomvc/views"
)

func main() {
	vecty.SetTitle("Flux • TodoMVC")
	vecty.AddStylesheet("css/base.css")
	vecty.AddStylesheet("css/index.css")

	ctx, app := stores.NewApp(context.Background())
	app.Todos = stores.NewTodoStore(ctx)
	app.Dispatcher = flux.NewDispatcher(app.Todos)

	page := views.NewPageView(ctx)
	vecty.RenderAsBody(page)
}
