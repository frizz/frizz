package stores

import (
	"context"

	"frizz.io/flux"
)

type App struct {
	Dispatcher flux.DispatcherInterface
	Watcher    flux.WatcherInterface
	Notifier   flux.NotifierInterface
	Fail       chan error

	Client *ClientStore
}

func (app *App) Init(ctx context.Context) {

	n := flux.NewNotifier()
	app.Notifier = n
	app.Watcher = n

	app.Client = NewClientStore(ctx, app)
	app.Dispatcher = flux.NewDispatcher(
		app.Notifier, // NotifierInterface
		app.Client,   // StoreInterface...
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
