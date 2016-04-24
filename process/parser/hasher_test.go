package parser

import (
	"testing"

	"kego.io/kerr/assert"
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
	assert.Equal(t, uint64(0x9f31cab0a9900716), h)
}
