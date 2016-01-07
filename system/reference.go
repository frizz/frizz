package system

import (
	"fmt"
	"strings"

	"strconv"

	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/json"
	"kego.io/kerr"
)

type Reference struct {
	Package string
	Name    string
}

func (r Reference) Value() string {
	if r.Package == "" && r.Name == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s", r.Package, r.Name)
}

func (r Reference) ValueContext(ctx context.Context) (string, error) {

	env, ok := envctx.FromContextOrNil(ctx)
	if !ok || env.Path == "" {
		return r.Value(), nil
	}

	if r.Package == "" && r.Name == "" {
		return "", nil
	}
	if r.Package == env.Path {
		return r.Name, nil
	}
	if r.Package == "kego.io/json" {
		return fmt.Sprintf("json:%s", r.Name), nil
	}
	if r.Package == "kego.io/system" {
		return fmt.Sprintf("system:%s", r.Name), nil
	}
	if alias, ok := env.Aliases[r.Package]; ok {
		return fmt.Sprintf("%s:%s", alias, r.Name), nil
	}
	return "", kerr.New("WGCDQQCFAD", nil, "Package %s not found in aliases", r.Package)
}

func NewReference(packagePath string, typeName string) *Reference {
	r := &Reference{}
	r.Package = packagePath
	r.Name = typeName
	return r
}

func NewReferenceFromString(ctx context.Context, in string) (*Reference, error) {
	r := &Reference{}
	err := r.Unpack(ctx, json.Pack(in))
	if err != nil {
		return nil, kerr.New("VXRGOQHWNB", err, "Unpack")
	}
	return r, nil
}

func (out *Reference) Unpack(ctx context.Context, in json.Packed) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("MOQVSKJXRB", nil, "Called Reference.Unpack with nil value")
	}
	if in.Type() != json.J_STRING {
		return kerr.New("RFLQSBPMYM", nil, "Can't unpack %s into *system.Reference", in.Type())
	}
	path, name, err := json.GetReferencePartsFromTypeString(ctx, in.String())
	if err != nil {
		// We need to clear the reference, because when we're scanning for
		// aliases we need to tolerate unknown import errors here
		out.Name = ""
		out.Package = ""
		if p, ok := err.(json.UnknownPackageError); ok {
			// if GetReferencePartsFromTypeString returns an UnknownPackageError we should
			// not wrap it in kerr
			return p
		}
		return kerr.New("MSXBLEIGVJ", err, "json.GetReferencePartsFromTypeString")
	}
	out.Package = path
	out.Name = name
	return nil
}

var _ json.Unpacker = (*Reference)(nil)

func (out *Reference) UnmarshalInterface(ctx context.Context, in interface{}) error {
	s, ok := in.(string)
	if !ok || s == "" {
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := json.GetReferencePartsFromTypeString(ctx, s)
		if err != nil {
			// We need to clear the reference, because when we're scanning for
			// aliases we need to tolerate unknown import errors here
			out.Name = ""
			out.Package = ""
			return err
		}
		out.Package = path
		out.Name = name
	}
	return nil
}

var _ json.Marshaler = (*Reference)(nil)

func (r *Reference) MarshalJSON(ctx context.Context) ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	val, err := r.ValueContext(ctx)
	if err != nil {
		return nil, kerr.New("VQCYFSTPQD", err, "ValueContext")
	}
	return []byte(strconv.Quote(val)), nil
}

func (r Reference) String() string {
	return r.Value()
}

func (r Reference) GetType(ctx context.Context) (*Type, bool) {
	if r.Package == "" || r.Name == "" {
		return nil, false
	}
	t, ok := GetTypeFromCache(ctx, r.Package, r.Name)
	if !ok {
		return nil, false
	}
	return t, true
}

func (r Reference) GetReflectType(ctx context.Context) (reflect.Type, bool) {
	if t, ok := jsonctx.FromContext(ctx).GetType(r.Package, r.Name); ok {
		return t, true
	}
	return nil, false
}
func (r Reference) GetReflectInterface(ctx context.Context) (reflect.Type, bool) {
	if t, ok := jsonctx.FromContext(ctx).GetInterface(r.Package, r.Name); ok {
		return t, true
	}
	return nil, false
}

func (r Reference) ChangeToType() Reference {
	return r.changeTo("")
}
func (r Reference) ChangeToRule() Reference {
	return r.changeTo(json.RULE_PREFIX)
}
func (r Reference) changeTo(prefix string) Reference {
	if r.Package == "" || r.Name == "" {
		return Reference{}
	}
	n := r.Name
	if r.IsRule() {
		n = n[1:]
	}
	return *NewReference(r.Package, prefix+n)
}
func (r Reference) IsRule() bool {
	return strings.HasPrefix(r.Name, json.RULE_PREFIX)
}

func (s Reference) NativeString() string {
	return s.Value()
}

var _ NativeString = (*Reference)(nil)

type SortableReferences []*Reference

func (s SortableReferences) Len() int {
	return len(s)
}
func (s SortableReferences) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableReferences) Less(i, j int) bool {
	return s[i].Value() < s[j].Value()
}

func (r *ReferenceRule) GetDefault() interface{} {
	return r.Default
}

var _ DefaultRule = (*ReferenceRule)(nil)
