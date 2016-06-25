package stores // import "kego.io/editor/client/stores"

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/flux"
)

type App struct {
	Dispatcher flux.DispatcherInterface
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
	app.Rule = NewRuleStore(ctx)
	app.Dispatcher = flux.NewDispatcher(
		app.Package,
		app.Editors,
		app.Branches,
		app.Nodes,
		app.Panels,
		app.Types,
		app.Data,
		app.Misc,
		app.Rule,
	)
	app.Notifier = flux.NewNotifier()
}

func (a *App) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.Dispatcher.Dispatch(action)
}

func (a *App) Watch(object interface{}, notif ...flux.Notif) chan flux.NotifPayload {
	return a.Notifier.Watch(object, notif...)
}

func (a *App) Notify(object interface{}, notif flux.Notif) (done chan struct{}) {
	return a.Notifier.Notify(object, notif)
}

func (a *App) NotifyWithData(object interface{}, notif flux.Notif, data interface{}) (done chan struct{}) {
	return a.Notifier.NotifyWithData(object, notif, data)
}

func (a *App) Delete(c chan flux.NotifPayload) {
	a.Notifier.Delete(c)
}
