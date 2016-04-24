package flux

import (
	"testing"

	"kego.io/kerr/assert"
)

type st struct {
	*Store
}

func (s *st) Handle(payload *Payload) (finished bool) { return true }

func TestStore_Init(t *testing.T) {
	s := &st{Store: &Store{}}
	s.Init(s)
	assert.NotNil(t, s.self)
	assert.Equal(t, s.self, s)
}

func TestStore_Notify(t *testing.T) {
	s := &st{Store: &Store{}}
	done := make(chan struct{}, 1)

	// subscribers should be nil, so will return 0
	count := s.Notify("a", "b")
	assert.Equal(t, 0, count)

	c := s.Watch("a", "b")
	go func() {
		<-c
		close(done)
	}()

	// notificationType doesn't exist so will return 0
	count = s.Notify("b", "c")
	assert.Equal(t, 0, count)

	// watched object doesn't exist so will return 0
	count = s.Notify("a", "c")
	assert.Equal(t, 0, count)

	// should notify
	count = s.Notify("a", "b")
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify1")

	done = make(chan struct{}, 1)
	go func() {
		<-c
		close(done)
	}()
	count = s.Notify("a")
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify2")

	d := s.Watch("e")
	done = make(chan struct{}, 1)
	go func() {
		<-d
		close(done)
	}()
	count = s.Notify("e", "a")
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify3")
}

func TestStore_Watch(t *testing.T) {
	s := &st{Store: &Store{}}
	c := s.Watch("a", "b")
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers["a"]))
	assert.Equal(t, 1, len(s.subscribers["a"][all_subscribers]))
	assert.Equal(t, 1, len(s.subscribers["a"]["b"]))
	assert.Equal(t, c, s.subscribers["a"][all_subscribers][0])
	assert.Equal(t, c, s.subscribers["a"]["b"][0])

	e := s.Watch("a", "d")
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 3, len(s.subscribers["a"]))
	assert.Equal(t, 2, len(s.subscribers["a"][all_subscribers]))
	assert.Equal(t, 1, len(s.subscribers["a"]["b"]))
	assert.Equal(t, 1, len(s.subscribers["a"]["d"]))
	assert.Equal(t, c, s.subscribers["a"][all_subscribers][0])
	assert.Equal(t, e, s.subscribers["a"][all_subscribers][1])
	assert.Equal(t, c, s.subscribers["a"]["b"][0])
	assert.Equal(t, e, s.subscribers["a"]["d"][0])

	g := s.Watch("a", "b")
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 3, len(s.subscribers["a"]))
	assert.Equal(t, 3, len(s.subscribers["a"][all_subscribers]))
	assert.Equal(t, 2, len(s.subscribers["a"]["b"]))
	assert.Equal(t, 1, len(s.subscribers["a"]["d"]))
	assert.Equal(t, c, s.subscribers["a"][all_subscribers][0])
	assert.Equal(t, e, s.subscribers["a"][all_subscribers][1])
	assert.Equal(t, g, s.subscribers["a"][all_subscribers][2])
	assert.Equal(t, c, s.subscribers["a"]["b"][0])
	assert.Equal(t, g, s.subscribers["a"]["b"][1])
	assert.Equal(t, e, s.subscribers["a"]["d"][0])

	j := s.Watch("i", "h")
	assert.Equal(t, 2, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers["i"]))
	assert.Equal(t, 1, len(s.subscribers["i"][all_subscribers]))
	assert.Equal(t, 1, len(s.subscribers["i"]["h"]))
	assert.Equal(t, j, s.subscribers["i"][all_subscribers][0])
	assert.Equal(t, j, s.subscribers["i"]["h"][0])

	k := s.Watch("l")
	assert.Equal(t, 3, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers["l"]))
	assert.Equal(t, 1, len(s.subscribers["l"][all_subscribers]))
	assert.Equal(t, 1, len(s.subscribers["l"][all_values]))
	assert.Equal(t, k, s.subscribers["l"][all_subscribers][0])
	assert.Equal(t, k, s.subscribers["l"][all_values][0])
}

func TestStore_Delete(t *testing.T) {
	s := &st{Store: &Store{}}
	z := make(chan struct{})
	s.Delete(z)
	c := s.Watch("a", "b")
	e := s.Watch("a", "c")
	g := s.Watch("a", "b")
	j := s.Watch("i", "h")
	assert.Equal(t, 2, len(s.subscribers))
	assert.Equal(t, c, s.subscribers["a"]["b"][0])
	assert.Equal(t, e, s.subscribers["a"]["c"][0])
	assert.Equal(t, g, s.subscribers["a"]["b"][1])
	assert.Equal(t, j, s.subscribers["i"]["h"][0])
	assert.Equal(t, 3, len(s.subscribers["a"]))
	s.Delete(c)
	assert.Equal(t, 3, len(s.subscribers["a"]))
	assert.Equal(t, g, s.subscribers["a"]["b"][0])
	s.Delete(g)
	assert.Equal(t, 2, len(s.subscribers["a"]))
	assert.Equal(t, 2, len(s.subscribers))
	s.Delete(e)
	assert.Equal(t, 1, len(s.subscribers))
	s.Delete(j)
	assert.Equal(t, 0, len(s.subscribers))
}
