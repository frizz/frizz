package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Id: system.Reference{Package: "kego.io/demo/common/units", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr4 := &system.RuleBase{}
	ptr5 := &system.Int_rule{Base: ptr3, RuleBase: ptr4, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr6 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr7 := &system.RuleBase{}
	ptr8 := &system.Int_rule{Base: ptr6, RuleBase: ptr7, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr9 := &system.Type{Base: ptr2, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr5, "width": ptr8}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/demo/common/units", "@rectangle", ptr1, 0xe9ce1340e46b290e)
	system.RegisterType("kego.io/demo/common/units", "rectangle", ptr9, 0xe9ce1340e46b290e)
}
