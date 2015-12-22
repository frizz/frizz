package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for localizer", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@simple"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Automatically created basic rule for simple", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := system.String("object")
	ptr11 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface(nil), Native: &ptr10, Rule: (*system.Type)(nil)}
	ptr12 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@translation"}
	ptr13 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr14 := &system.Object{Description: "Automatically created basic rule for translation", Id: ptr12, Rules: []system.RuleInterface(nil), Type: ptr13}
	ptr15 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr16 := system.String("object")
	ptr17 := &system.Type{Object: ptr14, Embed: []*system.Reference{ptr15}, Fields: map[string]system.RuleInterface(nil), Native: &ptr16, Rule: (*system.Type)(nil)}
	ptr18 := &system.Reference{Package: "kego.io/demo/common/words", Name: "localizer"}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr20 := &system.Object{Id: ptr18, Rules: []system.RuleInterface(nil), Type: ptr19}
	ptr21 := system.String("object")
	ptr22 := &system.Type{Object: ptr20, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Native: &ptr21, Rule: (*system.Type)(nil)}
	ptr23 := &system.Reference{Package: "kego.io/demo/common/words", Name: "simple"}
	ptr24 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr25 := &system.Object{Id: ptr23, Rules: []system.RuleInterface(nil), Type: ptr24}
	ptr26 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr27 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr26}
	ptr28 := &system.Rule{}
	ptr29 := &system.StringRule{Object: ptr27, Rule: ptr28, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr30 := system.String("object")
	ptr31 := &system.Type{Object: ptr25, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"string": ptr29}, Native: &ptr30, Rule: (*system.Type)(nil)}
	ptr32 := &system.Reference{Package: "kego.io/demo/common/words", Name: "translation"}
	ptr33 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr34 := &system.Object{Description: "This represents a translated string", Id: ptr32, Rules: []system.RuleInterface(nil), Type: ptr33}
	ptr35 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr36 := &system.Object{Description: "The original English string", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr35}
	ptr37 := &system.Rule{Optional: true}
	ptr38 := system.String("http")
	ptr39 := &system.StringRule{Object: ptr36, Rule: ptr37, Default: &ptr38, Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr40 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr41 := &system.Object{Description: "The translated strings", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr40}
	ptr42 := &system.Rule{}
	ptr43 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr44 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr43}
	ptr45 := &system.Rule{}
	ptr46 := &system.StringRule{Object: ptr44, Rule: ptr45, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr47 := &system.MapRule{Object: ptr41, Rule: ptr42, Items: ptr46, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr48 := system.String("object")
	ptr49 := &system.Type{Object: ptr34, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"english": ptr39, "translations": ptr47}, Native: &ptr48, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/words", "@localizer", ptr5, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "@simple", ptr11, 0x66d71f5dedb65f21)
	system.Register("kego.io/demo/common/words", "@translation", ptr17, 0x2435101aebb14e59)
	system.Register("kego.io/demo/common/words", "localizer", ptr22, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "simple", ptr31, 0x66d71f5dedb65f21)
	system.Register("kego.io/demo/common/words", "translation", ptr49, 0x2435101aebb14e59)
}
