package system

import (
	"testing"

	"kego.io/uerr"

	"github.com/stretchr/testify/assert"
)

func TestNewNumber(t *testing.T) {
	n := NewNumber(1.2)
	assert.True(t, n.Exists)
	assert.Equal(t, 1.2, n.Value)
}

func TestNumberUnmarshalJSON(t *testing.T) {

	var n Number

	err := n.UnmarshalJSON([]byte("1.2"), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, n.Exists)
	assert.Equal(t, 1.2, n.Value)

	err = n.UnmarshalJSON([]byte("null"), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, n.Exists)
	assert.Equal(t, 0.0, n.Value)

	err = n.UnmarshalJSON([]byte("foo"), "", map[string]string{})
	uerr.Assert(t, err, "GXNBRBELFA")

}

func TestNumberMarshalJSON(t *testing.T) {

	n := Number{Exists: false, Value: 0.0}
	ba, err := n.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	n = Number{Exists: true, Value: 1.2}
	ba, err = n.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "1.2", string(ba))

	n = Number{Exists: true, Value: 1.0}
	ba, err = n.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "1", string(ba))

}

func TestNumberString(t *testing.T) {

	n := Number{Exists: false, Value: 0.0}
	s := n.String()
	assert.Equal(t, "", s)

	n = Number{Exists: true, Value: 1.2}
	s = n.String()
	assert.Equal(t, "1.2", s)

	n = Number{Exists: true, Value: 1.0}
	s = n.String()
	assert.Equal(t, "1", s)

}
