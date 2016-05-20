package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/context/envctx"
)

func TestArrayRule_Enforce(t *testing.T) {
	r := ArrayRule{MaxItems: NewInt(2)}
	ok, message, err := r.Enforce(envctx.Empty, []int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, []int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", message)
	assert.False(t, ok)

	r = ArrayRule{MinItems: NewInt(2)}
	ok, message, err = r.Enforce(envctx.Empty, []int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, []int{1})
	assert.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", message)
	assert.False(t, ok)

	r = ArrayRule{UniqueItems: true}
	ok, message, err = r.Enforce(envctx.Empty, []int{1, 2, 3, 4})
	assert.NoError(t, err)
	assert.Equal(t, "", message)
	assert.True(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, []int{1, 2, 3, 3})
	assert.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item 3", message)
	assert.False(t, ok)

	ok, message, err = r.Enforce(envctx.Empty, []string{"foo", "bar", "foo"})
	assert.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item foo", message)
	assert.False(t, ok)

	r = ArrayRule{}
	ok, message, err = r.Enforce(envctx.Empty, []int{1, 2})
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "", message)

	r = ArrayRule{MaxItems: NewInt(2)}
	ok, message, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	assert.IsError(t, err, "OWTAUVVFBL")

	s := &StringRule{}
	var cr CollectionRule = &ArrayRule{Items: s}
	assert.Equal(t, s, cr.GetItemsRule())
}
