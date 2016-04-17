package detector // import "kego.io/flux/detector"

// ke: {"package": {"complete": true}}

import "sync"

func New() *detector {
	return &detector{
		waiting: map[interface{}][]interface{}{},
	}
}

type detector struct {
	m       sync.RWMutex
	waiting map[interface{}][]interface{}
}

func (w *detector) FinishedWait(store interface{}) {
	w.m.Lock()
	defer w.m.Unlock()
	if _, ok := w.waiting[store]; ok {
		delete(w.waiting, store)
	}
}

func (w *detector) RequestWait(store interface{}, waitFor ...interface{}) (loopFound bool, loopStore interface{}) {
	// returns true if s1 is waiting for s2
	var isWaiting func(s1, s2 interface{}) bool
	isWaiting = func(s1, s2 interface{}) bool {
		w.m.RLock()
		waits, ok := w.waiting[s1]
		w.m.RUnlock()
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
