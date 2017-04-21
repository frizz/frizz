package ext

import (
	"testing"

	"frizz.io/process"
	"frizz.io/process/generate"
	"frizz.io/tests"
)

// notest

func InitialiseAndGenerate(t *testing.T, cb *tests.ContextBuilder, name string, files map[string]string) (string, error) {
	path, _ := cb.TempPackage(name, files)
	ctx, _, err := process.Initialise(cb.Ctx(), &process.Options{
		Path: path,
	})
	if err != nil {
		return "", err
	}
	cb.SetCtx(ctx)
	source, err := generate.Structs(ctx, cb.Env())
	if err != nil {
		return "", err
	}
	return string(source), nil
}
