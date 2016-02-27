package tree

import (
	"net/rpc"

	"kego.io/kerr"
)

type AsyncInterface interface {
	LoadContent(client *rpc.Client, fail chan error) (success chan bool, loading bool)
	Loaded() bool
	MakeRequest(*rpc.Client) (requestCall *rpc.Call, doneChannel chan *rpc.Call, data interface{})
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

func (a *Async) LoadContent(client *rpc.Client, fail chan error) (success chan bool, loading bool) {

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

	requestCall, done, data := a.self.MakeRequest(client)

	go func() {

		responseCall := <-done

		if requestCall != responseCall {
			fail <- kerr.New("TMPQETEMGC", "requestCall != responseCall")
		}

		if data == nil {
			fail <- kerr.New("XOFACTHWOV", "data == nil")
		}

		if err := a.self.ProcessResponse(data); err != nil {
			fail <- kerr.Wrap("DLTCGMSREX", err)
		}

		a.loaded = true
		a.loading = false

		successChannel <- true

	}()

	return successChannel, false

}
