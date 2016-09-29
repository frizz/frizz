package ext

import (
	"testing"

	"kego.io/process"
	"kego.io/process/generate"
	"kego.io/tests"
)

// ke: {"package": {"notest": true}}

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
