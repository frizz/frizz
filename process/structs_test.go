package process

import (
	"regexp"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
	"kego.io/system"
)

func TestGenerateSource(t *testing.T) {

	cb := tests.Context("b.c/d")

	ty := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "a"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("object"),
	}
	cb.Stype("a", ty)

	tyn := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "an"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("bool"),
	}
	cb.Stype("an", tyn)

	tyi := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "ai"),
			Type: system.NewReference("kego.io/system", "type")},
		Native:    system.NewString("object"),
		Interface: true,
	}
	cb.Stype("ai", tyi)

	cb.Alias("f.g/h", "e")

	source, err := Structs(cb.Ctx(), cb.Env())
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package d\n")
	imp := getImports(t, string(source))
	assert.Contains(t, imp, "\t\"golang.org/x/net/context\"\n")
	assert.Contains(t, imp, "\t\"kego.io/context/jsonctx\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t\"reflect\"\n")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "\ntype A struct {\n\t*system.Object\n}\n")
	assert.Contains(t, string(source), "pkg.InitType(\"a\", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())\n")

	// Type "Ai" is an interface
	assert.Contains(t, string(source), "pkg.InitType(\"ai\", reflect.TypeOf((*Ai)(nil)).Elem(), reflect.TypeOf((*AiRule)(nil)), nil)\n")

	// Type "An" should not be a struct because it's a bool native. However, it should be registered.
	assert.NotContains(t, string(source), "type An struct")
	assert.Contains(t, string(source), "pkg.InitType(\"an\", reflect.TypeOf((*An)(nil)), reflect.TypeOf((*AnRule)(nil)), reflect.TypeOf((*AnInterface)(nil)).Elem())\n")

}

func getImports(t *testing.T, source string) string {
	r := regexp.MustCompile(`(?m)import \(([^)]*)\)`)
	matches := r.FindStringSubmatch(source)
	assert.Equal(t, 2, len(matches))
	return matches[1]
}

func TestGenerateCommand_errors(t *testing.T) {

	cb := tests.Context("a.b/\"").Spkg("a.b/\"")

	_, err := Structs(cb.Ctx(), cb.Env())
	// Quote in the path will generate malformed source
	assert.HasError(t, err, "CRBYOUOHPG")

	ty := &system.Type{
		Object: &system.Object{
			Id:   system.NewReference("b.c/d", "a corrupt"),
			Type: system.NewReference("kego.io/system", "type")},
		Native: system.NewString("object"),
	}
	cb = tests.Context("b.c/d").Stype("a", ty)

	_, err = Structs(cb.Ctx(), cb.Env())
	// Corrupt type ID causes error from source formatter
	assert.HasError(t, err, "CRBYOUOHPG")

}
