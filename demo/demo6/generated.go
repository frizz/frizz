// info:{"Path":"kego.io/demo/demo6","Hash":7465480149826331644}
package demo6

// ke: {"file": {"notest": true}}

import (
	"context"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for person
type PersonRule struct {
	*system.Object
	*system.Rule
}
type Page struct {
	*system.Object
	Heading *system.String `json:"heading"`
	People  []*Person      `json:"people"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}

type Person struct {
	*system.Object
	Age  *system.Int    `json:"age"`
	Name *system.String `json:"name"`
}
type PersonInterface interface {
	GetPerson(ctx context.Context) *Person
}

func (o *Person) GetPerson(ctx context.Context) *Person {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo6", 7465480149826331644)
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
	pkg.InitType("person", reflect.TypeOf((*Person)(nil)), reflect.TypeOf((*PersonRule)(nil)), reflect.TypeOf((*PersonInterface)(nil)).Elem())
}
