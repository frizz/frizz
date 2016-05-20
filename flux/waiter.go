package flux

import "sync"

// Waiter waits for all signals to fire before signalling a done channel
type Waiter struct {
	m  sync.RWMutex
	wg *sync.WaitGroup
}

// Add a signal to the list
func (w *Waiter) Add(wait chan struct{}) {
	if w.wg == nil {
		w.wg = &sync.WaitGroup{}
	}
	w.m.Lock()
	defer w.m.Unlock()
	w.wg.Add(1)
	go func() {
		<-wait
		w.wg.Done()
	}()
}

// Go starts a gorouting that waits for all siglals before signalling the done channel.
func (w *Waiter) Go(done chan struct{}) {
	if w.wg == nil {
		w.wg = &sync.WaitGroup{}
	}
	go func() {
		w.wg.Wait()
		close(done)
	}()
}
