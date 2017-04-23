package parser

import (
	"testing"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestHasher(t *testing.T) {
	p := PackageHasher{
		Path: "a",
		Aliases: map[string]string{
			"b": "c",
			"d": "e",
		},
		Types: map[string]uint64{
			"f": uint64(2),
			"g": uint64(3),
		},
	}
	h, err := p.Hash(1)
	require.NoError(t, err)
	assert.Equal(t, uint64(0x1e7153b485e77b8b), h)
}
