package validator

import (
	"encoding/base64"
	"reflect"

	"frizz.io/frizz"
	"frizz.io/global"
	"frizz.io/system"
	"github.com/pkg/errors"
)

func Validate(v interface{}, local global.Package) (valid bool, message string, err error) {
	packages := map[string]global.Package{}
	local.Add(packages)
	return validate(global.NewStack("root"), reflect.ValueOf(v), local, packages)
}

func validate(stack global.Stack, value reflect.Value, local global.Package, packages map[string]global.Package) (valid bool, message string, err error) {

	if valid, message, err := validateType(stack, value, local, packages); err != nil || !valid {
		return valid, message, err
	}

	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Interface {
		// interface: recurse with elem
		return validate(stack, value.Elem(), local, packages)
	}

	switch value.Kind() {
	case reflect.Struct:
		// validate fields
		for i := 0; i < value.Type().NumField(); i++ {
			name := value.Type().Field(i).Name
			field := value.Field(i)
			inner := stack.Append(global.FieldItem(name))
			if valid, message, err := validate(inner, field, local, packages); err != nil || !valid {
				return valid, message, err
			}
		}
	case reflect.Map, reflect.Array, reflect.Slice:
		// validate items
		for i := 0; i < value.Len(); i++ {
			var item reflect.Value
			var inner global.Stack
			if value.Kind() == reflect.Map {
				key := value.MapKeys()[i]
				item = value.MapIndex(key)
				inner = stack.Append(global.MapItem(key.Interface().(string)))
			} else {
				item = value.Index(i)
				inner = stack.Append(global.ArrayItem(i))
			}
			if valid, message, err := validate(inner, item, local, packages); err != nil || !valid {
				return valid, message, err
			}
		}
	}
	return true, "", nil
}

func validateType(stack global.Stack, value reflect.Value, local global.Package, packages map[string]global.Package) (valid bool, message string, err error) {

	if (value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface) && value.IsNil() {
		// it is not possible to validate a nil value by it's type -> return valid
		return true, "", nil
	}

	t := value.Type()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.PkgPath() == "" {
		// native value -> return valid
		return true, "", nil
	}

	pkg, ok := packages[t.PkgPath()]
	if !ok {
		// no typer -> return valid
		return true, "", nil
	}

	typebase64 := pkg.Get(t.Name())
	if typebase64 == "" {
		// no type in typer -> return valid
		return true, "", nil
	}

	typebytes, err := base64.StdEncoding.DecodeString(typebase64)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: decoding base64 of type", stack)
	}

	c := frizz.New(local)
	typeiface, err := c.Unmarshal(typebytes)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: unmarshaling type", stack)
	}

	typ, ok := typeiface.(system.Type)
	if !ok {
		return false, "", errors.Errorf("%s: type should system.Type, got %T", stack, typeiface)
	}

	if valid, message, err := typ.ValidateValue(stack, value); err != nil || !valid {
		return valid, message, err
	}
	return true, "", nil
}
