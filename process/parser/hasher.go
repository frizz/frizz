package parser

import (
	"github.com/davelondon/kerr"
	"github.com/surge/cityhash"
	"kego.io/json"
)

type PackageHasher struct {
	Path    string
	Aliases map[string]string
	Types   map[string]uint64
	Version int
}

type AliasInfo struct {
	Alias string
	Hash  uint64
}

func (p *PackageHasher) Hash() (uint64, error) {

	p.Version = 5

	bytes, err := json.MarshalPlain(p)
	if err != nil {
		return 0, kerr.Wrap("TGAEJVECIF", err)
	}

	return cityhash.CityHash64(bytes, uint32(len(bytes))), nil

}
