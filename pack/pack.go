package pack

import (
	"bytes"

	"frizz.io/global"
	"frizz.io/utils"
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

func Unmarshal(context global.PackageContext, in []byte) (interface{}, error) {
	data, err := utils.DecodeReader(bytes.NewBuffer(in))
	if err != nil {
		return nil, err
	}
	v, null, err := Unpack(context, data)
	if err != nil || null {
		return nil, err
	}
	if null {
		return nil, nil
	}
	return v, nil
}
