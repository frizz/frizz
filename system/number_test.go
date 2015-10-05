package system

import (
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
)

func TestNumberRule_Enforce(t *testing.T) {
	r := Number_rule{Rule_base: &Rule_base{Optional: false}, Minimum: NewNumber(1.5)}
	ok, message, err := r.Enforce(Number{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewNumber(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(1.5), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum: value 1 must not be less than 1.5", message)
	assert.False(t, ok)

	r = Number_rule{Rule_base: &Rule_base{Optional: false}, Minimum: NewNumber(1.5), ExclusiveMinimum: true}
	ok, message, err = r.Enforce(NewNumber(1.5), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Minimum (exclusive): value 1.5 must be greater than 1.5", message)
	assert.False(t, ok)

	r = Number_rule{Rule_base: &Rule_base{Optional: false}, Maximum: NewNumber(1.5)}
	ok, message, err = r.Enforce(Number{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewNumber(1), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(1.5), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(2), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum: value 2 must not be greater than 1.5", message)
	assert.False(t, ok)

	r = Number_rule{Rule_base: &Rule_base{Optional: false}, Maximum: NewNumber(1.5), ExclusiveMaximum: true}
	ok, message, err = r.Enforce(NewNumber(1.5), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Maximum (exclusive): value 1.5 must be less than 1.5", message)
	assert.False(t, ok)

	r = Number_rule{Rule_base: &Rule_base{Optional: false}, MultipleOf: NewNumber(1.5)}
	ok, message, err = r.Enforce(Number{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value must exist", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewNumber(0), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(1.5), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(3), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(NewNumber(4), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MultipleOf: value 4 must be a multiple of 1.5", message)
	assert.False(t, ok)

}

func TestNewNumber(t *testing.T) {
	n := NewNumber(1.2)
	assert.True(t, n.Exists)
	assert.Equal(t, 1.2, n.Value)
}

func TestNumberUnmarshalJSON(t *testing.T) {

	var n Number

	err := n.Unpack(json.NewJsonUnpacker(1.2))
	assert.NoError(t, err)
	assert.True(t, n.Exists)
	assert.Equal(t, 1.2, n.Value)

	err = n.Unpack(json.NewJsonUnpacker(nil))
	assert.NoError(t, err)
	assert.False(t, n.Exists)
	assert.Equal(t, 0.0, n.Value)

	err = n.Unpack(json.NewJsonUnpacker("foo"))
	assert.IsError(t, err, "YHXBFTONCW")

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
