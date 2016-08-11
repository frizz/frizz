// info:{"Path":"kego.io/demo/demo7","Hash":16224128938049598805}
package demo7

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/demo/demo7/images"
	"kego.io/system"
)

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}
type Page struct {
	*system.Object
	Hero *images.Photo `json:"hero"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo7", 16224128938049598805)
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
}
