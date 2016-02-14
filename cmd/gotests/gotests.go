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
	"kego.io/kerr"
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
		profiles, err = tester.Get(baseDir)
		//profiles, err = tester.GetSingle(baseDir, "kego.io/system")
	}
	if err != nil {
		log.Fatal(err)
	}

	profilesMap := map[string]*cover.Profile{}
	for _, p := range profiles {
		profilesMap[p.FileName] = p
	}

	if err := excludeGenerated(&profiles); err != nil {
		log.Fatal(err)
	}

	if err := excludePackages(&profiles, source); err != nil {
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

	if err := excludeSkips(profilesMap, source); err != nil {
		log.Fatal(err)
	}

	if err := checkComplete(profiles, source); err != nil {
		log.Fatal(err)
	}

	if err := tester.Save(profiles, filepath.Join(baseDir, "coverage.out")); err != nil {
		log.Fatal(err)
	}
}

func checkComplete(profiles []*cover.Profile, source *scanner.Source) error {
	for _, profile := range profiles {
		pkg := getPackage(profile)
		_, ok := source.CompletePackages[pkg]
		if ok {
			for _, block := range profile.Blocks {
				if block.Count == 0 {
					return kerr.New("GNLYPXHTNF", "Untested code in %s:%d-%d", profile.FileName, block.StartLine, block.EndLine)
				}
			}
		}
	}
	return nil
}

func excludeGenerated(profiles *[]*cover.Profile) error {
	out := []*cover.Profile{}
	for _, profile := range *profiles {
		if strings.HasSuffix(profile.FileName, "/generated.go") {
			fmt.Println("Excluding", profile.FileName)
		} else {
			out = append(out, profile)
		}
	}
	*profiles = out

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

func excludePackages(profiles *[]*cover.Profile, source *scanner.Source) error {
	out := []*cover.Profile{}
	for _, profile := range *profiles {

		pkg := getPackage(profile)

		_, ok := source.ExcludedPackages[pkg]

		if ok {
			fmt.Printf("Excluding package %s - %s\n", pkg, profile.FileName)
		} else {
			out = append(out, profile)
		}
	}
	*profiles = out
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

func getPackage(profile *cover.Profile) string {
	dir, _ := filepath.Split(profile.FileName)
	if strings.HasSuffix(dir, "/") {
		dir = dir[:len(dir)-1]
	}
	// dir is relative to the gopath/src, so we can just convert to slashes to get the package path.
	return filepath.ToSlash(dir)
}
