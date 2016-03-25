package stores // import "kego.io/editor/client/stores"

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/flux"
)

type App struct {
	*flux.Dispatcher
	Fail chan error
	Conn *connection.Conn

	Package  *PackageStore
	Selected *SelectedStore
	Editors  *EditorStore
	Branches *BranchStore
	Panels   *PanelStore
	Types    *TypesStore
	Data     *DataStore
}

func (app *App) Init(ctx context.Context) {
	app.Package = NewPackageStore(ctx)
	app.Selected = NewSelectedStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Panels = NewPanelStore(ctx)
	app.Types = NewTypesStore(ctx)
	app.Data = NewDataStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Package,
		app.Selected,
		app.Editors,
		app.Branches,
		app.Panels,
		app.Types,
		app.Data,
	)
}
