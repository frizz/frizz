package json

import "fmt"

type Unpackable interface {
	UpType() Type // Unpackable will never be J_OBJECT, only J_MAP
	UpNumber() float64
	UpString() string
	UpBool() bool
	UpArray() []Unpackable
	UpMap() map[string]Unpackable
	UpInterface() interface{}
}

type JsonUnpacker struct {
	value interface{}
	a     []Unpackable
	m     map[string]Unpackable
}

func NewJsonUnpacker(v interface{}) *JsonUnpacker {
	return &JsonUnpacker{value: v}
}

var _ Unpackable = (*JsonUnpacker)(nil)

func (j *JsonUnpacker) UpType() Type {
	if j == nil {
		return J_NULL
	}
	switch j.value.(type) {
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
	panic(fmt.Sprintf("Illegal type %T", j.value))
}
func (j *JsonUnpacker) UpNumber() float64 {
	return j.value.(float64)
}
func (j *JsonUnpacker) UpString() string {
	return j.value.(string)
}
func (j *JsonUnpacker) UpBool() bool {
	return j.value.(bool)
}
func (j *JsonUnpacker) UpArray() []Unpackable {
	if j.a == nil {
		out := []Unpackable{}
		in := j.value.([]interface{})
		for _, v := range in {
			out = append(out, NewJsonUnpacker(v))
		}
		j.a = out
	}
	return j.a
}
func (j *JsonUnpacker) UpMap() map[string]Unpackable {
	if j.m == nil {
		out := map[string]Unpackable{}
		in := j.value.(map[string]interface{})
		for n, v := range in {
			out[n] = NewJsonUnpacker(v)
		}
		j.m = out
	}
	return j.m
}
func (j *JsonUnpacker) UpInterface() interface{} {
	return j.value
}
