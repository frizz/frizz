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
	ptr136 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr137 := &system.Rule{}
	ptr138 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr139 := &system.Rule{}
	ptr140 := &tests.InstanceItemRule{Object: ptr138, Rule: ptr139}
	ptr141 := &system.MapRule{Object: ptr136, Rule: ptr137, Items: ptr140, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr142 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr143 := &system.Rule{}
	ptr144 := &system.StringRule{Object: ptr142, Rule: ptr143, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr145 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr146 := &system.Rule{}
	ptr147 := &system.StringRule{Object: ptr145, Rule: ptr146, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr148 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr149 := &system.Rule{}
	ptr150 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr151 := &system.Rule{}
	ptr152 := &tests.RightscaleLinkRule{Object: ptr150, Rule: ptr151}
	ptr153 := &system.ArrayRule{Object: ptr148, Rule: ptr149, Items: ptr152, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr154 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr155 := &system.Rule{}
	ptr156 := &system.StringRule{Object: ptr154, Rule: ptr155, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr157 := &system.Type{Object: ptr135, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"child": ptr141, "cloud_type": ptr144, "display_name": ptr147, "links": ptr153, "name": ptr156}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr158 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "instanceItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr159 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr160 := &system.Rule{}
	ptr161 := &system.StringRule{Object: ptr159, Rule: ptr160, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr162 := &system.Type{Object: ptr158, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"name": ptr161}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr163 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "kid", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr164 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr165 := &system.Rule{}
	ptr166 := &system.StringRule{Object: ptr164, Rule: ptr165, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr167 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr168 := &system.Rule{}
	ptr169 := &system.StringRule{Object: ptr167, Rule: ptr168, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr170 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr171 := &system.Rule{Optional: true}
	ptr172 := &system.BoolRule{Object: ptr170, Rule: ptr171, Default: system.Bool{}}
	ptr173 := &system.Type{Object: ptr163, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"language": ptr166, "level": ptr169, "preferred": ptr172}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr174 := &system.Object{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr175 := &system.Object{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr176 := &system.Rule{}
	ptr177 := &system.StringRule{Object: ptr175, Rule: ptr176, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr178 := &system.Object{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr179 := &system.Rule{Optional: true}
	ptr180 := &system.StringRule{Object: ptr178, Rule: ptr179, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr181 := &system.Object{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr182 := &system.Rule{}
	ptr183 := &system.StringRule{Object: ptr181, Rule: ptr182, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr184 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rectangle", Exists: true}}
	ptr185 := &system.Rule{Optional: true}
	ptr186 := &tests.RectangleRule{Object: ptr184, Rule: ptr185}
	ptr187 := &system.Type{Object: ptr174, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"path": ptr177, "protocol": ptr180, "server": ptr183, "size": ptr186}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors/tests", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr188 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "polykids", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr189 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr190 := &system.Rule{}
	ptr191 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr192 := &system.Rule{}
	ptr193 := &tests.KidRule{Object: ptr191, Rule: ptr192}
	ptr194 := &system.ArrayRule{Object: ptr189, Rule: ptr190, Items: ptr193, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr195 := &system.Type{Object: ptr188, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr194}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr196 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rectangle", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr197 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr198 := &system.Rule{}
	ptr199 := &system.IntRule{Object: ptr197, Rule: ptr198, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr200 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr201 := &system.Rule{}
	ptr202 := &system.IntRule{Object: ptr200, Rule: ptr201, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr203 := &system.Type{Object: ptr196, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"height": ptr199, "width": ptr202}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr204 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscale", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr205 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr206 := &system.Rule{}
	ptr207 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@instanceItem", Exists: true}}
	ptr208 := &system.Rule{}
	ptr209 := &tests.InstanceItemRule{Object: ptr207, Rule: ptr208}
	ptr210 := &system.MapRule{Object: ptr205, Rule: ptr206, Items: ptr209, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr211 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr212 := &system.Rule{}
	ptr213 := &system.StringRule{Object: ptr211, Rule: ptr212, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr214 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr215 := &system.Rule{}
	ptr216 := &system.StringRule{Object: ptr214, Rule: ptr215, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr217 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr218 := &system.Rule{}
	ptr219 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscaleLink", Exists: true}}
	ptr220 := &system.Rule{}
	ptr221 := &tests.RightscaleLinkRule{Object: ptr219, Rule: ptr220}
	ptr222 := &system.ArrayRule{Object: ptr217, Rule: ptr218, Items: ptr221, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr223 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr224 := &system.Rule{}
	ptr225 := &system.StringRule{Object: ptr223, Rule: ptr224, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr226 := &system.Type{Object: ptr204, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"child": ptr210, "cloud_type": ptr213, "display_name": ptr216, "links": ptr222, "name": ptr225}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr227 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleLink", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr228 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr229 := &system.Rule{}
	ptr230 := &system.StringRule{Object: ptr228, Rule: ptr229, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr231 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr232 := &system.Rule{}
	ptr233 := &system.StringRule{Object: ptr231, Rule: ptr232, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr234 := &system.Type{Object: ptr227, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"href": ptr230, "rel": ptr233}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr235 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "rightscaleList", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr236 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr237 := &system.Rule{}
	ptr238 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@rightscale", Exists: true}}
	ptr239 := &system.Rule{}
	ptr240 := &tests.RightscaleRule{Object: ptr238, Rule: ptr239}
	ptr241 := &system.ArrayRule{Object: ptr236, Rule: ptr237, Items: ptr240, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr242 := &system.Type{Object: ptr235, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"foo": ptr241}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr243 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "sibling", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr244 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr245 := &system.Rule{}
	ptr246 := &system.NumberRule{Object: ptr244, Rule: ptr245, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr247 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr248 := &system.Rule{}
	ptr249 := &system.NumberRule{Object: ptr247, Rule: ptr248, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr250 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@c", Exists: true}}
	ptr251 := &system.Rule{}
	ptr252 := &tests.CRule{Object: ptr250, Rule: ptr251}
	ptr253 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr254 := &system.Rule{}
	ptr255 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr256 := &system.Rule{}
	ptr257 := &system.NumberRule{Object: ptr255, Rule: ptr256, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr258 := &system.MapRule{Object: ptr253, Rule: ptr254, Items: ptr257, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr259 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr260 := &system.Rule{}
	ptr261 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr262 := &system.Rule{}
	ptr263 := &system.NumberRule{Object: ptr261, Rule: ptr262, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr264 := &system.MapRule{Object: ptr259, Rule: ptr260, Items: ptr263, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr265 := &system.Type{Object: ptr243, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr246, "b": ptr249, "c": ptr252, "d": ptr258, "e": ptr264}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr266 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simple", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr267 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@simpleItem", Exists: true}}
	ptr268 := &system.Rule{}
	ptr269 := &tests.SimpleItemRule{Object: ptr267, Rule: ptr268}
	ptr270 := &system.Type{Object: ptr266, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"a": ptr269}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr271 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "simpleItem", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr272 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr273 := &system.Rule{}
	ptr274 := &system.StringRule{Object: ptr272, Rule: ptr273, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr275 := &system.Type{Object: ptr271, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"b": ptr274}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr276 := &system.Object{Id: system.Reference{Package: "kego.io/selectors/tests", Name: "typed", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr277 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@image", Exists: true}}
	ptr278 := &system.Rule{}
	ptr279 := &tests.ImageRule{Object: ptr277, Rule: ptr278}
	ptr280 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr281 := &system.Rule{}
	ptr282 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr283 := &system.Rule{}
	ptr284 := &system.StringRule{Object: ptr282, Rule: ptr283, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr285 := &system.ArrayRule{Object: ptr280, Rule: ptr281, Items: ptr284, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr286 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr287 := &system.Rule{}
	ptr288 := &system.StringRule{Object: ptr286, Rule: ptr287, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr289 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr290 := &system.Rule{}
	ptr291 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "@kid", Exists: true}}
	ptr292 := &system.Rule{}
	ptr293 := &tests.KidRule{Object: ptr291, Rule: ptr292}
	ptr294 := &system.MapRule{Object: ptr289, Rule: ptr290, Items: ptr293, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr295 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr296 := &system.Rule{}
	ptr297 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr298 := &system.Rule{}
	ptr299 := &system.StringRule{Object: ptr297, Rule: ptr298, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr300 := &system.MapRule{Object: ptr295, Rule: ptr296, Items: ptr299, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr301 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr302 := &system.Rule{}
	ptr303 := &system.NumberRule{Object: ptr301, Rule: ptr302, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr304 := &system.Type{Object: ptr276, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"avatar": ptr279, "drinkPreference": ptr285, "favoriteColor": ptr288, "kids": ptr294, "name": ptr300, "weight": ptr303}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/selectors/tests", "@basic", ptr1, 0xe2f725e9d48b69d4)
	system.Register("kego.io/selectors/tests", "@c", ptr3, 0x6a4bf0b03046c68a)
	system.Register("kego.io/selectors/tests", "@collision", ptr5, 0x2f6382cf84b27ee7)
	system.Register("kego.io/selectors/tests", "@diagram", ptr7, 0x2076b0eaf334855b)
	system.Register("kego.io/selectors/tests", "@expr", ptr9, 0x6214e678b1df35e3)
	system.Register("kego.io/selectors/tests", "@gallery", ptr11, 0xa2261dbb985b3d3)
	system.Register("kego.io/selectors/tests", "@image", ptr13, 0x2a2a6f416ac31013)
	system.Register("kego.io/selectors/tests", "@instance", ptr15, 0xa7637e9d5bf458e9)
	system.Register("kego.io/selectors/tests", "@instanceItem", ptr17, 0x2dca71ec3918c621)
	system.Register("kego.io/selectors/tests", "@kid", ptr19, 0x3b576fd38b04a2eb)
	system.Register("kego.io/selectors/tests", "@photo", ptr21, 0x6c010530199ae92f)
	system.Register("kego.io/selectors/tests", "@polykids", ptr23, 0xebaffecb552df4d1)
	system.Register("kego.io/selectors/tests", "@rectangle", ptr25, 0x21539507256190ab)
	system.Register("kego.io/selectors/tests", "@rightscale", ptr27, 0xde456fd6916c6f51)
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
	system.Register("kego.io/selectors/tests", "instance", ptr157, 0xa7637e9d5bf458e9)
	system.Register("kego.io/selectors/tests", "instanceItem", ptr162, 0x2dca71ec3918c621)
	system.Register("kego.io/selectors/tests", "kid", ptr173, 0x3b576fd38b04a2eb)
	system.Register("kego.io/selectors/tests", "photo", ptr187, 0x6c010530199ae92f)
	system.Register("kego.io/selectors/tests", "polykids", ptr195, 0xebaffecb552df4d1)
	system.Register("kego.io/selectors/tests", "rectangle", ptr203, 0x21539507256190ab)
	system.Register("kego.io/selectors/tests", "rightscale", ptr226, 0xde456fd6916c6f51)
	system.Register("kego.io/selectors/tests", "rightscaleLink", ptr234, 0x127d9489c45697ea)
	system.Register("kego.io/selectors/tests", "rightscaleList", ptr242, 0x6a9527067852e20b)
	system.Register("kego.io/selectors/tests", "sibling", ptr265, 0xac837d672ff74ed6)
	system.Register("kego.io/selectors/tests", "simple", ptr270, 0x7054dc6469f4db46)
	system.Register("kego.io/selectors/tests", "simpleItem", ptr275, 0xac167d73687ccb1f)
	system.Register("kego.io/selectors/tests", "typed", ptr304, 0x216d0e1d4c22bb2d)
}
