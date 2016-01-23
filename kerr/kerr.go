// Package kerr is an error with a unique ID
package kerr // import "kego.io/kerr"

import (
	"fmt"
	"runtime"

	"os"
	"path/filepath"
	"strings"
)

// Struct is an error with a unique Id.
type Struct struct {
	// Unique ID for this error - used to identify this error in tests
	Id string
	// The inner error. Nil if this is the source error.
	Inner error
	// Path of the file that this error occurred in
	File string
	// Line number of the source where this error occurred
	Line int
	// Package path where the error occurred
	Package string
	// Function name where this error occurred
	Function string
	// Description is a description of the error if there is no inner error, or else
	// a description of the fucntion that returned the inner error
	Description string
	// Array of unique IDs of the error stack
	Stack []string
}

type Interface interface {
	ErrorId() string
	ErrorStack() []string
	ErrorInner() error
}

// New creates a new kerr.Struct
func New(id string, descriptionFormat string, descriptionArgs ...interface{}) Struct {
	return get(id, nil, descriptionFormat, descriptionArgs...)
}

// Wrap wraps an error in a kerr.Struct
func Wrap(id string, inner error) Struct {
	return get(id, inner, "")
}

func get(id string, inner error, descriptionFormat string, descriptionArgs ...interface{}) Struct {
	stack := []string{id}
	if i, ok := inner.(Interface); ok {
		stack = append(i.ErrorStack(), id)
	}

	packageName, functionName := "", ""
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		f := runtime.FuncForPC(pc)
		caller := f.Name()
		i := strings.LastIndex(caller, ".")
		if i > -1 {
			packageName = caller[:i]
			functionName = caller[i+1:]
		}
	}

	return Struct{
		Id:          id,
		Inner:       inner,
		File:        file,
		Line:        line,
		Package:     packageName,
		Function:    functionName,
		Description: fmt.Sprintf(descriptionFormat, descriptionArgs...),
		Stack:       stack,
	}
}

// Error returns a string of the error
func (e Struct) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("\n%s error in %s:%d %s: %s.\n", e.Id, getRelPath(e.File), e.Line, e.Function, e.Description)
	}
	inner := e.Inner.Error()
	if strings.HasPrefix(inner, "\n") {
		// Remove the leading new-line from inner error
		inner = inner[1:]
	}
	return fmt.Sprintf("\n%s error in %s:%d %s: \n%v", e.Id, getRelPath(e.File), e.Line, e.Function, inner)
}

func getRelPath(filePath string) (out string) {
	wd, err := os.Getwd()
	if err != nil {
		out = filePath
		return
	}
	out, err = filepath.Rel(wd, filePath)
	if err != nil {
		out = filePath
		return
	}
	return
}

// ErrorId returns the unique id of the error
func (e Struct) ErrorId() string {
	return e.Id
}

// ErrorStack returns the error id stack
func (e Struct) ErrorStack() []string {
	return e.Stack
}

// ErrorInner returns the inner error
func (e Struct) ErrorInner() error {
	return e.Inner
}

// Source gets the error at the bottom of the error stack. This can't be a method on kerr.Struct
// because it needs to return the whole error when called on an error that embeds kerr.Struct
func Source(e error) error {
	if i, ok := e.(Interface); ok && i.ErrorInner() != nil {
		return Source(i.ErrorInner())
	}
	return e
}
