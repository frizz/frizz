package unpacker

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"frizz.io/tests/unpacker/sub"

	"frizz.io/system"
	"github.com/pkg/errors"
)

func TestInterfaceFieldSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected InterfaceField
	}{
		"interface field": {`
			{"Iface": {"_type": "Impi", "Int": 1}}`,
			InterfaceField{Iface: Impi{Int: 1}},
		},
		"interface slice": {`
			{"Slice": [{"_type": "Impi", "Int": 1}, {"_type": "Imps", "String": "a"}]}`,
			InterfaceField{Slice: []Interface{Impi{Int: 1}, Imps{String: "a"}}},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		ctx := system.NewContext()
		uc := ctx.Value(system.UnpackContextKey).(*system.UnpackContext)
		uc.Path = "frizz.io/tests/unpacker"
		result, err := Unpackers.InterfaceField(ctx, v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestUnpackInterfaceSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected interface{}
	}{
		"imterface string": {`
			{"_type": "string", "_value": "a"}`,
			"a",
		},
		"interface natives": {`
			{"_type": "Natives", "_value": {"String": "a"}}`,
			Natives{String: "a"},
		},
		"interface sub": {`
			{"_type": "sub.Sub", "_value": {"String": "a"}}`,
			sub.Sub{String: "a"},
		},
		"interface natives no value": {`
			{"_type": "Natives", "String": "a"}`,
			Natives{String: "a"},
		},
		"interface sub no value": {`
			{"_type": "sub.Sub", "String": "a"}`,
			sub.Sub{String: "a"},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		ctx := system.NewContext()
		uc := ctx.Value(system.UnpackContextKey).(*system.UnpackContext)
		uc.Path = "frizz.io/tests/unpacker"
		uc.Set("sub", "frizz.io/tests/unpacker/sub")
		result, err := system.UnpackInterface(ctx, v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestPrivateSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Private
	}{
		"private": {`
			{"i": 1, "s": "a"}`,
			Private{i: 1, s: "a"},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Private(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasSubSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected AliasSub
	}{
		"alias sub": {`
			{"String": "a"}`,
			AliasSub{String: "a"},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.AliasSub(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasSliceSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected AliasSlice
	}{
		"alias slice": {`
			[1, 2, 3]`,
			AliasSlice{1, 2, 3},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.AliasSlice(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasArraySuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected AliasArray
	}{
		"alias array": {`
			["a", "b", "c"]`,
			AliasArray{"a", "b", "c"},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.AliasArray(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasArrayFail(t *testing.T) {
	tests := map[string]struct {
		json  string
		error string
	}{
		"alias array too long": {
			`["a", "b", "c", "d"]`,
			"data length 4 does not fit in array of length 3",
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatalf("%s error decoding %s", name, err)
		}
		_, errtest := Unpackers.AliasArray(system.NewContext(), v)
		if errtest == nil {
			t.Fatalf("%s expected error %s, got nil", name, test.error)
		}
		if !strings.Contains(errors.Cause(errtest).Error(), test.error) {
			t.Fatalf("%s error expected to contain '%s': %s", name, test.error, errors.Cause(errtest))
		}
	}
}

func TestAliasMapSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected AliasMap
	}{
		"alias map": {`
			{"a": {"Sub": {"String": "a"}}, "b": {"Sub": {"String": "b"}}}`,
			AliasMap{
				"a": &Qual{Sub: sub.Sub{String: "a"}},
				"b": &Qual{Sub: sub.Sub{String: "b"}},
			},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.AliasMap(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasPointerSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected AliasPointer
	}{
		"alias pointer": {`
			4`,
			AliasPointer(func() *Int { i := Int(4); return &i }()),
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.AliasPointer(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestAliasSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Alias
	}{
		"alias": {`
			3`,
			Alias(3),
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Alias(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestIntSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Int
	}{
		"int": {`
			2`,
			Int(2),
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Int(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestStringSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected String
	}{
		"string": {`
			"a"`,
			String("a"),
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.String(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestQualSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Qual
	}{
		"qual": {`
			{
				"Sub": {"String": "a"}
			}`,
			Qual{
				Sub: sub.Sub{String: "a"},
			},
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Qual(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestPointersSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Pointers
	}{
		"pointers": {`
			{
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
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Pointers(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestMapsSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Maps
	}{
		"maps": {`
			{
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
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Maps(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestSlicesSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Slices
	}{
		"slices": {`
			{
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
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Slices(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestSlicesErrors(t *testing.T) {
	tests := map[string]struct {
		json  string
		error string
	}{
		"array lit too long": {
			`{"ArrayLit": ["a", "b", "c", "d", "e", "f"]}`,
			"data length 6 does not fit in array of length 5",
		},
		"array expr too long": {
			`{"ArrayExpr": [1, 2, 3, 4, 5]}`,
			"data length 5 does not fit in array of length 4",
		},
		"wrong type map": {
			`{"Strings": {"a": "b"}}`,
			"error unpacking into slice, value should be an array",
		},
		"wrong type int": {
			`{"Ints": 1}`,
			"error unpacking into slice, value should be an array",
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatalf("%s error decoding %s", name, err)
		}
		_, errtest := Unpackers.Slices(system.NewContext(), v)
		if errtest == nil {
			t.Fatalf("%s expected error %s, got nil", name, test.error)
		}
		if !strings.Contains(errors.Cause(errtest).Error(), test.error) {
			t.Fatalf("%s error expected to contain '%s': %s", name, test.error, errors.Cause(errtest))
		}
	}
}

func TestStructsSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Structs
	}{
		"structs": {`
			{
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
		},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Structs(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if result != test.expected {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestNativesSuccess(t *testing.T) {
	tests := map[string]struct {
		json     string
		expected Natives
	}{
		"int":              {`{"Int": 1}`, Natives{Int: 1}},
		"int negative":     {`{"Int": -1}`, Natives{Int: -1}},
		"int8 small":       {`{"Int8": -128}`, Natives{Int8: -128}},
		"int8 big":         {`{"Int8": 127}`, Natives{Int8: 127}},
		"int16 small":      {`{"Int16": -32768}`, Natives{Int16: -32768}},
		"int16 big":        {`{"Int16": 32767}`, Natives{Int16: 32767}},
		"int32 small":      {`{"Int32": -2147483648}`, Natives{Int32: -2147483648}},
		"int32 big":        {`{"Int32": 2147483647}`, Natives{Int32: 2147483647}},
		"int64 small":      {`{"Int64": -9223372036854775808}`, Natives{Int64: -9223372036854775808}},
		"int64 big":        {`{"Int64": 9223372036854775807}`, Natives{Int64: 9223372036854775807}},
		"uint":             {`{"Uint": 1}`, Natives{Uint: 1}},
		"uint8 big":        {`{"Uint8": 255}`, Natives{Uint8: 255}},
		"uint16 big":       {`{"Uint16": 65535}`, Natives{Uint16: 65535}},
		"uint32 big":       {`{"Uint32": 4294967295}`, Natives{Uint32: 4294967295}},
		"uint64 big":       {`{"Uint64": 18446744073709551615}`, Natives{Uint64: 18446744073709551615}},
		"string":           {`{"String": "a"}`, Natives{String: "a"}},
		"rune":             {`{"Rune": "😀"}`, Natives{Rune: '😀'}},
		"bool true":        {`{"Bool": true}`, Natives{Bool: true}},
		"bool false":       {`{"Bool": false}`, Natives{Bool: false}},
		"byte":             {`{"Byte": 123}`, Natives{Byte: byte(123)}},
		"float32":          {`{"Float32": 1.5}`, Natives{Float32: 1.5}},
		"float32 negative": {`{"Float32": -1.5}`, Natives{Float32: -1.5}},
		"float64":          {`{"Float64": 1.5}`, Natives{Float64: 1.5}},
		"float64 negative": {`{"Float64": -1.5}`, Natives{Float64: -1.5}},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatal("Error decoding", err)
		}
		result, err := Unpackers.Natives(system.NewContext(), v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		if result != test.expected {
			t.Fatalf("Result %#v not what expected while unpacking %s", result, name)
		}
	}
}

func TestNativesErrors(t *testing.T) {
	tests := map[string]struct {
		json  string
		error string
	}{
		"int fraction":       {`{"Int": 1.5}`, `strconv.ParseInt: parsing "1.5": invalid syntax`},
		"int string":         {`{"Int": "a"}`, `unpacking into int, value "a" should be json.Number`},
		"int bool":           {`{"Int": false}`, `unpacking into int, value false should be json.Number`},
		"int8 too small":     {`{"Int8": -129}`, `strconv.ParseInt: parsing "-129": value out of range`},
		"int8 too big":       {`{"Int8": 128}`, `strconv.ParseInt: parsing "128": value out of range`},
		"int8 wrong type":    {`{"Int8": "a"}`, `unpacking into int8, value "a" should be json.Number`},
		"int16 too small":    {`{"Int16": -32769}`, `strconv.ParseInt: parsing "-32769": value out of range`},
		"int16 too big":      {`{"Int16": 32768}`, `strconv.ParseInt: parsing "32768": value out of range`},
		"int16 wrong type":   {`{"Int16": "a"}`, `unpacking into int16, value "a" should be json.Number`},
		"int32 too small":    {`{"Int32": -2147483649}`, `strconv.ParseInt: parsing "-2147483649": value out of range`},
		"int32 too big":      {`{"Int32": 2147483648}`, `strconv.ParseInt: parsing "2147483648": value out of range`},
		"int32 wrong type":   {`{"Int32": "a"}`, `unpacking into int32, value "a" should be json.Number`},
		"int64 too small":    {`{"Int64": -9223372036854775809}`, `strconv.ParseInt: parsing "-9223372036854775809": value out of range`},
		"int64 too big":      {`{"Int64": 9223372036854775808}`, `strconv.ParseInt: parsing "9223372036854775808": value out of range`},
		"int64 wrong type":   {`{"Int64": "a"}`, `unpacking into int64, value "a" should be json.Number`},
		"uint negative":      {`{"Uint": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint wrong type":    {`{"Uint": "a"}`, `unpacking into uint, value "a" should be json.Number`},
		"uint8 negative":     {`{"Uint8": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint8 too big":      {`{"Uint8": 256}`, `strconv.ParseUint: parsing "256": value out of range`},
		"uint8 wrong type":   {`{"Uint8": "a"}`, `unpacking into uint8, value "a" should be json.Number`},
		"uint16 negative":    {`{"Uint16": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint16 too big":     {`{"Uint16": 65536}`, `strconv.ParseUint: parsing "65536": value out of range`},
		"uint16 wrong type":  {`{"Uint16": "a"}`, `unpacking into uint16, value "a" should be json.Number`},
		"uint32 negative":    {`{"Uint32": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint32 too big":     {`{"Uint32": 4294967296}`, `strconv.ParseUint: parsing "4294967296": value out of range`},
		"uint32 wrong type":  {`{"Uint32": "a"}`, `unpacking into uint32, value "a" should be json.Number`},
		"uint64 negative":    {`{"Uint64": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"uint64 too big":     {`{"Uint64": 18446744073709551616}`, `strconv.ParseUint: parsing "18446744073709551616": value out of range`},
		"uint64 wrong type":  {`{"Uint64": "a"}`, `unpacking into uint64, value "a" should be json.Number`},
		"string number":      {`{"String": 1.0}`, `unpacking into string, value "1.0" should be string`},
		"string bool":        {`{"String": false}`, `unpacking into string, value false should be string`},
		"rune wrong type":    {`{"Rune": 1}`, `unpacking into rune, value "1" should be string`},
		"rune too long":      {`{"Rune": "aa"}`, `unpacking into rune: string should have a single rune`},
		"bool number":        {`{"Bool": 1.0}`, `unpacking into bool, value "1.0" should be bool`},
		"bool string":        {`{"Bool": "a"}`, `unpacking into bool, value "a" should be bool`},
		"byte negative":      {`{"Byte": -1}`, `strconv.ParseUint: parsing "-1": invalid syntax`},
		"byte too big":       {`{"Byte": 256}`, `strconv.ParseUint: parsing "256": value out of range`},
		"byte wrong type":    {`{"Byte": "a"}`, `unpacking into byte, value "a" should be json.Number`},
		"float32 wrong type": {`{"Float32": "a"}`, `unpacking into float32, value "a" should be json.Number`},
		"float32 too big":    {`{"Float32": 1e99}`, `strconv.ParseFloat: parsing "1e99": value out of range`},
		"float32 too small":  {`{"Float32": -1e99}`, `strconv.ParseFloat: parsing "-1e99": value out of range`},
		"float64 wrong type": {`{"Float64": "a"}`, `unpacking into float64, value "a" should be json.Number`},
		"float64 too big":    {`{"Float64": 1e999}`, `strconv.ParseFloat: parsing "1e999": value out of range`},
		"float64 too small":  {`{"Float64": -1e999}`, `strconv.ParseFloat: parsing "-1e999": value out of range`},
	}
	for name, test := range tests {
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatalf("%s error decoding %s", name, err)
		}
		_, errtest := Unpackers.Natives(system.NewContext(), v)
		if errtest == nil {
			t.Fatalf("%s expected error %s, got nil", name, test.error)
		}
		if !strings.Contains(errors.Cause(errtest).Error(), test.error) {
			t.Fatalf("%s error expected to contain '%s': %s", name, test.error, errors.Cause(errtest))
		}
	}
}
