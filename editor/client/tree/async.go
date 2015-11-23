package tree

import (
	"kego.io/editor/shared/messages"
	"kego.io/kerr"
)

type AsyncInterface interface {
	LoadContent() chan bool
	Loaded() bool
	ContentRequest() messages.MessageInterface
	ProcessResponse(messages.MessageInterface) error
	Update()
	Tree() *Tree
}

func NewAsync(self AsyncInterface) *Async {
	a := &Async{}
	a.self = self
	return a
}

type Async struct {
	self   AsyncInterface
	loaded bool
}

func (a *Async) Loaded() bool {
	return a.loaded
}

func (a *Async) LoadContent() chan bool {

	successChannel := make(chan bool, 1)

	if a.loaded {
		// if the data is already loaded, we return the success channel
		// and immediately signal it.
		successChannel <- true
		return successChannel
	}

	responseChannel := a.self.Tree().Conn.Request(a.self.ContentRequest())

	go func() {

		m := <-responseChannel

		if err := a.self.ProcessResponse(m); err != nil {
			a.self.Tree().Fail <- kerr.New("DLTCGMSREX", err, "awaitSourceResponse")
		}

		a.loaded = true

		a.self.Update()

		successChannel <- true

	}()

	return successChannel

}
