/*
kerr is an error with a unique ID
*/
package kerr // import "kego.io/kerr"

import (
	"fmt"

	"strings"
)

// New creates a new kerr.Error
func New(id string, inner error, location string, descriptionFormat string, descriptionArgs ...interface{}) UniqueError {
	stack := []string{id}
	if ui, ok := inner.(UniqueError); ok {
		stack = append(ui.Stack, id)
	}
	return UniqueError{
		Id:          id,
		Inner:       inner,
		Location:    location,
		Description: fmt.Sprintf(descriptionFormat, descriptionArgs...),
		Stack:       stack,
	}
}

// UniqueError is an error with a unique Id.
type UniqueError struct {
	Id          string
	Inner       error
	Location    string
	Description string
	Stack       []string
}

// Unique returns the unique Id of the error
func (e UniqueError) Unique() string {
	return e.Id
}

// Error returns a string of the error
func (e UniqueError) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("\n%s error in %s: %s.\n", e.Unique(), e.Location, e.Description)
	}
	inner := e.Inner.Error()
	if strings.HasPrefix(inner, "\n") {
		inner = inner[1:]
	}
	return fmt.Sprintf("\n%s error in %s: %s returned an error: \n%v", e.Unique(), e.Location, e.Description, inner)
}

// Unique things have a globally unique string
type Unique interface {
	Unique() string
}
