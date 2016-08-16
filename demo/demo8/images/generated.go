// info:{"Path":"kego.io/demo/demo8/images","Hash":6005185369334051609}
package images

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"context"

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
	pkg := jsonctx.InitPackage("kego.io/demo/demo8/images", 6005185369334051609)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
