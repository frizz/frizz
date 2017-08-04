package pack

import (
	"frizz.io/global"
	"github.com/pkg/errors"
)

func Unpack(c global.Context, in interface{}) (value interface{}, null bool, err error) {
	r := NewRoot()
	if err := r.Parse(in); err != nil {
		return nil, false, errors.Wrap(err, "parsing imports")
	}
	s := global.NewStack("root")
	return UnpackInterface(c, r, s, in)
}

func Repack(c global.Context, in interface{}) (value interface{}, dict bool, null bool, err error) {
	r := NewRoot()
	s := global.NewStack("root")
	return RepackInterface(c, r, s, true, in)
}
