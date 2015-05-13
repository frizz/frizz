/*
uerr is an error with a unique ID
*/
package uerr // import "kego.io/uerr"

import (
	"fmt"

	"strings"

	"github.com/stretchr/testify/assert"
)

// Assert works with the assert package to test for a specific error
func Assert(t assert.TestingT, theError error, expectedId string, msgAndArgs ...interface{}) bool {

	message := messageFromMsgAndArgs(msgAndArgs...)
	if !assert.NotNil(t, theError, "An error is expected but got nil. %s", message) {
		return false
	}
	i, ok := theError.(Unique)
	if !assert.True(t, ok, "Error should implement uerr.Unique", message) {
		return false
	}
	return assert.Equal(t, expectedId, i.Unique(), "Expected %s but got %s. %s", expectedId, i.Unique(), message)

}
func Stack(t assert.TestingT, theError error, expectedId string, msgAndArgs ...interface{}) bool {
	message := messageFromMsgAndArgs(msgAndArgs...)
	if !assert.NotNil(t, theError, "An error is expected but got nil. %s", message) {
		return false
	}
	u, ok := theError.(UniqueError)
	if !assert.True(t, ok, "Error should be UniqueError", message) {
		return false
	}
	for _, i := range u.Stack {
		if i == expectedId {
			return true
		}
	}
	return assert.Fail(t, fmt.Sprintf("Didn't find error %s on stack.", expectedId), msgAndArgs...)
}
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		return msgAndArgs[0].(string)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// New creates a new uerr.Error
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
