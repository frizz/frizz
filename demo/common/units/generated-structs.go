package units

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/json"
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
	json.Register("kego.io/demo/common/units", "@rectangle", reflect.TypeOf((*RectangleRule)(nil)), nil, 16847424425470667022)
	json.Register("kego.io/demo/common/units", "rectangle", reflect.TypeOf((*Rectangle)(nil)), reflect.TypeOf((*RectangleInterface)(nil)).Elem(), 16847424425470667022)
}
