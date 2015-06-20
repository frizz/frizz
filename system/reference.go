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

func NewReference(packagePath string, typeName string) Reference {
	r := Reference{}
	r.Exists = true
	r.Package = packagePath
	r.Name = typeName
	return r
}

func (r *Reference) RuleToParentType() (*Reference, error) {
	if !r.Exists {
		return nil, kerr.New("OSQKOWGVWX", nil, "Reference.RuleToParentType", "Reference is nil")
	}
	if !strings.HasPrefix(r.Name, "@") {
		return nil, kerr.New("HBKCDXQBYG", nil, "Reference.RuleToParentType", "Type %s is not a rule type", r.Name)
	}
	newType := r.Name[1:]
	newRef := &Reference{
		Package: r.Package,
		Name:    newType,
		Exists:  r.Exists,
	}
	return newRef, nil
}

func NewReferenceFromString(in string, path string, imports map[string]string) (*Reference, error) {
	r := &Reference{}
	err := r.UnmarshalJSON([]byte(strconv.Quote(in)), path, imports)
	if err != nil {
		return nil, kerr.New("VXRGOQHWNB", err, "system.NewReferenceFromString", "UnmarshalJSON")
	}
	return r, nil
}

func (out *Reference) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, imports); err != nil {
		return kerr.New("BBWVFPNNTT", err, "Reference.UnmarshalJSON", "json.UnmarshalPlain: %s", in)
	}
	if s == nil {
		out.Exists = false
		out.Name = ""
		out.Package = ""
	} else {
		path, name, err := json.GetReferencePartsFromTypeString(*s, path, imports)
		if err != nil {
			// We don't want to throw an error here, because when we're scanning for
			// imports we need to tolerate unknown imports
			out.Exists = false
			out.Name = ""
			out.Package = ""
			return nil
		}
		out.Exists = true
		out.Package = path
		out.Name = name
	}
	return nil
}

func (r *Reference) MarshalJSON() ([]byte, error) {
	if !r.Exists {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(r.Value())), nil
}

func (r *Reference) String() string {
	if !r.Exists {
		return ""
	}
	return r.Value()
}

func (r *Reference) GoReference(path string, imports map[string]string) (string, error) {
	s, err := IdToGoReference(r.Package, r.Name, path, imports)
	if err != nil {
		return "", kerr.New("LVHAQUOQGR", err, "Reference.GoReference", "IdToGoReference")
	}
	return s, nil
}

func (r *Reference) GetType() (*Type, bool) {
	if !r.Exists {
		return nil, false
	}
	return GetType(r.Package, r.Name)
}

func (s Reference) NativeString() (value string, exists bool) {
	return s.Value(), s.Exists
}
