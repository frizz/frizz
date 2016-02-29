package progress // import "kego.io/editor/client/flux/progress"

import (
	"fmt"
	"sync"
)

type Info struct {
	state       state
	Done        chan struct{}
	m           sync.Mutex
	subscribers []chan struct{}
}

func New() *Info {
	p := &Info{
		state: starting,
		Done:  make(chan struct{}),
	}
	go func() {
		<-p.Done
		// we don't want the notify function to run at the same time as this function, so we use a
		// mutex
		p.m.Lock()
		defer p.m.Unlock()
		p.state = finished
		for _, s := range p.subscribers {
			close(s)
		}
	}()
	return p
}

func (i *Info) Start() {
	i.state = running
}

func (p *Info) Notify() chan struct{} {
	// we don't want the done action to run at the same time as this function, so we use a mutex
	p.m.Lock()
	defer p.m.Unlock()
	n := make(chan struct{}, 1)
	if p.state == finished {
		fmt.Println("Notify called... state is already closed, so closing now.")
		close(n)
		return n
	}
	fmt.Println("Notify called... adding sub.")
	p.subscribers = append(p.subscribers, n)
	return n
}

type state string

const (
	starting state = "starting"
	running  state = "running"
	finished state = "finished"
)
