package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729606880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499392 := &system.Base{Context: ptr8729606880, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729617616 := &system.Type{Base: ptr8729499392, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729804832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499904 := &system.Base{Context: ptr8729804832, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729132960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500800 := &system.Base{Context: ptr8729132960, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729132832 := &system.Bool_rule{Base: ptr8729500800, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729130656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500032 := &system.Base{Context: ptr8729130656, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231840 := &system.Number_rule{Base: ptr8729500032, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729131104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500160 := &system.Base{Context: ptr8729131104, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231936 := &system.Number_rule{Base: ptr8729500160, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729131552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500288 := &system.Base{Context: ptr8729131552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729618496 := &system.String_rule{Base: ptr8729500288, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729131872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500416 := &system.Base{Context: ptr8729131872, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729618672 := &system.String_rule{Base: ptr8729500416, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729132192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500544 := &system.Base{Context: ptr8729132192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729618848 := &system.String_rule{Base: ptr8729500544, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729132608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500672 := &system.Base{Context: ptr8729132608, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729132416 := &system.Bool_rule{Base: ptr8729500672, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729618320 := &system.Type{Base: ptr8729499904, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"false": ptr8729132832, "int": ptr8729231840, "float": ptr8729231936, "string": ptr8729618496, "string2": ptr8729618672, "null": ptr8729618848, "true": ptr8729132416}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729812512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839232 := &system.Base{Context: ptr8729812512, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729811712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839744 := &system.Base{Context: ptr8729811712, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729811552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839872 := &system.Base{Context: ptr8729811552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729232224 := &system.Number_rule{Base: ptr8729839872, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729650944 := &system.Map_rule{Base: ptr8729839744, RuleBase: (*system.RuleBase)(nil), Items: ptr8729232224, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729812384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840000 := &system.Base{Context: ptr8729812384, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729812224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840128 := &system.Base{Context: ptr8729812224, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729232320 := &system.Number_rule{Base: ptr8729840128, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729651008 := &system.Map_rule{Base: ptr8729840000, RuleBase: (*system.RuleBase)(nil), Items: ptr8729232320, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729810144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839360 := &system.Base{Context: ptr8729810144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729232032 := &system.Number_rule{Base: ptr8729839360, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729810592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839488 := &system.Base{Context: ptr8729810592, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729232128 := &system.Number_rule{Base: ptr8729839488, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729811040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839616 := &system.Base{Context: ptr8729811040, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@c", Package: "kego.io/selectors", Type: "@c", Exists: true}}

	ptr8729848496 := &selectors.C_rule{Base: ptr8729839616, RuleBase: (*system.RuleBase)(nil)}

	ptr8729622016 := &system.Type{Base: ptr8729839232, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"d": ptr8729650944, "e": ptr8729651008, "a": ptr8729232032, "b": ptr8729232128, "c": ptr8729848496}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729602592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497728 := &system.Base{Context: ptr8729602592, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729600960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729496704 := &system.Base{Context: ptr8729600960, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729600928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729496832 := &system.Base{Context: ptr8729600928, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729616384 := &system.String_rule{Base: ptr8729496832, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729608192 := &system.Array_rule{Base: ptr8729496704, RuleBase: (*system.RuleBase)(nil), Items: ptr8729616384, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729601600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729496960 := &system.Base{Context: ptr8729601600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729601568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497088 := &system.Base{Context: ptr8729601568, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729616560 := &system.String_rule{Base: ptr8729497088, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729608272 := &system.Array_rule{Base: ptr8729496960, RuleBase: (*system.RuleBase)(nil), Items: ptr8729616560, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729602336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497216 := &system.Base{Context: ptr8729602336, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231456 := &system.Number_rule{Base: ptr8729497216, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729606976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497856 := &system.Base{Context: ptr8729606976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729606944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498112 := &system.Base{Context: ptr8729606944, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729361984 := &system.String_rule{Base: ptr8729498112, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729650112 := &system.Map_rule{Base: ptr8729497856, RuleBase: (*system.RuleBase)(nil), Items: ptr8729361984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729607296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498240 := &system.Base{Context: ptr8729607296, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729362160 := &system.String_rule{Base: ptr8729498240, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729600352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498368 := &system.Base{Context: ptr8729600352, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729600320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498496 := &system.Base{Context: ptr8729600320, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729600288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729496576 := &system.Base{Context: ptr8729600288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729361808 := &system.String_rule{Base: ptr8729496576, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729650176 := &system.Map_rule{Base: ptr8729498496, RuleBase: (*system.RuleBase)(nil), Items: ptr8729361808, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729611552 := &system.Array_rule{Base: ptr8729498368, RuleBase: (*system.RuleBase)(nil), Items: ptr8729650176, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729361632 := &system.Type{Base: ptr8729497728, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"seatingPreference": ptr8729608192, "drinkPreference": ptr8729608272, "weight": ptr8729231456, "name": ptr8729650112, "favoriteColor": ptr8729362160, "languagesSpoken": ptr8729611552}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729606432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499008 := &system.Base{Context: ptr8729606432, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729606272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499136 := &system.Base{Context: ptr8729606272, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729606240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499264 := &system.Base{Context: ptr8729606240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729617440 := &system.String_rule{Base: ptr8729499264, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729649664 := &system.Map_rule{Base: ptr8729499136, RuleBase: (*system.RuleBase)(nil), Items: ptr8729617440, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729617264 := &system.Type{Base: ptr8729499008, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729649664}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729605344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498880 := &system.Base{Context: ptr8729605344, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729617088 := &system.Type{Base: ptr8729498880, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729806784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838080 := &system.Base{Context: ptr8729806784, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729805440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838208 := &system.Base{Context: ptr8729805440, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729805408 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729620960 := &system.String_rule{Base: ptr8729838208, RuleBase: ptr8729805408, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729805856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838336 := &system.Base{Context: ptr8729805856, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729621136 := &system.String_rule{Base: ptr8729838336, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729806464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838464 := &system.Base{Context: ptr8729806464, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729621312 := &system.String_rule{Base: ptr8729838464, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729620608 := &system.Type{Base: ptr8729838080, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8729620960, "server": ptr8729621136, "path": ptr8729621312}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729607904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499520 := &system.Base{Context: ptr8729607904, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729607712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499648 := &system.Base{Context: ptr8729607712, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729617968 := &system.String_rule{Base: ptr8729499648, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729617792 := &system.Type{Base: ptr8729499520, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729617968}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729808672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729371904 := &system.Base{Context: ptr8729808672, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729807840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729837568 := &system.Base{Context: ptr8729807840, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729620080 := &system.String_rule{Base: ptr8729837568, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729808160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729837696 := &system.Base{Context: ptr8729808160, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729620256 := &system.String_rule{Base: ptr8729837696, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729808544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729837824 := &system.Base{Context: ptr8729808544, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729808512 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729808384 := &system.Bool_rule{Base: ptr8729837824, RuleBase: ptr8729808512, Default: system.Bool{Value: false, Exists: false}}

	ptr8729619904 := &system.Type{Base: ptr8729371904, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729620080, "level": ptr8729620256, "preferred": ptr8729808384}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729812928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840256 := &system.Base{Context: ptr8729812928, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729622192 := &system.Type{Base: ptr8729840256, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729603936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840384 := &system.Base{Context: ptr8729603936, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729600992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840512 := &system.Base{Context: ptr8729600992, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729600864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840640 := &system.Base{Context: ptr8729600864, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729622544 := &system.String_rule{Base: ptr8729840640, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729651200 := &system.Map_rule{Base: ptr8729840512, RuleBase: (*system.RuleBase)(nil), Items: ptr8729622544, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729601376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840768 := &system.Base{Context: ptr8729601376, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729622720 := &system.String_rule{Base: ptr8729840768, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729601984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729840896 := &system.Base{Context: ptr8729601984, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729601952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841024 := &system.Base{Context: ptr8729601952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729850656 := &selectors.Kid_rule{Base: ptr8729841024, RuleBase: (*system.RuleBase)(nil)}

	ptr8729651264 := &system.Map_rule{Base: ptr8729840896, RuleBase: (*system.RuleBase)(nil), Items: ptr8729850656, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729602368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841152 := &system.Base{Context: ptr8729602368, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8729850816 := &selectors.Image_rule{Base: ptr8729841152, RuleBase: (*system.RuleBase)(nil)}

	ptr8729603200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841280 := &system.Base{Context: ptr8729603200, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729603168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841408 := &system.Base{Context: ptr8729603168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729622896 := &system.String_rule{Base: ptr8729841408, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729609712 := &system.Array_rule{Base: ptr8729841280, RuleBase: (*system.RuleBase)(nil), Items: ptr8729622896, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729603680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841536 := &system.Base{Context: ptr8729603680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729232416 := &system.Number_rule{Base: ptr8729841536, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729622368 := &system.Type{Base: ptr8729840384, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729651200, "favoriteColor": ptr8729622720, "kids": ptr8729651264, "avatar": ptr8729850816, "drinkPreference": ptr8729609712, "weight": ptr8729232416}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729808800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838720 := &system.Base{Context: ptr8729808800, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729808576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838848 := &system.Base{Context: ptr8729808576, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729808448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838976 := &system.Base{Context: ptr8729808448, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729847200 := &selectors.Kid_rule{Base: ptr8729838976, RuleBase: (*system.RuleBase)(nil)}

	ptr8729608752 := &system.Array_rule{Base: ptr8729838848, RuleBase: (*system.RuleBase)(nil), Items: ptr8729847200, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729621664 := &system.Type{Base: ptr8729838720, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729608752}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729809600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729839104 := &system.Base{Context: ptr8729809600, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729621840 := &system.Type{Base: ptr8729839104, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729806880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729501568 := &system.Base{Context: ptr8729806880, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729619552 := &system.Type{Base: ptr8729501568, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729603008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497344 := &system.Base{Context: ptr8729603008, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729616736 := &system.Type{Base: ptr8729497344, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729604928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497472 := &system.Base{Context: ptr8729604928, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729603552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497600 := &system.Base{Context: ptr8729603552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231552 := &system.Number_rule{Base: ptr8729497600, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729604000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729497984 := &system.Base{Context: ptr8729604000, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231648 := &system.Number_rule{Base: ptr8729497984, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729604800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498624 := &system.Base{Context: ptr8729604800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729604640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729498752 := &system.Base{Context: ptr8729604640, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729231744 := &system.Number_rule{Base: ptr8729498752, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729649472 := &system.Map_rule{Base: ptr8729498624, RuleBase: (*system.RuleBase)(nil), Items: ptr8729231744, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729616912 := &system.Type{Base: ptr8729497472, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729231552, "b": ptr8729231648, "c": ptr8729649472}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729806592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729501440 := &system.Base{Context: ptr8729806592, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729619376 := &system.Type{Base: ptr8729501440, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729130080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729499776 := &system.Base{Context: ptr8729130080, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729618144 := &system.Type{Base: ptr8729499776, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729807520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729838592 := &system.Base{Context: ptr8729807520, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729621488 := &system.Type{Base: ptr8729838592, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729809088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729837952 := &system.Base{Context: ptr8729809088, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729620432 := &system.Type{Base: ptr8729837952, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729805248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729500928 := &system.Base{Context: ptr8729805248, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729619024 := &system.Type{Base: ptr8729500928, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729604448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729841664 := &system.Base{Context: ptr8729604448, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729623072 := &system.Type{Base: ptr8729841664, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729807296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729370624 := &system.Base{Context: ptr8729807296, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729619728 := &system.Type{Base: ptr8729370624, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729806176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729501056 := &system.Base{Context: ptr8729806176, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729806048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729501184 := &system.Base{Context: ptr8729806048, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729806016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729501312 := &system.Base{Context: ptr8729806016, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8729044608 := &selectors.Image_rule{Base: ptr8729501312, RuleBase: (*system.RuleBase)(nil)}

	ptr8729650368 := &system.Map_rule{Base: ptr8729501184, RuleBase: (*system.RuleBase)(nil), Items: ptr8729044608, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729619200 := &system.Type{Base: ptr8729501056, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729650368}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors:@basic", ptr8729616736)

	system.RegisterType("kego.io/selectors:@c", ptr8729617088)

	system.RegisterType("kego.io/selectors:@collision", ptr8729617616)

	system.RegisterType("kego.io/selectors:@diagram", ptr8729618144)

	system.RegisterType("kego.io/selectors:@expr", ptr8729619024)

	system.RegisterType("kego.io/selectors:@gallery", ptr8729619376)

	system.RegisterType("kego.io/selectors:@image", ptr8729619728)

	system.RegisterType("kego.io/selectors:@kid", ptr8729620432)

	system.RegisterType("kego.io/selectors:@photo", ptr8729621488)

	system.RegisterType("kego.io/selectors:@polykids", ptr8729621840)

	system.RegisterType("kego.io/selectors:@sibling", ptr8729622192)

	system.RegisterType("kego.io/selectors:@typed", ptr8729623072)

	system.RegisterType("kego.io/selectors:basic", ptr8729361632)

	system.RegisterType("kego.io/selectors:c", ptr8729616912)

	system.RegisterType("kego.io/selectors:collision", ptr8729617264)

	system.RegisterType("kego.io/selectors:diagram", ptr8729617792)

	system.RegisterType("kego.io/selectors:expr", ptr8729618320)

	system.RegisterType("kego.io/selectors:gallery", ptr8729619200)

	system.RegisterType("kego.io/selectors:image", ptr8729619552)

	system.RegisterType("kego.io/selectors:kid", ptr8729619904)

	system.RegisterType("kego.io/selectors:photo", ptr8729620608)

	system.RegisterType("kego.io/selectors:polykids", ptr8729621664)

	system.RegisterType("kego.io/selectors:sibling", ptr8729622016)

	system.RegisterType("kego.io/selectors:typed", ptr8729622368)

}
