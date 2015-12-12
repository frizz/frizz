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
	ptr45 := &system.Reference{Package: "kego.io/system", Name: "@rule"}
	ptr46 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr45}
	ptr47 := &system.Rule{Interface: true}
	ptr48 := &system.RuleRule{Object: ptr46, Rule: ptr47}
	ptr49 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr50 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr49}
	ptr51 := &system.Rule{Interface: true}
	ptr52 := &system.IntRule{Object: ptr50, Rule: ptr51, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: (*system.Int)(nil), MultipleOf: (*system.Int)(nil)}
	ptr53 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr54 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr53}
	ptr55 := &system.Rule{Interface: true}
	ptr56 := &system.StringRule{Object: ptr54, Rule: ptr55, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr57 := system.String("object")
	ptr58 := &system.Type{Object: ptr44, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"bar": ptr48, "foo": ptr52, "title": ptr56}, Is: []*system.Reference(nil), Native: &ptr57, Rule: (*system.Type)(nil)}
	ptr59 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery1a"}
	ptr60 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr61 := &system.Object{Description: "This represents a gallery - it has a title and an image", Id: ptr59, Rules: []system.RuleInterface(nil), Type: ptr60}
	ptr62 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@icon"}
	ptr63 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr62}
	ptr64 := &system.Rule{}
	ptr65 := &images.IconRule{Object: ptr63, Rule: ptr64}
	ptr66 := &system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle"}
	ptr67 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr66}
	ptr68 := &system.Rule{}
	ptr69 := &units.RectangleRule{Object: ptr67, Rule: ptr68}
	ptr70 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr71 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr70}
	ptr72 := &system.Rule{}
	ptr73 := system.Int(20)
	ptr74 := &system.StringRule{Object: ptr71, Rule: ptr72, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: &ptr73, MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr75 := system.String("object")
	ptr76 := &system.Type{Object: ptr61, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"image": ptr65, "size": ptr69, "title": ptr74}, Is: []*system.Reference(nil), Native: &ptr75, Rule: (*system.Type)(nil)}
	ptr77 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2"}
	ptr78 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr79 := &system.Object{Description: "This represents a gallery - it has a title and a map of photos", Id: ptr77, Rules: []system.RuleInterface(nil), Type: ptr78}
	ptr80 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr81 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr80}
	ptr82 := &system.Rule{}
	ptr83 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr84 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr83}
	ptr85 := &system.Rule{}
	ptr86 := &images.PhotoRule{Object: ptr84, Rule: ptr85}
	ptr87 := &system.MapRule{Object: ptr81, Rule: ptr82, Items: ptr86, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr88 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr89 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr88}
	ptr90 := &system.Rule{}
	ptr91 := &system.StringRule{Object: ptr89, Rule: ptr90, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr92 := system.String("object")
	ptr93 := &system.Type{Object: ptr79, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr87, "title": ptr91}, Is: []*system.Reference(nil), Native: &ptr92, Rule: (*system.Type)(nil)}
	ptr94 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2a"}
	ptr95 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr96 := &system.Object{Description: "This represents a gallery - it has a title and a map of images with a restriction rule", Id: ptr94, Rules: []system.RuleInterface(nil), Type: ptr95}
	ptr97 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr98 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr97}
	ptr99 := &system.Rule{}
	ptr100 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr101 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr100}
	ptr102 := &system.Rule{Selector: ".protocol"}
	ptr103 := system.String("https")
	ptr104 := &system.StringRule{Object: ptr101, Rule: ptr102, Default: (*system.String)(nil), Enum: []string(nil), Equal: &ptr103, Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr105 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr106 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface{ptr104}, Type: ptr105}
	ptr107 := &system.Rule{}
	ptr108 := &images.PhotoRule{Object: ptr106, Rule: ptr107}
	ptr109 := &system.MapRule{Object: ptr98, Rule: ptr99, Items: ptr108, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr110 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr111 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr110}
	ptr112 := &system.Rule{}
	ptr113 := &system.StringRule{Object: ptr111, Rule: ptr112, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr114 := system.String("object")
	ptr115 := &system.Type{Object: ptr96, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr109, "title": ptr113}, Is: []*system.Reference(nil), Native: &ptr114, Rule: (*system.Type)(nil)}
	ptr116 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery2b"}
	ptr117 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr118 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr117}
	ptr119 := &system.Rule{Selector: "{images:photo} .width"}
	ptr120 := system.Int(800)
	ptr121 := &system.IntRule{Object: ptr118, Rule: ptr119, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: &ptr120, MultipleOf: (*system.Int)(nil)}
	ptr122 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr123 := &system.Object{Description: "This represents a gallery - it has a title and a list of images with a restriction rule", Id: ptr116, Rules: []system.RuleInterface{ptr121}, Type: ptr122}
	ptr124 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr125 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr124}
	ptr126 := &system.Rule{}
	ptr127 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr128 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr127}
	ptr129 := &system.Rule{}
	ptr130 := &images.PhotoRule{Object: ptr128, Rule: ptr129}
	ptr131 := &system.MapRule{Object: ptr125, Rule: ptr126, Items: ptr130, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr132 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr133 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr132}
	ptr134 := &system.Rule{}
	ptr135 := &system.StringRule{Object: ptr133, Rule: ptr134, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr136 := system.String("object")
	ptr137 := &system.Type{Object: ptr123, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr131, "title": ptr135}, Is: []*system.Reference(nil), Native: &ptr136, Rule: (*system.Type)(nil)}
	ptr138 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3"}
	ptr139 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr140 := &system.Object{Description: "This represents a gallery - it has a localizer title and a map of images", Id: ptr138, Rules: []system.RuleInterface(nil), Type: ptr139}
	ptr141 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr142 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr141}
	ptr143 := &system.Rule{}
	ptr144 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr145 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr144}
	ptr146 := &system.Rule{}
	ptr147 := &images.ImageRule{Object: ptr145, Rule: ptr146, Secure: (*system.Bool)(nil)}
	ptr148 := &system.MapRule{Object: ptr142, Rule: ptr143, Items: ptr147, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr149 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr150 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr149}
	ptr151 := &system.Rule{}
	ptr152 := &words.LocalizerRule{Object: ptr150, Rule: ptr151}
	ptr153 := system.String("object")
	ptr154 := &system.Type{Object: ptr140, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr148, "title": ptr152}, Is: []*system.Reference(nil), Native: &ptr153, Rule: (*system.Type)(nil)}
	ptr155 := &system.Reference{Package: "kego.io/demo/site", Name: "gallery3a"}
	ptr156 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr157 := &system.Object{Description: "This represents a gallery - it has a localizer title and a map of images with a custom rule", Id: ptr155, Rules: []system.RuleInterface(nil), Type: ptr156}
	ptr158 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr159 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr158}
	ptr160 := &system.Rule{}
	ptr161 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr162 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr161}
	ptr163 := &system.Rule{}
	ptr164 := system.Bool(true)
	ptr165 := &images.ImageRule{Object: ptr162, Rule: ptr163, Secure: &ptr164}
	ptr166 := &system.MapRule{Object: ptr159, Rule: ptr160, Items: ptr165, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr167 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr168 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr167}
	ptr169 := &system.Rule{}
	ptr170 := &words.LocalizerRule{Object: ptr168, Rule: ptr169}
	ptr171 := system.String("object")
	ptr172 := &system.Type{Object: ptr157, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"images": ptr166, "title": ptr170}, Is: []*system.Reference(nil), Native: &ptr171, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/site", "@gallery1", ptr5, 0xfb88eab775368c2b)
	system.Register("kego.io/demo/site", "@gallery1a", ptr11, 0x6466a314b49930a6)
	system.Register("kego.io/demo/site", "@gallery2", ptr17, 0x1ae7f9edab96cccc)
	system.Register("kego.io/demo/site", "@gallery2a", ptr23, 0x36e8e3d8deac6786)
	system.Register("kego.io/demo/site", "@gallery2b", ptr29, 0xbacab45e6c8f30a9)
	system.Register("kego.io/demo/site", "@gallery3", ptr35, 0x674e3fdc86422374)
	system.Register("kego.io/demo/site", "@gallery3a", ptr41, 0xc3da8d08e80fb86e)
	system.Register("kego.io/demo/site", "gallery1", ptr58, 0xfb88eab775368c2b)
	system.Register("kego.io/demo/site", "gallery1a", ptr76, 0x6466a314b49930a6)
	system.Register("kego.io/demo/site", "gallery2", ptr93, 0x1ae7f9edab96cccc)
	system.Register("kego.io/demo/site", "gallery2a", ptr115, 0x36e8e3d8deac6786)
	system.Register("kego.io/demo/site", "gallery2b", ptr137, 0xbacab45e6c8f30a9)
	system.Register("kego.io/demo/site", "gallery3", ptr154, 0x674e3fdc86422374)
	system.Register("kego.io/demo/site", "gallery3a", ptr172, 0xc3da8d08e80fb86e)
}
