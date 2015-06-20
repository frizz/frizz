package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729641216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567488 := &system.Base{Context: ptr8729641216, Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729295104 := &system.Type{Base: ptr8729567488, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729504192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318912 := &system.Base{Context: ptr8729504192, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729503616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316992 := &system.Base{Context: ptr8729503616, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729503584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317120 := &system.Base{Context: ptr8729503584, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729298832 := &system.String_rule{Base: ptr8729317120, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729387168 := &system.Array_rule{Base: ptr8729316992, RuleBase: (*system.RuleBase)(nil), Items: ptr8729298832, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729503936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317248 := &system.Base{Context: ptr8729503936, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729681920 := &system.Number_rule{Base: ptr8729317248, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729501792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319040 := &system.Base{Context: ptr8729501792, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729501760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319168 := &system.Base{Context: ptr8729501760, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729298128 := &system.String_rule{Base: ptr8729319168, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729030528 := &system.Map_rule{Base: ptr8729319040, RuleBase: (*system.RuleBase)(nil), Items: ptr8729298128, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729502048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316224 := &system.Base{Context: ptr8729502048, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729298304 := &system.String_rule{Base: ptr8729316224, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729502592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316352 := &system.Base{Context: ptr8729502592, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729502560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316480 := &system.Base{Context: ptr8729502560, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729502528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316608 := &system.Base{Context: ptr8729502528, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729298480 := &system.String_rule{Base: ptr8729316608, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729027328 := &system.Map_rule{Base: ptr8729316480, RuleBase: (*system.RuleBase)(nil), Items: ptr8729298480, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729387008 := &system.Array_rule{Base: ptr8729316352, RuleBase: (*system.RuleBase)(nil), Items: ptr8729027328, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729503104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316736 := &system.Base{Context: ptr8729503104, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729503072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729316864 := &system.Base{Context: ptr8729503072, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729298656 := &system.String_rule{Base: ptr8729316864, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729387088 := &system.Array_rule{Base: ptr8729316736, RuleBase: (*system.RuleBase)(nil), Items: ptr8729298656, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729292704 := &system.Type{Base: ptr8729318912, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"drinkPreference": ptr8729387168, "weight": ptr8729681920, "name": ptr8729030528, "favoriteColor": ptr8729298304, "languagesSpoken": ptr8729387008, "seatingPreference": ptr8729387088}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729507968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318784 := &system.Base{Context: ptr8729507968, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729507840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319296 := &system.Base{Context: ptr8729507840, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729299360 := &system.String_rule{Base: ptr8729319296, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729293504 := &system.Type{Base: ptr8729318784, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729299360}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729642816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568128 := &system.Base{Context: ptr8729642816, Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729295744 := &system.Type{Base: ptr8729568128, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729642880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568768 := &system.Base{Context: ptr8729642880, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729641056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569024 := &system.Base{Context: ptr8729641056, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682592 := &system.Number_rule{Base: ptr8729569024, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729641440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569152 := &system.Base{Context: ptr8729641440, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}

	ptr8729020336 := &selectors.C_rule{Base: ptr8729569152, RuleBase: (*system.RuleBase)(nil)}

	ptr8729642048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569280 := &system.Base{Context: ptr8729642048, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729641888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569408 := &system.Base{Context: ptr8729641888, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682688 := &system.Number_rule{Base: ptr8729569408, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729028608 := &system.Map_rule{Base: ptr8729569280, RuleBase: (*system.RuleBase)(nil), Items: ptr8729682688, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729642720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569536 := &system.Base{Context: ptr8729642720, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729642496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569664 := &system.Base{Context: ptr8729642496, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682784 := &system.Number_rule{Base: ptr8729569664, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729028672 := &system.Map_rule{Base: ptr8729569536, RuleBase: (*system.RuleBase)(nil), Items: ptr8729682784, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729644128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568896 := &system.Base{Context: ptr8729644128, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682496 := &system.Number_rule{Base: ptr8729568896, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729296224 := &system.Type{Base: ptr8729568768, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"b": ptr8729682592, "c": ptr8729020336, "d": ptr8729028608, "e": ptr8729028672, "a": ptr8729682496}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729643136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569792 := &system.Base{Context: ptr8729643136, Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729296384 := &system.Type{Base: ptr8729569792, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729507008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318272 := &system.Base{Context: ptr8729507008, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729506880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318400 := &system.Base{Context: ptr8729506880, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729506848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318528 := &system.Base{Context: ptr8729506848, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729299184 := &system.String_rule{Base: ptr8729318528, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729028928 := &system.Map_rule{Base: ptr8729318400, RuleBase: (*system.RuleBase)(nil), Items: ptr8729299184, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729293184 := &system.Type{Base: ptr8729318272, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729028928}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729643744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568640 := &system.Base{Context: ptr8729643744, Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729296064 := &system.Type{Base: ptr8729568640, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729646080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729571200 := &system.Base{Context: ptr8729646080, Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729296704 := &system.Type{Base: ptr8729571200, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729640960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729157632 := &system.Base{Context: ptr8729640960, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728862016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729158912 := &system.Base{Context: ptr8728862016, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729300064 := &system.String_rule{Base: ptr8729158912, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728862272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567232 := &system.Base{Context: ptr8728862272, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729300240 := &system.String_rule{Base: ptr8729567232, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728862592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567360 := &system.Base{Context: ptr8728862592, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728862560 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728862464 := &system.Bool_rule{Base: ptr8729567360, RuleBase: ptr8728862560, Default: system.Bool{Value: false, Exists: false}}

	ptr8729294944 := &system.Type{Base: ptr8729157632, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729300064, "level": ptr8729300240, "preferred": ptr8728862464}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729506336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318144 := &system.Base{Context: ptr8729506336, Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729293024 := &system.Type{Base: ptr8729318144, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728861632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729321344 := &system.Base{Context: ptr8728861632, Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729294784 := &system.Type{Base: ptr8729321344, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729642560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567616 := &system.Base{Context: ptr8729642560, Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729641792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567744 := &system.Base{Context: ptr8729641792, Description: "The protocol for the url - e.g. http or https", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729641760 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729300416 := &system.String_rule{Base: ptr8729567744, RuleBase: ptr8729641760, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729642080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729567872 := &system.Base{Context: ptr8729642080, Description: "The server for the url - e.g. www.google.com", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729300592 := &system.String_rule{Base: ptr8729567872, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729642432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568000 := &system.Base{Context: ptr8729642432, Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729300768 := &system.String_rule{Base: ptr8729568000, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729295424 := &system.Type{Base: ptr8729567616, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8729300416, "server": ptr8729300592, "path": ptr8729300768}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728861376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729321216 := &system.Base{Context: ptr8728861376, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729294624 := &system.Type{Base: ptr8729321216, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729504448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317376 := &system.Base{Context: ptr8729504448, Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729295264 := &system.Type{Base: ptr8729317376, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729507392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318656 := &system.Base{Context: ptr8729507392, Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729293344 := &system.Type{Base: ptr8729318656, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728860288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320576 := &system.Base{Context: ptr8728860288, Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729294144 := &system.Type{Base: ptr8729320576, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728860928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320704 := &system.Base{Context: ptr8728860928, Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728860800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320832 := &system.Base{Context: ptr8728860800, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728860768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320960 := &system.Base{Context: ptr8728860768, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729597440 := &selectors.Image_rule{Base: ptr8729320960, RuleBase: (*system.RuleBase)(nil)}

	ptr8729029056 := &system.Map_rule{Base: ptr8729320832, RuleBase: (*system.RuleBase)(nil), Items: ptr8729597440, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729294304 := &system.Type{Base: ptr8729320704, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729029056}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728860032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319552 := &system.Base{Context: ptr8728860032, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728859328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320192 := &system.Base{Context: ptr8728859328, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729299888 := &system.String_rule{Base: ptr8729320192, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728859552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320320 := &system.Base{Context: ptr8728859552, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728859456 := &system.Bool_rule{Base: ptr8729320320, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8728859904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320448 := &system.Base{Context: ptr8728859904, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728859808 := &system.Bool_rule{Base: ptr8729320448, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729508640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319680 := &system.Base{Context: ptr8729508640, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682304 := &system.Number_rule{Base: ptr8729319680, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729509024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319808 := &system.Base{Context: ptr8729509024, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682400 := &system.Number_rule{Base: ptr8729319808, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729509408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319936 := &system.Base{Context: ptr8729509408, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729299536 := &system.String_rule{Base: ptr8729319936, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729509664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729320064 := &system.Base{Context: ptr8729509664, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729299712 := &system.String_rule{Base: ptr8729320064, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729293984 := &system.Type{Base: ptr8729319552, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"null": ptr8729299888, "true": ptr8728859456, "false": ptr8728859808, "int": ptr8729682304, "float": ptr8729682400, "string": ptr8729299536, "string2": ptr8729299712}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729508224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729319424 := &system.Base{Context: ptr8729508224, Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729293824 := &system.Type{Base: ptr8729319424, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729506080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317504 := &system.Base{Context: ptr8729506080, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729504832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317632 := &system.Base{Context: ptr8729504832, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682016 := &system.Number_rule{Base: ptr8729317632, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729505248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317760 := &system.Base{Context: ptr8729505248, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682112 := &system.Number_rule{Base: ptr8729317760, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729505952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729317888 := &system.Base{Context: ptr8729505952, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729505792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729318016 := &system.Base{Context: ptr8729505792, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682208 := &system.Number_rule{Base: ptr8729318016, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729028864 := &system.Map_rule{Base: ptr8729317888, RuleBase: (*system.RuleBase)(nil), Items: ptr8729682208, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729292864 := &system.Type{Base: ptr8729317504, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729682016, "b": ptr8729682112, "c": ptr8729028864}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728861184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729321088 := &system.Base{Context: ptr8728861184, Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729294464 := &system.Type{Base: ptr8729321088, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729645824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729569920 := &system.Base{Context: ptr8729645824, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729643904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570048 := &system.Base{Context: ptr8729643904, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729643872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570176 := &system.Base{Context: ptr8729643872, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729300944 := &system.String_rule{Base: ptr8729570176, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729028800 := &system.Map_rule{Base: ptr8729570048, RuleBase: (*system.RuleBase)(nil), Items: ptr8729300944, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729644192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570304 := &system.Base{Context: ptr8729644192, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729301120 := &system.String_rule{Base: ptr8729570304, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729644576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570432 := &system.Base{Context: ptr8729644576, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729644544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570560 := &system.Base{Context: ptr8729644544, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729022048 := &selectors.Kid_rule{Base: ptr8729570560, RuleBase: (*system.RuleBase)(nil)}

	ptr8729030464 := &system.Map_rule{Base: ptr8729570432, RuleBase: (*system.RuleBase)(nil), Items: ptr8729022048, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729644768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570688 := &system.Base{Context: ptr8729644768, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729022416 := &selectors.Image_rule{Base: ptr8729570688, RuleBase: (*system.RuleBase)(nil)}

	ptr8729645248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570816 := &system.Base{Context: ptr8729645248, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729645216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729570944 := &system.Base{Context: ptr8729645216, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729301296 := &system.String_rule{Base: ptr8729570944, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729387648 := &system.Array_rule{Base: ptr8729570816, RuleBase: (*system.RuleBase)(nil), Items: ptr8729301296, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729645568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729571072 := &system.Base{Context: ptr8729645568, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729682880 := &system.Number_rule{Base: ptr8729571072, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729296544 := &system.Type{Base: ptr8729569920, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"name": ptr8729028800, "favoriteColor": ptr8729301120, "kids": ptr8729030464, "avatar": ptr8729022416, "drinkPreference": ptr8729387648, "weight": ptr8729682880}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729643488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568256 := &system.Base{Context: ptr8729643488, Description: "", Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729643296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568384 := &system.Base{Context: ptr8729643296, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729643264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729568512 := &system.Base{Context: ptr8729643264, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729020272 := &selectors.Kid_rule{Base: ptr8729568512, RuleBase: (*system.RuleBase)(nil)}

	ptr8729390768 := &system.Array_rule{Base: ptr8729568384, RuleBase: (*system.RuleBase)(nil), Items: ptr8729020272, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729295904 := &system.Type{Base: ptr8729568256, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729390768}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors", "expr", ptr8729293984)

	system.RegisterType("kego.io/selectors", "@gallery", ptr8729294464)

	system.RegisterType("kego.io/selectors", "@photo", ptr8729295744)

	system.RegisterType("kego.io/selectors", "sibling", ptr8729296224)

	system.RegisterType("kego.io/selectors", "@polykids", ptr8729296064)

	system.RegisterType("kego.io/selectors", "@image", ptr8729294784)

	system.RegisterType("kego.io/selectors", "photo", ptr8729295424)

	system.RegisterType("kego.io/selectors", "@collision", ptr8729293344)

	system.RegisterType("kego.io/selectors", "polykids", ptr8729295904)

	system.RegisterType("kego.io/selectors", "typed", ptr8729296544)

	system.RegisterType("kego.io/selectors", "@kid", ptr8729295104)

	system.RegisterType("kego.io/selectors", "basic", ptr8729292704)

	system.RegisterType("kego.io/selectors", "diagram", ptr8729293504)

	system.RegisterType("kego.io/selectors", "@sibling", ptr8729296384)

	system.RegisterType("kego.io/selectors", "@expr", ptr8729294144)

	system.RegisterType("kego.io/selectors", "@diagram", ptr8729293824)

	system.RegisterType("kego.io/selectors", "@typed", ptr8729296704)

	system.RegisterType("kego.io/selectors", "gallery", ptr8729294304)

	system.RegisterType("kego.io/selectors", "collision", ptr8729293184)

	system.RegisterType("kego.io/selectors", "kid", ptr8729294944)

	system.RegisterType("kego.io/selectors", "@c", ptr8729293024)

	system.RegisterType("kego.io/selectors", "image", ptr8729294624)

	system.RegisterType("kego.io/selectors", "@basic", ptr8729295264)

	system.RegisterType("kego.io/selectors", "c", ptr8729292864)

	var _ selectors.Nothing

}
