package validate

import (
	"testing"

	"kego.io/process/parser"
	"kego.io/system/node"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	_ "kego.io/process/validate/tests"
	"kego.io/tests"
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

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")
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

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")
}

func TestRuleHasExtraRules(t *testing.T) {
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
			d: foo
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")
	assert.Equal(t, "MinLength: length of \"foo\" must not be less than 7", errors[0].Description)
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

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")
}

// TODO: Recreate these
/*
func TestValidateObjectChildren(t *testing.T) {
	cb := tests.New()
	n := &node.Node{Value: nil}
	errors, err := validateObjectChildren(cb.Ctx(), n)
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

	n = &node.Node{Value: 1, Null: true}
	errors, err = validateObjectChildren(cb.Ctx(), n)
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

	n = &node.Node{Value: 1, Missing: true}
	errors, err = validateObjectChildren(cb.Ctx(), n)
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

	r := &system.StringRule{Rule: &system.Rule{}}
	ty := &system.Type{Fields: map[string]system.RuleInterface{
		"A": r,
	}}
	n = &node.Node{Value: 1, Type: ty}
	_, err = validateObjectChildren(cb.Ctx(), n)
	assert.IsError(t, err, "ETODESNSET")

	r.Optional = true

	// Rule that doesn't implement ObjectInterface shouldn't happen
	assert.SkipError("XRTVWVUAMP")

}
*/

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
					min-length: 5
			a:
				- abcabc
				- bcdbcd
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

	cb.TempFile("c.yaml", `
			type: tests:e
			id: c
			rules:
				-
					selector: ".a>{system:string}"
					type: system:@string
					min-length: 5
			a:
				- abcabc
				- abc
		`)

	errors, err = ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")

	cb.RemoveTempFile("c.yaml")

	cb.TempFile("d.yaml", `
			type: tests:e
			id: d
			rules:
				-
					selector: ".b>{system:string}"
					type: system:@string
					min-length: 5
			b:
				a: abcabc
				b: bcdbcd
		`)

	errors, err = ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

	cb.TempFile("e.yaml", `
			type: tests:e
			id: e
			rules:
				-
					selector: ".b>{system:string}"
					type: system:@string
					min-length: 5
			b:
				a: abcabc
				b: bcd
		`)

	errors, err = ValidatePackage(cb.Ctx())
	assert.IsError(t, errors[0], "HLKQWDCMRN")
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

	errors, err := ValidatePackage(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, 0, len(errors))
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

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")
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
					max-length: 5
		`,
	})

	cb.Path(path).Dir(dir).Alias("tests", "kego.io/process/validate/tests").Jauto().Sauto(parser.Parse)

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.IsError(t, errors[0], "HLKQWDCMRN")

}

func TestValidateNode(t *testing.T) {
	cb := tests.New()
	errors, err := ValidateNode(cb.Ctx(), &node.Node{Value: nil})
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))
	errors, err = ValidateNode(cb.Ctx(), &node.Node{Value: 1, Null: true})
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))
	errors, err = ValidateNode(cb.Ctx(), &node.Node{Value: 1, Missing: true})
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))
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

	errors, err := ValidatePackage(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, 0, len(errors))

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
					"min-length": 10,
					"max-length": 5
				}
			}
		}`,
	})

	cb.Path(path).Dir(dir).Jsystem().Sauto(parser.Parse)

	errors, err := ValidatePackage(cb.Ctx())
	// @string is invalid because minLength > maxLength
	require.NoError(t, err)
	assert.IsError(t, errors[0], "KULDIJUYFB")

}
