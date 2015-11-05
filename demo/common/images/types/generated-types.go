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
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@image"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Restriction rules for images", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := &system.Reference{Package: "kego.io/system", Name: "@bool"}
	ptr11 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr10}
	ptr12 := &system.Rule{Optional: true}
	ptr13 := &system.BoolRule{Object: ptr11, Rule: ptr12, Default: (*system.Bool)(nil)}
	ptr14 := system.String("object")
	ptr15 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface{"secure": ptr13}, Is: []*system.Reference(nil), Native: &ptr14, Rule: (*system.Type)(nil)}
	ptr16 := &system.Reference{Package: "kego.io/demo/common/images", Name: "@photo"}
	ptr17 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr18 := &system.Object{Description: "Automatically created basic rule for photo", Id: ptr16, Rules: []system.RuleInterface(nil), Type: ptr17}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr20 := system.String("object")
	ptr21 := &system.Type{Object: ptr18, Embed: []*system.Reference{ptr19}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr20, Rule: (*system.Type)(nil)}
	ptr22 := &system.Reference{Package: "kego.io/demo/common/images", Name: "icon"}
	ptr23 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr24 := &system.Object{Description: "This is a type of image, which just contains the url of the image", Id: ptr22, Rules: []system.RuleInterface(nil), Type: ptr23}
	ptr25 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr26 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr25}
	ptr27 := &system.Rule{}
	ptr28 := &system.StringRule{Object: ptr26, Rule: ptr27, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr29 := &system.Reference{Package: "kego.io/demo/common/images", Name: "image"}
	ptr30 := system.String("object")
	ptr31 := &system.Type{Object: ptr24, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"url": ptr28}, Is: []*system.Reference{ptr29}, Native: &ptr30, Rule: (*system.Type)(nil)}
	ptr32 := &system.Reference{Package: "kego.io/demo/common/images", Name: "image"}
	ptr33 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr34 := &system.Object{Description: "This interface type represents all images - icon and photo", Id: ptr32, Rules: []system.RuleInterface(nil), Type: ptr33}
	ptr35 := system.String("object")
	ptr36 := &system.Type{Object: ptr34, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Is: []*system.Reference(nil), Native: &ptr35, Rule: ptr15}
	ptr37 := &system.Reference{Package: "kego.io/demo/common/images", Name: "photo"}
	ptr38 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr39 := &system.Object{Description: "This represents an image, and contains path, server and protocol separately", Id: ptr37, Rules: []system.RuleInterface(nil), Type: ptr38}
	ptr40 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr41 := &system.Object{Description: "The path for the url - e.g. /foo/bar.jpg", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr40}
	ptr42 := &system.Rule{}
	ptr43 := system.String("^/.*$")
	ptr44 := &system.StringRule{Object: ptr41, Rule: ptr42, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: &ptr43}
	ptr45 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr46 := &system.Object{Description: "The protocol for the url - e.g. http or https", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr45}
	ptr47 := &system.Rule{Optional: true}
	ptr48 := system.String("http")
	ptr49 := &system.StringRule{Object: ptr46, Rule: ptr47, Default: &ptr48, Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr50 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr51 := &system.Object{Description: "The server for the url - e.g. www.google.com", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr50}
	ptr52 := &system.Rule{}
	ptr53 := &system.StringRule{Object: ptr51, Rule: ptr52, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr54 := &system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle"}
	ptr55 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr54}
	ptr56 := &system.Rule{}
	ptr57 := &units.RectangleRule{Object: ptr55, Rule: ptr56}
	ptr58 := &system.Reference{Package: "kego.io/demo/common/images", Name: "image"}
	ptr59 := system.String("object")
	ptr60 := &system.Type{Object: ptr39, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"path": ptr44, "protocol": ptr49, "server": ptr53, "size": ptr57}, Is: []*system.Reference{ptr58}, Native: &ptr59, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/images", "@icon", ptr5, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "@image", ptr15, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "@photo", ptr21, 0xf2b64533e434a543)
	system.Register("kego.io/demo/common/images", "icon", ptr31, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "image", ptr36, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "photo", ptr60, 0xf2b64533e434a543)
}
