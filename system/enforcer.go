package system

import (
	"fmt"

	"kego.io/uerr"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(data interface{}, path string, imports map[string]string) (bool, string, error)
}

func (r *String_rule) Enforce(data interface{}, path string, imports map[string]string) (bool, string, error) {
	s, ok := data.(String)
	if !ok {
		return false, "", uerr.New("SXFBXGQSEA", nil, "String_rule.Enforce", "data %T should be system.String.", data)
	}
	if r.Equal.Exists {
		if !s.Exists {
			return false, fmt.Sprintf("Equal: Value must exist, and equal '%s'", r.Equal.Value), nil
		}
		if s.Value != r.Equal.Value {
			return false, fmt.Sprintf("Equal: Value must equal '%s'", r.Equal.Value), nil
		}
	}
	if r.MaxLength.Exists {
		if s.Exists && len(s.Value) > r.MaxLength.Value {
			return false, fmt.Sprintf("MaxLength: Length must not be greater than %d", r.MaxLength.Value), nil
		}
	}
	// TODO: Implement rules
	return true, "", nil
}
func (r *Number_rule) Enforce(data interface{}, path string, imports map[string]string) (bool, string, error) {
	// TODO: Implement rules
	return true, "", nil
}
func (r *Bool_rule) Enforce(data interface{}, path string, imports map[string]string) (bool, string, error) {
	// TODO: Implement rules
	return true, "", nil
}
func (r *Map_rule) Enforce(data interface{}, path string, imports map[string]string) (bool, string, error) {
	// TODO: Implement rules
	return true, "", nil
}
func (r *Array_rule) Enforce(data interface{}, path string, imports map[string]string) (bool, string, error) {
	// TODO: Implement rules
	return true, "", nil
}
