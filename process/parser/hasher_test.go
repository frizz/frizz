package parser

import (
	"testing"

	"github.com/davelondon/ktest/assert"
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
	assert.NoError(t, err)
	assert.Equal(t, uint64(0x64290df7e1cd5bb2), h)
}
