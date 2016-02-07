package main

// ke: {"notest":true}

import (
	"log"
	"path/filepath"

	"strings"

	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/tools/cover"
	"kego.io/cmd/gotests/coverage"
	"kego.io/cmd/gotests/scanner"
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

	for _, w := range source.Wraps {
		p, ok := m[w.File]
		if !ok {
			continue
		}
		for _, b := range p.Blocks {
			if b.StartLine <= w.Line && b.EndLine >= w.Line && b.Count != 1 {
				b.Count = 1
			}
		}
	}
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
	profiles, err := coverage.Get(baseDir)
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
	err = coverage.Save(profiles, filepath.Join(baseDir, "coverage.out"))
	if err != nil {
		log.Fatal(err)
	}
}
