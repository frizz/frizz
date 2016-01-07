// info:{"Path":"kego.io/demo/common/units","Hash":1636266296275141462}
package units

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for rectangle
type RectangleRule struct {
	*system.Object
	*system.Rule
}
type Rectangle struct {
	*system.Object
	Height *system.Int `json:"height"`
	Width  *system.Int `json:"width"`
}
type RectangleInterface interface {
	GetRectangle(ctx context.Context) *Rectangle
}

func (o *Rectangle) GetRectangle(ctx context.Context) *Rectangle {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/common/units", 1636266296275141462)
	pkg.InitType("rectangle", reflect.TypeOf((*Rectangle)(nil)), reflect.TypeOf((*RectangleRule)(nil)), reflect.TypeOf((*RectangleInterface)(nil)).Elem())
}
