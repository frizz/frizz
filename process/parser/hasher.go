package parser

import (
	"encoding/json"

	"github.com/dave/kerr"
	"github.com/surge/cityhash"
)

type PackageHasher struct {
	Path    string
	Aliases map[string]string
	Types   map[string]uint64
	Exports map[string]uint64
	Version int
}

func (p *PackageHasher) Hash() (uint64, error) {

	p.Version = 6

	bytes, err := json.Marshal(p)
	if err != nil {
		return 0, kerr.Wrap("TGAEJVECIF", err)
	}

	return cityhash.CityHash64(bytes, uint32(len(bytes))), nil

}
