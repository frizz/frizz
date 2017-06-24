package sub

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/tests/packer/sub"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, err error) {
	switch name {
	case "Sub":
		return p.UnpackSub(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value Sub, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			String string
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Sub(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Sub":
		return p.RepackSub(root, stack, in.(Sub))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackSub(root *frizz.Root, stack frizz.Stack, in Sub) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		String string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		String string
	})(in))
}
