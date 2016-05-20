package builder

import (
	"fmt"
	"sort"

	"github.com/davelondon/kerr"
)

type Import struct {
	// Go package path
	Path string
	// Alias used in the source, if different from the name
	Alias string
	// Package name - this is determined by executing "go list -f '{{.Name}}' [path]"
	Name string
}

type Imports map[string]Import

func (i Imports) Anonymous(path string) {
	i.add(path, true)
}

func (i Imports) Add(path string) string {
	return i.add(path, false)
}

func (i Imports) add(path string, anonymous bool) string {

	imp, found := i[path]
	if found && imp.Alias != "_" {
		// if the path already exists in the imports with an alias, we don't need to do anything.
		return imp.Alias
	} else if found && anonymous {
		// if the path already exists in the imports, and we're adding an anonymous import, we
		// don't need to do anything.
		return "_"
	}

	name := ""
	if found {
		name = imp.Name
	} else {
		name = getPackageName(path)
	}

	// either !found or (found && currentAlias == "_" && !anonymous)

	if anonymous {
		// if we're adding anonymously, we don't need to work out an alias
		i[path] = Import{Path: path, Alias: "_", Name: name}
		return "_"
	}

	// the package name is not guaranteed to be unique, so we find a unique alias by
	// appending digits until it is unique and not a keyword.
	alias := i.alias(name)

	i[path] = Import{Path: path, Alias: alias, Name: name}
	return alias
}

func (i Imports) packages() []string {
	paths := make([]string, len(i))
	count := 0
	for path, _ := range i {
		paths[count] = path
		count++
	}
	sort.Strings(paths)
	return paths
}

func (i Imports) alias(preferredAlias string) string {
	new := preferredAlias
	count := 0
	for {
		if count > 0 {
			new = fmt.Sprintf("%s%v", preferredAlias, count)
		}
		if !i.hasAlias(new) && !keyword[new] {
			return new
		}
		count++
		if count > 100 {
			// ke: {"block": {"notest": true}}
			panic(kerr.New("FOVRTYCGSI", "too many iterations").Error())
		}
	}
}

func (i Imports) hasAlias(alias string) bool {
	for _, a := range i {
		if a.Alias == alias {
			return true
		}
	}
	return false
}

var keyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}
