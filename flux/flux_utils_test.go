package flux

import (
	"sync"
	"testing"
	"time"

	"kego.io/kerr/assert"
)

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

func waitForGroup(t *testing.T, wg *sync.WaitGroup, description string) {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
	// ok!
	case <-time.After(50 * time.Millisecond):
		assert.Fail(t, description)
	}
}
