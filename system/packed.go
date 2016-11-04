package system

import (
	"context"
	"fmt"

	"encoding/json"
)

type JsonType string

const (
	J_NULL     JsonType = "null"
	J_STRING   JsonType = "string"
	J_BOOL     JsonType = "bool"
	J_NUMBER   JsonType = "number"
	J_ARRAY    JsonType = "array"
	J_MAP      JsonType = "map"
	J_OBJECT   JsonType = "object"
	J_OPERATOR JsonType = "operator" // special type user by selectors
)

type Packed interface {
	Type() JsonType // packed will never be J_OBJECT, only J_MAP
	Number() float64
	String() string
	Bool() bool
	Array() []Packed
	Map() map[string]Packed
	Interface() interface{}
}

// Unpacker unpacks the data from in into the object. If iface is true, we're
// unpacking into an interface. In this situation, map types are always
// specified in typed form: {"type": "mytype", "value": {"foo": "bar"}}.
type Unpacker interface {
	Unpack(ctx context.Context, in Packed, iface bool) error
}

type Repacker interface {
	// Repack packs the object into json data ready for marshaling to a string.
	// If iface is true, we are packing into an interface type, so map types are
	// always specified in specific typed form: {"type": "mytype", "value":
	// {"foo": "bar"}}
	// Repack(ctx context.Context, iface bool) (interface{}, error)
	Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error)
}

type packed struct {
	v interface{}
	a []Packed
	m map[string]Packed
}

func Pack(v interface{}) *packed {
	return &packed{v: v}
}

func MustPackString(s string) *packed {
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		panic(err.Error())
	}
	return &packed{v: v}
}

var _ Packed = (*packed)(nil)

func (j *packed) Type() JsonType {
	if j == nil {
		return J_NULL
	}
	switch j.v.(type) {
	case nil:
		return J_NULL
	case float64:
		return J_NUMBER
	case bool:
		return J_BOOL
	case string:
		return J_STRING
	case map[string]interface{}:
		return J_MAP
	case []interface{}:
		return J_ARRAY
	}
	// ke: {"block": {"notest": true}}
	panic(fmt.Sprintf("Illegal type %T", j.v))
}
func (j *packed) Number() float64 {
	if j.v == nil {
		return 0.0
	}
	return j.v.(float64)
}
func (j *packed) String() string {
	if j.v == nil {
		return ""
	}
	return j.v.(string)
}
func (j *packed) Bool() bool {
	if j.v == nil {
		return false
	}
	return j.v.(bool)
}
func (j *packed) Array() []Packed {
	if j.a == nil {
		out := []Packed{}
		in, ok := j.v.([]interface{})
		if ok {
			for _, v := range in {
				out = append(out, Pack(v))
			}
		}
		j.a = out
	}
	return j.a
}
func (j *packed) Map() map[string]Packed {
	if j.m == nil {
		out := map[string]Packed{}
		in, ok := j.v.(map[string]interface{})
		if ok {
			for n, v := range in {
				out[n] = Pack(v)
			}
		}
		j.m = out
	}
	return j.m
}
func (j *packed) Interface() interface{} {
	return j.v
}
