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
	ptr0 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery1"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for gallery1", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery1a"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Automatically created basic rule for gallery1a", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := system.String("object")
	ptr11 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr10, Rule: (*system.Type)(nil)}
	ptr12 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery2"}
	ptr13 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr14 := &system.Object{Description: "Automatically created basic rule for gallery2", Id: ptr12, Rules: []system.RuleInterface(nil), Type: ptr13}
	ptr15 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr16 := system.String("object")
	ptr17 := &system.Type{Object: ptr14, Embed: []*system.Reference{ptr15}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr16, Rule: (*system.Type)(nil)}
	ptr18 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery2a"}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr20 := &system.Object{Description: "Automatically created basic rule for gallery2a", Id: ptr18, Rules: []system.RuleInterface(nil), Type: ptr19}
	ptr21 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr22 := system.String("object")
	ptr23 := &system.Type{Object: ptr20, Embed: []*system.Reference{ptr21}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr22, Rule: (*system.Type)(nil)}
	ptr24 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery2b"}
	ptr25 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr26 := &system.Object{Description: "Automatically created basic rule for gallery2b", Id: ptr24, Rules: []system.RuleInterface(nil), Type: ptr25}
	ptr27 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr28 := system.String("object")
	ptr29 := &system.Type{Object: ptr26, Embed: []*system.Reference{ptr27}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr28, Rule: (*system.Type)(nil)}
	ptr30 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery3"}
	ptr31 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr32 := &system.Object{Description: "Automatically created basic rule for gallery3", Id: ptr30, Rules: []system.RuleInterface(nil), Type: ptr31}
	ptr33 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr34 := system.String("object")
	ptr35 := &system.Type{Object: ptr32, Embed: []*system.Reference{ptr33}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr34, Rule: (*system.Type)(nil)}
	ptr36 := &system.Reference{Package: "kego.io/demo/site", Name: "@gallery3a"}
	ptr37 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr38 := &system.Object{Description: "Automatically created basic rule for gallery3a", Id: ptr36, Rules: []system.RuleInterface(nil), Type: ptr37}
	ptr39 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr40 := system.String("object")
	ptr41 := &system.Type{Object: ptr38, Embed: []*system.Reference{ptr39}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr40, Rule: (*system.Type)(nil)}
	ptr42 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery1"}
	ptr43 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr44 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr42, Rules: []system.RuleInterface(nil), Type: ptr43}
	ptr45 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr46 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr45}
	ptr47 := &system.Rule{}
	ptr48 := &system.StringRule{Object: ptr46, Rule: ptr47, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr49 := system.String("object")
	ptr50 := &system.Type{Object: ptr44, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"title": ptr48}, Is: []*system.Reference(nil), Native: &ptr49, Rule: (*system.Type)(nil)}
	ptr51 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery1a"}
	ptr52 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr53 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr51, Rules: []system.RuleInterface(nil), Type: ptr52}
	ptr54 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr55 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr54}
	ptr56 := &system.Rule{}
	ptr57 := system.Int(20)
	ptr58 := &system.StringRule{Object: ptr55, Rule: ptr56, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: &ptr57, MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr59 := system.String("object")
	ptr60 := &system.Type{Object: ptr53, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"title": ptr58}, Is: []*system.Reference(nil), Native: &ptr59, Rule: (*system.Type)(nil)}
	ptr61 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2"}
	ptr62 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr63 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr61, Rules: []system.RuleInterface(nil), Type: ptr62}
	ptr64 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr65 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr64}
	ptr66 := &system.Rule{}
	ptr67 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr68 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr67}
	ptr69 := &system.Rule{}
	ptr70 := &images.PhotoRule{Object: ptr68, Rule: ptr69}
	ptr71 := &system.MapRule{Object: ptr65, Rule: ptr66, Items: ptr70, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr72 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr73 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr72}
	ptr74 := &system.Rule{}
	ptr75 := &system.StringRule{Object: ptr73, Rule: ptr74, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr76 := system.String("object")
	ptr77 := &system.Type{Object: ptr63, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr71, "title": ptr75}, Is: []*system.Reference(nil), Native: &ptr76, Rule: (*system.Type)(nil)}
	ptr78 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2a"}
	ptr79 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr80 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr78, Rules: []system.RuleInterface(nil), Type: ptr79}
	ptr81 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr82 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr81}
	ptr83 := &system.Rule{}
	ptr84 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr85 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr84}
	ptr86 := &system.Rule{Selector: ".protocol"}
	ptr87 := system.String("https")
	ptr88 := &system.StringRule{Object: ptr85, Rule: ptr86, Default: (*system.String)(nil), Enum: []string(nil), Equal: &ptr87, Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr89 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr90 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface{ptr88}, Type: ptr89}
	ptr91 := &system.Rule{}
	ptr92 := &images.PhotoRule{Object: ptr90, Rule: ptr91}
	ptr93 := &system.MapRule{Object: ptr82, Rule: ptr83, Items: ptr92, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr94 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr95 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr94}
	ptr96 := &system.Rule{}
	ptr97 := &system.StringRule{Object: ptr95, Rule: ptr96, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr98 := system.String("object")
	ptr99 := &system.Type{Object: ptr80, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr93, "title": ptr97}, Is: []*system.Reference(nil), Native: &ptr98, Rule: (*system.Type)(nil)}
	ptr100 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2b"}
	ptr101 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr102 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr101}
	ptr103 := &system.Rule{Selector: "{images:photo} .width"}
	ptr104 := system.Int(800)
	ptr105 := &system.IntRule{Object: ptr102, Rule: ptr103, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: &ptr104, MultipleOf: (*system.Int)(nil)}
	ptr106 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr107 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr100, Rules: []system.RuleInterface{ptr105}, Type: ptr106}
	ptr108 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr109 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr108}
	ptr110 := &system.Rule{}
	ptr111 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr112 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr111}
	ptr113 := &system.Rule{}
	ptr114 := &images.PhotoRule{Object: ptr112, Rule: ptr113}
	ptr115 := &system.MapRule{Object: ptr109, Rule: ptr110, Items: ptr114, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr116 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr117 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr116}
	ptr118 := &system.Rule{}
	ptr119 := &system.StringRule{Object: ptr117, Rule: ptr118, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr120 := system.String("object")
	ptr121 := &system.Type{Object: ptr107, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr115, "title": ptr119}, Is: []*system.Reference(nil), Native: &ptr120, Rule: (*system.Type)(nil)}
	ptr122 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3"}
	ptr123 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr124 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr122, Rules: []system.RuleInterface(nil), Type: ptr123}
	ptr125 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr126 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr125}
	ptr127 := &system.Rule{}
	ptr128 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr129 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr128}
	ptr130 := &system.Rule{}
	ptr131 := &images.ImageRule{Object: ptr129, Rule: ptr130, Secure: (*system.Bool)(nil)}
	ptr132 := &system.MapRule{Object: ptr126, Rule: ptr127, Items: ptr131, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr133 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr134 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr133}
	ptr135 := &system.Rule{}
	ptr136 := &words.LocalizerRule{Object: ptr134, Rule: ptr135}
	ptr137 := system.String("object")
	ptr138 := &system.Type{Object: ptr124, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr132, "title": ptr136}, Is: []*system.Reference(nil), Native: &ptr137, Rule: (*system.Type)(nil)}
	ptr139 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3a"}
	ptr140 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr141 := &system.Object{Description: "This represents a gallery - it's just a list of images.", Id: ptr139, Rules: []system.RuleInterface(nil), Type: ptr140}
	ptr142 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr143 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr142}
	ptr144 := &system.Rule{}
	ptr145 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr146 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr145}
	ptr147 := &system.Rule{}
	ptr148 := system.Bool(true)
	ptr149 := &images.ImageRule{Object: ptr146, Rule: ptr147, Secure: &ptr148}
	ptr150 := &system.MapRule{Object: ptr143, Rule: ptr144, Items: ptr149, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr151 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr152 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr151}
	ptr153 := &system.Rule{}
	ptr154 := &words.LocalizerRule{Object: ptr152, Rule: ptr153}
	ptr155 := system.String("object")
	ptr156 := &system.Type{Object: ptr141, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr150, "title": ptr154}, Is: []*system.Reference(nil), Native: &ptr155, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/site", "@gallery1", ptr5, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "@gallery1a", ptr11, 0xf6237017f97fc838)
	system.Register("kego.io/demo/site", "@gallery2", ptr17, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "@gallery2a", ptr23, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "@gallery2b", ptr29, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "@gallery3", ptr35, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "@gallery3a", ptr41, 0x6c37f443b6dc9d05)
	system.Register("kego.io/demo/site", "gallery1", ptr50, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "gallery1a", ptr60, 0xf6237017f97fc838)
	system.Register("kego.io/demo/site", "gallery2", ptr77, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "gallery2a", ptr99, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "gallery2b", ptr121, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "gallery3", ptr138, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "gallery3a", ptr156, 0x6c37f443b6dc9d05)
}
