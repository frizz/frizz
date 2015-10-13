package units

import (
	"reflect"

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
	Height system.Int `json:"height"`
	Width  system.Int `json:"width"`
}
type RectangleInterface interface {
	GetRectangle() *Rectangle
}

func (o *Rectangle) GetRectangle() *Rectangle {
	if o == nil {
		return &Rectangle{}
	}
	return o
}
func init() {
	json.Register("kego.io/demo/common/units", "@rectangle", reflect.TypeOf(&RectangleRule{}), nil, 16847424425470667022)
	json.Register("kego.io/demo/common/units", "rectangle", reflect.TypeOf(&Rectangle{}), reflect.TypeOf((*RectangleInterface)(nil)).Elem(), 16847424425470667022)
}
