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
	ptr34 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr35 := &system.Type{Base: ptr34, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr36 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr37 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr38 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr39 := &system.String_rule{Base: ptr38, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr40 := &system.Array_rule{Base: ptr37, RuleBase: (*system.RuleBase)(nil), Items: ptr39, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr41 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr42 := &system.String_rule{Base: ptr41, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr43 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr44 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr45 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr46 := &system.String_rule{Base: ptr45, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr47 := &system.Map_rule{Base: ptr44, RuleBase: (*system.RuleBase)(nil), Items: ptr46, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr48 := &system.Array_rule{Base: ptr43, RuleBase: (*system.RuleBase)(nil), Items: ptr47, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr49 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr50 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr51 := &system.String_rule{Base: ptr50, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr52 := &system.Map_rule{Base: ptr49, RuleBase: (*system.RuleBase)(nil), Items: ptr51, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr53 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr54 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr55 := &system.String_rule{Base: ptr54, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr56 := &system.Array_rule{Base: ptr53, RuleBase: (*system.RuleBase)(nil), Items: ptr55, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr57 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr58 := &system.Number_rule{Base: ptr57, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr59 := &system.Type{Base: ptr36, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"drinkPreference": ptr40, "favoriteColor": ptr42, "languagesSpoken": ptr48, "name": ptr52, "seatingPreference": ptr56, "weight": ptr58}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr60 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr61 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr62 := &system.Number_rule{Base: ptr61, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr63 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr64 := &system.Number_rule{Base: ptr63, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr65 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr66 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr67 := &system.Number_rule{Base: ptr66, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr68 := &system.Map_rule{Base: ptr65, RuleBase: (*system.RuleBase)(nil), Items: ptr67, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr69 := &system.Type{Base: ptr60, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr62, "b": ptr64, "c": ptr68}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr70 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr71 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr72 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr73 := &system.String_rule{Base: ptr72, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr74 := &system.Map_rule{Base: ptr71, RuleBase: (*system.RuleBase)(nil), Items: ptr73, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr75 := &system.Type{Base: ptr70, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr74}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr76 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr77 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr78 := &system.String_rule{Base: ptr77, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr79 := &system.Type{Base: ptr76, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr78}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr80 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr81 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr82 := &system.Bool_rule{Base: ptr81, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr83 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr84 := &system.Number_rule{Base: ptr83, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr85 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr86 := &system.Number_rule{Base: ptr85, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr87 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr88 := &system.String_rule{Base: ptr87, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr89 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr90 := &system.String_rule{Base: ptr89, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr91 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr92 := &system.String_rule{Base: ptr91, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr93 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr94 := &system.Bool_rule{Base: ptr93, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr95 := &system.Type{Base: ptr80, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"false": ptr82, "float": ptr84, "int": ptr86, "null": ptr88, "string": ptr90, "string2": ptr92, "true": ptr94}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr96 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr97 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr98 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr99 := &tests.Image_rule{Base: ptr98, RuleBase: (*system.RuleBase)(nil)}
	ptr100 := &system.Map_rule{Base: ptr97, RuleBase: (*system.RuleBase)(nil), Items: ptr99, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr101 := &system.Type{Base: ptr96, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr100}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Type{Base: ptr102, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr104 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instance", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr105 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr106 := &system.String_rule{Base: ptr105, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr107 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr108 := &system.String_rule{Base: ptr107, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr109 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr110 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr111 := &tests.RightscaleLink_rule{Base: ptr110, RuleBase: (*system.RuleBase)(nil)}
	ptr112 := &system.Array_rule{Base: ptr109, RuleBase: (*system.RuleBase)(nil), Items: ptr111, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr113 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr114 := &system.String_rule{Base: ptr113, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr115 := &system.Type{Base: ptr104, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"cloud_type": ptr106, "display_name": ptr108, "links": ptr112, "name": ptr114}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr116 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instanceItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr117 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr118 := &system.String_rule{Base: ptr117, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr119 := &system.Type{Base: ptr116, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr118}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr120 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr121 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr122 := &system.String_rule{Base: ptr121, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr123 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr124 := &system.String_rule{Base: ptr123, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr125 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr126 := &system.RuleBase{Optional: true}
	ptr127 := &system.Bool_rule{Base: ptr125, RuleBase: ptr126, Default: system.Bool{}}
	ptr128 := &system.Type{Base: ptr120, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr122, "level": ptr124, "preferred": ptr127}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr129 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr130 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr131 := &system.String_rule{Base: ptr130, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr132 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr133 := &system.RuleBase{Optional: true}
	ptr134 := &system.String_rule{Base: ptr132, RuleBase: ptr133, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr135 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr136 := &system.String_rule{Base: ptr135, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr137 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr138 := &system.RuleBase{Optional: true}
	ptr139 := &tests.Rectangle_rule{Base: ptr137, RuleBase: ptr138}
	ptr140 := &system.Type{Base: ptr129, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr131, "protocol": ptr134, "server": ptr136, "size": ptr139}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr141 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr142 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr143 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr144 := &tests.Kid_rule{Base: ptr143, RuleBase: (*system.RuleBase)(nil)}
	ptr145 := &system.Array_rule{Base: ptr142, RuleBase: (*system.RuleBase)(nil), Items: ptr144, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr146 := &system.Type{Base: ptr141, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr145}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr147 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr148 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr149 := &system.Int_rule{Base: ptr148, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr150 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr151 := &system.Int_rule{Base: ptr150, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr152 := &system.Type{Base: ptr147, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr149, "width": ptr151}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr153 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscale", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr154 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr155 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr156 := &tests.InstanceItem_rule{Base: ptr155, RuleBase: (*system.RuleBase)(nil)}
	ptr157 := &system.Map_rule{Base: ptr154, RuleBase: (*system.RuleBase)(nil), Items: ptr156, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr158 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr159 := &system.String_rule{Base: ptr158, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr160 := &system.Type{Base: ptr153, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"child": ptr157, "name": ptr159}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr161 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleLink", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr162 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr163 := &system.String_rule{Base: ptr162, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr164 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr165 := &system.String_rule{Base: ptr164, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr166 := &system.Type{Base: ptr161, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"href": ptr163, "rel": ptr165}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr167 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleList", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr168 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr169 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}}
	ptr170 := &tests.Rightscale_rule{Base: ptr169, RuleBase: (*system.RuleBase)(nil)}
	ptr171 := &system.Array_rule{Base: ptr168, RuleBase: (*system.RuleBase)(nil), Items: ptr170, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr172 := &system.Type{Base: ptr167, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"foo": ptr171}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr173 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr174 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr175 := &system.Number_rule{Base: ptr174, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr176 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr177 := &system.Number_rule{Base: ptr176, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr178 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr179 := &tests.C_rule{Base: ptr178, RuleBase: (*system.RuleBase)(nil)}
	ptr180 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr181 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr182 := &system.Number_rule{Base: ptr181, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr183 := &system.Map_rule{Base: ptr180, RuleBase: (*system.RuleBase)(nil), Items: ptr182, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr184 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr185 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr186 := &system.Number_rule{Base: ptr185, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr187 := &system.Map_rule{Base: ptr184, RuleBase: (*system.RuleBase)(nil), Items: ptr186, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr188 := &system.Type{Base: ptr173, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr175, "b": ptr177, "c": ptr179, "d": ptr183, "e": ptr187}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr189 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr190 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr191 := &tests.Image_rule{Base: ptr190, RuleBase: (*system.RuleBase)(nil)}
	ptr192 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr193 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr194 := &system.String_rule{Base: ptr193, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr195 := &system.Array_rule{Base: ptr192, RuleBase: (*system.RuleBase)(nil), Items: ptr194, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr196 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr197 := &system.String_rule{Base: ptr196, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr198 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr199 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr200 := &tests.Kid_rule{Base: ptr199, RuleBase: (*system.RuleBase)(nil)}
	ptr201 := &system.Map_rule{Base: ptr198, RuleBase: (*system.RuleBase)(nil), Items: ptr200, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr202 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr203 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr204 := &system.String_rule{Base: ptr203, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr205 := &system.Map_rule{Base: ptr202, RuleBase: (*system.RuleBase)(nil), Items: ptr204, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr206 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr207 := &system.Number_rule{Base: ptr206, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr208 := &system.Type{Base: ptr189, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"avatar": ptr191, "drinkPreference": ptr195, "favoriteColor": ptr197, "kids": ptr201, "name": ptr205, "weight": ptr207}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors/tests", "@basic", ptr1, 0xe2f725e9d48b69d4)
	system.RegisterType("kego.io/selectors/tests", "@c", ptr3, 0x6a4bf0b03046c68a)
	system.RegisterType("kego.io/selectors/tests", "@collision", ptr5, 0x2f6382cf84b27ee7)
	system.RegisterType("kego.io/selectors/tests", "@diagram", ptr7, 0x2076b0eaf334855b)
	system.RegisterType("kego.io/selectors/tests", "@expr", ptr9, 0x6214e678b1df35e3)
	system.RegisterType("kego.io/selectors/tests", "@gallery", ptr11, 0xa2261dbb985b3d3)
	system.RegisterType("kego.io/selectors/tests", "@image", ptr13, 0x2a2a6f416ac31013)
	system.RegisterType("kego.io/selectors/tests", "@instance", ptr15, 0x91954502a09aa42a)
	system.RegisterType("kego.io/selectors/tests", "@instanceItem", ptr17, 0x2dca71ec3918c621)
	system.RegisterType("kego.io/selectors/tests", "@kid", ptr19, 0x3b576fd38b04a2eb)
	system.RegisterType("kego.io/selectors/tests", "@photo", ptr21, 0x9b14a4b75f8931a8)
	system.RegisterType("kego.io/selectors/tests", "@polykids", ptr23, 0xebaffecb552df4d1)
	system.RegisterType("kego.io/selectors/tests", "@rectangle", ptr25, 0x21539507256190ab)
	system.RegisterType("kego.io/selectors/tests", "@rightscale", ptr27, 0x19db6e12b53f9744)
	system.RegisterType("kego.io/selectors/tests", "@rightscaleLink", ptr29, 0x127d9489c45697ea)
	system.RegisterType("kego.io/selectors/tests", "@rightscaleList", ptr31, 0x6a9527067852e20b)
	system.RegisterType("kego.io/selectors/tests", "@sibling", ptr33, 0xac837d672ff74ed6)
	system.RegisterType("kego.io/selectors/tests", "@typed", ptr35, 0x216d0e1d4c22bb2d)
	system.RegisterType("kego.io/selectors/tests", "basic", ptr59, 0xe2f725e9d48b69d4)
	system.RegisterType("kego.io/selectors/tests", "c", ptr69, 0x6a4bf0b03046c68a)
	system.RegisterType("kego.io/selectors/tests", "collision", ptr75, 0x2f6382cf84b27ee7)
	system.RegisterType("kego.io/selectors/tests", "diagram", ptr79, 0x2076b0eaf334855b)
	system.RegisterType("kego.io/selectors/tests", "expr", ptr95, 0x6214e678b1df35e3)
	system.RegisterType("kego.io/selectors/tests", "gallery", ptr101, 0xa2261dbb985b3d3)
	system.RegisterType("kego.io/selectors/tests", "image", ptr103, 0x2a2a6f416ac31013)
	system.RegisterType("kego.io/selectors/tests", "instance", ptr115, 0x91954502a09aa42a)
	system.RegisterType("kego.io/selectors/tests", "instanceItem", ptr119, 0x2dca71ec3918c621)
	system.RegisterType("kego.io/selectors/tests", "kid", ptr128, 0x3b576fd38b04a2eb)
	system.RegisterType("kego.io/selectors/tests", "photo", ptr140, 0x9b14a4b75f8931a8)
	system.RegisterType("kego.io/selectors/tests", "polykids", ptr146, 0xebaffecb552df4d1)
	system.RegisterType("kego.io/selectors/tests", "rectangle", ptr152, 0x21539507256190ab)
	system.RegisterType("kego.io/selectors/tests", "rightscale", ptr160, 0x19db6e12b53f9744)
	system.RegisterType("kego.io/selectors/tests", "rightscaleLink", ptr166, 0x127d9489c45697ea)
	system.RegisterType("kego.io/selectors/tests", "rightscaleList", ptr172, 0x6a9527067852e20b)
	system.RegisterType("kego.io/selectors/tests", "sibling", ptr188, 0xac837d672ff74ed6)
	system.RegisterType("kego.io/selectors/tests", "typed", ptr208, 0x216d0e1d4c22bb2d)
}
