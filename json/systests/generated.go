// info:{"Path":"kego.io/json/systests","Hash":14515556620459574619}
package systests

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
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
}
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/json/systests", 14515556620459574619)
	pkg.InitType("a", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())
}
