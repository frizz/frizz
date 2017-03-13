package flux

import (
	"testing"

	"github.com/dave/ktest/assert"
)

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

func TestNotifier_Notify(t *testing.T) {
	n := &Notifier{}
	done := make(chan struct{}, 1)

	// subscribers should be nil, so will return 0
	finished := n.Notify("a", bNotif)
	assert.WaitFor(t, finished, false)

	// subscribers should be nil, so will return 0
	count, finished := n.notifyCount("a", bNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 0, count)

	c := n.Watch("a", bNotif)
	go func() {
		notif := <-c
		close(notif.Done)
		close(done)
	}()

	// notif doesn't exist so will return 0
	count, finished = n.notifyCount("a", cNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 0, count)

	// object doesn't exist so will return 0
	count, finished = n.notifyCount("b", bNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 0, count)

	// should notify
	count, finished = n.notifyCount("a", bNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 1, count)
	assert.WaitFor(t, done, false, "Timed out waiting for notify1")

	done = make(chan struct{}, 1)
	go func() {
		notif := <-c
		s, ok := notif.Data.(string)
		assert.True(t, ok)
		assert.Equal(t, "h", s)
		close(notif.Done)
		close(done)
	}()
	count, finished = n.notifyCount(nil, bNotif, "h")
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 1, count)
	assert.WaitFor(t, done, false, "Timed out waiting for notify2")

	d := n.Watch(nil, eNotif)
	done = make(chan struct{}, 1)
	go func() {
		notif := <-d
		close(notif.Done)
		close(done)
	}()
	count, finished = n.notifyCount("f", eNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 1, count)
	assert.WaitFor(t, done, false, "Timed out waiting for notify3")

	g := n.Watch("h", iNotif, jNotif)
	done = make(chan struct{}, 1)
	go func() {
		notif1 := <-g
		close(notif1.Done)
		notif2 := <-g
		close(notif2.Done)
		close(done)
	}()
	count, finished = n.notifyCount("h", jNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 1, count)
	count, finished = n.notifyCount("h", iNotif, nil)
	assert.WaitFor(t, finished, false)
	assert.Equal(t, 1, count)
	assert.WaitFor(t, done, false, "Timed out waiting for notify4")

	done = make(chan struct{}, 1)
	go func() {
		notif := <-c
		s, ok := notif.Data.(string)
		assert.True(t, ok)
		assert.Equal(t, "i", s)
		close(notif.Done)
		close(done)
	}()
	finished = n.NotifyWithData(nil, bNotif, "i")
	assert.WaitFor(t, finished, false)
	assert.WaitFor(t, done, false, "Timed out waiting for notify5")
}

func TestNotifier_Watch(t *testing.T) {
	n := NewNotifier()
	c := n.Watch("a", bNotif)
	assert.Equal(t, 1, len(n.subscribers))
	assert.Equal(t, 2, len(n.subscribers[bNotif]))
	assert.Equal(t, 1, len(n.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 1, len(n.subscribers[bNotif]["a"]))
	assert.Equal(t, c, n.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, c, n.subscribers[bNotif]["a"][0])

	e := n.Watch("d", bNotif)
	assert.Equal(t, 1, len(n.subscribers))
	assert.Equal(t, 3, len(n.subscribers[bNotif]))
	assert.Equal(t, 2, len(n.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 1, len(n.subscribers[bNotif]["a"]))
	assert.Equal(t, 1, len(n.subscribers[bNotif]["d"]))
	assert.Equal(t, c, n.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, e, n.subscribers[bNotif][allSubscribers][1])
	assert.Equal(t, c, n.subscribers[bNotif]["a"][0])
	assert.Equal(t, e, n.subscribers[bNotif]["d"][0])

	g := n.Watch("a", bNotif)
	assert.Equal(t, 1, len(n.subscribers))
	assert.Equal(t, 3, len(n.subscribers[bNotif]))
	assert.Equal(t, 3, len(n.subscribers[bNotif][allSubscribers]))
	assert.Equal(t, 2, len(n.subscribers[bNotif]["a"]))
	assert.Equal(t, 1, len(n.subscribers[bNotif]["d"]))
	assert.Equal(t, c, n.subscribers[bNotif][allSubscribers][0])
	assert.Equal(t, e, n.subscribers[bNotif][allSubscribers][1])
	assert.Equal(t, g, n.subscribers[bNotif][allSubscribers][2])
	assert.Equal(t, c, n.subscribers[bNotif]["a"][0])
	assert.Equal(t, g, n.subscribers[bNotif]["a"][1])
	assert.Equal(t, e, n.subscribers[bNotif]["d"][0])

	j := n.Watch("i", hNotif)
	assert.Equal(t, 2, len(n.subscribers))
	assert.Equal(t, 2, len(n.subscribers[hNotif]))
	assert.Equal(t, 1, len(n.subscribers[hNotif][allSubscribers]))
	assert.Equal(t, 1, len(n.subscribers[hNotif]["i"]))
	assert.Equal(t, j, n.subscribers[hNotif][allSubscribers][0])
	assert.Equal(t, j, n.subscribers[hNotif]["i"][0])

	k := n.Watch(nil, iNotif)
	assert.Equal(t, 3, len(n.subscribers))
	assert.Equal(t, 2, len(n.subscribers[iNotif]))
	assert.Equal(t, 1, len(n.subscribers[iNotif][allSubscribers]))
	assert.Equal(t, 1, len(n.subscribers[iNotif][allObjects]))
	assert.Equal(t, k, n.subscribers[iNotif][allSubscribers][0])
	assert.Equal(t, k, n.subscribers[iNotif][allObjects][0])
}

func TestNotifier_Delete(t *testing.T) {
	n := &Notifier{}
	z := make(chan NotifPayload)
	n.Delete(z)
	c := n.Watch("b", aNotif)
	e := n.Watch("c", aNotif)
	g := n.Watch("b", aNotif)
	j := n.Watch("h", iNotif)
	assert.Equal(t, 2, len(n.subscribers))
	assert.Equal(t, c, n.subscribers[aNotif]["b"][0])
	assert.Equal(t, e, n.subscribers[aNotif]["c"][0])
	assert.Equal(t, g, n.subscribers[aNotif]["b"][1])
	assert.Equal(t, j, n.subscribers[iNotif]["h"][0])
	assert.Equal(t, 3, len(n.subscribers[aNotif]))
	n.Delete(c)
	assert.Equal(t, 3, len(n.subscribers[aNotif]))
	assert.Equal(t, g, n.subscribers[aNotif]["b"][0])
	n.Delete(g)
	assert.Equal(t, 2, len(n.subscribers[aNotif]))
	assert.Equal(t, 2, len(n.subscribers))
	n.Delete(e)
	assert.Equal(t, 1, len(n.subscribers))
	n.Delete(j)
	assert.Equal(t, 0, len(n.subscribers))
}
