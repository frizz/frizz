package system

import (
	"fmt"
	"strings"

	"strconv"

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

func (r Reference) ValueContext(path string, aliases map[string]string) (string, error) {
	if r.Package == "" && r.Name == "" {
		return "", nil
	}
	if r.Package == path {
		return r.Name, nil
	}
	if r.Package == "kego.io/json" {
		return fmt.Sprintf("json:%s", r.Name), nil
	}
	if r.Package == "kego.io/system" {
		return fmt.Sprintf("system:%s", r.Name), nil
	}
	if alias, ok := aliases[r.Package]; ok {
		return fmt.Sprintf("%s:%s", alias, r.Name), nil
	}
	return "", kerr.New("WGCDQQCFAD", nil, "Package %s not found in aliases", r.Package)
	// Should we return an error here?
	//return fmt.Sprintf("%s:%s", r.Package, r.Name)
}

func NewReference(packagePath string, typeName string) *Reference {
	r := &Reference{}
	r.Package = packagePath
	r.Name = typeName
	return r
}

func NewReferenceFromString(in string, path string, aliases map[string]string) (*Reference, error) {
	r := &Reference{}
	err := r.Unpack(json.Pack(in), path, aliases)
	if err != nil {
		return nil, kerr.New("VXRGOQHWNB", err, "Unpack")
	}
	return r, nil
}

func (out *Reference) Unpack(in json.Packed, path string, aliases map[string]string) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("MOQVSKJXRB", nil, "Called Reference.Unpack with nil value")
	}
	if in.Type() != json.J_STRING {
		return kerr.New("RFLQSBPMYM", nil, "Can't unpack %s into *system.Reference", in.Type())
	}
	path, name, err := json.GetReferencePartsFromTypeString(in.String(), path, aliases)
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

var _ json.ContextUnpacker = (*Reference)(nil)

func (out *Reference) UnmarshalInterface(in interface{}, path string, aliases map[string]string) error {
	s, ok := in.(string)
	if !ok || s == "" {
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := json.GetReferencePartsFromTypeString(s, path, aliases)
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

func (r *Reference) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(r.Value())), nil
}

var _ json.Marshaler = (*Reference)(nil)

func (r *Reference) MarshalContextJSON(path string, aliases map[string]string) ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	val, err := r.ValueContext(path, aliases)
	if err != nil {
		return nil, kerr.New("VQCYFSTPQD", err, "ValueContext")
	}
	return []byte(strconv.Quote(val)), nil
}

var _ json.ContextMarshaler = (*Reference)(nil)

func (r Reference) String() string {
	return r.Value()
}

func (r Reference) GetType() (*Type, bool) {
	if r.Package == "" || r.Name == "" {
		return nil, false
	}
	if h, ok := GetGlobal(r.Package, r.Name); ok {
		if t, ok := h.Object.(*Type); ok {
			return t, true
		}
	}
	return nil, false
}

func (r Reference) ChangeToType() Reference {
	return r.changeTo("")
}
func (r Reference) ChangeToRule() Reference {
	return r.changeTo(RULE_PREFIX)
}
func (r Reference) ChangeToInterface() Reference {
	return r.changeTo(INTERFACE_PREFIX)
}
func (r Reference) changeTo(prefix string) Reference {
	if r.Package == "" || r.Name == "" {
		return Reference{}
	}
	n := r.Name
	if strings.HasPrefix(n, INTERFACE_PREFIX) || strings.HasPrefix(n, RULE_PREFIX) {
		n = n[1:]
	}
	return *NewReference(r.Package, prefix+n)
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
