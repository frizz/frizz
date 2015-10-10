package connection

import (
	"testing"

	"time"

	"kego.io/editor/shared/messages"
	"kego.io/kerr/assert"
)

func TestTimeout(t *testing.T) {

	input, outcome := testTimeout(t)

	time.Sleep(time.Millisecond * 15)
	input <- messages.NewSourceRequest("a")

	err := <-outcome
	assert.HasError(t, err, "QKTKOKWSDG")

}

func TestNoTimeout(t *testing.T) {

	input, outcome := testTimeout(t)

	time.Sleep(time.Millisecond * 5)
	input <- messages.NewSourceRequest("a")

	err := <-outcome
	assert.NoError(t, err)

}

func testTimeout(t *testing.T) (chan messages.Message, chan error) {
	fail := make(chan error)
	c := New(nil, fail, false, "kego.io/editor/shared/messages", map[string]string{})

	outcome := make(chan error, 1)

	input := make(chan messages.Message, 1)
	output := c.applyTimeout(time.Millisecond*10, input)

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
