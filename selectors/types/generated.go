package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729460224 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729460352 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729135696 := &system.String_rule{Base: ptr8729460352, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729460480 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729136400 := &system.String_rule{Base: ptr8729460480, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729460608 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728731872 := &system.RuleBase{Optional: true}

	ptr8728736800 := &system.Bool_rule{Base: ptr8729460608, RuleBase: ptr8728731872, Default: system.Bool{}}

	ptr8729460096 := &system.Type{Base: ptr8729460224, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr8729135696, "level": ptr8729136400, "preferred": ptr8728736800}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459968 := &system.Base{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459840 := &system.Type{Base: ptr8729459968, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459712 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459584 := &system.Type{Base: ptr8729459712, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729454848 := &system.Base{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729454720 := &system.Type{Base: ptr8729454848, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729457536 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729458176 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729135520 := &system.String_rule{Base: ptr8729458176, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729458304 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728734720 := &system.Bool_rule{Base: ptr8729458304, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729458432 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728734976 := &system.Bool_rule{Base: ptr8729458432, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729457664 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479616 := &system.Number_rule{Base: ptr8729457664, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729457792 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479712 := &system.Number_rule{Base: ptr8729457792, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729457920 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729135168 := &system.String_rule{Base: ptr8729457920, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729458048 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729135344 := &system.String_rule{Base: ptr8729458048, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729457408 := &system.Type{Base: ptr8729457536, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"null": ptr8729135520, "true": ptr8728734720, "false": ptr8728734976, "int": ptr8729479616, "float": ptr8729479712, "string": ptr8729135168, "string2": ptr8729135344}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729453184 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729453312 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729453440 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729135872 := &system.String_rule{Base: ptr8729453440, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728930688 := &system.Map_rule{Base: ptr8729453312, RuleBase: (*system.RuleBase)(nil), Items: ptr8729135872, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729453568 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729136048 := &system.String_rule{Base: ptr8729453568, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729453696 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729453824 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729453952 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729136224 := &system.String_rule{Base: ptr8729453952, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728930752 := &system.Map_rule{Base: ptr8729453824, RuleBase: (*system.RuleBase)(nil), Items: ptr8729136224, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728924144 := &system.Array_rule{Base: ptr8729453696, RuleBase: (*system.RuleBase)(nil), Items: ptr8728930752, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729454080 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729454208 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729134288 := &system.String_rule{Base: ptr8729454208, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728924224 := &system.Array_rule{Base: ptr8729454080, RuleBase: (*system.RuleBase)(nil), Items: ptr8729134288, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729454336 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729454464 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729134464 := &system.String_rule{Base: ptr8729454464, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728924304 := &system.Array_rule{Base: ptr8729454336, RuleBase: (*system.RuleBase)(nil), Items: ptr8729134464, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729454592 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479232 := &system.Number_rule{Base: ptr8729454592, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729453056 := &system.Type{Base: ptr8729453184, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr8728930688, "favoriteColor": ptr8729136048, "languagesSpoken": ptr8728924144, "seatingPreference": ptr8728924224, "drinkPreference": ptr8728924304, "weight": ptr8729479232}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729458688 := &system.Base{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729458560 := &system.Type{Base: ptr8729458688, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729156480 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729156608 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728734688 := &system.RuleBase{Optional: true}

	ptr8729136576 := &system.String_rule{Base: ptr8729156608, RuleBase: ptr8728734688, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729156736 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729136752 := &system.String_rule{Base: ptr8729156736, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729156864 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729136928 := &system.String_rule{Base: ptr8729156864, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729156352 := &system.Type{Base: ptr8729156480, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"protocol": ptr8729136576, "server": ptr8729136752, "path": ptr8729136928}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729156224 := &system.Base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729156096 := &system.Type{Base: ptr8729156224, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456640 := &system.Base{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729456512 := &system.Type{Base: ptr8729456640, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729457280 := &system.Base{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729457152 := &system.Type{Base: ptr8729457280, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729452800 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729452928 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729477120 := &system.Number_rule{Base: ptr8729452928, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729272320 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729477216 := &system.Number_rule{Base: ptr8729272320, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729272448 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}

	ptr8729299856 := &selectors.C_rule{Base: ptr8729272448, RuleBase: (*system.RuleBase)(nil)}

	ptr8729272576 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729272704 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729477312 := &system.Number_rule{Base: ptr8729272704, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728929792 := &system.Map_rule{Base: ptr8729272576, RuleBase: (*system.RuleBase)(nil), Items: ptr8729477312, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729272832 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729272960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729477408 := &system.Number_rule{Base: ptr8729272960, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728929856 := &system.Map_rule{Base: ptr8729272832, RuleBase: (*system.RuleBase)(nil), Items: ptr8729477408, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729452672 := &system.Type{Base: ptr8729452800, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr8729477120, "b": ptr8729477216, "c": ptr8729299856, "d": ptr8728929792, "e": ptr8728929856}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729274880 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729274752 := &system.Type{Base: ptr8729274880, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729157120 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729156992 := &system.Type{Base: ptr8729157120, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729452544 := &system.Base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729034880 := &system.Type{Base: ptr8729452544, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729273216 := &system.Base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729273088 := &system.Type{Base: ptr8729273216, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729459456 := &system.Base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459328 := &system.Type{Base: ptr8729459456, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456896 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729457024 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729134992 := &system.String_rule{Base: ptr8729457024, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729456768 := &system.Type{Base: ptr8729456896, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr8729134992}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729157376 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729157504 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729034752 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729299152 := &selectors.Kid_rule{Base: ptr8729034752, RuleBase: (*system.RuleBase)(nil)}

	ptr8728921264 := &system.Array_rule{Base: ptr8729157504, RuleBase: (*system.RuleBase)(nil), Items: ptr8729299152, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729157248 := &system.Type{Base: ptr8729157376, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr8728921264}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729458944 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729459072 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729459200 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729302416 := &selectors.Image_rule{Base: ptr8729459200, RuleBase: (*system.RuleBase)(nil)}

	ptr8728931328 := &system.Map_rule{Base: ptr8729459072, RuleBase: (*system.RuleBase)(nil), Items: ptr8729302416, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729458816 := &system.Type{Base: ptr8729458944, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr8728931328}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729273472 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729273600 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729273728 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729137104 := &system.String_rule{Base: ptr8729273728, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728930368 := &system.Map_rule{Base: ptr8729273600, RuleBase: (*system.RuleBase)(nil), Items: ptr8729137104, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729273856 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729137280 := &system.String_rule{Base: ptr8729273856, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729273984 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729274112 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729301296 := &selectors.Kid_rule{Base: ptr8729274112, RuleBase: (*system.RuleBase)(nil)}

	ptr8728930496 := &system.Map_rule{Base: ptr8729273984, RuleBase: (*system.RuleBase)(nil), Items: ptr8729301296, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729274240 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729301616 := &selectors.Image_rule{Base: ptr8729274240, RuleBase: (*system.RuleBase)(nil)}

	ptr8729274368 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729274496 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729137456 := &system.String_rule{Base: ptr8729274496, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728923584 := &system.Array_rule{Base: ptr8729274368, RuleBase: (*system.RuleBase)(nil), Items: ptr8729137456, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729274624 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729477504 := &system.Number_rule{Base: ptr8729274624, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729273344 := &system.Type{Base: ptr8729273472, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr8728930368, "favoriteColor": ptr8729137280, "kids": ptr8728930496, "avatar": ptr8729301616, "drinkPreference": ptr8728923584, "weight": ptr8729477504}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729456128 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729456256 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729456384 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729134816 := &system.String_rule{Base: ptr8729456384, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728931136 := &system.Map_rule{Base: ptr8729456256, RuleBase: (*system.RuleBase)(nil), Items: ptr8729134816, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729456000 := &system.Type{Base: ptr8729456128, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr8728931136}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729455872 := &system.Base{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729455744 := &system.Type{Base: ptr8729455872, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729455104 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729455232 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479328 := &system.Number_rule{Base: ptr8729455232, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729455360 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479424 := &system.Number_rule{Base: ptr8729455360, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729455488 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729455616 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729479520 := &system.Number_rule{Base: ptr8729455616, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728931072 := &system.Map_rule{Base: ptr8729455488, RuleBase: (*system.RuleBase)(nil), Items: ptr8729479520, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729454976 := &system.Type{Base: ptr8729455104, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr8729479328, "b": ptr8729479424, "c": ptr8728931072}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors", "@expr", ptr8729458560)

	system.RegisterType("kego.io/selectors", "@kid", ptr8729156096)

	system.RegisterType("kego.io/selectors", "diagram", ptr8729456768)

	system.RegisterType("kego.io/selectors", "expr", ptr8729457408)

	system.RegisterType("kego.io/selectors", "basic", ptr8729453056)

	system.RegisterType("kego.io/selectors", "@gallery", ptr8729459328)

	system.RegisterType("kego.io/selectors", "polykids", ptr8729157248)

	system.RegisterType("kego.io/selectors", "collision", ptr8729456000)

	system.RegisterType("kego.io/selectors", "photo", ptr8729156352)

	system.RegisterType("kego.io/selectors", "@polykids", ptr8729034880)

	system.RegisterType("kego.io/selectors", "sibling", ptr8729452672)

	system.RegisterType("kego.io/selectors", "@typed", ptr8729274752)

	system.RegisterType("kego.io/selectors", "@photo", ptr8729156992)

	system.RegisterType("kego.io/selectors", "@sibling", ptr8729273088)

	system.RegisterType("kego.io/selectors", "@c", ptr8729455744)

	system.RegisterType("kego.io/selectors", "c", ptr8729454976)

	system.RegisterType("kego.io/selectors", "image", ptr8729459584)

	system.RegisterType("kego.io/selectors", "@collision", ptr8729456512)

	system.RegisterType("kego.io/selectors", "@basic", ptr8729454720)

	system.RegisterType("kego.io/selectors", "@diagram", ptr8729457152)

	system.RegisterType("kego.io/selectors", "gallery", ptr8729458816)

	system.RegisterType("kego.io/selectors", "typed", ptr8729273344)

	system.RegisterType("kego.io/selectors", "kid", ptr8729460096)

	system.RegisterType("kego.io/selectors", "@image", ptr8729459840)

}

var _ selectors.Nothing
