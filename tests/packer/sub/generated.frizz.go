package sub

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer/sub"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Sub":
		return p.UnpackSub(context, in)
	case "SubInterface":
		return p.UnpackSubInterface(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackSub(context global.DataContext, in interface{}) (value Sub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		String string
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			String string
		}
		if v, ok := m["String"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("String")))
			u, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackString(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Sub(out), false, nil
}
func (p packageType) UnpackSubInterface(context global.DataContext, in interface{}) (value SubInterface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value interface{}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := pack.UnpackInterface(context, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface{})
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
	return SubInterface(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Sub":
		return p.RepackSub(context, in.(Sub))
	case "SubInterface":
		return p.RepackSubInterface(context, in.(SubInterface))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackSub(context global.DataContext, in Sub) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		String string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		String string
	})(in))
}
func (p packageType) RepackSubInterface(context global.DataContext, in SubInterface) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in interface{}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return pack.RepackInterface(context, false, in)
	}(context, (interface{})(in))
}
func (p packageType) GetData(filename string) string {
	return ""
}
func (p packageType) GetType(name string) string {
	switch name {
	case "Sub":
		return ""
	case "SubInterface":
		return ""
	}
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/packer/sub"] = Package
}
func (p packageType) Loader(loader global.Loader) dataType {
	return dataType{loader}
}

var Data = Package.Loader(pack.DefaultLoader)

type dataType struct {
	loader global.Loader
}
