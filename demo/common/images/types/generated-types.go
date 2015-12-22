package types

import (
	"kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@icon"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for icon", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Restriction rules for images", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := &system.Reference{Package: "kego.io/system", Name: "@bool"}
	ptr11 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr10}
	ptr12 := &system.Rule{Optional: true}
	ptr13 := &system.BoolRule{Object: ptr11, Rule: ptr12, Default: (*system.Bool)(nil)}
	ptr14 := system.String("object")
	ptr15 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface{"secure": ptr13}, Native: &ptr14, Rule: (*system.Type)(nil)}
	ptr16 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr17 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr18 := &system.Object{Description: "Automatically created basic rule for photo", Id: ptr16, Rules: []system.RuleInterface(nil), Type: ptr17}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr20 := system.String("object")
	ptr21 := &system.Type{Object: ptr18, Embed: []*system.Reference{ptr19}, Fields: map[string]system.RuleInterface(nil), Native: &ptr20, Rule: (*system.Type)(nil)}
	ptr22 := &system.Reference{Package: "kego.io/demo/common/images", Name: "icon"}
	ptr23 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr24 := &system.Object{Description: "This is a type of image, which just contains the url of the image", Id: ptr22, Rules: []system.RuleInterface(nil), Type: ptr23}
	ptr25 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr26 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr25}
	ptr27 := &system.Rule{}
	ptr28 := &system.StringRule{Object: ptr26, Rule: ptr27, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr29 := system.String("object")
	ptr30 := &system.Type{Object: ptr24, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"url": ptr28}, Native: &ptr29, Rule: (*system.Type)(nil)}
	ptr31 := &system.Reference{Package: "kego.io/demo/common/images", Name: "image"}
	ptr32 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr33 := &system.Object{Description: "This interface type represents all images - icon and photo", Id: ptr31, Rules: []system.RuleInterface(nil), Type: ptr32}
	ptr34 := system.String("object")
	ptr35 := &system.Type{Object: ptr33, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Native: &ptr34, Rule: ptr15}
	ptr36 := &system.Reference{Package: "kego.io/demo/common/images", Name: "photo"}
	ptr37 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr38 := &system.Object{Description: "This represents an image, and contains path, server and protocol separately", Id: ptr36, Rules: []system.RuleInterface(nil), Type: ptr37}
	ptr39 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr40 := &system.Object{Description: "The path for the url - e.g. /foo/bar.jpg", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr39}
	ptr41 := &system.Rule{}
	ptr42 := system.String("^/.*$")
	ptr43 := &system.StringRule{Object: ptr40, Rule: ptr41, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: &ptr42}
	ptr44 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr45 := &system.Object{Description: "The protocol for the url - e.g. http or https", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr44}
	ptr46 := &system.Rule{Optional: true}
	ptr47 := system.String("http")
	ptr48 := &system.StringRule{Object: ptr45, Rule: ptr46, Default: &ptr47, Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr49 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr50 := &system.Object{Description: "The server for the url - e.g. www.google.com", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr49}
	ptr51 := &system.Rule{}
	ptr52 := &system.StringRule{Object: ptr50, Rule: ptr51, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr53 := &system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle"}
	ptr54 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr53}
	ptr55 := &system.Rule{}
	ptr56 := &units.RectangleRule{Object: ptr54, Rule: ptr55}
	ptr57 := system.String("object")
	ptr58 := &system.Type{Object: ptr38, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"path": ptr43, "protocol": ptr48, "server": ptr52, "size": ptr56}, Native: &ptr57, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/images", "@icon", ptr5, 0x4c8f82b332c9df85)
	system.Register("kego.io/demo/common/images", "@image", ptr15, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "@photo", ptr21, 0x9256fcb626fc81cf)
	system.Register("kego.io/demo/common/images", "icon", ptr30, 0x4c8f82b332c9df85)
	system.Register("kego.io/demo/common/images", "image", ptr35, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "photo", ptr58, 0x9256fcb626fc81cf)
}
