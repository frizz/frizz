// info:{"Path":"kego.io/packer/systests","Hash":9500717051976719187}
package systests

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"context"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for a
type ARule struct {
	*system.Object
	*system.Rule
}
type A struct {
	*system.Object
	A system.StringInterface `json:"a"`
	B system.NumberInterface `json:"b"`
	C system.BoolInterface   `json:"c"`
	D system.StringInterface `json:"d"`
}
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/packer/systests", 9500717051976719187)
	pkg.InitType("a", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())
}
