package validation_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"frizz.io/frizz"
	"frizz.io/system"
	"frizz.io/tests/packer"
	"frizz.io/validators"
)

func TestNumbers(t *testing.T) {
	vals := map[string]valDef{
		"gt int": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.GreaterThan", "_value": 2}
			]}`,
			tests: map[string]testDef{
				"int success":   {data: `{"_type": "Int", "_value": 3}`},
				"int fail eq":   {data: `{"_type": "Int", "_value": 2}`, msg: `root: value 2 must be greater than 2`},
				"int fail lt":   {data: `{"_type": "Int", "_value": 1}`, msg: `root: value 1 must be greater than 2`},
				"float success": {data: `{"_type": "Float64", "_value": 3.1}`},
				"float fail eq": {data: `{"_type": "Float64", "_value": 2.0}`, msg: `root: value 2 must be greater than 2`},
				"float fail lt": {data: `{"_type": "Float64", "_value": 1.1}`, msg: `root: value 1.1 must be greater than 2`},
			},
		},
		"gt int neg": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.GreaterThan", "_value": -2}
			]}`,
			tests: map[string]testDef{
				"uint err":      {data: `{"_type": "Uint", "_value": 3}`, err: `root: type packer.Uint can only be compared with uint64, not -2: strconv.ParseUint: parsing "-2": invalid syntax`},
				"int success":   {data: `{"_type": "Int", "_value": -1}`},
				"int fail eq":   {data: `{"_type": "Int", "_value": -2}`, msg: "root: value -2 must be greater than -2"},
				"int fail lt":   {data: `{"_type": "Int", "_value": -3}`, msg: "root: value -3 must be greater than -2"},
				"float success": {data: `{"_type": "Float64", "_value": -1.9}`},
				"float fail eq": {data: `{"_type": "Float64", "_value": -2.0}`, msg: `root: value -2 must be greater than -2`},
				"float fail lt": {data: `{"_type": "Float64", "_value": -2.1}`, msg: `root: value -2.1 must be greater than -2`},
			},
		},
		"gt float": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.GreaterThan", "_value": 2.1}
			]}`,
			tests: map[string]testDef{
				"int err":       {data: `{"_type": "Int", "_value": 3}`, err: `root: type packer.Int can only be compared with int64, not 2.1: strconv.ParseInt: parsing "2.1": invalid syntax`},
				"float success": {data: `{"_type": "Float64", "_value": 3.1}`},
				"float fail eq": {data: `{"_type": "Float64", "_value": 2.1}`, msg: `root: value 2.1 must be greater than 2.1`},
				"float fail lt": {data: `{"_type": "Float64", "_value": 1.1}`, msg: `root: value 1.1 must be greater than 2.1`},
			},
		},
		"eq int": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Equal", "_value": 2}
			]}`,
			tests: map[string]testDef{
				"success": {data: `{"_type": "Int", "_value": 2}`},
				"fail":    {data: `{"_type": "Int", "_value": 1}`, msg: "root: value 1 must be equal to 2"},
			},
		},
		"eq string": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Equal", "_value": "a"}
			]}`,
			tests: map[string]testDef{
				"success": {data: `{"_type": "String", "_value": "a"}`},
				"fail":    {data: `{"_type": "String", "_value": "b"}`, msg: "root: value \"b\" must be equal to \"a\""},
			},
		},
	}
	run(t, "Numbers", vals)
}

func TestStructs(t *testing.T) {
	vals := map[string]valDef{
		"equal": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"String": [{"_type": "validators.Equal", "_value": "a"}]}}
			]}`,
			tests: map[string]testDef{
				"success": {data: `{"_type": "Natives", "String": "a"}`},
				"fail":    {data: `{"_type": "Natives", "String": "b"}`, msg: `root.String: value "b" must be equal to "a"`},
				"empty":   {data: `{"_type": "Natives"}`, msg: `root.String: value "" must be equal to "a"`},
			},
		},
		"unknown": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"Foo": [{"_type": "validators.Equal", "_value": "a"}]}}
			]}`,
			tests: map[string]testDef{
				"err": {data: `{"_type": "Natives", "String": "a"}`, err: `root: field Foo not found in packer.Natives`},
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
			valid, message, err := typ.Validate(frizz.Stack{frizz.RootItem("root")}, iface)
			if test.err == "" && err != nil {
				t.Errorf("%s - %s - %s: error when none expepected: %s", name, valName, testName, err.Error())
			}
			if test.err != "" && err == nil {
				t.Errorf("%s - %s - %s: no error when one expepected: %s", name, valName, testName, test.err)
			}
			if test.err != "" && err != nil && err.Error() != test.err {
				t.Errorf("%s - %s - %s: unexpected error: %s. Expected: %s", name, valName, testName, err.Error(), test.err)
			}
			if test.err == "" && test.msg == "" && !valid {
				t.Errorf("%s - %s - %s: expected valid, but result was invalid %#v", name, valName, testName, message)
			}
			if test.err == "" && test.msg != "" && valid {
				t.Errorf("%s - %s - %s: expected invalid %#v, but result was valid", name, valName, testName, test.msg)
			}
			if test.err == "" && test.msg != "" && !valid && message != test.msg {
				t.Errorf("%s - %s - %s: expected %#v, but message was %#v", name, valName, testName, test.msg, message)
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

type valDef struct {
	typeFile string // type
	tests    map[string]testDef
}

type testDef struct {
	data string // json input
	msg  string // expect invalid with this message
	err  string // expect error containing this string
}
