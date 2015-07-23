package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestMapRule_Enforce(t *testing.T) {
	r := Map_rule{MaxItems: NewInt(2)}
	ok, message, err := r.Enforce(map[string]int{"foo": 1, "bar": 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(map[string]int{"foo": 1, "bar": 2, "baz": 3}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", message)
	assert.False(t, ok)

	r = Map_rule{MinItems: NewInt(2)}
	ok, message, err = r.Enforce(map[string]int{"foo": 1, "bar": 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(map[string]int{"foo": 1}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", message)
	assert.False(t, ok)

}
