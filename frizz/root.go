package frizz

import (
	"fmt"
	"regexp"
	"strings"
)

func (r *Root) RegisterPackage(path string) string {
	for a, p := range r.Imports {
		if path == p {
			return a
		}
	}
	preferredAlias := guessAlias(path)
	alias := preferredAlias
	count := 0
	for {
		if _, ok := r.Imports[alias]; !ok {
			r.Imports[alias] = path
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
