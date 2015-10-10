package connection

import (
	"io"

	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/system"
)

// Register registers a channel for messages of a specific type. Replies
// are not send to the registered channels.
func (c *Conn) Subscribe(t system.Reference) chan messages.Message {
	ch := make(chan messages.Message)
	c.subs[t] = ch
	return ch
}

func (c *Conn) Receive() error {
	for {

		var i interface{}
		// Here we unmarshal with messages.Path and messages.Alias because this
		// is just a message, not an object from the local repo.
		if err := ke.NewDecoder(c.socket, messages.Path, messages.Aliases).Decode(&i); err != nil {
			if err == io.EOF {
				// Closing the fail channel exits the app gracefully
				if c.debug {
					c.fail <- kerr.New("BFWXLLOSHQ", nil, "Connection closed")
					return nil
				}
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
				return kerr.New("SRXAVLUFGW", nil, "Response received, but request %s not found", requestGuid)
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
