package kego // import "kego.io/kego"

import (
	"io"
	"os"

	"kego.io/json"
	"kego.io/kerr"
	_ "kego.io/system"
)

func Open(filename string, path string, imports map[string]string) (value interface{}, err error) {

	file, err := os.Open(filename)
	if err != nil {
		err = kerr.New("NDJKHCDCIW", err, "kego.Open", "os.Open")
		return
	}
	defer file.Close()

	err = json.NewDecoder(file, path, imports).Decode(&value)
	return
}

func Unmarshal(data []byte, v *interface{}, path string, imports map[string]string) error {
	return json.Unmarshal(data, v, path, imports)
}

func NewDecoder(r io.Reader, path string, imports map[string]string) *json.Decoder {
	return json.NewDecoder(r, path, imports)
}
