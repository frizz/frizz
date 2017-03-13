package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

func TestWaiter_Add(t *testing.T) {
	c := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	w := &Waiter{}
	w.Add(c)
	close(c)
	w.Go(done)
	assert.WaitFor(t, done, false, "")

	w = &Waiter{}
	done = make(chan struct{}, 1)
	w.Go(done)
	assert.WaitFor(t, done, false, "")
}
