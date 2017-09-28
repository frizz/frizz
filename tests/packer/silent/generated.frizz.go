package silent

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer/silent"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Silent":
		return p.UnpackSilent(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackSilent(context global.DataContext, in interface{}) (value Silent, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value interface {
		Silent()
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := pack.UnpackInterface(context, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface {
			Silent()
		})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into interface, type %T does not implement interface", context.Location(), out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Silent(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Silent":
		return p.RepackSilent(context, in.(Silent))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackSilent(context global.DataContext, in Silent) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in interface {
		Silent()
	}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return pack.RepackInterface(context, false, in)
	}(context, (interface {
		Silent()
	})(in))
}
func (p packageType) GetData(filename string) string {
	return ""
}
func (p packageType) GetType(name string) string {
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/packer/silent"] = Package
}
