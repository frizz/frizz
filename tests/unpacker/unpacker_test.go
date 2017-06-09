package unpacker

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"frizz.io/frizz"
	"frizz.io/tests/unpacker/sub"

	"go/ast"

	"github.com/pkg/errors"
)

type test struct {
	json     string
	expected interface{}
	error    string
}

func TestType(t *testing.T) {
	tests := map[string]test{
		"type qual": {
			`"foo.Bar"`,
			Type{Path: "github.com/foo", Name: "Bar"},
			"",
		},
		"type ident": {
			`"Foo"`,
			Type{Path: "frizz.io/tests/unpacker", Name: "Foo"},
			"",
		},
	}
	for name, test := range tests {
		if name != "type ident" {
			continue
		}
		v := decode(t, name, test.json)

		r, s := root()
		r.Imports["foo"] = "github.com/foo"
		result, err := Unpackers.Type(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestCustom(t *testing.T) {
	tests := map[string]test{
		"custom ident": {
			`"Foo"`,
			&ast.Ident{NamePos: 1, Name: "Foo", Obj: &ast.Object{}},
			"",
		},
		"custom sel": {
			`"foo.Bar"`,
			&ast.SelectorExpr{
				X:   &ast.Ident{NamePos: 1, Name: "foo", Obj: &ast.Object{}},
				Sel: &ast.Ident{NamePos: 5, Name: "Bar", Obj: nil},
			},
			""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Custom(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestImports(t *testing.T) {
	tests := map[string]test{
		"imports": {
			`{
				"_import": {"sub": "frizz.io/tests/unpacker/sub"},
				"_type": "sub.Sub",
				"String": "a"
			}`,
			sub.Sub{String: "a"},
			"",
		},
		"imports not map": {`{"_type": "Natives", "_import": []}`, nil, "unpacking into interface, _import should be a map"},
		"imports values":  {`{"_type": "Natives", "_import": {"foo": 1}}`, nil, "unpacking into interface, _import values should be strings"},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		u := frizz.New("frizz.io/tests/unpacker")
		result, err := u.Unpack(v)

		ensure(t, name, test, err, result)
	}
}

func TestInterfaceField(t *testing.T) {
	tests := map[string]test{
		"interface field": {
			`{
				"Iface": {"_type": "Impi", "Int": 1}
			}`,
			InterfaceField{Iface: Impi{Int: 1}},
			"",
		},
		"interface field slice": {
			`{
				"Slice": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			InterfaceField{Slice: []Interface{Impi{Int: 1}, Imps{String: "a"}}},
			"",
		},
		"interface field array": {
			`{
				"Array": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			InterfaceField{Array: [3]Interface{Impi{Int: 1}, Imps{String: "a"}}},
			"",
		},
		"interface field array too long": {
			`{
				"Array": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Impi", "Int": 1},
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			nil,
			"data length 4 does not fit in array of length 3",
		},
		"interface field map": {
			`{
				"Map": {
					"a": {"_type": "Impi", "Int": 1},
					"b": {"_type": "Imps", "String": "a"}
				}
			}`,
			InterfaceField{Map: map[string]Interface{"a": Impi{Int: 1}, "b": Imps{String: "a"}}},
			"",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.InterfaceField(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestUnpackInterface(t *testing.T) {
	tests := map[string]test{
		"interface string":           {`{"_type": "string", "_value": "a"}`, "a", ""},
		"interface natives":          {`{"_type": "Natives", "_value": {"String": "a"}}`, Natives{String: "a"}, ""},
		"interface sub":              {`{"_type": "sub.Sub", "_value": {"String": "a"}}`, sub.Sub{String: "a"}, ""},
		"interface natives no value": {`{"_type": "Natives", "String": "a"}`, Natives{String: "a"}, ""},
		"interface sub no value":     {`{"_type": "sub.Sub", "String": "a"}`, sub.Sub{String: "a"}, ""},
		"interface type not found":   {`{"_type": "Foo"}`, nil, "unpacking into interface, can't find frizz.io/tests/unpacker.Foo in type registry"},
		"interface no _type":         {`{"foo": "bar"}`, nil, "unpacking into interface, _type field missing"},
		"interface not map":          {`[1, 2]`, nil, "unpacking into interface, value should be a map"},
		"interface _type wrong type": {`{"_type": true}`, nil, "unpacking into interface, _type should be a string"},
		"interface parsing type":     {`{"_type": "^"}`, nil, "1:2: expected operand, found 'EOF'"},
		"interface missing value":    {`{"_type": "int"}`, nil, "unpacking native type into interface, _value field missing"},
		"interface sel not string":   {`{"_type": "(1).Foo"}`, nil, "unpacking into interface, SelectorExpr.X should be *ast.Ident"},
		"interface import not found": {`{"_type": "foo.Foo"}`, nil, "unpacking into interface, can't find foo in imports"},
		"interface qual not found":   {`{"_type": "sub.Foo"}`, nil, "unpacking into interface, can't find frizz.io/tests/unpacker/sub.Foo in registry"},
		"interface unsupported type": {`{"_type": "[]foo"}`, nil, "unsupported type []foo"},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		r.Imports["sub"] = "frizz.io/tests/unpacker/sub"
		result, err := r.UnpackInterface(s, v)

		ensure(t, name, test, err, result)
	}
}

func TestUnpackInterfaceNoPath(t *testing.T) {
	tests := map[string]test{
		"interface local path not set": {`{"_type": "Natives"}`, nil, "unpacking into interface, local path is not set"},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		r.Path = ""
		result, err := r.UnpackInterface(s, v)

		ensure(t, name, test, err, result)
	}
}

func TestPrivate(t *testing.T) {
	tests := map[string]test{
		"private": {`{"i": 1, "s": "a"}`, Private{i: 1, s: "a"}, ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Private(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAliasSub(t *testing.T) {
	tests := map[string]test{
		"alias sub": {`{"String": "a"}`, AliasSub{String: "a"}, ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.AliasSub(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAliasSlice(t *testing.T) {
	tests := map[string]test{
		"alias slice": {`[1, 2, 3]`, AliasSlice{1, 2, 3}, ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.AliasSlice(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAliasArray(t *testing.T) {
	tests := map[string]test{
		"alias array":          {`["a", "b", "c"]`, AliasArray{"a", "b", "c"}, ""},
		"alias array too long": {`["a", "b", "c", "d"]`, nil, "data length 4 does not fit in array of length 3"},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.AliasArray(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAliasMap(t *testing.T) {
	tests := map[string]test{
		"alias map": {
			`{"a": {"Sub": {"String": "a"}}, "b": {"Sub": {"String": "b"}}}`,
			AliasMap{
				"a": &Qual{Sub: sub.Sub{String: "a"}},
				"b": &Qual{Sub: sub.Sub{String: "b"}},
			},
			"",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.AliasMap(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAliasPointe(t *testing.T) {
	tests := map[string]test{
		"alias pointer": {`4`, AliasPointer(func() *Int { i := Int(4); return &i }()), ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.AliasPointer(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestAlias(t *testing.T) {
	tests := map[string]test{
		"alias": {`3`, Alias(3), ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Alias(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestInt(t *testing.T) {
	tests := map[string]test{
		"int": {`2`, Int(2), ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Int(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestString(t *testing.T) {
	tests := map[string]test{
		"string": {`"a"`, String("a"), ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.String(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestQual(t *testing.T) {
	tests := map[string]test{
		"qual": {`{"Sub": {"String": "a"}}`, Qual{Sub: sub.Sub{String: "a"}}, ""},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Qual(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestPointers(t *testing.T) {
	tests := map[string]test{
		"pointers": {
			`{
				"String": "a",
				"Int": 2,
				"Sub": {"String": "a"},
				"Array": [1, 2, 3],
				"Slice": ["a", "b", "c"],
				"Map": {"a": 1, "b": 2},
				"SliceString": ["a"],
				"SliceInt": [1],
				"SliceSub": [{"String": "a"}]
			}`,
			Pointers{
				String:      func() *string { s := "a"; return &s }(),
				Int:         func() *Int { i := Int(2); return &i }(),
				Sub:         &sub.Sub{String: "a"},
				Array:       &[3]int{1, 2, 3},
				Slice:       &[]string{"a", "b", "c"},
				Map:         &map[string]int{"a": 1, "b": 2},
				SliceString: []*string{func() *string { s := "a"; return &s }()},
				SliceInt:    []*Int{func() *Int { i := Int(1); return &i }()},
				SliceSub:    []*sub.Sub{{String: "a"}},
			},
			"",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Pointers(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestMaps(t *testing.T) {
	tests := map[string]test{
		"maps": {
			`{
				"Ints": {"a": 1, "b": 2, "c": 3},
				"Strings": {"a": "b", "c": "d", "e": "f"},
				"Slices": {"a": ["b", "c"], "d": ["e", "f", "g"]},
				"Arrays": {"a": [1, 2], "b": [3, 4]},
				"Maps": {"a": {"b": "c"}, "d": {"e": "f"}}
			}`,
			Maps{
				Ints:    map[string]int{"a": 1, "b": 2, "c": 3},
				Strings: map[string]string{"a": "b", "c": "d", "e": "f"},
				Slices:  map[string][]string{"a": {"b", "c"}, "d": {"e", "f", "g"}},
				Arrays:  map[string][2]int{"a": {1, 2}, "b": {3, 4}},
				Maps:    map[string]map[string]string{"a": {"b": "c"}, "d": {"e": "f"}},
			},
			"",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Maps(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestSlices(t *testing.T) {
	tests := map[string]test{
		"slices": {
			`{
				"Ints": [1, 2, 3],
				"Strings": ["a", "b", "c"],
				"ArrayLit": ["a", "b", "c", "d", "e"],
				"ArrayExpr": [1, 2, 3, 4],
				"Structs": [{"Int": 1}, {"Int": 2}],
				"Arrays": [["a", "b"], ["c", "d"]]
			}`,
			Slices{
				Ints:      []int{1, 2, 3},
				Strings:   []string{"a", "b", "c"},
				ArrayLit:  [5]string{"a", "b", "c", "d", "e"},
				ArrayExpr: [2 * N]int{1, 2, 3, 4},
				Structs: []struct{ Int int }{
					{Int: 1},
					{Int: 2},
				},
				Arrays: [][]string{
					{"a", "b"},
					{"c", "d"},
				},
			},
			"",
		},
		"array lit too long":  {`{"ArrayLit": ["a", "b", "c", "d", "e", "f"]}`, nil, "data length 6 does not fit in array of length 5"},
		"array expr too long": {`{"ArrayExpr": [1, 2, 3, 4, 5]}`, nil, "data length 5 does not fit in array of length 4"},
		"wrong type map":      {`{"Strings": {"a": "b"}}`, nil, "unpacking into slice, value should be an array"},
		"wrong type int":      {`{"Ints": 1}`, nil, "unpacking into slice, value should be an array"},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Slices(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestStructs(t *testing.T) {
	tests := map[string]test{
		"structs": {
			`{
				"Simple": {
					"Int": 1, 
					"Bool": true
				}, 
				"Complex": {
					"String": "a", 
					"Inner": {
						"Float32": 1.1
					}
				}
			}`,
			Structs{
				Simple: struct {
					Int  int
					Bool bool
				}{
					Int:  1,
					Bool: true,
				},
				Complex: struct {
					String string
					Inner  struct {
						Float32 float32
					}
				}{
					String: "a",
					Inner: struct {
						Float32 float32
					}{
						Float32: 1.1,
					},
				},
			},
			"",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Structs(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func TestNatives(t *testing.T) {
	tests := map[string]test{
		"int":                {`{"Int": 1}`, Natives{Int: 1}, ""},
		"int negative":       {`{"Int": -1}`, Natives{Int: -1}, ""},
		"int8 small":         {`{"Int8": -128}`, Natives{Int8: -128}, ""},
		"int8 big":           {`{"Int8": 127}`, Natives{Int8: 127}, ""},
		"int16 small":        {`{"Int16": -32768}`, Natives{Int16: -32768}, ""},
		"int16 big":          {`{"Int16": 32767}`, Natives{Int16: 32767}, ""},
		"int32 small":        {`{"Int32": -2147483648}`, Natives{Int32: -2147483648}, ""},
		"int32 big":          {`{"Int32": 2147483647}`, Natives{Int32: 2147483647}, ""},
		"int64 small":        {`{"Int64": -9223372036854775808}`, Natives{Int64: -9223372036854775808}, ""},
		"int64 big":          {`{"Int64": 9223372036854775807}`, Natives{Int64: 9223372036854775807}, ""},
		"uint":               {`{"Uint": 1}`, Natives{Uint: 1}, ""},
		"uint8 big":          {`{"Uint8": 255}`, Natives{Uint8: 255}, ""},
		"uint16 big":         {`{"Uint16": 65535}`, Natives{Uint16: 65535}, ""},
		"uint32 big":         {`{"Uint32": 4294967295}`, Natives{Uint32: 4294967295}, ""},
		"uint64 big":         {`{"Uint64": 18446744073709551615}`, Natives{Uint64: 18446744073709551615}, ""},
		"string":             {`{"String": "a"}`, Natives{String: "a"}, ""},
		"rune":               {`{"Rune": "😀"}`, Natives{Rune: '😀'}, ""},
		"bool true":          {`{"Bool": true}`, Natives{Bool: true}, ""},
		"bool false":         {`{"Bool": false}`, Natives{Bool: false}, ""},
		"byte":               {`{"Byte": 123}`, Natives{Byte: byte(123)}, ""},
		"float32":            {`{"Float32": 1.5}`, Natives{Float32: 1.5}, ""},
		"float32 negative":   {`{"Float32": -1.5}`, Natives{Float32: -1.5}, ""},
		"float64":            {`{"Float64": 1.5}`, Natives{Float64: 1.5}, ""},
		"float64 negative":   {`{"Float64": -1.5}`, Natives{Float64: -1.5}, ""},
		"int fraction":       {`{"Int": 1.5}`, nil, `strconv.ParseInt: parsing "1.5": invalid syntax`},
		"int string":         {`{"Int": "a"}`, nil, `unpacking into int, value "a" should be json.Number`},
		"int bool":           {`{"Int": false}`, nil, `unpacking into int, value false should be json.Number`},
		"int8 too small":     {`{"Int8": -129}`, nil, `strconv.ParseInt: parsing "-129": value out of range`},
		"int8 too big":       {`{"Int8": 128}`, nil, `strconv.ParseInt: parsing "128": value out of range`},
		"int8 wrong type":    {`{"Int8": "a"}`, nil, `unpacking into int8, value "a" should be json.Number`},
		"int16 too small":    {`{"Int16": -32769}`, nil, `strconv.ParseInt: parsing "-32769": value out of range`},
		"int16 too big":      {`{"Int16": 32768}`, nil, `strconv.ParseInt: parsing "32768": value out of range`},
		"int16 wrong type":   {`{"Int16": "a"}`, nil, `unpacking into int16, value "a" should be json.Number`},
		"int32 too small":    {`{"Int32": -2147483649}`, nil, `strconv.ParseInt: parsing "-2147483649": value out of range`},
		"int32 too big":      {`{"Int32": 2147483648}`, nil, `strconv.ParseInt: parsing "2147483648": value out of range`},
		"int32 wrong type":   {`{"Int32": "a"}`, nil, `unpacking into int32, value "a" should be json.Number`},
		"int64 too small":    {`{"Int64": -9223372036854775809}`, nil, `strconv.ParseInt: parsing "-9223372036854775809": value out of range`},
		"int64 too big":      {`{"Int64": 9223372036854775808}`, nil, `strconv.ParseInt: parsing "9223372036854775808": value out of range`},
		"int64 wrong type":   {`{"Int64": "a"}`, nil, `unpacking into int64, value "a" should be json.Number`},
		"uint negative":      {`{"Uint": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint wrong type":    {`{"Uint": "a"}`, nil, `unpacking into uint, value "a" should be json.Number`},
		"uint8 negative":     {`{"Uint8": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint8 too big":      {`{"Uint8": 256}`, nil, `strconv.ParseUint: parsing "256": value out of range`},
		"uint8 wrong type":   {`{"Uint8": "a"}`, nil, `unpacking into uint8, value "a" should be json.Number`},
		"uint16 negative":    {`{"Uint16": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint16 too big":     {`{"Uint16": 65536}`, nil, `strconv.ParseUint: parsing "65536": value out of range`},
		"uint16 wrong type":  {`{"Uint16": "a"}`, nil, `unpacking into uint16, value "a" should be json.Number`},
		"uint32 negative":    {`{"Uint32": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint32 too big":     {`{"Uint32": 4294967296}`, nil, `strconv.ParseUint: parsing "4294967296": value out of range`},
		"uint32 wrong type":  {`{"Uint32": "a"}`, nil, `unpacking into uint32, value "a" should be json.Number`},
		"uint64 negative":    {`{"Uint64": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint64 too big":     {`{"Uint64": 18446744073709551616}`, nil, `strconv.ParseUint: parsing "18446744073709551616": value out of range`},
		"uint64 wrong type":  {`{"Uint64": "a"}`, nil, `unpacking into uint64, value "a" should be json.Number`},
		"string number":      {`{"String": 1.0}`, nil, `unpacking into string, value "1.0" should be string`},
		"string bool":        {`{"String": false}`, nil, `unpacking into string, value false should be string`},
		"rune wrong type":    {`{"Rune": 1}`, nil, `unpacking into rune, value "1" should be string`},
		"rune too long":      {`{"Rune": "aa"}`, nil, `unpacking into rune: string should have a single rune`},
		"bool number":        {`{"Bool": 1.0}`, nil, `unpacking into bool, value "1.0" should be bool`},
		"bool string":        {`{"Bool": "a"}`, nil, `unpacking into bool, value "a" should be bool`},
		"byte negative":      {`{"Byte": -1}`, nil, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"byte too big":       {`{"Byte": 256}`, nil, `strconv.ParseUint: parsing "256": value out of range`},
		"byte wrong type":    {`{"Byte": "a"}`, nil, `unpacking into byte, value "a" should be json.Number`},
		"float32 wrong type": {`{"Float32": "a"}`, nil, `unpacking into float32, value "a" should be json.Number`},
		"float32 too big":    {`{"Float32": 1e99}`, nil, `strconv.ParseFloat: parsing "1e99": value out of range`},
		"float32 too small":  {`{"Float32": -1e99}`, nil, `strconv.ParseFloat: parsing "-1e99": value out of range`},
		"float64 wrong type": {`{"Float64": "a"}`, nil, `unpacking into float64, value "a" should be json.Number`},
		"float64 too big":    {`{"Float64": 1e999}`, nil, `strconv.ParseFloat: parsing "1e999": value out of range`},
		"float64 too small":  {`{"Float64": -1e999}`, nil, `strconv.ParseFloat: parsing "-1e999": value out of range`},
	}
	for name, test := range tests {
		v := decode(t, name, test.json)

		r, s := root()
		result, err := Unpackers.Natives(r, s, v)

		ensure(t, name, test, err, result)
	}
}

func root() (*frizz.Root, frizz.Stack) {
	r := &frizz.Root{
		Unpacker: frizz.New("frizz.io/tests/unpacker"),
		Imports:  make(map[string]string),
	}
	s := frizz.Stack{frizz.RootItem("root")}
	return r, s
}

func decode(t *testing.T, name, s string) interface{} {
	var v interface{}
	d := json.NewDecoder(bytes.NewBuffer([]byte(s)))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		t.Fatalf("%s error decoding %s", name, err)
	}
	return v
}

func ensure(t *testing.T, name string, test test, err error, result interface{}) {
	if test.error != "" {
		if err == nil {
			t.Fatalf("%s: expected error '%s', got nil", name, test.error)
		}
		if !strings.Contains(errors.Cause(err).Error(), test.error) {
			t.Fatalf("%s: expected error '%s', got %s", name, test.error, errors.Cause(err))
		}
	} else {
		if err != nil {
			t.Fatalf("%s: error unpacking: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("%s: result %#v not what expected", name, result)
		}
	}
}
