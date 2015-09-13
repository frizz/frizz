package main

import (
	_ "kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("kego.io/demo/common/units"); err != nil {
		console.Error(err.Error())
	}
}
