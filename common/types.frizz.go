package common

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

type Packer struct{}

func (Packer) Path() string {
	return "frizz.io/common"
}
func (p Packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "Validator":
		return p.UnpackValidator(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p Packer) UnpackValidator(root *frizz.Root, stack frizz.Stack, in interface{}) (value Validator, err error) {
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
