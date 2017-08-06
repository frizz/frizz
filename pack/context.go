package pack

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"frizz.io/global"
)

func NewDataContext(p global.PackageContext, r global.RootContext, l global.Location) dataContext {
	return dataContext{p, r, l}
}

type dataContext struct {
	p global.PackageContext
	r global.RootContext
	l global.Location
}

func (d dataContext) New(location global.Location) global.DataContext {
	return dataContext{d.p, d.r, location}
}

func (d dataContext) Package() global.PackageContext {
	return d.p
}

func (d dataContext) Root() global.RootContext {
	return d.r
}

func (d dataContext) Location() global.Location {
	return d.l
}

func NewValidationContext(p global.PackageContext, l global.Location) validationContext {
	return validationContext{p, l}
}

type validationContext struct {
	p global.PackageContext
	l global.Location
}

func (v validationContext) Package() global.PackageContext {
	return v.p
}

func (v validationContext) Location() global.Location {
	return v.l
}

func NewRootContext(p global.PackageContext) rootContext {
	return rootContext{
		p:       p,
		imports: map[string]string{},
	}
}

type rootContext struct {
	p       global.PackageContext
	imports map[string]string
}

func (r rootContext) Package() global.PackageContext {
	return r.p
}

func (r rootContext) Imports() map[string]string {
	return r.imports
}

func (r rootContext) ParseImports(v interface{}) error {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	im, ok := m["_import"]
	if !ok {
		return nil
	}
	imp, ok := im.(map[string]interface{})
	if !ok {
		return errors.Errorf("parsing _import, should be a map, found %T", im)
	}
	for alias, pathi := range imp {
		path, ok := pathi.(string)
		if !ok {
			return errors.Errorf("parsing _import, values should be strings, found %T", pathi)
		}
		r.imports[alias] = path
	}
	return nil
}

func (r rootContext) RegisterImport(path string) string {
	for a, p := range r.imports {
		if path == p {
			return a
		}
	}
	preferredAlias := guessAlias(path)
	alias := preferredAlias
	count := 0
	for {
		if _, ok := r.imports[alias]; !ok {
			r.imports[alias] = path
			return alias
		}
		count++
		alias = fmt.Sprintf("%s%d", preferredAlias, count)
	}
}

func guessAlias(path string) string {
	alias := path

	if strings.HasSuffix(alias, "/") {
		// training slashes are usually tolerated, so we can get rid of one if
		// it exists
		alias = alias[:len(alias)-1]
	}

	if strings.Contains(alias, "/") {
		// if the path contains a "/", use the last part
		alias = alias[strings.LastIndex(alias, "/")+1:]
	}

	// alias should be lower case
	alias = strings.ToLower(alias)

	// alias should now only contain alphanumerics
	importsRegex := regexp.MustCompile(`[^a-z0-9]`)
	alias = importsRegex.ReplaceAllString(alias, "")

	return alias
}

func NewPackageContext(path string, packages map[string]global.Package) packageContext {
	return packageContext{
		path:     path,
		packages: packages,
	}
}

type packageContext struct {
	path     string
	packages map[string]global.Package
}

func (p packageContext) Path() string {
	return p.path
}

func (p packageContext) Register(packages ...global.Package) {
	for _, pkg := range packages {
		p.packages[pkg.Path()] = pkg
	}
}

func (p packageContext) Get(path string) global.Package {
	pkg, ok := p.packages[path]
	if !ok {
		return nil
	}
	return pkg
}
