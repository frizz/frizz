package tests

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/internal/pkgtest"
)

func TestRules(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	currentDir, namespaceDir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)
	defer os.RemoveAll(namespaceDir)

	_, err = runKego(namespaceDir, "a", map[string]string{
		"gallery.yaml": `
			type: system:type
			id: gallery
			fields:
				str:
					type: system:@string
					rules:
						-
							selector: ":root"
							type: system:@string
							equal: foo`,
		"faces.yaml": `
			type: gallery
			id: faces
			str: foo`,
	})
	assert.NoError(t, err)

	_, err = runKego(namespaceDir, "b", map[string]string{
		"gallery.yaml": `
			type: system:type
			id: gallery
			fields:
				str:
					type: system:@string
					rules:
						-
							selector: ":root"
							type: system:@string
							equal: foo`,
		"faces.yaml": `
			type: gallery
			id: faces
			str: bar`,
	})
	assert.IsError(t, err, "UDDSSMQRHA")
	assert.Contains(t, err.Error(), "Equal: value must equal 'foo'")

}

func TestFoo(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping long-running end-to-end tests")
	}

	currentDir, namespaceDir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)
	defer os.RemoveAll(namespaceDir)

	path, err := runKego(namespaceDir, "elements", map[string]string{
		"title.yaml": `
			type: system:type
			id: title
			fields:
				text:
					type: system:@string`,
		"foo.yaml": `
			type: title
			id: foo
			text: Heading`,
		"page.yaml": `
			type: system:type
			id: page
			fields:
				title:
					type: "@title"`,
	})
	assert.NoError(t, err)

	_, err = runKego(namespaceDir, "site", map[string]string{
		"front.yaml": `
			type: elements:page
			id: front
			title:
				type: elements:title
				text: Heading`,
		"imports.yaml": fmt.Sprintf(`
			type: system:imports
			imports:
				%s: elements`, path),
		"b_test.go": `
			package site
			import (
				"testing"
				"kego.io/kerr/assert"
			)
			func TestMain(t *testing.T) {
				assert.Equal(t, "Heading", Front.Title.Text.Value)
			}`,
	})
	assert.NoError(t, err)
}

func runKego(namespaceDir string, packageDir string, files map[string]string) (string, error) {

	runTests := false
	for name, contents := range files {
		if strings.HasSuffix(name, ".yaml") {
			files[name] = strings.Replace(contents, "\t", "    ", -1)
		}
		if strings.HasSuffix(name, "_test.go") {
			runTests = true
		}
	}

	_, err := pkgtest.CreateTemporaryPackage(namespaceDir, packageDir, files)
	if err != nil {
		return "", err
	}

	dir, update, recursive, verbose, path, aliases, err := process.Initialise()
	if err != nil {
		return "", err
	}

	if err := process.KegoCmd(dir, update, recursive, verbose, path, aliases); err != nil {
		return "", err
	}

	if runTests {
		if out, err := exec.Command("go", "test").CombinedOutput(); err != nil {
			return "", fmt.Errorf("%s", string(out))
		}
	}

	return path, nil
}
