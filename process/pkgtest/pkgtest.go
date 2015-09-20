package pkgtest // import "kego.io/process/pkgtest"

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateTemporaryPackage creates a package directory and writes the files provided. We
// return the package path (e.g. github.com/x/y) and the full dir of the created directory
func CreateTemporaryPackage(namespaceDir string, dirName string, files map[string]string) (path string, dir string, tests bool, err error) {

	for name, contents := range files {
		if strings.HasSuffix(name, ".yaml") {
			// our Go files are tab indented, but yaml files don't like tabs.
			files[name] = strings.Replace(contents, "\t", "    ", -1)
		}
		if strings.HasSuffix(name, "_test.go") {
			// if we add a xxx_test.go file we should also run "go test"
			tests = true
		}
	}

	dir = filepath.Join(namespaceDir, dirName)
	if err = os.MkdirAll(dir, 0700); err != nil {
		return
	}

	for name, contents := range files {
		if err = ioutil.WriteFile(filepath.Join(dir, name), []byte(contents), 0700); err != nil {
			return
		}
	}

	namespace := namespaceDir[strings.LastIndex(namespaceDir, string(os.PathSeparator))+1:]

	path = fmt.Sprint(namespace, "/", dirName)

	return

}

// CreateTemporaryNamespace creates a temporary namespace in your GOPATH/src directory. This
// should be deleted after use. We return the newly created dir.
func CreateTemporaryNamespace() (string, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	srcDir := filepath.Join(getCurrentGopath(currentDir, os.Getenv("GOPATH")), "src")
	namespaceDir, err := ioutil.TempDir(srcDir, "tmp")
	if err != nil {
		return "", err
	}

	return namespaceDir, nil
}

func getCurrentGopath(currentDir string, gopathEnv string) string {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		if strings.HasPrefix(currentDir, gopath) {
			return gopath
		}
	}
	return gopaths[0]
}
