package system

import (
	"sort"

	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/sysctx"
	"kego.io/json"
	"kego.io/kerr"
)

func GetAllTypesThatImplementInterface(ctx context.Context, typ *Type) []*Type {
	scache := sysctx.FromContext(ctx)

	out := []*Type{}

	var reflectType reflect.Type
	if typ.Interface {
		// The type provided is an interface type
		rt, ok := typ.Id.GetReflectType(ctx)
		if !ok {
			return nil
		}
		reflectType = rt
	} else {
		// The type provided is not an interface type, so we get it's automatic generated interface
		rt, ok := typ.Id.GetReflectInterface(ctx)
		if !ok {
			return nil
		}
		reflectType = rt
	}

	for _, pkgName := range scache.Keys() {
		pkgInfo, ok := scache.Get(pkgName)
		if !ok {
			continue
		}
		for _, typName := range pkgInfo.Types.Keys() {
			typ, ok := pkgInfo.Types.Get(typName)
			if !ok {
				continue
			}
			if typ.(*Type).Interface {
				continue
			}
			if typ.(*Type).Implements(ctx, reflectType) {
				out = append(out, typ.(*Type))
			}
		}
	}

	return out
}

func GetTypeFromCache(ctx context.Context, path string, name string) (*Type, bool) {
	scache := sysctx.FromContext(ctx)
	pcache, ok := scache.Get(path)
	if !ok {
		return nil, false
	}
	t, ok := pcache.Types.Get(name)
	if !ok {
		return nil, false
	}
	return t.(*Type), true
}

func (t *Type) ZeroValue(ctx context.Context) (interface{}, error) {
	if t.IsNativeCollection() {
		return nil, kerr.New("PGUHCGBJWE", "ZeroValue must not be used with collection type")
	}
	rt, ok := t.Id.GetReflectType(ctx)
	if !ok {
		return nil, kerr.New("RSWTEOTNBD", "Type not found for %s", t.Id)
	}
	return reflect.Zero(rt).Interface(), nil
}

func (t *Type) Implements(ctx context.Context, i reflect.Type) bool {
	ob, ok := t.Id.GetReflectType(ctx)
	if !ok {
		return false
	}
	return ob.Implements(i)
}

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(ctx context.Context) (bool, string, error)
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

func (t *Type) NativeJsonType() json.Type {
	switch t.Native.Value() {
	case "number":
		return json.J_NUMBER
	case "string":
		return json.J_STRING
	case "bool":
		return json.J_BOOL
	case "map":
		return json.J_MAP
	case "object":
		return json.J_OBJECT
	case "array":
		return json.J_ARRAY
	default:
		return json.J_NULL
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
		return "", kerr.New("TXQIDRBJRH", "Native type not found: %v", jsonNativeType)
	}
}

func (t *Type) IsJsonValue() bool {
	return t.Id.Package == "kego.io/json"
}
func (t *Type) IsNativeValue() bool {
	return nativeTypeClass(t.Native.Value()) == nativeValue
}
func (t *Type) IsNativeCollection() bool {
	return nativeTypeClass(t.Native.Value()) == nativeCollection
}
func (t *Type) IsNativeMap() bool {
	return t.Native.Value() == "map"
}
func (t *Type) IsNativeArray() bool {
	return t.Native.Value() == "array"
}
func (t *Type) IsNativeObject() bool {
	return nativeTypeClass(t.Native.Value()) == nativeObject
}
func (t *Type) NativeValueGolangType() (string, error) {
	return nativeGoType(t.Native.Value())
}

func (t *Type) GoName() string {
	return GoName(t.Id.Name)
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
	Origin *Reference
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
