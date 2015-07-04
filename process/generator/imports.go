package generator

import (
	"fmt"
	"sort"
	"strings"
)

type imports map[string]string

// Only used in tests
func NewImports_test(args ...map[string]string) imports {
	if len(args) > 0 {
		return imports(args[0])
	}
	return imports{}
}

// Only used in tests
func (i imports) Add_test(path string) string {
	return i.add(path, false)
}

// Only used in tests
func (i imports) Anon_test(path string) string {
	return i.add(path, true)
}

func (i imports) add(path string, anonymous bool) string {

	currentAlias, found := i[path]
	if found && currentAlias != "_" {
		// if the path already exists in the imports with an alias, we don't need to do anything.
		return currentAlias
	} else if found && anonymous {
		// if the path already exists in the imports, and we're adding an anonymous import, we
		// don't need to do anything.
		return "_"
	}

	// either !found or (found && currentAlias == "_" && !anonymous)

	if anonymous {
		// if we're adding anonymously, we don't need to work out an alias
		i[path] = "_"
		return "_"
	}

	// lets find a preferred alias
	preferred := path
	if strings.Contains(path, "/") {
		preferred = path[strings.LastIndex(path, "/")+1:]
	}

	// the preferred alias is not guaranteed to be unique, so we find a unique alias by
	// appending digits until it is unique and not a keyword.
	alias := i.alias(preferred)

	i[path] = alias
	return alias
}

func (i imports) packages() []string {
	paths := make([]string, len(i))
	count := 0
	for path, _ := range i {
		paths[count] = path
		count++
	}
	sort.Strings(paths)
	return paths
}

func (i imports) alias(preferredAlias string) string {
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
			panic("too many iterations")
		}
	}
}

func (i imports) hasAlias(alias string) bool {
	for _, a := range i {
		if a == alias {
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
