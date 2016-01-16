package tests

import (
	"os"

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
