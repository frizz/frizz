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
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "Validator":
		return p.UnpackValidator(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackValidator(root *frizz.Root, stack frizz.Stack, in interface{}) (value Validator, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface {
		Validate(input interface{}) (valid bool, message string, err error)
	}, err error) {
		// interfaceUnpacker
		out, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, err
		}
		iface, ok := out.(interface {
			Validate(input interface{}) (valid bool, message string, err error)
		})
		if !ok {
			return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		return iface, nil
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Validator(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, bool, error) {
	switch name {
	case "Validator":
		return p.RepackValidator(root, stack, in.(Validator))
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackValidator(root *frizz.Root, stack frizz.Stack, in Validator) (interface{}, bool, error) {
	return func(root *frizz.Root, stack frizz.Stack, in interface {
		Validate(input interface{}) (valid bool, message string, err error)
	}) (interface{}, bool, error) {
		// interfaceRepacker
		return frizz.RepackInterface(root, stack, in)
	}(root, stack, (interface {
		Validate(input interface{}) (valid bool, message string, err error)
	})(in))
}
