package main

import (
	"kego.io/editor/client"
	_ "kego.io/editor/shared/messages"
	_ "kego.io/editor/shared/messages/types"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("kego.io/editor/shared/messages"); err != nil {
		console.Error(err.Error())
	}
}
