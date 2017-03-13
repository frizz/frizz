package node

import (
	"testing"

	"context"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestNodeHasher_Hash(t *testing.T) {
	ctx := context.WithValue(context.Background(), nodeHasherVersionKey, 1)
	h := NodeHasher{
		Missing: true,
		Null:    true,
		String:  "a",
		Number:  2.0,
		Bool:    true,
		Map:     []MapItem{{Key: "b", Hash: 1}},
		Array:   []uint64{1, 2},
	}
	hash, err := h.Hash(ctx)
	require.NoError(t, err)
	assert.Equal(t, uint64(0x678307ecafd3a89d), hash)

	// run with no forced version to get complete code coverage, but ignore
	// value because it will change
	_, err = h.Hash(context.Background())
	require.NoError(t, err)
}

func TestNode_Hash(t *testing.T) {
	ctx := context.WithValue(context.Background(), nodeHasherVersionKey, 1)
	n1 := &Node{
		ValueString: "a",
		ValueNumber: 2.0,
		ValueBool:   true,
	}
	n2 := &Node{
		ValueString: "b",
		ValueNumber: 3.0,
		ValueBool:   false,
		Map:         map[string]*Node{"a": n1},
		Array:       []*Node{n1},
	}
	err := n2.RecomputeHash(ctx, true)
	require.NoError(t, err)
	hash := n2.Hash()
	assert.Equal(t, uint64(0x61e2ccdb839475f1), hash)
}
