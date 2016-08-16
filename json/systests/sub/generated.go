// info:{"Path":"kego.io/json/systests/sub","Hash":7354942973798892629}
package sub

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
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/json/systests/sub", 7354942973798892629)
	pkg.InitType("a", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())
}
