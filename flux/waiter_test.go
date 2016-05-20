package flux

import "testing"

func TestWaiter_Add(t *testing.T) {
	c := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	w := &Waiter{}
	w.Add(c)
	close(c)
	w.Go(done)
	waitFor(t, done, false, "")

	w = &Waiter{}
	done = make(chan struct{}, 1)
	w.Go(done)
	waitFor(t, done, false, "")
}
