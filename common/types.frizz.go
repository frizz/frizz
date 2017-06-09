package common

import (
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Validator func(*frizz.Root, frizz.Stack, interface{}) (Validator, error)
}{Validator: unpacker_Validator}

func unpacker_Validator(r *frizz.Root, s frizz.Stack, in interface{}) (value Validator, err error) {
	// interfaceUnpacker
	out, err := r.UnpackInterface(s, in)
	if err != nil {
		return value, err
	}
	iface, ok := out.(Validator)
	if !ok {
		return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
	}
	return iface, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/common", "Validator", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Validator(r, s, in)
	})
}
