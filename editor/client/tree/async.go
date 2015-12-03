package tree

import (
	"kego.io/editor/shared/connection"
	"kego.io/editor/shared/messages"
	"kego.io/kerr"
)

type AsyncInterface interface {
	LoadContent(conn *connection.Conn, fail chan error) chan bool
	Loaded() bool
	ContentRequest() messages.MessageInterface
	ProcessResponse(messages.MessageInterface) error
}

func NewAsync(self AsyncInterface) *Async {
	return &Async{self: self}
}

type Async struct {
	self   AsyncInterface
	loaded bool
}

func (a *Async) Loaded() bool {
	return a.loaded
}

func (a *Async) LoadContent(conn *connection.Conn, fail chan error) chan bool {

	successChannel := make(chan bool, 1)

	if a.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it.
		successChannel <- true
		return successChannel
	}

	responseChannel := conn.Request(a.self.ContentRequest())

	go func() {

		m := <-responseChannel

		if err := a.self.ProcessResponse(m); err != nil {
			fail <- kerr.New("DLTCGMSREX", err, "awaitSourceResponse")
		}

		a.loaded = true

		successChannel <- true

	}()

	return successChannel

}
