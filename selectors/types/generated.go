package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729190848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8728876736 := &system.Base{Context: ptr8729190848, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729296544 := &system.Type{Base: ptr8728876736, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728851552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245936 := &system.Base{Context: ptr8728851552, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728851360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246048 := &system.Base{Context: ptr8728851360, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8728851328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246160 := &system.Base{Context: ptr8728851328, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "@kid", Exists: true}}

	ptr8729458688 := &selectors.Kid_rule{Base: ptr8729246160, RuleBase: (*system.RuleBase)(nil)}

	ptr8729209424 := &system.Array_rule{Base: ptr8729246048, RuleBase: (*system.RuleBase)(nil), Items: ptr8729458688, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729296064 := &system.Type{Base: ptr8729245936, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729209424}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729640416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245824 := &system.Base{Context: ptr8729640416, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729295904 := &system.Type{Base: ptr8729245824, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729197664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241456 := &system.Base{Context: ptr8729197664, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729195296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241792 := &system.Base{Context: ptr8729195296, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282096 := &system.String_rule{Base: ptr8729241792, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729195936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242016 := &system.Base{Context: ptr8729195936, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8729195904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242240 := &system.Base{Context: ptr8729195904, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729195872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242464 := &system.Base{Context: ptr8729195872, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282272 := &system.String_rule{Base: ptr8729242464, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729502336 := &system.Map_rule{Base: ptr8729242240, RuleBase: (*system.RuleBase)(nil), Items: ptr8729282272, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729207504 := &system.Array_rule{Base: ptr8729242016, RuleBase: (*system.RuleBase)(nil), Items: ptr8729502336, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729196480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242576 := &system.Base{Context: ptr8729196480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8729196448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242688 := &system.Base{Context: ptr8729196448, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282448 := &system.String_rule{Base: ptr8729242688, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729207584 := &system.Array_rule{Base: ptr8729242576, RuleBase: (*system.RuleBase)(nil), Items: ptr8729282448, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729197056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242800 := &system.Base{Context: ptr8729197056, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8729197024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729242912 := &system.Base{Context: ptr8729197024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282624 := &system.String_rule{Base: ptr8729242912, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729207664 := &system.Array_rule{Base: ptr8729242800, RuleBase: (*system.RuleBase)(nil), Items: ptr8729282624, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729197408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243024 := &system.Base{Context: ptr8729197408, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729307872 := &system.Number_rule{Base: ptr8729243024, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729195008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241568 := &system.Base{Context: ptr8729195008, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729194976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241680 := &system.Base{Context: ptr8729194976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729281920 := &system.String_rule{Base: ptr8729241680, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729502272 := &system.Map_rule{Base: ptr8729241568, RuleBase: (*system.RuleBase)(nil), Items: ptr8729281920, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729295264 := &system.Type{Base: ptr8729241456, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"favoriteColor": ptr8729282096, "languagesSpoken": ptr8729207504, "seatingPreference": ptr8729207584, "drinkPreference": ptr8729207664, "weight": ptr8729307872, "name": ptr8729502272}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729638624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245264 := &system.Base{Context: ptr8729638624, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729295424 := &system.Type{Base: ptr8729245264, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729634976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243920 := &system.Base{Context: ptr8729634976, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729634848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244032 := &system.Base{Context: ptr8729634848, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729634816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244144 := &system.Base{Context: ptr8729634816, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282800 := &system.String_rule{Base: ptr8729244144, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729502720 := &system.Map_rule{Base: ptr8729244032, RuleBase: (*system.RuleBase)(nil), Items: ptr8729282800, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729293344 := &system.Type{Base: ptr8729243920, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729502720}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729636000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244480 := &system.Base{Context: ptr8729636000, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729635872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244592 := &system.Base{Context: ptr8729635872, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729282976 := &system.String_rule{Base: ptr8729244592, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729293664 := &system.Type{Base: ptr8729244480, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729282976}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Id: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729198016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243136 := &system.Base{Context: ptr8729198016, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729292864 := &system.Type{Base: ptr8729243136, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729634240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243808 := &system.Base{Context: ptr8729634240, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729293184 := &system.Type{Base: ptr8729243808, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729190496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246384 := &system.Base{Context: ptr8729190496, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728852480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246496 := &system.Base{Context: ptr8728852480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729305568 := &system.Number_rule{Base: ptr8729246496, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728852896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246608 := &system.Base{Context: ptr8728852896, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729305664 := &system.Number_rule{Base: ptr8729246608, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728853248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246720 := &system.Base{Context: ptr8728853248, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "@c", Exists: true}}

	ptr8729460016 := &selectors.C_rule{Base: ptr8729246720, RuleBase: (*system.RuleBase)(nil)}

	ptr8728853888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246832 := &system.Base{Context: ptr8728853888, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8728853728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729247056 := &system.Base{Context: ptr8728853728, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729305760 := &system.Number_rule{Base: ptr8729247056, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729501952 := &system.Map_rule{Base: ptr8729246832, RuleBase: (*system.RuleBase)(nil), Items: ptr8729305760, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728854496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729247280 := &system.Base{Context: ptr8728854496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8728854336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729247392 := &system.Base{Context: ptr8728854336, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729305856 := &system.Number_rule{Base: ptr8729247392, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729502080 := &system.Map_rule{Base: ptr8729247280, RuleBase: (*system.RuleBase)(nil), Items: ptr8729305856, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729296384 := &system.Type{Base: ptr8729246384, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729305568, "b": ptr8729305664, "c": ptr8729460016, "d": ptr8729501952, "e": ptr8729502080}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729193920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729237296 := &system.Base{Context: ptr8729193920, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729175616 := &system.Type{Base: ptr8729237296, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729636352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244704 := &system.Base{Context: ptr8729636352, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729293984 := &system.Type{Base: ptr8729244704, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729635328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244368 := &system.Base{Context: ptr8729635328, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729293504 := &system.Type{Base: ptr8729244368, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729635456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240224 := &system.Base{Context: ptr8729635456, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729294304 := &system.Type{Base: ptr8729240224, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729634944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244816 := &system.Base{Context: ptr8729634944, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729636768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729244928 := &system.Base{Context: ptr8729636768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729308256 := &system.Number_rule{Base: ptr8729244928, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729632832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245040 := &system.Base{Context: ptr8729632832, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729308352 := &system.Number_rule{Base: ptr8729245040, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729633312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729239664 := &system.Base{Context: ptr8729633312, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729283152 := &system.String_rule{Base: ptr8729239664, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729239776 := &system.Base{Context: ptr8729633632, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729283328 := &system.String_rule{Base: ptr8729239776, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729239888 := &system.Base{Context: ptr8729633984, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729283504 := &system.String_rule{Base: ptr8729239888, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729634432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240000 := &system.Base{Context: ptr8729634432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8729634336 := &system.Bool_rule{Base: ptr8729240000, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729634752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240112 := &system.Base{Context: ptr8729634752, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8729634656 := &system.Bool_rule{Base: ptr8729240112, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729294144 := &system.Type{Base: ptr8729244816, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"int": ptr8729308256, "float": ptr8729308352, "string": ptr8729283152, "string2": ptr8729283328, "null": ptr8729283504, "true": ptr8729634336, "false": ptr8729634656}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729638272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241120 := &system.Base{Context: ptr8729638272, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729637504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241232 := &system.Base{Context: ptr8729637504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729283680 := &system.String_rule{Base: ptr8729241232, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729637792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241344 := &system.Base{Context: ptr8729637792, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729283856 := &system.String_rule{Base: ptr8729241344, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729638144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245152 := &system.Base{Context: ptr8729638144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8729638112 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729638016 := &system.Bool_rule{Base: ptr8729245152, RuleBase: ptr8729638112, Default: system.Bool{Value: false, Exists: false}}

	ptr8729295104 := &system.Type{Base: ptr8729241120, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729283680, "level": ptr8729283856, "preferred": ptr8729638016}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729636480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240336 := &system.Base{Context: ptr8729636480, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729636192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240448 := &system.Base{Context: ptr8729636192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729636128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240560 := &system.Base{Context: ptr8729636128, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "@image", Exists: true}}

	ptr8729455056 := &selectors.Image_rule{Base: ptr8729240560, RuleBase: (*system.RuleBase)(nil)}

	ptr8729501888 := &system.Map_rule{Base: ptr8729240448, RuleBase: (*system.RuleBase)(nil), Items: ptr8729455056, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729294464 := &system.Type{Base: ptr8729240336, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729501888}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729640928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729241008 := &system.Base{Context: ptr8729640928, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729294944 := &system.Type{Base: ptr8729241008, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729640064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245376 := &system.Base{Context: ptr8729640064, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729639232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245488 := &system.Base{Context: ptr8729639232, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729639200 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729284032 := &system.String_rule{Base: ptr8729245488, RuleBase: ptr8729639200, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729639552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245600 := &system.Base{Context: ptr8729639552, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729284208 := &system.String_rule{Base: ptr8729245600, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729639936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729245712 := &system.Base{Context: ptr8729639936, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729284384 := &system.String_rule{Base: ptr8729245712, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729295584 := &system.Type{Base: ptr8729245376, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8729284032, "server": ptr8729284208, "path": ptr8729284384}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Id: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729636864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240784 := &system.Base{Context: ptr8729636864, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729294624 := &system.Type{Base: ptr8729240784, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729637088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729240896 := &system.Base{Context: ptr8729637088, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729294784 := &system.Type{Base: ptr8729240896, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728852064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729246272 := &system.Base{Context: ptr8728852064, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729296224 := &system.Type{Base: ptr8729246272, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729193568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8728876848 := &system.Base{Context: ptr8729193568, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729191488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8728877072 := &system.Base{Context: ptr8729191488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729191424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8728877296 := &system.Base{Context: ptr8729191424, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729284560 := &system.String_rule{Base: ptr8728877296, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729502144 := &system.Map_rule{Base: ptr8728877072, RuleBase: (*system.RuleBase)(nil), Items: ptr8729284560, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729191840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729236288 := &system.Base{Context: ptr8729191840, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729284736 := &system.String_rule{Base: ptr8729236288, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729192256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729236512 := &system.Base{Context: ptr8729192256, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729192224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729236736 := &system.Base{Context: ptr8729192224, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "@kid", Exists: true}}

	ptr8729011872 := &selectors.Kid_rule{Base: ptr8729236736, RuleBase: (*system.RuleBase)(nil)}

	ptr8729502208 := &system.Map_rule{Base: ptr8729236512, RuleBase: (*system.RuleBase)(nil), Items: ptr8729011872, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729192480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729236848 := &system.Base{Context: ptr8729192480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Id: "@image", Exists: true}}

	ptr8729013072 := &selectors.Image_rule{Base: ptr8729236848, RuleBase: (*system.RuleBase)(nil)}

	ptr8729192960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729236960 := &system.Base{Context: ptr8729192960, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8729192928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729237072 := &system.Base{Context: ptr8729192928, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729284912 := &system.String_rule{Base: ptr8729237072, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729210384 := &system.Array_rule{Base: ptr8729236960, RuleBase: (*system.RuleBase)(nil), Items: ptr8729284912, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729193312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729237184 := &system.Base{Context: ptr8729193312, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729305952 := &system.Number_rule{Base: ptr8729237184, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729296704 := &system.Type{Base: ptr8728876848, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729502144, "favoriteColor": ptr8729284736, "kids": ptr8729502208, "avatar": ptr8729013072, "drinkPreference": ptr8729210384, "weight": ptr8729305952}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729633888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243248 := &system.Base{Context: ptr8729633888, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729198432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243360 := &system.Base{Context: ptr8729198432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729307968 := &system.Number_rule{Base: ptr8729243360, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729633024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243472 := &system.Base{Context: ptr8729633024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729308064 := &system.Number_rule{Base: ptr8729243472, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729633760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243584 := &system.Base{Context: ptr8729633760, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729633600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729243696 := &system.Base{Context: ptr8729633600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729308160 := &system.Number_rule{Base: ptr8729243696, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729502656 := &system.Map_rule{Base: ptr8729243584, RuleBase: (*system.RuleBase)(nil), Items: ptr8729308160, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729293024 := &system.Type{Base: ptr8729243248, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729307968, "b": ptr8729308064, "c": ptr8729502656}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors:@basic", ptr8729292864)

	system.RegisterType("kego.io/selectors:@c", ptr8729293184)

	system.RegisterType("kego.io/selectors:@collision", ptr8729293504)

	system.RegisterType("kego.io/selectors:@diagram", ptr8729293984)

	system.RegisterType("kego.io/selectors:@expr", ptr8729294304)

	system.RegisterType("kego.io/selectors:@gallery", ptr8729294624)

	system.RegisterType("kego.io/selectors:@image", ptr8729294944)

	system.RegisterType("kego.io/selectors:@kid", ptr8729295424)

	system.RegisterType("kego.io/selectors:@photo", ptr8729295904)

	system.RegisterType("kego.io/selectors:@polykids", ptr8729296224)

	system.RegisterType("kego.io/selectors:@sibling", ptr8729296544)

	system.RegisterType("kego.io/selectors:@typed", ptr8729175616)

	system.RegisterType("kego.io/selectors:basic", ptr8729295264)

	system.RegisterType("kego.io/selectors:c", ptr8729293024)

	system.RegisterType("kego.io/selectors:collision", ptr8729293344)

	system.RegisterType("kego.io/selectors:diagram", ptr8729293664)

	system.RegisterType("kego.io/selectors:expr", ptr8729294144)

	system.RegisterType("kego.io/selectors:gallery", ptr8729294464)

	system.RegisterType("kego.io/selectors:image", ptr8729294784)

	system.RegisterType("kego.io/selectors:kid", ptr8729295104)

	system.RegisterType("kego.io/selectors:photo", ptr8729295584)

	system.RegisterType("kego.io/selectors:polykids", ptr8729296064)

	system.RegisterType("kego.io/selectors:sibling", ptr8729296384)

	system.RegisterType("kego.io/selectors:typed", ptr8729296704)

	var _ selectors.Nothing

}
