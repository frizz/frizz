package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/context"
	"kego.io/kerr"
	"kego.io/process/packages"
)

// CreateTemporaryPackage creates a package directory and writes the files provided. We return the
// package path (e.g. github.com/x/y) and the full dir of the created directory
func CreateTemporaryPackage(ctx context.Context, namespace string, name string, files map[string]string) (path string, packageDir string, err error) {

	gopath := packages.GetCurrentGopath(ctx)
	namespaceDir := filepath.Join(gopath, "src", namespace)
	if err = os.MkdirAll(namespaceDir, 0700); err != nil {
		return "", "", kerr.Wrap("QANOAKBLTC", err)
	}
	subDir, err := ioutil.TempDir(namespaceDir, "js")
	if err != nil {
		return "", "", kerr.Wrap("XDVGJTBVNY", err)
	}

	packageDir = filepath.Join(subDir, name)
	if err = os.MkdirAll(packageDir, 0700); err != nil {
		return "", "", kerr.Wrap("SROBIUHXRA", err)
	}

	for name, contents := range files {
		if err = ioutil.WriteFile(filepath.Join(packageDir, name), []byte(contents), 0700); err != nil {
			return "", "", kerr.Wrap("WLFYDXRJVR", err)
		}
	}

	tempDirName := subDir[strings.LastIndex(subDir, string(os.PathSeparator))+1:]

	path = fmt.Sprint(namespace, "/", tempDirName, "/", name)

	return

}
