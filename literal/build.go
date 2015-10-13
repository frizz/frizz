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

	value := reflect.ValueOf(object)
	if value.Kind() != reflect.Ptr {
		panic("Must be pointer")
	}
	if pointer, ok := findPointer(pointers, value.Pointer()); ok {
		return pointer
	}

	p := newPrinter()
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
	pointer := Pointer{Value: value.Pointer(), Name: fmt.Sprint("ptr", len(*p.pointers)), Source: s}
	*p.pointers = append(*p.pointers, pointer)
	return pointer

}
