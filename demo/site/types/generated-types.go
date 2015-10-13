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
	ptr0 := &system.Object{Description: "Automatically created basic rule for gallery1", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object{Description: "Automatically created basic rule for gallery1a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery1a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object{Description: "Automatically created basic rule for gallery2", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object{Description: "Automatically created basic rule for gallery2a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Object: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object{Description: "Automatically created basic rule for gallery2b", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery2b", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object{Description: "Automatically created basic rule for gallery3", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object: ptr10, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object{Description: "Automatically created basic rule for gallery3a", Id: system.Reference{Package: "kego.io/demo/site", Name: "@gallery3a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object: ptr12, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.Rule{}
	ptr17 := &system.StringRule{Object: ptr15, Rule: ptr16, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr18 := &system.Type{Object: ptr14, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"title": ptr17}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery1a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.Rule{}
	ptr22 := &system.StringRule{Object: ptr20, Rule: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 20, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Type{Object: ptr19, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"title": ptr22}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr26 := &system.Rule{}
	ptr27 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr28 := &system.Rule{}
	ptr29 := &images.PhotoRule{Object: ptr27, Rule: ptr28}
	ptr30 := &system.MapRule{Object: ptr25, Rule: ptr26, Items: ptr29, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr31 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr32 := &system.Rule{}
	ptr33 := &system.StringRule{Object: ptr31, Rule: ptr32, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34 := &system.Type{Object: ptr24, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr30, "title": ptr33}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr35 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr36 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr37 := &system.Rule{}
	ptr38 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr39 := &system.Rule{Selector: ".protocol"}
	ptr40 := &system.StringRule{Object: ptr38, Rule: ptr39, Default: system.String{}, Enum: []string(nil), Equal: system.String{Value: "https", Exists: true}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr41 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface{ptr40}, Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr42 := &system.Rule{}
	ptr43 := &images.PhotoRule{Object: ptr41, Rule: ptr42}
	ptr44 := &system.MapRule{Object: ptr36, Rule: ptr37, Items: ptr43, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr45 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr46 := &system.Rule{}
	ptr47 := &system.StringRule{Object: ptr45, Rule: ptr46, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr48 := &system.Type{Object: ptr35, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr44, "title": ptr47}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr49 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr50 := &system.Rule{Selector: "{images:photo} .width"}
	ptr51 := &system.IntRule{Object: ptr49, Rule: ptr50, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 800, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr52 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery2b", Exists: true}, Rules: []system.RuleInterface{ptr51}, Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr53 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr54 := &system.Rule{}
	ptr55 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}}
	ptr56 := &system.Rule{}
	ptr57 := &images.PhotoRule{Object: ptr55, Rule: ptr56}
	ptr58 := &system.MapRule{Object: ptr53, Rule: ptr54, Items: ptr57, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr59 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr60 := &system.Rule{}
	ptr61 := &system.StringRule{Object: ptr59, Rule: ptr60, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr62 := &system.Type{Object: ptr52, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr58, "title": ptr61}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr63 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr64 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr65 := &system.Rule{}
	ptr66 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr67 := &system.Rule{}
	ptr68 := &images.ImageRule{Object: ptr66, Rule: ptr67, Secure: system.Bool{}}
	ptr69 := &system.MapRule{Object: ptr64, Rule: ptr65, Items: ptr68, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr70 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr71 := &system.Rule{}
	ptr72 := &words.LocalizerRule{Object: ptr70, Rule: ptr71}
	ptr73 := &system.Type{Object: ptr63, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr69, "title": ptr72}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr74 := &system.Object{Description: "This represents a gallery - it's just a list of images.", Id: system.Reference{Package: "kego.io/demo/site", Name: "gallery3a", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr75 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr76 := &system.Rule{}
	ptr77 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}}
	ptr78 := &system.Rule{}
	ptr79 := &images.ImageRule{Object: ptr77, Rule: ptr78, Secure: system.Bool{Value: true, Exists: true}}
	ptr80 := &system.MapRule{Object: ptr75, Rule: ptr76, Items: ptr79, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr81 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}}
	ptr82 := &system.Rule{}
	ptr83 := &words.LocalizerRule{Object: ptr81, Rule: ptr82}
	ptr84 := &system.Type{Object: ptr74, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr80, "title": ptr83}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
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
