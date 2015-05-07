package system

import (
	"fmt"
	"strings"

	"strconv"

	"kego.io/json"
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
		return nil, Err(nil, "Reference.RuleToParentType", "Reference is nil")
	}
	if !strings.HasPrefix(r.Type, "@") {
		return nil, Err(nil, "Reference.RuleToParentType", "Type %s is not a rule type", r.Type)
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

func (out *Reference) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, imports); err != nil {
		return err
	}
	if s == nil {
		out.Exists = false
	} else if *s == "" {
		// Special case for Object
		out.Exists = false
	} else {
		out.Exists = true
		path, name, err := json.GetReferencePartsFromTypeString(*s, path, imports)
		if err != nil {
			return Err(err, "Reference.UnmarshalJSON", "json.GetReferencePartsFromTypeString")
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

func (r *Reference) GoReference(localImports map[string]string, localPackagePath string) (string, error) {
	s, err := IdToGoReference(r.Type, r.Package, localImports, localPackagePath)
	if err != nil {
		return "", Err(err, "Reference.GoReference", "IdToGoReference")
	}
	return s, nil
}

func (r *Reference) GetType() (*Type, bool) {
	if !r.Exists {
		return nil, false
	}
	return GetType(r.Value)
}
