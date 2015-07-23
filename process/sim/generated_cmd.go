package main

import (
	"fmt"
	"os"

	_ "github.com/davelondon/ke_common/images"
	_ "github.com/davelondon/ke_common/images/types"
	_ "github.com/davelondon/ke_common/units"
	_ "github.com/davelondon/ke_common/units/types"
	_ "github.com/davelondon/ke_common/words"
	_ "github.com/davelondon/ke_common/words/types"
	_ "github.com/davelondon/ke_site"
	_ "github.com/davelondon/ke_site/types"
	"kego.io/process"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	update := false
	recursive := false
	globals := false
	path := "github.com/davelondon/ke_site"
	set, err := process.InitialiseCommand(update, recursive, globals, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := process.EditCommand(set); err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
}
