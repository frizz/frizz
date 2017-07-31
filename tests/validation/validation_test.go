package validation_test

import (
	"testing"

	"strings"

	"bytes"
	"encoding/json"

	"frizz.io/frizz"
	"frizz.io/system"
	"frizz.io/tests/packer"
	"frizz.io/validators"
)

type testDef struct {
	ty string // type
	js string // json input
	ms string // expect invalid with this message
	er string // expect error containing this string
}

func TestValidate(t *testing.T) {

	tests := map[string]testDef{
		"regex success": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"String": [{"_type": "validators.Regex","Regex": "^foo.*$"}]}
				}]
			}`,
			js: `{"_type": "Natives", "String": "foo"}`,
		},
		"regex fail": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"String": [{"_type": "validators.Regex","Regex": "^foo.*$"}]}
				}]
			}`,
			js: `{"_type": "Natives", "String": "bar"}`,
			ms: `input "bar" did not match regex "^foo.*$"`,
		},
		"gt success": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"Int": [{"_type": "validators.GreaterThan","_value": 2}]}
				}]
			}`,
			js: `{"_type": "Natives","Int": 3}`,
		},
		"gt fail": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"Int": [{"_type": "validators.GreaterThan","_value": 2}]}
				}]
			}`,
			js: `{"_type": "Natives","Int": 1}`,
			ms: "value 1 must be greater than 1",
		},
		"gt success zero": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"Int": [{"_type": "validators.GreaterThan","_value": 2}]}
				}]
			}`,
			js: `{"_type": "Natives","Int": 0}`,
		},
		"gt success empty": {
			ty: `{
				"_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"},
				"_type": "system.Type",
				"Validators": [{
					"_type": "validators.Struct",
					"_value": {"Int": [{"_type": "validators.GreaterThan","_value": 2}]}
				}]
			}`,
			js: `{"_type": "Natives"}`,
		},
	}

	for name, td := range tests {
		r := &frizz.Root{
			Context: frizz.New(packer.Imports).Register(system.Packer, validators.Packer),
			Imports: make(map[string]string),
		}
		v := unmarshal(t, td.ty)
		r.ParseImports(v)
		typ, _, err := system.Packer.UnpackType(r, frizz.Stack{frizz.RootItem("root")}, v)

		f := frizz.New(packer.Imports)
		iface, err := f.Unmarshal([]byte(td.js))
		if err != nil {
			t.Fatal(err)
		}
		valid, message, err := typ.Validate(iface)
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

func unmarshal(t *testing.T, in string) interface{} {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer([]byte(in)))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		t.Fatal(err)
	}
	return v
}
