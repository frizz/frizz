package progress // import "kego.io/editor/client/flux/progress"

// ke: {"package": {"complete": true}}

import "sync"

type Info struct {
	Done        chan struct{}
	finished    bool
	m           sync.Mutex
	subscribers []chan struct{}
}

func New() *Info {
	i := &Info{
		Done: make(chan struct{}),
	}
	go func() {
		<-i.Done
		// We don't want the notify function to run at the same time as this function, so we use a
		// mutex
		i.m.Lock()
		defer i.m.Unlock()
		i.finished = true
		for _, s := range i.subscribers {
			close(s)
		}
	}()
	return i
}

func (i *Info) Notify() chan struct{} {
	// we don't want the done action to run at the same time as this function, so we use a mutex
	i.m.Lock()
	defer i.m.Unlock()
	n := make(chan struct{}, 1)
	if i.finished {
		// state is already finished, so we close the channel before returning
		close(n)
		return n
	}
	i.subscribers = append(i.subscribers, n)
	return n
}
