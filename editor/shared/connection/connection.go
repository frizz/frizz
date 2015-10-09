//go:generate mockgen -destination mocks/io.go -package mocks io ReadWriteCloser
//go:generate mockgen -destination mocks/messages.go -package mocks kego.io/editor/shared/messages Message
package connection // import "kego.io/editor/shared/connection"

import (
	"io"
	"time"

	"sync"

	"kego.io/editor/shared/messages"
	"kego.io/kerr"
	"kego.io/system"
)

// Messages are all localhost so we shouldn't need more than 200ms timeout.
var TIMEOUT = time.Millisecond * 200

// Pass messages as strings makes debugging easier. Perhaps binary in production?
var MESSAGE_TYPE = M_STRING

type Conn struct {
	socket   io.ReadWriteCloser
	fail     chan error
	path     string
	aliases  map[string]string
	in       chan messages.Message
	out      chan messages.Message
	requests map[string]chan messages.Message
	subs     map[system.Reference]chan messages.Message

	// This wait group is only used in the tests to signal when all the
	// outgoing messages have finished sending.
	outwg *sync.WaitGroup
}

func New(socket io.ReadWriteCloser, fail chan error, path string, aliases map[string]string) *Conn {
	c := &Conn{
		socket:   socket,
		fail:     fail,
		path:     path,
		aliases:  aliases,
		in:       make(chan messages.Message),
		out:      make(chan messages.Message),
		requests: map[string]chan messages.Message{},
		subs:     map[system.Reference]chan messages.Message{},
		outwg:    &sync.WaitGroup{},
	}
	go c.handle(c.sender)
	return c
}

func (c *Conn) Close() error {
	return c.socket.Close()
}

func (c *Conn) handle(f func() error) {
	if err := f(); err != nil {
		c.fail <- err
	}
}

func (c *Conn) applyTimeout(duration time.Duration, input chan messages.Message) chan messages.Message {
	output := make(chan messages.Message)
	go func() {
		select {
		case <-time.After(duration):
			c.fail <- kerr.New("QKTKOKWSDG", nil, "Timed out waiting for message")
			return
		case m := <-input:
			output <- m
			return
		}
	}()
	return output
}

type MessageType string

const (
	M_STRING MessageType = "string"
	M_BINARY             = "binary"
)
