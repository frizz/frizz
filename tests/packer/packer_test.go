package packer

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"frizz.io/frizz"
	"frizz.io/tests/packer/sub"

	"go/ast"

	"github.com/pkg/errors"
)

type test struct {
	js    string      // json input
	ex    interface{} // expected unpacked output
	re    string      // expected repacked output (if different from js)
	err   string      // if err != "", there should be an error, and will contain this string
	dict  bool        // should the repacked output be a dict?
	unull bool        // should the unpacked output be null?
	rnull bool        // should the repacked output be null?
}

func TestCustomSub(t *testing.T) {
	tests := map[string]test{
		"custom sub": {
			js: `{"String": "a"}`,
			ex: CustomSub(sub.Sub{String: "a-b"}),
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackCustomSub(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackCustomSub(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAges(t *testing.T) {
	tests := map[string]test{
		"ages": {
			js: `"dave:39,john:37"`,
			ex: Ages{"dave": 39, "john": 37},
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAges(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAges(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestCsv(t *testing.T) {
	tests := map[string]test{
		"csv": {
			js: `"1,2,3"`,
			ex: Csv{1, 2, 3},
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackCsv(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackCsv(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestType(t *testing.T) {
	tests := map[string]test{
		"type qual": {
			js: `"foo.Bar"`,
			ex: Type{Path: "github.com/foo", Name: "Bar"},
		},
		"type ident": {
			js: `"Foo"`,
			ex: Type{Path: "frizz.io/tests/packer", Name: "Foo"},
		},
	}
	for name, test := range tests {
		if name != "type ident" {
			continue
		}
		v := decode(t, name, test.js)

		r, s := root()
		r.Imports["foo"] = "github.com/foo"
		unpacked, null1, err1 := Packer.UnpackType(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackType(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestCustom(t *testing.T) {
	tests := map[string]test{
		"custom ident": {
			js: `"Foo"`,
			ex: Custom{&ast.Ident{NamePos: 1, Name: "Foo", Obj: &ast.Object{}}},
		},
		"custom sel": {
			js: `"foo.Bar"`,
			ex: Custom{&ast.SelectorExpr{
				X:   &ast.Ident{NamePos: 1, Name: "foo", Obj: &ast.Object{}},
				Sel: &ast.Ident{NamePos: 5, Name: "Bar", Obj: nil},
			}},
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackCustom(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackCustom(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}


func TestImportsFail(t *testing.T) {
	tests := map[string]test{
		"imports": {
			js: `{
				"_import": {"sub": "frizz.io/tests/packer/sub"},
				"_type": "sub.Sub",
				"String": "a"
			}`,
			err: "packer for frizz.io/tests/packer not registered",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		u := frizz.New(Imports)
		unpacked, null1, err1 := u.Unpack(v)
		repacked, dict, null2, err2 := u.Repack(unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestImports(t *testing.T) {
	tests := map[string]test{
		"imports": {
			js: `{
				"_import": {"sub": "frizz.io/tests/packer/sub"},
				"_type": "sub.Sub",
				"String": "a"
			}`,
			ex: sub.Sub{String: "a"},
		},
		"imports not map": {
			js:  `{"_type": "Natives", "_import": []}`,
			err: "unpacking into interface, _import should be a map",
		},
		"imports values": {
			js:  `{"_type": "Natives", "_import": {"foo": 1}}`,
			err: "unpacking into interface, _import values should be strings",
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		u := frizz.New(Imports)
		u.Register(sub.Packer)
		unpacked, null1, err1 := u.Unpack(v)
		repacked, dict, null2, err2 := u.Repack(unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestInterfaceField(t *testing.T) {
	tests := map[string]test{
		"interface field": {
			js: `{
				"Iface": {"_type": "Impi", "Int": 1}
			}`,
			ex: InterfaceField{Iface: Impi{Int: 1}},
		},
		"interface field slice": {
			js: `{
				"Slice": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			ex: InterfaceField{Slice: []Interface{Impi{Int: 1}, Imps{String: "a"}}},
		},
		"interface field array": {
			js: `{
				"Array": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			ex: InterfaceField{Array: [3]Interface{Impi{Int: 1}, Imps{String: "a"}}},
			re: `{
				"Array": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"},
					null
				]
			}`,
			// TODO: trim null items from end of slices
		},
		"interface field array too long": {
			js: `{
				"Array": [
					{"_type": "Impi", "Int": 1},
					{"_type": "Impi", "Int": 1},
					{"_type": "Impi", "Int": 1},
					{"_type": "Imps", "String": "a"}
				]
			}`,
			err: "data length 4 does not fit in array of length 3",
		},
		"interface field map": {
			js: `{
				"Map": {
					"a": {"_type": "Impi", "Int": 1},
					"b": {"_type": "Imps", "String": "a"}
				}
			}`,
			ex: InterfaceField{Map: map[string]Interface{"a": Impi{Int: 1}, "b": Imps{String: "a"}}},
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackInterfaceField(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackInterfaceField(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestUnpackInterface(t *testing.T) {
	tests := map[string]test{
		"interface string":           {js: `{"_type": "string", "_value": "a"}`, ex: "a"},
		"interface natives":          {js: `{"_type": "Natives", "_value": {"String": "a"}}`, ex: Natives{String: "a"}, re: `{"_type": "Natives", "String": "a"}`},
		"interface sub":              {js: `{"_type": "sub.Sub", "_value": {"String": "a"}}`, ex: sub.Sub{String: "a"}, re: `{"_type": "sub.Sub", "String": "a"}`},
		"interface natives no value": {js: `{"_type": "Natives", "String": "a"}`, ex: Natives{String: "a"}},
		"interface sub no value":     {js: `{"_type": "sub.Sub", "String": "a"}`, ex: sub.Sub{String: "a"}},
		"interface type not found":   {js: `{"_type": "Foo"}`, err: "type Foo not found"},
		"interface no _type":         {js: `{"foo": "bar"}`, err: "unpacking into interface, _type field missing"},
		"interface not map":          {js: `[1, 2]`, err: "unpacking into interface, value should be a map"},
		"interface _type wrong type": {js: `{"_type": true}`, err: "unpacking into interface, _type should be a string"},
		"interface parsing type":     {js: `{"_type": "^"}`, err: "1:2: expected operand, found 'EOF'"},
		"interface missing value":    {js: `{"_type": "int"}`, err: "unpacking native type into interface, _value field missing"},
		"interface sel not string":   {js: `{"_type": "(1).Foo"}`, err: "parsing reference (1).Foo, SelectorExpr.X should be *ast.Ident"},
		"interface import not found": {js: `{"_type": "foo.Foo"}`, err: "parsing reference foo.Foo, can't find foo in imports"},
		"interface qual not found":   {js: `{"_type": "sub.Foo"}`, err: "type Foo not found"},
		"interface unsupported type": {js: `{"_type": "[]foo"}`, err: "unsupported type []foo"},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		r.Imports["sub"] = "frizz.io/tests/packer/sub"
		r.Register(sub.Packer)

		unpacked, null1, err1 := r.UnpackInterface(s, v)
		repacked, dict, null2, err2 := r.RepackInterface(s, false, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestUnpackInterfaceNoPath(t *testing.T) {
	tests := map[string]test{
		"interface local path not set": {js: `{"_type": "Natives"}`, err: "unpacking into interface, local path is not set"},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		r.Path = ""
		unpacked, null1, err1 := r.UnpackInterface(s, v)
		repacked, dict, null2, err2 := r.RepackInterface(s, false, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestPrivate(t *testing.T) {
	tests := map[string]test{
		"private": {js: `{"i": 1, "s": "a"}`, ex: Private{i: 1, s: "a"}},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackPrivate(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackPrivate(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAliasSub(t *testing.T) {
	tests := map[string]test{
		"alias sub": {js: `{"String": "a"}`, ex: AliasSub{String: "a"}},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAliasSub(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAliasSub(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAliasSlice(t *testing.T) {
	tests := map[string]test{
		"alias slice": {js: `[1, 2, 3]`, ex: AliasSlice{1, 2, 3}},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAliasSlice(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAliasSlice(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAliasArray(t *testing.T) {
	tests := map[string]test{
		"alias array":          {js: `["a", "b", "c"]`, ex: AliasArray{"a", "b", "c"}},
		"alias array too long": {js: `["a", "b", "c", "d"]`, err: "data length 4 does not fit in array of length 3"},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAliasArray(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAliasArray(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAliasMap(t *testing.T) {
	tests := map[string]test{
		"alias map": {
			js: `{"a": {"Sub": {"String": "a"}}, "b": {"Sub": {"String": "b"}}}`,
			ex: AliasMap{
				"a": &Qual{Sub: sub.Sub{String: "a"}},
				"b": &Qual{Sub: sub.Sub{String: "b"}},
			},
			dict: true,
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAliasMap(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAliasMap(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAliasPointe(t *testing.T) {
	tests := map[string]test{
		"alias pointer": {js: `4`, ex: AliasPointer(func() *Int { i := Int(4); return &i }())},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAliasPointer(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAliasPointer(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestAlias(t *testing.T) {
	tests := map[string]test{
		"alias": {js: `3`, ex: Alias(3)},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackAlias(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackAlias(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestInt(t *testing.T) {
	tests := map[string]test{
		"int": {js: `2`, ex: Int(2)},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackInt(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackInt(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestString(t *testing.T) {
	tests := map[string]test{
		"string": {js: `"a"`, ex: String("a")},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackString(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackString(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestQual(t *testing.T) {
	tests := map[string]test{
		"qual": {js: `{"Sub": {"String": "a"}}`, ex: Qual{Sub: sub.Sub{String: "a"}}},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackQual(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackQual(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestPointers(t *testing.T) {
	tests := map[string]test{
		"pointers": {
			js: `{
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
			ex: Pointers{
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
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackPointers(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackPointers(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestMaps(t *testing.T) {
	tests := map[string]test{
		"maps": {
			js: `{
				"Ints": {"a": 1, "b": 2, "c": 3},
				"Strings": {"a": "b", "c": "d", "e": "f"},
				"Slices": {"a": ["b", "c"], "d": ["e", "f", "g"]},
				"Arrays": {"a": [1, 2], "b": [3, 4]},
				"Maps": {"a": {"b": "c"}, "d": {"e": "f"}}
			}`,
			ex: Maps{
				Ints:    map[string]int{"a": 1, "b": 2, "c": 3},
				Strings: map[string]string{"a": "b", "c": "d", "e": "f"},
				Slices:  map[string][]string{"a": {"b", "c"}, "d": {"e", "f", "g"}},
				Arrays:  map[string][2]int{"a": {1, 2}, "b": {3, 4}},
				Maps:    map[string]map[string]string{"a": {"b": "c"}, "d": {"e": "f"}},
			},
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackMaps(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackMaps(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestSlices(t *testing.T) {
	tests := map[string]test{
		"slices": {
			js: `{
				"Ints": [1, 2, 3],
				"Strings": ["a", "b", "c"],
				"ArrayLit": ["a", "b", "c", "d", "e"],
				"ArrayExpr": [1, 2, 3, 4],
				"Structs": [{"Int": 1}, {"Int": 2}],
				"Arrays": [["a", "b"], ["c", "d"]]
			}`,
			ex: Slices{
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
		},
		"array lit too long":  {js: `{"ArrayLit": ["a", "b", "c", "d", "e", "f"]}`, err: "data length 6 does not fit in array of length 5"},
		"array expr too long": {js: `{"ArrayExpr": [1, 2, 3, 4, 5]}`, err: "data length 5 does not fit in array of length 4"},
		"wrong type map":      {js: `{"Strings": {"a": "b"}}`, err: "unpacking into slice, value should be an array"},
		"wrong type int":      {js: `{"Ints": 1}`, err: "unpacking into slice, value should be an array"},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackSlices(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackSlices(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestStructs(t *testing.T) {
	tests := map[string]test{
		"structs": {
			js: `{
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
			ex: Structs{
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
		},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackStructs(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackStructs(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func TestNatives(t *testing.T) {
	tests := map[string]test{
		"int":                {js: `{"Int": 1}`, ex: Natives{Int: 1}},
		"int zero":           {js: `{"Int": 0}`, ex: Natives{Int: 0}, re: "{}", rnull: true},
		"int negative":       {js: `{"Int": -1}`, ex: Natives{Int: -1}},
		"int8 small":         {js: `{"Int8": -128}`, ex: Natives{Int8: -128}},
		"int8 big":           {js: `{"Int8": 127}`, ex: Natives{Int8: 127}},
		"int8 zero":          {js: `{"Int8": 0}`, ex: Natives{Int8: 0}, re: "{}", rnull: true},
		"int16 small":        {js: `{"Int16": -32768}`, ex: Natives{Int16: -32768}},
		"int16 big":          {js: `{"Int16": 32767}`, ex: Natives{Int16: 32767}},
		"int32 small":        {js: `{"Int32": -2147483648}`, ex: Natives{Int32: -2147483648}},
		"int32 big":          {js: `{"Int32": 2147483647}`, ex: Natives{Int32: 2147483647}},
		"int64 small":        {js: `{"Int64": -9223372036854775808}`, ex: Natives{Int64: -9223372036854775808}},
		"int64 big":          {js: `{"Int64": 9223372036854775807}`, ex: Natives{Int64: 9223372036854775807}},
		"uint":               {js: `{"Uint": 1}`, ex: Natives{Uint: 1}},
		"uint8 big":          {js: `{"Uint8": 255}`, ex: Natives{Uint8: 255}},
		"uint16 big":         {js: `{"Uint16": 65535}`, ex: Natives{Uint16: 65535}},
		"uint32 big":         {js: `{"Uint32": 4294967295}`, ex: Natives{Uint32: 4294967295}},
		"uint64 big":         {js: `{"Uint64": 18446744073709551615}`, ex: Natives{Uint64: 18446744073709551615}},
		"string":             {js: `{"String": "a"}`, ex: Natives{String: "a"}},
		"rune":               {js: `{"Rune": "😀"}`, ex: Natives{Rune: '😀'}},
		"bool true":          {js: `{"Bool": true}`, ex: Natives{Bool: true}},
		"bool false":         {js: `{"Bool": false}`, ex: Natives{Bool: false}, re: "{}", rnull: true},
		"byte":               {js: `{"Byte": 123}`, ex: Natives{Byte: byte(123)}},
		"float32":            {js: `{"Float32": 1.5}`, ex: Natives{Float32: 1.5}},
		"float32 negative":   {js: `{"Float32": -1.5}`, ex: Natives{Float32: -1.5}},
		"float32 zero":       {js: `{"Float32": 0}`, ex: Natives{Float32: 0}, re: "{}", rnull: true},
		"float64":            {js: `{"Float64": 1.5}`, ex: Natives{Float64: 1.5}},
		"float64 negative":   {js: `{"Float64": -1.5}`, ex: Natives{Float64: -1.5}},
		"int fraction":       {js: `{"Int": 1.5}`, err: `strconv.ParseInt: parsing "1.5": invalid syntax`},
		"int string":         {js: `{"Int": "a"}`, err: `unpacking into int, value "a" should be json.Number`},
		"int bool":           {js: `{"Int": false}`, err: `unpacking into int, value false should be json.Number`},
		"int8 too small":     {js: `{"Int8": -129}`, err: `strconv.ParseInt: parsing "-129": value out of range`},
		"int8 too big":       {js: `{"Int8": 128}`, err: `strconv.ParseInt: parsing "128": value out of range`},
		"int8 wrong type":    {js: `{"Int8": "a"}`, err: `unpacking into int8, value "a" should be json.Number`},
		"int16 too small":    {js: `{"Int16": -32769}`, err: `strconv.ParseInt: parsing "-32769": value out of range`},
		"int16 too big":      {js: `{"Int16": 32768}`, err: `strconv.ParseInt: parsing "32768": value out of range`},
		"int16 wrong type":   {js: `{"Int16": "a"}`, err: `unpacking into int16, value "a" should be json.Number`},
		"int32 too small":    {js: `{"Int32": -2147483649}`, err: `strconv.ParseInt: parsing "-2147483649": value out of range`},
		"int32 too big":      {js: `{"Int32": 2147483648}`, err: `strconv.ParseInt: parsing "2147483648": value out of range`},
		"int32 wrong type":   {js: `{"Int32": "a"}`, err: `unpacking into int32, value "a" should be json.Number`},
		"int64 too small":    {js: `{"Int64": -9223372036854775809}`, err: `strconv.ParseInt: parsing "-9223372036854775809": value out of range`},
		"int64 too big":      {js: `{"Int64": 9223372036854775808}`, err: `strconv.ParseInt: parsing "9223372036854775808": value out of range`},
		"int64 wrong type":   {js: `{"Int64": "a"}`, err: `unpacking into int64, value "a" should be json.Number`},
		"uint negative":      {js: `{"Uint": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint wrong type":    {js: `{"Uint": "a"}`, err: `unpacking into uint, value "a" should be json.Number`},
		"uint8 negative":     {js: `{"Uint8": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint8 too big":      {js: `{"Uint8": 256}`, err: `strconv.ParseUint: parsing "256": value out of range`},
		"uint8 wrong type":   {js: `{"Uint8": "a"}`, err: `unpacking into uint8, value "a" should be json.Number`},
		"uint16 negative":    {js: `{"Uint16": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint16 too big":     {js: `{"Uint16": 65536}`, err: `strconv.ParseUint: parsing "65536": value out of range`},
		"uint16 wrong type":  {js: `{"Uint16": "a"}`, err: `unpacking into uint16, value "a" should be json.Number`},
		"uint32 negative":    {js: `{"Uint32": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint32 too big":     {js: `{"Uint32": 4294967296}`, err: `strconv.ParseUint: parsing "4294967296": value out of range`},
		"uint32 wrong type":  {js: `{"Uint32": "a"}`, err: `unpacking into uint32, value "a" should be json.Number`},
		"uint64 negative":    {js: `{"Uint64": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint64 too big":     {js: `{"Uint64": 18446744073709551616}`, err: `strconv.ParseUint: parsing "18446744073709551616": value out of range`},
		"uint64 wrong type":  {js: `{"Uint64": "a"}`, err: `unpacking into uint64, value "a" should be json.Number`},
		"string number":      {js: `{"String": 1.0}`, err: `unpacking into string, value "1.0" should be string`},
		"string bool":        {js: `{"String": false}`, err: `unpacking into string, value false should be string`},
		"rune wrong type":    {js: `{"Rune": 1}`, err: `unpacking into rune, value "1" should be string`},
		"rune too long":      {js: `{"Rune": "aa"}`, err: `unpacking into rune: string should have a single rune`},
		"bool number":        {js: `{"Bool": 1.0}`, err: `unpacking into bool, value "1.0" should be bool`},
		"bool string":        {js: `{"Bool": "a"}`, err: `unpacking into bool, value "a" should be bool`},
		"byte negative":      {js: `{"Byte": -1}`, err: `strconv.ParseUint: parsing "-1": invalid syntax`},
		"byte too big":       {js: `{"Byte": 256}`, err: `strconv.ParseUint: parsing "256": value out of range`},
		"byte wrong type":    {js: `{"Byte": "a"}`, err: `unpacking into byte, value "a" should be json.Number`},
		"float32 wrong type": {js: `{"Float32": "a"}`, err: `unpacking into float32, value "a" should be json.Number`},
		"float32 too big":    {js: `{"Float32": 1e99}`, err: `strconv.ParseFloat: parsing "1e99": value out of range`},
		"float32 too small":  {js: `{"Float32": -1e99}`, err: `strconv.ParseFloat: parsing "-1e99": value out of range`},
		"float64 wrong type": {js: `{"Float64": "a"}`, err: `unpacking into float64, value "a" should be json.Number`},
		"float64 too big":    {js: `{"Float64": 1e999}`, err: `strconv.ParseFloat: parsing "1e999": value out of range`},
		"float64 too small":  {js: `{"Float64": -1e999}`, err: `strconv.ParseFloat: parsing "-1e999": value out of range`},
		"ptr string":         {js: `{"PtrString": "a"}`, ex: Natives{PtrString: func() *string { v := "a"; return &v }()}},
		"ptr string empty":   {js: `{"PtrString": ""}`, ex: Natives{PtrString: func() *string { v := ""; return &v }()}},
		"ptr int":            {js: `{"PtrInt": 1}`, ex: Natives{PtrInt: func() *int { v := 1; return &v }()}},
		"ptr int zero":       {js: `{"PtrInt": 0}`, ex: Natives{PtrInt: func() *int { v := 0; return &v }()}},
		"ptr int null":       {js: `{"PtrInt": null}`, ex: Natives{PtrInt: nil}, re: "{}", rnull: true},
	}
	for name, test := range tests {
		v := decode(t, name, test.js)

		r, s := root()
		unpacked, null1, err1 := Packer.UnpackNatives(r, s, v)
		repacked, dict, null2, err2 := Packer.RepackNatives(r, s, unpacked)

		ensure(t, name, test, unpacked, null1, err1, repacked, dict, null2, err2)
	}
}

func root() (*frizz.Root, frizz.Stack) {
	r := &frizz.Root{
		Context: frizz.New(Imports),
		Imports: make(map[string]string),
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

func ensure(t *testing.T, name string, test test, unpacked interface{}, null1 bool, err1 error, repacked interface{}, dict bool, null2 bool, err2 error) {
	if test.err != "" {
		if err1 == nil && err2 == nil {
			t.Fatalf("%s: expected error '%s', got nil", name, test.err)
		}
		err1match := err1 != nil && strings.Contains(errors.Cause(err1).Error(), test.err)
		err2match := err2 != nil && strings.Contains(errors.Cause(err2).Error(), test.err)
		if !err1match && !err2match {
			t.Fatalf("%s: expected error '%s', got '%s' and '%s'", name, test.err, errors.Cause(err1), errors.Cause(err2))
		}
	} else {
		if err1 != nil {
			t.Fatalf("%s: error unpacking: %s", name, err1)
		}
		if err2 != nil {
			t.Fatalf("%s: error repacking: %s", name, err2)
		}
		if !reflect.DeepEqual(unpacked, test.ex) {
			t.Fatalf("%s: unpacked result: %#v not what expected: %#v", name, unpacked, test.ex)
		}
		var v interface{}
		if test.re != "" {
			v = decode(t, name, test.re)
		} else {
			v = decode(t, name, test.js)
		}
		if !reflect.DeepEqual(repacked, v) {
			t.Fatalf("%s: repacked result: %#v not what expected: %#v", name, repacked, v)
		}
		if dict != test.dict {
			t.Fatalf("%s: repacked dict = %v, expected %v", name, dict, test.dict)
		}
		if null1 != test.unull {
			t.Fatalf("%s: unpacked null = %v, expected %v", name, null1, test.unull)
		}
		if null2 != test.rnull {
			t.Fatalf("%s: repacked null = %v, expected %v", name, null2, test.rnull)
		}
	}
}
