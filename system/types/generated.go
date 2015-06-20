package types

import (
	"kego.io/system"
)

func init() {

	ptr8728963712 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963840 := &system.Base{Description: "Type of the object.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729140992 := &system.Reference_rule{Base: ptr8728963840, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8728963968 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729276928 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729141056 := &system.Reference_rule{Base: ptr8728963968, RuleBase: ptr8729276928, Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8728964096 := &system.Base{Description: "Description for the developer", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729277248 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729005440 := &system.JsonString_rule{Base: ptr8728964096, RuleBase: ptr8729277248}

	ptr8728964224 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729277536 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728964352 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729005664 := &system.Rule_rule{Base: ptr8728964352, RuleBase: (*system.RuleBase)(nil)}

	ptr8729076432 := &system.Array_rule{Base: ptr8728964224, RuleBase: ptr8729277536, Items: ptr8729005664, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729232160 := &system.Type{Base: ptr8728963712, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729140992, "id": ptr8729141056, "description": ptr8729005440, "rules": ptr8729076432}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728964608 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728964736 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728964864 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729278528 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729278464 := &system.Bool_rule{Base: ptr8728964864, RuleBase: ptr8729278528, Default: system.Bool{Value: false, Exists: false}}

	ptr8729232640 := &system.Type{Base: ptr8728964736, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729278464}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729232480 := &system.Type{Base: ptr8728964608, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729232640}

	ptr8728967936 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728968064 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729280480 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729006064 := &system.JsonBool_rule{Base: ptr8728968064, RuleBase: ptr8729280480}

	ptr8728968192 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8728732480 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729006272 := &system.JsonString_rule{Base: ptr8728968192, RuleBase: ptr8728732480}

	ptr8729236160 := &system.Type{Base: ptr8728967936, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729006064, "selector": ptr8729006272}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729110656 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729237280 := &system.Type{Base: ptr8729110656, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728967424 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728967552 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729279040 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728816960 := &system.Reference_rule{Base: ptr8728967552, RuleBase: ptr8729279040, Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729235360 := &system.Type{Base: ptr8728967424, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728816960}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728967808 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729236000 := &system.Type{Base: ptr8728967808, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728964480 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729232320 := &system.Type{Base: ptr8728964480, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728963072 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963584 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729275872 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729275776 := &system.Bool_rule{Base: ptr8728963584, RuleBase: ptr8729275872, Default: system.Bool{Value: false, Exists: true}}

	ptr8728963200 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729004352 := &system.Rule_rule{Base: ptr8728963200, RuleBase: (*system.RuleBase)(nil)}

	ptr8728963328 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729275328 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076272 := &system.Int_rule{Base: ptr8728963328, RuleBase: ptr8729275328, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728963456 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729275584 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076352 := &system.Int_rule{Base: ptr8728963456, RuleBase: ptr8729275584, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729231680 := &system.Type{Base: ptr8728963072, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"uniqueItems": ptr8729275776, "items": ptr8729004352, "minItems": ptr8729076272, "maxItems": ptr8729076352}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728964992 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728965248 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728965376 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729157808 := &system.String_rule{Base: ptr8728965376, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728817536 := &system.Map_rule{Base: ptr8728965248, RuleBase: (*system.RuleBase)(nil), Items: ptr8729157808, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729233760 := &system.Type{Base: ptr8728964992, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728817536}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728966656 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728962560 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729273024 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076752 := &system.Int_rule{Base: ptr8728962560, RuleBase: ptr8729273024, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728966784 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729272352 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076512 := &system.Int_rule{Base: ptr8728966784, RuleBase: ptr8729272352, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728961024 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729272576 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076592 := &system.Int_rule{Base: ptr8728961024, RuleBase: ptr8729272576, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728962304 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729272800 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076672 := &system.Int_rule{Base: ptr8728962304, RuleBase: ptr8729272800, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729234240 := &system.Type{Base: ptr8728966656, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"maximum": ptr8729076752, "default": ptr8729076512, "multipleOf": ptr8729076592, "minimum": ptr8729076672}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728966528 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729234080 := &system.Type{Base: ptr8728966528, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729234240}

	ptr8728965504 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729233920 := &system.Type{Base: ptr8728965504, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728965888 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728966016 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728966144 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729275712 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728838816 := &system.Number_rule{Base: ptr8728966144, RuleBase: ptr8729275712, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728966272 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729276160 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728838912 := &system.Number_rule{Base: ptr8728966272, RuleBase: ptr8729276160, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728966400 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729276576 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728838240 := &system.Number_rule{Base: ptr8728966400, RuleBase: ptr8729276576, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728966912 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729277056 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729276960 := &system.Bool_rule{Base: ptr8728966912, RuleBase: ptr8729277056, Default: system.Bool{Value: false, Exists: true}}

	ptr8728967040 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729277344 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728838336 := &system.Number_rule{Base: ptr8728967040, RuleBase: ptr8729277344, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728967168 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729277792 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729277696 := &system.Bool_rule{Base: ptr8728967168, RuleBase: ptr8729277792, Default: system.Bool{Value: false, Exists: true}}

	ptr8729234720 := &system.Type{Base: ptr8728966016, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728838816, "multipleOf": ptr8728838912, "minimum": ptr8728838240, "exclusiveMinimum": ptr8729276960, "maximum": ptr8728838336, "exclusiveMaximum": ptr8729277696}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729234560 := &system.Type{Base: ptr8728965888, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729234720}

	ptr8728962688 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728962816 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728965120 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729002848 := &system.Rule_rule{Base: ptr8728965120, RuleBase: (*system.RuleBase)(nil)}

	ptr8728965632 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729274304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076832 := &system.Int_rule{Base: ptr8728965632, RuleBase: ptr8729274304, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728965760 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729274560 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729076912 := &system.Int_rule{Base: ptr8728965760, RuleBase: ptr8729274560, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729233120 := &system.Type{Base: ptr8728962816, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729002848, "minItems": ptr8729076832, "maxItems": ptr8729076912}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729231360 := &system.Type{Base: ptr8728962688, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729233120}

	ptr8728967296 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729235200 := &system.Type{Base: ptr8728967296, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729235360}

	ptr8728968320 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729236320 := &system.Type{Base: ptr8728968320, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728968448 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728968576 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728968704 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728733760 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729158512 := &system.String_rule{Base: ptr8728968704, RuleBase: ptr8728733760, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728968832 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728734016 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728968960 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729157984 := &system.String_rule{Base: ptr8728968960, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729077152 := &system.Array_rule{Base: ptr8728968832, RuleBase: ptr8728734016, Items: ptr8729157984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728969088 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728734304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729077312 := &system.Int_rule{Base: ptr8728969088, RuleBase: ptr8728734304, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729108480 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728734528 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729077392 := &system.Int_rule{Base: ptr8729108480, RuleBase: ptr8728734528, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729108608 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728734752 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729158160 := &system.String_rule{Base: ptr8729108608, RuleBase: ptr8728734752, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729108736 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728734976 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729158336 := &system.String_rule{Base: ptr8729108736, RuleBase: ptr8728734976, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729108864 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728735456 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729158688 := &system.String_rule{Base: ptr8729108864, RuleBase: ptr8728735456, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729236640 := &system.Type{Base: ptr8728968576, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729158512, "enum": ptr8729077152, "minLength": ptr8729077312, "maxLength": ptr8729077392, "equal": ptr8729158160, "pattern": ptr8729158336, "format": ptr8729158688}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729236480 := &system.Type{Base: ptr8728968448, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729236640}

	ptr8728962944 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729231520 := &system.Type{Base: ptr8728962944, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729231680}

	ptr8728967680 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729235840 := &system.Type{Base: ptr8728967680, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729108992 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729109120 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728731840 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729002032 := &system.JsonBool_rule{Base: ptr8729109120, RuleBase: ptr8728731840}

	ptr8729109248 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728732800 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729109376 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728813824 := &system.Reference_rule{Base: ptr8729109376, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729077632 := &system.Array_rule{Base: ptr8729109248, RuleBase: ptr8728732800, Items: ptr8728813824, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729109504 := &system.Base{Description: "Type which this should extend", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728733504 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728814208 := &system.Reference_rule{Base: ptr8729109504, RuleBase: ptr8728733504, Default: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}}

	ptr8729110016 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728734848 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729003424 := &system.JsonBool_rule{Base: ptr8729110016, RuleBase: ptr8728734848}

	ptr8729110272 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728735488 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729110400 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729003872 := &system.Rule_rule{Base: ptr8729110400, RuleBase: (*system.RuleBase)(nil)}

	ptr8728815104 := &system.Map_rule{Base: ptr8729110272, RuleBase: ptr8728735488, Items: ptr8729003872, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729110528 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8728735744 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729004048 := &system.Type_rule{Base: ptr8729110528, RuleBase: ptr8728735744}

	ptr8729109632 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728733824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729109760 := &system.Base{Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728814464 := &system.Reference_rule{Base: ptr8729109760, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729077712 := &system.Array_rule{Base: ptr8729109632, RuleBase: ptr8728733824, Items: ptr8728814464, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729109888 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728734464 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729157632 := &system.String_rule{Base: ptr8729109888, RuleBase: ptr8728734464, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729110144 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728735168 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729003616 := &system.JsonBool_rule{Base: ptr8729110144, RuleBase: ptr8728735168}

	ptr8729237120 := &system.Type{Base: ptr8729108992, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8729002032, "embed": ptr8729077632, "extends": ptr8728814208, "interface": ptr8729003424, "fields": ptr8728815104, "rule": ptr8729004048, "is": ptr8729077712, "native": ptr8729157632, "exclude": ptr8729003616}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "base", ptr8729232160)

	system.RegisterType("kego.io/system", "bool", ptr8729232480)

	system.RegisterType("kego.io/system", "@reference", ptr8729235360)

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729236320)

	system.RegisterType("kego.io/system", "type", ptr8729237120)

	system.RegisterType("kego.io/system", "map", ptr8729231360)

	system.RegisterType("kego.io/system", "reference", ptr8729235200)

	system.RegisterType("kego.io/system", "string", ptr8729236480)

	system.RegisterType("kego.io/system", "@number", ptr8729234720)

	system.RegisterType("kego.io/system", "rule", ptr8729235840)

	system.RegisterType("kego.io/system", "@string", ptr8729236640)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729236160)

	system.RegisterType("kego.io/system", "@base", ptr8729232320)

	system.RegisterType("kego.io/system", "@array", ptr8729231680)

	system.RegisterType("kego.io/system", "int", ptr8729234080)

	system.RegisterType("kego.io/system", "@imports", ptr8729233920)

	system.RegisterType("kego.io/system", "number", ptr8729234560)

	system.RegisterType("kego.io/system", "@map", ptr8729233120)

	system.RegisterType("kego.io/system", "@type", ptr8729237280)

	system.RegisterType("kego.io/system", "@rule", ptr8729236000)

	system.RegisterType("kego.io/system", "imports", ptr8729233760)

	system.RegisterType("kego.io/system", "@int", ptr8729234240)

	system.RegisterType("kego.io/system", "@bool", ptr8729232640)

	system.RegisterType("kego.io/system", "array", ptr8729231520)

}
