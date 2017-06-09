package common

// frizz
type Validator interface {
	Validate(input interface{}) (valid bool, message string, err error)
}
