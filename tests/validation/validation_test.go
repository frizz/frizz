package validation_test

import (
	"testing"

	"fmt"

	"frizz.io/frizz"
	"frizz.io/frizz/validator"
	"frizz.io/tests/validation"
)

func TestValidate(t *testing.T) {

	f := frizz.New(validation.Imports)
	iface, err := f.Unmarshal([]byte(`{
		"_type": "Simple",
		"String": "foo"
	}`))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(validator.Validate(iface, validation.Imports))

}
