package common

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/common"
}
func (p packageType) Unpack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Validator":
		return p.UnpackValidator(context, root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) UnpackValidator(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Validator, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value interface {
		Validate(stack global.Stack, input interface{}) (valid bool, message string, err error)
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := pack.UnpackInterface(context, root, stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface {
			Validate(stack global.Stack, input interface{}) (valid bool, message string, err error)
		})
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
	return Validator(out), false, nil
}
func (p packageType) Repack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Validator":
		return p.RepackValidator(context, root, stack, in.(Validator))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) RepackValidator(context global.Context, root global.Root, stack global.Stack, in Validator) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in interface {
		Validate(stack global.Stack, input interface{}) (valid bool, message string, err error)
	}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return pack.RepackInterface(context, root, stack, false, in)
	}(context, root, stack, (interface {
		Validate(stack global.Stack, input interface{}) (valid bool, message string, err error)
	})(in))
}
func (p packageType) Get(name string) string {
	return ""
}
func (p packageType) Add(packages map[string]global.Package) {
	packages["frizz.io/common"] = Package
}
