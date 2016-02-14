package main

// ke: {"package": {"notest": true}}

import (
	"log"
	"path/filepath"

	"strings"

	"fmt"

	"flag"

	"golang.org/x/net/context"
	"golang.org/x/tools/cover"
	"kego.io/cmd/gotests/scanner"
	"kego.io/cmd/gotests/tester"
	"kego.io/process/packages"
)

var baseDir string

func main() {

	all := flag.Bool("all", false, "If -all is spefified, gotests will always run all the tests")
	flag.Parse()

	var err error
	baseDir, err = packages.GetDirFromPackage(context.Background(), "kego.io")
	if err != nil {
		log.Fatal(err)
	}

	source, err := scanner.Get(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	//if err := tester.Js(source.JsTestPackages); err != nil {
	//	log.Fatal(err)
	//}

	var profiles []*cover.Profile
	if all != nil && *all {
		profiles, err = tester.Get(baseDir)
	} else {
		profiles, err = tester.GetSingle(baseDir, "kego.io/process/validate/command")
	}
	if err != nil {
		log.Fatal(err)
	}

	profilesMap := map[string]*cover.Profile{}
	for _, p := range profiles {
		profilesMap[p.FileName] = p
	}

	if err := excludeGenerated(profilesMap); err != nil {
		log.Fatal(err)
	}

	if err := excludeWraps(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := excludeBlocks(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := excludeFuncs(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := excludePackages(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := excludeSkips(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := tester.Save(profiles, filepath.Join(baseDir, "coverage.out")); err != nil {
		log.Fatal(err)
	}
}

func excludeGenerated(profiles map[string]*cover.Profile) error {
	for filename, _ := range profiles {
		if strings.HasSuffix(filename, "/generated.go") {
			fmt.Println("Excluding", filename)
			delete(profiles, filename)
			/*
				// summarize the original Profile
				statements := 0
				lastLine := 0
				lastCol := 0
				for _, pb := range p.Blocks {
					statements += pb.NumStmt
					if pb.EndLine >= lastLine {
						lastLine = pb.EndLine
						if pb.EndCol >= lastCol {
							lastCol = pb.EndCol
						}
					}
				}
				// overwrite with a single block
				p.Blocks = []cover.ProfileBlock{
					cover.ProfileBlock{
						StartLine: 0,
						StartCol:  0,
						EndLine:   lastLine,
						EndCol:    lastCol,
						NumStmt:   statements,
						Count:     1,
					},
				}
			*/
		}
	}

	return nil
}

func excludeWraps(profiles map[string]*cover.Profile, source *scanner.Source) error {
	for _, def := range source.Wraps {
		p, ok := profiles[def.File]
		if !ok {
			continue
		}
		for i, b := range p.Blocks {
			if b.StartLine <= def.Line && b.EndLine >= def.Line && b.Count != 1 {
				b.Count = 1
				p.Blocks[i] = b
				fmt.Printf("Excluding Wrap %s from %s:%d\n", def.Id, def.File, def.Line)
			}
		}
	}
	return nil
}

func excludeBlocks(profiles map[string]*cover.Profile, source *scanner.Source) error {
	for _, eb := range source.ExcludedBlocks {
		p, ok := profiles[eb.File]
		if !ok {
			continue
		}
		for i, b := range p.Blocks {
			if b.StartLine <= eb.Line && b.EndLine >= eb.Line && b.Count != 1 {
				b.Count = 1
				p.Blocks[i] = b
				fmt.Printf("Excluding block from %s:%d\n", eb.File, eb.Line)
			}
		}
	}
	return nil
}

func excludeFuncs(profiles map[string]*cover.Profile, source *scanner.Source) error {
	for _, ef := range source.ExcludedFuncs {
		p, ok := profiles[ef.File]
		if !ok {
			continue
		}
		for i, b := range p.Blocks {
			if b.StartLine <= ef.LineEnd && b.EndLine >= ef.LineStart && b.Count != 1 {
				b.Count = 1
				p.Blocks[i] = b
				fmt.Printf("Excluding func from %s:%d-%d\n", ef.File, ef.LineStart, ef.LineEnd)
			}
		}
	}
	return nil
}

func excludePackages(profiles map[string]*cover.Profile, source *scanner.Source) error {
	for path, _ := range source.ExcludedPackages {
		for filename, _ := range profiles {
			dir, _ := filepath.Split(filename)
			if strings.HasSuffix(dir, "/") {
				dir = dir[:len(dir)-1]
			}
			// dir is relative to the gopath/src, so we can just convert to slashes to get the
			// package path.
			pkg := filepath.ToSlash(dir)

			if path == pkg {
				delete(profiles, filename)
			}
		}
	}
	return nil
}

func excludeSkips(profiles map[string]*cover.Profile, source *scanner.Source) error {

	for id, _ := range source.Skipped {
		def, ok := source.All[id]
		if ok {
			p, ok := profiles[def.File]
			if !ok {
				continue
			}
			for i, b := range p.Blocks {
				if b.StartLine <= def.Line && b.EndLine >= def.Line && b.Count != 1 {
					b.Count = 1
					p.Blocks[i] = b
					fmt.Printf("Excluding skipped error %s from %s:%d\n", id, def.File, def.Line)
				}
			}
		}
	}
	return nil
}
