package flux

// ke: {"package": {"complete": true}}

import (
	"fmt"
	"sync"

	"kego.io/editor/client/flux/detector"
	"kego.io/editor/client/flux/progress"
)

type Dispatcher struct {
	storesMutex sync.Mutex
	stores      []StoreInterface
}

func NewDispatcher(stores ...StoreInterface) *Dispatcher {
	d := &Dispatcher{}
	for _, s := range stores {
		d.Register(s)
	}
	return d
}

func (d *Dispatcher) Register(store StoreInterface) {
	d.storesMutex.Lock()
	defer d.storesMutex.Unlock()
	d.stores = append(d.stores, store)
}

func (d *Dispatcher) Dispatch(action ActionInterface) {
	d.storesMutex.Lock()
	defer d.storesMutex.Unlock()

	info := map[StoreInterface]*progress.Info{}
	wg := sync.WaitGroup{}
	wg.Add(len(d.stores))

	loop := detector.New()

	for _, store := range d.stores {
		info[store] = progress.New()
	}
	for _, storeRange := range d.stores {
		store := storeRange
		p := info[store]
		waitFor := func(stores ...StoreInterface) {
			// First check to see if any of the stores we want to wait for are waiting for this one.
			// That would be a deadlock.
			ifaces := make([]interface{}, len(stores))
			for i := range stores {
				ifaces[i] = stores[i]
			}
			if found, loopStore := loop.RequestWait(store, ifaces...); found {
				panic(fmt.Errorf("%T and %T are waiting for each other.", store, loopStore))
			}
			for _, s := range stores {
				<-info[s].Finished()
			}
			loop.FinishedWait(store)
		}
		go func() {
			finished := store.Handle(&Payload{Action: action, WaitFor: waitFor, Done: p.Done})
			if finished {
				close(p.Done)
			}
		}()
		go func() {
			<-p.Finished()
			wg.Done()
		}()
	}
	wg.Wait()
}
