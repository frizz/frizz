package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8729175040 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729175808 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729175936 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729092384 := &system.Number_rule{Base: ptr8729175936, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961344 := &system.Map_rule{Base: ptr8729175808, RuleBase: (*system.RuleBase)(nil), Items: ptr8729092384, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729175168 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729092096 := &system.Number_rule{Base: ptr8729175168, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729175296 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729092192 := &system.Number_rule{Base: ptr8729175296, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729175424 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729147776 := &selectors.C_rule{Base: ptr8729175424, RuleBase: (*system.RuleBase)(nil)}
	ptr8729175552 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729175680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729092288 := &system.Number_rule{Base: ptr8729175680, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961216 := &system.Map_rule{Base: ptr8729175552, RuleBase: (*system.RuleBase)(nil), Items: ptr8729092288, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728890432 := &system.Type{Base: ptr8729175040, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"e": ptr8728961344, "a": ptr8729092096, "b": ptr8729092192, "c": ptr8729147776, "d": ptr8728961216}, Rule: (*system.Type)(nil)}
	ptr8728838016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil)}
	ptr8728888080 := &system.Type{Base: ptr8728838016, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728835968 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728836096 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098816 := &system.Number_rule{Base: ptr8728836096, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728836224 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098912 := &system.Number_rule{Base: ptr8728836224, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728836352 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660384 := &system.String_rule{Base: ptr8728836352, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728836480 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660560 := &system.String_rule{Base: ptr8728836480, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728836608 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660736 := &system.String_rule{Base: ptr8728836608, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728836736 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729152352 := &system.Bool_rule{Base: ptr8728836736, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728836864 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729152640 := &system.Bool_rule{Base: ptr8728836864, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728889200 := &system.Type{Base: ptr8728835968, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"int": ptr8729098816, "float": ptr8729098912, "string": ptr8728660384, "string2": ptr8728660560, "null": ptr8728660736, "true": ptr8729152352, "false": ptr8729152640}, Rule: (*system.Type)(nil)}
	ptr8728836992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8728887632 := &system.Type{Base: ptr8728836992, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728837120 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728889424 := &system.Type{Base: ptr8728837120, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729176064 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil)}
	ptr8728890544 := &system.Type{Base: ptr8729176064, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729177472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil)}
	ptr8728890768 := &system.Type{Base: ptr8729177472, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729174528 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729174656 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729174784 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729146944 := &selectors.Kid_rule{Base: ptr8729174784, RuleBase: (*system.RuleBase)(nil)}
	ptr8728954992 := &system.Array_rule{Base: ptr8729174656, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8729146944, MinItems: system.Int{Value: 0}}
	ptr8728890208 := &system.Type{Base: ptr8729174528, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8728954992}, Rule: (*system.Type)(nil)}
	ptr8728837888 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil)}
	ptr8728889760 := &system.Type{Base: ptr8728837888, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729174912 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil)}
	ptr8728890320 := &system.Type{Base: ptr8729174912, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728835456 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil)}
	ptr8728888528 := &system.Type{Base: ptr8728835456, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8729176192 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729176576 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661968 := &system.String_rule{Base: ptr8729176576, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729176704 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729176832 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729149312 := &selectors.Kid_rule{Base: ptr8729176832, RuleBase: (*system.RuleBase)(nil)}
	ptr8728961472 := &system.Map_rule{Base: ptr8729176704, RuleBase: (*system.RuleBase)(nil), Items: ptr8729149312, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729176960 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729149408 := &selectors.Image_rule{Base: ptr8729176960, RuleBase: (*system.RuleBase)(nil)}
	ptr8729177088 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729177216 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728662144 := &system.String_rule{Base: ptr8729177216, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728955872 := &system.Array_rule{Base: ptr8729177088, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728662144, MinItems: system.Int{Value: 0}}
	ptr8729177344 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729092480 := &system.Number_rule{Base: ptr8729177344, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729176320 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729176448 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661792 := &system.String_rule{Base: ptr8729176448, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961408 := &system.Map_rule{Base: ptr8729176320, RuleBase: (*system.RuleBase)(nil), Items: ptr8728661792, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728890656 := &system.Type{Base: ptr8729176192, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"favoriteColor": ptr8728661968, "kids": ptr8728961472, "avatar": ptr8729149408, "drinkPreference": ptr8728955872, "weight": ptr8729092480, "name": ptr8728961408}, Rule: (*system.Type)(nil)}
	ptr8728832256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728832512 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728832640 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729142880 := &selectors.Image_rule{Base: ptr8728832640, RuleBase: (*system.RuleBase)(nil)}
	ptr8728961152 := &system.Map_rule{Base: ptr8728832512, RuleBase: (*system.RuleBase)(nil), Items: ptr8729142880, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728887520 := &system.Type{Base: ptr8728832256, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728961152}, Rule: (*system.Type)(nil)}
	ptr8728835072 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728835200 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728835328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660032 := &system.String_rule{Base: ptr8728835328, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961984 := &system.Map_rule{Base: ptr8728835200, RuleBase: (*system.RuleBase)(nil), Items: ptr8728660032, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728888416 := &system.Type{Base: ptr8728835072, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728961984}, Rule: (*system.Type)(nil)}
	ptr8728832768 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728833664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728833792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659504 := &system.String_rule{Base: ptr8728833792, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958912 := &system.Array_rule{Base: ptr8728833664, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728659504, MinItems: system.Int{Value: 0}}
	ptr8728833920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728834048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659680 := &system.String_rule{Base: ptr8728834048, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958992 := &system.Array_rule{Base: ptr8728833920, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728659680, MinItems: system.Int{Value: 0}}
	ptr8728834176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098432 := &system.Number_rule{Base: ptr8728834176, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728832896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728833024 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728658976 := &system.String_rule{Base: ptr8728833024, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961536 := &system.Map_rule{Base: ptr8728832896, RuleBase: (*system.RuleBase)(nil), Items: ptr8728658976, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728833152 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659152 := &system.String_rule{Base: ptr8728833152, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728833280 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728833408 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728833536 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659328 := &system.String_rule{Base: ptr8728833536, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961600 := &system.Map_rule{Base: ptr8728833408, RuleBase: (*system.RuleBase)(nil), Items: ptr8728659328, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728958832 := &system.Array_rule{Base: ptr8728833280, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728961600, MinItems: system.Int{Value: 0}}
	ptr8728887744 := &system.Type{Base: ptr8728832768, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"seatingPreference": ptr8728958912, "drinkPreference": ptr8728958992, "weight": ptr8729098432, "name": ptr8728961536, "favoriteColor": ptr8728659152, "languagesSpoken": ptr8728958832}, Rule: (*system.Type)(nil)}
	ptr8729174400 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8728890096 := &system.Type{Base: ptr8729174400, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728835584 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728835712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660208 := &system.String_rule{Base: ptr8728835712, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728888640 := &system.Type{Base: ptr8728835584, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728660208}, Rule: (*system.Type)(nil)}
	ptr8728666112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8729174016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8729152736 := &system.RuleBase{Optional: true}
	ptr8728661264 := &system.String_rule{Base: ptr8729174016, RuleBase: ptr8729152736, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729174144 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728661440 := &system.String_rule{Base: ptr8729174144, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729174272 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728661616 := &system.String_rule{Base: ptr8729174272, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728889872 := &system.Type{Base: ptr8728666112, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728661264, "server": ptr8728661440, "path": ptr8728661616}, Rule: (*system.Type)(nil)}
	ptr8728837376 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728837504 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660912 := &system.String_rule{Base: ptr8728837504, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728837632 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661088 := &system.String_rule{Base: ptr8728837632, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728837760 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729151680 := &system.RuleBase{Optional: true}
	ptr8729151584 := &system.Bool_rule{Base: ptr8728837760, RuleBase: ptr8729151680, Default: system.Bool{}}
	ptr8728889648 := &system.Type{Base: ptr8728837376, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728660912, "level": ptr8728661088, "preferred": ptr8729151584}, Rule: (*system.Type)(nil)}
	ptr8728834304 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728834432 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098528 := &system.Number_rule{Base: ptr8728834432, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728834560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098624 := &system.Number_rule{Base: ptr8728834560, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728834688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728834816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729098720 := &system.Number_rule{Base: ptr8728834816, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961920 := &system.Map_rule{Base: ptr8728834688, RuleBase: (*system.RuleBase)(nil), Items: ptr8729098720, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728888192 := &system.Type{Base: ptr8728834304, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729098528, "b": ptr8729098624, "c": ptr8728961920}, Rule: (*system.Type)(nil)}
	ptr8728834944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c", Rules: []system.Rule(nil)}
	ptr8728888304 := &system.Type{Base: ptr8728834944, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728837248 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image", Rules: []system.Rule(nil)}
	ptr8728889536 := &system.Type{Base: ptr8728837248, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728835840 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil)}
	ptr8728888976 := &system.Type{Base: ptr8728835840, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728832128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil)}
	ptr8728887296 := &system.Type{Base: ptr8728832128, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728890320)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728887520)
	system.RegisterType("kego.io/selectors", "collision", ptr8728888416)
	system.RegisterType("kego.io/selectors", "basic", ptr8728887744)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728890096)
	system.RegisterType("kego.io/selectors", "photo", ptr8728889872)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728890432)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728890208)
	system.RegisterType("kego.io/selectors", "@image", ptr8728889536)
	system.RegisterType("kego.io/selectors", "kid", ptr8728889648)
	system.RegisterType("kego.io/selectors", "c", ptr8728888192)
	system.RegisterType("kego.io/selectors", "typed", ptr8728890656)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728888976)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728888080)
	system.RegisterType("kego.io/selectors", "image", ptr8728889424)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728888640)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728887296)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728887632)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728889760)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728890768)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728888528)
	system.RegisterType("kego.io/selectors", "@c", ptr8728888304)
	system.RegisterType("kego.io/selectors", "expr", ptr8728889200)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728890544)
}
