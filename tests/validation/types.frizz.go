package validation

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Simple func(*frizz.Root, frizz.Stack, interface{}) (Simple, error)
}{Simple: unpacker_Simple}

func unpacker_Simple(root *frizz.Root, stack frizz.Stack, in interface{}) (value Simple, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Simple
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
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/tests/validation", "Simple", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Simple(root, stack, in)
	})
}
