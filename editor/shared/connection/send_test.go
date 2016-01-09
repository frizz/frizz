package connection

import (
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"kego.io/editor/shared/connection/mocks"
	"kego.io/editor/shared/messages"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
	"kego.io/system"
)

func doTest(t *testing.T, f func(*Conn, *mocks.MockReadWriteCloser)) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	socket := mocks.NewMockReadWriteCloser(ctl)
	fail := make(chan error)

	c := New(tests.Context("kego.io/editor/shared/messages").Ctx(), socket, fail, false)

	f(c, socket)

	socket.EXPECT().Close()
	c.Close()

	c.outwg.Wait()
}

func TestSend(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewDataRequest("a.b/c", "d", "e")

		expected, _ := ke.Marshal(m)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		c.Send(m)

	})
}

func TestRequest(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewDataRequest("a.b/c", "d", "e")

		expected, _ := ke.Marshal(m)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		_ = c.Request(m)

	})
}

func TestRequestResponseChannel(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewDataRequest("a.b/c", "d", "e")

		expected, _ := ke.Marshal(m)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		responseChannel := c.sendRequestAndReturnResponseChannel(m)

		// sendRequestAndWaitForResponse is the first part of the Request method,
		// which sends a request and stores the response channel in the requests
		// map. We should check the response channel exists in the requests map.
		r, ok := c.requests[m.GetMessage(nil).Guid.Value()]
		assert.True(t, ok)
		assert.Equal(t, r, responseChannel)

	})
}

func TestRespond(t *testing.T) {
	doTest(t, func(c *Conn, socket *mocks.MockReadWriteCloser) {

		m := messages.NewDataResponse("a.b/c", "d", true, "f")

		// The Respond method adds the guid of the request message to the "Request"
		// field, so we must clone the request and add it.
		m1, err := clone(tests.Context("kego.io/editor/shared/messages").Jauto().Ctx(), m)
		assert.NoError(t, err)
		m1.(messages.MessageInterface).GetMessage(nil).Request = system.NewString("c")
		expected, _ := ke.Marshal(m1)
		expected = append(expected, byte('\n'))

		socket.EXPECT().Write(expected)
		c.Respond(m, "c")

	})
}

func clone(ctx context.Context, in system.ObjectInterface) (system.ObjectInterface, error) {
	b, err := ke.Marshal(in)
	if err != nil {
		return nil, kerr.New("WSTYPJHIVG", err, "ke.Marshal")
	}
	var i interface{}
	err = ke.Unmarshal(ctx, b, &i)
	if err != nil {
		return nil, kerr.New("NGELVDDBRF", err, "ke.Unmarshal")
	}
	o, ok := i.(system.ObjectInterface)
	if !ok {
		return nil, kerr.New("HLMRUOUQEM", nil, "Unmarshaled object is not a system.ObjectInterface")
	}
	return o, nil
}
