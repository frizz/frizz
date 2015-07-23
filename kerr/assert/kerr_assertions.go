package assert

import (
	"fmt"

	"kego.io/kerr"
)

func (a *Assertions) IsError(theError error, expectedId string, msgAndArgs ...interface{}) bool {
	return IsError(a.t, theError, expectedId, msgAndArgs...)
}

// IsError works with the kerr package to test for a specific error
func IsError(t TestingT, theError error, expectedId string, msgAndArgs ...interface{}) bool {

	message := messageFromMsgAndArgs(msgAndArgs...)
	if !NotNil(t, theError, "An error is expected but got nil. %s", message) {
		return false
	}
	i, ok := theError.(kerr.Unique)
	if !True(t, ok, "Error should implement kerr.Unique", message) {
		return false
	}
	return Equal(t, expectedId, i.Unique(), "Expected %s but got %s. %s:\n%s", expectedId, i.Unique(), message, theError)

}

func (a *Assertions) HasError(theError error, expectedId string, msgAndArgs ...interface{}) bool {
	return HasError(a.t, theError, expectedId, msgAndArgs...)
}

// HasError works with the kerr package to test for a specific error on the error stack
func HasError(t TestingT, theError error, expectedId string, msgAndArgs ...interface{}) bool {
	message := messageFromMsgAndArgs(msgAndArgs...)
	if !NotNil(t, theError, "An error is expected but got nil. %s", message) {
		return false
	}
	u, ok := theError.(kerr.UniqueError)
	if !True(t, ok, "Error should be UniqueError", message) {
		return false
	}
	for _, i := range u.Stack {
		if i == expectedId {
			return true
		}
	}
	return Fail(t, fmt.Sprintf("Didn't find error %s on stack:\n%s", expectedId, theError), msgAndArgs...)
}