package stores // import "kego.io/editor/client/stores"

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/flux"
)

type App struct {
	Dispatcher *flux.Dispatcher
	Fail       chan error
	Conn       *connection.Conn

	Nodes    *NodeStore
	Editors  *EditorStore
	Branches *BranchStore
}

func (app *App) Init(ctx context.Context) {
	app.Nodes = NewNodeStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Nodes,
		app.Editors,
		app.Branches,
	)
}
