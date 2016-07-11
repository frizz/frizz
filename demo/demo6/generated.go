// info:{"Path":"kego.io/demo/demo6","Hash":12288996034802306009}
package demo6

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for bar
type BarRule struct {
	*system.Object
	*system.Rule
}

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
type Bar struct {
	*system.Object
	Foo *Foo `json:"foo"`
}
type BarInterface interface {
	GetBar(ctx context.Context) *Bar
}

func (o *Bar) GetBar(ctx context.Context) *Bar {
	return o
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
	Bar *Bar `json:"bar"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo6", 12288996034802306009)
	pkg.InitType("bar", reflect.TypeOf((*Bar)(nil)), reflect.TypeOf((*BarRule)(nil)), reflect.TypeOf((*BarInterface)(nil)).Elem())
	pkg.InitType("foo", reflect.TypeOf((*Foo)(nil)), reflect.TypeOf((*FooRule)(nil)), reflect.TypeOf((*FooInterface)(nil)).Elem())
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
}
