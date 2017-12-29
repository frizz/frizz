package flux

import "sync"

type Dispatcher struct {
	m        sync.Mutex
	notifier NotifierInterface
	stores   []StoreInterface
}

// NewDispatcher creates a new dispatcher and registers the provided stores
func NewDispatcher(notifier NotifierInterface, stores ...StoreInterface) *Dispatcher {
	d := &Dispatcher{notifier: notifier}
	for _, s := range stores {
		d.Register(s)
	}
	return d
}

// Register adds a store to the dispatcher
func (d *Dispatcher) Register(store StoreInterface) {
	d.m.Lock()
	defer d.m.Unlock()
	d.stores = append(d.stores, store)
}

// Dispatch sends an action to all registered stores
func (d *Dispatcher) Dispatch(action ActionInterface) chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		payloads := make(map[StoreInterface]*Payload, len(d.stores))
		d.handleAction(action, payloads)
		d.sendNotifs(payloads)
		close(done)
	}()
	return done
}

func (d *Dispatcher) handleAction(action ActionInterface, payloads map[StoreInterface]*Payload) {
	d.m.Lock()
	defer d.m.Unlock()

	// Create a waitgroup and add the number of stores
	wg := sync.WaitGroup{}
	wg.Add(len(d.stores))

	// Create the loop detector
	loop := newLoopDetector()

	for _, store := range d.stores {
		payloads[store] = newPayload(action, store, payloads, loop, d.notifier)
	}

	for _, sr := range d.stores {
		// The store will be used inside a goroutine, so we can't use sr, which is re-used by
		// the range.
		store := sr

		payload := payloads[store]

		go func() {
			// Start the store handler.
			finished := store.Handle(payload)

			// If we finished syncronously, we close the tracker's Done channel. If we are
			// still processing asyncronously, we leave it to be closed when the handler has
			// finished.
			if finished {
				close(payload.Done)
			}
		}()

		go func() {
			// We wait for the tracker to finish
			<-payload.finished()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (d *Dispatcher) sendNotifs(payloads map[StoreInterface]*Payload) {
	// all stores have finished processing, so we can start firing off
	// notifs...
	for _, p := range payloads {
		for _, n := range p.notifs {
			<-d.notifier.Notify(n.object, n.notif, n.data)
		}
	}
}
