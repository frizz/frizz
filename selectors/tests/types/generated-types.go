package types

import (
	"kego.io/selectors/tests"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object_base{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object_base{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object_base: ptr2, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object_base{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object_base: ptr4, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object_base{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Object_base: ptr6, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object_base{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object_base: ptr8, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object_base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object_base: ptr10, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object_base{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object_base: ptr12, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object_base{Description: "Automatically created basic rule for instance", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instance", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Object_base: ptr14, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Object_base{Description: "Automatically created basic rule for instanceItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Object_base: ptr16, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Object_base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Type{Object_base: ptr18, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr20 := &system.Object_base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr21 := &system.Type{Object_base: ptr20, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Object_base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Type{Object_base: ptr22, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Object_base{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Type{Object_base: ptr24, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr26 := &system.Object_base{Description: "Automatically created basic rule for rightscale", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr27 := &system.Type{Object_base: ptr26, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr28 := &system.Object_base{Description: "Automatically created basic rule for rightscaleLink", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr29 := &system.Type{Object_base: ptr28, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr30 := &system.Object_base{Description: "Automatically created basic rule for rightscaleList", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleList", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr31 := &system.Type{Object_base: ptr30, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr32 := &system.Object_base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr33 := &system.Type{Object_base: ptr32, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34 := &system.Object_base{Description: "Automatically created basic rule for simple", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr35 := &system.Type{Object_base: ptr34, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr36 := &system.Object_base{Description: "Automatically created basic rule for simpleItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr37 := &system.Type{Object_base: ptr36, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr38 := &system.Object_base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr39 := &system.Type{Object_base: ptr38, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr40 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr41 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr42 := &system.Rule_base{}
	ptr43 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr44 := &system.Rule_base{}
	ptr45 := &system.String_rule{Object_base: ptr43, Rule_base: ptr44, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr46 := &system.Array_rule{Object_base: ptr41, Rule_base: ptr42, Items: ptr45, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr47 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr48 := &system.Rule_base{}
	ptr49 := &system.String_rule{Object_base: ptr47, Rule_base: ptr48, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr50 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr51 := &system.Rule_base{}
	ptr52 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr53 := &system.Rule_base{}
	ptr54 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr55 := &system.Rule_base{}
	ptr56 := &system.String_rule{Object_base: ptr54, Rule_base: ptr55, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr57 := &system.Map_rule{Object_base: ptr52, Rule_base: ptr53, Items: ptr56, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr58 := &system.Array_rule{Object_base: ptr50, Rule_base: ptr51, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr60 := &system.Rule_base{}
	ptr61 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr62 := &system.Rule_base{}
	ptr63 := &system.String_rule{Object_base: ptr61, Rule_base: ptr62, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr64 := &system.Map_rule{Object_base: ptr59, Rule_base: ptr60, Items: ptr63, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr65 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr66 := &system.Rule_base{}
	ptr67 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr68 := &system.Rule_base{}
	ptr69 := &system.String_rule{Object_base: ptr67, Rule_base: ptr68, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr70 := &system.Array_rule{Object_base: ptr65, Rule_base: ptr66, Items: ptr69, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr71 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr72 := &system.Rule_base{}
	ptr73 := &system.Number_rule{Object_base: ptr71, Rule_base: ptr72, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr74 := &system.Type{Object_base: ptr40, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"drinkPreference": ptr46, "favoriteColor": ptr49, "languagesSpoken": ptr58, "name": ptr64, "seatingPreference": ptr70, "weight": ptr73}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr77 := &system.Rule_base{}
	ptr78 := &system.Number_rule{Object_base: ptr76, Rule_base: ptr77, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr79 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr80 := &system.Rule_base{}
	ptr81 := &system.Number_rule{Object_base: ptr79, Rule_base: ptr80, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr82 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr83 := &system.Rule_base{}
	ptr84 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr85 := &system.Rule_base{}
	ptr86 := &system.Number_rule{Object_base: ptr84, Rule_base: ptr85, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr87 := &system.Map_rule{Object_base: ptr82, Rule_base: ptr83, Items: ptr86, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr88 := &system.Type{Object_base: ptr75, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr78, "b": ptr81, "c": ptr87}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr89 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr90 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr91 := &system.Rule_base{}
	ptr92 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr93 := &system.Rule_base{}
	ptr94 := &system.String_rule{Object_base: ptr92, Rule_base: ptr93, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr95 := &system.Map_rule{Object_base: ptr90, Rule_base: ptr91, Items: ptr94, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr96 := &system.Type{Object_base: ptr89, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"number": ptr95}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr97 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr98 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr99 := &system.Rule_base{}
	ptr100 := &system.String_rule{Object_base: ptr98, Rule_base: ptr99, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr101 := &system.Type{Object_base: ptr97, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr100}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr104 := &system.Rule_base{}
	ptr105 := &system.Bool_rule{Object_base: ptr103, Rule_base: ptr104, Default: system.Bool{}}
	ptr106 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr107 := &system.Rule_base{}
	ptr108 := &system.Number_rule{Object_base: ptr106, Rule_base: ptr107, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr109 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr110 := &system.Rule_base{}
	ptr111 := &system.Number_rule{Object_base: ptr109, Rule_base: ptr110, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr112 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr113 := &system.Rule_base{}
	ptr114 := &system.String_rule{Object_base: ptr112, Rule_base: ptr113, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr115 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr116 := &system.Rule_base{}
	ptr117 := &system.String_rule{Object_base: ptr115, Rule_base: ptr116, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr118 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr119 := &system.Rule_base{}
	ptr120 := &system.String_rule{Object_base: ptr118, Rule_base: ptr119, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr121 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr122 := &system.Rule_base{}
	ptr123 := &system.Bool_rule{Object_base: ptr121, Rule_base: ptr122, Default: system.Bool{}}
	ptr124 := &system.Type{Object_base: ptr102, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"false": ptr105, "float": ptr108, "int": ptr111, "null": ptr114, "string": ptr117, "string2": ptr120, "true": ptr123}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr125 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr126 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr127 := &system.Rule_base{}
	ptr128 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr129 := &system.Rule_base{}
	ptr130 := &tests.Image_rule{Object_base: ptr128, Rule_base: ptr129}
	ptr131 := &system.Map_rule{Object_base: ptr126, Rule_base: ptr127, Items: ptr130, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr132 := &system.Type{Object_base: ptr125, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr131}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr133 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr134 := &system.Type{Object_base: ptr133, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr135 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instance", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr136 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr137 := &system.Rule_base{}
	ptr138 := &system.String_rule{Object_base: ptr136, Rule_base: ptr137, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr139 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr140 := &system.Rule_base{}
	ptr141 := &system.String_rule{Object_base: ptr139, Rule_base: ptr140, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr142 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr143 := &system.Rule_base{}
	ptr144 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr145 := &system.Rule_base{}
	ptr146 := &tests.RightscaleLink_rule{Object_base: ptr144, Rule_base: ptr145}
	ptr147 := &system.Array_rule{Object_base: ptr142, Rule_base: ptr143, Items: ptr146, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr148 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr149 := &system.Rule_base{}
	ptr150 := &system.String_rule{Object_base: ptr148, Rule_base: ptr149, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr151 := &system.Type{Object_base: ptr135, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"cloud_type": ptr138, "display_name": ptr141, "links": ptr147, "name": ptr150}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr152 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instanceItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr153 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr154 := &system.Rule_base{}
	ptr155 := &system.String_rule{Object_base: ptr153, Rule_base: ptr154, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr156 := &system.Type{Object_base: ptr152, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr155}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr157 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr158 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr159 := &system.Rule_base{}
	ptr160 := &system.String_rule{Object_base: ptr158, Rule_base: ptr159, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr161 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr162 := &system.Rule_base{}
	ptr163 := &system.String_rule{Object_base: ptr161, Rule_base: ptr162, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr164 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr165 := &system.Rule_base{Optional: true}
	ptr166 := &system.Bool_rule{Object_base: ptr164, Rule_base: ptr165, Default: system.Bool{}}
	ptr167 := &system.Type{Object_base: ptr157, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"language": ptr160, "level": ptr163, "preferred": ptr166}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr168 := &system.Object_base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr169 := &system.Object_base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr170 := &system.Rule_base{}
	ptr171 := &system.String_rule{Object_base: ptr169, Rule_base: ptr170, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr172 := &system.Object_base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr173 := &system.Rule_base{Optional: true}
	ptr174 := &system.String_rule{Object_base: ptr172, Rule_base: ptr173, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr175 := &system.Object_base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr176 := &system.Rule_base{}
	ptr177 := &system.String_rule{Object_base: ptr175, Rule_base: ptr176, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr178 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr179 := &system.Rule_base{Optional: true}
	ptr180 := &tests.Rectangle_rule{Object_base: ptr178, Rule_base: ptr179}
	ptr181 := &system.Type{Object_base: ptr168, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr171, "protocol": ptr174, "server": ptr177, "size": ptr180}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr182 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr183 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr184 := &system.Rule_base{}
	ptr185 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr186 := &system.Rule_base{}
	ptr187 := &tests.Kid_rule{Object_base: ptr185, Rule_base: ptr186}
	ptr188 := &system.Array_rule{Object_base: ptr183, Rule_base: ptr184, Items: ptr187, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr189 := &system.Type{Object_base: ptr182, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr188}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr190 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr191 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr192 := &system.Rule_base{}
	ptr193 := &system.Int_rule{Object_base: ptr191, Rule_base: ptr192, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr194 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr195 := &system.Rule_base{}
	ptr196 := &system.Int_rule{Object_base: ptr194, Rule_base: ptr195, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr197 := &system.Type{Object_base: ptr190, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr193, "width": ptr196}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr198 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscale", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr199 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr200 := &system.Rule_base{}
	ptr201 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr202 := &system.Rule_base{}
	ptr203 := &tests.InstanceItem_rule{Object_base: ptr201, Rule_base: ptr202}
	ptr204 := &system.Map_rule{Object_base: ptr199, Rule_base: ptr200, Items: ptr203, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr205 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr206 := &system.Rule_base{}
	ptr207 := &system.String_rule{Object_base: ptr205, Rule_base: ptr206, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr208 := &system.Type{Object_base: ptr198, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"child": ptr204, "name": ptr207}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr209 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleLink", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr210 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr211 := &system.Rule_base{}
	ptr212 := &system.String_rule{Object_base: ptr210, Rule_base: ptr211, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr213 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr214 := &system.Rule_base{}
	ptr215 := &system.String_rule{Object_base: ptr213, Rule_base: ptr214, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr216 := &system.Type{Object_base: ptr209, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"href": ptr212, "rel": ptr215}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr217 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleList", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr218 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr219 := &system.Rule_base{}
	ptr220 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}}
	ptr221 := &system.Rule_base{}
	ptr222 := &tests.Rightscale_rule{Object_base: ptr220, Rule_base: ptr221}
	ptr223 := &system.Array_rule{Object_base: ptr218, Rule_base: ptr219, Items: ptr222, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr224 := &system.Type{Object_base: ptr217, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"foo": ptr223}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr225 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr226 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr227 := &system.Rule_base{}
	ptr228 := &system.Number_rule{Object_base: ptr226, Rule_base: ptr227, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr229 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr230 := &system.Rule_base{}
	ptr231 := &system.Number_rule{Object_base: ptr229, Rule_base: ptr230, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr232 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr233 := &system.Rule_base{}
	ptr234 := &tests.C_rule{Object_base: ptr232, Rule_base: ptr233}
	ptr235 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr236 := &system.Rule_base{}
	ptr237 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr238 := &system.Rule_base{}
	ptr239 := &system.Number_rule{Object_base: ptr237, Rule_base: ptr238, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr240 := &system.Map_rule{Object_base: ptr235, Rule_base: ptr236, Items: ptr239, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr241 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr242 := &system.Rule_base{}
	ptr243 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr244 := &system.Rule_base{}
	ptr245 := &system.Number_rule{Object_base: ptr243, Rule_base: ptr244, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr246 := &system.Map_rule{Object_base: ptr241, Rule_base: ptr242, Items: ptr245, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr247 := &system.Type{Object_base: ptr225, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr228, "b": ptr231, "c": ptr234, "d": ptr240, "e": ptr246}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr248 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr249 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}}
	ptr250 := &system.Rule_base{}
	ptr251 := &tests.SimpleItem_rule{Object_base: ptr249, Rule_base: ptr250}
	ptr252 := &system.Type{Object_base: ptr248, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"a": ptr251}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr253 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simpleItem", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr254 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr255 := &system.Rule_base{}
	ptr256 := &system.String_rule{Object_base: ptr254, Rule_base: ptr255, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr257 := &system.Type{Object_base: ptr253, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"b": ptr256}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr258 := &system.Object_base{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr259 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr260 := &system.Rule_base{}
	ptr261 := &tests.Image_rule{Object_base: ptr259, Rule_base: ptr260}
	ptr262 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr263 := &system.Rule_base{}
	ptr264 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr265 := &system.Rule_base{}
	ptr266 := &system.String_rule{Object_base: ptr264, Rule_base: ptr265, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr267 := &system.Array_rule{Object_base: ptr262, Rule_base: ptr263, Items: ptr266, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr268 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr269 := &system.Rule_base{}
	ptr270 := &system.String_rule{Object_base: ptr268, Rule_base: ptr269, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr271 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr272 := &system.Rule_base{}
	ptr273 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr274 := &system.Rule_base{}
	ptr275 := &tests.Kid_rule{Object_base: ptr273, Rule_base: ptr274}
	ptr276 := &system.Map_rule{Object_base: ptr271, Rule_base: ptr272, Items: ptr275, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr277 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr278 := &system.Rule_base{}
	ptr279 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr280 := &system.Rule_base{}
	ptr281 := &system.String_rule{Object_base: ptr279, Rule_base: ptr280, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr282 := &system.Map_rule{Object_base: ptr277, Rule_base: ptr278, Items: ptr281, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr283 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr284 := &system.Rule_base{}
	ptr285 := &system.Number_rule{Object_base: ptr283, Rule_base: ptr284, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr286 := &system.Type{Object_base: ptr258, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"avatar": ptr261, "drinkPreference": ptr267, "favoriteColor": ptr270, "kids": ptr276, "name": ptr282, "weight": ptr285}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
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
