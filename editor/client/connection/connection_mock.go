package connection

import "net/rpc"

type Mock struct {
	// Reply - if set, calls to Go will immediatly return this reply
	Reply interface{}

	// Log - a log of all calls to Go
	Log []struct {
		serviceMethod string
		args          interface{}
		reply         interface{}
		done          chan *rpc.Call
		fail          chan error
	}

	// Closed is true if Closed has been called
	Closed bool
}

func (m *Mock) Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call, fail chan error) *rpc.Call {
	m.Log = append(m.Log, struct {
		serviceMethod string
		args          interface{}
		reply         interface{}
		done          chan *rpc.Call
		fail          chan error
	}{serviceMethod, args, reply, done, fail})

	if m.Reply == nil {
		return &rpc.Call{}
	}

	call := &rpc.Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         m.Reply,
		Error:         nil,
		Done:          done,
	}

	done <- call

	return call

}

func (m *Mock) Close() {
	m.Closed = true
}
