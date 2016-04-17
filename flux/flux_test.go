package flux

import (
	"testing"

	"time"

	"kego.io/kerr/assert"
)

func TestWatchMulti(t *testing.T) {
	done := make(chan struct{})
	go func() {
		// We watch two channels, a and b
		a := make(chan struct{}, 1)
		b := make(chan struct{}, 1)
		c := WatchMulti(a, b)

		// We send a signal to a and wait for the response on the
		a <- struct{}{}
		waitFor(t, c, true, "A")

		// We send a signal to a and b and ensure two signals are returned
		a <- struct{}{}
		b <- struct{}{}
		waitFor(t, c, true, "B")
		waitFor(t, c, true, "C")

		// We close a and send on b to ensure it still works
		close(a)
		b <- struct{}{}
		waitFor(t, c, true, "D")

		// We close the second input channel, and check that the signal channel is also closed
		close(b)
		waitFor(t, c, false, "E")

		close(done)
	}()
	waitFor(t, done, false, "G")
}
func waitFor(t *testing.T, c chan struct{}, shouldBeOpen bool, description string) {
	select {
	case _, open := <-c:
		if open != shouldBeOpen {
			assert.Fail(t, description+" (not in correct state) ")
		}
		return
	case <-time.After(time.Millisecond * 50):
		assert.Fail(t, description)
	}
}
