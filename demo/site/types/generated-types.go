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
	ptr0 := &system.Object_base{Description: "Automatically created basic rule for gallery1", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object_base{Description: "Automatically created basic rule for gallery1a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object_base: ptr2, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object_base{Description: "Automatically created basic rule for gallery2", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object_base: ptr4, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object_base{Description: "Automatically created basic rule for gallery2a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Object_base: ptr6, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object_base{Description: "Automatically created basic rule for gallery2b", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2b", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object_base: ptr8, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object_base{Description: "Automatically created basic rule for gallery3", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object_base: ptr10, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object_base{Description: "Automatically created basic rule for gallery3a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object_base: ptr12, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.Rule_base{}
	ptr17 := &system.String_rule{Object_base: ptr15, Rule_base: ptr16, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr18 := &system.Type{Object_base: ptr14, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr17}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.Rule_base{}
	ptr22 := &system.String_rule{Object_base: ptr20, Rule_base: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 20, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Type{Object_base: ptr19, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr22}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr26 := &system.Rule_base{}
	ptr27 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr28 := &system.Rule_base{}
	ptr29 := &images.Photo_rule{Object_base: ptr27, Rule_base: ptr28}
	ptr30 := &system.Map_rule{Object_base: ptr25, Rule_base: ptr26, Items: ptr29, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr31 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr32 := &system.Rule_base{}
	ptr33 := &system.String_rule{Object_base: ptr31, Rule_base: ptr32, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34 := &system.Type{Object_base: ptr24, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr30, "title": ptr33}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr35 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr36 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr37 := &system.Rule_base{}
	ptr38 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr39 := &system.Rule_base{Selector: ".protocol"}
	ptr40 := &system.String_rule{Object_base: ptr38, Rule_base: ptr39, Default: system.String{}, Enum: []string(nil), Equal: system.String{Value: "https", Exists: true}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr41 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule{ptr40}, Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr42 := &system.Rule_base{}
	ptr43 := &images.Photo_rule{Object_base: ptr41, Rule_base: ptr42}
	ptr44 := &system.Map_rule{Object_base: ptr36, Rule_base: ptr37, Items: ptr43, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr45 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr46 := &system.Rule_base{}
	ptr47 := &system.String_rule{Object_base: ptr45, Rule_base: ptr46, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr48 := &system.Type{Object_base: ptr35, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr44, "title": ptr47}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr49 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr50 := &system.Rule_base{Selector: "{images:photo} .width"}
	ptr51 := &system.Int_rule{Object_base: ptr49, Rule_base: ptr50, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 800, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr52 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2b", Exists: true}, Rules: []system.Rule{ptr51}, Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr53 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr54 := &system.Rule_base{}
	ptr55 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr56 := &system.Rule_base{}
	ptr57 := &images.Photo_rule{Object_base: ptr55, Rule_base: ptr56}
	ptr58 := &system.Map_rule{Object_base: ptr53, Rule_base: ptr54, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr60 := &system.Rule_base{}
	ptr61 := &system.String_rule{Object_base: ptr59, Rule_base: ptr60, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr62 := &system.Type{Object_base: ptr52, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr58, "title": ptr61}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr63 := &system.Object_base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr64 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr65 := &system.Rule_base{}
	ptr66 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr67 := &system.Rule_base{}
	ptr68 := &images.Image_rule{Object_base: ptr66, Rule_base: ptr67, Secure: system.Bool{}}
	ptr69 := &system.Map_rule{Object_base: ptr64, Rule_base: ptr65, Items: ptr68, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr70 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr71 := &system.Rule_base{}
	ptr72 := &words.Localizer_rule{Object_base: ptr70, Rule_base: ptr71}
	ptr73 := &system.Type{Object_base: ptr63, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr69, "title": ptr72}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr74 := &system.Object_base{Description: "This represents a gallery - it's just a list of images.", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr75 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr76 := &system.Rule_base{}
	ptr77 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr78 := &system.Rule_base{}
	ptr79 := &images.Image_rule{Object_base: ptr77, Rule_base: ptr78, Secure: system.Bool{Value: true, Exists: true}}
	ptr80 := &system.Map_rule{Object_base: ptr75, Rule_base: ptr76, Items: ptr79, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr81 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr82 := &system.Rule_base{}
	ptr83 := &words.Localizer_rule{Object_base: ptr81, Rule_base: ptr82}
	ptr84 := &system.Type{Object_base: ptr74, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr80, "title": ptr83}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/site", "@gallery1", ptr1, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "@gallery1a", ptr3, 0xf6237017f97fc838)
	system.Register("kego.io/demo/site", "@gallery2", ptr5, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "@gallery2a", ptr7, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "@gallery2b", ptr9, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "@gallery3", ptr11, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "@gallery3a", ptr13, 0x6c37f443b6dc9d05)
	system.Register("kego.io/demo/site", "gallery1", ptr18, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "gallery1a", ptr23, 0xf6237017f97fc838)
	system.Register("kego.io/demo/site", "gallery2", ptr34, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "gallery2a", ptr48, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "gallery2b", ptr62, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "gallery3", ptr73, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "gallery3a", ptr84, 0x6c37f443b6dc9d05)
}
