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
			{"iface": {"__type": "Impi", "int": 1}}`,
			InterfaceField{Iface: Impi{Int: 1}},
		},
		"interface slice": {`
			{"slice": [{"__type": "Impi", "int": 1}, {"__type": "Imps", "string": "a"}]}`,
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
			{"__type": "string", "__value": "a"}`,
			"a",
		},
		"interface natives": {`
			{"__type": "Natives", "__value": {"string": "a"}}`,
			Natives{String: "a"},
		},
		"interface sub": {`
			{"__type": "sub.Sub", "__value": {"string": "a"}}`,
			sub.Sub{String: "a"},
		},
		"interface natives no value": {`
			{"__type": "Natives", "string": "a"}`,
			Natives{String: "a"},
		},
		"interface sub no value": {`
			{"__type": "sub.Sub", "string": "a"}`,
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
			{"_i": 1, "_s": "a"}`,
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
			{"string": "a"}`,
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
			{"a": {"sub": {"string": "a"}}, "b": {"sub": {"string": "b"}}}`,
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
				"sub": {"string": "a"}
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
				"string": "a",
				"int": 2,
				"sub": {"string": "a"},
				"array": [1, 2, 3],
				"slice": ["a", "b", "c"],
				"map": {"a": 1, "b": 2},
				"slice-string": ["a"],
				"slice-int": [1],
				"slice-sub": [{"string": "a"}]
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
				"ints": {"a": 1, "b": 2, "c": 3},
				"strings": {"a": "b", "c": "d", "e": "f"},
				"slices": {"a": ["b", "c"], "d": ["e", "f", "g"]},
				"arrays": {"a": [1, 2], "b": [3, 4]},
				"maps": {"a": {"b": "c"}, "d": {"e": "f"}}
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
				"ints": [1, 2, 3],
				"strings": ["a", "b", "c"],
				"array-lit": ["a", "b", "c", "d", "e"],
				"array-expr": [1, 2, 3, 4],
				"structs": [{"int": 1}, {"int": 2}],
				"arrays": [["a", "b"], ["c", "d"]]
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
			`{"array-lit": ["a", "b", "c", "d", "e", "f"]}`,
			"data length 6 does not fit in array of length 5",
		},
		"array expr too long": {
			`{"array-expr": [1, 2, 3, 4, 5]}`,
			"data length 5 does not fit in array of length 4",
		},
		"wrong type map": {
			`{"strings": {"a": "b"}}`,
			"error unpacking into slice, value should be an array",
		},
		"wrong type int": {
			`{"ints": 1}`,
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
				"simple": {
					"int": 1, 
					"bool": true
				}, 
				"complex": {
					"string": "a", 
					"inner": {
						"float-32": 1.1
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
		"int":              {`{"int": 1}`, Natives{Int: 1}},
		"int negative":     {`{"int": -1}`, Natives{Int: -1}},
		"int8 small":       {`{"int-8": -128}`, Natives{Int8: -128}},
		"int8 big":         {`{"int-8": 127}`, Natives{Int8: 127}},
		"int16 small":      {`{"int-16": -32768}`, Natives{Int16: -32768}},
		"int16 big":        {`{"int-16": 32767}`, Natives{Int16: 32767}},
		"int32 small":      {`{"int-32": -2147483648}`, Natives{Int32: -2147483648}},
		"int32 big":        {`{"int-32": 2147483647}`, Natives{Int32: 2147483647}},
		"int64 small":      {`{"int-64": -9223372036854775808}`, Natives{Int64: -9223372036854775808}},
		"int64 big":        {`{"int-64": 9223372036854775807}`, Natives{Int64: 9223372036854775807}},
		"uint":             {`{"uint": 1}`, Natives{Uint: 1}},
		"uint8 big":        {`{"uint-8": 255}`, Natives{Uint8: 255}},
		"uint16 big":       {`{"uint-16": 65535}`, Natives{Uint16: 65535}},
		"uint32 big":       {`{"uint-32": 4294967295}`, Natives{Uint32: 4294967295}},
		"uint64 big":       {`{"uint-64": 18446744073709551615}`, Natives{Uint64: 18446744073709551615}},
		"string":           {`{"string": "a"}`, Natives{String: "a"}},
		"rune":             {`{"rune": "😀"}`, Natives{Rune: '😀'}},
		"bool true":        {`{"bool": true}`, Natives{Bool: true}},
		"bool false":       {`{"bool": false}`, Natives{Bool: false}},
		"byte":             {`{"byte": 123}`, Natives{Byte: byte(123)}},
		"float32":          {`{"float-32": 1.5}`, Natives{Float32: 1.5}},
		"float32 negative": {`{"float-32": -1.5}`, Natives{Float32: -1.5}},
		"float64":          {`{"float-64": 1.5}`, Natives{Float64: 1.5}},
		"float64 negative": {`{"float-64": -1.5}`, Natives{Float64: -1.5}},
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
		"int fraction":       {`{"int": 1.5}`, "UGYQH"},
		"int string":         {`{"int": "a"}`, "DMSHN"},
		"int bool":           {`{"int": false}`, "DMSHN"},
		"int8 too small":     {`{"int-8": -129}`, "JWHEX"},
		"int8 too big":       {`{"int-8": 128}`, "JWHEX"},
		"int8 wrong type":    {`{"int-8": "a"}`, "PQDKT"},
		"int16 too small":    {`{"int-16": -32769}`, "VJGCV"},
		"int16 too big":      {`{"int-16": 32768}`, "VJGCV"},
		"int16 wrong type":   {`{"int-16": "a"}`, "YBJBD"},
		"int32 too small":    {`{"int-32": -2147483649}`, "BRQTC"},
		"int32 too big":      {`{"int-32": 2147483648}`, "BRQTC"},
		"int32 wrong type":   {`{"int-32": "a"}`, "DCWYT"},
		"int64 too small":    {`{"int-64": -9223372036854775809}`, "EPAHS"},
		"int64 too big":      {`{"int-64": 9223372036854775808}`, "EPAHS"},
		"int64 wrong type":   {`{"int-64": "a"}`, "DOFWU"},
		"uint negative":      {`{"uint": -1}`, "FMPHN"},
		"uint wrong type":    {`{"uint": "a"}`, "CIJOH"},
		"uint8 negative":     {`{"uint-8": -1}`, "HBACM"},
		"uint8 too big":      {`{"uint-8": 256}`, "HBACM"},
		"uint8 wrong type":   {`{"uint-8": "a"}`, "DBUNO"},
		"uint16 negative":    {`{"uint-16": -1}`, "CXCUQ"},
		"uint16 too big":     {`{"uint-16": 65536}`, "CXCUQ"},
		"uint16 wrong type":  {`{"uint-16": "a"}`, "SKVIO"},
		"uint32 negative":    {`{"uint-32": -1}`, "ILJLU"},
		"uint32 too big":     {`{"uint-32": 4294967296}`, "ILJLU"},
		"uint32 wrong type":  {`{"uint-32": "a"}`, "UJMCO"},
		"uint64 negative":    {`{"uint-64": -1}`, "QSFRS"},
		"uint64 too big":     {`{"uint-64": 18446744073709551616}`, "QSFRS"},
		"uint64 wrong type":  {`{"uint-64": "a"}`, "KCTKC"},
		"string number":      {`{"string": 1.0}`, "YTFIW"},
		"string bool":        {`{"string": false}`, "YTFIW"},
		"rune wrong type":    {`{"rune": 1}`, "MPHSY"},
		"rune too long":      {`{"rune": "aa"}`, "HGHCR"},
		"bool number":        {`{"bool": 1.0}`, "VPGKI"},
		"bool string":        {`{"bool": "a"}`, "VPGKI"},
		"byte negative":      {`{"byte": -1}`, "JPAPE"},
		"byte too big":       {`{"byte": 256}`, "JPAPE"},
		"byte wrong type":    {`{"byte": "a"}`, "LKYXV"},
		"float32 wrong type": {`{"float-32": "a"}`, "RBEXS"},
		"float32 too big":    {`{"float-32": 1e99}`, "SAHPU"},
		"float32 too small":  {`{"float-32": -1e99}`, "SAHPU"},
		"float64 wrong type": {`{"float-64": "a"}`, "YRENT"},
		"float64 too big":    {`{"float-64": 1e999}`, "INKLX"},
		"float64 too small":  {`{"float-64": -1e999}`, "INKLX"},
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
		//if !strings.Contains(errors.Cause(errtest).Error(), test.errstr) {
		//	t.Fatalf("%s error expected to contain '%s': %s", name, test.error, errors.Cause(errtest))
		//}

		// TODO: get annie working with code coverage.
		//a, err := ann.Lookup(errtest)
		//if err != nil {
		//	t.Fatalf("%s error looking up annotation: %s", name, err)
		//}
		//if a != test.error {
		//	t.Fatalf("%s unexpected error. Expected %s, got %s: %s", name, test.error, a, errtest)
		//}
	}
}
