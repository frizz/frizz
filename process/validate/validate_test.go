package validate

import (
	"testing"

	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/system/node"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
	_ "kego.io/process/validate/tests"
)

func TestFieldExtraMap(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:f
			id: b
			c:
				a:
					type: tests:a
					b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestFieldExtraArray(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:f
			id: b
			b:
				-
					type: tests:a
					b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestFieldExtraRulesObject(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:f
			id: b
			a:
				type: tests:a
				b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestValidateObjectChildren(t *testing.T) {
	cb := tests.New()
	n := &node.Node{Value: nil}
	err := validateObjectChildren(cb.Ctx(), n)
	assert.NoError(t, err)

	n = &node.Node{Value: 1, Null: true}
	err = validateObjectChildren(cb.Ctx(), n)
	assert.NoError(t, err)

	n = &node.Node{Value: 1, Missing: true}
	err = validateObjectChildren(cb.Ctx(), n)
	assert.NoError(t, err)

	r := &system.StringRule{Rule: &system.Rule{}}
	ty := &system.Type{Fields: map[string]system.RuleInterface{
		"A": r,
	}}
	n = &node.Node{Value: 1, Type: ty}
	err = validateObjectChildren(cb.Ctx(), n)
	assert.IsError(t, err, "ETODESNSET")

	r.Optional = true

	// Rule that doesn't implement ObjectInterface shouldn't happen
	assert.SkipError("XRTVWVUAMP")

}

func TestValidateCollection(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:e
			id: b
			rules:
				-
					selector: ".a>{system:string}"
					type: system:@string
					minLength: 5
			a:
				- abcabc
				- bcdbcd
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.NoError(t, err)

	cb.TempFile("c.yaml", `
			type: tests:e
			id: c
			rules:
				-
					selector: ".a>{system:string}"
					type: system:@string
					minLength: 5
			a:
				- abcabc
				- abc
		`)

	err = ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")

	cb.RemoveTempFile("c.yaml")

	cb.TempFile("d.yaml", `
			type: tests:e
			id: d
			rules:
				-
					selector: ".b>{system:string}"
					type: system:@string
					minLength: 5
			b:
				a: abcabc
				b: bcdbcd
		`)

	err = ValidatePackage(cb.Ctx())
	assert.NoError(t, err)

	cb.TempFile("e.yaml", `
			type: tests:e
			id: e
			rules:
				-
					selector: ".b>{system:string}"
					type: system:@string
					minLength: 5
			b:
				a: abcabc
				b: bcd
		`)

	err = ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestRulesEnforcer(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:a
			id: b
			b: foo
			rules:
				-
					type: tests:@a
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "ABVWHMMXGG")
}

func TestInterface(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:d
			id: b
			a:
				type: tests:a
				b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HLKQWDCMRN")
}

func TestTestRulesApplyToObjects(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/validate/tests")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				tests: kego.io/process/validate/tests
		`,
		"b.yml": `
			type: tests:a
			id: b
			b: foobar
			rules:
				-   selector: ".b"
					type: "system:@string"
					maxLength: 5
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")

}

func TestValidateNode(t *testing.T) {
	cb := tests.New()
	err := validateNode(cb.Ctx(), &node.Node{Value: nil})
	assert.NoError(t, err)
	err = validateNode(cb.Ctx(), &node.Node{Value: 1, Null: true})
	assert.NoError(t, err)
	err = validateNode(cb.Ctx(), &node.Node{Value: 1, Missing: true})
	assert.NoError(t, err)
}

func TestValidate_NeedsTypes(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"description": "a",
			"type": "system:type",
			"id": "a",
			"fields": {
				"a": {
					"type": "system:@string"
				}
			}
		}`,
	})

	cb.Path(path).Dir(dir).Jsystem().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.NoError(t, err)

}

func TestValidate_error1(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("b", map[string]string{
		"b.json": `{
			"description": "b",
			"type": "system:type",
			"id": "b",
			"fields": {
				"b": {
					"type": "system:@string",
					"minLength": 10,
					"maxLength": 5
				}
			}
		}`,
	})

	cb.Path(path).Dir(dir).Jsystem().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "YLONAMFUAG")

}
