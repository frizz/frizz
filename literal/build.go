package literal

import (
	"reflect"

	"fmt"
)

type Pointer struct {
	Value  uintptr
	Name   string
	Source string
}

func findPointer(pointers *[]Pointer, v uintptr) (Pointer, bool) {
	for _, p := range *pointers {
		if p.Value == v {
			return p, true
		}
	}
	return Pointer{}, false
}

func Build(object interface{}, pointers *[]Pointer, path string, getAlias func(string) string) Pointer {
	return buildValue(reflect.ValueOf(object), pointers, path, getAlias)
}
func buildValue(value reflect.Value, pointers *[]Pointer, path string, getAlias func(string) string) Pointer {

	if value.Kind() != reflect.Ptr {
		panic("Must be pointer")
	}
	ptr := value.Pointer()
	if pointer, ok := findPointer(pointers, ptr); ok {
		return pointer
	}

	switch value.Elem().Kind() {
	case reflect.Bool,
		reflect.Float64, reflect.Float32,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		value = value.Elem()
	}

	p := newPrinter()
	p.build = true
	p.path = path
	p.getAlias = getAlias
	p.pointers = pointers
	p.reordered = false
	p.goodArgNum = true
	p.fmt.clearflags()
	p.fmt.sharpV = true
	p.printValue(value, 'v', 0)
	s := string(p.buf)
	p.free()
	pointer := Pointer{Value: ptr, Name: fmt.Sprint("ptr", len(*p.pointers)), Source: s}
	*p.pointers = append(*p.pointers, pointer)
	return pointer

}
