package flux

import (
	"testing"

	"sync"

	"time"

	"kego.io/kerr/assert"
)

func TestProgress(t *testing.T) {
	st1 := &st{}
	env := map[StoreInterface]*Payload{}
	loop := newLoopDetector()

	env[st1] = newPayload("a", st1, env, loop)

	p := env[st1]
	assert.False(t, p.complete)

	wg1 := &sync.WaitGroup{}
	wg1.Add(2)
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	wg3 := &sync.WaitGroup{}
	wg3.Add(1)
	wg4 := &sync.WaitGroup{}
	wg4.Add(2)
	a := false
	go func() {
		wg1.Done()
		<-p.finished()
		waitWithTimeout(t, wg2)
		a = true
		wg4.Done()
	}()
	b := false
	go func() {
		wg1.Done()
		<-p.finished()
		waitWithTimeout(t, wg3)
		b = true
		wg4.Done()
	}()

	waitWithTimeout(t, wg1)

	close(p.Done)

	wg2.Done()
	wg3.Done()

	waitWithTimeout(t, wg4)

	assert.True(t, p.complete)

	assert.True(t, a)
	assert.True(t, b)

	wg5 := &sync.WaitGroup{}
	wg5.Add(1)
	c := false
	go func() {
		<-p.finished()
		wg5.Done()
		c = true
	}()

	waitWithTimeout(t, wg5)
	assert.True(t, c)

}

func waitWithTimeout(t *testing.T, wg *sync.WaitGroup) {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		// ok!
	case <-time.After(200 * time.Millisecond):
		assert.Fail(t, "Timeout")
	}
}
