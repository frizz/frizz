package system

import "frizz.io/frizz"

// frizz
type Raw struct {
	Value interface{}
}

func (r *Raw) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) (null bool, err error) {
	r.Value = in
	return r.Value == nil, nil
}

func (r *Raw) Repack(root *frizz.Root, stack frizz.Stack) (value interface{}, dict bool, null bool, err error) {
	return r.Value, false, r.Value == nil, nil
}
