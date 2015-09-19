package types

import (
	"kego.io/selectors/tests"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Base{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Base: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Base{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Base: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Base{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Base: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Base: ptr10, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Base{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Base: ptr12, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Base{Description: "Automatically created basic rule for instance", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instance", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Base: ptr14, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Base{Description: "Automatically created basic rule for instanceItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Base: ptr16, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Type{Base: ptr18, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr20 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr21 := &system.Type{Base: ptr20, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Type{Base: ptr22, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Base{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Type{Base: ptr24, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr26 := &system.Base{Description: "Automatically created basic rule for rightscale", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr27 := &system.Type{Base: ptr26, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr28 := &system.Base{Description: "Automatically created basic rule for rightscaleLink", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr29 := &system.Type{Base: ptr28, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr30 := &system.Base{Description: "Automatically created basic rule for rightscaleList", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleList", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr31 := &system.Type{Base: ptr30, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr32 := &system.Base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr33 := &system.Type{Base: ptr32, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34 := &system.Base{Description: "Automatically created basic rule for simple", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr35 := &system.Type{Base: ptr34, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr36 := &system.Base{Description: "Automatically created basic rule for simpleItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr37 := &system.Type{Base: ptr36, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr38 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr39 := &system.Type{Base: ptr38, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr40 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr41 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr42 := &system.RuleBase{}
	ptr43 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr44 := &system.RuleBase{}
	ptr45 := &system.String_rule{Base: ptr43, RuleBase: ptr44, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr46 := &system.Array_rule{Base: ptr41, RuleBase: ptr42, Items: ptr45, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr47 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr48 := &system.RuleBase{}
	ptr49 := &system.String_rule{Base: ptr47, RuleBase: ptr48, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr50 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr51 := &system.RuleBase{}
	ptr52 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr53 := &system.RuleBase{}
	ptr54 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr55 := &system.RuleBase{}
	ptr56 := &system.String_rule{Base: ptr54, RuleBase: ptr55, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr57 := &system.Map_rule{Base: ptr52, RuleBase: ptr53, Items: ptr56, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr58 := &system.Array_rule{Base: ptr50, RuleBase: ptr51, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr60 := &system.RuleBase{}
	ptr61 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr62 := &system.RuleBase{}
	ptr63 := &system.String_rule{Base: ptr61, RuleBase: ptr62, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr64 := &system.Map_rule{Base: ptr59, RuleBase: ptr60, Items: ptr63, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr65 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr66 := &system.RuleBase{}
	ptr67 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr68 := &system.RuleBase{}
	ptr69 := &system.String_rule{Base: ptr67, RuleBase: ptr68, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr70 := &system.Array_rule{Base: ptr65, RuleBase: ptr66, Items: ptr69, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr71 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr72 := &system.RuleBase{}
	ptr73 := &system.Number_rule{Base: ptr71, RuleBase: ptr72, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr74 := &system.Type{Base: ptr40, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"drinkPreference": ptr46, "favoriteColor": ptr49, "languagesSpoken": ptr58, "name": ptr64, "seatingPreference": ptr70, "weight": ptr73}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr77 := &system.RuleBase{}
	ptr78 := &system.Number_rule{Base: ptr76, RuleBase: ptr77, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr79 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr80 := &system.RuleBase{}
	ptr81 := &system.Number_rule{Base: ptr79, RuleBase: ptr80, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr82 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr83 := &system.RuleBase{}
	ptr84 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr85 := &system.RuleBase{}
	ptr86 := &system.Number_rule{Base: ptr84, RuleBase: ptr85, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr87 := &system.Map_rule{Base: ptr82, RuleBase: ptr83, Items: ptr86, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr88 := &system.Type{Base: ptr75, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr78, "b": ptr81, "c": ptr87}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr89 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr90 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr91 := &system.RuleBase{}
	ptr92 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr93 := &system.RuleBase{}
	ptr94 := &system.String_rule{Base: ptr92, RuleBase: ptr93, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr95 := &system.Map_rule{Base: ptr90, RuleBase: ptr91, Items: ptr94, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr96 := &system.Type{Base: ptr89, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr95}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr97 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr98 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr99 := &system.RuleBase{}
	ptr100 := &system.String_rule{Base: ptr98, RuleBase: ptr99, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr101 := &system.Type{Base: ptr97, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr100}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr104 := &system.RuleBase{}
	ptr105 := &system.Bool_rule{Base: ptr103, RuleBase: ptr104, Default: system.Bool{}}
	ptr106 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr107 := &system.RuleBase{}
	ptr108 := &system.Number_rule{Base: ptr106, RuleBase: ptr107, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr109 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr110 := &system.RuleBase{}
	ptr111 := &system.Number_rule{Base: ptr109, RuleBase: ptr110, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr112 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr113 := &system.RuleBase{}
	ptr114 := &system.String_rule{Base: ptr112, RuleBase: ptr113, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr115 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr116 := &system.RuleBase{}
	ptr117 := &system.String_rule{Base: ptr115, RuleBase: ptr116, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr118 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr119 := &system.RuleBase{}
	ptr120 := &system.String_rule{Base: ptr118, RuleBase: ptr119, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr121 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr122 := &system.RuleBase{}
	ptr123 := &system.Bool_rule{Base: ptr121, RuleBase: ptr122, Default: system.Bool{}}
	ptr124 := &system.Type{Base: ptr102, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"false": ptr105, "float": ptr108, "int": ptr111, "null": ptr114, "string": ptr117, "string2": ptr120, "true": ptr123}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr125 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr126 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr127 := &system.RuleBase{}
	ptr128 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr129 := &system.RuleBase{}
	ptr130 := &tests.Image_rule{Base: ptr128, RuleBase: ptr129}
	ptr131 := &system.Map_rule{Base: ptr126, RuleBase: ptr127, Items: ptr130, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr132 := &system.Type{Base: ptr125, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr131}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr133 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr134 := &system.Type{Base: ptr133, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr135 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instance", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr136 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr137 := &system.RuleBase{}
	ptr138 := &system.String_rule{Base: ptr136, RuleBase: ptr137, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr139 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr140 := &system.RuleBase{}
	ptr141 := &system.String_rule{Base: ptr139, RuleBase: ptr140, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr142 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr143 := &system.RuleBase{}
	ptr144 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr145 := &system.RuleBase{}
	ptr146 := &tests.RightscaleLink_rule{Base: ptr144, RuleBase: ptr145}
	ptr147 := &system.Array_rule{Base: ptr142, RuleBase: ptr143, Items: ptr146, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr148 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr149 := &system.RuleBase{}
	ptr150 := &system.String_rule{Base: ptr148, RuleBase: ptr149, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr151 := &system.Type{Base: ptr135, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"cloud_type": ptr138, "display_name": ptr141, "links": ptr147, "name": ptr150}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr152 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instanceItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr153 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr154 := &system.RuleBase{}
	ptr155 := &system.String_rule{Base: ptr153, RuleBase: ptr154, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr156 := &system.Type{Base: ptr152, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr155}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr157 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr158 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr159 := &system.RuleBase{}
	ptr160 := &system.String_rule{Base: ptr158, RuleBase: ptr159, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr161 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr162 := &system.RuleBase{}
	ptr163 := &system.String_rule{Base: ptr161, RuleBase: ptr162, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr164 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr165 := &system.RuleBase{Optional: true}
	ptr166 := &system.Bool_rule{Base: ptr164, RuleBase: ptr165, Default: system.Bool{}}
	ptr167 := &system.Type{Base: ptr157, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr160, "level": ptr163, "preferred": ptr166}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr168 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr169 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr170 := &system.RuleBase{}
	ptr171 := &system.String_rule{Base: ptr169, RuleBase: ptr170, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr172 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr173 := &system.RuleBase{Optional: true}
	ptr174 := &system.String_rule{Base: ptr172, RuleBase: ptr173, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr175 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr176 := &system.RuleBase{}
	ptr177 := &system.String_rule{Base: ptr175, RuleBase: ptr176, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr178 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr179 := &system.RuleBase{Optional: true}
	ptr180 := &tests.Rectangle_rule{Base: ptr178, RuleBase: ptr179}
	ptr181 := &system.Type{Base: ptr168, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr171, "protocol": ptr174, "server": ptr177, "size": ptr180}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr182 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr183 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr184 := &system.RuleBase{}
	ptr185 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr186 := &system.RuleBase{}
	ptr187 := &tests.Kid_rule{Base: ptr185, RuleBase: ptr186}
	ptr188 := &system.Array_rule{Base: ptr183, RuleBase: ptr184, Items: ptr187, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr189 := &system.Type{Base: ptr182, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr188}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr190 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr191 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr192 := &system.RuleBase{}
	ptr193 := &system.Int_rule{Base: ptr191, RuleBase: ptr192, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr194 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr195 := &system.RuleBase{}
	ptr196 := &system.Int_rule{Base: ptr194, RuleBase: ptr195, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr197 := &system.Type{Base: ptr190, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr193, "width": ptr196}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr198 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscale", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr199 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr200 := &system.RuleBase{}
	ptr201 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr202 := &system.RuleBase{}
	ptr203 := &tests.InstanceItem_rule{Base: ptr201, RuleBase: ptr202}
	ptr204 := &system.Map_rule{Base: ptr199, RuleBase: ptr200, Items: ptr203, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr205 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr206 := &system.RuleBase{}
	ptr207 := &system.String_rule{Base: ptr205, RuleBase: ptr206, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr208 := &system.Type{Base: ptr198, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"child": ptr204, "name": ptr207}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr209 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleLink", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr210 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr211 := &system.RuleBase{}
	ptr212 := &system.String_rule{Base: ptr210, RuleBase: ptr211, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr213 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr214 := &system.RuleBase{}
	ptr215 := &system.String_rule{Base: ptr213, RuleBase: ptr214, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr216 := &system.Type{Base: ptr209, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"href": ptr212, "rel": ptr215}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr217 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleList", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr218 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr219 := &system.RuleBase{}
	ptr220 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}}
	ptr221 := &system.RuleBase{}
	ptr222 := &tests.Rightscale_rule{Base: ptr220, RuleBase: ptr221}
	ptr223 := &system.Array_rule{Base: ptr218, RuleBase: ptr219, Items: ptr222, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr224 := &system.Type{Base: ptr217, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"foo": ptr223}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr225 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr226 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr227 := &system.RuleBase{}
	ptr228 := &system.Number_rule{Base: ptr226, RuleBase: ptr227, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr229 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr230 := &system.RuleBase{}
	ptr231 := &system.Number_rule{Base: ptr229, RuleBase: ptr230, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr232 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr233 := &system.RuleBase{}
	ptr234 := &tests.C_rule{Base: ptr232, RuleBase: ptr233}
	ptr235 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr236 := &system.RuleBase{}
	ptr237 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr238 := &system.RuleBase{}
	ptr239 := &system.Number_rule{Base: ptr237, RuleBase: ptr238, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr240 := &system.Map_rule{Base: ptr235, RuleBase: ptr236, Items: ptr239, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr241 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr242 := &system.RuleBase{}
	ptr243 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr244 := &system.RuleBase{}
	ptr245 := &system.Number_rule{Base: ptr243, RuleBase: ptr244, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr246 := &system.Map_rule{Base: ptr241, RuleBase: ptr242, Items: ptr245, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr247 := &system.Type{Base: ptr225, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr228, "b": ptr231, "c": ptr234, "d": ptr240, "e": ptr246}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr248 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr249 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}}
	ptr250 := &system.RuleBase{}
	ptr251 := &tests.SimpleItem_rule{Base: ptr249, RuleBase: ptr250}
	ptr252 := &system.Type{Base: ptr248, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr251}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr253 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simpleItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr254 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr255 := &system.RuleBase{}
	ptr256 := &system.String_rule{Base: ptr254, RuleBase: ptr255, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr257 := &system.Type{Base: ptr253, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"b": ptr256}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr258 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr259 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr260 := &system.RuleBase{}
	ptr261 := &tests.Image_rule{Base: ptr259, RuleBase: ptr260}
	ptr262 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr263 := &system.RuleBase{}
	ptr264 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr265 := &system.RuleBase{}
	ptr266 := &system.String_rule{Base: ptr264, RuleBase: ptr265, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr267 := &system.Array_rule{Base: ptr262, RuleBase: ptr263, Items: ptr266, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr268 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr269 := &system.RuleBase{}
	ptr270 := &system.String_rule{Base: ptr268, RuleBase: ptr269, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr271 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr272 := &system.RuleBase{}
	ptr273 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr274 := &system.RuleBase{}
	ptr275 := &tests.Kid_rule{Base: ptr273, RuleBase: ptr274}
	ptr276 := &system.Map_rule{Base: ptr271, RuleBase: ptr272, Items: ptr275, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr277 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr278 := &system.RuleBase{}
	ptr279 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr280 := &system.RuleBase{}
	ptr281 := &system.String_rule{Base: ptr279, RuleBase: ptr280, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr282 := &system.Map_rule{Base: ptr277, RuleBase: ptr278, Items: ptr281, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr283 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr284 := &system.RuleBase{}
	ptr285 := &system.Number_rule{Base: ptr283, RuleBase: ptr284, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr286 := &system.Type{Base: ptr258, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"avatar": ptr261, "drinkPreference": ptr267, "favoriteColor": ptr270, "kids": ptr276, "name": ptr282, "weight": ptr285}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/selectors/tests", "@basic", ptr1, 0xe2f725e9d48b69d4)
	system.Register("kego.io/selectors/tests", "@c", ptr3, 0x6a4bf0b03046c68a)
	system.Register("kego.io/selectors/tests", "@collision", ptr5, 0x2f6382cf84b27ee7)
	system.Register("kego.io/selectors/tests", "@diagram", ptr7, 0x2076b0eaf334855b)
	system.Register("kego.io/selectors/tests", "@expr", ptr9, 0x6214e678b1df35e3)
	system.Register("kego.io/selectors/tests", "@gallery", ptr11, 0xa2261dbb985b3d3)
	system.Register("kego.io/selectors/tests", "@image", ptr13, 0x2a2a6f416ac31013)
	system.Register("kego.io/selectors/tests", "@instance", ptr15, 0x91954502a09aa42a)
	system.Register("kego.io/selectors/tests", "@instanceItem", ptr17, 0x2dca71ec3918c621)
	system.Register("kego.io/selectors/tests", "@kid", ptr19, 0x3b576fd38b04a2eb)
	system.Register("kego.io/selectors/tests", "@photo", ptr21, 0x9b14a4b75f8931a8)
	system.Register("kego.io/selectors/tests", "@polykids", ptr23, 0xebaffecb552df4d1)
	system.Register("kego.io/selectors/tests", "@rectangle", ptr25, 0x21539507256190ab)
	system.Register("kego.io/selectors/tests", "@rightscale", ptr27, 0x19db6e12b53f9744)
	system.Register("kego.io/selectors/tests", "@rightscaleLink", ptr29, 0x127d9489c45697ea)
	system.Register("kego.io/selectors/tests", "@rightscaleList", ptr31, 0x6a9527067852e20b)
	system.Register("kego.io/selectors/tests", "@sibling", ptr33, 0xac837d672ff74ed6)
	system.Register("kego.io/selectors/tests", "@simple", ptr35, 0x7054dc6469f4db46)
	system.Register("kego.io/selectors/tests", "@simpleItem", ptr37, 0xac167d73687ccb1f)
	system.Register("kego.io/selectors/tests", "@typed", ptr39, 0x216d0e1d4c22bb2d)
	system.Register("kego.io/selectors/tests", "basic", ptr74, 0xe2f725e9d48b69d4)
	system.Register("kego.io/selectors/tests", "c", ptr88, 0x6a4bf0b03046c68a)
	system.Register("kego.io/selectors/tests", "collision", ptr96, 0x2f6382cf84b27ee7)
	system.Register("kego.io/selectors/tests", "diagram", ptr101, 0x2076b0eaf334855b)
	system.Register("kego.io/selectors/tests", "expr", ptr124, 0x6214e678b1df35e3)
	system.Register("kego.io/selectors/tests", "gallery", ptr132, 0xa2261dbb985b3d3)
	system.Register("kego.io/selectors/tests", "image", ptr134, 0x2a2a6f416ac31013)
	system.Register("kego.io/selectors/tests", "instance", ptr151, 0x91954502a09aa42a)
	system.Register("kego.io/selectors/tests", "instanceItem", ptr156, 0x2dca71ec3918c621)
	system.Register("kego.io/selectors/tests", "kid", ptr167, 0x3b576fd38b04a2eb)
	system.Register("kego.io/selectors/tests", "photo", ptr181, 0x9b14a4b75f8931a8)
	system.Register("kego.io/selectors/tests", "polykids", ptr189, 0xebaffecb552df4d1)
	system.Register("kego.io/selectors/tests", "rectangle", ptr197, 0x21539507256190ab)
	system.Register("kego.io/selectors/tests", "rightscale", ptr208, 0x19db6e12b53f9744)
	system.Register("kego.io/selectors/tests", "rightscaleLink", ptr216, 0x127d9489c45697ea)
	system.Register("kego.io/selectors/tests", "rightscaleList", ptr224, 0x6a9527067852e20b)
	system.Register("kego.io/selectors/tests", "sibling", ptr247, 0xac837d672ff74ed6)
	system.Register("kego.io/selectors/tests", "simple", ptr252, 0x7054dc6469f4db46)
	system.Register("kego.io/selectors/tests", "simpleItem", ptr257, 0xac167d73687ccb1f)
	system.Register("kego.io/selectors/tests", "typed", ptr286, 0x216d0e1d4c22bb2d)
}
