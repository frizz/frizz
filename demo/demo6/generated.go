// info:{"Path":"kego.io/demo/demo6","Hash":15595427421812101520}
package demo6

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for foo
type FooRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}
type Foo struct {
	*system.Object
	Foo *system.String `json:"foo"`
}
type FooInterface interface {
	GetFoo(ctx context.Context) *Foo
}

func (o *Foo) GetFoo(ctx context.Context) *Foo {
	return o
}

type Page struct {
	*system.Object
	Title []*Foo `json:"title"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo6", 15595427421812101520)
	pkg.InitType("foo", reflect.TypeOf((*Foo)(nil)), reflect.TypeOf((*FooRule)(nil)), reflect.TypeOf((*FooInterface)(nil)).Elem())
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
}
