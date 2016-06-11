package stores // import "kego.io/editor/client/stores"

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/flux"
)

type App struct {
	Dispatcher flux.DispatcherInterface
	Fail       chan error
	Conn       connection.Interface

	Package  *PackageStore
	Editors  *EditorStore
	Branches *BranchStore
	Nodes    *NodeStore
	Panels   *PanelStore
	Types    *TypeStore
	Data     *DataStore
	Misc     *MiscStore
}

func (app *App) Init(ctx context.Context) {
	app.Package = NewPackageStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Nodes = NewNodeStore(ctx)
	app.Panels = NewPanelStore(ctx)
	app.Types = NewTypeStore(ctx)
	app.Data = NewDataStore(ctx)
	app.Misc = NewMiscStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Package,
		app.Editors,
		app.Branches,
		app.Nodes,
		app.Panels,
		app.Types,
		app.Data,
		app.Misc,
	)
}

func (a *App) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.Dispatcher.Dispatch(action)
}
