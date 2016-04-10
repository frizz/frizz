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
	Editors  *EditorStore
	Branches *BranchStore
	Panels   *PanelStore
	Types    *TypeStore
	Data     *DataStore
}

func (app *App) Init(ctx context.Context) {
	app.Package = NewPackageStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Panels = NewPanelStore(ctx)
	app.Types = NewTypeStore(ctx)
	app.Data = NewDataStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Package,
		app.Editors,
		app.Branches,
		app.Panels,
		app.Types,
		app.Data,
	)
}
