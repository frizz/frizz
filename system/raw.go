package system

import (
	"frizz.io/global"
)

// frizz
type Raw struct {
	Value interface{}
}

func (r *Raw) Unpack(context global.DataContext, in interface{}) (null bool, err error) {
	r.Value = in
	return r.Value == nil, nil
}

func (r *Raw) Repack(context global.DataContext) (value interface{}, dict bool, null bool, err error) {
	return r.Value, false, r.Value == nil, nil
}
