package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

func DecodeFile(fpath string) (interface{}, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return DecodeReader(f)
}

func DecodeReader(r io.Reader) (interface{}, error) {
	var v interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}
	return v, nil
}
