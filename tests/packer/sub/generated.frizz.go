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
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Sub":
		return p.UnpackSub(root, stack, in)
	case "SubInterface":
		return p.UnpackSubInterface(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value Sub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into struct, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			String string
		}
		if v, ok := m["String"]; ok {
			stack := stack.Append(frizz.FieldItem("String"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackString(stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Sub(out), false, nil
}
func (p packer) UnpackSubInterface(root *frizz.Root, stack frizz.Stack, in interface{}) (value SubInterface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface{}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface{})
		if !ok {
			return value, false, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return SubInterface(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Sub":
		return p.RepackSub(root, stack, in.(Sub))
	case "SubInterface":
		return p.RepackSubInterface(root, stack, in.(SubInterface))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackSubInterface(root *frizz.Root, stack frizz.Stack, in SubInterface) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return root.RepackInterface(stack, false, in)
	}(root, stack, (interface{})(in))
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
func AddImports(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	if packers != nil {
		packers["frizz.io/tests/packer/sub"] = Packer
	}
}
