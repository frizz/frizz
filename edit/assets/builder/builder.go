package main

import (
	"log"

	"path/filepath"

	"frizz.io/edit/assets"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/shurcooL/vfsgen"
)

func main() {
	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/assets")
	if err != nil {
		log.Fatalln(err)
	}
	err = vfsgen.Generate(assets.Assets, vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
		Filename:     filepath.Join(dir, "assets_vfsdata.go"),
	})
	if err != nil {
		log.Fatalln(err)
	}
}
