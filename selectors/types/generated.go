package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728828928 := &system.Base{Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}}
	ptr8728828800 := &system.Type{Base: ptr8728828928, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729142272 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}}
	ptr8729142400 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728666112 := &system.Number_rule{Base: ptr8729142400, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729145984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728666784 := &system.Number_rule{Base: ptr8729145984, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729146112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651664 := &system.String_rule{Base: ptr8729146112, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729146240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651840 := &system.String_rule{Base: ptr8729146240, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729146368 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652016 := &system.String_rule{Base: ptr8729146368, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729146496 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728555840 := &system.Bool_rule{Base: ptr8729146496, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8729146624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728556128 := &system.Bool_rule{Base: ptr8729146624, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8729142144 := &system.Type{Base: ptr8729142272, Fields: map[string]system.Rule{"int": ptr8728666112, "float": ptr8728666784, "string": ptr8728651664, "string2": ptr8728651840, "null": ptr8728652016, "true": ptr8728555840, "false": ptr8728556128}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729145344 := &system.Base{Description: "Automatically created basic rule for c", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}
	ptr8729145216 := &system.Type{Base: ptr8729145344, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728829184 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}}
	ptr8728829312 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728829440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729154816 := &selectors.Kid_rule{Base: ptr8728829440, RuleBase: (*system.RuleBase)(nil)}
	ptr8728963984 := &system.Array_rule{Base: ptr8728829312, RuleBase: (*system.RuleBase)(nil), Items: ptr8729154816, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728829056 := &system.Type{Base: ptr8728829184, Fields: map[string]system.Rule{"a": ptr8728963984}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729149312 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}
	ptr8728828672 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653424 := &system.String_rule{Base: ptr8728828672, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}}
	ptr8728828416 := &system.Base{Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728558752 := &system.RuleBase{Optional: true}
	ptr8728653072 := &system.String_rule{Base: ptr8728828416, RuleBase: ptr8728558752, Pattern: system.String{}, Format: system.String{}, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728828544 := &system.Base{Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653248 := &system.String_rule{Base: ptr8728828544, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729149184 := &system.Type{Base: ptr8729149312, Fields: map[string]system.Rule{"path": ptr8728653424, "protocol": ptr8728653072, "server": ptr8728653248}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729141632 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}
	ptr8729141760 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651488 := &system.String_rule{Base: ptr8729141760, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729141504 := &system.Type{Base: ptr8729141632, Fields: map[string]system.Rule{"url": ptr8728651488}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729142016 := &system.Base{Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}}
	ptr8729141888 := &system.Type{Base: ptr8729142016, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729147136 := &system.Base{Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}
	ptr8729147264 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729147392 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729152128 := &selectors.Image_rule{Base: ptr8729147392, RuleBase: (*system.RuleBase)(nil)}
	ptr8728977536 := &system.Map_rule{Base: ptr8729147264, RuleBase: (*system.RuleBase)(nil), Items: ptr8729152128, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729147008 := &system.Type{Base: ptr8729147136, Fields: map[string]system.Rule{"images": ptr8728977536}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729142656 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}}
	ptr8729143168 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729143296 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729143424 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652544 := &system.String_rule{Base: ptr8729143424, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728978304 := &system.Map_rule{Base: ptr8729143296, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652544, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728964944 := &system.Array_rule{Base: ptr8729143168, RuleBase: (*system.RuleBase)(nil), Items: ptr8728978304, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729143552 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729143680 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652720 := &system.String_rule{Base: ptr8729143680, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728965024 := &system.Array_rule{Base: ptr8729143552, RuleBase: (*system.RuleBase)(nil), Items: ptr8728652720, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729143808 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729143936 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649904 := &system.String_rule{Base: ptr8729143936, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728965104 := &system.Array_rule{Base: ptr8729143808, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649904, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729144064 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728670816 := &system.Number_rule{Base: ptr8729144064, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729142784 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729142912 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728649728 := &system.String_rule{Base: ptr8729142912, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728978240 := &system.Map_rule{Base: ptr8729142784, RuleBase: (*system.RuleBase)(nil), Items: ptr8728649728, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729143040 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652368 := &system.String_rule{Base: ptr8729143040, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729142528 := &system.Type{Base: ptr8729142656, Fields: map[string]system.Rule{"languagesSpoken": ptr8728964944, "seatingPreference": ptr8728965024, "drinkPreference": ptr8728965104, "weight": ptr8728670816, "name": ptr8728978240, "favoriteColor": ptr8728652368}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729019264 := &system.Base{Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}}
	ptr8729019136 := &system.Type{Base: ptr8729019264, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729149056 := &system.Base{Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}
	ptr8729148928 := &system.Type{Base: ptr8729149056, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729144576 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}}
	ptr8729144704 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728670912 := &system.Number_rule{Base: ptr8729144704, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729144832 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728671008 := &system.Number_rule{Base: ptr8729144832, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729144960 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729145088 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728671104 := &system.Number_rule{Base: ptr8729145088, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728978624 := &system.Map_rule{Base: ptr8729144960, RuleBase: (*system.RuleBase)(nil), Items: ptr8728671104, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729144448 := &system.Type{Base: ptr8729144576, Fields: map[string]system.Rule{"a": ptr8728670912, "b": ptr8728671008, "c": ptr8728978624}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729146880 := &system.Base{Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}}
	ptr8729146752 := &system.Type{Base: ptr8729146880, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729020928 := &system.Base{Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}}
	ptr8729020800 := &system.Type{Base: ptr8729020928, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728657920 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}}
	ptr8729018368 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728666976 := &system.Number_rule{Base: ptr8729018368, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729018496 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}}
	ptr8729155488 := &selectors.C_rule{Base: ptr8729018496, RuleBase: (*system.RuleBase)(nil)}
	ptr8729018624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729018752 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728667072 := &system.Number_rule{Base: ptr8729018752, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728977600 := &system.Map_rule{Base: ptr8729018624, RuleBase: (*system.RuleBase)(nil), Items: ptr8728667072, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729018880 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729019008 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728667168 := &system.Number_rule{Base: ptr8729019008, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728977664 := &system.Map_rule{Base: ptr8729018880, RuleBase: (*system.RuleBase)(nil), Items: ptr8728667168, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728658048 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728666880 := &system.Number_rule{Base: ptr8728658048, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728829824 := &system.Type{Base: ptr8728657920, Fields: map[string]system.Rule{"b": ptr8728666976, "c": ptr8729155488, "d": ptr8728977600, "e": ptr8728977664, "a": ptr8728666880}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729147648 := &system.Base{Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}}
	ptr8729147520 := &system.Type{Base: ptr8729147648, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729147904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}
	ptr8729147776 := &system.Type{Base: ptr8729147904, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true}
	ptr8729148160 := &system.Base{Description: "Automatically created basic rule for image", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}
	ptr8729148032 := &system.Type{Base: ptr8729148160, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729019520 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}}
	ptr8729020032 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729020160 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729156496 := &selectors.Kid_rule{Base: ptr8729020160, RuleBase: (*system.RuleBase)(nil)}
	ptr8728977792 := &system.Map_rule{Base: ptr8729020032, RuleBase: (*system.RuleBase)(nil), Items: ptr8729156496, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729020288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729156592 := &selectors.Image_rule{Base: ptr8729020288, RuleBase: (*system.RuleBase)(nil)}
	ptr8729020416 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729020544 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653952 := &system.String_rule{Base: ptr8729020544, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728965264 := &system.Array_rule{Base: ptr8729020416, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729020672 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728667264 := &system.Number_rule{Base: ptr8729020672, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729019648 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729019776 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653600 := &system.String_rule{Base: ptr8729019776, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728977728 := &system.Map_rule{Base: ptr8729019648, RuleBase: (*system.RuleBase)(nil), Items: ptr8728653600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729019904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728653776 := &system.String_rule{Base: ptr8729019904, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729019392 := &system.Type{Base: ptr8729019520, Fields: map[string]system.Rule{"kids": ptr8728977792, "avatar": ptr8729156592, "drinkPreference": ptr8728965264, "weight": ptr8728667264, "name": ptr8728977728, "favoriteColor": ptr8728653776}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728829696 := &system.Base{Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}}
	ptr8728829568 := &system.Type{Base: ptr8728829696, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729148416 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}}
	ptr8729148544 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652192 := &system.String_rule{Base: ptr8729148544, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729148672 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728652896 := &system.String_rule{Base: ptr8729148672, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8729148800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728558048 := &system.RuleBase{Optional: true}
	ptr8728557952 := &system.Bool_rule{Base: ptr8729148800, RuleBase: ptr8728558048, Default: system.Bool{}}
	ptr8729148288 := &system.Type{Base: ptr8729148416, Fields: map[string]system.Rule{"language": ptr8728652192, "level": ptr8728652896, "preferred": ptr8728557952}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729141376 := &system.Base{Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}}
	ptr8729141248 := &system.Type{Base: ptr8729141376, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729144320 := &system.Base{Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}}
	ptr8729144192 := &system.Type{Base: ptr8729144320, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729145600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}}
	ptr8729145728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729145856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728651312 := &system.String_rule{Base: ptr8729145856, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728978688 := &system.Map_rule{Base: ptr8729145728, RuleBase: (*system.RuleBase)(nil), Items: ptr8728651312, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729145472 := &system.Type{Base: ptr8729145600, Fields: map[string]system.Rule{"number": ptr8728978688}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	system.RegisterType("kego.io/selectors", "@c", ptr8729145216)
	system.RegisterType("kego.io/selectors", "basic", ptr8729142528)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8729019136)
	system.RegisterType("kego.io/selectors", "c", ptr8729144448)
	system.RegisterType("kego.io/selectors", "@typed", ptr8729020800)
	system.RegisterType("kego.io/selectors", "@image", ptr8729148032)
	system.RegisterType("kego.io/selectors", "expr", ptr8729142144)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8729147520)
	system.RegisterType("kego.io/selectors", "image", ptr8729147776)
	system.RegisterType("kego.io/selectors", "typed", ptr8729019392)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728829568)
	system.RegisterType("kego.io/selectors", "kid", ptr8729148288)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728828800)
	system.RegisterType("kego.io/selectors", "photo", ptr8729149184)
	system.RegisterType("kego.io/selectors", "gallery", ptr8729147008)
	system.RegisterType("kego.io/selectors", "@kid", ptr8729148928)
	system.RegisterType("kego.io/selectors", "@expr", ptr8729146752)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728829824)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728829056)
	system.RegisterType("kego.io/selectors", "diagram", ptr8729141504)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8729141888)
	system.RegisterType("kego.io/selectors", "@collision", ptr8729141248)
	system.RegisterType("kego.io/selectors", "@basic", ptr8729144192)
	system.RegisterType("kego.io/selectors", "collision", ptr8729145472)
}
