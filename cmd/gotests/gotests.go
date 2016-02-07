package main

// ke: {"package": {"notest":true}}

import (
	"log"
	"path/filepath"

	"strings"

	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/tools/cover"
	"kego.io/cmd/gotests/scanner"
	"kego.io/cmd/gotests/tester"
	"kego.io/process/packages"
)

var baseDir string

func excludeWrap(profiles []*cover.Profile) error {
	source, err := scanner.Get(baseDir)
	if err != nil {
		return err
	}
	m := map[string]*cover.Profile{}
	for _, p := range profiles {
		m[p.FileName] = p
	}

	for _, def := range source.Wraps {
		p, ok := m[def.File]
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

	for _, pos := range source.Notests {
		p, ok := m[pos.File]
		if !ok {
			continue
		}
		for i, b := range p.Blocks {
			if b.StartLine <= pos.Line && b.EndLine >= pos.Line && b.Count != 1 {
				b.Count = 1
				p.Blocks[i] = b
				fmt.Printf("Excluding block from %s:%d\n", pos.File, pos.Line)
			}
		}
	}

	/*
		for id, _ := range source.Skipped {
			def, ok := source.All[id]
			if ok {
				p, ok := m[def.File]
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
	*/
	return err
}
func excludeGenerated(profiles []*cover.Profile) error {
	for _, p := range profiles {
		if strings.HasSuffix(p.FileName, "/generated.go") {
			fmt.Println("Excluding", p.FileName)
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
		}
	}

	return nil
}

func main() {
	var err error
	baseDir, err = packages.GetDirFromPackage(context.Background(), "kego.io")
	if err != nil {
		log.Fatal(err)
	}
	profiles, err := tester.Get(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	err = excludeGenerated(profiles)
	if err != nil {
		log.Fatal(err)
	}
	err = excludeWrap(profiles)
	if err != nil {
		log.Fatal(err)
	}
	err = tester.Save(profiles, filepath.Join(baseDir, "coverage.out"))
	if err != nil {
		log.Fatal(err)
	}
}
