package flux // import "frizz.io/flux"

// frizz: {"package": {"complete": true}}

type ActionInterface interface{}

type AppInterface interface {
	WatcherInterface
	DispatcherInterface
}

type DispatcherInterface interface {
	Dispatch(action ActionInterface) chan struct{}
}

type WatcherInterface interface {
	// Watch returns a channel subscribed to reveive notifs notifications about
	// object. If object is nil, the subscription is for all objects in the
	// store.
	Watch(object interface{}, notif ...Notif) chan NotifPayload
	Delete(c chan NotifPayload)
}

type NotifierInterface interface {
	// Notify sends the notif notification to all subscribers of that
	// notification for object. If object is nil, the notification is sent to
	// all subscribers. The chanel returned is closed when the notify action
	// has finished.
	Notify(object interface{}, notif Notif) (done chan struct{})
	NotifyWithData(object interface{}, notif Notif, data interface{}) (done chan struct{})
}
