package process_test

import (
	"testing"

	"io/ioutil"
	"path/filepath"

	"fmt"

	"regexp"

	"github.com/dave/ktest/require"
	"kego.io/process/generate"
	"kego.io/process/parser"
	"kego.io/tests"
)

func TestGenerateSystem(t *testing.T) {

	testGenerated(t, "kego.io/system")

}

func TestGenerateData(t *testing.T) {

	testGenerated(t, "kego.io/tests/data")

}

func testGenerated(t *testing.T, path string) {
	cb := tests.New().Path(path).Jauto().Sauto(parser.Parse)
	pi, err := parser.Parse(cb.Ctx(), path)
	require.NoError(t, err)

	generatedBytes, err := generate.Structs(cb.Ctx(), pi.Env)
	require.NoError(t, err)
	generatedString := string(generatedBytes)

	existingFilePath := filepath.Join(pi.Dir, "generated.go")
	existingBytes, err := ioutil.ReadFile(existingFilePath)
	require.NoError(t, err)
	existingString := string(existingBytes)

	// TODO: The "goimports" tool will often re-order the imports, so this is
	// a kludge to remove it before comparing. This is not ideal!
	importsRegex := regexp.MustCompile(`(?ms:\nimport \(\n.*\n\)\n)`)
	generatedString = importsRegex.ReplaceAllString(generatedString, "-")
	existingString = importsRegex.ReplaceAllString(existingString, "-")

	if generatedString != existingString {
		fmt.Println("Generated code for " + path + " is not what is present:")
		fmt.Println(generatedString)
	}
	require.Equal(t, generatedString, existingString)
}
