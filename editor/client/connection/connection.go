package connection

import (
	"net/rpc"
	"time"

	"frizz.io/editor/client/common"
	"frizz.io/editor/shared"
	"github.com/dave/kerr"
)

type Interface interface {
	Go(method shared.Method, args interface{}, reply interface{}, fail chan error) chan error
	Close()
}

type Conn struct {
	client *rpc.Client
}

func New(client *rpc.Client) *Conn {
	return &Conn{
		client: client,
	}
}

func (c *Conn) Go(method shared.Method, args interface{}, reply interface{}, fail chan error) chan error {
	done := make(chan error, 1)
	call := c.client.Go(string(method), args, reply, make(chan *rpc.Call, 1))

	go func() {
		select {
		case <-call.Done:
			if call.Error != nil {
				done <- call.Error
			} else {
				close(done)
			}
		case <-time.After(common.ClientConnectionTimeout):
			fail <- kerr.New("CWOTFNPITL", "Timeout")
		}
	}()

	return done

}

func (c *Conn) Close() {
	c.client.Close()
}
