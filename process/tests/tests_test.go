package tests

import (
	"fmt"
	"os"
	"testing"

	"kego.io/kerr/assert"
)

func TestRules(t *testing.T) {

	current, namespace, err := dirs()
	assert.NoError(t, err)
	defer os.Chdir(current)
	defer os.RemoveAll(namespace)

	_, err = pkg(namespace, "elements", map[string]string{
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

}

func TestFoo(t *testing.T) {

	current, namespace, err := dirs()
	assert.NoError(t, err)
	defer os.Chdir(current)
	defer os.RemoveAll(namespace)

	path, err := pkg(namespace, "elements", map[string]string{
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

	_, err = pkg(namespace, "site", map[string]string{
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
