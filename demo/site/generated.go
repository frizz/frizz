// info:{"Path":"kego.io/demo/site","Hash":3726678967244446909}
package site

// ke: {"file": {"notest": true}}

import (
	"context"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/demo/common/images"
	"kego.io/system"
)

// Automatically created basic rule for body
type BodyRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for columns
type ColumnsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for hero
type HeroRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for page
type PageRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for section
type SectionRule struct {
	*system.Object
	*system.Rule
}
type Body struct {
	*system.Object
	Align *system.String         `json:"align"`
	Copy  system.StringInterface `json:"copy"`
	Title system.StringInterface `json:"title"`
}
type BodyInterface interface {
	GetBody(ctx context.Context) *Body
}

func (o *Body) GetBody(ctx context.Context) *Body {
	return o
}

type Columns struct {
	*system.Object
	Columns []Section `json:"columns"`
}
type ColumnsInterface interface {
	GetColumns(ctx context.Context) *Columns
}

func (o *Columns) GetColumns(ctx context.Context) *Columns {
	return o
}

type Hero struct {
	*system.Object
	Head    system.StringInterface `json:"head"`
	Image   images.Image           `json:"image"`
	Subhead system.StringInterface `json:"subhead"`
}
type HeroInterface interface {
	GetHero(ctx context.Context) *Hero
}

func (o *Hero) GetHero(ctx context.Context) *Hero {
	return o
}

type Page struct {
	*system.Object
	Sections []Section              `json:"sections"`
	Title    system.StringInterface `json:"title"`
}
type PageInterface interface {
	GetPage(ctx context.Context) *Page
}

func (o *Page) GetPage(ctx context.Context) *Page {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/site", 3726678967244446909)
	pkg.InitType("body", reflect.TypeOf((*Body)(nil)), reflect.TypeOf((*BodyRule)(nil)), reflect.TypeOf((*BodyInterface)(nil)).Elem())
	pkg.InitType("columns", reflect.TypeOf((*Columns)(nil)), reflect.TypeOf((*ColumnsRule)(nil)), reflect.TypeOf((*ColumnsInterface)(nil)).Elem())
	pkg.InitType("hero", reflect.TypeOf((*Hero)(nil)), reflect.TypeOf((*HeroRule)(nil)), reflect.TypeOf((*HeroInterface)(nil)).Elem())
	pkg.InitType("page", reflect.TypeOf((*Page)(nil)), reflect.TypeOf((*PageRule)(nil)), reflect.TypeOf((*PageInterface)(nil)).Elem())
	pkg.InitType("section", reflect.TypeOf((*Section)(nil)).Elem(), reflect.TypeOf((*SectionRule)(nil)), nil)
}
