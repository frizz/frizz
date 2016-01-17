package connection

import (
	"io"

	"golang.org/x/net/context"

	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
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

	unpackContext := context.Background()
	unpackContext = jsonctx.ManualContext(unpackContext)
	jcache := jsonctx.FromContext(unpackContext)
	jcache.Packages.InitManual("kego.io/json", "kego.io/system", "kego.io/editor/shared/messages")
	jcache.Dummies.InitAuto()
	unpackContext = envctx.Dummy(unpackContext, "kego.io/editor/shared/messages", map[string]string{})

	for {

		var i interface{}
		// Here we unmarshal with the dummy messages.Ctx because this is just a message, not an
		// object from the local repo.
		if err := ke.NewDecoder(unpackContext, c.socket).Decode(&i); err != nil {
			if err == io.EOF {
				// Closing the fail channel exits the app gracefully
				if c.debug {
					c.fail <- kerr.New("BFWXLLOSHQ", "Connection closed")
					return nil
				}
				close(c.fail)
				return nil
			}
			return kerr.Wrap("GJTKHMTWYY", err)
		}

		m, ok := i.(messages.MessageInterface)
		if !ok {
			return kerr.New("FYWKSPGYXY", "Received type %T does not implement messages.MessageInterface", i)
		}

		requestGuid := m.GetMessage(nil).Request
		if requestGuid != nil {
			reply, ok := c.requests[requestGuid.Value()]
			if !ok {
				return kerr.New("SRXAVLUFGW", "Response received, but request %s not found", requestGuid)
			}
			reply <- m
			delete(c.requests, requestGuid.Value())
			continue
		}

		t := m.(system.ObjectInterface).GetObject(nil).Type
		subscriber, ok := c.subs[*t]
		if !ok {
			return kerr.New("USAYKUNHSC", "Subscriber not found for %s", t.Value())
		}
		subscriber <- m
	}
}
