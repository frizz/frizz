package editor

import (
	"github.com/gopherjs/gopherjs/js"
	"kego.io/system"
)

func Edit(path string) {
	g := system.GetAllGlobalsInPackage(path)
	for _, o := range g {
		js.Global.Get("document").Call("write", o.GetBase().Id.Value()+"<br>")
	}
}
