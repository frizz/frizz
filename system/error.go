package system

import "fmt"

// Set RETURN_ERRORS to true to bubble up plain errors without extra formatting
var RETURN_ERRORS = false

func Err(inner error, location string, descriptionFormat string, descriptionArgs ...interface{}) error {
	if inner != nil && RETURN_ERRORS {
		return inner
	}
	return Error{
		Inner:       inner,
		Location:    location,
		Description: fmt.Sprintf(descriptionFormat, descriptionArgs...),
	}
}

type Error struct {
	Inner       error
	Location    string
	Description string
}

func (e Error) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("Error in %s: %s.", e.Location, e.Description)
	}
	return fmt.Sprintf("Error in %s: %s returned an error: \n%v\n", e.Location, e.Description, e.Inner)
}
