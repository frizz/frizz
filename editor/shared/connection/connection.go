//go:generate mockgen -destination mocks/io.go -package mocks io ReadWriteCloser
//go:generate mockgen -destination mocks/messages.go -package mocks kego.io/editor/shared/messages Message
package connection // import "kego.io/editor/shared/connection"

import (
	"io"
	"time"

	"sync"

	"golang.org/x/net/context"
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
	ctx      context.Context
	in       chan messages.MessageInterface
	out      chan messages.MessageInterface
	requests map[string]chan messages.MessageInterface
	subs     map[system.Reference]chan messages.MessageInterface

	// This wait group is only used in the tests to signal when all the
	// outgoing messages have finished sending.
	outwg *sync.WaitGroup
	debug bool // in debug mode we don't exit the server on connection close
}

func New(ctx context.Context, socket io.ReadWriteCloser, fail chan error, debug bool) *Conn {
	c := &Conn{
		socket:   socket,
		fail:     fail,
		ctx:      ctx,
		in:       make(chan messages.MessageInterface),
		out:      make(chan messages.MessageInterface),
		requests: map[string]chan messages.MessageInterface{},
		subs:     map[system.Reference]chan messages.MessageInterface{},
		outwg:    &sync.WaitGroup{},
		debug:    debug,
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

func (c *Conn) applyTimeout(duration time.Duration, input chan messages.MessageInterface) chan messages.MessageInterface {
	output := make(chan messages.MessageInterface)
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
