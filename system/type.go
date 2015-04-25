package system

import (
	"fmt"

	kegofmt "kego.io/fmt"
)

func (t *Type) HasExtends() bool {
	// Only the actual system:object type should have an empty string here.
	return t.Extends.Value != ""
}

type nativeTypeClasses int

const (
	nativeValue nativeTypeClasses = iota
	nativeCollection
	nativeType
)

func nativeTypeClass(nativeTypeString string) nativeTypeClasses {
	switch nativeTypeString {
	case "number", "string", "bool":
		return nativeValue
	case "map", "array":
		return nativeCollection
	default:
		return nativeType
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
		return "", fmt.Errorf("Native type not found: %v\n", jsonNativeType)
	}
}

func (t *Type) IsNativeValue() bool {
	return nativeTypeClass(t.Native.Value) == nativeValue
}
func (t *Type) IsNativeCollection() bool {
	return nativeTypeClass(t.Native.Value) == nativeCollection
}
func (t *Type) IsNativeType() bool {
	return nativeTypeClass(t.Native.Value) == nativeType
}
func (t *Type) NativeValueGolangType() (string, error) {
	return nativeGoType(t.Native.Value)
}

func (t *Type) GoName() string {
	return IdToGoName(t.Id.Value)
}

// GoTypeReference outputs a Go source code reference to the name of this type. If we're in
// the local package, it just outputs the name e.g. "Object". If we're in a different package,
// it looks up the alias of the package in the imports and appends that to the start.
// e.g. "system.Object".
func (t *Type) GoTypeReference(localImports map[string]string, localPackagePath string) (string, error) {
	return IdToGoReference(t.Id.Value, t.Context.Package.Value, localImports, localPackagePath)
}

func (t *Type) ShouldImplementBasic() bool {
	return t.IsNativeType() && t.Extends.Exists && t.Extends.Value == "kego.io/system:object"
}

func (t *Type) GoSyntax(localPackage string, imports map[string]string) string {
	return kegofmt.GoSyntax(localPackage, imports, *t)
}

func (t *Type) Defaulter() (name string, property *Property, ok bool) {
	for name, prop := range t.Properties {
		if prop.Defaulter.Value {
			return name, prop, true
		}
	}
	return "", nil, false
}
