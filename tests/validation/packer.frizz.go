package validation

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/tests/validation"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, err error) {
	switch name {
	case "Simple":
		return p.UnpackSimple(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackSimple(root *frizz.Root, stack frizz.Stack, in interface{}) (value Simple, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
		Int    int
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			String string
			Int    int
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Simple(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Simple":
		return p.RepackSimple(root, stack, in.(Simple))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackSimple(root *frizz.Root, stack frizz.Stack, in Simple) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		String string
		Int    int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
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
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		String string
		Int    int
	})(in))
}
