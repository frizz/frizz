package sub

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Sub func(*frizz.Root, frizz.Stack, interface{}) (Sub, error)
}{Sub: unpacker_Sub}

func unpacker_Sub(r *frizz.Root, s frizz.Stack, in interface{}) (value Sub, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Sub
	if v, ok := m["String"]; ok {
		s := s.Append(frizz.FieldItem("String"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker/sub", "Sub", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Sub(r, s, in)
	})
}
