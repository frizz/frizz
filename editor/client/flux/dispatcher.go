package flux

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

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
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

	loopDetector := detector.New()

	for _, store := range d.stores {
		info[store] = progress.New()
	}
	for _, storeRange := range d.stores {
		store := storeRange
		p := info[store]
		waitFor := func(stores ...StoreInterface) {
			fmt.Printf("%T is waiting for %T...\n", store, stores[0])
			// First check to see if any of the stores we want to wait for are waiting for this one.
			// That would be a deadlock.
			loopFound, loopStore := loopDetector.RequestWait(store, convertStores(stores))
			if loopFound {
				panic(fmt.Sprintf("%T and %T are waiting for each other.", store, loopStore))
			}
			for _, s := range stores {
				<-info[s].Notify()
			}
			loopDetector.FinishedWait(store)
		}
		go func() {
			p.Start()
			handledSyncronously := store.Handle(&Payload{Action: action, WaitFor: waitFor, Done: p.Done})
			if handledSyncronously {
				close(p.Done)
			}
		}()
		go func() {
			<-p.Notify()
			wg.Done()
		}()
	}
	wg.Wait()
}

type Payload struct {
	Action  ActionInterface
	WaitFor func(stores ...StoreInterface)
	Done    chan struct{}
}

func convertStores(stores []StoreInterface) []detector.Store {
	d := []detector.Store{}
	for _, s := range stores {
		d = append(d, s.(detector.Store))
	}
	return d
}
