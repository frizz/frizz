package flux

import "sync"

func newLoopDetector() *loopDetector {
	return &loopDetector{
		waiting: map[StoreInterface][]StoreInterface{},
	}
}

type loopDetector struct {
	m       sync.RWMutex
	waiting map[StoreInterface][]StoreInterface
}

// finished marks a store as finished processing
func (w *loopDetector) finished(store StoreInterface) {
	w.m.Lock()
	defer w.m.Unlock()
	if _, ok := w.waiting[store]; ok {
		delete(w.waiting, store)
	}
}

// request requests a wait. We return if a loop was found, and if one is found we return the store
// that was waiting for this store.
func (w *loopDetector) request(store StoreInterface, waitFor ...StoreInterface) (loopFound bool, loopStore StoreInterface) {
	// returns true if s1 is waiting for s2
	var isWaiting func(s1, s2 StoreInterface) bool
	isWaiting = func(s1, s2 StoreInterface) bool {
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
