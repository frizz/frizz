package tests // import "kego.io/process/tests"

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
type Ctx struct {
	Path      string
	Aliases   map[string]string
	Dir       string
	Edit      bool
	Update    bool
	Recursive bool
	Log       bool
	Debug     bool
	Types     []*jsonctx.TypeInfo
	Dummies   bool
	Packages  []string
}

func PathCtx(path string) context.Context {
	return AllCtx(Ctx{Path: path})
}
func EnvCtx(path string, aliases map[string]string) context.Context {
	return AllCtx(Ctx{Path: path, Aliases: aliases})
}
func JsonTypeCtx(path string, name string, typ reflect.Type) context.Context {
	return JsonTypeAliasesCtx(path, name, typ, map[string]string{})
}
func JsonTypeAliasesCtx(path string, name string, typ reflect.Type, aliases map[string]string) context.Context {

	ctx := envctx.NewContext(context.Background(), &envctx.Env{
		Path:    path,
		Aliases: aliases,
	})
	ctx = jsonctx.ManualContext(ctx, false)

	jcache := jsonctx.FromContext(ctx)
	p := jcache.Packages.Set(path, 0)
	p.Types.Set(name, &jsonctx.TypeInfo{Name: name, Type: typ})

	return ctx
}
func SystemTypeCtx(path string, name string, typ interface{}) context.Context {

	env := &envctx.Env{
		Path:    path,
		Aliases: map[string]string{},
	}
	ctx := envctx.NewContext(context.Background(), env)
	ctx = cachectx.NewContext(ctx)

	cache := cachectx.FromContext(ctx)
	p := cache.Set(env)
	p.Types.Set(name, typ)

	return ctx
}
func AllCtx(c Ctx) context.Context {
	if c.Aliases == nil {
		c.Aliases = map[string]string{}
	}
	ctx := context.Background()
	ctx = wgctx.NewContext(ctx)
	ctx = envctx.NewContext(ctx, &envctx.Env{
		Path:      c.Path,
		Aliases:   c.Aliases,
		Recursive: c.Recursive,
	})
	ctx = cmdctx.NewContext(ctx, &cmdctx.Cmd{
		Dir:    c.Dir,
		Edit:   c.Edit,
		Update: c.Update,
		Log:    c.Log,
	})
	ctx = jsonctx.ManualContext(ctx, c.Dummies, c.Packages...)
	jcache := jsonctx.FromContext(ctx)
	p := jcache.Packages.Set(c.Path, 0)
	for _, t := range c.Types {
		p.Types.Set(t.Name, t)
	}
	return ctx
}*/

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
