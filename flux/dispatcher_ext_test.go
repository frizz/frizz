package flux_test

import (
	"testing"

	"frizz.io/flux"
	"frizz.io/flux/mock_flux"
	"github.com/golang/mock/gomock"
)

func TestDispatcher_Notify(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	not := mock_flux.NewMockNotifierInterface(m)
	a := &st1{}
	gomock.InOrder(not.EXPECT().NotifyWithData("a", notif1{}, nil).Return(closedChannel()))
	d := flux.NewDispatcher(not, a)
	done := d.Dispatch(false)
	<-done
}

func closedChannel() chan struct{} {
	c := make(chan struct{}, 1)
	close(c)
	return c
}

func TestDispatcher_NotifyAfter(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	not := mock_flux.NewMockNotifierInterface(m)
	a := &st1{}
	gomock.InOrder(not.EXPECT().NotifyWithData("b", notif1{}, nil).Return(closedChannel()))
	d := flux.NewDispatcher(not, a)
	done := d.Dispatch(true)
	<-done
	close(handleDone)
	<-notifyDone
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
