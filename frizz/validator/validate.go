package validator

import (
	"encoding/base64"
	"reflect"

	"frizz.io/frizz"
	"frizz.io/system"
	"github.com/pkg/errors"
)

func Validate(v interface{}, imports frizz.Importer) (valid bool, message string, err error) {
	types := map[string]frizz.Typer{}
	imports.Add(nil, types)
	return validate(frizz.Stack{frizz.RootItem("root")}, reflect.ValueOf(v), imports, types)
}

func validate(stack frizz.Stack, value reflect.Value, imports frizz.Importer, types map[string]frizz.Typer) (valid bool, message string, err error) {
	valid, message, err = validateType(stack, value, imports, types)
	if err != nil {
		return false, "", err
	}
	if !valid {
		return false, message, nil
	}

	// validate fields
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.Type().NumField(); i++ {
			name := value.Type().Field(i).Name
			field := value.Field(i)
			inner := stack.Append(frizz.FieldItem(name))
			valid, message, err = validate(inner, field, imports, types)
			if err != nil {
				return false, "", err
			}
			if !valid {
				return false, message, nil
			}
		}
	case reflect.Map, reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			var item reflect.Value
			var inner frizz.Stack
			if value.Kind() == reflect.Map {
				key := value.MapKeys()[i]
				item = value.MapIndex(key)
				inner = stack.Append(frizz.MapItem(key.Interface().(string)))
			} else {
				item = value.Index(i)
				inner = stack.Append(frizz.ArrayItem(i))
			}
			valid, message, err = validate(inner, item, imports, types)
			if err != nil {
				return false, "", err
			}
			if !valid {
				return false, message, nil
			}
		}
	}
	return true, "", nil
}

func validateType(stack frizz.Stack, value reflect.Value, imports frizz.Importer, types map[string]frizz.Typer) (valid bool, message string, err error) {

	typer, ok := types[value.Type().PkgPath()]
	if !ok {
		// no typer -> return valid
		return true, "", nil
	}

	typebase64 := typer.Get(value.Type().Name())
	if typebase64 == "" {
		// no type in typer -> return valid
		return true, "", nil
	}

	typebytes, err := base64.StdEncoding.DecodeString(typebase64)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: decoding base64 of type", stack)
	}

	f := frizz.New(imports)
	typeiface, err := f.Unmarshal(typebytes)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: unmarshaling type", stack)
	}

	t, ok := typeiface.(system.Type)
	if !ok {
		return false, "", errors.Errorf("%s: type should system.Type, got %T", stack, typeiface)
	}

	valid, message, err = t.ValidateValue(stack, value)
	if err != nil {
		return false, "", err
	}
	if !valid {
		return false, message, nil
	}
	return true, "", nil
}
