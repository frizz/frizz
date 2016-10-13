package system

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/envctx"
)

func TestArrayRule_Enforce(t *testing.T) {
	r := ArrayRule{MaxItems: NewInt(2)}
	fail, messages, err := r.Enforce(envctx.Empty, []int{1, 2})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, []int{1, 2, 3})
	require.NoError(t, err)
	assert.Equal(t, "MaxItems: length 3 should not be greater than 2", messages[0])
	assert.True(t, fail)

	r = ArrayRule{MinItems: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, []int{1, 2})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, []int{1})
	require.NoError(t, err)
	assert.Equal(t, "MinItems: length 1 should not be less than 2", messages[0])
	assert.True(t, fail)

	r = ArrayRule{UniqueItems: true}
	fail, messages, err = r.Enforce(envctx.Empty, []int{1, 2, 3, 4})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, []int{1, 2, 3, 3})
	require.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item 3", messages[0])
	assert.True(t, fail)

	fail, messages, err = r.Enforce(envctx.Empty, []string{"foo", "bar", "foo"})
	require.NoError(t, err)
	assert.Equal(t, "UniqueItems: array contains duplicate item foo", messages[0])
	assert.True(t, fail)

	r = ArrayRule{}
	fail, messages, err = r.Enforce(envctx.Empty, []int{1, 2})
	require.NoError(t, err)
	assert.Equal(t, 0, len(messages))
	assert.False(t, fail)

	r = ArrayRule{MaxItems: NewInt(2)}
	fail, messages, err = r.Enforce(envctx.Empty, map[string]int{"foo": 1, "bar": 2})
	assert.IsError(t, err, "OWTAUVVFBL")

	s := &StringRule{}
	var cr CollectionRule = &ArrayRule{Items: s}
	assert.Equal(t, s, cr.GetItemsRule())
}
