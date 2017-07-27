package common

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/common"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Validator":
		return p.UnpackValidator(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackValidator(root *frizz.Root, stack frizz.Stack, in interface{}) (value Validator, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface {
		Validate(input interface{}) (valid bool, message string, err error)
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface {
			Validate(input interface{}) (valid bool, message string, err error)
		})
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
	return Validator(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Validator":
		return p.RepackValidator(root, stack, in.(Validator))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackValidator(root *frizz.Root, stack frizz.Stack, in Validator) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in interface {
		Validate(input interface{}) (valid bool, message string, err error)
	}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return root.RepackInterface(stack, false, in)
	}(root, stack, (interface {
		Validate(input interface{}) (valid bool, message string, err error)
	})(in))
}

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/common"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	if packers != nil {
		packers["frizz.io/common"] = Packer
	}
}
