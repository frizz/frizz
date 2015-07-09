package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"kego.io/process"
)

func pkg(namespace string, pkgName string, files map[string]string) (string, error) {

	dir := filepath.Join(namespace, pkgName)
	os.MkdirAll(dir, 0700)
	os.Chdir(dir)

	runTests := false
	for name, contents := range files {
		if strings.HasSuffix(name, ".yaml") {
			contents = strings.Replace(contents, "\t", "    ", -1)
		}
		if strings.HasSuffix(name, "_test.go") {
			runTests = true
		}
		if err := ioutil.WriteFile(filepath.Join(dir, name), []byte(contents), 0700); err != nil {
			return "", err
		}
	}

	dir, test, recursive, _, path, imports, err := process.Initialise()
	if err != nil {
		return "", err
	}

	if err := process.KegoCmd(dir, test, recursive, true, path, imports); err != nil {
		return "", err
	}

	if runTests {
		if out, err := exec.Command("go", "test").CombinedOutput(); err != nil {
			return "", fmt.Errorf("%s", string(out))
		}
	}

	return path, nil
}

func dirs() (current string, namespace string, err error) {

	current, err = os.Getwd()
	if err != nil {
		return
	}

	gopath := getCurrentGopath(current, os.Getenv("GOPATH"))
	src := filepath.Join(gopath, "src")
	namespace, err = ioutil.TempDir(src, "tmp")
	if err != nil {
		return
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
