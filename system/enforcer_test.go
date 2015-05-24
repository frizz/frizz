package system

import (
	"testing"

	"kego.io/assert"
)

func TestStringRule_Enforce(t *testing.T) {
	r := String_rule{Equal: NewString("a"), MaxLength: NewInt(1)}
	ok, message, err := r.Enforce(NewString("a"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Equal: Value must exist, and equal 'a'", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(NewString("b"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "Equal: Value must equal 'a'", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce("a", "", map[string]string{})
	assert.IsError(t, err, "SXFBXGQSEA")

	r = String_rule{MaxLength: NewInt(1)}
	ok, message, err = r.Enforce(NewString("ab"), "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "MaxLength: Length must not be greater than 1", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(String{}, "", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

}
