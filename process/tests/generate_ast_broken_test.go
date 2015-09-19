package tests

import (
	"os"
	"testing"

	"fmt"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/pkgtest"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

// TODO: this is WIP - currently stuck getting comments to line up.
// TODO: See http://stackoverflow.com/questions/31628613 for my question.
func SkipTestNonDestructiveGeneration(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	namespaceDir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespaceDir)

	path, _, _, err := pkgtest.CreateTemporaryPackage(namespaceDir, "a", map[string]string{
		"a.yaml": `
			type: system:type
			id: b
			fields:
				c:
					type: system:@string
				d:
					type: system:@number`,
		"b.go": `package a

		import "kego.io/system"

		// This is a struct comment
		type B struct {
			// This is a field comment
			C system.String
		}`,
	})

	verbose := false
	set, err := process.InitialiseManually(false, false, false, verbose, path)
	assert.NoError(t, err)

	s, err := process.GenerateSource(process.S_STRUCTS, set)
	assert.NoError(t, err)

	assert.Contains(t, string(s), "package a")
	fmt.Println(string(s))

}
