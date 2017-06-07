package sub

import (
	system "frizz.io/system"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Sub func(interface{}) (Sub, error)
}{Sub: unpacker_Sub}

func unpacker_Sub(in interface{}) (value Sub, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Sub
	if v, ok := m["string"]; ok {
		u, err := func(in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
