package validation_test

import (
	"testing"

	"strings"

	"frizz.io/frizz"
	"frizz.io/frizz/validator"
	"frizz.io/tests/validation"
)

type testDef struct {
	js string // json input
	ms string // expect invalid with this message
	er string // expect error containing this string
}

func TestValidate(t *testing.T) {

	tests := map[string]testDef{
		"regex success": {
			js: `{"_type": "Simple","String": "foo"}`,
		},
		"regex fail": {
			js: `{"_type": "Simple","String": "bar"}`,
			ms: `input "bar" did not match regex "^foo.*$"`,
		},
		"gt success": {
			js: `{"_type": "Simple","Int": 3}`,
		},
		"gt fail": {
			js: `{"_type": "Simple","Int": 1}`,
			ms: "value 1 must be greater than 1",
		},
		"gt success zero": {
			js: `{"_type": "Simple","Int": 0}`,
		},
		"gt success empty": {
			js: `{"_type": "Simple"}`,
		},
	}

	for name, td := range tests {
		f := frizz.New(validation.Imports)
		iface, err := f.Unmarshal([]byte(td.js))
		if err != nil {
			t.Fatal(err)
		}
		valid, message, err := validator.Validate(iface, validation.Imports)
		if td.er == "" && err != nil {
			t.Fatalf("%s: error when none expepected: %s", name, err.Error())
		}
		if td.er != "" && err == nil {
			t.Fatalf("%s: no error when one expepected: %s", name, td.er)
		}
		if td.er != "" && err != nil && !strings.Contains(err.Error(), td.er) {
			t.Fatalf("%s: unexpected error: %s. Expected: %s", name, err.Error(), td.er)
		}
		if td.ms == "" && !valid {
			t.Fatalf("%s: expected valid, but result was invalid %#v", name, message)
		}
		if td.ms != "" && valid {
			t.Fatalf("%s: expected invalid %#v, but result was valid", name, td.ms)
		}
	}

}
