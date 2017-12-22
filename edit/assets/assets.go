// +build dev

package assets

import (
	"net/http"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

// Assets contains project assets.
var Assets http.FileSystem

func init() {
	dir, err := patsy.Dir(vos.Os(), "frizz.io/edit/assets/data")
	if err != nil {
		panic(err.Error())
	}
	Assets = http.Dir(dir)
}
