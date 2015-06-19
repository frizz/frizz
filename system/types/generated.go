package types

import (
	"kego.io/system"
)

func init() {

	ptr8729404448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729108704 := &system.Base{Context: ptr8729404448, Description: "Lists imports used in this package.", Id: "imports", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729404288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729108816 := &system.Base{Context: ptr8729404288, Description: "Map of import name to path.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729404256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729108928 := &system.Base{Context: ptr8729404256, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729206960 := &system.String_rule{Base: ptr8729108928, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728881792 := &system.Map_rule{Base: ptr8729108816, RuleBase: (*system.RuleBase)(nil), Items: ptr8729206960, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729364832 := &system.Type{Base: ptr8729108704, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728881792}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729409152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112512 := &system.Base{Context: ptr8729409152, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729408864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112624 := &system.Base{Context: ptr8729408864, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729408416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112848 := &system.Base{Context: ptr8729408416, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729408384 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388448 := &system.Int_rule{Base: ptr8729112848, RuleBase: ptr8729408384, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729408736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112288 := &system.Base{Context: ptr8729408736, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729408704 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388528 := &system.Int_rule{Base: ptr8729112288, RuleBase: ptr8729408704, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729408096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112736 := &system.Base{Context: ptr8729408096, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@rule", Exists: true}}

	ptr8729398144 := &system.Rule_rule{Base: ptr8729112736, RuleBase: (*system.RuleBase)(nil)}

	ptr8729365312 := &system.Type{Base: ptr8729112624, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"minItems": ptr8729388448, "maxItems": ptr8729388528, "items": ptr8729398144}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729365152 := &system.Type{Base: ptr8729112512, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729365312}

	ptr8728800416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113968 := &system.Base{Context: ptr8728800416, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728800128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114192 := &system.Base{Context: ptr8728800128, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728800000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114304 := &system.Base{Context: ptr8728800000, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@reference", Exists: true}}

	ptr8728799968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728885504 := &system.Reference_rule{Base: ptr8729114304, RuleBase: ptr8728799968, Default: system.Reference{Package: "", Id: "", Exists: false}}

	ptr8729366752 := &system.Type{Base: ptr8729114192, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728885504}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729366592 := &system.Type{Base: ptr8729113968, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729366752}

	ptr8729406976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109376 := &system.Base{Context: ptr8729406976, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729405760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109488 := &system.Base{Context: ptr8729405760, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729405728 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388128 := &system.Int_rule{Base: ptr8729109488, RuleBase: ptr8729405728, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729406272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109712 := &system.Base{Context: ptr8729406272, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729406240 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388208 := &system.Int_rule{Base: ptr8729109712, RuleBase: ptr8729406240, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729406560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109824 := &system.Base{Context: ptr8729406560, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729406528 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388288 := &system.Int_rule{Base: ptr8729109824, RuleBase: ptr8729406528, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729406848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112400 := &system.Base{Context: ptr8729406848, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8729406816 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388368 := &system.Int_rule{Base: ptr8729112400, RuleBase: ptr8729406816, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729364672 := &system.Type{Base: ptr8729109376, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729388128, "multipleOf": ptr8729388208, "minimum": ptr8729388288, "maximum": ptr8729388368}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728800832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114416 := &system.Base{Context: ptr8728800832, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729367232 := &system.Type{Base: ptr8729114416, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109936 := &system.Base{Context: ptr8728803840, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728803552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110048 := &system.Base{Context: ptr8728803552, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728803072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110384 := &system.Base{Context: ptr8728803072, Description: "This is the maximum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8728803040 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729387968 := &system.Int_rule{Base: ptr8729110384, RuleBase: ptr8728803040, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728803424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110496 := &system.Base{Context: ptr8728803424, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8728803392 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728803296 := &system.Bool_rule{Base: ptr8729110496, RuleBase: ptr8728803392, Default: system.Bool{Value: false, Exists: true}}

	ptr8728802432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110160 := &system.Base{Context: ptr8728802432, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@rule", Exists: true}}

	ptr8729398768 := &system.Rule_rule{Base: ptr8729110160, RuleBase: (*system.RuleBase)(nil)}

	ptr8728802752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110272 := &system.Base{Context: ptr8728802752, Description: "This is the minimum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8728802720 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729387888 := &system.Int_rule{Base: ptr8729110272, RuleBase: ptr8728802720, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729362752 := &system.Type{Base: ptr8729110048, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"maxItems": ptr8729387968, "uniqueItems": ptr8728803296, "items": ptr8729398768, "minItems": ptr8729387888}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729362592 := &system.Type{Base: ptr8729109936, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729362752}

	ptr8729404224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110608 := &system.Base{Context: ptr8729404224, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729403552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111392 := &system.Base{Context: ptr8729403552, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@context", Exists: true}}

	ptr8729403520 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400416 := &system.Context_rule{Base: ptr8729111392, RuleBase: ptr8729403520}

	ptr8729404000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111504 := &system.Base{Context: ptr8729404000, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8729403968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729403936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111616 := &system.Base{Context: ptr8729403936, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@rule", Exists: true}}

	ptr8729400704 := &system.Rule_rule{Base: ptr8729111616, RuleBase: (*system.RuleBase)(nil)}

	ptr8729388048 := &system.Array_rule{Base: ptr8729111504, RuleBase: ptr8729403968, Items: ptr8729400704, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728804512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110720 := &system.Base{Context: ptr8728804512, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@reference", Exists: true}}

	ptr8728879104 := &system.Reference_rule{Base: ptr8729110720, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Id: "", Exists: false}}

	ptr8728804864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729110944 := &system.Base{Context: ptr8728804864, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@string", Exists: true}}

	ptr8728804832 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400016 := &system.JsonString_rule{Base: ptr8729110944, RuleBase: ptr8728804832}

	ptr8728805248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111168 := &system.Base{Context: ptr8728805248, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@string", Exists: true}}

	ptr8728805216 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400208 := &system.JsonString_rule{Base: ptr8729111168, RuleBase: ptr8728805216}

	ptr8729363232 := &system.Type{Base: ptr8729110608, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"context": ptr8729400416, "rules": ptr8729388048, "type": ptr8728879104, "id": ptr8729400016, "description": ptr8729400208}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728802528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114976 := &system.Base{Context: ptr8728802528, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729367712 := &system.Type{Base: ptr8729114976, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729407264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109264 := &system.Base{Context: ptr8729407264, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729364512 := &system.Type{Base: ptr8729109264, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729364672}

	ptr8729404928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729109152 := &system.Base{Context: ptr8729404928, Description: "Automatically created basic rule for imports", Id: "@imports", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729364352 := &system.Type{Base: ptr8729109152, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729406944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743312 := &system.Base{Context: ptr8729406944, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729368672 := &system.Type{Base: ptr8728743312, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729403584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729108592 := &system.Base{Context: ptr8729403584, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729362432 := &system.Type{Base: ptr8729108592, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729406464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739952 := &system.Base{Context: ptr8729406464, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728803008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741072 := &system.Base{Context: ptr8728803008, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@bool", Exists: true}}

	ptr8728802976 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729398048 := &system.JsonBool_rule{Base: ptr8728741072, RuleBase: ptr8728802976}

	ptr8729404480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742192 := &system.Base{Context: ptr8729404480, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@bool", Exists: true}}

	ptr8729404416 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729399984 := &system.JsonBool_rule{Base: ptr8728742192, RuleBase: ptr8729404416}

	ptr8729405152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742304 := &system.Base{Context: ptr8729405152, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@bool", Exists: true}}

	ptr8729405056 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400320 := &system.JsonBool_rule{Base: ptr8728742304, RuleBase: ptr8729405056}

	ptr8729405920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742416 := &system.Base{Context: ptr8729405920, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@map", Exists: true}}

	ptr8729405824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729405696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742640 := &system.Base{Context: ptr8729405696, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@rule", Exists: true}}

	ptr8729400768 := &system.Rule_rule{Base: ptr8728742640, RuleBase: (*system.RuleBase)(nil)}

	ptr8728882368 := &system.Map_rule{Base: ptr8728742416, RuleBase: ptr8729405824, Items: ptr8729400768, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729406336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742976 := &system.Base{Context: ptr8729406336, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@type", Exists: true}}

	ptr8729406304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400976 := &system.Type_rule{Base: ptr8728742976, RuleBase: ptr8729406304}

	ptr8728803616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741296 := &system.Base{Context: ptr8728803616, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8728803584 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728803520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741408 := &system.Base{Context: ptr8728803520, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@reference", Exists: true}}

	ptr8728882048 := &system.Reference_rule{Base: ptr8728741408, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Id: "", Exists: false}}

	ptr8729389248 := &system.Array_rule{Base: ptr8728741296, RuleBase: ptr8728803584, Items: ptr8728882048, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728804288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741632 := &system.Base{Context: ptr8728804288, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@reference", Exists: true}}

	ptr8728804224 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728882112 := &system.Reference_rule{Base: ptr8728741632, RuleBase: ptr8728804224, Default: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}}

	ptr8728804768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741744 := &system.Base{Context: ptr8728804768, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8728804736 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728804704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741856 := &system.Base{Context: ptr8728804704, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@reference", Exists: true}}

	ptr8728882304 := &system.Reference_rule{Base: ptr8728741856, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Id: "", Exists: false}}

	ptr8729389328 := &system.Array_rule{Base: ptr8728741744, RuleBase: ptr8728804736, Items: ptr8728882304, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729403744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742080 := &system.Base{Context: ptr8729403744, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729403712 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729208016 := &system.String_rule{Base: ptr8728742080, RuleBase: ptr8729403712, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729368512 := &system.Type{Base: ptr8728739952, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8729398048, "interface": ptr8729399984, "exclude": ptr8729400320, "fields": ptr8728882368, "rule": ptr8729400976, "embed": ptr8729389248, "extends": ptr8728882112, "is": ptr8729389328, "native": ptr8729208016}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728801824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115088 := &system.Base{Context: ptr8728801824, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728801472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115200 := &system.Base{Context: ptr8728801472, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728800544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729116208 := &system.Base{Context: ptr8728800544, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8728800512 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729207664 := &system.String_rule{Base: ptr8729116208, RuleBase: ptr8728800512, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728801344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729116320 := &system.Base{Context: ptr8728801344, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8728801312 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729207840 := &system.String_rule{Base: ptr8729116320, RuleBase: ptr8728801312, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728797888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115312 := &system.Base{Context: ptr8728797888, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8728797856 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729207136 := &system.String_rule{Base: ptr8729115312, RuleBase: ptr8728797856, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728798912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115424 := &system.Base{Context: ptr8728798912, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@array", Exists: true}}

	ptr8728798880 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728798848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115536 := &system.Base{Context: ptr8728798848, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8729207312 := &system.String_rule{Base: ptr8729115536, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729388768 := &system.Array_rule{Base: ptr8729115424, RuleBase: ptr8728798880, Items: ptr8729207312, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728799488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115648 := &system.Base{Context: ptr8728799488, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8728799456 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729388928 := &system.Int_rule{Base: ptr8729115648, RuleBase: ptr8728799456, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728799808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115760 := &system.Base{Context: ptr8728799808, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@int", Exists: true}}

	ptr8728799776 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729389008 := &system.Int_rule{Base: ptr8729115760, RuleBase: ptr8728799776, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728800192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729115984 := &system.Base{Context: ptr8728800192, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@string", Exists: true}}

	ptr8728800160 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729207488 := &system.String_rule{Base: ptr8729115984, RuleBase: ptr8728800160, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729368032 := &system.Type{Base: ptr8729115200, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"pattern": ptr8729207664, "format": ptr8729207840, "default": ptr8729207136, "enum": ptr8729388768, "minLength": ptr8729388928, "maxLength": ptr8729389008, "equal": ptr8729207488}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729367872 := &system.Type{Base: ptr8729115088, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729368032}

	ptr8729405632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111840 := &system.Base{Context: ptr8729405632, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729405344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111952 := &system.Base{Context: ptr8729405344, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729405216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112064 := &system.Base{Context: ptr8729405216, Description: "Default value if this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8729405184 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729405120 := &system.Bool_rule{Base: ptr8729112064, RuleBase: ptr8729405184, Default: system.Bool{Value: false, Exists: false}}

	ptr8729363712 := &system.Type{Base: ptr8729111952, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729405120}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729363552 := &system.Type{Base: ptr8729111840, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729363712}

	ptr8728802144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114640 := &system.Base{Context: ptr8728802144, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8728801632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114752 := &system.Base{Context: ptr8728801632, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@bool", Exists: true}}

	ptr8728801600 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729402400 := &system.JsonBool_rule{Base: ptr8729114752, RuleBase: ptr8728801600}

	ptr8728801984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114864 := &system.Base{Context: ptr8728801984, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Id: "@string", Exists: true}}

	ptr8728801952 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729402592 := &system.JsonString_rule{Base: ptr8729114864, RuleBase: ptr8728801952}

	ptr8729367552 := &system.Type{Base: ptr8729114640, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729402400, "selector": ptr8729402592}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729404576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729111728 := &system.Base{Context: ptr8729404576, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729363392 := &system.Type{Base: ptr8729111728, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728798720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112960 := &system.Base{Context: ptr8728798720, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729411328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113632 := &system.Base{Context: ptr8729411328, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8729411296 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729411200 := &system.Bool_rule{Base: ptr8729113632, RuleBase: ptr8729411296, Default: system.Bool{Value: false, Exists: true}}

	ptr8728797344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113744 := &system.Base{Context: ptr8728797344, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8728797280 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904160 := &system.Number_rule{Base: ptr8729113744, RuleBase: ptr8728797280, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728798400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113856 := &system.Base{Context: ptr8728798400, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@bool", Exists: true}}

	ptr8728798368 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728798112 := &system.Bool_rule{Base: ptr8729113856, RuleBase: ptr8728798368, Default: system.Bool{Value: false, Exists: true}}

	ptr8729410016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113296 := &system.Base{Context: ptr8729410016, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729409984 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728903776 := &system.Number_rule{Base: ptr8729113296, RuleBase: ptr8729409984, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729410432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113408 := &system.Base{Context: ptr8729410432, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729410400 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904256 := &system.Number_rule{Base: ptr8729113408, RuleBase: ptr8729410400, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729410848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113520 := &system.Base{Context: ptr8729410848, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "@number", Exists: true}}

	ptr8729410816 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904064 := &system.Number_rule{Base: ptr8729113520, RuleBase: ptr8729410816, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729366112 := &system.Type{Base: ptr8729112960, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMinimum": ptr8729411200, "maximum": ptr8728904160, "exclusiveMaximum": ptr8728798112, "default": ptr8728903776, "multipleOf": ptr8728904256, "minimum": ptr8728904064}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729406112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729112176 := &system.Base{Context: ptr8729406112, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729364192 := &system.Type{Base: ptr8729112176, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Package: "kego.io/system", Id: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728799104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729113072 := &system.Base{Context: ptr8728799104, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729365952 := &system.Type{Base: ptr8729113072, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Id: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729366112}

	ptr8728801184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729114528 := &system.Base{Context: ptr8728801184, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Id: "type", Exists: true}}

	ptr8729367392 := &system.Type{Base: ptr8729114528, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Id: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Id: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Id: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729362752)

	system.RegisterType("kego.io/system:@base", ptr8729363392)

	system.RegisterType("kego.io/system:@bool", ptr8729363712)

	system.RegisterType("kego.io/system:@context", ptr8729362432)

	system.RegisterType("kego.io/system:@imports", ptr8729364352)

	system.RegisterType("kego.io/system:@int", ptr8729364672)

	system.RegisterType("kego.io/system:@map", ptr8729365312)

	system.RegisterType("kego.io/system:@number", ptr8729366112)

	system.RegisterType("kego.io/system:@reference", ptr8729366752)

	system.RegisterType("kego.io/system:@rule", ptr8729367392)

	system.RegisterType("kego.io/system:@ruleBase", ptr8729367712)

	system.RegisterType("kego.io/system:@string", ptr8729368032)

	system.RegisterType("kego.io/system:@type", ptr8729368672)

	system.RegisterType("kego.io/system:array", ptr8729362592)

	system.RegisterType("kego.io/system:base", ptr8729363232)

	system.RegisterType("kego.io/system:bool", ptr8729363552)

	system.RegisterType("kego.io/system:context", ptr8729364192)

	system.RegisterType("kego.io/system:imports", ptr8729364832)

	system.RegisterType("kego.io/system:int", ptr8729364512)

	system.RegisterType("kego.io/system:map", ptr8729365152)

	system.RegisterType("kego.io/system:number", ptr8729365952)

	system.RegisterType("kego.io/system:reference", ptr8729366592)

	system.RegisterType("kego.io/system:rule", ptr8729367232)

	system.RegisterType("kego.io/system:ruleBase", ptr8729367552)

	system.RegisterType("kego.io/system:string", ptr8729367872)

	system.RegisterType("kego.io/system:type", ptr8729368512)

}
