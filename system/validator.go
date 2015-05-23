package system

import "fmt"

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(path string, imports map[string]string) (bool, string, error)
}

func (r *String_rule) Validate(path string, imports map[string]string) (ok bool, message string, err error) {
	if r.MaxLength.Exists && r.MinLength.Exists {
		if r.MaxLength.Value < r.MinLength.Value {
			return false, fmt.Sprintf("MaxLength %d must not be lass than MinLength %d", r.MaxLength.Value, r.MinLength.Value), nil
		}
	}
	return true, "", nil
}
