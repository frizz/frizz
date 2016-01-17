package tree

import (
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/kerr"
)

type AsyncInterface interface {
	LoadContent(conn *connection.Conn, fail chan error) (success chan bool, loading bool)
	Loaded() bool
	ContentRequest() messages.MessageInterface
	ProcessResponse(messages.MessageInterface) error
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

func (a *Async) LoadContent(conn *connection.Conn, fail chan error) (success chan bool, loading bool) {

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

	responseChannel := conn.Request(a.self.ContentRequest())

	go func() {

		m := <-responseChannel

		if err := a.self.ProcessResponse(m); err != nil {
			fail <- kerr.Wrap("DLTCGMSREX", err)
		}

		a.loaded = true
		a.loading = false

		successChannel <- true

	}()

	return successChannel, false

}
