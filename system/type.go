package system

import (
	"sort"

	"kego.io/kerr"
)

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(path string, aliases map[string]string) (bool, string, error)
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
		return "", kerr.New("TXQIDRBJRH", nil, "Native type not found: %v", jsonNativeType)
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

func (t *Type) SortedFields() []Field {
	fields := SortableFields{}
	for name, rule := range t.Fields {
		fields = append(fields, Field{Name: name, Rule: rule, Origin: t.Id})
	}
	sort.Sort(fields)
	return []Field(fields)
}

type Field struct {
	Name   string
	Rule   RuleInterface
	Origin Reference
}
type SortableFields []Field

func (s SortableFields) Len() int {
	return len(s)
}
func (s SortableFields) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableFields) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
