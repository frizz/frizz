package validation_test

import (
	"testing"

	"bytes"
	"encoding/base64"
	"encoding/json"

	"frizz.io/frizz"
	"frizz.io/global"
	"frizz.io/pack"
	"frizz.io/system"
	"frizz.io/tests/packer"
	"frizz.io/tests/validation"
	"frizz.io/validator"
	"frizz.io/validators"
)

func TestGeneration(t *testing.T) {
	f := frizz.New(validation.Package)
	vt, err := frizz.Unmarshal(f, []byte(`{
		"_type": "ValidateTest",
		"Int": 3
	}`))
	if err != nil {
		t.Fatal(err)
	}
	valid, message, err := frizz.Validate(f, vt)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("Should be valid")
	}
	if message != "" {
		t.Fatal("Message should be empty")
	}

	vt, err = frizz.Unmarshal(f, []byte(`{
		"_type": "ValidateTest",
		"Int": 2
	}`))
	if err != nil {
		t.Fatal(err)
	}
	valid, message, err = frizz.Validate(f, vt)
	if err != nil {
		t.Fatal(err)
	}
	if valid {
		t.Fatal("Should not be valid")
	}
	expected := `root.Int: value 2 must be greater than 2`
	if message != expected {
		t.Fatalf("Message should be %#v. Found %#v", expected, message)
	}
}

type msss map[string]map[string]string
type mss map[string]string

func TestInterface(t *testing.T) {
	vals := map[string]valDef{
		"validator": {
			typeFile: `{"_type": "system.Type", "_import": {"system": "frizz.io/system", "validators": "frizz.io/validators"}, "Validators": [
				{"_type": "InterfaceValidator", "_value": "1"}
			]}`,
			tests: map[string]testDef{
				"success": {data: `{"_type": "Impi", "Int": 1}`},
				"fail":    {data: `{"_type": "Impi", "Int": 2}`, msg: "root: value \"2\" should be equal to \"1\""},
				"zero":    {data: `{"_type": "Impi", "Int": 0}`, msg: "root: value \"0\" should be equal to \"1\""},
				"missing": {data: `{"_type": "Impi"}`, msg: "root: value \"0\" should be equal to \"1\""},
			},
		},
	}
	run(t, "Interface", vals)
}

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
				"not null": {data: `{"_type": "Pointers", "Int": 1}`, msg: "root.Int: value 1 must be null"},
				"success":  {data: `{"_type": "Pointers"}`},
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
			iface, err := frizz.Unmarshal(frizz.New(packer.Package), []byte(test.data))
			if err != nil {
				t.Fatalf("%s - %s - %s: %s", name, valName, testName, err.Error())
			}

			var valid bool
			var message string
			if val.typeFile != "" {
				typ := unmarshalType(t, name, valName, testName, val.typeFile)
				context := pack.NewValidationContext(frizz.New(packer.Package), global.NewStack("root"))
				valid, message, err = typ.Validate(context, iface)
			} else {
				mi := mockPackage{
					path: packer.Package.Path(),
					packages: map[string]global.Package{
						packer.Package.Path():     mockPackage{path: packer.Package.Path(), inner: packer.Package},
						system.Package.Path():     mockPackage{path: system.Package.Path(), inner: system.Package},
						validators.Package.Path(): mockPackage{path: validators.Package.Path(), inner: validators.Package},
					},
					types: map[string]string{},
					inner: packer.Package,
				}
				for path, types := range val.types {
					if _, ok := mi.packages[path]; ok {
						mp := mi.packages[path].(mockPackage)
						mi.packages[path] = mockPackage{path: path, inner: mp.inner, types: types}
					} else {
						mi.packages[path] = mockPackage{path: path, types: types}
					}
				}
				context := frizz.New(mi)
				valid, message, err = validator.Validate(context, iface)
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
	// packer.Imports does not automatically import system or validators, so we must register manually
	c := frizz.New(packer.Package, system.Package, validators.Package)
	r := pack.NewRootContext(c)
	v, err := unmarshal(in)
	if err != nil {
		t.Fatalf("%s - %s - %s: unmarshaling typeFile: %s", name, val, test, err.Error())
	}
	if err := r.ParseImports(v); err != nil {
		t.Fatalf("%s - %s - %s: parsing typeFile imports: %s", name, val, test, err.Error())
	}
	d := pack.NewDataContext(c, r, global.NewStack("root"))
	typ, _, err := system.Package.UnpackType(d, v)
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

type mockPackage struct {
	path     string
	packages map[string]global.Package
	types    map[string]string
	inner    global.Package
}

func (m mockPackage) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	return m.inner.Unpack(context, in, name)
}

func (m mockPackage) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	return m.inner.Repack(context, in, name)
}

func (m mockPackage) GetType(name string) string {
	return name
}

func (m mockPackage) GetData(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(m.types[name]))
}

func (m mockPackage) GetImportedPackages(packages map[string]global.Package) {
	for _, v := range m.packages {
		packages[v.Path()] = v
	}
}

func (m mockPackage) Path() string {
	return m.path
}
