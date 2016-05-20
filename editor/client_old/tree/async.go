package tree

import (
	"net/rpc"

	"github.com/davelondon/kerr"
)

type AsyncInterface interface {
	LoadContent(fail chan error) (success chan bool, loading bool)
	Loaded() bool
	MakeCall(fail chan error) (requestCall *rpc.Call)
	ProcessResponse(interface{}) error
	Cancel()
}

func NewAsync(self AsyncInterface) *Async {
	return &Async{self: self}
}

type Async struct {
	self    AsyncInterface
	loading bool
	loaded  bool
}

func (a *Async) Cancel() {
	a.loading = false
}

func (a *Async) Loaded() bool {
	return a.loaded
}

func (a *Async) LoadContent(fail chan error) (success chan bool, loading bool) {

	if a.loading {
		return nil, true
	}

	successChannel := make(chan bool, 1)

	if a.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it.
		successChannel <- true
		return successChannel, false
	}

	a.loading = true

	call := a.self.MakeCall(fail)

	go func() {

		<-call.Done

		if err := a.self.ProcessResponse(call.Reply); err != nil {
			fail <- kerr.Wrap("DLTCGMSREX", err)
		}

		a.loaded = true
		a.loading = false

		successChannel <- true

	}()

	return successChannel, false

}
