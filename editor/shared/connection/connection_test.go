package connection

import (
	"testing"

	"time"

	"kego.io/editor/shared/messages"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestTimeout(t *testing.T) {

	input, outcome := testTimeout(t)

	time.Sleep(time.Millisecond * 100)
	input <- messages.NewSourceRequest("a")

	err := <-outcome
	assert.HasError(t, err, "QKTKOKWSDG")

}

func TestNoTimeout(t *testing.T) {

	input, outcome := testTimeout(t)

	time.Sleep(time.Millisecond * 10)
	input <- messages.NewSourceRequest("a")

	err := <-outcome
	assert.NoError(t, err)

}

func testTimeout(t *testing.T) (chan messages.MessageInterface, chan error) {
	fail := make(chan error)
	c := New(tests.PathCtx("kego.io/editor/shared/messages"), nil, fail, false)

	outcome := make(chan error, 1)

	input := make(chan messages.MessageInterface, 1)
	output := c.applyTimeout(time.Millisecond*50, input)

	go func() {
		select {
		case err := <-c.fail:
			outcome <- err
		case <-output:
			outcome <- nil
		}
	}()

	return input, outcome
}
