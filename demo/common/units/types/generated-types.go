package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object{Id: system.Reference{Package: "kego.io/demo/common/units", Name: "rectangle", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr4 := &system.Rule{}
	ptr5 := &system.IntRule{Object: ptr3, Rule: ptr4, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr6 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr7 := &system.Rule{}
	ptr8 := &system.IntRule{Object: ptr6, Rule: ptr7, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr9 := &system.Type{Object: ptr2, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"height": ptr5, "width": ptr8}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/units", "@rectangle", ptr1, 0xe9ce1340e46b290e)
	system.Register("kego.io/demo/common/units", "rectangle", ptr9, 0xe9ce1340e46b290e)
}
