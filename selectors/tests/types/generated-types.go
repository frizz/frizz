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
	ptr14 := &system.Base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Base: ptr14, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Base: ptr16, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Type{Base: ptr18, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr20 := &system.Base{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr21 := &system.Type{Base: ptr20, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Type{Base: ptr22, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Type{Base: ptr24, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr26 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr27 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr28 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr29 := &system.String_rule{Base: ptr28, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr30 := &system.Array_rule{Base: ptr27, RuleBase: (*system.RuleBase)(nil), Items: ptr29, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr31 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr32 := &system.String_rule{Base: ptr31, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr33 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr34 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr35 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr36 := &system.String_rule{Base: ptr35, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr37 := &system.Map_rule{Base: ptr34, RuleBase: (*system.RuleBase)(nil), Items: ptr36, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr38 := &system.Array_rule{Base: ptr33, RuleBase: (*system.RuleBase)(nil), Items: ptr37, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr39 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr40 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr41 := &system.String_rule{Base: ptr40, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr42 := &system.Map_rule{Base: ptr39, RuleBase: (*system.RuleBase)(nil), Items: ptr41, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr43 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr44 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr45 := &system.String_rule{Base: ptr44, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr46 := &system.Array_rule{Base: ptr43, RuleBase: (*system.RuleBase)(nil), Items: ptr45, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr47 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr48 := &system.Number_rule{Base: ptr47, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr49 := &system.Type{Base: ptr26, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"drinkPreference": ptr30, "favoriteColor": ptr32, "languagesSpoken": ptr38, "name": ptr42, "seatingPreference": ptr46, "weight": ptr48}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr50 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr51 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr52 := &system.Number_rule{Base: ptr51, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr53 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr54 := &system.Number_rule{Base: ptr53, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr55 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr56 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr57 := &system.Number_rule{Base: ptr56, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr58 := &system.Map_rule{Base: ptr55, RuleBase: (*system.RuleBase)(nil), Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Type{Base: ptr50, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr52, "b": ptr54, "c": ptr58}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr60 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr61 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr62 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr63 := &system.String_rule{Base: ptr62, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr64 := &system.Map_rule{Base: ptr61, RuleBase: (*system.RuleBase)(nil), Items: ptr63, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr65 := &system.Type{Base: ptr60, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr64}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr66 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr67 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr68 := &system.String_rule{Base: ptr67, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr69 := &system.Type{Base: ptr66, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr68}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr70 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr71 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr72 := &system.Bool_rule{Base: ptr71, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr73 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr74 := &system.Number_rule{Base: ptr73, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr75 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr76 := &system.Number_rule{Base: ptr75, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr77 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr78 := &system.String_rule{Base: ptr77, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr79 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr80 := &system.String_rule{Base: ptr79, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr81 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr82 := &system.String_rule{Base: ptr81, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr83 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr84 := &system.Bool_rule{Base: ptr83, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}
	ptr85 := &system.Type{Base: ptr70, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"false": ptr72, "float": ptr74, "int": ptr76, "null": ptr78, "string": ptr80, "string2": ptr82, "true": ptr84}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr86 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr87 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr88 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr89 := &tests.Image_rule{Base: ptr88, RuleBase: (*system.RuleBase)(nil)}
	ptr90 := &system.Map_rule{Base: ptr87, RuleBase: (*system.RuleBase)(nil), Items: ptr89, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr91 := &system.Type{Base: ptr86, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr90}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr92 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr93 := &system.Type{Base: ptr92, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr94 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr95 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr96 := &system.String_rule{Base: ptr95, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr97 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr98 := &system.String_rule{Base: ptr97, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr99 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr100 := &system.RuleBase{Optional: true}
	ptr101 := &system.Bool_rule{Base: ptr99, RuleBase: ptr100, Default: system.Bool{}}
	ptr102 := &system.Type{Base: ptr94, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr96, "level": ptr98, "preferred": ptr101}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr103 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr104 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr105 := &system.String_rule{Base: ptr104, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr106 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr107 := &system.RuleBase{Optional: true}
	ptr108 := &system.String_rule{Base: ptr106, RuleBase: ptr107, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr109 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr110 := &system.String_rule{Base: ptr109, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr111 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr112 := &system.RuleBase{Optional: true}
	ptr113 := &tests.Rectangle_rule{Base: ptr111, RuleBase: ptr112}
	ptr114 := &system.Type{Base: ptr103, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr105, "protocol": ptr108, "server": ptr110, "size": ptr113}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr115 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr116 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr117 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr118 := &tests.Kid_rule{Base: ptr117, RuleBase: (*system.RuleBase)(nil)}
	ptr119 := &system.Array_rule{Base: ptr116, RuleBase: (*system.RuleBase)(nil), Items: ptr118, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr120 := &system.Type{Base: ptr115, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr119}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr121 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr122 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr123 := &system.Int_rule{Base: ptr122, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr124 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr125 := &system.Int_rule{Base: ptr124, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr126 := &system.Type{Base: ptr121, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr123, "width": ptr125}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr127 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr128 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr129 := &system.Number_rule{Base: ptr128, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr130 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr131 := &system.Number_rule{Base: ptr130, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr132 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr133 := &tests.C_rule{Base: ptr132, RuleBase: (*system.RuleBase)(nil)}
	ptr134 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr135 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr136 := &system.Number_rule{Base: ptr135, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr137 := &system.Map_rule{Base: ptr134, RuleBase: (*system.RuleBase)(nil), Items: ptr136, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr138 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr139 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr140 := &system.Number_rule{Base: ptr139, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr141 := &system.Map_rule{Base: ptr138, RuleBase: (*system.RuleBase)(nil), Items: ptr140, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr142 := &system.Type{Base: ptr127, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr129, "b": ptr131, "c": ptr133, "d": ptr137, "e": ptr141}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr143 := &system.Base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr144 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr145 := &tests.Image_rule{Base: ptr144, RuleBase: (*system.RuleBase)(nil)}
	ptr146 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr147 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr148 := &system.String_rule{Base: ptr147, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr149 := &system.Array_rule{Base: ptr146, RuleBase: (*system.RuleBase)(nil), Items: ptr148, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr150 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr151 := &system.String_rule{Base: ptr150, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr152 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr153 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr154 := &tests.Kid_rule{Base: ptr153, RuleBase: (*system.RuleBase)(nil)}
	ptr155 := &system.Map_rule{Base: ptr152, RuleBase: (*system.RuleBase)(nil), Items: ptr154, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr156 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr157 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr158 := &system.String_rule{Base: ptr157, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr159 := &system.Map_rule{Base: ptr156, RuleBase: (*system.RuleBase)(nil), Items: ptr158, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr160 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr161 := &system.Number_rule{Base: ptr160, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr162 := &system.Type{Base: ptr143, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"avatar": ptr145, "drinkPreference": ptr149, "favoriteColor": ptr151, "kids": ptr155, "name": ptr159, "weight": ptr161}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/selectors/tests", "@basic", ptr1, 0xe2f725e9d48b69d4)
	system.RegisterType("kego.io/selectors/tests", "@c", ptr3, 0x6a4bf0b03046c68a)
	system.RegisterType("kego.io/selectors/tests", "@collision", ptr5, 0x2f6382cf84b27ee7)
	system.RegisterType("kego.io/selectors/tests", "@diagram", ptr7, 0x2076b0eaf334855b)
	system.RegisterType("kego.io/selectors/tests", "@expr", ptr9, 0x6214e678b1df35e3)
	system.RegisterType("kego.io/selectors/tests", "@gallery", ptr11, 0xa2261dbb985b3d3)
	system.RegisterType("kego.io/selectors/tests", "@image", ptr13, 0x2a2a6f416ac31013)
	system.RegisterType("kego.io/selectors/tests", "@kid", ptr15, 0x3b576fd38b04a2eb)
	system.RegisterType("kego.io/selectors/tests", "@photo", ptr17, 0x9b14a4b75f8931a8)
	system.RegisterType("kego.io/selectors/tests", "@polykids", ptr19, 0xebaffecb552df4d1)
	system.RegisterType("kego.io/selectors/tests", "@rectangle", ptr21, 0x21539507256190ab)
	system.RegisterType("kego.io/selectors/tests", "@sibling", ptr23, 0xac837d672ff74ed6)
	system.RegisterType("kego.io/selectors/tests", "@typed", ptr25, 0x216d0e1d4c22bb2d)
	system.RegisterType("kego.io/selectors/tests", "basic", ptr49, 0xe2f725e9d48b69d4)
	system.RegisterType("kego.io/selectors/tests", "c", ptr59, 0x6a4bf0b03046c68a)
	system.RegisterType("kego.io/selectors/tests", "collision", ptr65, 0x2f6382cf84b27ee7)
	system.RegisterType("kego.io/selectors/tests", "diagram", ptr69, 0x2076b0eaf334855b)
	system.RegisterType("kego.io/selectors/tests", "expr", ptr85, 0x6214e678b1df35e3)
	system.RegisterType("kego.io/selectors/tests", "gallery", ptr91, 0xa2261dbb985b3d3)
	system.RegisterType("kego.io/selectors/tests", "image", ptr93, 0x2a2a6f416ac31013)
	system.RegisterType("kego.io/selectors/tests", "kid", ptr102, 0x3b576fd38b04a2eb)
	system.RegisterType("kego.io/selectors/tests", "photo", ptr114, 0x9b14a4b75f8931a8)
	system.RegisterType("kego.io/selectors/tests", "polykids", ptr120, 0xebaffecb552df4d1)
	system.RegisterType("kego.io/selectors/tests", "rectangle", ptr126, 0x21539507256190ab)
	system.RegisterType("kego.io/selectors/tests", "sibling", ptr142, 0xac837d672ff74ed6)
	system.RegisterType("kego.io/selectors/tests", "typed", ptr162, 0x216d0e1d4c22bb2d)
}
