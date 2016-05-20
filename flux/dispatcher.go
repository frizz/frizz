package flux

import "sync"

type DispatcherInterface interface {
	Dispatch(action ActionInterface) chan struct{}
}

type Dispatcher struct {
	m      sync.Mutex
	stores []StoreInterface
}

// NewDispatcher creates a new dispatcher and registers the provided stores
func NewDispatcher(stores ...StoreInterface) *Dispatcher {
	d := &Dispatcher{}
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

		d.m.Lock()
		defer d.m.Unlock()

		// Create a waitgroup and add the number of stores
		wg := sync.WaitGroup{}
		wg.Add(len(d.stores))

		// Create the loop detector
		loop := newLoopDetector()

		payloads := make(map[StoreInterface]*Payload, len(d.stores))
		for _, store := range d.stores {
			payloads[store] = newPayload(action, store, payloads, loop)
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
		close(done)
	}()
	return done
}
