package types

import (
	selectors "kego.io/selectors"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728861952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Description: "Automatically created basic rule for image", Rules: []system.Rule(nil)}
	ptr8728928704 := &system.Type{Base: ptr8728861952, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728861056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728860544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660208 := &system.String_rule{Base: ptr8728860544, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728930384 := &system.Type{Base: ptr8728861056, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728660208}, Rule: (*system.Type)(nil)}
	ptr8728970240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728970368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729182208 := &system.Number_rule{Base: ptr8728970368, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728970496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729182304 := &system.Number_rule{Base: ptr8728970496, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728970624 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728851248 := &selectors.C_rule{Base: ptr8728970624, RuleBase: (*system.RuleBase)(nil)}
	ptr8728970752 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728970880 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729182400 := &system.Number_rule{Base: ptr8728970880, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961216 := &system.Map_rule{Base: ptr8728970752, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8729182400}
	ptr8728971008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728971136 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729182496 := &system.Number_rule{Base: ptr8728971136, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961344 := &system.Map_rule{Base: ptr8728971008, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8729182496}
	ptr8728931056 := &system.Type{Base: ptr8728970240, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729182208, "b": ptr8729182304, "c": ptr8728851248, "d": ptr8728961216, "e": ptr8728961344}, Rule: (*system.Type)(nil)}
	ptr8728856960 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728857088 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728857216 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728846864 := &selectors.Image_rule{Base: ptr8728857216, RuleBase: (*system.RuleBase)(nil)}
	ptr8728961152 := &system.Map_rule{Base: ptr8728857088, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728846864}
	ptr8728928368 := &system.Type{Base: ptr8728856960, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728961152}, Rule: (*system.Type)(nil)}
	ptr8728857600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728857728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728857856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728657920 := &system.String_rule{Base: ptr8728857856, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961536 := &system.Map_rule{Base: ptr8728857728, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728657920}
	ptr8728857984 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659152 := &system.String_rule{Base: ptr8728857984, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728858112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728858240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728858368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659328 := &system.String_rule{Base: ptr8728858368, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961600 := &system.Map_rule{Base: ptr8728858240, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728659328}
	ptr8728985088 := &system.Array_rule{Base: ptr8728858112, RuleBase: (*system.RuleBase)(nil), Items: ptr8728961600, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728858496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728858624 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659504 := &system.String_rule{Base: ptr8728858624, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728985168 := &system.Array_rule{Base: ptr8728858496, RuleBase: (*system.RuleBase)(nil), Items: ptr8728659504, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728858752 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728858880 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659680 := &system.String_rule{Base: ptr8728858880, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728985248 := &system.Array_rule{Base: ptr8728858752, RuleBase: (*system.RuleBase)(nil), Items: ptr8728659680, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728859008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189024 := &system.Number_rule{Base: ptr8728859008, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728929376 := &system.Type{Base: ptr8728857600, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8728961536, "favoriteColor": ptr8728659152, "languagesSpoken": ptr8728985088, "seatingPreference": ptr8728985168, "drinkPreference": ptr8728985248, "weight": ptr8729189024}, Rule: (*system.Type)(nil)}
	ptr8728857472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728928592 := &system.Type{Base: ptr8728857472, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728857344 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8728928480 := &system.Type{Base: ptr8728857344, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728859904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Description: "Automatically created basic rule for c", Rules: []system.Rule(nil)}
	ptr8728929824 := &system.Type{Base: ptr8728859904, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728859264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728859392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189120 := &system.Number_rule{Base: ptr8728859392, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728859520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189216 := &system.Number_rule{Base: ptr8728859520, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728859648 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728859776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189312 := &system.Number_rule{Base: ptr8728859776, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728961920 := &system.Map_rule{Base: ptr8728859648, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8729189312}
	ptr8728929712 := &system.Type{Base: ptr8728859264, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729189120, "b": ptr8729189216, "c": ptr8728961920}, Rule: (*system.Type)(nil)}
	ptr8728860416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Description: "Automatically created basic rule for collision", Rules: []system.Rule(nil)}
	ptr8728930272 := &system.Type{Base: ptr8728860416, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728862592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Description: "Automatically created basic rule for kid", Rules: []system.Rule(nil)}
	ptr8728928928 := &system.Type{Base: ptr8728862592, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728856704 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Description: "Automatically created basic rule for expr", Rules: []system.Rule(nil)}
	ptr8728928256 := &system.Type{Base: ptr8728856704, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728969600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8728929264 := &system.Type{Base: ptr8728969600, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728666112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8728969216 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8729110080 := &system.RuleBase{Optional: true}
	ptr8728661264 := &system.String_rule{Base: ptr8728969216, RuleBase: ptr8729110080, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728969344 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728661440 := &system.String_rule{Base: ptr8728969344, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728969472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728661616 := &system.String_rule{Base: ptr8728969472, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728929040 := &system.Type{Base: ptr8728666112, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8728661264, "server": ptr8728661440, "path": ptr8728661616}, Rule: (*system.Type)(nil)}
	ptr8728971264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Description: "Automatically created basic rule for sibling", Rules: []system.Rule(nil)}
	ptr8728931168 := &system.Type{Base: ptr8728971264, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728860800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728861184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189504 := &system.Number_rule{Base: ptr8728861184, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728861312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660384 := &system.String_rule{Base: ptr8728861312, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728861440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660560 := &system.String_rule{Base: ptr8728861440, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728861568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660736 := &system.String_rule{Base: ptr8728861568, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728861696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729109472 := &system.Bool_rule{Base: ptr8728861696, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728861824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729109728 := &system.Bool_rule{Base: ptr8728861824, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr8728860928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729189408 := &system.Number_rule{Base: ptr8728860928, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728930608 := &system.Type{Base: ptr8728860800, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"float": ptr8729189504, "string": ptr8728660384, "string2": ptr8728660560, "null": ptr8728660736, "true": ptr8729109472, "false": ptr8729109728, "int": ptr8729189408}, Rule: (*system.Type)(nil)}
	ptr8728971392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728971904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728972032 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728852352 := &selectors.Kid_rule{Base: ptr8728972032, RuleBase: (*system.RuleBase)(nil)}
	ptr8728961472 := &system.Map_rule{Base: ptr8728971904, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728852352}
	ptr8728972160 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728852448 := &selectors.Image_rule{Base: ptr8728972160, RuleBase: (*system.RuleBase)(nil)}
	ptr8728972288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728972416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728662144 := &system.String_rule{Base: ptr8728972416, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728988880 := &system.Array_rule{Base: ptr8728972288, RuleBase: (*system.RuleBase)(nil), Items: ptr8728662144, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728972544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729182592 := &system.Number_rule{Base: ptr8728972544, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728971520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728971648 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661792 := &system.String_rule{Base: ptr8728971648, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961408 := &system.Map_rule{Base: ptr8728971520, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728661792}
	ptr8728971776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661968 := &system.String_rule{Base: ptr8728971776, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728931280 := &system.Type{Base: ptr8728971392, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"kids": ptr8728961472, "avatar": ptr8728852448, "drinkPreference": ptr8728988880, "weight": ptr8729182592, "name": ptr8728961408, "favoriteColor": ptr8728661968}, Rule: (*system.Type)(nil)}
	ptr8728969728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728969856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728969984 := &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728850304 := &selectors.Kid_rule{Base: ptr8728969984, RuleBase: (*system.RuleBase)(nil)}
	ptr8728988000 := &system.Array_rule{Base: ptr8728969856, RuleBase: (*system.RuleBase)(nil), Items: ptr8728850304, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728930720 := &system.Type{Base: ptr8728969728, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8728988000}, Rule: (*system.Type)(nil)}
	ptr8728862080 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728862208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660912 := &system.String_rule{Base: ptr8728862208, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728862336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728661088 := &system.String_rule{Base: ptr8728862336, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728862464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729116288 := &system.RuleBase{Optional: true}
	ptr8729116192 := &system.Bool_rule{Base: ptr8728862464, RuleBase: ptr8729116288, Default: system.Bool{}}
	ptr8728928816 := &system.Type{Base: ptr8728862080, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8728660912, "level": ptr8728661088, "preferred": ptr8729116192}, Rule: (*system.Type)(nil)}
	ptr8728972672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Description: "Automatically created basic rule for typed", Rules: []system.Rule(nil)}
	ptr8728931392 := &system.Type{Base: ptr8728972672, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728860032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728860160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728860288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660032 := &system.String_rule{Base: ptr8728860288, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728961984 := &system.Map_rule{Base: ptr8728860160, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728660032}
	ptr8728930160 := &system.Type{Base: ptr8728860032, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728961984}, Rule: (*system.Type)(nil)}
	ptr8728859136 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Description: "Automatically created basic rule for basic", Rules: []system.Rule(nil)}
	ptr8728929488 := &system.Type{Base: ptr8728859136, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728970112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Description: "Automatically created basic rule for polykids", Rules: []system.Rule(nil)}
	ptr8728930832 := &system.Type{Base: ptr8728970112, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728860672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Description: "Automatically created basic rule for diagram", Rules: []system.Rule(nil)}
	ptr8728930496 := &system.Type{Base: ptr8728860672, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors", "image", ptr8728928592)
	system.RegisterType("kego.io/selectors", "@gallery", ptr8728928480)
	system.RegisterType("kego.io/selectors", "@kid", ptr8728928928)
	system.RegisterType("kego.io/selectors", "polykids", ptr8728930720)
	system.RegisterType("kego.io/selectors", "@typed", ptr8728931392)
	system.RegisterType("kego.io/selectors", "basic", ptr8728929376)
	system.RegisterType("kego.io/selectors", "@collision", ptr8728930272)
	system.RegisterType("kego.io/selectors", "@expr", ptr8728928256)
	system.RegisterType("kego.io/selectors", "@photo", ptr8728929264)
	system.RegisterType("kego.io/selectors", "@sibling", ptr8728931168)
	system.RegisterType("kego.io/selectors", "typed", ptr8728931280)
	system.RegisterType("kego.io/selectors", "kid", ptr8728928816)
	system.RegisterType("kego.io/selectors", "@basic", ptr8728929488)
	system.RegisterType("kego.io/selectors", "@image", ptr8728928704)
	system.RegisterType("kego.io/selectors", "diagram", ptr8728930384)
	system.RegisterType("kego.io/selectors", "c", ptr8728929712)
	system.RegisterType("kego.io/selectors", "photo", ptr8728929040)
	system.RegisterType("kego.io/selectors", "collision", ptr8728930160)
	system.RegisterType("kego.io/selectors", "sibling", ptr8728931056)
	system.RegisterType("kego.io/selectors", "gallery", ptr8728928368)
	system.RegisterType("kego.io/selectors", "@c", ptr8728929824)
	system.RegisterType("kego.io/selectors", "expr", ptr8728930608)
	system.RegisterType("kego.io/selectors", "@polykids", ptr8728930832)
	system.RegisterType("kego.io/selectors", "@diagram", ptr8728930496)
}
