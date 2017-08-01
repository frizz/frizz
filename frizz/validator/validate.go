package validator

import (
	"encoding/base64"
	"reflect"

	"frizz.io/frizz"
	"github.com/pkg/errors"
	"frizz.io/system"
)

func Validate(v interface{}, imports frizz.Importer) (valid bool, message string, err error) {
	types := map[string]frizz.Typer{}
	imports.Add(nil, types)

	value := reflect.ValueOf(v)

	typer, ok := types[value.Type().PkgPath()]
	if !ok {
		// no validator -> valid?
		return true, "", nil
	}

	typebase64 := typer.Get(value.Type().Name())
	if typebase64 == "" {
		// no validator -> valid?
		return true, "", nil
	}

	typebytes, err := base64.StdEncoding.DecodeString(typebase64)
	if err != nil {
		return false, "", errors.Wrap(err, "decoding base64 of type")
	}

	f := frizz.New(imports)
	typeiface, err := f.Unmarshal(typebytes)
	if err != nil {
		return false, "", errors.Wrap(err, "unmarshaling type")
	}

	t, ok := typeiface.(system.Type)
	if !ok {
		return false, "", errors.Errorf("type should system.Type, got %T", typeiface)
	}

	return t.Validate(frizz.Stack{frizz.RootItem("root")}, v)

}
