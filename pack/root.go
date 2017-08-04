package pack

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func NewRoot() *root {
	return &root{
		imports: map[string]string{},
	}
}

type root struct {
	imports map[string]string
}

func (r *root) Imports() map[string]string {
	return r.imports
}

func (r *root) Parse(v interface{}) error {
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

func (r *root) Register(path string) string {
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
