package system

import (
	"sync"

	"kego.io/kerr"
)

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(path string, imports map[string]string) (bool, string, error)
}

var types struct {
	sync.RWMutex
	m map[Reference]*Type
}

func RegisterType(path string, name string, typ *Type) {
	types.Lock()
	defer types.Unlock()
	if types.m == nil {
		types.m = make(map[Reference]*Type)
	}
	types.m[NewReference(path, name)] = typ
}
func UnregisterType(path string, name string) {
	types.Lock()
	defer types.Unlock()
	if types.m == nil {
		return
	}
	delete(types.m, NewReference(path, name))
}

// TODO: Perhaps this should not return a pointer if it will
// TODO: be used concurrently?
func GetType(path string, name string) (t *Type, found bool) {
	types.RLock()
	defer types.RUnlock()
	t, found = types.m[NewReference(path, name)]
	return
}

// TODO: Perhaps this should not return pointers if it will
// TODO: be used concurrently?
func GetAllTypesInPackage(path string) map[Reference]*Type {
	out := map[Reference]*Type{}
	types.RLock()
	defer types.RUnlock()
	for ref, t := range types.m {
		if ref.Package == path {
			out[ref] = t
		}
	}
	return out
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
	return GoName(t.Id.Name)
}

func (t *Type) FullName() string {
	return t.Id.Value()
}
