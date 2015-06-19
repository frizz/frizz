package system

import (
	"fmt"
	"sync"

	"strings"

	"kego.io/kerr"
)

var types struct {
	sync.RWMutex
	m map[string]*Type
}

func RegisterType(name string, typ *Type) {
	types.Lock()
	defer types.Unlock()
	if types.m == nil {
		types.m = make(map[string]*Type)
	}
	types.m[name] = typ
}
func UnregisterType(name string) {
	types.Lock()
	defer types.Unlock()
	if types.m == nil {
		return
	}
	delete(types.m, name)
}

// TODO: Perhaps this should not return a pointer if it will
// TODO: be used concurrently?
func GetType(name string) (t *Type, found bool) {
	types.RLock()
	defer types.RUnlock()
	t, found = types.m[name]
	return
}

// TODO: Perhaps this should not return pointers if it will
// TODO: be used concurrently?
func GetAllTypesInPackage(path string) map[string]*Type {
	out := map[string]*Type{}
	types.RLock()
	defer types.RUnlock()
	for k, t := range types.m {
		if strings.HasPrefix(k, fmt.Sprintf("%s:", path)) {
			out[k] = t
		}
	}
	return out
}

func (t *Type) HasExtends() bool {
	// Only the actual system:object type should have an empty string here.
	return t.Extends.Value() != ""
}

type nativeTypeClasses string

const (
	nativeValue      nativeTypeClasses = "value"
	nativeCollection nativeTypeClasses = "collection"
	nativeObject     nativeTypeClasses = "object"
)

func nativeTypeClass(nativeTypeString string) nativeTypeClasses {
	switch nativeTypeString {
	case "number", "string", "bool":
		return nativeValue
	case "map", "array":
		return nativeCollection
	default:
		return nativeObject
	}
}

func nativeGoType(jsonNativeType string) (string, error) {
	switch jsonNativeType {
	case "number":
		return "float64", nil
	case "string":
		return "string", nil
	case "bool":
		return "bool", nil
	default:
		return "", kerr.New("TXQIDRBJRH", nil, "nativeGoType", "Native type not found: %v", jsonNativeType)
	}
}

func (t *Type) IsNativeValue() bool {
	return nativeTypeClass(t.Native.Value) == nativeValue
}
func (t *Type) IsNativeCollection() bool {
	return nativeTypeClass(t.Native.Value) == nativeCollection
}
func (t *Type) IsNativeMap() bool {
	return t.Native.Value == "map"
}
func (t *Type) IsNativeArray() bool {
	return t.Native.Value == "array"
}
func (t *Type) IsNativeObject() bool {
	return nativeTypeClass(t.Native.Value) == nativeObject
}
func (t *Type) NativeValueGolangType() (string, error) {
	return nativeGoType(t.Native.Value)
}

func (t *Type) GoName() string {
	return IdToGoName(t.Id)
}

func (t *Type) FullName() string {
	return fmt.Sprintf("%s:%s", t.Context.Package, t.Id)
}

// GoTypeReference outputs a Go source code reference to the name of this type. If we're in
// the local package, it just outputs the name e.g. "String". If we're in a different package,
// it looks up the alias of the package in the imports and appends that to the start.
// e.g. "system.String".
func (t *Type) GoTypeReference(path string, imports map[string]string) (string, error) {
	return IdToGoReference(t.Context.Package, t.Id, path, imports)
}

/*
func (t *Type) Defaulter() (name string, property *Property, ok bool) {
	for name, prop := range t.Properties {
		if prop.Defaulter {
			return name, prop, true
		}
	}
	return "", nil, false
}
*/
