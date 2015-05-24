package system

import (
	"testing"

	"kego.io/assert"
	"kego.io/uerr"
)

func TestNewBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)
}

func TestBoolUnmarshalJSON(t *testing.T) {

	var b Bool

	err := b.UnmarshalJSON([]byte("true"), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.True(t, b.Value)

	err = b.UnmarshalJSON([]byte("false"), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, b.Exists)
	assert.False(t, b.Value)

	err = b.UnmarshalJSON([]byte("null"), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, b.Exists)
	assert.False(t, b.Value)

	err = b.UnmarshalJSON([]byte("foo"), "", map[string]string{})
	uerr.Assert(t, err, "CJMOICJGJG")

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
