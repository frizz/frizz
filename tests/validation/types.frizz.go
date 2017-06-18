package validation

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

type Packer struct{}

func (Packer) Path() string {
	return "frizz.io/tests/validation"
}
func (p Packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "Simple":
		return p.UnpackSimple(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p Packer) UnpackSimple(root *frizz.Root, stack frizz.Stack, in interface{}) (value Simple, err error) {
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
