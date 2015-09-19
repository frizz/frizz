package words

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for localizer
type Localizer_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for simple
type Simple_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for translation
type Translation_rule struct {
	*system.Base
	*system.RuleBase
}
type Simple struct {
	*system.Base
	String system.String `json:"string"`
}

// This represents a translated string
type Translation struct {
	*system.Base
	// The original English string
	English system.String `kego:"{\"default\":{\"value\":\"http\"}}" json:"english"`
	// The translated strings
	Translations map[string]system.String `json:"translations"`
}

func init() {
	json.Register("kego.io/demo/common/words", "@localizer", reflect.TypeOf(&Localizer_rule{}), 0x72d207b91bb0393)
	json.Register("kego.io/demo/common/words", "@simple", reflect.TypeOf(&Simple_rule{}), 0xb367436bb1ddac59)
	json.Register("kego.io/demo/common/words", "@translation", reflect.TypeOf(&Translation_rule{}), 0x20d3acc377a84c88)
	json.Register("kego.io/demo/common/words", "simple", reflect.TypeOf(&Simple{}), 0xb367436bb1ddac59)
	json.Register("kego.io/demo/common/words", "translation", reflect.TypeOf(&Translation{}), 0x20d3acc377a84c88)
}
