// info:{"Path":"kego.io/demo/demo5/translation","Hash":15213008964671100022}
package translation

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for localized
type LocalizedRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for smartling
type SmartlingRule struct {
	*system.Object
	*system.Rule
}
type Simple struct {
	*system.Object
	Text *system.String `json:"text"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}

type Smartling struct {
	*system.Object
	English      *system.String            `json:"english"`
	Translations map[string]*system.String `json:"translations"`
}
type SmartlingInterface interface {
	GetSmartling(ctx context.Context) *Smartling
}

func (o *Smartling) GetSmartling(ctx context.Context) *Smartling {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/demo5/translation", 15213008964671100022)
	pkg.InitType("localized", reflect.TypeOf((*Localized)(nil)).Elem(), reflect.TypeOf((*LocalizedRule)(nil)), nil)
	pkg.InitType("simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
	pkg.InitType("smartling", reflect.TypeOf((*Smartling)(nil)), reflect.TypeOf((*SmartlingRule)(nil)), reflect.TypeOf((*SmartlingInterface)(nil)).Elem())
}
