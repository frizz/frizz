package validation_test

import (
	"testing"

	"bytes"
	"encoding/base64"
	"encoding/json"

	"frizz.io/frizz"
	"frizz.io/frizz/validator"
	"frizz.io/system"
	"frizz.io/tests/packer"
	"frizz.io/validators"
)

type msss map[string]map[string]string
type mss map[string]string

func TestPointersNull(t *testing.T) {
	vals := map[string]valDef{
		"false": {
			types: msss{"frizz.io/tests/packer": mss{"Pointers": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"Int": [{"_type": "validators.NotNull"}]}}
			]}`}},
			tests: map[string]testDef{
				"success": {data: `{"_type": "Pointers", "Int": 1}`},
				"null":    {data: `{"_type": "Pointers"}`, msg: "root.Int: value must not be null"},
			},
		},
		"true": {
			types: msss{"frizz.io/tests/packer": mss{"Pointers": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"Int": [{"_type": "validators.IsNull"}]}}
			]}`}},
			tests: map[string]testDef{
				"success": {data: `{"_type": "Pointers", "Int": 1}`, msg: "root.Int: value 1 must be null"},
				"null":    {data: `{"_type": "Pointers"}`},
			},
		},
	}
	run(t, "PointersNull", vals)
}

func TestPointersByStruct(t *testing.T) {
	vals := map[string]valDef{
		"eq": {
			types: msss{"frizz.io/tests/packer": mss{"Pointers": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"Int": [{"_type": "validators.Equal", "_value": 2}]}}
			]}`}},
			tests: map[string]testDef{
				"success": {data: `{"_type": "Pointers", "Int": 2}`},
				"null":    {data: `{"_type": "Pointers"}`},
				"zero":    {data: `{"_type": "Pointers", "Int": 0}`, msg: "root.Int: value 0 must be equal to 2"},
				"fail":    {data: `{"_type": "Pointers", "Int": 1}`, msg: "root.Int: value 1 must be equal to 2"},
			},
		},
	}
	run(t, "PointersByStruct", vals)
}

func TestPointers(t *testing.T) {
	vals := map[string]valDef{
		"eq": {
			types: msss{"frizz.io/tests/packer": mss{"Int": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Equal", "_value": 2}
			]}`}},
			tests: map[string]testDef{
				"success": {data: `{"_type": "Pointers", "Int": 2}`},
				"null":    {data: `{"_type": "Pointers"}`},
				"zero":    {data: `{"_type": "Pointers", "Int": 0}`, msg: "root.Int: value 0 must be equal to 2"},
				"fail":    {data: `{"_type": "Pointers", "Int": 1}`, msg: "root.Int: value 1 must be equal to 2"},
			},
		},
	}
	run(t, "Pointers", vals)
}

func TestType(t *testing.T) {
	vals := map[string]valDef{
		"sub": {
			types: msss{"frizz.io/tests/packer/sub": mss{"Sub": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.Struct", "_value": {"String": [{"_type": "validators.Equal", "_value": "a"}]}}
			]}`}},
			tests: map[string]testDef{
				"struct": {data: `{"_type": "Qual", "Sub": {"String": "b"}}`, msg: `root.Sub.String: value "b" must be equal to "a"`},
				"map":    {data: `{"_type": "SubMap", "Map": {"a": {"String": "b"}}}`, msg: `root.Map["a"].String: value "b" must be equal to "a"`},
				"slice":  {data: `{"_type": "SubSlice", "Slice": [{"String": "b"}]}`, msg: `root.Slice[0].String: value "b" must be equal to "a"`},
			},
		},
	}
	run(t, "Type", vals)
}

func TestMocks(t *testing.T) {
	vals := map[string]valDef{
		"gt int": {
			types: msss{"frizz.io/tests/packer": mss{"Int": `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "validators.GreaterThan", "_value": 2}
			]}`}},
			tests: map[string]testDef{
				"int success": {data: `{"_type": "Int", "_value": 3}`},
				"int fail eq": {data: `{"_type": "Int", "_value": 2}`, msg: `root: value 2 must be greater than 2`},
				"int fail lt": {data: `{"_type": "Int", "_value": 1}`, msg: `root: value 1 must be greater than 2`},
			},
		},
	}
	run(t, "Mocks", vals)
}

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
			iface, err := frizz.New(packer.Imports).Unmarshal([]byte(test.data))
			if err != nil {
				t.Fatalf("%s - %s - %s: %s", name, valName, testName, err.Error())
			}

			var valid bool
			var message string
			if val.typeFile != "" {
				typ := unmarshalType(t, name, valName, testName, val.typeFile)
				valid, message, err = typ.Validate(frizz.Stack{frizz.RootItem("root")}, iface)
			} else {
				mi := mockImporter{
					path:    "frizz.io/tests/packer",
					packers: []frizz.Packer{packer.Packer, system.Packer, validators.Packer},
					typers:  map[string]frizz.Typer{},
				}
				for typPkg, m := range val.types {
					mi.typers[typPkg] = mockTyper{path: typPkg, types: m}
				}
				valid, message, err = validator.Validate(iface, mi)
			}

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

func unmarshalType(t *testing.T, name, val, test, in string) system.Type {
	r := &frizz.Root{
		// packer.Imports does not automatically import system or validators, so we must register manually
		Context: frizz.New(packer.Imports).Register(system.Packer, validators.Packer),
		Imports: make(map[string]string),
	}
	v, err := unmarshal(in)
	if err != nil {
		t.Fatalf("%s - %s - %s: unmarshaling typeFile: %s", name, val, test, err.Error())
	}
	if err := r.ParseImports(v); err != nil {
		t.Fatalf("%s - %s - %s: parsing typeFile imports: %s", name, val, test, err.Error())
	}
	typ, _, err := system.Packer.UnpackType(r, frizz.Stack{frizz.RootItem("root")}, v)
	if err != nil {
		t.Fatalf("%s - %s - %s: unpacking typeFile: %s", name, val, test, err.Error())
	}
	return typ
}

func unmarshal(in string) (interface{}, error) {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer([]byte(in)))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

type valDef struct {
	types    map[string]map[string]string
	typeFile string // type
	tests    map[string]testDef
}

type testDef struct {
	data string // json input
	msg  string // expect invalid with this message
	err  string // expect error containing this string
}

type mockImporter struct {
	path    string
	packers []frizz.Packer
	typers  map[string]frizz.Typer
}

func (m mockImporter) Path() string {
	return m.path
}

func (m mockImporter) Add(p map[string]frizz.Packer, t map[string]frizz.Typer) {
	if p != nil {
		for _, v := range m.packers {
			p[v.Path()] = v
		}
	}
	if t != nil {
		for k, v := range m.typers {
			t[k] = v
		}
	}
}

type mockTyper struct {
	path  string
	types map[string]string
}

func (m mockTyper) Path() string {
	return m.path
}

func (m mockTyper) Get(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(m.types[name]))
}
