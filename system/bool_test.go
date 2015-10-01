package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestNewBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)
}

func TestBoolUnmarshalJSON(t *testing.T) {

	var b Bool

	err := b.Unpack(true)
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)

	err = b.Unpack(false)
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.False(t, b.Value)

	err = b.Unpack(nil)
	assert.NoError(t, err)
	assert.False(t, b.Exists)
	assert.False(t, b.Value)

	err = b.Unpack("foo")
	assert.IsError(t, err, "GXQGNEPJYS")

}
func TestBoolMarshalJSON(t *testing.T) {

	b := Bool{Exists: false, Value: false}
	ba, err := b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	b = Bool{Exists: true, Value: false}
	ba, err = b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "false", string(ba))

	b = Bool{Exists: true, Value: true}
	ba, err = b.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "true", string(ba))

}
func TestBoolString(t *testing.T) {

	b := Bool{Exists: false, Value: false}
	s := b.String()
	assert.Equal(t, "", s)

	b = Bool{Exists: true, Value: false}
	s = b.String()
	assert.Equal(t, "false", s)

	b = Bool{Exists: true, Value: true}
	s = b.String()
	assert.Equal(t, "true", s)

}
