package system

// Validator is a type that needs to have it's data validated.
type Validator interface {
	Validate(path string, imports map[string]string) (bool, string, error)
}
