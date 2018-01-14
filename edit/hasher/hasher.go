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

func HashWithImports(files map[string][]byte, imports map[string]uint64) (uint64, error) {
	b, err := json.Marshal(map[string]interface{}{
		"files":   files,
		"imports": imports,
	})
	if err != nil {
		return 0, err
	}
	return farmhash.Hash64(b), nil
}
