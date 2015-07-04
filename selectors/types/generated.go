package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728922624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed"}
	ptr8728922496 := &system.Type{Base: ptr8728922624, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728921216 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}}
	ptr8728921728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728921856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729139280 := &selectors.Kid_rule{Base: ptr8728921856, RuleBase: (*system.RuleBase)(nil)}
	ptr8728912320 := &system.Map_rule{Base: ptr8728921728, RuleBase: (*system.RuleBase)(nil), Items: ptr8729139280, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728921984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729139376 := &selectors.Image_rule{Base: ptr8728921984, RuleBase: (*system.RuleBase)(nil)}
	ptr8728922112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728922240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653952 := &system.String_rule{Base: ptr8728922240, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729038272 := &system.Array_rule{Base: ptr8728922112, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728922368 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728658880 := &system.Number_rule{Base: ptr8728922368, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728921344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728921472 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653600 := &system.String_rule{Base: ptr8728921472, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728912256 := &system.Map_rule{Base: ptr8728921344, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728921600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653776 := &system.String_rule{Base: ptr8728921600, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728921088 := &system.Type{Base: ptr8728921216, Fields: map[string]system.Rule{"kids": ptr8728912320, "avatar": ptr8729139376, "drinkPreference": ptr8729038272, "weight": ptr8728658880, "name": ptr8728912256, "favoriteColor": ptr8728653776}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728803840 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}
	ptr8728803712 := &system.Type{Base: ptr8728803840, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true}
	ptr8728805248 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately"}
	ptr8728820224 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https"}
	ptr8728548320 := &system.RuleBase{Optional: true}
	ptr8728653072 := &system.String_rule{Base: ptr8728820224, RuleBase: ptr8728548320, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728820352 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com"}
	ptr8728653248 := &system.String_rule{Base: ptr8728820352, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728820480 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg"}
	ptr8728653424 := &system.String_rule{Base: ptr8728820480, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728805120 := &system.Type{Base: ptr8728805248, Fields: map[string]system.Rule{"protocol": ptr8728653072, "server": ptr8728653248, "path": ptr8728653424}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728821504 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids"}
	ptr8728821376 := &system.Type{Base: ptr8728821504, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728801920 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}
	ptr8728802048 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651488 := &system.String_rule{Base: ptr8728802048, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728801792 := &system.Type{Base: ptr8728801920, Fields: map[string]system.Rule{"url": ptr8728651488}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728801152 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}}
	ptr8728801280 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728801408 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651312 := &system.String_rule{Base: ptr8728801408, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728912832 := &system.Map_rule{Base: ptr8728801280, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651312, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728801024 := &system.Type{Base: ptr8728801152, Fields: map[string]system.Rule{"number": ptr8728912832}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728797440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr"}
	ptr8728797312 := &system.Type{Base: ptr8728797440, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728803584 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery"}
	ptr8728803456 := &system.Type{Base: ptr8728803584, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728802560 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}}
	ptr8728803072 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651840 := &system.String_rule{Base: ptr8728803072, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728803200 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652016 := &system.String_rule{Base: ptr8728803200, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728803328 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728547104 := &system.Bool_rule{Base: ptr8728803328, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728797184 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728545024 := &system.Bool_rule{Base: ptr8728797184, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728802688 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728661280 := &system.Number_rule{Base: ptr8728802688, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728802816 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728661376 := &system.Number_rule{Base: ptr8728802816, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728802944 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651664 := &system.String_rule{Base: ptr8728802944, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728802432 := &system.Type{Base: ptr8728802560, Fields: map[string]system.Rule{"string2": ptr8728651840, "null": ptr8728652016, "true": ptr8728547104, "false": ptr8728545024, "int": ptr8728661280, "float": ptr8728661376, "string": ptr8728651664}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728804096 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image"}
	ptr8728803968 := &system.Type{Base: ptr8728804096, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728820992 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}}
	ptr8728821120 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728821248 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729137184 := &selectors.Kid_rule{Base: ptr8728821248, RuleBase: (*system.RuleBase)(nil)}
	ptr8729036992 := &system.Array_rule{Base: ptr8728821120, RuleBase: (*system.RuleBase)(nil), Items: ptr8729137184, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728820864 := &system.Type{Base: ptr8728820992, Fields: map[string]system.Rule{"a": ptr8729036992}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728800896 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c"}
	ptr8728800768 := &system.Type{Base: ptr8728800896, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728800128 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}}
	ptr8728800256 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728660992 := &system.Number_rule{Base: ptr8728800256, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728800384 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728661088 := &system.Number_rule{Base: ptr8728800384, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728800512 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728800640 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728661184 := &system.Number_rule{Base: ptr8728800640, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728912768 := &system.Map_rule{Base: ptr8728800512, RuleBase: (*system.RuleBase)(nil), Items: ptr8728661184, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728800000 := &system.Type{Base: ptr8728800128, Fields: map[string]system.Rule{"a": ptr8728660992, "b": ptr8728661088, "c": ptr8728912768}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728801664 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision"}
	ptr8728801536 := &system.Type{Base: ptr8728801664, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728802304 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram"}
	ptr8728802176 := &system.Type{Base: ptr8728802304, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728797696 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images"}
	ptr8728797824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728797952 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729133952 := &selectors.Image_rule{Base: ptr8728797952, RuleBase: (*system.RuleBase)(nil)}
	ptr8728912000 := &system.Map_rule{Base: ptr8728797824, RuleBase: (*system.RuleBase)(nil), Items: ptr8729133952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728797568 := &system.Type{Base: ptr8728797696, Fields: map[string]system.Rule{"images": ptr8728912000}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728804992 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid"}
	ptr8728804864 := &system.Type{Base: ptr8728804992, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728820736 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo"}
	ptr8728820608 := &system.Type{Base: ptr8728820736, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728799872 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic"}
	ptr8728799744 := &system.Type{Base: ptr8728799872, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728633472 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}}
	ptr8728920576 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728920704 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728658784 := &system.Number_rule{Base: ptr8728920704, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728912192 := &system.Map_rule{Base: ptr8728920576, RuleBase: (*system.RuleBase)(nil), Items: ptr8728658784, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728633344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728657920 := &system.Number_rule{Base: ptr8728633344, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728920064 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728658592 := &system.Number_rule{Base: ptr8728920064, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728920192 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}}
	ptr8729138032 := &selectors.C_rule{Base: ptr8728920192, RuleBase: (*system.RuleBase)(nil)}
	ptr8728920320 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728920448 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728658688 := &system.Number_rule{Base: ptr8728920448, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728912064 := &system.Map_rule{Base: ptr8728920320, RuleBase: (*system.RuleBase)(nil), Items: ptr8728658688, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728821632 := &system.Type{Base: ptr8728633472, Fields: map[string]system.Rule{"e": ptr8728912192, "a": ptr8728657920, "b": ptr8728658592, "c": ptr8729138032, "d": ptr8728912064}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728798208 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}}
	ptr8728798336 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728798464 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649728 := &system.String_rule{Base: ptr8728798464, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728912384 := &system.Map_rule{Base: ptr8728798336, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728798592 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652368 := &system.String_rule{Base: ptr8728798592, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728798720 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728798848 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728798976 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652544 := &system.String_rule{Base: ptr8728798976, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728912448 := &system.Map_rule{Base: ptr8728798848, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652544, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729037632 := &system.Array_rule{Base: ptr8728798720, RuleBase: (*system.RuleBase)(nil), Items: ptr8728912448, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728799104 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728799232 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652720 := &system.String_rule{Base: ptr8728799232, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729037712 := &system.Array_rule{Base: ptr8728799104, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652720, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728799360 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728799488 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649904 := &system.String_rule{Base: ptr8728799488, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729037792 := &system.Array_rule{Base: ptr8728799360, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649904, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728799616 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728660896 := &system.Number_rule{Base: ptr8728799616, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728798080 := &system.Type{Base: ptr8728798208, Fields: map[string]system.Rule{"name": ptr8728912384, "favoriteColor": ptr8728652368, "languagesSpoken": ptr8729037632, "seatingPreference": ptr8729037712, "drinkPreference": ptr8729037792, "weight": ptr8728660896}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728920960 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling"}
	ptr8728920832 := &system.Type{Base: ptr8728920960, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728804352 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}}
	ptr8728804608 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652896 := &system.String_rule{Base: ptr8728804608, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728804736 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728547584 := &system.RuleBase{Optional: true}
	ptr8728547488 := &system.Bool_rule{Base: ptr8728804736, RuleBase: ptr8728547584, Default: system.Bool{}}
	ptr8728804480 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652192 := &system.String_rule{Base: ptr8728804480, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728804224 := &system.Type{Base: ptr8728804352, Fields: map[string]system.Rule{"level": ptr8728652896, "preferred": ptr8728547488, "language": ptr8728652192}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	system.RegisterType("kego.io/selectors", "@typed", ptr8728922496)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728801792)
	system.RegisterType("kego.io/selectors", "c", ptr8728800000)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728797568)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728799744)
	system.RegisterType("kego.io/selectors", "image", ptr8728803712)
	system.RegisterType("kego.io/selectors", "photo", ptr8728805120)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728802176)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728820864)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728801536)
	system.RegisterType("kego.io/selectors", "basic", ptr8728798080)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728920832)
	system.RegisterType("kego.io/selectors", "kid", ptr8728804224)
	system.RegisterType("kego.io/selectors", "collision", ptr8728801024)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728803456)
	system.RegisterType("kego.io/selectors", "expr", ptr8728802432)
	system.RegisterType("kego.io/selectors", "@image", ptr8728803968)
	system.RegisterType("kego.io/selectors", "@c", ptr8728800768)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728804864)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728820608)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728821632)
	system.RegisterType("kego.io/selectors", "typed", ptr8728921088)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728821376)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728797312)
}
