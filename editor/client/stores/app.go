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

	Root     *RootStore
	Selected *SelectedStore
	Editors  *EditorStore
	Branches *BranchStore
	Panels   *PanelStore
}

func (app *App) Init(ctx context.Context) {
	app.Root = NewRootStore(ctx)
	app.Selected = NewSelectedStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Panels = NewPanelStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Root,
		app.Selected,
		app.Editors,
		app.Branches,
		app.Panels,
	)
}
