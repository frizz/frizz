package packer

import (
	"fmt"

	"frizz.io/global"

	"github.com/pkg/errors"
)

// frizz
type InterfaceValidator string

func (i InterfaceValidator) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	iface, ok := input.(Interface)
	if !ok {
		return false, "", errors.Errorf("%s: InterfaceValidator can only validate type packer.Interface. Found %T", context.Location(), input)
	}
	if iface.Foo() != string(i) {
		return false, fmt.Sprintf("%s: value %#v should be equal to %#v", context.Location(), iface.Foo(), string(i)), nil
	}
	return true, "", nil
}
