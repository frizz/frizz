// info:{"Path":"kego.io/demo/demo7/images","Hash":12815993352109631826}
package images

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

type PhotoRule struct {
	*system.Object
	*system.Rule
	Big *system.Bool `json:"big"`
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
	pkg := jsonctx.InitPackage("kego.io/demo/demo7/images", 12815993352109631826)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
