package sub

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Sub func(*frizz.Root, frizz.Stack, interface{}) (Sub, error)
}{Sub: unpacker_Sub}

func unpacker_Sub(root *frizz.Root, stack frizz.Stack, in interface{}) (value Sub, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Sub
	if v, ok := m["String"]; ok {
		stack := stack.Append(frizz.FieldItem("String"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker/sub", "Sub", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Sub(root, stack, in)
	})
}
