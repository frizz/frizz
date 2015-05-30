package literal

import (
	"reflect"
	"strconv"

	"github.com/go-errors/errors"
)

func Build(object interface{}, pointers map[string]string, order *[]string, path string, imports map[string]string) string {

	value := reflect.ValueOf(object)
	if value.Kind() != reflect.Ptr {
		panic(errors.New("Must be pointer"))
	}
	name := pointerLiteralName(value.Pointer())
	if pointers[name] != "" {
		return name
	}

	p := newPrinter()
	p.path = path
	p.imports = imports
	p.pointers = pointers
	p.order = order
	p.reordered = false
	p.goodArgNum = true
	p.fmt.clearflags()
	p.fmt.sharpV = true
	p.printValue(value, 'v', 0)
	s := string(p.buf)
	p.free()
	*p.order = append(*p.order, name)
	p.pointers[name] = s
	return name

}

func pointerLiteralName(addr uintptr) string {
	s := "ptr" + strconv.Itoa(int(addr))
	return s
}
