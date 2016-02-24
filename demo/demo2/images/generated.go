// info:{"Path":"kego.io/demo/demo2/images","Hash":13983040157067132389}
package images

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
	Url *system.String `json:"url"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo2/images", 13983040157067132389)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
