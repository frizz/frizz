package client

import (
	"io"

	"github.com/gopherjs/websocket"
	"kego.io/editor/shared/connection"
)

// We create a socket type that will implement ReadWriteCloser that
// allows us to control whether the websocket sends string or binary
// messages
type socket struct {
	c  *websocket.Conn
	mt connection.MessageType
}

var _ io.ReadWriteCloser = (*socket)(nil)

func (s *socket) Read(p []byte) (n int, err error) {
	return s.c.Read(p)
}
func (s *socket) Write(p []byte) (n int, err error) {
	if s.mt == connection.M_STRING {
		return s.c.WriteString(string(p))
	}
	return s.c.Write(p)
}
func (s *socket) Close() error {
	return s.c.Close()
}
