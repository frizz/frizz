package main

import (
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/selectors/tests"
	_ "kego.io/selectors/tests/types"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("kego.io/selectors/tests"); err != nil {
		console.Error(err.Error())
	}
}
