package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestArrayRule_Enforce(t *testing.T) {
	r := Array_rule{MaxItems: NewInt(2)}
	ok, message, err := r.Enforce([]int{1, 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce([]int{1, 2, 3}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", message)
	assert.False(t, ok)

	r = Array_rule{MinItems: NewInt(2)}
	ok, message, err = r.Enforce([]int{1, 2}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce([]int{1}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", message)
	assert.False(t, ok)

	r = Array_rule{UniqueItems: true}
	ok, message, err = r.Enforce([]int{1, 2, 3, 4}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce([]int{1, 2, 3, 3}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item 3", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce([]string{"foo", "bar", "foo"}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item foo", message)
	assert.False(t, ok)
}
