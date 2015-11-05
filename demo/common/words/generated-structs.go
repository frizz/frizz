package words

import (
	"reflect"

	"kego.io/json"
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
	GetSimple() *Simple
}

func (o *Simple) GetSimple() *Simple {
	return o
}

// This represents a translated string
type Translation struct {
	*system.Object
	// The original English string
	English *system.String `kego:"{\"default\":{\"value\":\"http\"}}" json:"english"`
	// The translated strings
	Translations map[string]*system.String `json:"translations"`
}
type TranslationInterface interface {
	GetTranslation() *Translation
}

func (o *Translation) GetTranslation() *Translation {
	return o
}
func init() {
	json.Register("kego.io/demo/common/words", "@localizer", reflect.TypeOf((*LocalizerRule)(nil)), nil, 517105247315493779)
	json.Register("kego.io/demo/common/words", "@simple", reflect.TypeOf((*SimpleRule)(nil)), nil, 12927375385214954585)
	json.Register("kego.io/demo/common/words", "@translation", reflect.TypeOf((*TranslationRule)(nil)), nil, 2365424184825760904)
	json.Register("kego.io/demo/common/words", "localizer", reflect.TypeOf((*Localizer)(nil)).Elem(), nil, 517105247315493779)
	json.Register("kego.io/demo/common/words", "simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem(), 12927375385214954585)
	json.Register("kego.io/demo/common/words", "translation", reflect.TypeOf((*Translation)(nil)), reflect.TypeOf((*TranslationInterface)(nil)).Elem(), 2365424184825760904)
}
