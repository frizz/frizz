package system

import (
	"sort"

	"reflect"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/context/sysctx"
	"kego.io/json"
)

func GetAllTypesThatImplementInterface(ctx context.Context, typ *Type) []*Type {
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

	return GetAllTypesThatImplementReflectInterface(ctx, reflectType)
}

func GetAllTypesThatImplementReflectInterface(ctx context.Context, reflectType reflect.Type) []*Type {

	if reflectType.Kind() != reflect.Interface {
		panic(kerr.New("JUCCMVNDLR", "%v is not an interface", reflectType).Error())
	}

	scache := sysctx.FromContext(ctx)

	out := []*Type{}

	for _, pkgName := range scache.Keys() {
		pkgInfo, ok := scache.Get(pkgName)
		if !ok {
			// ke: {"block": {"notest": true}}
			continue
		}
		for _, typName := range pkgInfo.Types.Keys() {
			typ, ok := pkgInfo.Types.Get(typName)
			if !ok {
				// ke: {"block": {"notest": true}}
				continue
			}
			t := typ.Type.(*Type)
			if t.Interface {
				continue
			}
			if t.Implements(ctx, reflectType) {
				out = append(out, t)
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
	return t.Type.(*Type), true
}

func (t *Type) ZeroValue(ctx context.Context, null bool) (reflect.Value, error) {
	if t.IsNativeCollection() {
		return reflect.Value{}, kerr.New("PGUHCGBJWE", "ZeroValue must not be used with collection type")
	}
	rt, ok := t.Id.GetReflectType(ctx)
	if !ok {
		return reflect.Value{}, kerr.New("RSWTEOTNBD", "Type not found for %s", t.Id)
	}
	return zeroValue(rt, null), nil
}

func zeroValue(rt reflect.Type, null bool) reflect.Value {
	if rt.Kind() == reflect.String || rt.Kind() == reflect.Float64 || rt.Kind() == reflect.Bool {
		return reflect.Zero(rt)
	}
	if null {
		return reflect.Zero(rt)
	}
	if rt.Kind() == reflect.Slice {
		return reflect.MakeSlice(rt, 0, 5)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(rt)
	}
	rv := reflect.New(rt.Elem())
	if rv.Elem().Kind() == reflect.Struct {
		zeroEmbed(rv.Elem())
	}
	return rv
}

func zeroEmbed(v reflect.Value) {
	// We loop round the fields, and initialise any anonymous
	// fields that are nil.
	for i := 0; i < v.Type().NumField(); i++ {
		sf := v.Type().Field(i)
		fv := v.FieldByName(sf.Name)
		if sf.Anonymous && fv.IsNil() {
			fv.Set(reflect.New(sf.Type.Elem()))
			zeroEmbed(fv.Elem())
		}
	}
}

func (t *Type) Implements(ctx context.Context, i reflect.Type) bool {
	if t.IsNativeCollection() {
		return false
	}
	rt, ok := t.Id.GetReflectType(ctx)
	if !ok {
		return false
	}
	return rt.Implements(i)
}

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(ctx context.Context) (fail bool, messages []string, err error)
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
	return t.IsNativeValue() && t.Id.Package == "kego.io/json"
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

func (t *Type) FieldOrigins() []*Reference {
	out := []*Reference{t.Id}
	out = append(out, t.Embed...)
	if !t.Basic {
		out = append(out, NewReference("kego.io/system", "object"))
	}
	return out
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
