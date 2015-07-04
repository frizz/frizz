package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728816256 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately"}
	ptr8729116672 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https"}
	ptr8728546816 := &system.RuleBase{Optional: true}
	ptr8728653072 := &system.String_rule{Base: ptr8729116672, RuleBase: ptr8728546816, Format: system.String{}, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729116800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com"}
	ptr8728653248 := &system.String_rule{Base: ptr8729116800, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729116928 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg"}
	ptr8728653424 := &system.String_rule{Base: ptr8729116928, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr8728854640 := &system.Type{Base: ptr8728816256, Fields: map[string]system.Rule{"protocol": ptr8728653072, "server": ptr8728653248, "path": ptr8728653424}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728818048 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}}
	ptr8728818176 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728818304 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651488 := &system.String_rule{Base: ptr8728818304, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728644416 := &system.Map_rule{Base: ptr8728818176, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651488, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728855536 := &system.Type{Base: ptr8728818048, Fields: map[string]system.Rule{"number": ptr8728644416}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729117568 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids"}
	ptr8728857104 := &system.Type{Base: ptr8729117568, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728817280 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}}
	ptr8728817664 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728817792 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729201088 := &system.Number_rule{Base: ptr8728817792, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728644352 := &system.Map_rule{Base: ptr8728817664, RuleBase: (*system.RuleBase)(nil), Items: ptr8729201088, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728817408 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729200896 := &system.Number_rule{Base: ptr8728817408, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728817536 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729200992 := &system.Number_rule{Base: ptr8728817536, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728855200 := &system.Type{Base: ptr8728817280, Fields: map[string]system.Rule{"c": ptr8728644352, "a": ptr8729200896, "b": ptr8729200992}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728820864 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images"}
	ptr8728820992 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728821120 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729065152 := &selectors.Image_rule{Base: ptr8728821120, RuleBase: (*system.RuleBase)(nil)}
	ptr8728644608 := &system.Map_rule{Base: ptr8728820992, RuleBase: (*system.RuleBase)(nil), Items: ptr8729065152, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728856320 := &system.Type{Base: ptr8728820864, Fields: map[string]system.Rule{"images": ptr8728644608}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729118720 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling"}
	ptr8728857328 := &system.Type{Base: ptr8729118720, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728821376 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}
	ptr8728856544 := &system.Type{Base: ptr8728821376, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true}
	ptr8728818944 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}}
	ptr8728819200 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729201280 := &system.Number_rule{Base: ptr8728819200, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728819328 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651840 := &system.String_rule{Base: ptr8728819328, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728819456 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652016 := &system.String_rule{Base: ptr8728819456, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728819584 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652192 := &system.String_rule{Base: ptr8728819584, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728819712 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729116352 := &system.Bool_rule{Base: ptr8728819712, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728819840 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728543328 := &system.Bool_rule{Base: ptr8728819840, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728819072 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729201184 := &system.Number_rule{Base: ptr8728819072, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728855984 := &system.Type{Base: ptr8728818944, Fields: map[string]system.Rule{"float": ptr8729201280, "string": ptr8728651840, "string2": ptr8728652016, "null": ptr8728652192, "true": ptr8729116352, "false": ptr8728543328, "int": ptr8729201184}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728817920 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c"}
	ptr8728855424 := &system.Type{Base: ptr8728817920, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728821632 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}}
	ptr8728633344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651312 := &system.String_rule{Base: ptr8728633344, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728815744 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652720 := &system.String_rule{Base: ptr8728815744, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728816000 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728545664 := &system.RuleBase{Optional: true}
	ptr8728545568 := &system.Bool_rule{Base: ptr8728816000, RuleBase: ptr8728545664, Default: system.Bool{}}
	ptr8728856768 := &system.Type{Base: ptr8728821632, Fields: map[string]system.Rule{"language": ptr8728651312, "level": ptr8728652720, "preferred": ptr8728545568}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728816384 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}}
	ptr8728820608 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729200800 := &system.Number_rule{Base: ptr8728820608, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728816512 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728816640 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649728 := &system.String_rule{Base: ptr8728816640, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728643456 := &system.Map_rule{Base: ptr8728816512, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728816768 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649904 := &system.String_rule{Base: ptr8728816768, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728816896 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728817024 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728817152 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652368 := &system.String_rule{Base: ptr8728817152, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728644032 := &system.Map_rule{Base: ptr8728817024, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652368, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729217856 := &system.Array_rule{Base: ptr8728816896, RuleBase: (*system.RuleBase)(nil), Items: ptr8728644032, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728820096 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728820224 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652544 := &system.String_rule{Base: ptr8728820224, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729217936 := &system.Array_rule{Base: ptr8728820096, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652544, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728820352 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728820480 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651136 := &system.String_rule{Base: ptr8728820480, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729218016 := &system.Array_rule{Base: ptr8728820352, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651136, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728854864 := &system.Type{Base: ptr8728816384, Fields: map[string]system.Rule{"weight": ptr8729200800, "name": ptr8728643456, "favoriteColor": ptr8728649904, "languagesSpoken": ptr8729217856, "seatingPreference": ptr8729217936, "drinkPreference": ptr8729218016}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728818560 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}
	ptr8728818688 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651664 := &system.String_rule{Base: ptr8728818688, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728855760 := &system.Type{Base: ptr8728818560, Fields: map[string]system.Rule{"url": ptr8728651664}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728818816 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram"}
	ptr8728855872 := &system.Type{Base: ptr8728818816, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729117056 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo"}
	ptr8728854752 := &system.Type{Base: ptr8729117056, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728818432 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision"}
	ptr8728855648 := &system.Type{Base: ptr8728818432, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728819968 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr"}
	ptr8728856096 := &system.Type{Base: ptr8728819968, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729118848 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}}
	ptr8729120000 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729198976 := &system.Number_rule{Base: ptr8729120000, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729118976 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729119104 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653600 := &system.String_rule{Base: ptr8729119104, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728642624 := &system.Map_rule{Base: ptr8729118976, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729119232 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653776 := &system.String_rule{Base: ptr8729119232, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729119360 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729119488 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729064688 := &selectors.Kid_rule{Base: ptr8729119488, RuleBase: (*system.RuleBase)(nil)}
	ptr8728642816 := &system.Map_rule{Base: ptr8729119360, RuleBase: (*system.RuleBase)(nil), Items: ptr8729064688, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729119616 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729064784 := &selectors.Image_rule{Base: ptr8729119616, RuleBase: (*system.RuleBase)(nil)}
	ptr8729119744 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729119872 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653952 := &system.String_rule{Base: ptr8729119872, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729216816 := &system.Array_rule{Base: ptr8729119744, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728857440 := &system.Type{Base: ptr8729118848, Fields: map[string]system.Rule{"weight": ptr8729198976, "name": ptr8728642624, "favoriteColor": ptr8728653776, "kids": ptr8728642816, "avatar": ptr8729064784, "drinkPreference": ptr8729216816}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729120128 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed"}
	ptr8728857552 := &system.Type{Base: ptr8729120128, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728820736 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic"}
	ptr8728855088 := &system.Type{Base: ptr8728820736, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729117184 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}}
	ptr8729117312 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729117440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729062304 := &selectors.Kid_rule{Base: ptr8729117440, RuleBase: (*system.RuleBase)(nil)}
	ptr8729215856 := &system.Array_rule{Base: ptr8729117312, RuleBase: (*system.RuleBase)(nil), Items: ptr8729062304, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728856880 := &system.Type{Base: ptr8729117184, Fields: map[string]system.Rule{"a": ptr8729215856}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728821248 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery"}
	ptr8728856432 := &system.Type{Base: ptr8728821248, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728821504 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image"}
	ptr8728856656 := &system.Type{Base: ptr8728821504, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728816128 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid"}
	ptr8728854528 := &system.Type{Base: ptr8728816128, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729117696 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}}
	ptr8729117824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729198592 := &system.Number_rule{Base: ptr8729117824, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729117952 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729198688 := &system.Number_rule{Base: ptr8729117952, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729118080 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}}
	ptr8729062992 := &selectors.C_rule{Base: ptr8729118080, RuleBase: (*system.RuleBase)(nil)}
	ptr8729118208 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729118336 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729198784 := &system.Number_rule{Base: ptr8729118336, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728641664 := &system.Map_rule{Base: ptr8729118208, RuleBase: (*system.RuleBase)(nil), Items: ptr8729198784, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729118464 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729118592 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729198880 := &system.Number_rule{Base: ptr8729118592, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728642240 := &system.Map_rule{Base: ptr8729118464, RuleBase: (*system.RuleBase)(nil), Items: ptr8729198880, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728857216 := &system.Type{Base: ptr8729117696, Fields: map[string]system.Rule{"a": ptr8729198592, "b": ptr8729198688, "c": ptr8729062992, "d": ptr8728641664, "e": ptr8728642240}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	system.RegisterType("kego.io/selectors", "photo", ptr8728854640)
	system.RegisterType("kego.io/selectors", "expr", ptr8728855984)
	system.RegisterType("kego.io/selectors", "basic", ptr8728854864)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728855872)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728854528)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728856432)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728857216)
	system.RegisterType("kego.io/selectors", "c", ptr8728855200)
	system.RegisterType("kego.io/selectors", "image", ptr8728856544)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728855648)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728856096)
	system.RegisterType("kego.io/selectors", "typed", ptr8728857440)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728857104)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728856320)
	system.RegisterType("kego.io/selectors", "kid", ptr8728856768)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728855760)
	system.RegisterType("kego.io/selectors", "@image", ptr8728856656)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728855088)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728856880)
	system.RegisterType("kego.io/selectors", "collision", ptr8728855536)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728857328)
	system.RegisterType("kego.io/selectors", "@c", ptr8728855424)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728854752)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728857552)
}
