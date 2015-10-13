package types

import (
	"kego.io/selectors/tests"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@basic", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@collision", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@diagram", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Object: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@expr", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@gallery", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object: ptr10, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object: ptr12, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object{Description: "Automatically created basic rule for instance", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instance", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Object: ptr14, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Object{Description: "Automatically created basic rule for instanceItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Object: ptr16, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Object{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Type{Object: ptr18, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr20 := &system.Object{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@photo", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr21 := &system.Type{Object: ptr20, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Object{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@polykids", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Type{Object: ptr22, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Object{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Type{Object: ptr24, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr26 := &system.Object{Description: "Automatically created basic rule for rightscale", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr27 := &system.Type{Object: ptr26, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr28 := &system.Object{Description: "Automatically created basic rule for rightscaleLink", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr29 := &system.Type{Object: ptr28, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr30 := &system.Object{Description: "Automatically created basic rule for rightscaleList", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleList", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr31 := &system.Type{Object: ptr30, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr32 := &system.Object{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@sibling", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr33 := &system.Type{Object: ptr32, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34 := &system.Object{Description: "Automatically created basic rule for simple", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simple", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr35 := &system.Type{Object: ptr34, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr36 := &system.Object{Description: "Automatically created basic rule for simpleItem", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr37 := &system.Type{Object: ptr36, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr38 := &system.Object{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "@typed", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr39 := &system.Type{Object: ptr38, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr40 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "basic", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr41 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr42 := &system.Rule{}
	ptr43 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr44 := &system.Rule{}
	ptr45 := &system.StringRule{Object: ptr43, Rule: ptr44, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr46 := &system.ArrayRule{Object: ptr41, Rule: ptr42, Items: ptr45, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr47 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr48 := &system.Rule{}
	ptr49 := &system.StringRule{Object: ptr47, Rule: ptr48, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr50 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr51 := &system.Rule{}
	ptr52 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr53 := &system.Rule{}
	ptr54 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr55 := &system.Rule{}
	ptr56 := &system.StringRule{Object: ptr54, Rule: ptr55, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr57 := &system.MapRule{Object: ptr52, Rule: ptr53, Items: ptr56, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr58 := &system.ArrayRule{Object: ptr50, Rule: ptr51, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr60 := &system.Rule{}
	ptr61 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr62 := &system.Rule{}
	ptr63 := &system.StringRule{Object: ptr61, Rule: ptr62, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr64 := &system.MapRule{Object: ptr59, Rule: ptr60, Items: ptr63, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr65 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr66 := &system.Rule{}
	ptr67 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr68 := &system.Rule{}
	ptr69 := &system.StringRule{Object: ptr67, Rule: ptr68, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr70 := &system.ArrayRule{Object: ptr65, Rule: ptr66, Items: ptr69, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr71 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr72 := &system.Rule{}
	ptr73 := &system.NumberRule{Object: ptr71, Rule: ptr72, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr74 := &system.Type{Object: ptr40, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"drinkPreference": ptr46, "favoriteColor": ptr49, "languagesSpoken": ptr58, "name": ptr64, "seatingPreference": ptr70, "weight": ptr73}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "c", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr77 := &system.Rule{}
	ptr78 := &system.NumberRule{Object: ptr76, Rule: ptr77, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr79 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr80 := &system.Rule{}
	ptr81 := &system.NumberRule{Object: ptr79, Rule: ptr80, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr82 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr83 := &system.Rule{}
	ptr84 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr85 := &system.Rule{}
	ptr86 := &system.NumberRule{Object: ptr84, Rule: ptr85, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr87 := &system.MapRule{Object: ptr82, Rule: ptr83, Items: ptr86, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr88 := &system.Type{Object: ptr75, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr78, "b": ptr81, "c": ptr87}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr89 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "collision", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr90 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr91 := &system.Rule{}
	ptr92 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr93 := &system.Rule{}
	ptr94 := &system.StringRule{Object: ptr92, Rule: ptr93, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr95 := &system.MapRule{Object: ptr90, Rule: ptr91, Items: ptr94, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr96 := &system.Type{Object: ptr89, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"number": ptr95}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr97 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr98 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr99 := &system.Rule{}
	ptr100 := &system.StringRule{Object: ptr98, Rule: ptr99, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr101 := &system.Type{Object: ptr97, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"url": ptr100}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "expr", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr104 := &system.Rule{}
	ptr105 := &system.BoolRule{Object: ptr103, Rule: ptr104, Default: system.Bool{}}
	ptr106 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr107 := &system.Rule{}
	ptr108 := &system.NumberRule{Object: ptr106, Rule: ptr107, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr109 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr110 := &system.Rule{}
	ptr111 := &system.NumberRule{Object: ptr109, Rule: ptr110, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr112 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr113 := &system.Rule{}
	ptr114 := &system.StringRule{Object: ptr112, Rule: ptr113, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr115 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr116 := &system.Rule{}
	ptr117 := &system.StringRule{Object: ptr115, Rule: ptr116, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr118 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr119 := &system.Rule{}
	ptr120 := &system.StringRule{Object: ptr118, Rule: ptr119, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr121 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr122 := &system.Rule{}
	ptr123 := &system.BoolRule{Object: ptr121, Rule: ptr122, Default: system.Bool{}}
	ptr124 := &system.Type{Object: ptr102, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"false": ptr105, "float": ptr108, "int": ptr111, "null": ptr114, "string": ptr117, "string2": ptr120, "true": ptr123}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr125 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr126 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr127 := &system.Rule{}
	ptr128 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr129 := &system.Rule{}
	ptr130 := &tests.ImageRule{Object: ptr128, Rule: ptr129}
	ptr131 := &system.MapRule{Object: ptr126, Rule: ptr127, Items: ptr130, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr132 := &system.Type{Object: ptr125, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr131}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr133 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr134 := &system.Type{Object: ptr133, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr135 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instance", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr136 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr137 := &system.Rule{}
	ptr138 := &system.StringRule{Object: ptr136, Rule: ptr137, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr139 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr140 := &system.Rule{}
	ptr141 := &system.StringRule{Object: ptr139, Rule: ptr140, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr142 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr143 := &system.Rule{}
	ptr144 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr145 := &system.Rule{}
	ptr146 := &tests.RightscaleLinkRule{Object: ptr144, Rule: ptr145}
	ptr147 := &system.ArrayRule{Object: ptr142, Rule: ptr143, Items: ptr146, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr148 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr149 := &system.Rule{}
	ptr150 := &system.StringRule{Object: ptr148, Rule: ptr149, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr151 := &system.Type{Object: ptr135, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"cloud_type": ptr138, "display_name": ptr141, "links": ptr147, "name": ptr150}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr152 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instanceItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr153 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr154 := &system.Rule{}
	ptr155 := &system.StringRule{Object: ptr153, Rule: ptr154, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr156 := &system.Type{Object: ptr152, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"name": ptr155}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr157 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr158 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr159 := &system.Rule{}
	ptr160 := &system.StringRule{Object: ptr158, Rule: ptr159, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr161 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr162 := &system.Rule{}
	ptr163 := &system.StringRule{Object: ptr161, Rule: ptr162, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr164 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr165 := &system.Rule{Optional: true}
	ptr166 := &system.BoolRule{Object: ptr164, Rule: ptr165, Default: system.Bool{}}
	ptr167 := &system.Type{Object: ptr157, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"language": ptr160, "level": ptr163, "preferred": ptr166}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr168 := &system.Object{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr169 := &system.Object{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr170 := &system.Rule{}
	ptr171 := &system.StringRule{Object: ptr169, Rule: ptr170, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr172 := &system.Object{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr173 := &system.Rule{Optional: true}
	ptr174 := &system.StringRule{Object: ptr172, Rule: ptr173, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr175 := &system.Object{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr176 := &system.Rule{}
	ptr177 := &system.StringRule{Object: ptr175, Rule: ptr176, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr178 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr179 := &system.Rule{Optional: true}
	ptr180 := &tests.RectangleRule{Object: ptr178, Rule: ptr179}
	ptr181 := &system.Type{Object: ptr168, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"path": ptr171, "protocol": ptr174, "server": ptr177, "size": ptr180}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr182 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr183 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr184 := &system.Rule{}
	ptr185 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr186 := &system.Rule{}
	ptr187 := &tests.KidRule{Object: ptr185, Rule: ptr186}
	ptr188 := &system.ArrayRule{Object: ptr183, Rule: ptr184, Items: ptr187, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr189 := &system.Type{Object: ptr182, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr188}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr190 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr191 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr192 := &system.Rule{}
	ptr193 := &system.IntRule{Object: ptr191, Rule: ptr192, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr194 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr195 := &system.Rule{}
	ptr196 := &system.IntRule{Object: ptr194, Rule: ptr195, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr197 := &system.Type{Object: ptr190, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"height": ptr193, "width": ptr196}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr198 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscale", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr199 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr200 := &system.Rule{}
	ptr201 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr202 := &system.Rule{}
	ptr203 := &tests.InstanceItemRule{Object: ptr201, Rule: ptr202}
	ptr204 := &system.MapRule{Object: ptr199, Rule: ptr200, Items: ptr203, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr205 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr206 := &system.Rule{}
	ptr207 := &system.StringRule{Object: ptr205, Rule: ptr206, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr208 := &system.Type{Object: ptr198, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"child": ptr204, "name": ptr207}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr209 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleLink", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr210 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr211 := &system.Rule{}
	ptr212 := &system.StringRule{Object: ptr210, Rule: ptr211, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr213 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr214 := &system.Rule{}
	ptr215 := &system.StringRule{Object: ptr213, Rule: ptr214, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr216 := &system.Type{Object: ptr209, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"href": ptr212, "rel": ptr215}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr217 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleList", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr218 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr219 := &system.Rule{}
	ptr220 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}}
	ptr221 := &system.Rule{}
	ptr222 := &tests.RightscaleRule{Object: ptr220, Rule: ptr221}
	ptr223 := &system.ArrayRule{Object: ptr218, Rule: ptr219, Items: ptr222, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr224 := &system.Type{Object: ptr217, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"foo": ptr223}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr225 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr226 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr227 := &system.Rule{}
	ptr228 := &system.NumberRule{Object: ptr226, Rule: ptr227, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr229 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr230 := &system.Rule{}
	ptr231 := &system.NumberRule{Object: ptr229, Rule: ptr230, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr232 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr233 := &system.Rule{}
	ptr234 := &tests.CRule{Object: ptr232, Rule: ptr233}
	ptr235 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr236 := &system.Rule{}
	ptr237 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr238 := &system.Rule{}
	ptr239 := &system.NumberRule{Object: ptr237, Rule: ptr238, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr240 := &system.MapRule{Object: ptr235, Rule: ptr236, Items: ptr239, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr241 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr242 := &system.Rule{}
	ptr243 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr244 := &system.Rule{}
	ptr245 := &system.NumberRule{Object: ptr243, Rule: ptr244, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr246 := &system.MapRule{Object: ptr241, Rule: ptr242, Items: ptr245, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr247 := &system.Type{Object: ptr225, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr228, "b": ptr231, "c": ptr234, "d": ptr240, "e": ptr246}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr248 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simple", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr249 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}}
	ptr250 := &system.Rule{}
	ptr251 := &tests.SimpleItemRule{Object: ptr249, Rule: ptr250}
	ptr252 := &system.Type{Object: ptr248, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr251}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr253 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simpleItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr254 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr255 := &system.Rule{}
	ptr256 := &system.StringRule{Object: ptr254, Rule: ptr255, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr257 := &system.Type{Object: ptr253, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"b": ptr256}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr258 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr259 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr260 := &system.Rule{}
	ptr261 := &tests.ImageRule{Object: ptr259, Rule: ptr260}
	ptr262 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr263 := &system.Rule{}
	ptr264 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr265 := &system.Rule{}
	ptr266 := &system.StringRule{Object: ptr264, Rule: ptr265, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr267 := &system.ArrayRule{Object: ptr262, Rule: ptr263, Items: ptr266, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr268 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr269 := &system.Rule{}
	ptr270 := &system.StringRule{Object: ptr268, Rule: ptr269, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr271 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr272 := &system.Rule{}
	ptr273 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr274 := &system.Rule{}
	ptr275 := &tests.KidRule{Object: ptr273, Rule: ptr274}
	ptr276 := &system.MapRule{Object: ptr271, Rule: ptr272, Items: ptr275, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr277 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr278 := &system.Rule{}
	ptr279 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr280 := &system.Rule{}
	ptr281 := &system.StringRule{Object: ptr279, Rule: ptr280, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr282 := &system.MapRule{Object: ptr277, Rule: ptr278, Items: ptr281, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr283 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr284 := &system.Rule{}
	ptr285 := &system.NumberRule{Object: ptr283, Rule: ptr284, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr286 := &system.Type{Object: ptr258, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"avatar": ptr261, "drinkPreference": ptr267, "favoriteColor": ptr270, "kids": ptr276, "name": ptr282, "weight": ptr285}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
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
