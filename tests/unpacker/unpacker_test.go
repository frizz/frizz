package unpacker

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/dave/annie"
)

var ann = annie.New()

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
		result, err := unpackNatives(v)
		if err != nil {
			t.Fatalf("Error while unpacking %s: %s", name, err)
		}
		n, ok := result.(*Natives)
		if !ok {
			t.Fatalf("Result from natives unpack should be *Natives, got %T while unpacking %s", result, name)
		}
		if *n != test.expected {
			t.Fatalf("Result %#v not what expected while unpacking %s", *n, name)
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
		if name != "int fraction" {
			continue
		}
		var v interface{}
		d := json.NewDecoder(bytes.NewBuffer([]byte(test.json)))
		d.UseNumber()
		if err := d.Decode(&v); err != nil {
			t.Fatalf("%s error decoding %s", name, err)
		}
		_, errtest := unpackNatives(v)
		if errtest == nil {
			t.Fatalf("%s expected error %s, got nil", name, test.error)
		}
		a, err := ann.Lookup(errtest)
		if err != nil {
			t.Fatalf("%s error looking up annotation: %s", name, err)
		}
		if a != test.error {
			t.Fatalf("%s unexpected error. Expected %s, got %s: %s", name, test.error, a, errtest)
		}
	}
}
