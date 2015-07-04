package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8729224192 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729224320 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280704 := &system.Number_rule{Base: ptr8729224320, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729224448 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280800 := &system.Number_rule{Base: ptr8729224448, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729224576 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729238064 := &selectors.C_rule{Base: ptr8729224576, RuleBase: (*system.RuleBase)(nil)}
	ptr8729224704 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729224832 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280896 := &system.Number_rule{Base: ptr8729224832, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729215232 := &system.Map_rule{Base: ptr8729224704, RuleBase: (*system.RuleBase)(nil), Items: ptr8729280896, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729224960 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729225088 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729281376 := &system.Number_rule{Base: ptr8729225088, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729215296 := &system.Map_rule{Base: ptr8729224960, RuleBase: (*system.RuleBase)(nil), Items: ptr8729281376, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729029472 := &system.Type{Base: ptr8729224192, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729280704, "b": ptr8729280800, "c": ptr8729238064, "d": ptr8729215232, "e": ptr8729215296}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729223680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729223808 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729223936 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729237536 := &selectors.Kid_rule{Base: ptr8729223936, RuleBase: (*system.RuleBase)(nil)}
	ptr8729160992 := &system.Array_rule{Base: ptr8729223808, RuleBase: (*system.RuleBase)(nil), Items: ptr8729237536, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729029136 := &system.Type{Base: ptr8729223680, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729160992}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728956032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil)}
	ptr8729027008 := &system.Type{Base: ptr8728956032, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729223552 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8729029024 := &system.Type{Base: ptr8729223552, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729226624 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil)}
	ptr8729029808 := &system.Type{Base: ptr8729226624, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728959488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728959616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728959744 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728797184 := &system.String_rule{Base: ptr8728959744, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729216064 := &system.Map_rule{Base: ptr8728959616, RuleBase: (*system.RuleBase)(nil), Items: ptr8728797184, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728959872 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728797360 := &system.String_rule{Base: ptr8728959872, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728960000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728956416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728956544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728799824 := &system.String_rule{Base: ptr8728956544, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729216128 := &system.Map_rule{Base: ptr8728956416, RuleBase: (*system.RuleBase)(nil), Items: ptr8728799824, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729164912 := &system.Array_rule{Base: ptr8728960000, RuleBase: (*system.RuleBase)(nil), Items: ptr8729216128, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728956672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728956800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728800000 := &system.String_rule{Base: ptr8728956800, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729164992 := &system.Array_rule{Base: ptr8728956672, RuleBase: (*system.RuleBase)(nil), Items: ptr8728800000, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728956928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728957056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728798592 := &system.String_rule{Base: ptr8728957056, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729165072 := &system.Array_rule{Base: ptr8728956928, RuleBase: (*system.RuleBase)(nil), Items: ptr8728798592, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728957184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280992 := &system.Number_rule{Base: ptr8728957184, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729027792 := &system.Type{Base: ptr8728959488, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729216064, "favoriteColor": ptr8728797360, "languagesSpoken": ptr8729164912, "seatingPreference": ptr8729164992, "drinkPreference": ptr8729165072, "weight": ptr8729280992}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728960256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image", Rules: []system.Rule(nil)}
	ptr8729028576 := &system.Type{Base: ptr8728960256, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8729224064 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil)}
	ptr8729029360 := &system.Type{Base: ptr8729224064, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728958848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728958976 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728959104 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729234816 := &selectors.Image_rule{Base: ptr8728959104, RuleBase: (*system.RuleBase)(nil)}
	ptr8729215168 := &system.Map_rule{Base: ptr8728958976, RuleBase: (*system.RuleBase)(nil), Items: ptr8729234816, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729027568 := &system.Type{Base: ptr8728958848, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729215168}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728955264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728955392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728955520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728798944 := &system.String_rule{Base: ptr8728955520, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729215040 := &system.Map_rule{Base: ptr8728955392, RuleBase: (*system.RuleBase)(nil), Items: ptr8728798944, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729026672 := &system.Type{Base: ptr8728955264, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729215040}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729225344 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729225472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729225600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728801056 := &system.String_rule{Base: ptr8729225600, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729215360 := &system.Map_rule{Base: ptr8729225472, RuleBase: (*system.RuleBase)(nil), Items: ptr8728801056, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729225728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728801232 := &system.String_rule{Base: ptr8729225728, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729225856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729225984 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729239072 := &selectors.Kid_rule{Base: ptr8729225984, RuleBase: (*system.RuleBase)(nil)}
	ptr8729215424 := &system.Map_rule{Base: ptr8729225856, RuleBase: (*system.RuleBase)(nil), Items: ptr8729239072, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729226112 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729239168 := &selectors.Image_rule{Base: ptr8729226112, RuleBase: (*system.RuleBase)(nil)}
	ptr8729226240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729226368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728801408 := &system.String_rule{Base: ptr8729226368, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729161952 := &system.Array_rule{Base: ptr8729226240, RuleBase: (*system.RuleBase)(nil), Items: ptr8728801408, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729226496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729281472 := &system.Number_rule{Base: ptr8729226496, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729029696 := &system.Type{Base: ptr8729225344, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729215360, "favoriteColor": ptr8728801232, "kids": ptr8729215424, "avatar": ptr8729239168, "drinkPreference": ptr8729161952, "weight": ptr8729281472}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728955008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c", Rules: []system.Rule(nil)}
	ptr8729026560 := &system.Type{Base: ptr8728955008, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728958720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil)}
	ptr8729027344 := &system.Type{Base: ptr8728958720, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728955776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728955904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728799120 := &system.String_rule{Base: ptr8728955904, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729026896 := &system.Type{Base: ptr8728955776, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728799120}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}}
	ptr8728788992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8729223168 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8728690464 := &system.RuleBase{Optional: true}
	ptr8728800528 := &system.String_rule{Base: ptr8729223168, RuleBase: ptr8728690464, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729223296 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728800704 := &system.String_rule{Base: ptr8729223296, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729223424 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728800880 := &system.String_rule{Base: ptr8729223424, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8729028912 := &system.Type{Base: ptr8728788992, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728800528, "server": ptr8728800704, "path": ptr8728800880}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}}
	ptr8729225216 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil)}
	ptr8729029584 := &system.Type{Base: ptr8729225216, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728960128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil)}
	ptr8729028464 := &system.Type{Base: ptr8728960128, Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728957312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil)}
	ptr8729028016 := &system.Type{Base: ptr8728957312, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728957440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728957568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729281088 := &system.Number_rule{Base: ptr8728957568, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728957696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729281184 := &system.Number_rule{Base: ptr8728957696, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728957824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728957952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729281280 := &system.Number_rule{Base: ptr8728957952, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729216448 := &system.Map_rule{Base: ptr8728957824, RuleBase: (*system.RuleBase)(nil), Items: ptr8729281280, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729028128 := &system.Type{Base: ptr8728957440, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729281088, "b": ptr8729281184, "c": ptr8729216448}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728955648 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil)}
	ptr8729026784 := &system.Type{Base: ptr8728955648, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728956160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728959360 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280608 := &system.Number_rule{Base: ptr8728959360, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728958080 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728799296 := &system.String_rule{Base: ptr8728958080, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728799472 := &system.String_rule{Base: ptr8728958208, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728799648 := &system.String_rule{Base: ptr8728958336, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728958464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728686400 := &system.Bool_rule{Base: ptr8728958464, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728958592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728686656 := &system.Bool_rule{Base: ptr8728958592, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728956288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729280512 := &system.Number_rule{Base: ptr8728956288, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8729027232 := &system.Type{Base: ptr8728956160, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"float": ptr8729280608, "string": ptr8728799296, "string2": ptr8728799472, "null": ptr8728799648, "true": ptr8728686400, "false": ptr8728686656, "int": ptr8729280512}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728959232 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8729027680 := &system.Type{Base: ptr8728959232, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728960896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil)}
	ptr8729028800 := &system.Type{Base: ptr8728960896, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728960384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728960512 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728800176 := &system.String_rule{Base: ptr8728960512, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728960640 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728800352 := &system.String_rule{Base: ptr8728960640, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728960768 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728688640 := &system.RuleBase{Optional: true}
	ptr8728688544 := &system.Bool_rule{Base: ptr8728960768, RuleBase: ptr8728688640, Default: system.Bool{}}
	ptr8729028688 := &system.Type{Base: ptr8728960384, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728800176, "level": ptr8728800352, "preferred": ptr8728688544}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	system.RegisterType("kego.io/selectors", "collision", ptr8729026672)
	system.RegisterType("kego.io/selectors", "diagram", ptr8729026896)
	system.RegisterType("kego.io/selectors", "typed", ptr8729029696)
	system.RegisterType("kego.io/selectors", "image", ptr8729028464)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8729027680)
	system.RegisterType("kego.io/selectors", "polykids", ptr8729029136)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8729027008)
	system.RegisterType("kego.io/selectors", "@photo", ptr8729029024)
	system.RegisterType("kego.io/selectors", "@image", ptr8729028576)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8729029360)
	system.RegisterType("kego.io/selectors", "@kid", ptr8729028800)
	system.RegisterType("kego.io/selectors", "kid", ptr8729028688)
	system.RegisterType("kego.io/selectors", "expr", ptr8729027232)
	system.RegisterType("kego.io/selectors", "sibling", ptr8729029472)
	system.RegisterType("kego.io/selectors", "basic", ptr8729027792)
	system.RegisterType("kego.io/selectors", "photo", ptr8729028912)
	system.RegisterType("kego.io/selectors", "c", ptr8729028128)
	system.RegisterType("kego.io/selectors", "@collision", ptr8729026784)
	system.RegisterType("kego.io/selectors", "@basic", ptr8729028016)
	system.RegisterType("kego.io/selectors", "@typed", ptr8729029808)
	system.RegisterType("kego.io/selectors", "gallery", ptr8729027568)
	system.RegisterType("kego.io/selectors", "@c", ptr8729026560)
	system.RegisterType("kego.io/selectors", "@expr", ptr8729027344)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8729029584)
}
