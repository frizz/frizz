package hasher

import (
	"encoding/json"

	"github.com/leemcloughlin/gofarmhash"
)

func Hash(in map[string][]byte) (uint64, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	return farmhash.Hash64(b), nil
}
