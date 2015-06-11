package system

import (
	"fmt"
	"strings"

	"strconv"

	"kego.io/json"
	"kego.io/kerr"
)

type Reference struct {
	Value   string
	Package string
	Type    string
	Exists  bool
}

func NewReference(packagePath string, typeName string) Reference {
	r := Reference{}
	r.Exists = true
	r.Package = packagePath
	r.Type = typeName
	r.Value = fmt.Sprintf("%s:%s", r.Package, r.Type)
	return r
}

func (r *Reference) RuleToParentType() (*Reference, error) {
	if !r.Exists {
		return nil, kerr.New("OSQKOWGVWX", nil, "Reference.RuleToParentType", "Reference is nil")
	}
	if !strings.HasPrefix(r.Type, "@") {
		return nil, kerr.New("HBKCDXQBYG", nil, "Reference.RuleToParentType", "Type %s is not a rule type", r.Type)
	}
	newType := r.Type[1:]
	newRef := &Reference{
		Value:   fmt.Sprintf("%s:%s", r.Package, newType),
		Package: r.Package,
		Type:    newType,
		Exists:  r.Exists,
	}
	return newRef, nil
}

func NewReferenceFromString(in string, path string, imports map[string]string) (*Reference, error) {
	r := &Reference{}
	err := r.UnmarshalJSON([]byte(in), path, imports)
	if err != nil {
		return nil, kerr.New("VXRGOQHWNB", err, "system.NewReferenceFromString", "UnmarshalJSON")
	}
	return r, nil
}

func (out *Reference) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, imports); err != nil {
		return kerr.New("BBWVFPNNTT", err, "Reference.UnmarshalJSON", "json.UnmarshalPlain")
	}
	if s == nil {
		out.Exists = false
		out.Value = ""
		out.Type = ""
		out.Package = ""
	} else if *s == "" {
		// Special case for Object
		out.Exists = false
		out.Value = ""
		out.Type = ""
		out.Package = ""
	} else {
		out.Exists = true
		path, name, err := json.GetReferencePartsFromTypeString(*s, path, imports)
		if err != nil {
			return kerr.New("DBQPULKKUH", err, "Reference.UnmarshalJSON", "json.GetReferencePartsFromTypeString")
		}
		out.Package = path
		out.Type = name
		out.Value = fmt.Sprintf("%s:%s", out.Package, out.Type)
	}
	return nil
}

func (r *Reference) MarshalJSON() ([]byte, error) {
	if !r.Exists {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(r.Value)), nil
}

func (r *Reference) GoReference(path string, imports map[string]string) (string, error) {
	s, err := IdToGoReference(r.Package, r.Type, path, imports)
	if err != nil {
		return "", kerr.New("LVHAQUOQGR", err, "Reference.GoReference", "IdToGoReference")
	}
	return s, nil
}

func (r *Reference) GetType() (*Type, bool) {
	if !r.Exists {
		return nil, false
	}
	return GetType(r.Value)
}

func (s Reference) NativeString() (value string, exists bool) {
	return s.Value, s.Exists
}

func (s Reference) NativeExists() bool {
	return s.Exists
}
