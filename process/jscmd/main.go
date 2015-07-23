package main

import (
	"fmt"

	_ "github.com/davelondon/ke_common/images"
	_ "github.com/davelondon/ke_common/images/types"
	_ "github.com/davelondon/ke_common/units"
	_ "github.com/davelondon/ke_common/units/types"
	_ "github.com/davelondon/ke_common/words"
	_ "github.com/davelondon/ke_common/words/types"
	"github.com/davelondon/ke_site"
	_ "github.com/davelondon/ke_site"
	_ "github.com/davelondon/ke_site/types"
	"github.com/gopherjs/gopherjs/js"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	js.Global.Get("document").Call("write", fmt.Sprint("Hello world! ", ke_site.Site3.Title.Localize([]string{"fr"})))
}
