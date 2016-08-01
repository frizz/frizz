package connection

import (
	"net/rpc"
	"time"

	"github.com/davelondon/kerr"
	"kego.io/editor/client/common"
	"kego.io/editor/shared"
)

type Interface interface {
	Go(method shared.Method, args interface{}, reply interface{}, done chan *rpc.Call, fail chan error) *rpc.Call
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

func (c *Conn) Go(method shared.Method, args interface{}, reply interface{}, done chan *rpc.Call, fail chan error) *rpc.Call {
	rpcCall := c.client.Go(string(method), args, reply, make(chan *rpc.Call, 1))

	call := &rpc.Call{
		ServiceMethod: string(method),
		Args:          args,
		Reply:         reply,
		Error:         rpcCall.Error,
		Done:          done,
	}

	go func() {
		select {
		case <-rpcCall.Done:
			done <- call
		case <-time.After(common.ClientConnectionTimeout):
			fail <- kerr.New("CWOTFNPITL", "Timeout")
		}
	}()

	return call

}

func (c *Conn) Close() {
	c.client.Close()
}
