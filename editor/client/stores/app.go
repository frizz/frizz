package stores // import "kego.io/editor/client/stores"

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/flux"
)

type App struct {
	Dispatcher flux.DispatcherInterface
	Watcher    flux.WatcherInterface
	Notifier   flux.NotifierInterface
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
	Rule     *RuleStore
	Actions  *ActionStore
	Files    *FileStore
}

func (app *App) Init(ctx context.Context) {

	n := flux.NewNotifier()
	app.Notifier = n
	app.Watcher = n

	app.Package = NewPackageStore(ctx)
	app.Editors = NewEditorStore(ctx)
	app.Branches = NewBranchStore(ctx)
	app.Nodes = NewNodeStore(ctx)
	app.Panels = NewPanelStore(ctx)
	app.Types = NewTypeStore(ctx)
	app.Data = NewDataStore(ctx)
	app.Misc = NewMiscStore(ctx)
	app.Rule = NewRuleStore(ctx)
	app.Actions = NewActionStore(ctx)
	app.Files = NewFileStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Notifier, // NotifierInterface
		app.Package,  // StoreInterface...
		app.Editors,
		app.Branches,
		app.Nodes,
		app.Panels,
		app.Types,
		app.Data,
		app.Misc,
		app.Rule,
		app.Actions,
		app.Files,
	)
}

func (a *App) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.Dispatcher.Dispatch(action)
}

func (a *App) Watch(object interface{}, notif ...flux.Notif) chan flux.NotifPayload {
	return a.Watcher.Watch(object, notif...)
}

func (a *App) Delete(c chan flux.NotifPayload) {
	a.Watcher.Delete(c)
}
