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
	ptr54 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr55 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr54}
	ptr56 := &system.Rule{}
	ptr57 := &images.ImageRule{Object: ptr55, Rule: ptr56, Secure: (*system.Bool)(nil)}
	ptr58 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr59 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr58}
	ptr60 := &system.Rule{}
	ptr61 := system.Int(20)
	ptr62 := &system.StringRule{Object: ptr59, Rule: ptr60, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: &ptr61, MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr63 := system.String("object")
	ptr64 := &system.Type{Object: ptr53, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"image": ptr57, "title": ptr62}, Is: []*system.Reference(nil), Native: &ptr63, Rule: (*system.Type)(nil)}
	ptr65 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2"}
	ptr66 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr67 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr65, Rules: []system.RuleInterface(nil), Type: ptr66}
	ptr68 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr69 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr68}
	ptr70 := &system.Rule{}
	ptr71 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr72 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr71}
	ptr73 := &system.Rule{}
	ptr74 := &images.PhotoRule{Object: ptr72, Rule: ptr73}
	ptr75 := &system.MapRule{Object: ptr69, Rule: ptr70, Items: ptr74, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr76 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr77 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr76}
	ptr78 := &system.Rule{}
	ptr79 := &system.StringRule{Object: ptr77, Rule: ptr78, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr80 := system.String("object")
	ptr81 := &system.Type{Object: ptr67, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr75, "title": ptr79}, Is: []*system.Reference(nil), Native: &ptr80, Rule: (*system.Type)(nil)}
	ptr82 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2a"}
	ptr83 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr84 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr82, Rules: []system.RuleInterface(nil), Type: ptr83}
	ptr85 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr86 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr85}
	ptr87 := &system.Rule{}
	ptr88 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr89 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr88}
	ptr90 := &system.Rule{Selector: ".protocol"}
	ptr91 := system.String("https")
	ptr92 := &system.StringRule{Object: ptr89, Rule: ptr90, Default: (*system.String)(nil), Enum: []string(nil), Equal: &ptr91, Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr93 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr94 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface{ptr92}, Type: ptr93}
	ptr95 := &system.Rule{}
	ptr96 := &images.PhotoRule{Object: ptr94, Rule: ptr95}
	ptr97 := &system.MapRule{Object: ptr86, Rule: ptr87, Items: ptr96, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr98 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr99 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr98}
	ptr100 := &system.Rule{}
	ptr101 := &system.StringRule{Object: ptr99, Rule: ptr100, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr102 := system.String("object")
	ptr103 := &system.Type{Object: ptr84, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr97, "title": ptr101}, Is: []*system.Reference(nil), Native: &ptr102, Rule: (*system.Type)(nil)}
	ptr104 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2b"}
	ptr105 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr106 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr105}
	ptr107 := &system.Rule{Selector: "{images:photo} .width"}
	ptr108 := system.Int(800)
	ptr109 := &system.IntRule{Object: ptr106, Rule: ptr107, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: &ptr108, MultipleOf: (*system.Int)(nil)}
	ptr110 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr111 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr104, Rules: []system.RuleInterface{ptr109}, Type: ptr110}
	ptr112 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr113 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr112}
	ptr114 := &system.Rule{}
	ptr115 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr116 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr115}
	ptr117 := &system.Rule{}
	ptr118 := &images.PhotoRule{Object: ptr116, Rule: ptr117}
	ptr119 := &system.MapRule{Object: ptr113, Rule: ptr114, Items: ptr118, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr120 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr121 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr120}
	ptr122 := &system.Rule{}
	ptr123 := &system.StringRule{Object: ptr121, Rule: ptr122, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr124 := system.String("object")
	ptr125 := &system.Type{Object: ptr111, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr119, "title": ptr123}, Is: []*system.Reference(nil), Native: &ptr124, Rule: (*system.Type)(nil)}
	ptr126 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3"}
	ptr127 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr128 := &system.Object{Description: "This represents a gallery - it's just a list of images", Id: ptr126, Rules: []system.RuleInterface(nil), Type: ptr127}
	ptr129 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr130 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr129}
	ptr131 := &system.Rule{}
	ptr132 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr133 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr132}
	ptr134 := &system.Rule{}
	ptr135 := &images.ImageRule{Object: ptr133, Rule: ptr134, Secure: (*system.Bool)(nil)}
	ptr136 := &system.MapRule{Object: ptr130, Rule: ptr131, Items: ptr135, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr137 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr138 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr137}
	ptr139 := &system.Rule{}
	ptr140 := &words.LocalizerRule{Object: ptr138, Rule: ptr139}
	ptr141 := system.String("object")
	ptr142 := &system.Type{Object: ptr128, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr136, "title": ptr140}, Is: []*system.Reference(nil), Native: &ptr141, Rule: (*system.Type)(nil)}
	ptr143 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3a"}
	ptr144 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr145 := &system.Object{Description: "This represents a gallery - it's just a list of images.", Id: ptr143, Rules: []system.RuleInterface(nil), Type: ptr144}
	ptr146 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr147 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr146}
	ptr148 := &system.Rule{}
	ptr149 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr150 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr149}
	ptr151 := &system.Rule{}
	ptr152 := system.Bool(true)
	ptr153 := &images.ImageRule{Object: ptr150, Rule: ptr151, Secure: &ptr152}
	ptr154 := &system.MapRule{Object: ptr147, Rule: ptr148, Items: ptr153, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr155 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr156 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr155}
	ptr157 := &system.Rule{}
	ptr158 := &words.LocalizerRule{Object: ptr156, Rule: ptr157}
	ptr159 := system.String("object")
	ptr160 := &system.Type{Object: ptr145, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr154, "title": ptr158}, Is: []*system.Reference(nil), Native: &ptr159, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/site", "@gallery1", ptr5, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "@gallery1a", ptr11, 0xe078528a307382b0)
	system.Register("kego.io/demo/site", "@gallery2", ptr17, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "@gallery2a", ptr23, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "@gallery2b", ptr29, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "@gallery3", ptr35, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "@gallery3a", ptr41, 0x6c37f443b6dc9d05)
	system.Register("kego.io/demo/site", "gallery1", ptr50, 0x8997ec74747e71d5)
	system.Register("kego.io/demo/site", "gallery1a", ptr64, 0xe078528a307382b0)
	system.Register("kego.io/demo/site", "gallery2", ptr81, 0x3ff7f11dfc3e88c0)
	system.Register("kego.io/demo/site", "gallery2a", ptr103, 0x4a0bb74542b31d47)
	system.Register("kego.io/demo/site", "gallery2b", ptr125, 0x8dc28cbd83066e3d)
	system.Register("kego.io/demo/site", "gallery3", ptr142, 0x327c5e90f012bd49)
	system.Register("kego.io/demo/site", "gallery3a", ptr160, 0x6c37f443b6dc9d05)
}
