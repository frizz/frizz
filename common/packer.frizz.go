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
	// interfaceUnpacker
	out, err := root.UnpackInterface(stack, in)
	if err != nil {
		return value, err
	}
	iface, ok := out.(Validator)
	if !ok {
		return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
	}
	return iface, nil
}
