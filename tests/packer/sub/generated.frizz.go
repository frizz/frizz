package sub

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer/sub"
}
func (p packageType) Unpack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Sub":
		return p.UnpackSub(context, root, stack, in)
	case "SubInterface":
		return p.UnpackSubInterface(context, root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) UnpackSub(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Sub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value struct {
		String string
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map", stack)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			String string
		}
		if v, ok := m["String"]; ok {
			stack := stack.Append(global.FieldItem("String"))
			u, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackString(stack, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Sub(out), false, nil
}
func (p packageType) UnpackSubInterface(context global.Context, root global.Root, stack global.Stack, in interface{}) (value SubInterface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value interface{}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := pack.UnpackInterface(context, root, stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into interface, type %T does not implement interface", stack, out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return SubInterface(out), false, nil
}
func (p packageType) Repack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Sub":
		return p.RepackSub(context, root, stack, in.(Sub))
	case "SubInterface":
		return p.RepackSubInterface(context, root, stack, in.(SubInterface))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) RepackSub(context global.Context, root global.Root, stack global.Stack, in Sub) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in struct {
		String string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.Context, root global.Root, stack global.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(context, root, stack, (struct {
		String string
	})(in))
}
func (p packageType) RepackSubInterface(context global.Context, root global.Root, stack global.Stack, in SubInterface) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return pack.RepackInterface(context, root, stack, false, in)
	}(context, root, stack, (interface{})(in))
}
func (p packageType) Get(name string) string {
	return ""
}
func (p packageType) Add(packages map[string]global.Package) {
	packages["frizz.io/tests/packer/sub"] = Package
}
