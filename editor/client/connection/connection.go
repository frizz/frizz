package connection

import (
	"net/rpc"
	"time"

	"kego.io/kerr"
)

type Conn struct {
	client *rpc.Client
}

func New(client *rpc.Client) *Conn {
	return &Conn{
		client: client,
	}
}

func (c *Conn) Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call, fail chan error) *rpc.Call {
	rpcCall := c.client.Go(serviceMethod, args, reply, make(chan *rpc.Call, 1))

	call := &rpc.Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Error:         rpcCall.Error,
		Done:          done,
	}

	go func() {
		select {
		case <-rpcCall.Done:
			done <- call
		case <-time.After(time.Millisecond * 200):
			fail <- kerr.New("CWOTFNPITL", "Timeout")
		}
	}()

	return call

}

func (c *Conn) Close() {
	c.Close()
}
