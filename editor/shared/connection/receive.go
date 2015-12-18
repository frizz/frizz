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
func (c *Conn) Subscribe(t system.Reference) chan messages.MessageInterface {
	ch := make(chan messages.MessageInterface)
	c.subs[t] = ch
	return ch
}

func (c *Conn) Receive() error {
	for {

		var i interface{}
		// Here we unmarshal with the dummy messages.Ctx because this is just a message, not an
		// object from the local repo.
		if err := ke.NewDecoder(messages.Ctx, c.socket).Decode(&i); err != nil {
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

		m, ok := i.(messages.MessageInterface)
		if !ok {
			return kerr.New("FYWKSPGYXY", nil, "Received type %T does not implement messages.MessageInterface", i)
		}

		requestGuid := m.GetMessage().Request
		if requestGuid != nil {
			reply, ok := c.requests[requestGuid.Value()]
			if !ok {
				return kerr.New("SRXAVLUFGW", nil, "Response received, but request %s not found", requestGuid)
			}
			reply <- m
			delete(c.requests, requestGuid.Value())
			continue
		}

		t := m.(system.ObjectInterface).GetObject().Type
		subscriber, ok := c.subs[*t]
		if !ok {
			return kerr.New("USAYKUNHSC", nil, "Subscriber not found for %s", t.Value())
		}
		subscriber <- m
	}
}
