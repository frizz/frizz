package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729379968 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729380096 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729380224 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729161504 := &system.String_rule{Base: ptr8729380224, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729069760 := &system.Map_rule{Base: ptr8729380096, RuleBase: (*system.RuleBase)(nil), Items: ptr8729161504, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729380352 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729161680 := &system.String_rule{Base: ptr8729380352, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729380480 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729380608 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729282320 := &selectors.Kid_rule{Base: ptr8729380608, RuleBase: (*system.RuleBase)(nil)}

	ptr8729070976 := &system.Map_rule{Base: ptr8729380480, RuleBase: (*system.RuleBase)(nil), Items: ptr8729282320, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729380736 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729282416 := &selectors.Image_rule{Base: ptr8729380736, RuleBase: (*system.RuleBase)(nil)}

	ptr8729380864 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729380992 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729161856 := &system.String_rule{Base: ptr8729380992, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728945600 := &system.Array_rule{Base: ptr8729380864, RuleBase: (*system.RuleBase)(nil), Items: ptr8729161856, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729381120 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289664 := &system.Number_rule{Base: ptr8729381120, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729379840 := &system.Type{Base: ptr8729379968, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr8729069760, "favoriteColor": ptr8729161680, "kids": ptr8729070976, "avatar": ptr8729282416, "drinkPreference": ptr8728945600, "weight": ptr8729289664}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729059328 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729059456 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289280 := &system.Number_rule{Base: ptr8729059456, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729378816 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289376 := &system.Number_rule{Base: ptr8729378816, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729378944 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}

	ptr8729280704 := &selectors.C_rule{Base: ptr8729378944, RuleBase: (*system.RuleBase)(nil)}

	ptr8729379072 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729379200 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289472 := &system.Number_rule{Base: ptr8729379200, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729069120 := &system.Map_rule{Base: ptr8729379072, RuleBase: (*system.RuleBase)(nil), Items: ptr8729289472, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729379328 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729379456 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289568 := &system.Number_rule{Base: ptr8729379456, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729069632 := &system.Map_rule{Base: ptr8729379328, RuleBase: (*system.RuleBase)(nil), Items: ptr8729289568, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729173888 := &system.Type{Base: ptr8729059328, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr8729289280, "b": ptr8729289376, "c": ptr8729280704, "d": ptr8729069120, "e": ptr8729069632}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729458432 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729458560 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729458688 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729285088 := &selectors.Image_rule{Base: ptr8729458688, RuleBase: (*system.RuleBase)(nil)}

	ptr8729070016 := &system.Map_rule{Base: ptr8729458560, RuleBase: (*system.RuleBase)(nil), Items: ptr8729285088, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729458304 := &system.Type{Base: ptr8729458432, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr8729070016}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456768 := &system.Base{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729456640 := &system.Type{Base: ptr8729456768, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729458176 := &system.Base{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729458048 := &system.Type{Base: ptr8729458176, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729460608 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729172480 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728838784 := &system.RuleBase{Optional: true}

	ptr8729160976 := &system.String_rule{Base: ptr8729172480, RuleBase: ptr8728838784, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729172608 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729161152 := &system.String_rule{Base: ptr8729172608, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729172736 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729161328 := &system.String_rule{Base: ptr8729172736, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729460480 := &system.Type{Base: ptr8729460608, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"protocol": ptr8729160976, "server": ptr8729161152, "path": ptr8729161328}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729455360 := &system.Base{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729455232 := &system.Type{Base: ptr8729455360, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729381376 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729381248 := &system.Type{Base: ptr8729381376, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729457024 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729457152 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289088 := &system.Number_rule{Base: ptr8729457152, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729457280 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729289184 := &system.Number_rule{Base: ptr8729457280, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729457408 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729160096 := &system.String_rule{Base: ptr8729457408, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729457536 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729160272 := &system.String_rule{Base: ptr8729457536, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729457664 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729160448 := &system.String_rule{Base: ptr8729457664, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729457792 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729376128 := &system.Bool_rule{Base: ptr8729457792, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729457920 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729376384 := &system.Bool_rule{Base: ptr8729457920, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729456896 := &system.Type{Base: ptr8729457024, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"int": ptr8729289088, "float": ptr8729289184, "string": ptr8729160096, "string2": ptr8729160272, "null": ptr8729160448, "true": ptr8729376128, "false": ptr8729376384}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729454592 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729454848 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729288896 := &system.Number_rule{Base: ptr8729454848, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729454976 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729455104 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729288992 := &system.Number_rule{Base: ptr8729455104, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729069824 := &system.Map_rule{Base: ptr8729454976, RuleBase: (*system.RuleBase)(nil), Items: ptr8729288992, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729454720 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729288800 := &system.Number_rule{Base: ptr8729454720, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729454464 := &system.Type{Base: ptr8729454592, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"b": ptr8729288896, "c": ptr8729069824, "a": ptr8729288800}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459200 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459072 := &system.Type{Base: ptr8729459200, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729453184 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729452672 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729452800 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729452928 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729159040 := &system.String_rule{Base: ptr8729452928, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729067776 := &system.Map_rule{Base: ptr8729452800, RuleBase: (*system.RuleBase)(nil), Items: ptr8729159040, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728944960 := &system.Array_rule{Base: ptr8729452672, RuleBase: (*system.RuleBase)(nil), Items: ptr8729067776, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729453568 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729453696 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729159216 := &system.String_rule{Base: ptr8729453696, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728945040 := &system.Array_rule{Base: ptr8729453568, RuleBase: (*system.RuleBase)(nil), Items: ptr8729159216, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729453824 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729453952 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729159392 := &system.String_rule{Base: ptr8729453952, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728945120 := &system.Array_rule{Base: ptr8729453824, RuleBase: (*system.RuleBase)(nil), Items: ptr8729159392, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729454080 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729288704 := &system.Number_rule{Base: ptr8729454080, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729453312 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729453440 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729158688 := &system.String_rule{Base: ptr8729453440, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729071040 := &system.Map_rule{Base: ptr8729453312, RuleBase: (*system.RuleBase)(nil), Items: ptr8729158688, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729452544 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729158864 := &system.String_rule{Base: ptr8729452544, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729453056 := &system.Type{Base: ptr8729453184, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"languagesSpoken": ptr8728944960, "seatingPreference": ptr8728945040, "drinkPreference": ptr8728945120, "weight": ptr8729288704, "name": ptr8729071040, "favoriteColor": ptr8729158864}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459712 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459840 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729160624 := &system.String_rule{Base: ptr8729459840, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729459968 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729160800 := &system.String_rule{Base: ptr8729459968, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729460096 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729378304 := &system.RuleBase{Optional: true}

	ptr8729378208 := &system.Bool_rule{Base: ptr8729460096, RuleBase: ptr8729378304, Default: system.Bool{}}

	ptr8729459584 := &system.Type{Base: ptr8729459712, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr8729160624, "level": ptr8729160800, "preferred": ptr8729378208}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729455616 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729455744 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729455872 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729159744 := &system.String_rule{Base: ptr8729455872, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729069888 := &system.Map_rule{Base: ptr8729455744, RuleBase: (*system.RuleBase)(nil), Items: ptr8729159744, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729455488 := &system.Type{Base: ptr8729455616, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr8729069888}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729172992 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729172864 := &system.Type{Base: ptr8729172992, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729458944 := &system.Base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729458816 := &system.Type{Base: ptr8729458944, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456128 := &system.Base{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729456000 := &system.Type{Base: ptr8729456128, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729379712 := &system.Base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729379584 := &system.Type{Base: ptr8729379712, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729460352 := &system.Base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729460224 := &system.Type{Base: ptr8729460352, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456384 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729456512 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729159920 := &system.String_rule{Base: ptr8729456512, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729456256 := &system.Type{Base: ptr8729456384, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr8729159920}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729454336 := &system.Base{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729454208 := &system.Type{Base: ptr8729454336, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459456 := &system.Base{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459328 := &system.Type{Base: ptr8729459456, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729173760 := &system.Base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729173632 := &system.Type{Base: ptr8729173760, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729173248 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729173376 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729173504 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729287376 := &selectors.Kid_rule{Base: ptr8729173504, RuleBase: (*system.RuleBase)(nil)}

	ptr8728950320 := &system.Array_rule{Base: ptr8729173376, RuleBase: (*system.RuleBase)(nil), Items: ptr8729287376, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729173120 := &system.Type{Base: ptr8729173248, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr8728950320}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors", "typed", ptr8729379840)

	system.RegisterType("kego.io/selectors", "@typed", ptr8729381248)

	system.RegisterType("kego.io/selectors", "kid", ptr8729459584)

	system.RegisterType("kego.io/selectors", "@photo", ptr8729172864)

	system.RegisterType("kego.io/selectors", "@sibling", ptr8729379584)

	system.RegisterType("kego.io/selectors", "@polykids", ptr8729173632)

	system.RegisterType("kego.io/selectors", "gallery", ptr8729458304)

	system.RegisterType("kego.io/selectors", "@diagram", ptr8729456640)

	system.RegisterType("kego.io/selectors", "@c", ptr8729455232)

	system.RegisterType("kego.io/selectors", "expr", ptr8729456896)

	system.RegisterType("kego.io/selectors", "image", ptr8729459072)

	system.RegisterType("kego.io/selectors", "basic", ptr8729453056)

	system.RegisterType("kego.io/selectors", "@kid", ptr8729460224)

	system.RegisterType("kego.io/selectors", "@expr", ptr8729458048)

	system.RegisterType("kego.io/selectors", "@gallery", ptr8729458816)

	system.RegisterType("kego.io/selectors", "@basic", ptr8729454208)

	system.RegisterType("kego.io/selectors", "polykids", ptr8729173120)

	system.RegisterType("kego.io/selectors", "sibling", ptr8729173888)

	system.RegisterType("kego.io/selectors", "photo", ptr8729460480)

	system.RegisterType("kego.io/selectors", "c", ptr8729454464)

	system.RegisterType("kego.io/selectors", "collision", ptr8729455488)

	system.RegisterType("kego.io/selectors", "@collision", ptr8729456000)

	system.RegisterType("kego.io/selectors", "diagram", ptr8729456256)

	system.RegisterType("kego.io/selectors", "@image", ptr8729459328)

}

var _ selectors.Nothing
