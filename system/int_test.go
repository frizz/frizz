package system

import (
	"testing"

	"kego.io/uerr"

	"github.com/stretchr/testify/assert"
)

func TestNewInt(t *testing.T) {
	n := NewInt(2)
	assert.True(t, n.Exists)
	assert.Equal(t, 2, n.Value)
}

func TestIntUnmarshalJSON(t *testing.T) {

	var i Int

	err := i.UnmarshalJSON([]byte("2"), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, i.Exists)
	assert.Equal(t, 2, i.Value)

	err = i.UnmarshalJSON([]byte("-12"), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, i.Exists)
	assert.Equal(t, -12, i.Value)

	err = i.UnmarshalJSON([]byte("null"), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, i.Exists)
	assert.Equal(t, 0, i.Value)

	err = i.UnmarshalJSON([]byte("foo"), "", map[string]string{})
	uerr.Assert(t, err, "WCXYWVMOTT")

	err = i.UnmarshalJSON([]byte("1.2"), "", map[string]string{})
	uerr.Assert(t, err, "KVEOETSIJY")

}

func TestIntMarshalJSON(t *testing.T) {

	i := Int{Exists: false, Value: 0}
	ba, err := i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	i = Int{Exists: true, Value: 12}
	ba, err = i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "12", string(ba))

	i = Int{Exists: true, Value: -101}
	ba, err = i.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "-101", string(ba))

}

func TestIntString(t *testing.T) {

	i := Int{Exists: false, Value: 0}
	s := i.String()
	assert.Equal(t, "", s)

	i = Int{Exists: true, Value: 12}
	s = i.String()
	assert.Equal(t, "12", s)

	i = Int{Exists: true, Value: -101}
	s = i.String()
	assert.Equal(t, "-101", s)

}
