package flux

import "sync"

type StoreInterface interface {
	// Handle consumes the action in the payload. If the operation was performed syncronously,
	// return true. If not, return false and close the payload.Done channel when finished.
	Handle(payload *Payload) (finished bool)

	// Change returns a channel that is signalled when there's a change to this
	Changed() chan struct{}

	// notify is use by this store to send the change signal to all change chanels
	Notify()
}

type Store struct {
	m           sync.Mutex
	subscribers []chan struct{}
}

func (s *Store) Changed() chan struct{} {
	s.m.Lock()
	defer s.m.Unlock()
	c := make(chan struct{})
	s.subscribers = append(s.subscribers, c)
	return c
}

func (s *Store) Notify() {
	s.m.Lock()
	defer s.m.Unlock()
	for _, sub := range s.subscribers {
		sub <- struct{}{}
	}
}
