package flux

import "sync"

type StoreInterface interface {
	// Handle consumes the action in the payload. If the operation was performed syncronously,
	// return true. If not, return false and close the payload.Done channel when finished.
	Handle(payload *Payload) (finished bool)

	// Change returns a channel that is signalled when there's a change to this
	WatchAll(watch ...interface{}) chan struct{}

	WatchSingle(notificationType interface{}, watch ...interface{}) chan struct{}

	// notify is use by this store to send the change signal to all change chanels
	NotifyAll(changed ...interface{})

	NotifySingle(notificationType interface{}, changed ...interface{})

	DeleteAll(c chan struct{})

	DeleteSingle(notificationType interface{}, c chan struct{})
}

type Store struct {
	m           sync.Mutex
	self        StoreInterface
	subscribers map[interface{}]notifier
}

func (s *Store) Init(si StoreInterface) {
	s.self = si
}

type key string

const all_notifications key = "all_notifications"
const all_subscribers key = "all_subscribers"
const all_values key = "all_values"

func (s *Store) DeleteAll(c chan struct{}) {
	s.self.DeleteSingle(all_notifications, c)
}
func (s *Store) DeleteSingle(notificationType interface{}, c chan struct{}) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		return
	}
	if s.subscribers[notificationType] == nil {
		return
	}
	n := s.subscribers[notificationType]
	for v, a := range n {
		for i, ch := range a {
			if c == ch {
				if len(n[v]) == 1 {
					delete(n, v)
					break
				}
				n[v] = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
}

func (s *Store) WatchAll(watch ...interface{}) chan struct{} {
	return s.self.WatchSingle(all_notifications, watch...)
}

func (s *Store) WatchSingle(notificationType interface{}, watch ...interface{}) chan struct{} {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		s.subscribers = map[interface{}]notifier{}
	}
	if s.subscribers[notificationType] == nil {
		s.subscribers[notificationType] = notifier{}
	}
	return s.subscribers[notificationType].watch(watch)

}

func (s *Store) NotifyAll(changed ...interface{}) {
	s.self.NotifySingle(all_notifications, changed...)
}

func (s *Store) NotifySingle(notificationType interface{}, changed ...interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		return
	}
	if s.subscribers[notificationType] == nil {
		return
	}
	s.subscribers[notificationType].notify(changed)
}

type notifier map[interface{}][]chan struct{}

func (n notifier) watch(watch []interface{}) chan struct{} {
	c := make(chan struct{})
	// all_subscribers contains all the subscribers reguardless what they are watching.
	n[all_subscribers] = append(n[all_subscribers], c)
	if len(watch) == 0 {
		// all_values contains only the subscribers that are watching all values.
		n[all_values] = append(n[all_values], c)
	}
	for _, v := range watch {
		n[v] = append(n[v], c)
	}
	return c
}

func (n notifier) notify(changed []interface{}) {
	matching := map[chan struct{}]bool{}
	if len(changed) == 0 {
		// all_subscribers contains all the subscribers reguardless what they are watching.
		for _, c := range n[all_subscribers] {
			matching[c] = true
		}
	} else {
		// all_values contains only the subscribers that are watching all values.
		for _, c := range n[all_values] {
			matching[c] = true
		}
		for _, v := range changed {
			for _, c := range n[v] {
				matching[c] = true
			}
		}
	}

	for c, _ := range matching {
		c <- struct{}{}
	}
}
