package ke // import "kego.io/ke"

import (
	"io"
	"os"

	"kego.io/json"
	"kego.io/kerr"
	_ "kego.io/system"
)

func Open(filename string, path string, aliases map[string]string) (value interface{}, err error) {

	file, err := os.Open(filename)
	if err != nil {
		err = kerr.New("NDJKHCDCIW", err, "kego.Open", "os.Open")
		return
	}
	defer file.Close()

	err = json.NewDecoder(file, path, aliases).Decode(&value)
	return
}

func Unmarshal(data []byte, v *interface{}, path string, aliases map[string]string) error {
	return json.Unmarshal(data, v, path, aliases)
}

func NewDecoder(r io.Reader, path string, aliases map[string]string) *json.Decoder {
	return json.NewDecoder(r, path, aliases)
}
