package generate

import (
	"testing"

	"regexp"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
	"kego.io/system"
)

func TestGenerateCommand_errors(t *testing.T) {

	_, err := Structs(tests.PathCtx("a.b/\""))
	// Quote in the path will generate malformed source
	assert.HasError(t, err, "CRBYOUOHPG")

	ty := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "a corrupt"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("object"),
	}
	system.Register("b.c/d", "a", ty, 0)
	defer system.Unregister("b.c/d", "a")

	_, err = Structs(tests.PathCtx("b.c/d"))
	// Corrupt type ID causes error from source formatter
	assert.HasError(t, err, "CRBYOUOHPG")

	_, err = Types(tests.EnvCtx("b.c/d", map[string]string{"\"": "\""}))
	// Quote in the alias path will generate malformed source
	assert.HasError(t, err, "CRBYOUOHPG")
}
func getImports(t *testing.T, source string) string {
	r := regexp.MustCompile(`(?m)import \(([^)]*)\)`)
	matches := r.FindStringSubmatch(source)
	assert.Equal(t, 2, len(matches))
	return matches[1]
}
func TestGenerateSource(t *testing.T) {

	ty := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "a"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("object"),
	}
	system.Register("b.c/d", "a", ty, 0)
	defer system.Unregister("b.c/d", "a")

	tyn := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "an"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("bool"),
	}
	system.Register("b.c/d", "an", tyn, 0)
	defer system.Unregister("b.c/d", "an")

	tyi := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "ai"),
			Type: system.NewReference("kego.io/system", "type")},
		Native:    system.NewString("object"),
		Interface: true,
	}
	system.Register("b.c/d", "ai", tyi, 0)
	defer system.Unregister("b.c/d", "ai")

	source, err := Structs(tests.EnvCtx("b.c/d", map[string]string{"f.g/h": "e"}))
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package d\n")
	imp := getImports(t, string(source))
	assert.Contains(t, imp, "\t\"kego.io/json\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t\"reflect\"\n")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "\ntype A struct {\n\t*system.Object\n}\n")
	assert.Contains(t, string(source), "json.Register(\"b.c/d\", \"a\", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem(), 0)\n")

	source, err = Structs(tests.EnvCtx("b.c/d", map[string]string{"f.g/h": "e"}))
	assert.NoError(t, err)
	assert.Contains(t, string(source), "json.Register(\"b.c/d\", \"an\", reflect.TypeOf((*An)(nil)), reflect.TypeOf((*AnInterface)(nil)).Elem(), 0)\n")

	source, err = Structs(tests.EnvCtx("b.c/d", map[string]string{"f.g/h": "e"}))
	assert.NoError(t, err)
	assert.Contains(t, string(source), "json.Register(\"b.c/d\", \"ai\", reflect.TypeOf((*Ai)(nil)).Elem(), nil, 0)\n")

	source, err = Types(tests.EnvCtx("b.c/d", map[string]string{"f.g/h": "e"}))
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package types")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "\t_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t_ \"kego.io/system/types\"\n")
	assert.NotContains(t, imp, "\"b.c/d\"")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "system.Register(\"b.c/d\", \"a\", ptr4, 0x0)")

	source, err = TypesCommand(tests.EnvCtx("b.c/d", map[string]string{"f.g/h": "e"}))
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package main\n")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "_ \"b.c/d\"\n")
	assert.Contains(t, imp, "_ \"f.g/h\"\n")
	assert.Contains(t, imp, "_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "\"fmt\"\n")
	assert.Contains(t, imp, "\t\"kego.io/process\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system/types\"\n")
	assert.Contains(t, imp, "\"os\"\n")
	assert.Contains(t, string(source), "process.Generate(ctx, process.S_TYPES)")

}
