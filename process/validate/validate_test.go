package validate

import (
	"testing"

	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/system/node"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
	_ "kego.io/process/tests/simple"
)

func TestFieldExtraMap(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:f
			id: b
			c:
				a:
					type: simple:a
					b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestFieldExtraArray(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:f
			id: b
			b:
				-
					type: simple:a
					b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}

func TestFieldExtraRulesObject(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:f
			id: b
			a:
				type: simple:a
				b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

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
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:e
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

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.NoError(t, err)

	cb.TempFile("c.yaml", `
			type: simple:e
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
			type: simple:e
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
			type: simple:e
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
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:a
			id: b
			b: foo
			rules:
				-
					type: simple:@a
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "ABVWHMMXGG")
}

func TestInterface(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:d
			id: b
			a:
				type: simple:a
				b: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HLKQWDCMRN")
}

/*
func TestTestRulesApplyToInterface(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:b
			id: b
			c:
				type: simple:a
				b: abcabc
			rules:
				-   selector: ".c"
					type: "system:@string"
					maxLength: 5
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

	err := ValidatePackage(cb.Ctx())
	assert.IsError(t, err, "KWLWXKWHLF")
	assert.HasError(t, err, "HAOXUVTFEX")
}
*/

func TestTestRulesApplyToObjects(t *testing.T) {
	cb := tests.New().TempGopath(true).CopyToTemp("kego.io/process/tests/simple")
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
		"a.yml": `
			type: system:package
			id: a
			aliases:
				"kego.io/process/tests/simple": simple
		`,
		"b.yml": `
			type: simple:a
			id: b
			b: foobar
			rules:
				-   selector: ".b"
					type: "system:@string"
					maxLength: 5
		`,
	})

	cb.Path(path).Dir(dir).Alias("kego.io/process/tests/simple", "simple").Jauto().Sauto(parser.Parse)

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
