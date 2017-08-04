package pack

import (
	"reflect"

	"github.com/pkg/errors"
	"frizz.io/global"
	"frizz.io/utils"
)

func UnpackInterface(context global.Context, root global.Root, stack global.Stack, in interface{}) (value interface{}, null bool, err error) {
	if in == nil {
		return nil, true, nil
	}
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, value should be a map", stack)
	}
	ti, ok := m["_type"]
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, _type field missing", stack)
	}
	ts, ok := ti.(string)
	if !ok {
		return nil, false, errors.Errorf("%s: unpacking into interface, _type should be a string", stack)
	}

	v := m["_value"]

	if context.Path() == "" {
		return nil, false, errors.Errorf("%s: unpacking into interface, local path is not set", stack)
	}

	path, name, err := utils.ParseReference(ts, context.Path(), root.Imports())
	if err != nil {
		return nil, false, errors.Wrapf(err, "%s: unpacking into interface", stack)
	} else if path != "" && name != "" {
		// we have a path and a name => unpack with packer
		p := context.Package(path)
		if p == nil {
			return nil, false, errors.Errorf("%s: unpacking into interface, packer for %s not registered", stack, context.Path())
		}
		if v != nil {
			return p.Unpack(context, root, stack, v, name)
		}
		return p.Unpack(context, root, stack, in, name)
	} else if name != "" {
		// no path path but we have a name => native type
		if v == nil {
			return nil, false, errors.Errorf("%s: unpacking native type into interface, _value field missing", stack)
		}
		return UnpackNative(stack, name, v)
	} else {
		// no path or name => type not found
		return nil, false, errors.Errorf("%s: unsupported type %s", stack, ts)
	}
}

func RepackInterface(context global.Context, root global.Root, stack global.Stack, isroot bool, in interface{}) (value interface{}, dict bool, null bool, err error) {
	if in == nil {
		return nil, false, true, nil
	}
	t := reflect.TypeOf(in)
	if t.PkgPath() == "" {
		switch t.Name() {
		case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64", "rune", "string":
			value, dict, null, err = RepackNative(stack, t.Name(), in)
			if err != nil {
				return nil, false, false, err
			}
		default:
			return nil, false, false, errors.Errorf("%s: unsupported type %s", stack, t.Name())
		}
	} else {
		p := context.Package(t.PkgPath())
		if p == nil {
			return nil, false, false, errors.Errorf("%s: can't find packer for %s", stack, t.PkgPath())
		}
		value, dict, null, err = p.Repack(context, root, stack, in, t.Name())
		if err != nil {
			return nil, false, false, err
		}
	}

	typeName := t.Name()
	if t.PkgPath() != "" && t.PkgPath() != context.Path() {
		alias := root.Register(t.PkgPath())
		typeName = alias + "." + t.Name()
	}

	out := map[string]interface{}{}

	vmap, isMap := value.(map[string]interface{})
	isStruct := isMap && !dict

	if isStruct {
		// if value is a struct, use the value as out
		out = vmap
	} else if !null {
		// if value is not a struct, add the value to _value only if not null
		out["_value"] = value
	}

	// always add type information
	out["_type"] = typeName

	// when repacking into a root, add imports to _import
	if isroot {
		imports := make(map[string]interface{}, len(root.Imports()))
		for a, p := range root.Imports() {
			imports[a] = p
		}
		out["_import"] = imports
	}

	return out, false, false, nil
}
