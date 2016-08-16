// info:{"Path":"kego.io/demo/demo3/images","Hash":10497271076607888210}
package images

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"context"

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
	pkg := jsonctx.InitPackage("kego.io/demo/demo3/images", 10497271076607888210)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
