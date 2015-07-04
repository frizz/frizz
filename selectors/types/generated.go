package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728958848 := &system.Base{Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}}
	ptr8728958976 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728959104 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729203952 := &selectors.Image_rule{Base: ptr8728959104, RuleBase: (*system.RuleBase)(nil)}
	ptr8728783488 := &system.Map_rule{Base: ptr8728958976, RuleBase: (*system.RuleBase)(nil), Items: ptr8729203952, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728995024 := &system.Type{Base: ptr8728958848, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728783488}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728960896 := &system.Base{Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}
	ptr8728995696 := &system.Type{Base: ptr8728960896, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729305600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}}
	ptr8729305728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729305856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729200336 := &selectors.Kid_rule{Base: ptr8729305856, RuleBase: (*system.RuleBase)(nil)}
	ptr8729281232 := &system.Array_rule{Base: ptr8729305728, RuleBase: (*system.RuleBase)(nil), Items: ptr8729200336, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728996032 := &system.Type{Base: ptr8729305600, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729281232}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729306112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}}
	ptr8729306240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729363008 := &system.Number_rule{Base: ptr8729306240, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729306368 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729363104 := &system.Number_rule{Base: ptr8729306368, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729306496 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}}
	ptr8729201248 := &selectors.C_rule{Base: ptr8729306496, RuleBase: (*system.RuleBase)(nil)}
	ptr8729306624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729306752 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729363200 := &system.Number_rule{Base: ptr8729306752, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728781888 := &system.Map_rule{Base: ptr8729306624, RuleBase: (*system.RuleBase)(nil), Items: ptr8729363200, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729306880 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729307008 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729363296 := &system.Number_rule{Base: ptr8729307008, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728782080 := &system.Map_rule{Base: ptr8729306880, RuleBase: (*system.RuleBase)(nil), Items: ptr8729363296, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728996368 := &system.Type{Base: ptr8729306112, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729363008, "b": ptr8729363104, "c": ptr8729201248, "d": ptr8728781888, "e": ptr8728782080}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728957184 := &system.Base{Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}}
	ptr8728994352 := &system.Type{Base: ptr8728957184, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729307136 := &system.Base{Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}}
	ptr8728996480 := &system.Type{Base: ptr8729307136, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728959744 := &system.Base{Description: "Automatically created basic rule for c", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}
	ptr8728994128 := &system.Type{Base: ptr8728959744, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729305984 := &system.Base{Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}}
	ptr8728996144 := &system.Type{Base: ptr8729305984, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728957696 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}}
	ptr8728958592 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728934496 := &system.Bool_rule{Base: ptr8728958592, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728957824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362816 := &system.Number_rule{Base: ptr8728957824, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728957952 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362912 := &system.Number_rule{Base: ptr8728957952, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728958080 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791104 := &system.String_rule{Base: ptr8728958080, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958208 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791280 := &system.String_rule{Base: ptr8728958208, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958336 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791456 := &system.String_rule{Base: ptr8728958336, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958464 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728934240 := &system.Bool_rule{Base: ptr8728958464, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728994800 := &system.Type{Base: ptr8728957696, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"false": ptr8728934496, "int": ptr8729362816, "float": ptr8729362912, "string": ptr8728791104, "string2": ptr8728791280, "null": ptr8728791456, "true": ptr8728934240}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728959232 := &system.Base{Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}}
	ptr8728995136 := &system.Type{Base: ptr8728959232, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728956800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}}
	ptr8728956928 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362528 := &system.Number_rule{Base: ptr8728956928, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728959360 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362624 := &system.Number_rule{Base: ptr8728959360, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728959488 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728959616 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362720 := &system.Number_rule{Base: ptr8728959616, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728782784 := &system.Map_rule{Base: ptr8728959488, RuleBase: (*system.RuleBase)(nil), Items: ptr8729362720, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728994016 := &system.Type{Base: ptr8728956800, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729362528, "b": ptr8729362624, "c": ptr8728782784}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728797184 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}}
	ptr8729305088 := &system.Base{Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728685024 := &system.RuleBase{Optional: true}
	ptr8728792336 := &system.String_rule{Base: ptr8729305088, RuleBase: ptr8728685024, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729305216 := &system.Base{Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728792512 := &system.String_rule{Base: ptr8729305216, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729305344 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728790576 := &system.String_rule{Base: ptr8729305344, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728995808 := &system.Type{Base: ptr8728797184, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728792336, "server": ptr8728792512, "path": ptr8728790576}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}}
	ptr8728959872 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}}
	ptr8728960000 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728957056 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728790752 := &system.String_rule{Base: ptr8728957056, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728783360 := &system.Map_rule{Base: ptr8728960000, RuleBase: (*system.RuleBase)(nil), Items: ptr8728790752, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728994240 := &system.Type{Base: ptr8728959872, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728783360}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728957568 := &system.Base{Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}}
	ptr8728994688 := &system.Type{Base: ptr8728957568, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728960384 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}}
	ptr8728960512 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791984 := &system.String_rule{Base: ptr8728960512, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728960640 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728792160 := &system.String_rule{Base: ptr8728960640, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728960768 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728682624 := &system.RuleBase{Optional: true}
	ptr8728936384 := &system.Bool_rule{Base: ptr8728960768, RuleBase: ptr8728682624, Default: system.Bool{}}
	ptr8728995584 := &system.Type{Base: ptr8728960384, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728791984, "level": ptr8728792160, "preferred": ptr8728936384}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728957312 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}}
	ptr8728957440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728790928 := &system.String_rule{Base: ptr8728957440, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728994464 := &system.Type{Base: ptr8728957312, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728790928}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}}
	ptr8729307264 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}}
	ptr8729307776 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729307904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}}
	ptr8729202640 := &selectors.Kid_rule{Base: ptr8729307904, RuleBase: (*system.RuleBase)(nil)}
	ptr8728782464 := &system.Map_rule{Base: ptr8729307776, RuleBase: (*system.RuleBase)(nil), Items: ptr8729202640, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729308032 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}}
	ptr8729202752 := &selectors.Image_rule{Base: ptr8729308032, RuleBase: (*system.RuleBase)(nil)}
	ptr8729308160 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729308288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728793216 := &system.String_rule{Base: ptr8729308288, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729282112 := &system.Array_rule{Base: ptr8729308160, RuleBase: (*system.RuleBase)(nil), Items: ptr8728793216, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729308416 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729363392 := &system.Number_rule{Base: ptr8729308416, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8729307392 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8729307520 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728792864 := &system.String_rule{Base: ptr8729307520, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728782144 := &system.Map_rule{Base: ptr8729307392, RuleBase: (*system.RuleBase)(nil), Items: ptr8728792864, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729307648 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728793040 := &system.String_rule{Base: ptr8729307648, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728996592 := &system.Type{Base: ptr8729307264, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"kids": ptr8728782464, "avatar": ptr8729202752, "drinkPreference": ptr8729282112, "weight": ptr8729363392, "name": ptr8728782144, "favoriteColor": ptr8728793040}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728960256 := &system.Base{Description: "Automatically created basic rule for image", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}
	ptr8728995360 := &system.Type{Base: ptr8728960256, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728955008 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}}
	ptr8728956288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728956416 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728790400 := &system.String_rule{Base: ptr8728956416, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729281152 := &system.Array_rule{Base: ptr8728956288, RuleBase: (*system.RuleBase)(nil), Items: ptr8728790400, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728956544 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729362432 := &system.Number_rule{Base: ptr8728956544, RuleBase: (*system.RuleBase)(nil), Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728955264 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728955392 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728788992 := &system.String_rule{Base: ptr8728955392, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728780928 := &system.Map_rule{Base: ptr8728955264, RuleBase: (*system.RuleBase)(nil), Items: ptr8728788992, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728955520 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728789168 := &system.String_rule{Base: ptr8728955520, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728955648 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728955776 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728955904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791632 := &system.String_rule{Base: ptr8728955904, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728781504 := &system.Map_rule{Base: ptr8728955776, RuleBase: (*system.RuleBase)(nil), Items: ptr8728791632, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729280992 := &system.Array_rule{Base: ptr8728955648, RuleBase: (*system.RuleBase)(nil), Items: ptr8728781504, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728956032 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728956160 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728791808 := &system.String_rule{Base: ptr8728956160, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729281072 := &system.Array_rule{Base: ptr8728956032, RuleBase: (*system.RuleBase)(nil), Items: ptr8728791808, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728993792 := &system.Type{Base: ptr8728955008, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"drinkPreference": ptr8729281152, "weight": ptr8729362432, "name": ptr8728780928, "favoriteColor": ptr8728789168, "languagesSpoken": ptr8729280992, "seatingPreference": ptr8729281072}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729308544 := &system.Base{Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}}
	ptr8728996704 := &system.Type{Base: ptr8729308544, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728958720 := &system.Base{Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}}
	ptr8728994912 := &system.Type{Base: ptr8728958720, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729305472 := &system.Base{Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}}
	ptr8728995920 := &system.Type{Base: ptr8729305472, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728956672 := &system.Base{Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}}
	ptr8728993904 := &system.Type{Base: ptr8728956672, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728960128 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}
	ptr8728995248 := &system.Type{Base: ptr8728960128, Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728996144)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728995136)
	system.RegisterType("kego.io/selectors", "@image", ptr8728995360)
	system.RegisterType("kego.io/selectors", "basic", ptr8728993792)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728995696)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728996480)
	system.RegisterType("kego.io/selectors", "@c", ptr8728994128)
	system.RegisterType("kego.io/selectors", "collision", ptr8728994240)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728994464)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728994912)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728995024)
	system.RegisterType("kego.io/selectors", "expr", ptr8728994800)
	system.RegisterType("kego.io/selectors", "photo", ptr8728995808)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728996704)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728995920)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728993904)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728996368)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728994352)
	system.RegisterType("kego.io/selectors", "c", ptr8728994016)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728994688)
	system.RegisterType("kego.io/selectors", "kid", ptr8728995584)
	system.RegisterType("kego.io/selectors", "typed", ptr8728996592)
	system.RegisterType("kego.io/selectors", "image", ptr8728995248)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728996032)
}
