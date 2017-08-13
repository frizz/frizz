package validator

import (
	"encoding/base64"
	"reflect"

	"frizz.io/global"
	"frizz.io/pack"
	"frizz.io/system"
	"github.com/pkg/errors"
)

func Validate(p global.PackageContext, v interface{}) (valid bool, message string, err error) {
	context := pack.NewValidationContext(p, global.NewStack("root"))
	return validate(context, reflect.ValueOf(v))
}

func validate(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {

	if valid, message, err := validateType(context, value); err != nil || !valid {
		return valid, message, err
	}

	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Interface {
		// interface: recurse with elem
		return validate(context, value.Elem())
	}

	switch value.Kind() {
	case reflect.Struct:
		// validate fields
		for i := 0; i < value.Type().NumField(); i++ {
			name := value.Type().Field(i).Name
			field := value.Field(i)
			inner := pack.NewValidationContext(context.Package(), context.Location().Child(global.FieldItem(name)))
			if valid, message, err := validate(inner, field); err != nil || !valid {
				return valid, message, err
			}
		}
	case reflect.Map, reflect.Array, reflect.Slice:
		// validate items
		for i := 0; i < value.Len(); i++ {
			var item reflect.Value
			var inner global.ValidationContext
			if value.Kind() == reflect.Map {
				key := value.MapKeys()[i]
				item = value.MapIndex(key)
				inner = pack.NewValidationContext(context.Package(), context.Location().Child(global.MapItem(key.Interface().(string))))
			} else {
				item = value.Index(i)
				inner = pack.NewValidationContext(context.Package(), context.Location().Child(global.ArrayItem(i)))
			}
			if valid, message, err := validate(inner, item); err != nil || !valid {
				return valid, message, err
			}
		}
	}
	return true, "", nil
}

func validateType(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {

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

	pkg := context.Package().Get(t.PkgPath())
	if pkg == nil {
		// no typer -> return valid
		return true, "", nil
	}

	typebase64 := pkg.GetData(pkg.GetType(t.Name()))
	if typebase64 == "" {
		// no type in typer -> return valid
		return true, "", nil
	}

	typebytes, err := base64.StdEncoding.DecodeString(typebase64)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: decoding base64 of type", context.Location())
	}

	typeiface, err := pack.Unmarshal(context.Package(), typebytes)
	if err != nil {
		return false, "", errors.Wrapf(err, "%s: unmarshaling type", context.Location())
	}

	typ, ok := typeiface.(system.Type)
	if !ok {
		return false, "", errors.Errorf("%s: type should system.Type, got %T", context.Location(), typeiface)
	}

	if valid, message, err := typ.ValidateValue(context, value); err != nil || !valid {
		return valid, message, err
	}
	return true, "", nil
}
