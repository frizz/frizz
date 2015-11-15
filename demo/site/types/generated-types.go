package types

import (
	"kego.io/demo/common/images"
	_ "kego.io/demo/common/images/types"
	"kego.io/demo/common/units"
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
	ptr44 := &system.Object{Description: "This represents a gallery - it just has a title", Id: ptr42, Rules: []system.RuleInterface(nil), Type: ptr43}
	ptr45 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr46 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr45}
	ptr47 := &system.Rule{}
	ptr48 := &system.StringRule{Object: ptr46, Rule: ptr47, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr49 := system.String("object")
	ptr50 := &system.Type{Object: ptr44, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"title": ptr48}, Is: []*system.Reference(nil), Native: &ptr49, Rule: (*system.Type)(nil)}
	ptr51 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery1a"}
	ptr52 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr53 := &system.Object{Description: "This represents a gallery - it has a title and an image", Id: ptr51, Rules: []system.RuleInterface(nil), Type: ptr52}
	ptr54 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@icon"}
	ptr55 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr54}
	ptr56 := &system.Rule{}
	ptr57 := &images.IconRule{Object: ptr55, Rule: ptr56}
	ptr58 := &system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle"}
	ptr59 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr58}
	ptr60 := &system.Rule{}
	ptr61 := &units.RectangleRule{Object: ptr59, Rule: ptr60}
	ptr62 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr63 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr62}
	ptr64 := &system.Rule{}
	ptr65 := system.Int(20)
	ptr66 := &system.StringRule{Object: ptr63, Rule: ptr64, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: &ptr65, MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr67 := system.String("object")
	ptr68 := &system.Type{Object: ptr53, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"image": ptr57, "size": ptr61, "title": ptr66}, Is: []*system.Reference(nil), Native: &ptr67, Rule: (*system.Type)(nil)}
	ptr69 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2"}
	ptr70 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr71 := &system.Object{Description: "This represents a gallery - it has a title and a map of photos", Id: ptr69, Rules: []system.RuleInterface(nil), Type: ptr70}
	ptr72 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr73 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr72}
	ptr74 := &system.Rule{}
	ptr75 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr76 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr75}
	ptr77 := &system.Rule{}
	ptr78 := &images.PhotoRule{Object: ptr76, Rule: ptr77}
	ptr79 := &system.MapRule{Object: ptr73, Rule: ptr74, Items: ptr78, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr80 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr81 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr80}
	ptr82 := &system.Rule{}
	ptr83 := &system.StringRule{Object: ptr81, Rule: ptr82, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr84 := system.String("object")
	ptr85 := &system.Type{Object: ptr71, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr79, "title": ptr83}, Is: []*system.Reference(nil), Native: &ptr84, Rule: (*system.Type)(nil)}
	ptr86 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2a"}
	ptr87 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr88 := &system.Object{Description: "This represents a gallery - it has a title and a map of images with a restriction rule", Id: ptr86, Rules: []system.RuleInterface(nil), Type: ptr87}
	ptr89 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr90 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr89}
	ptr91 := &system.Rule{}
	ptr92 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr93 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr92}
	ptr94 := &system.Rule{Selector: ".protocol"}
	ptr95 := system.String("https")
	ptr96 := &system.StringRule{Object: ptr93, Rule: ptr94, Default: (*system.String)(nil), Enum: []string(nil), Equal: &ptr95, Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr97 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr98 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface{ptr96}, Type: ptr97}
	ptr99 := &system.Rule{}
	ptr100 := &images.PhotoRule{Object: ptr98, Rule: ptr99}
	ptr101 := &system.MapRule{Object: ptr90, Rule: ptr91, Items: ptr100, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr102 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr103 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr102}
	ptr104 := &system.Rule{}
	ptr105 := &system.StringRule{Object: ptr103, Rule: ptr104, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr106 := system.String("object")
	ptr107 := &system.Type{Object: ptr88, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr101, "title": ptr105}, Is: []*system.Reference(nil), Native: &ptr106, Rule: (*system.Type)(nil)}
	ptr108 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2b"}
	ptr109 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr110 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr109}
	ptr111 := &system.Rule{Selector: "{images:photo} .width"}
	ptr112 := system.Int(800)
	ptr113 := &system.IntRule{Object: ptr110, Rule: ptr111, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: &ptr112, MultipleOf: (*system.Int)(nil)}
	ptr114 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr115 := &system.Object{Description: "This represents a gallery - it has a title and a list of images with a restriction rule", Id: ptr108, Rules: []system.RuleInterface{ptr113}, Type: ptr114}
	ptr116 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr117 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr116}
	ptr118 := &system.Rule{}
	ptr119 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr120 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr119}
	ptr121 := &system.Rule{}
	ptr122 := &images.PhotoRule{Object: ptr120, Rule: ptr121}
	ptr123 := &system.MapRule{Object: ptr117, Rule: ptr118, Items: ptr122, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr124 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr125 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr124}
	ptr126 := &system.Rule{}
	ptr127 := &system.StringRule{Object: ptr125, Rule: ptr126, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr128 := system.String("object")
	ptr129 := &system.Type{Object: ptr115, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr123, "title": ptr127}, Is: []*system.Reference(nil), Native: &ptr128, Rule: (*system.Type)(nil)}
	ptr130 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3"}
	ptr131 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr132 := &system.Object{Description: "This represents a gallery - it has a localizer title and a map of images", Id: ptr130, Rules: []system.RuleInterface(nil), Type: ptr131}
	ptr133 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr134 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr133}
	ptr135 := &system.Rule{}
	ptr136 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr137 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr136}
	ptr138 := &system.Rule{}
	ptr139 := &images.ImageRule{Object: ptr137, Rule: ptr138, Secure: (*system.Bool)(nil)}
	ptr140 := &system.MapRule{Object: ptr134, Rule: ptr135, Items: ptr139, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr141 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr142 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr141}
	ptr143 := &system.Rule{}
	ptr144 := &words.LocalizerRule{Object: ptr142, Rule: ptr143}
	ptr145 := system.String("object")
	ptr146 := &system.Type{Object: ptr132, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr140, "title": ptr144}, Is: []*system.Reference(nil), Native: &ptr145, Rule: (*system.Type)(nil)}
	ptr147 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3a"}
	ptr148 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr149 := &system.Object{Description: "This represents a gallery - it has a localizer title and a map of images with a custom rule", Id: ptr147, Rules: []system.RuleInterface(nil), Type: ptr148}
	ptr150 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr151 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr150}
	ptr152 := &system.Rule{}
	ptr153 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr154 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr153}
	ptr155 := &system.Rule{}
	ptr156 := system.Bool(true)
	ptr157 := &images.ImageRule{Object: ptr154, Rule: ptr155, Secure: &ptr156}
	ptr158 := &system.MapRule{Object: ptr151, Rule: ptr152, Items: ptr157, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr159 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr160 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr159}
	ptr161 := &system.Rule{}
	ptr162 := &words.LocalizerRule{Object: ptr160, Rule: ptr161}
	ptr163 := system.String("object")
	ptr164 := &system.Type{Object: ptr149, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr158, "title": ptr162}, Is: []*system.Reference(nil), Native: &ptr163, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/site", "@gallery1", ptr5, 0x22f817cfb011224a)
	system.Register("kego.io/demo/site", "@gallery1a", ptr11, 0x6466a314b49930a6)
	system.Register("kego.io/demo/site", "@gallery2", ptr17, 0x1ae7f9edab96cccc)
	system.Register("kego.io/demo/site", "@gallery2a", ptr23, 0x36e8e3d8deac6786)
	system.Register("kego.io/demo/site", "@gallery2b", ptr29, 0xbacab45e6c8f30a9)
	system.Register("kego.io/demo/site", "@gallery3", ptr35, 0x674e3fdc86422374)
	system.Register("kego.io/demo/site", "@gallery3a", ptr41, 0xc3da8d08e80fb86e)
	system.Register("kego.io/demo/site", "gallery1", ptr50, 0x22f817cfb011224a)
	system.Register("kego.io/demo/site", "gallery1a", ptr68, 0x6466a314b49930a6)
	system.Register("kego.io/demo/site", "gallery2", ptr85, 0x1ae7f9edab96cccc)
	system.Register("kego.io/demo/site", "gallery2a", ptr107, 0x36e8e3d8deac6786)
	system.Register("kego.io/demo/site", "gallery2b", ptr129, 0xbacab45e6c8f30a9)
	system.Register("kego.io/demo/site", "gallery3", ptr146, 0x674e3fdc86422374)
	system.Register("kego.io/demo/site", "gallery3a", ptr164, 0xc3da8d08e80fb86e)
}
