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
	Exists  bool
}

func (r Reference) Value() string {
	if !r.Exists {
		return ""
	}
	return fmt.Sprintf("%s:%s", r.Package, r.Name)
}

func (r Reference) ValueCompact(path string, aliases map[string]string) (string, error) {
	if !r.Exists {
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

func NewReference(packagePath string, typeName string) Reference {
	r := Reference{}
	r.Exists = true
	r.Package = packagePath
	r.Name = typeName
	return r
}

func (r *Reference) RuleToParentType() (*Reference, error) {
	if !r.Exists {
		return nil, kerr.New("OSQKOWGVWX", nil, "Reference is nil")
	}
	if !strings.HasPrefix(r.Name, "@") {
		return nil, kerr.New("HBKCDXQBYG", nil, "Type %s is not a rule type", r.Name)
	}
	newType := r.Name[1:]
	newRef := &Reference{
		Package: r.Package,
		Name:    newType,
		Exists:  r.Exists,
	}
	return newRef, nil
}

func NewReferenceFromString(in string, path string, aliases map[string]string) (*Reference, error) {
	r := &Reference{}
	err := r.UnmarshalJSON([]byte(strconv.Quote(in)), path, aliases)
	if err != nil {
		return nil, kerr.New("VXRGOQHWNB", err, "UnmarshalJSON")
	}
	return r, nil
}

func (out *Reference) UnmarshalJSON(in []byte, path string, aliases map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, aliases); err != nil {
		return kerr.New("BBWVFPNNTT", err, "json.UnmarshalPlain: %s", in)
	}
	if s == nil {
		out.Exists = false
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := json.GetReferencePartsFromTypeString(*s, path, aliases)
		if err != nil {
			// We need to clear the reference, because when we're scanning for
			// aliases we need to tolerate unknown import errors here
			out.Exists = false
			out.Name = ""
			out.Package = ""
			return err
		}
		out.Exists = true
		out.Package = path
		out.Name = name
	}
	return nil
}

var _ json.Unmarshaler = (*Reference)(nil)

func (r Reference) MarshalJSON() ([]byte, error) {
	if !r.Exists {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(r.Value())), nil
}

var _ json.Marshaler = (*Reference)(nil)

func (r Reference) MarshalCompactJSON(path string, aliases map[string]string) ([]byte, error) {
	if !r.Exists {
		return []byte("null"), nil
	}
	val, err := r.ValueCompact(path, aliases)
	if err != nil {
		return nil, kerr.New("VQCYFSTPQD", err, "ValueCompact")
	}
	return []byte(strconv.Quote(val)), nil
}

var _ json.CompactMarshaler = (*Reference)(nil)

func (r *Reference) String() string {
	if !r.Exists {
		return ""
	}
	return r.Value()
}

func (r *Reference) GetType() (*Type, bool) {
	if !r.Exists {
		return nil, false
	}
	if h, ok := GetGlobal(r.Package, r.Name); ok {
		if t, ok := h.Object.(*Type); ok {
			return t, true
		}
	}
	return nil, false
}

func (s Reference) NativeString() (value string, exists bool) {
	return s.Value(), s.Exists
}

type SortableReferences []Reference

func (s SortableReferences) Len() int {
	return len(s)
}
func (s SortableReferences) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableReferences) Less(i, j int) bool {
	return s[i].Value() < s[j].Value()
}

// We satisfy the json.EmptyAware interface to allow intelligent omission of
// empty values when marshalling
func (r Reference) Empty() bool {
	return !r.Exists
}

var _ json.EmptyAware = (*Reference)(nil)
