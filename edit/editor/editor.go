package editor

import (
	"fmt"

	"frizz.io/global"
)

func Edit(p global.Package) {
	fmt.Println(p.Path())
}
