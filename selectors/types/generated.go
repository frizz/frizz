package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728820864 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728820992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728821120 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729081344 := &selectors.Image_rule{Base: ptr8728821120, RuleBase: (*system.RuleBase)(nil)}
	ptr8728636480 := &system.Map_rule{Base: ptr8728820992, RuleBase: (*system.RuleBase)(nil), Items: ptr8729081344, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728491600 := &system.Type{Base: ptr8728820864, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728636480}}
	ptr8729068032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729068160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729068288 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729078144 := &selectors.Kid_rule{Base: ptr8729068288, RuleBase: (*system.RuleBase)(nil)}
	ptr8729174896 := &system.Array_rule{Base: ptr8729068160, RuleBase: (*system.RuleBase)(nil), Items: ptr8729078144, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728492384 := &system.Type{Base: ptr8729068032, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729174896}}
	ptr8728820736 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil)}
	ptr8728491376 := &system.Type{Base: ptr8728820736, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8729069696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729069824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729069952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653600 := &system.String_rule{Base: ptr8729069952, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728634560 := &system.Map_rule{Base: ptr8729069824, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729070080 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653776 := &system.String_rule{Base: ptr8729070080, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729070208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729070336 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729080816 := &selectors.Kid_rule{Base: ptr8729070336, RuleBase: (*system.RuleBase)(nil)}
	ptr8728634688 := &system.Map_rule{Base: ptr8729070208, RuleBase: (*system.RuleBase)(nil), Items: ptr8729080816, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729070464 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729080960 := &selectors.Image_rule{Base: ptr8729070464, RuleBase: (*system.RuleBase)(nil)}
	ptr8729070592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729070720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728653952 := &system.String_rule{Base: ptr8729070720, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729175936 := &system.Array_rule{Base: ptr8729070592, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729070848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729166208 := &system.Number_rule{Base: ptr8729070848, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728492944 := &system.Type{Base: ptr8729069696, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8728634560, "favoriteColor": ptr8728653776, "kids": ptr8728634688, "avatar": ptr8729080960, "drinkPreference": ptr8729175936, "weight": ptr8729166208}}
	ptr8728817920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil)}
	ptr8728490368 := &system.Type{Base: ptr8728817920, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728819584 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil)}
	ptr8728491040 := &system.Type{Base: ptr8728819584, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728821504 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image", Rules: []system.Rule(nil)}
	ptr8728492160 := &system.Type{Base: ptr8728821504, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728816384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728817280 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728817408 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651312 := &system.String_rule{Base: ptr8728817408, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729176256 := &system.Array_rule{Base: ptr8728817280, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651312, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728817536 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728817664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651488 := &system.String_rule{Base: ptr8728817664, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729176336 := &system.Array_rule{Base: ptr8728817536, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651488, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728817792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167360 := &system.Number_rule{Base: ptr8728817792, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728816512 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728816640 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728649728 := &system.String_rule{Base: ptr8728816640, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728635840 := &system.Map_rule{Base: ptr8728816512, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728816768 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728649904 := &system.String_rule{Base: ptr8728816768, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728816896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728817024 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728817152 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651136 := &system.String_rule{Base: ptr8728817152, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728635904 := &system.Map_rule{Base: ptr8728817024, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651136, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729176176 := &system.Array_rule{Base: ptr8728816896, RuleBase: (*system.RuleBase)(nil), Items: ptr8728635904, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728490144 := &system.Type{Base: ptr8728816384, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"seatingPreference": ptr8729176256, "drinkPreference": ptr8729176336, "weight": ptr8729167360, "name": ptr8728635840, "favoriteColor": ptr8728649904, "languagesSpoken": ptr8729176176}}
	ptr8729067904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8728490032 := &system.Type{Base: ptr8729067904, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728819328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728819456 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652016 := &system.String_rule{Base: ptr8728819456, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728490928 := &system.Type{Base: ptr8728819328, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728652016}}
	ptr8728816256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8729067520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8728946112 := &system.RuleBase{Optional: true}
	ptr8728653072 := &system.String_rule{Base: ptr8729067520, RuleBase: ptr8728946112, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729067648 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728653248 := &system.String_rule{Base: ptr8729067648, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8729067776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728653424 := &system.String_rule{Base: ptr8729067776, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}}
	ptr8728489808 := &system.Type{Base: ptr8728816256, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728653072, "server": ptr8728653248, "path": ptr8728653424}}
	ptr8728819712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728819840 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167744 := &system.Number_rule{Base: ptr8728819840, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728819968 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167840 := &system.Number_rule{Base: ptr8728819968, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728820096 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652192 := &system.String_rule{Base: ptr8728820096, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728820224 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652368 := &system.String_rule{Base: ptr8728820224, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728820352 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652544 := &system.String_rule{Base: ptr8728820352, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728820480 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728945280 := &system.Bool_rule{Base: ptr8728820480, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728820608 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728945536 := &system.Bool_rule{Base: ptr8728820608, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728491152 := &system.Type{Base: ptr8728819712, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"int": ptr8729167744, "float": ptr8729167840, "string": ptr8728652192, "string2": ptr8728652368, "null": ptr8728652544, "true": ptr8728945280, "false": ptr8728945536}}
	ptr8728818048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728818176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167456 := &system.Number_rule{Base: ptr8728818176, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728818304 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167552 := &system.Number_rule{Base: ptr8728818304, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728818432 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728818560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729167648 := &system.Number_rule{Base: ptr8728818560, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728636224 := &system.Map_rule{Base: ptr8728818432, RuleBase: (*system.RuleBase)(nil), Items: ptr8729167648, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728490480 := &system.Type{Base: ptr8728818048, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729167456, "b": ptr8729167552, "c": ptr8728636224}}
	ptr8728821632 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728657920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651664 := &system.String_rule{Base: ptr8728657920, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728815744 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728652720 := &system.String_rule{Base: ptr8728815744, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728816000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728945376 := &system.RuleBase{Optional: true}
	ptr8728945248 := &system.Bool_rule{Base: ptr8728816000, RuleBase: ptr8728945376, Default: system.Bool{}}
	ptr8728492272 := &system.Type{Base: ptr8728821632, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728651664, "level": ptr8728652720, "preferred": ptr8728945248}}
	ptr8728818688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c", Rules: []system.Rule(nil)}
	ptr8728490592 := &system.Type{Base: ptr8728818688, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728821376 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728492048 := &system.Type{Base: ptr8728821376, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil)}
	ptr8729068416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil)}
	ptr8728492496 := &system.Type{Base: ptr8729068416, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728819200 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil)}
	ptr8728490816 := &system.Type{Base: ptr8728819200, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728818816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728818944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728819072 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728651840 := &system.String_rule{Base: ptr8728819072, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728636288 := &system.Map_rule{Base: ptr8728818944, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651840, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728490704 := &system.Type{Base: ptr8728818816, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728636288}}
	ptr8728821248 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8728491712 := &system.Type{Base: ptr8728821248, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8729070976 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil)}
	ptr8728493056 := &system.Type{Base: ptr8729070976, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8729068544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729068672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729165824 := &system.Number_rule{Base: ptr8729068672, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8729068800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729165920 := &system.Number_rule{Base: ptr8729068800, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8729068928 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729079280 := &selectors.C_rule{Base: ptr8729068928, RuleBase: (*system.RuleBase)(nil)}
	ptr8729069056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729069184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729166016 := &system.Number_rule{Base: ptr8729069184, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728634048 := &system.Map_rule{Base: ptr8729069056, RuleBase: (*system.RuleBase)(nil), Items: ptr8729166016, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729069312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729069440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729166112 := &system.Number_rule{Base: ptr8729069440, RuleBase: (*system.RuleBase)(nil), Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728634432 := &system.Map_rule{Base: ptr8729069312, RuleBase: (*system.RuleBase)(nil), Items: ptr8729166112, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728492608 := &system.Type{Base: ptr8729068544, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729165824, "b": ptr8729165920, "c": ptr8729079280, "d": ptr8728634048, "e": ptr8728634432}}
	ptr8728816128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil)}
	ptr8728486000 := &system.Type{Base: ptr8728816128, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8729069568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil)}
	ptr8728492832 := &system.Type{Base: ptr8729069568, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728492832)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728492384)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728491040)
	system.RegisterType("kego.io/selectors", "@image", ptr8728492160)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728492496)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728492608)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728491376)
	system.RegisterType("kego.io/selectors", "basic", ptr8728490144)
	system.RegisterType("kego.io/selectors", "c", ptr8728490480)
	system.RegisterType("kego.io/selectors", "@c", ptr8728490592)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728486000)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728490816)
	system.RegisterType("kego.io/selectors", "collision", ptr8728490704)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728493056)
	system.RegisterType("kego.io/selectors", "typed", ptr8728492944)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728490928)
	system.RegisterType("kego.io/selectors", "photo", ptr8728489808)
	system.RegisterType("kego.io/selectors", "expr", ptr8728491152)
	system.RegisterType("kego.io/selectors", "image", ptr8728492048)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728491600)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728490368)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728490032)
	system.RegisterType("kego.io/selectors", "kid", ptr8728492272)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728491712)
}
