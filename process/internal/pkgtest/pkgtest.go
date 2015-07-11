package pkgtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateTemporaryPackage creates a package directory and writes the files provided. We
// return the package path
func CreateTemporaryPackage(namespaceDir string, dirName string, files map[string]string) (string, error) {

	dir := filepath.Join(namespaceDir, dirName)
	os.MkdirAll(dir, 0700)
	os.Chdir(dir)

	for name, contents := range files {
		if err := ioutil.WriteFile(filepath.Join(dir, name), []byte(contents), 0700); err != nil {
			return "", err
		}
	}

	namespace := namespaceDir[strings.LastIndex(namespaceDir, string(os.PathSeparator))+1:]

	return fmt.Sprint(namespace, "/", dirName), nil

}

// CreateTemporaryNamespace creates a temporary namespace in your GOPATH/src directory. This
// should be deleted after use. We return the current dir and the namespace dir.
func CreateTemporaryNamespace() (currentDir string, namespaceDir string, err error) {

	currentDir, err = os.Getwd()
	if err != nil {
		return "", "", err
	}

	gopath := getCurrentGopath(currentDir, os.Getenv("GOPATH"))
	src := filepath.Join(gopath, "src")
	namespaceDir, err = ioutil.TempDir(src, "tmp")
	if err != nil {
		return "", "", err
	}

	return
}

func getCurrentGopath(dir string, gopathEnv string) string {
	gopaths := filepath.SplitList(gopathEnv)
	for _, gopath := range gopaths {
		if strings.HasPrefix(dir, gopath) {
			return gopath
		}
	}
	panic(fmt.Sprintf("Current dir %s is not in gopath %s", dir, gopathEnv))
}
