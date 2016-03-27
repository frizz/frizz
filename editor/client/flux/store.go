package flux

import "sync"

type StoreInterface interface {
	// Handle consumes the action in the payload. If the operation was performed syncronously,
	// return true. If not, return false and close the payload.Done channel when finished.
	Handle(payload *Payload) (finished bool)

	// Change returns a channel that is signalled when there's a change to this
	Changed(watch ...interface{}) chan struct{}

	// notify is use by this store to send the change signal to all change chanels
	Notify(changed ...interface{})

	Delete(chan struct{})
}

type Store struct {
	m           sync.Mutex
	subscribers map[interface{}][]chan struct{}
}

type key string

const all_subscribers key = "all_subscribers"
const all_values key = "all_values"

func (s *Store) Delete(c chan struct{}) {
	s.m.Lock()
	defer s.m.Unlock()
	for v, a := range s.subscribers {
		for i, ch := range a {
			if c == ch {
				if len(s.subscribers[v]) == 1 {
					delete(s.subscribers, v)
					break
				}
				s.subscribers[v] = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
}

func (s *Store) Changed(watch ...interface{}) chan struct{} {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		s.subscribers = map[interface{}][]chan struct{}{}
	}
	c := make(chan struct{})
	// all_subscribers contains all the subscribers reguardless what they are watching.
	s.subscribers[all_subscribers] = append(s.subscribers[all_subscribers], c)
	if len(watch) == 0 {
		// all_values contains only the subscribers that are watching all values.
		s.subscribers[all_values] = append(s.subscribers[all_values], c)
	}
	for _, v := range watch {
		s.subscribers[v] = append(s.subscribers[v], c)
	}
	return c
}

func (s *Store) Notify(changed ...interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		return
	}
	matching := map[chan struct{}]bool{}
	if len(changed) == 0 {
		// all_subscribers contains all the subscribers reguardless what they are watching.
		for _, c := range s.subscribers[all_subscribers] {
			matching[c] = true
		}
	} else {
		// all_values contains only the subscribers that are watching all values.
		for _, c := range s.subscribers[all_values] {
			matching[c] = true
		}
		for _, v := range changed {
			for _, c := range s.subscribers[v] {
				matching[c] = true
			}
		}
	}

	for c, _ := range matching {
		c <- struct{}{}
	}

}
