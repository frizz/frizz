package connection

import (
	"net/rpc"
	"testing"

	"kego.io/kerr/assert"
)

func TestNew(t *testing.T) {
	client := &rpc.Client{}
	conn := New(client)
	assert.Equal(t, client, conn.client)
}
