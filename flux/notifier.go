package flux

import "sync"

func NewNotifier() *Notifier {
	n := &Notifier{}
	n.subscribers = map[Notif]notifHelper{}
	return n
}

type Notifier struct {
	m           sync.Mutex
	subscribers map[Notif]notifHelper
}

type NotifPayload struct {
	Type Notif
	Done chan struct{}
	Data interface{}
}

type Notif interface {
	IsNotif()
}

type key string

const (
	allSubscribers key = "allSubscribers"
	allObjects     key = "allObjects"
)

func (s *Notifier) Delete(c chan NotifPayload) {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		return
	}

NextHelper:
	for notif, helper := range s.subscribers {
	NextObject:
		for object, chans := range helper {
			for i, ch := range chans {
				if c == ch {
					if len(chans) == 1 {
						if len(helper) == 1 {
							delete(s.subscribers, notif)
							continue NextHelper
						}
						delete(helper, object)
						continue NextObject
					}
					helper[object] = append(chans[:i], chans[i+1:]...)
					continue NextObject
				}
			}
		}
	}
	close(c)
}

// Watch returns a channel subscribed to reveive notifs notifications about object. If object is
// nil, the subscription is for all objects in the store.
func (s *Notifier) Watch(object interface{}, notifs ...Notif) chan NotifPayload {
	if len(notifs) == 0 {
		panic("You must specify notification types to watch for")
	}
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		s.subscribers = map[Notif]notifHelper{}
	}
	c := make(chan NotifPayload)
	for _, notif := range notifs {
		if s.subscribers[notif] == nil {
			s.subscribers[notif] = notifHelper{}
		}
		s.subscribers[notif].watch(object, c)
	}
	return c
}

// Notify sends the notif notification to all subscribers of that notification for object. If
// object is nil, the notification is sent to all subscribers.
func (s *Notifier) Notify(object interface{}, notif Notif) (done chan struct{}) {
	_, done = s.notifyCount(object, notif, nil)
	return done
}

func (s *Notifier) NotifyWithData(object interface{}, notif Notif, data interface{}) (done chan struct{}) {
	_, done = s.notifyCount(object, notif, data)
	return done
}

func (s *Notifier) notifyCount(object interface{}, notif Notif, data interface{}) (count int, done chan struct{}) {
	if sub := s.getSubscriber(notif); sub != nil {
		return sub.notify(object, notif, data)
	}
	done = make(chan struct{}, 1)
	close(done)
	return 0, done
}

func (s *Notifier) getSubscriber(notif Notif) notifHelper {
	s.m.Lock()
	defer s.m.Unlock()
	if s.subscribers == nil {
		return nil
	}
	if s.subscribers[notif] == nil {
		return nil
	}
	return s.subscribers[notif]
}

type notifHelper map[interface{}][]chan NotifPayload

func (n notifHelper) watch(object interface{}, c chan NotifPayload) {
	// allSubscribers contains all the subscribers reguardless what they are watching.
	n[allSubscribers] = append(n[allSubscribers], c)
	if object == nil {
		// allObjects contains only the subscribers that are watching all objects.
		n[allObjects] = append(n[allObjects], c)
	} else {
		n[object] = append(n[object], c)
	}
}

func (n notifHelper) notify(object interface{}, notif Notif, data interface{}) (count int, done chan struct{}) {
	matching := map[chan NotifPayload]bool{}
	if object == nil {
		// allSubscribers contains all the subscribers reguardless what they are watching.
		for _, c := range n[allSubscribers] {
			matching[c] = true
		}
	} else {
		// allObjects contains only the subscribers that are watching all objects.
		for _, c := range n[allObjects] {
			matching[c] = true
		}
		for _, c := range n[object] {
			matching[c] = true
		}
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(matching))
	for c := range matching {
		notifDone := make(chan struct{}, 1)
		go func() {
			<-notifDone
			wg.Done()
		}()
		c <- NotifPayload{Type: notif, Done: notifDone, Data: data}
	}
	done = make(chan struct{}, 1)
	go func() {
		wg.Wait()
		close(done)
	}()
	return len(matching), done
}
