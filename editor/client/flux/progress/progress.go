package progress // import "kego.io/editor/client/flux/progress"

import "sync"

type Info struct {
	Done        chan struct{}
	state       state
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
		// We don't want the notify function to run at the same time as this function, so we use a
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
		// state is already finished, so we close the channel before returning
		close(n)
		return n
	}
	p.subscribers = append(p.subscribers, n)
	return n
}

type state string

const (
	starting state = "starting"
	running  state = "running"
	finished state = "finished"
)
