package packer

import (
	"fmt"

	"frizz.io/frizz"
	"github.com/pkg/errors"
)

// frizz
type InterfaceValidator string

func (i InterfaceValidator) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	iface, ok := input.(Interface)
	if !ok {
		return false, "", errors.Errorf("%s: InterfaceValidator can only validate type packer.Interface. Found %T", stack, input)
	}
	if iface.Foo() != string(i) {
		return false, fmt.Sprintf("%s: value %#v should be equal to %#v", stack, iface.Foo(), string(i)), nil
	}
	return true, "", nil
}
