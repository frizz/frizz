package types

import (
	"kego.io/demo/common/images"
	_ "kego.io/demo/common/images/types"
	_ "kego.io/demo/common/units/types"
	"kego.io/demo/common/words"
	_ "kego.io/demo/common/words/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for gallery1", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Automatically created basic rule for gallery1a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Base{Description: "Automatically created basic rule for gallery2", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Base: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Base{Description: "Automatically created basic rule for gallery2a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Base: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Base{Description: "Automatically created basic rule for gallery2b", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2b", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Base: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Base{Description: "Automatically created basic rule for gallery3", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Base: ptr10, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Base{Description: "Automatically created basic rule for gallery3a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Base: ptr12, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.RuleBase{}
	ptr17 := &system.String_rule{Base: ptr15, RuleBase: ptr16, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr18 := &system.Type{Base: ptr14, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr17}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.RuleBase{}
	ptr22 := &system.String_rule{Base: ptr20, RuleBase: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 20, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Type{Base: ptr19, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr22}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr26 := &system.RuleBase{}
	ptr27 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr28 := &system.RuleBase{}
	ptr29 := &images.Photo_rule{Base: ptr27, RuleBase: ptr28}
	ptr30 := &system.Map_rule{Base: ptr25, RuleBase: ptr26, Items: ptr29, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr31 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr32 := &system.RuleBase{}
	ptr33 := &system.String_rule{Base: ptr31, RuleBase: ptr32, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34 := &system.Type{Base: ptr24, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr30, "title": ptr33}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr35 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr36 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr37 := &system.RuleBase{}
	ptr38 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr39 := &system.RuleBase{Selector: ".protocol"}
	ptr40 := &system.String_rule{Base: ptr38, RuleBase: ptr39, Default: system.String{}, Enum: []string(nil), Equal: system.String{Value: "https", Exists: true}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr41 := &system.Base{Id: system.Reference{}, Rules: []system.Rule{ptr40}, Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr42 := &system.RuleBase{}
	ptr43 := &images.Photo_rule{Base: ptr41, RuleBase: ptr42}
	ptr44 := &system.Map_rule{Base: ptr36, RuleBase: ptr37, Items: ptr43, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr45 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr46 := &system.RuleBase{}
	ptr47 := &system.String_rule{Base: ptr45, RuleBase: ptr46, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr48 := &system.Type{Base: ptr35, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr44, "title": ptr47}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr49 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr50 := &system.RuleBase{Selector: "{images:photo} .width"}
	ptr51 := &system.Int_rule{Base: ptr49, RuleBase: ptr50, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 800, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr52 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2b", Exists: true}, Rules: []system.Rule{ptr51}, Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr53 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr54 := &system.RuleBase{}
	ptr55 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr56 := &system.RuleBase{}
	ptr57 := &images.Photo_rule{Base: ptr55, RuleBase: ptr56}
	ptr58 := &system.Map_rule{Base: ptr53, RuleBase: ptr54, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr60 := &system.RuleBase{}
	ptr61 := &system.String_rule{Base: ptr59, RuleBase: ptr60, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr62 := &system.Type{Base: ptr52, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr58, "title": ptr61}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr63 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr64 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr65 := &system.RuleBase{}
	ptr66 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr67 := &system.RuleBase{}
	ptr68 := &images.Image_rule{Base: ptr66, RuleBase: ptr67, Secure: system.Bool{}}
	ptr69 := &system.Map_rule{Base: ptr64, RuleBase: ptr65, Items: ptr68, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr70 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr71 := &system.RuleBase{}
	ptr72 := &words.Localizer_rule{Base: ptr70, RuleBase: ptr71}
	ptr73 := &system.Type{Base: ptr63, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr69, "title": ptr72}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr74 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr75 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr76 := &system.RuleBase{}
	ptr77 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr78 := &system.RuleBase{}
	ptr79 := &images.Image_rule{Base: ptr77, RuleBase: ptr78, Secure: system.Bool{Value: true, Exists: true}}
	ptr80 := &system.Map_rule{Base: ptr75, RuleBase: ptr76, Items: ptr79, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr81 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr82 := &system.RuleBase{}
	ptr83 := &words.Localizer_rule{Base: ptr81, RuleBase: ptr82}
	ptr84 := &system.Type{Base: ptr74, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr80, "title": ptr83}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/demo/site", "@gallery1", ptr1, 0x8997ec74747e71d5)
	system.RegisterType("kego.io/demo/site", "@gallery1a", ptr3, 0xf6237017f97fc838)
	system.RegisterType("kego.io/demo/site", "@gallery2", ptr5, 0x3ff7f11dfc3e88c0)
	system.RegisterType("kego.io/demo/site", "@gallery2a", ptr7, 0x4a0bb74542b31d47)
	system.RegisterType("kego.io/demo/site", "@gallery2b", ptr9, 0x8dc28cbd83066e3d)
	system.RegisterType("kego.io/demo/site", "@gallery3", ptr11, 0x327c5e90f012bd49)
	system.RegisterType("kego.io/demo/site", "@gallery3a", ptr13, 0x89fb5c874cf0f786)
	system.RegisterType("kego.io/demo/site", "gallery1", ptr18, 0x8997ec74747e71d5)
	system.RegisterType("kego.io/demo/site", "gallery1a", ptr23, 0xf6237017f97fc838)
	system.RegisterType("kego.io/demo/site", "gallery2", ptr34, 0x3ff7f11dfc3e88c0)
	system.RegisterType("kego.io/demo/site", "gallery2a", ptr48, 0x4a0bb74542b31d47)
	system.RegisterType("kego.io/demo/site", "gallery2b", ptr62, 0x8dc28cbd83066e3d)
	system.RegisterType("kego.io/demo/site", "gallery3", ptr73, 0x327c5e90f012bd49)
	system.RegisterType("kego.io/demo/site", "gallery3a", ptr84, 0x89fb5c874cf0f786)
}
