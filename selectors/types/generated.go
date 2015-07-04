package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728871808 := &system.Base{Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}}
	ptr8728562080 := &system.Type{Base: ptr8728871808, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728505984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}}
	ptr8728506752 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728506880 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728507008 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728331648 := &system.String_rule{Base: ptr8728507008, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728601728 := &system.Map_rule{Base: ptr8728506880, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728331648, MinItems: system.Int{Value: 0}}
	ptr8728896528 := &system.Array_rule{Base: ptr8728506752, RuleBase: (*system.RuleBase)(nil), Items: ptr8728601728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728507136 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728509440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728331824 := &system.String_rule{Base: ptr8728509440, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728896608 := &system.Array_rule{Base: ptr8728507136, RuleBase: (*system.RuleBase)(nil), Items: ptr8728331824, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728509568 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728509696 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728333408 := &system.String_rule{Base: ptr8728509696, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728896688 := &system.Array_rule{Base: ptr8728509568, RuleBase: (*system.RuleBase)(nil), Items: ptr8728333408, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728504448 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674304 := &system.Number_rule{Base: ptr8728504448, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728506112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728506240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728330240 := &system.String_rule{Base: ptr8728506240, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728601664 := &system.Map_rule{Base: ptr8728506112, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728330240, MinItems: system.Int{Value: 0}}
	ptr8728506624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728330416 := &system.String_rule{Base: ptr8728506624, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728560736 := &system.Type{Base: ptr8728505984, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"languagesSpoken": ptr8728896528, "seatingPreference": ptr8728896608, "drinkPreference": ptr8728896688, "weight": ptr8728674304, "name": ptr8728601664, "favoriteColor": ptr8728330416}, Rule: (*system.Type)(nil)}
	ptr8728506496 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}
	ptr8728509824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332000 := &system.String_rule{Base: ptr8728509824, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728560176 := &system.Type{Base: ptr8728506496, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728332000}, Rule: (*system.Type)(nil)}
	ptr8728505472 := &system.Base{Description: "Automatically created basic rule for c", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}
	ptr8728559840 := &system.Type{Base: ptr8728505472, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728872960 := &system.Base{Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}}
	ptr8728562416 := &system.Type{Base: ptr8728872960, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728874368 := &system.Base{Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}}
	ptr8728562640 := &system.Type{Base: ptr8728874368, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728504704 := &system.Base{Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}}
	ptr8728559616 := &system.Type{Base: ptr8728504704, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728509056 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}
	ptr8728561296 := &system.Type{Base: ptr8728509056, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728507392 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}}
	ptr8728507648 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674784 := &system.Number_rule{Base: ptr8728507648, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728507776 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332176 := &system.String_rule{Base: ptr8728507776, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728507904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332352 := &system.String_rule{Base: ptr8728507904, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728508032 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332528 := &system.String_rule{Base: ptr8728508032, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728508160 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728653920 := &system.Bool_rule{Base: ptr8728508160, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728508288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728654176 := &system.Bool_rule{Base: ptr8728508288, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728507520 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674688 := &system.Number_rule{Base: ptr8728507520, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728560512 := &system.Type{Base: ptr8728507392, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"float": ptr8728674784, "string": ptr8728332176, "string2": ptr8728332352, "null": ptr8728332528, "true": ptr8728653920, "false": ptr8728654176, "int": ptr8728674688}, Rule: (*system.Type)(nil)}
	ptr8728871936 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}}
	ptr8728872064 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674880 := &system.Number_rule{Base: ptr8728872064, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728872192 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674976 := &system.Number_rule{Base: ptr8728872192, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728872320 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}}
	ptr8728853584 := &selectors.C_rule{Base: ptr8728872320, RuleBase: (*system.RuleBase)(nil)}
	ptr8728872448 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728872576 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728675072 := &system.Number_rule{Base: ptr8728872576, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728601024 := &system.Map_rule{Base: ptr8728872448, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728675072, MinItems: system.Int{Value: 0}}
	ptr8728872704 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728872832 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728675168 := &system.Number_rule{Base: ptr8728872832, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728601088 := &system.Map_rule{Base: ptr8728872704, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728675168, MinItems: system.Int{Value: 0}}
	ptr8728562192 := &system.Type{Base: ptr8728871936, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8728674880, "b": ptr8728674976, "c": ptr8728853584, "d": ptr8728601024, "e": ptr8728601088}, Rule: (*system.Type)(nil)}
	ptr8728509312 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}}
	ptr8728510208 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728656128 := &system.RuleBase{Optional: true}
	ptr8728656032 := &system.Bool_rule{Base: ptr8728510208, RuleBase: ptr8728656128, Default: system.Bool{}}
	ptr8728509952 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332704 := &system.String_rule{Base: ptr8728509952, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728510080 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728332880 := &system.String_rule{Base: ptr8728510080, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728561520 := &system.Type{Base: ptr8728509312, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"preferred": ptr8728656032, "language": ptr8728332704, "level": ptr8728332880}, Rule: (*system.Type)(nil)}
	ptr8728873088 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}}
	ptr8728873856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8728846944 := &selectors.Image_rule{Base: ptr8728873856, RuleBase: (*system.RuleBase)(nil)}
	ptr8728873984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728874112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728333584 := &system.String_rule{Base: ptr8728874112, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728895488 := &system.Array_rule{Base: ptr8728873984, RuleBase: (*system.RuleBase)(nil), Items: ptr8728333584, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728874240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728675264 := &system.Number_rule{Base: ptr8728874240, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728873216 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728873344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728334112 := &system.String_rule{Base: ptr8728873344, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728601152 := &system.Map_rule{Base: ptr8728873216, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728334112, MinItems: system.Int{Value: 0}}
	ptr8728873472 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728334288 := &system.String_rule{Base: ptr8728873472, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728873600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728873728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8728846624 := &selectors.Kid_rule{Base: ptr8728873728, RuleBase: (*system.RuleBase)(nil)}
	ptr8728600576 := &system.Map_rule{Base: ptr8728873600, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728846624, MinItems: system.Int{Value: 0}}
	ptr8728562528 := &system.Type{Base: ptr8728873088, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"avatar": ptr8728846944, "drinkPreference": ptr8728895488, "weight": ptr8728675264, "name": ptr8728601152, "favoriteColor": ptr8728334288, "kids": ptr8728600576}, Rule: (*system.Type)(nil)}
	ptr8728509184 := &system.Base{Description: "Automatically created basic rule for image", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}
	ptr8728561408 := &system.Type{Base: ptr8728509184, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728322048 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}
	ptr8728870912 := &system.Base{Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728656832 := &system.RuleBase{Optional: true}
	ptr8728333056 := &system.String_rule{Base: ptr8728870912, RuleBase: ptr8728656832, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728871040 := &system.Base{Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728333232 := &system.String_rule{Base: ptr8728871040, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728871168 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728333936 := &system.String_rule{Base: ptr8728871168, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728561744 := &system.Type{Base: ptr8728322048, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728333056, "server": ptr8728333232, "path": ptr8728333936}, Rule: (*system.Type)(nil)}
	ptr8728871424 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}}
	ptr8728871552 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728871680 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8728853056 := &selectors.Kid_rule{Base: ptr8728871680, RuleBase: (*system.RuleBase)(nil)}
	ptr8728899488 := &system.Array_rule{Base: ptr8728871552, RuleBase: (*system.RuleBase)(nil), Items: ptr8728853056, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728561968 := &system.Type{Base: ptr8728871424, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8728899488}, Rule: (*system.Type)(nil)}
	ptr8728506368 := &system.Base{Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}}
	ptr8728560064 := &system.Type{Base: ptr8728506368, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728510336 := &system.Base{Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}
	ptr8728561632 := &system.Type{Base: ptr8728510336, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728508544 := &system.Base{Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}
	ptr8728508672 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728508800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8728851040 := &selectors.Image_rule{Base: ptr8728508800, RuleBase: (*system.RuleBase)(nil)}
	ptr8728600960 := &system.Map_rule{Base: ptr8728508672, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728851040, MinItems: system.Int{Value: 0}}
	ptr8728560848 := &system.Type{Base: ptr8728508544, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728600960}, Rule: (*system.Type)(nil)}
	ptr8728871296 := &system.Base{Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}}
	ptr8728561856 := &system.Type{Base: ptr8728871296, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728508928 := &system.Base{Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}}
	ptr8728561184 := &system.Type{Base: ptr8728508928, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728507264 := &system.Base{Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}}
	ptr8728560400 := &system.Type{Base: ptr8728507264, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728504832 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}}
	ptr8728505216 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728505344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674592 := &system.Number_rule{Base: ptr8728505344, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728600768 := &system.Map_rule{Base: ptr8728505216, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728674592, MinItems: system.Int{Value: 0}}
	ptr8728504960 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674400 := &system.Number_rule{Base: ptr8728504960, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728505088 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728674496 := &system.Number_rule{Base: ptr8728505088, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728559728 := &system.Type{Base: ptr8728504832, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"c": ptr8728600768, "a": ptr8728674400, "b": ptr8728674496}, Rule: (*system.Type)(nil)}
	ptr8728505600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}}
	ptr8728505728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728505856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728333760 := &system.String_rule{Base: ptr8728505856, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728600832 := &system.Map_rule{Base: ptr8728505728, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728333760, MinItems: system.Int{Value: 0}}
	ptr8728559952 := &system.Type{Base: ptr8728505600, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728600832}, Rule: (*system.Type)(nil)}
	ptr8728508416 := &system.Base{Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}}
	ptr8728560624 := &system.Type{Base: ptr8728508416, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors", "@expr", ptr8728560624)
	system.RegisterType("kego.io/selectors", "image", ptr8728561296)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728561968)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728561632)
	system.RegisterType("kego.io/selectors", "c", ptr8728559728)
	system.RegisterType("kego.io/selectors", "collision", ptr8728559952)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728560176)
	system.RegisterType("kego.io/selectors", "expr", ptr8728560512)
	system.RegisterType("kego.io/selectors", "kid", ptr8728561520)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728561856)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728562080)
	system.RegisterType("kego.io/selectors", "basic", ptr8728560736)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728562416)
	system.RegisterType("kego.io/selectors", "typed", ptr8728562528)
	system.RegisterType("kego.io/selectors", "photo", ptr8728561744)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728560064)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728560848)
	system.RegisterType("kego.io/selectors", "@c", ptr8728559840)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728562640)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728559616)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728562192)
	system.RegisterType("kego.io/selectors", "@image", ptr8728561408)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728561184)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728560400)
}
