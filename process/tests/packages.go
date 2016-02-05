package tests

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/context"
	"kego.io/kerr"
	"kego.io/process/packages"
)

// RealGopath instructs the TempNamespace and TempPackage functions to use the real GOPATH
func (c *ContextBuilder) RealGopath() *ContextBuilder {
	c.gopathInitialized = true
	return c
}

// TempGopath creates a temporary GoPath, and optionally copies in the json files from system.
func (c *ContextBuilder) TempGopath(sys bool) *ContextBuilder {
	gopath, err := ioutil.TempDir("", "go")
	if err != nil {
		panic(kerr.Wrap("IGEJIBQJYB", err).Error())
	}
	c.tempDirs = append(c.tempDirs, gopath)
	os.Mkdir(filepath.Join(gopath, "src"), os.FileMode(0777))
	if sys {
		realSystemDir, err := packages.GetDirFromPackage(context.Background(), "kego.io/system")
		if err != nil {
			c.Cleanup()
			panic(kerr.Wrap("STHSVMNWLR", err).Error())
		}
		os.MkdirAll(filepath.Join(gopath, "src", "kego.io", "system"), os.FileMode(0777))
		copyDirJson(realSystemDir, filepath.Join(gopath, "src", "kego.io", "system"))
	}
	c.OsVar("GOPATH", gopath)
	c.gopathInitialized = true
	return c
}

func (c *ContextBuilder) TempNamespace() *ContextBuilder {
	if !c.gopathInitialized {
		panic(kerr.New("ITRXYQXCXD", "Gopath must be initialized with TempGopath or RealGopath").Error())
	}
	gopath := packages.GetCurrentGopath(c.Ctx())
	namespaceDir, err := ioutil.TempDir(filepath.Join(gopath, "src"), "ns")
	if err != nil {
		panic(kerr.Wrap("XHHBAUHAOT", err).Error())
	}
	c.tempDirs = append(c.tempDirs, namespaceDir)
	c.tempNamespace = namespaceDir
	return c
}

func (c *ContextBuilder) TempPackage(name string, files map[string]string) (path string, dir string) {

	if c.tempNamespace == "" {
		c.TempNamespace()
	}

	for name, contents := range files {
		if strings.HasSuffix(name, ".yaml") {
			// our Go files are tab indented, but yaml files don't like tabs.
			files[name] = strings.Replace(contents, "\t", "    ", -1)
		}
	}

	dir = filepath.Join(c.tempNamespace, name)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(kerr.Wrap("KJJMYPCMAL", err).Error())
	}

	c.tempDirs = append(c.tempDirs, dir)

	for name, contents := range files {
		if err := ioutil.WriteFile(filepath.Join(dir, name), []byte(contents), 0777); err != nil {
			panic(kerr.Wrap("AHDYGKKPID", err).Error())
		}
	}

	namespace := c.tempNamespace[strings.LastIndex(c.tempNamespace, string(os.PathSeparator))+1:]

	path = fmt.Sprint(namespace, "/", name)

	return
}

func (c *ContextBuilder) Cleanup() {
	for _, dir := range c.tempDirs {
		os.RemoveAll(dir)
	}
}

/*
func (c *ContextBuilder) TempNamespace() *ContextBuilder {
	c.TempNamespaceNamed("")
	return c
}

func (c *ContextBuilder) TempNamespaceNamed(name string) *ContextBuilder {
	gopath := GetCurrentGopath(ctx)

	srcDir := filepath.Join(gopath, "src")
	namespaceDir, err := ioutil.TempDir(srcDir, "tmp")
	if err != nil {
		return "", err
	}

	return namespaceDir, nil
	return c
}
*/

func copyFile(source string, dest string) (err error) {

	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return
}

func copyDirJson(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if !obj.IsDir() && strings.HasSuffix(obj.Name(), ".json") {
			// perform copy
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
