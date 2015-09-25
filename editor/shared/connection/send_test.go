package connection

import (
	"testing"

	"github.com/golang/mock/gomock"
	"kego.io/editor/shared/connection/mocks"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/system"
)

func doTest(t *testing.T, f func(*Conn, *mocks.MockReadWriteCloser)) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	socket := mocks.NewMockReadWriteCloser(ctl)
	fail := make(chan error)

	c := New(socket, fail, "kego.io/editor/shared/messages", map[string]string{})

	f(c, socket)

	socket.EXPECT().Close()
	c.Close()

	c.outwg.Wait()
}

func TestSend(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewGlobalRequest("a")

		expected, _ := ke.Marshal(m)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		c.Send(m)

	})
}

func TestRequest(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewGlobalRequest("a")

		expected, _ := ke.Marshal(m)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		responseChannel := c.sendRequestAndReturnResponseChannel(m)

		// sendRequestAndWaitForResponse is the first part of the Request method,
		// which sends a request and stores the response channel in the requests
		// map. We should check the response channel exists in the requests map.
		r, ok := c.requests[m.Message().Guid.Value]
		assert.True(t, ok)
		assert.Equal(t, r, responseChannel)

	})
}

func TestRespond(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewGlobalResponse("a", true, "b")

		// The Respond method adds the guid of the request message to the "Request"
		// field, so we must clone the request and add it.
		m1, err := clone(m, "kego.io/editor/shared/messages", map[string]string{})
		assert.NoError(t, err)
		m1.(messages.Message).Message().Request = system.NewString("c")
		expected, _ := ke.Marshal(m1)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		c.Respond(m, "c")

	})
}

func clone(in system.Object, path string, aliases map[string]string) (system.Object, error) {
	b, err := ke.Marshal(in)
	if err != nil {
		return nil, kerr.New("WSTYPJHIVG", err, "ke.Marshal")
	}
	var i interface{}
	err = ke.Unmarshal(b, &i, path, aliases)
	if err != nil {
		return nil, kerr.New("NGELVDDBRF", err, "ke.Unmarshal")
	}
	o, ok := i.(system.Object)
	if !ok {
		return nil, kerr.New("HLMRUOUQEM", nil, "Unmarshaled object is not a system.Object")
	}
	return o, nil
}
