package system

import (
	"fmt"
	"strings"

	"context"

	"regexp"

	"github.com/dave/kerr"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
)

type Reference struct {
	Package string
	Name    string
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
	err := r.Unpack(ctx, Pack(in), false)
	if err != nil {
		return nil, kerr.Wrap("VXRGOQHWNB", err)
	}
	return r, nil
}

func GetReferencePartsFromTypeString(ctx context.Context, typeString string) (path string, name string, err error) {

	env := envctx.FromContext(ctx)

	if strings.Contains(typeString, "/") {
		// If the type name contains a slash, I'm assuming it's a fully
		// qualified type name of the form "kego.io/system:type".
		// TODO: Improve this with a regex?
		parts := strings.Split(typeString, ":")

		// We hard-code system and json to prevent them having to always be
		// specified in the aliases
		if parts[0] == "kego.io/system" {
			return "kego.io/system", parts[1], nil
		} else if parts[0] == "kego.io/json" {
			return "kego.io/json", parts[1], nil
		}

		_, found := findKey(env.Aliases, parts[0])
		if !found && parts[0] != env.Path {
			return "", "", UnknownPackageError{
				Struct:         kerr.New("KJSOXDESFD", "Unknown package %s", parts[0]),
				UnknownPackage: parts[0],
			}
		}
		return parts[0], parts[1], nil
	} else if strings.Contains(typeString, ":") {
		// If the type name contains a colon, I'm assuming it's an abreviated
		// qualified type name of the form "system:type". We should look the
		// package name up in the aliases map.
		// TODO: Improve this with a regex?
		parts := strings.Split(typeString, ":")

		// We hard-code system and json to prevent them having to always be
		// specified in the aliases
		if parts[0] == "system" {
			return "kego.io/system", parts[1], nil
		} else if parts[0] == "json" {
			return "kego.io/json", parts[1], nil
		}

		packagePath, ok := env.Aliases[parts[0]]
		if !ok {
			return "", "", UnknownPackageError{
				Struct:         kerr.New("DKKFLKDKYI", "Unknown package %s", parts[0]),
				UnknownPackage: parts[0],
			}
		}
		return packagePath, parts[1], nil
	} else {
		return env.Path, typeString, nil
	}
}
func findKey(m map[string]string, value string) (string, bool) {
	for k, v := range m {
		if value == v {
			return k, true
		}
	}
	return "", false
}

type UnknownPackageError struct {
	kerr.Struct
	UnknownPackage string
}

type UnknownTypeError struct {
	kerr.Struct
	UnknownType string
}

func (v *Reference) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return kerr.New("MOQVSKJXRB", "Called Reference.Unpack with nil value")
	}
	if in.Type() == J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != J_STRING {
		return kerr.New("RFLQSBPMYM", "Can't unpack %s into *system.Reference", in.Type())
	}
	path, name, err := GetReferencePartsFromTypeString(ctx, in.String())
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

var _ Unpacker = (*Reference)(nil)

func (out *Reference) UnmarshalInterface(ctx context.Context, in interface{}) error {
	s, ok := in.(string)
	if !ok || s == "" {
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := GetReferencePartsFromTypeString(ctx, s)
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

var _ Repacker = (*Reference)(nil)

func (r *Reference) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType JsonType, err error) {
	if r == nil {
		return nil, "kego.io/system", "reference", J_NULL, nil
	}
	val, err := r.ValueContext(ctx)
	if err != nil {
		return nil, "", "", "", kerr.Wrap("VQCYFSTPQD", err)
	}
	return val, "kego.io/system", "reference", J_STRING, nil
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
