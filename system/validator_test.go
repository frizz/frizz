package system

import (
	"testing"

	"kego.io/assert"
)

func TestStringRule_Validate(t *testing.T) {
	r := &String_rule{}
	ok, message, err := r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &String_rule{MaxLength: NewInt(10)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &String_rule{MaxLength: NewInt(10), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &String_rule{MaxLength: NewInt(5), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = &String_rule{MaxLength: NewInt(4), MinLength: NewInt(5)}
	ok, message, err = r.Validate("", map[string]string{})
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "MaxLength 4 must not be lass than MinLength 5", message)
}
