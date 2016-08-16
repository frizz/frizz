// info:{"Path":"kego.io/demo/common/words","Hash":15839668451341961644}
package words

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"context"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for localizer
type LocalizerRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for translation
type TranslationRule struct {
	*system.Object
	*system.Rule
}
type Simple struct {
	*system.Object
	String *system.String `json:"string"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}

// This represents a translated string
type Translation struct {
	*system.Object
	// The original English string
	English *system.String `json:"english"`
	// The translated strings
	Translations map[string]*system.String `json:"translations"`
}
type TranslationInterface interface {
	GetTranslation(ctx context.Context) *Translation
}

func (o *Translation) GetTranslation(ctx context.Context) *Translation {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/demo/common/words", 15839668451341961644)
	pkg.InitType("localizer", reflect.TypeOf((*Localizer)(nil)).Elem(), reflect.TypeOf((*LocalizerRule)(nil)), nil)
	pkg.InitType("simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
	pkg.InitType("translation", reflect.TypeOf((*Translation)(nil)), reflect.TypeOf((*TranslationRule)(nil)), reflect.TypeOf((*TranslationInterface)(nil)).Elem())
}
