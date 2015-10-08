package system

import (
	"kego.io/json"
)

type unpackerFunc func([]byte, *interface{}, string, map[string]string) error

func unmarshalFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	return json.Unmarshal(data, i, path, aliases)
}
func unpackFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	var j interface{}
	if err := json.UnmarshalPlain(data, &j); err != nil {
		return err
	}
	return json.Unpack(json.NewJsonUnpacker(j), i, path, aliases)
}
func repackFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	var n Node
	if err := json.UnmarshalPlainContext(data, &n, path, aliases); err != nil {
		return err
	}
	return json.Unpack(&n, i, path, aliases)
}