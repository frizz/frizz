package flux_test

import (
	"testing"

	"kego.io/editor/client/ctests"
	"kego.io/flux"
)

func TestDispatcher_Notify(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()
	app := cb.GetApp()
	a := &st1{}
	app.Dispatcher = flux.NewDispatcher(app.Notifier, a)
	cb.ExpectNotified("a", notif1{}, nil)
	done := app.Dispatch(false)
	<-done
	cb.AssertAppSuccess()
}

func TestDispatcher_NotifyAfter(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()
	app := cb.GetApp()
	a := &st1{}
	app.Dispatcher = flux.NewDispatcher(app.Notifier, a)
	cb.ExpectNotified("b", notif1{}, nil)
	done := app.Dispatch(true)
	<-done
	close(handleDone)
	<-notifyDone
	cb.AssertAppSuccess()
}

var handleDone = make(chan struct{}, 1)
var notifyDone = make(chan struct{}, 1)

type notif1 struct{}

func (n notif1) IsNotif() {}

type st1 struct {
	*flux.Store
}

func (s *st1) Handle(payload *flux.Payload) (finished bool) {
	async := payload.Action.(bool)
	if async {
		go func() {
			<-handleDone
			payload.Notify("b", notif1{})
			close(notifyDone)
		}()
	} else {
		payload.Notify("a", notif1{})
	}
	return true
}
