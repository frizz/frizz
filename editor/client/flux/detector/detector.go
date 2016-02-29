package detector // import "kego.io/editor/client/flux/detector"

import "sync"

type Store interface{}

func New() *detector {
	return &detector{
		waiting: map[Store][]Store{},
	}
}

type detector struct {
	m       sync.RWMutex
	waiting map[Store][]Store
}

func (w *detector) FinishedWait(store Store) {
	w.m.Lock()
	defer w.m.Unlock()
	if _, ok := w.waiting[store]; ok {
		delete(w.waiting, store)
	}
}

func (w *detector) RequestWait(store Store, waitFor ...Store) (loopFound bool, loopStore Store) {
	// returns true if s1 is waiting for s2
	var isWaiting func(s1, s2 Store) bool
	isWaiting = func(s1, s2 Store) bool {
		w.m.RLock()
		defer w.m.RUnlock()
		waits, ok := w.waiting[s1]
		if !ok {
			return false
		}
		for _, inner := range waits {
			if inner == s2 {
				return true
			}
			if isWaiting(inner, s2) {
				return true
			}
		}
		return false
	}
	for _, requested := range waitFor {
		if isWaiting(requested, store) {
			return true, requested
		}
	}
	w.m.Lock()
	defer w.m.Unlock()
	w.waiting[store] = waitFor
	return false, nil
}
