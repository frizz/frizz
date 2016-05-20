package flux

import (
	"testing"

	"kego.io/kerr/assert"
)

type st struct {
	*Store
}

type testNotif string

func (t testNotif) IsNotif() {}

const (
	aNotif testNotif = "aNotif"
	bNotif testNotif = "bNotif"
	cNotif testNotif = "cNotif"
	dNotif testNotif = "dNotif"
	eNotif testNotif = "eNotif"
	fNotif testNotif = "fNotif"
	gNotif testNotif = "gNotif"
	hNotif testNotif = "hNotif"
	iNotif testNotif = "iNotif"
	jNotif testNotif = "jNotif"
)

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
	count, finished := s.Notify("a", bNotif)
	<-finished
	assert.Equal(t, 0, count)

	c := s.Watch("a", bNotif)
	go func() {
		notif := <-c
		close(notif.Done)
		close(done)
	}()

	// notif doesn't exist so will return 0
	count, finished = s.Notify("a", cNotif)
	<-finished
	assert.Equal(t, 0, count)

	// object doesn't exist so will return 0
	count, finished = s.Notify("b", bNotif)
	<-finished
	assert.Equal(t, 0, count)

	// should notify
	count, finished = s.Notify("a", bNotif)
	<-finished
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify1")

	done = make(chan struct{}, 1)
	go func() {
		notif := <-c
		close(notif.Done)
		close(done)
	}()
	count, finished = s.Notify(nil, bNotif)
	<-finished
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify2")

	d := s.Watch(nil, eNotif)
	done = make(chan struct{}, 1)
	go func() {
		notif := <-d
		close(notif.Done)
		close(done)
	}()
	count, finished = s.Notify("f", eNotif)
	<-finished
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify3")

	g := s.Watch("h", iNotif, jNotif)
	done = make(chan struct{}, 1)
	go func() {
		notif1 := <-g
		close(notif1.Done)
		notif2 := <-g
		close(notif2.Done)
		close(done)
	}()
	count, finished = s.Notify("h", jNotif)
	<-finished
	assert.Equal(t, 1, count)
	count, finished = s.Notify("h", iNotif)
	<-finished
	assert.Equal(t, 1, count)
	waitFor(t, done, false, "Timed out waiting for notify4")
}

func TestStore_Watch(t *testing.T) {
	s := &st{Store: &Store{}}
	c := s.Watch("a", bNotif)
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers[bNotif]))
	assert.Equal(t, 1, len(s.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 1, len(s.subscribers[bNotif]["a"]))
	assert.Equal(t, c, s.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, c, s.subscribers[bNotif]["a"][0])

	e := s.Watch("d", bNotif)
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 3, len(s.subscribers[bNotif]))
	assert.Equal(t, 2, len(s.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 1, len(s.subscribers[bNotif]["a"]))
	assert.Equal(t, 1, len(s.subscribers[bNotif]["d"]))
	assert.Equal(t, c, s.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, e, s.subscribers[bNotif][allSubscribers][1])
	assert.Equal(t, c, s.subscribers[bNotif]["a"][0])
	assert.Equal(t, e, s.subscribers[bNotif]["d"][0])

	g := s.Watch("a", bNotif)
	assert.Equal(t, 1, len(s.subscribers))
	assert.Equal(t, 3, len(s.subscribers[bNotif]))
	assert.Equal(t, 3, len(s.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 2, len(s.subscribers[bNotif]["a"]))
	assert.Equal(t, 1, len(s.subscribers[bNotif]["d"]))
	assert.Equal(t, c, s.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, e, s.subscribers[bNotif][allSubscribers][1])
	assert.Equal(t, g, s.subscribers[bNotif][allSubscribers][2])
	assert.Equal(t, c, s.subscribers[bNotif]["a"][0])
	assert.Equal(t, g, s.subscribers[bNotif]["a"][1])
	assert.Equal(t, e, s.subscribers[bNotif]["d"][0])

	j := s.Watch("i", hNotif)
	assert.Equal(t, 2, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers[hNotif]))
	assert.Equal(t, 1, len(s.subscribers[hNotif][allSubscribers]))
	assert.Equal(t, 1, len(s.subscribers[hNotif]["i"]))
	assert.Equal(t, j, s.subscribers[hNotif][allSubscribers][0])
	assert.Equal(t, j, s.subscribers[hNotif]["i"][0])

	k := s.Watch(nil, iNotif)
	assert.Equal(t, 3, len(s.subscribers))
	assert.Equal(t, 2, len(s.subscribers[iNotif]))
	assert.Equal(t, 1, len(s.subscribers[iNotif][allSubscribers]))
	assert.Equal(t, 1, len(s.subscribers[iNotif][allObjects]))
	assert.Equal(t, k, s.subscribers[iNotif][allSubscribers][0])
	assert.Equal(t, k, s.subscribers[iNotif][allObjects][0])
}

func TestStore_Delete(t *testing.T) {
	s := &st{Store: &Store{}}
	z := make(chan NotifPayload)
	s.Delete(z)
	c := s.Watch("b", aNotif)
	e := s.Watch("c", aNotif)
	g := s.Watch("b", aNotif)
	j := s.Watch("h", iNotif)
	assert.Equal(t, 2, len(s.subscribers))
	assert.Equal(t, c, s.subscribers[aNotif]["b"][0])
	assert.Equal(t, e, s.subscribers[aNotif]["c"][0])
	assert.Equal(t, g, s.subscribers[aNotif]["b"][1])
	assert.Equal(t, j, s.subscribers[iNotif]["h"][0])
	assert.Equal(t, 3, len(s.subscribers[aNotif]))
	s.Delete(c)
	assert.Equal(t, 3, len(s.subscribers[aNotif]))
	assert.Equal(t, g, s.subscribers[aNotif]["b"][0])
	s.Delete(g)
	assert.Equal(t, 2, len(s.subscribers[aNotif]))
	assert.Equal(t, 2, len(s.subscribers))
	s.Delete(e)
	assert.Equal(t, 1, len(s.subscribers))
	s.Delete(j)
	assert.Equal(t, 0, len(s.subscribers))
}
