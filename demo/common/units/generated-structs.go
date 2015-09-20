package units

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for rectangle
type Rectangle_rule struct {
	*system.Object_base
	*system.Rule_base
}
type Rectangle struct {
	*system.Object_base
	Height system.Int `json:"height"`
	Width  system.Int `json:"width"`
}

func init() {
	json.Register("kego.io/demo/common/units", "@rectangle", reflect.TypeOf(&Rectangle_rule{}), 0xe9ce1340e46b290e)
	json.Register("kego.io/demo/common/units", "rectangle", reflect.TypeOf(&Rectangle{}), 0xe9ce1340e46b290e)
}
