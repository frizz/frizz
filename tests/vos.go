package tests

import (
	"os"
	"strings"

	"kego.io/context/vosctx"
)

var _ vosctx.Vos = (*MockOs)(nil)

type MockOs struct {
	EnvironmentVariables map[string]string
	WorkingDirectory     string
}

func (o *MockOs) Getenv(key string) string {
	if o.EnvironmentVariables == nil || len(o.EnvironmentVariables) == 0 {
		return os.Getenv(key)
	}
	return o.EnvironmentVariables[key]
}

func (o *MockOs) Getwd() (string, error) {
	if o.WorkingDirectory == "" {
		return os.Getwd()
	}
	return o.WorkingDirectory, nil
}

// Environ returns a copy of strings representing the environment, in the form "key=value".
func (o *MockOs) Environ() []string {
	if o.EnvironmentVariables == nil || len(o.EnvironmentVariables) == 0 {
		return os.Environ()
	}
	out := []string{}
Outer:
	for _, real := range os.Environ() {
		for mockName, mockValue := range o.EnvironmentVariables {
			if strings.HasPrefix(real, mockName+"=") {
				out = append(out, mockName+"="+mockValue)
				continue Outer
			}
		}
		out = append(out, real)
	}
	return out
}
