package main

import (
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("kego.io/system"); err != nil {
		console.Error(err.Error())
	}
}
