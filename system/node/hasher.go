package node

import (
	"github.com/davelondon/kerr"
	"github.com/surge/cityhash"
	"kego.io/json"
)

type NodeHasher struct {
	String  string
	Number  float64
	Bool    bool
	Map     map[string]uint64
	Array   []uint64
	Version int
}

func (p *NodeHasher) Hash() (uint64, error) {

	p.Version = 1

	bytes, err := json.MarshalPlain(p)
	if err != nil {
		return 0, kerr.Wrap("QYEXVJIEOS", err)
	}

	return cityhash.CityHash64(bytes, uint32(len(bytes))), nil

}
