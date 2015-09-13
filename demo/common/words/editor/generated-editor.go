package main

import (
	_ "kego.io/demo/common/words"
	_ "kego.io/demo/common/words/types"
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("kego.io/demo/common/words"); err != nil {
		console.Error(err.Error())
	}
}
