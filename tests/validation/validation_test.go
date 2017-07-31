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

type valDef struct {
	typeFile string // type
	tests    map[string]testDef
}

type testDef struct {
	data string // json input
	msg  string // expect invalid with this message
	err  string // expect error containing this string
}

func TestStructs(t *testing.T) {
	vals := map[string]valDef{
		"regex": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"String": [{"_type": "validators.Regex","Regex": "^foo.*$"}]}}
			]}`,
			tests: map[string]testDef{
				"success": {data: `{"_type": "Natives", "String": "foo"}`},
				"fail":    {data: `{"_type": "Natives", "String": "bar"}`, msg: `input "bar" did not match regex "^foo.*$"`},
			},
		},
		"gt": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"Int": [{"_type": "validators.GreaterThan","_value": 2}]}}
			]}`,
			tests: map[string]testDef{
				"success":       {data: `{"_type": "Natives", "Int": 3}`},
				"fail":          {data: `{"_type": "Natives", "Int": 1}`, msg: "value 1 must be greater than 1"},
				"success zero":  {data: `{"_type": "Natives", "Int": 0}`},
				"success empty": {data: `{"_type": "Natives"}`},
			},
		},
	}
	run(t, "Structs", vals)
}

func run(t *testing.T, name string, vals map[string]valDef) {
	for valName, val := range vals {
		for testName, test := range val.tests {
			r := &frizz.Root{
				// packer.Imports does not automatically import system or validators, so we must register manually
				Context: frizz.New(packer.Imports).Register(system.Packer, validators.Packer),
				Imports: make(map[string]string),
			}
			v, err := unmarshal(t, val.typeFile)
			if err != nil {
				t.Fatalf("%s - %s - %s: unmarshaling typeFile: %s", name, valName, testName, err.Error())
			}
			if err := r.ParseImports(v); err != nil {
				t.Fatalf("%s - %s - %s: parsing typeFile imports: %s", name, valName, testName, err.Error())
			}
			typ, _, err := system.Packer.UnpackType(r, frizz.Stack{frizz.RootItem("root")}, v)
			if err != nil {
				t.Fatalf("%s - %s - %s: unpacking typeFile: %s", name, valName, testName, err.Error())
			}

			f := frizz.New(packer.Imports)
			iface, err := f.Unmarshal([]byte(test.data))
			if err != nil {
				t.Fatalf("%s - %s - %s: %s", name, valName, testName, err.Error())
			}
			valid, message, err := typ.Validate(iface)
			if test.err == "" && err != nil {
				t.Fatalf("%s - %s - %s: error when none expepected: %s", name, valName, testName, err.Error())
			}
			if test.err != "" && err == nil {
				t.Fatalf("%s - %s - %s: no error when one expepected: %s", name, valName, testName, test.err)
			}
			if test.err != "" && err != nil && !strings.Contains(err.Error(), test.err) {
				t.Fatalf("%s - %s - %s: unexpected error: %s. Expected: %s", name, valName, testName, err.Error(), test.err)
			}
			if test.msg == "" && !valid {
				t.Fatalf("%s - %s - %s: expected valid, but result was invalid %#v", name, valName, testName, message)
			}
			if test.msg != "" && valid {
				t.Fatalf("%s - %s - %s: expected invalid %#v, but result was valid", name, valName, testName, test.msg)
			}
		}
	}
}

func unmarshal(t *testing.T, in string) (interface{}, error) {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer([]byte(in)))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
