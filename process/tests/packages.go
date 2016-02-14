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
	c.gopathInitialized = true
	c.OsVar("GOPATH", gopath)
	if sys {
		c.CopyToTemp("kego.io/system")
	}
	return c
}

func (c *ContextBuilder) CopyToTemp(path string) *ContextBuilder {
	if !c.gopathInitialized {
		c.Cleanup()
		panic(kerr.New("FCUPXYWDKW", "Gopath must be initialized with TempGopath or RealGopath").Error())
	}

	gopath := packages.GetCurrentGopath(c.Ctx())

	realDir, err := packages.GetDirFromPackage(context.Background(), path)
	if err != nil {
		c.Cleanup()
		panic(kerr.Wrap("XNLAQAHXKC", err).Error())
	}

	// path is a Go path, so we separate with /
	parts := append([]string{gopath, "src"}, strings.Split(path, "/")...)
	newDir := filepath.Join(parts...)

	os.MkdirAll(newDir, os.FileMode(0777))
	if err := copyDataFiles(realDir, newDir); err != nil {
		panic(kerr.Wrap("MEIKIUHPWN", err).Error())
	}
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
		if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
			// our Go files are tab indented, but yaml files don't like tabs.
			files[name] = strings.Replace(contents, "\t", "    ", -1)
		}
	}

	dir = filepath.Join(c.tempNamespace, name)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(kerr.Wrap("KJJMYPCMAL", err).Error())
	}

	c.tempPackageDir = dir
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

func (c *ContextBuilder) SetTemp(path string) (dir string) {
	if !c.gopathInitialized {
		panic("Gopath must be initialized with TempGopath or RealGopath")
	}
	if c.tempPackageDir != "" {
		panic("tempPackage must be empty.")
	}
	if c.tempNamespace != "" {
		panic("tempNamespace must be empty.")
	}

	slash := strings.LastIndex(c.tempNamespace, string(os.PathSeparator))
	namespace := c.tempNamespace[slash+1:]

	gopath := packages.GetCurrentGopath(c.Ctx())
	dir = filepath.Join(gopath, "src", path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		panic(fmt.Sprintf("Dir %s does not exist", dir))
	}
	c.tempNamespace = namespace
	c.tempPackageDir = dir
	return dir
}

func (c *ContextBuilder) RemoveTempFile(name string) *ContextBuilder {
	if c.tempPackageDir == "" {
		panic("Need to call tempPackage first.")
	}
	filename := filepath.Join(c.tempPackageDir, name)
	if _, err := os.Stat(filename); err != nil {
		panic(err.Error())
	}
	if err := os.Remove(filename); err != nil {
		panic(err.Error())
	}
	return c
}

func (c *ContextBuilder) TempFile(name string, contents string) *ContextBuilder {
	if c.tempPackageDir == "" {
		panic("Need to call tempPackage first.")
	}
	if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
		// our Go files are tab indented, but yaml files don't like tabs.
		contents = strings.Replace(contents, "\t", "    ", -1)
	}
	if err := ioutil.WriteFile(filepath.Join(c.tempPackageDir, name), []byte(contents), 0777); err != nil {
		panic(kerr.Wrap("UYQNYOCWMP", err).Error())
	}
	return c
}

func (c *ContextBuilder) Cleanup() {
	for _, dir := range c.tempDirs {
		os.RemoveAll(dir)
	}

}

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

func copyDataFiles(source string, dest string) (err error) {

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

		if !obj.IsDir() && hasCorrectExtension(obj.Name()) {
			// perform copy
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func hasCorrectExtension(name string) bool {
	return strings.HasSuffix(name, ".json") ||
		strings.HasSuffix(name, ".yml") ||
		strings.HasSuffix(name, ".yaml")
}
