package main

import (
	"fmt"
	"os"

	_ "kego.io/demo/common/images"
	_ "kego.io/demo/common/images/types"
	_ "kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	_ "kego.io/demo/common/words"
	_ "kego.io/demo/common/words/types"
	_ "kego.io/demo/site"
	_ "kego.io/demo/site/types"
	"kego.io/editor/server"
	"kego.io/process"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

// This starts the local editor server and opens a browser. This
// simulates the ke command without running it.
func main() {
	if err := server.Start("kego.io/demo/site", true, true); err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
}
