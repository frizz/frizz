package connection // import "kego.io/editor/shared/connection"

import (
	"io"
	"time"

	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system"
)

const TIMEOUT = time.Second
const MESSAGE_TYPE = M_STRING

type MessageType string

const (
	M_STRING MessageType = "string"
	M_BINARY             = "binary"
)

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
	}
	go c.handle(c.sender)
	return c
}

type Conn struct {
	socket   io.ReadWriteCloser
	fail     chan error
	path     string
	aliases  map[string]string
	in       chan messages.Message
	out      chan messages.Message
	requests map[string]chan messages.Message
	subs     map[system.Reference]chan messages.Message
}

func (c *Conn) Close() error {
	return c.socket.Close()
}

// Register registers a chanel for messages of a specific type. Replies
// are not send to the registered channels.
func (c *Conn) Subscribe(t system.Reference) chan messages.Message {
	ch := make(chan messages.Message)
	c.subs[t] = ch
	return ch
}

// Send sends a message to the connection.
func (c *Conn) Send(message messages.Message) {
	c.out <- message
}

// Request sends a message and expects a response
func (c *Conn) Request(message messages.Message, fail chan error) chan messages.Message {

	responseChannel := make(chan messages.Message)
	c.requests[message.Message().Guid.Value] = responseChannel
	c.out <- message

	outputChannel := make(chan messages.Message)
	go func() {
		select {
		case <-time.After(TIMEOUT):
			fail <- kerr.New("QKTKOKWSDG", nil, "Timeout when waiting for reply")
			return
		case m := <-responseChannel:
			outputChannel <- m
			return
		}
	}()

	return outputChannel
}

// Respond sends a message as a reply to a request
func (c *Conn) Respond(message messages.Message, requestGuid string) {
	message.Message().Request = system.NewString(requestGuid)
	c.out <- message
}

func (c *Conn) sender() error {
	for {
		m := <-c.out
		if err := ke.NewEncoder(c.socket).Encode(m); err != nil {
			if err == io.EOF {
				// Closing the fail channel exits the app gracefully
				close(c.fail)
				return nil
			}
			return kerr.New("WIUXNWRXCQ", err, "Encode")
		}
	}
}

func (c *Conn) Receive() error {
	for {

		var i interface{}
		// Here we unmarshal with messages.Path and messages.Alias because this
		// is just a message, not an object from the local repo.
		if err := ke.NewDecoder(c.socket, messages.Path, messages.Aliases).Decode(&i); err != nil {
			if err == io.EOF {
				// Closing the fail channel exits the app gracefully
				close(c.fail)
				return nil
			}
			return kerr.New("GJTKHMTWYY", err, "ke.Unmarshal")
		}

		m, ok := i.(messages.Message)
		if !ok {
			return kerr.New("FYWKSPGYXY", nil, "Received type %T does not implement messages.Message", i)
		}

		requestGuid := m.Message().Request
		if requestGuid.Exists {
			reply, ok := c.requests[requestGuid.Value]
			if !ok {
				return kerr.New("SRXAVLUFGW", nil, "Request %s not found", requestGuid)
			}
			reply <- m
			delete(c.requests, requestGuid.Value)
			continue
		}

		t := m.(system.Object).Object().Type
		subscriber, ok := c.subs[t]
		if !ok {
			return kerr.New("USAYKUNHSC", nil, "Subscriber not found for %s", t.Value())
		}
		subscriber <- m
	}
}

func (c *Conn) handle(f func() error) {
	if err := f(); err != nil {
		c.fail <- err
	}
}
