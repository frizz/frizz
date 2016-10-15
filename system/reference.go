package system

import (
	"fmt"
	"strings"

	"reflect"

	"context"

	"regexp"

	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/packer"
)

type Reference struct {
	Package string
	Name    string
}

func New_Reference(ctx context.Context) interface{} {
	v := new(Reference)
	return v
}

func (r *Reference) Label(ctx context.Context) string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r Reference) Value() string {
	if r.Package == "" && r.Name == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s", r.Package, r.Name)
}

func (r Reference) ValueContext(ctx context.Context) (string, error) {

	env := envctx.FromContextOrNil(ctx)
	if env == nil || env.Path == "" {
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
	for alias, pkg := range env.Aliases {
		if pkg == r.Package {
			return fmt.Sprintf("%s:%s", alias, r.Name), nil
		}
	}
	return "", kerr.New("WGCDQQCFAD", "Package %s not found in aliases", r.Package)
}

func (r *ReferenceRule) Validate(ctx context.Context) (fail bool, messages []string, err error) {
	if r.Pattern != nil {
		if _, err := regexp.Compile(r.Pattern.Value()); err != nil {
			fail = true
			messages = append(messages, fmt.Sprintf("Pattern: regex does not compile: %s", r.Pattern.Value()))
		}
	}
	if r.PatternNot != nil {
		if _, err := regexp.Compile(r.PatternNot.Value()); err != nil {
			fail = true
			messages = append(messages, fmt.Sprintf("PatternNot: regex does not compile: %s", r.PatternNot.Value()))
		}
	}
	return
}

var _ Validator = (*ReferenceRule)(nil)

func (r *ReferenceRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if i, ok := data.(ReferenceInterface); ok && i != nil {
		data = i.GetReference(ctx)
	}

	v, ok := data.(*Reference)
	if !ok && data != nil {
		return true, nil, kerr.New("BYDVGGETWW", "Reference rule: value %T should be *system.Reference", data)
	}

	// Pattern restriction should be the same as StringRule
	var s *String
	if v != nil {
		s = NewString(v.Name)
	}
	sr := StringRule{
		Rule:       &Rule{Optional: r.Optional},
		Pattern:    r.Pattern,
		PatternNot: r.PatternNot,
	}
	if fail, messages, err = sr.Enforce(ctx, s); err != nil {
		return true, nil, kerr.Wrap("KYYJLYOSHT", err)
	}
	return
}

var _ Enforcer = (*ReferenceRule)(nil)

func NewReference(packagePath string, typeName string) *Reference {
	r := &Reference{}
	r.Package = packagePath
	r.Name = typeName
	return r
}

func NewReferenceFromString(ctx context.Context, in string) (*Reference, error) {
	r := &Reference{}
	err := r.Unpack(ctx, packer.Pack(in), false)
	if err != nil {
		return nil, kerr.Wrap("VXRGOQHWNB", err)
	}
	return r, nil
}

func (v *Reference) Unpack(ctx context.Context, in packer.Packed, iface bool) error {
	if in == nil || in.Type() == packer.J_NULL {
		return kerr.New("MOQVSKJXRB", "Called Reference.Unpack with nil value")
	}
	if v == nil {
		v = New_Reference(ctx).(*Reference)
	}
	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != packer.J_STRING {
		return kerr.New("RFLQSBPMYM", "Can't unpack %s into *system.Reference", in.Type())
	}
	path, name, err := packer.GetReferencePartsFromTypeString(ctx, in.String())
	if err != nil {
		// We need to clear the reference, because when we're scanning for
		// aliases we need to tolerate unknown import errors here
		v.Name = ""
		v.Package = ""
		return kerr.Wrap("MSXBLEIGVJ", err)
	}
	v.Package = path
	v.Name = name
	return nil
}

var _ packer.Unpacker = (*Reference)(nil)

func (out *Reference) UnmarshalInterface(ctx context.Context, in interface{}) error {
	s, ok := in.(string)
	if !ok || s == "" {
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := packer.GetReferencePartsFromTypeString(ctx, s)
		if err != nil {
			// We need to clear the reference, because when we're scanning for
			// aliases we need to tolerate unknown import errors here
			out.Name = ""
			out.Package = ""
			return kerr.Wrap("ETLPLMMWCC", err)
		}
		out.Package = path
		out.Name = name
	}
	return nil
}

var _ packer.Repacker = (*Reference)(nil)

func (r *Reference) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error) {
	if r == nil {
		return nil, "kego.io/system", "reference", nil
	}
	val, err := r.ValueContext(ctx)
	if err != nil {
		return nil, "", "", kerr.Wrap("VQCYFSTPQD", err)
	}
	return val, "kego.io/system", "reference", nil
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
	return r.changeTo(jsonctx.RULE_PREFIX)
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
	return strings.HasPrefix(r.Name, jsonctx.RULE_PREFIX)
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
	if r.Default == nil {
		return nil
	}
	return r.Default
}

var _ DefaultRule = (*ReferenceRule)(nil)
