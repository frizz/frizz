// deps produces a list of packages ordered by their level in the dependence tree. The include
// function allows irrelevant packages to be excluded.
package main

import (
	"go/build"

	"github.com/shurcooL/go/importgraphutil"

	"strings"

	"fmt"
)

func include(name string) bool {
	if !strings.HasPrefix(name, "kego.io/") {
		return false
	}
	if strings.HasPrefix(name, "kego.io/demo/") {
		return false
	}
	if strings.HasPrefix(name, "kego.io/process/validate/selectors/tests") {
		return false
	}
	if strings.HasSuffix(name, "/mocks") {
		return false
	}
	if name == "kego.io/kerr" {
		return false
	}
	if name == "kego.io/editor/client/console" {
		return false
	}

	return true
}

func main() {
	forward, reverse, errs := importgraphutil.BuildNoTests(&build.Default)
	if len(errs) > 0 {
		panic(fmt.Errorf("meh"))
	}

	forward = filter(forward)
	reverse = filter(reverse)

	pkglevel := map[string]int{}
	levels := []level{}

	for {
		next := getNextLevel(levels, forward)
		if next == nil {
			break
		}
		levels = append(levels, *next)
	}

	for i, level := range levels {
		for _, p := range level.packages {
			pkglevel[p.name] = i
		}
	}

	for i, level := range levels {
		fmt.Println("Level", i)
		fmt.Println("=======")
		for _, p := range level.packages {

			if len(reverse[p.name]) == 0 && i == 0 {
				continue
			}

			fmt.Print(p.name[8:])
			if len(p.imports) > 0 {
				fmt.Print(" (")
				first := true
				for j := i - 1; j >= 0; j-- {
					for n, _ := range p.imports {
						if pkglevel[n] == j {
							if !first {
								fmt.Print(", ")
							}
							fmt.Print(n[8:])
							first = false
						}
					}
				}
				fmt.Print(")")
			}
			//fmt.Print(" [", len(reverse[p.name]), "]")

			printReverseLinks := false
			if printReverseLinks {
				if len(reverse[p.name]) > 0 {
					fmt.Print(" [")
					first := true
					for n, _ := range reverse[p.name] {
						if !first {
							fmt.Print(", ")
						}
						fmt.Print(n[8:])
						first = false
					}
					fmt.Print("]")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

}

func getNextLevel(previous []level, packages map[string]map[string]bool) *level {

	next := &level{packages: []item{}}

Outer:
	for p, imports := range packages {
		//if the package has already been put into a previous level, ignore it
		for _, level := range previous {
			for _, p1 := range level.packages {
				if p == p1.name {
					continue Outer
				}
			}
		}
		if len(previous) == 0 {
			// if we're on the first level, we're only interested in packages
			// with no imports, so as soon as we find an import we can cancel
			if len(imports) == 0 {
				next.packages = append(next.packages, item{p, imports})
			}

			continue Outer
		}

		cantFindOneOfTheImports := false
		for i, _ := range imports {
			// Does this import occur in the previous levels?
			found := false
		Inner:
			for _, level := range previous {
				for _, item := range level.packages {
					if item.name == i {
						found = true
						continue Inner
					}
				}
			}
			if !found {
				cantFindOneOfTheImports = true
				break
			}
		}
		if cantFindOneOfTheImports {
			continue
		}
		next.packages = append(next.packages, item{p, imports})

	}
	if len(next.packages) > 0 {
		return next
	}
	return nil
}

type level struct {
	packages []item
}
type item struct {
	name    string
	imports map[string]bool
}

func filter(in map[string]map[string]bool) map[string]map[string]bool {
	filtered := map[string]map[string]bool{}

	for name, imports := range in {
		if include(name) {
			pkg := map[string]bool{}
			for imp, _ := range imports {
				if include(imp) {
					pkg[imp] = true
				}
			}
			filtered[name] = pkg
		}
	}
	return filtered
}
