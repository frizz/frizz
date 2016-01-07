package parse

import (
	"github.com/surge/cityhash"
	"kego.io/json"
	"kego.io/kerr"
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

	p.Version = 2

	bytes, err := json.MarshalPlain(p)
	if err != nil {
		return 0, kerr.New("TGAEJVECIF", err, "json.Marshal")
	}

	return cityhash.CityHash64(bytes, uint32(len(bytes))), nil

}
