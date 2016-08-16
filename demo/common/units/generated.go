// info:{"Path":"kego.io/demo/common/units","Hash":10988868301072060435}
package units

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"context"

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
	pkg := jsonctx.InitPackage("kego.io/demo/common/units", 10988868301072060435)
	pkg.InitType("rectangle", reflect.TypeOf((*Rectangle)(nil)), reflect.TypeOf((*RectangleRule)(nil)), reflect.TypeOf((*RectangleInterface)(nil)).Elem())
}
