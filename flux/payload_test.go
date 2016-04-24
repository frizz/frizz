package flux

import (
	"testing"

	"sync"

	"kego.io/kerr/assert"
)

func TestPayload_Wait(t *testing.T) {

}

func TestNewPayload(t *testing.T) {
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
		waitForGroup(t, wg2, "A")
		a = true
		wg4.Done()
	}()
	b := false
	go func() {
		wg1.Done()
		<-p.finished()
		waitForGroup(t, wg3, "B")
		b = true
		wg4.Done()
	}()

	waitForGroup(t, wg1, "C")

	close(p.Done)

	wg2.Done()
	wg3.Done()

	waitForGroup(t, wg4, "D")

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

	waitForGroup(t, wg5, "E")
	assert.True(t, c)

}
