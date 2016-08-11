// info:{"Path":"kego.io/demo/demo7/images","Hash":14109639835138036999}
package images

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
}
type Photo struct {
	*system.Object
	Height *system.Int    `json:"height"`
	Url    *system.String `json:"url"`
	Width  *system.Int    `json:"width"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo7/images", 14109639835138036999)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
