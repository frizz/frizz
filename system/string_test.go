package system

import (
	"testing"

	"kego.io/kerr/assert"

	"strconv"
)

func TestNewString(t *testing.T) {
	s := NewString("a")
	assert.True(t, s.Exists)
	assert.Equal(t, "a", s.Value)
}

func TestStringUnmarshalJSON(t *testing.T) {

	var s String

	err := s.UnmarshalJSON([]byte(strconv.Quote(`foo "bar"`)), "", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, s.Exists)
	assert.Equal(t, `foo "bar"`, s.Value)

	err = s.UnmarshalJSON([]byte("null"), "", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, s.Exists)
	assert.Equal(t, "", s.Value)

	err = s.UnmarshalJSON([]byte("foo"), "", map[string]string{})
	assert.IsError(t, err, "ACHMRKVFAB")

}

func TestStringMarshalJSON(t *testing.T) {

	s := String{Exists: false, Value: ""}
	ba, err := s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(ba))

	s = String{Exists: true, Value: `foo "bar"`}
	ba, err = s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `"foo \"bar\""`, string(ba))

}

func TestStringString(t *testing.T) {

	s := String{Exists: false, Value: ""}
	str := s.String()
	assert.Equal(t, "", str)

	s = String{Exists: true, Value: `foo "bar"`}
	str = s.String()
	assert.Equal(t, `foo "bar"`, str)

}
