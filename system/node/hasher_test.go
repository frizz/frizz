package node

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"golang.org/x/net/context"
)

func TestNodeHasher_Hash(t *testing.T) {
	ctx := context.WithValue(context.Background(), nodeHasherVersionKey, 1)
	h := NodeHasher{
		String: "a",
		Number: 2.0,
		Bool:   true,
		Map:    map[string]uint64{"b": 1},
		Array:  []uint64{1, 2},
	}
	hash, err := h.Hash(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0x566b54beeb09a9da), hash)

	// run with no forced version to get complete code coverage, but ignore
	// value because it will change
	_, err = h.Hash(context.Background())
	assert.NoError(t, err)
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
	hash, err := n2.Hash(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0x8c85083b1815ef45), hash)
}
