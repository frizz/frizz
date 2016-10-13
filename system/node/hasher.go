package node

import (
	"context"
	"encoding/json"

	"github.com/davelondon/kerr"
	"github.com/surge/cityhash"
)

const currentNodeHasherVersion = 1

type NodeHasher struct {
	String  string
	Number  float64
	Bool    bool
	Null    bool
	Missing bool
	Map     []MapItem
	Array   []uint64
	Version int
}
type MapItem struct {
	Key  string
	Hash uint64
}

type nodeHasherVersionKeyType int

var nodeHasherVersionKey nodeHasherVersionKeyType = 0

func (p *NodeHasher) Hash(ctx context.Context) (uint64, error) {

	if version, ok := ctx.Value(nodeHasherVersionKey).(int); ok {
		p.Version = version
	} else {
		p.Version = currentNodeHasherVersion
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		return 0, kerr.Wrap("QYEXVJIEOS", err)
	}

	return cityhash.CityHash64(bytes, uint32(len(bytes))), nil

}
