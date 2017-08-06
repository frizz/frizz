package pack

import (
	"frizz.io/global"
	"github.com/pkg/errors"
)

func Unpack(context global.PackageContext, in interface{}) (value interface{}, null bool, err error) {
	c := NewDataContext(context, NewRootContext(context), global.NewStack("root"))
	if err := c.Root().ParseImports(in); err != nil {
		return nil, false, errors.Wrap(err, "parsing imports")
	}
	return UnpackInterface(c, in)
}

func Repack(context global.PackageContext, in interface{}) (value interface{}, dict bool, null bool, err error) {
	c := NewDataContext(context, NewRootContext(context), global.NewStack("root"))
	return RepackInterface(c, true, in)
}
