// info:{"Path":"kego.io/demo/demo4","Hash":5974341007325696005}
package demo4

import (
	"reflect"

	"github.com/davelondon/images"
	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}
type Page struct {
	*system.Object
	Hero  *images.Photo  `json:"hero"`
	Title *system.String `json:"title"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo4", 5974341007325696005)
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
}
