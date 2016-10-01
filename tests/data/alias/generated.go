// info:{"Path":"kego.io/tests/data/alias","Hash":3681006963164671295}
package alias

// ke: {"file": {"notest": true}}

import (
	"context"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for alms
type AlmsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for main
type MainRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}
type Alms map[string]*Simple
type AlmsInterface interface {
	GetAlms(ctx context.Context) *Alms
}

func (o Alms) GetAlms(ctx context.Context) Alms {
	return o
}

type Main struct {
	*system.Object
	A Alms `json:"a"`
}
type MainInterface interface {
	GetMain(ctx context.Context) *Main
}

func (o *Main) GetMain(ctx context.Context) *Main {
	return o
}

type Simple struct {
	*system.Object
	Js string `json:"js"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/tests/data/alias", 3681006963164671295)
	pkg.InitType("alms", reflect.TypeOf((*Alms)(nil)), reflect.TypeOf((*AlmsRule)(nil)), reflect.TypeOf((*AlmsInterface)(nil)).Elem())
	pkg.InitType("main", reflect.TypeOf((*Main)(nil)), reflect.TypeOf((*MainRule)(nil)), reflect.TypeOf((*MainInterface)(nil)).Elem())
	pkg.InitType("simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
}
