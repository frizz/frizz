package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestIntRule_Enforce(t *testing.T) {
	r := Int_rule{Minimum: NewInt(2)}
	ok, message, err := r.Enforce(Int{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 2", message)
	assert.False(t, ok)

	r = Int_rule{Maximum: NewInt(2)}
	ok, message, err = r.Enforce(Int{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value 3 must not be greater than 2", message)
	assert.False(t, ok)

	r = Int_rule{MultipleOf: NewInt(3)}
	ok, message, err = r.Enforce(Int{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewInt(0), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(6), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewInt(4), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 3", message)
	assert.False(t, ok)

}

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
	assert.IsError(t, err, "WCXYWVMOTT")

	err = i.UnmarshalJSON([]byte("1.2"), "", map[string]string{})
	assert.IsError(t, err, "KVEOETSIJY")

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
