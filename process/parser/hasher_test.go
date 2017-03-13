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
	h, err := p.Hash()
	require.NoError(t, err)
	assert.Equal(t, uint64(0x466636fd7d88d949), h)
}
