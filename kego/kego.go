package kego // import "kego.io/kego"

import (
	"io"

	"kego.io/json"
	_ "kego.io/system"
)

func Unmarshal(data []byte, v *interface{}, path string, imports map[string]string) error {
	return json.Unmarshal(data, v, path, imports)
}

func NewDecoder(r io.Reader, path string, imports map[string]string) *json.Decoder {
	return json.NewDecoder(r, path, imports)
}
