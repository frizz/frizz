package types

import (
	"kego.io/system"
)

func init() {

	ptr8728776896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729010176 := &system.Base{Context: ptr8728776896, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728776576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729011456 := &system.Base{Context: ptr8728776576, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728775232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729011712 := &system.Base{Context: ptr8728775232, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728694144 := &system.Rule_rule{Base: ptr8729011712, RuleBase: (*system.RuleBase)(nil)}

	ptr8728775648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729011840 := &system.Base{Context: ptr8728775648, Description: "This is the minimum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728775616 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728871392 := &system.Int_rule{Base: ptr8729011840, RuleBase: ptr8728775616, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728776000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729011968 := &system.Base{Context: ptr8728776000, Description: "This is the maximum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728775968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728871472 := &system.Int_rule{Base: ptr8729011968, RuleBase: ptr8728775968, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728776448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012096 := &system.Base{Context: ptr8728776448, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728776416 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728776224 := &system.Bool_rule{Base: ptr8729012096, RuleBase: ptr8728776416, Default: system.Bool{Value: false, Exists: true}}

	ptr8729004096 := &system.Type{Base: ptr8729011456, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728694144, "minItems": ptr8728871392, "maxItems": ptr8728871472, "uniqueItems": ptr8728776224}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729002160 := &system.Type{Base: ptr8729010176, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729004096}

	ptr8729273728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014656 := &system.Base{Context: ptr8729273728, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729273408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014784 := &system.Base{Context: ptr8729273408, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729272576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014912 := &system.Base{Context: ptr8729272576, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729267696 := &system.Rule_rule{Base: ptr8729014912, RuleBase: (*system.RuleBase)(nil)}

	ptr8729272928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015040 := &system.Base{Context: ptr8729272928, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729272896 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728872672 := &system.Int_rule{Base: ptr8729015040, RuleBase: ptr8729272896, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729273280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015168 := &system.Base{Context: ptr8729273280, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729273248 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728874192 := &system.Int_rule{Base: ptr8729015168, RuleBase: ptr8729273248, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729005680 := &system.Type{Base: ptr8729014784, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729267696, "minItems": ptr8728872672, "maxItems": ptr8728874192}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729005504 := &system.Type{Base: ptr8729014656, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729005680}

	ptr8728776608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013760 := &system.Base{Context: ptr8728776608, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729004976 := &system.Type{Base: ptr8729013760, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728773312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013120 := &system.Base{Context: ptr8728773312, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729002336 := &system.Type{Base: ptr8729013120, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729281600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051520 := &system.Base{Context: ptr8729281600, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728774784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051648 := &system.Base{Context: ptr8728774784, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728774752 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729269728 := &system.JsonBool_rule{Base: ptr8729051648, RuleBase: ptr8728774752}

	ptr8728777248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052032 := &system.Base{Context: ptr8728777248, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728777216 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728877472 := &system.Reference_rule{Base: ptr8729052032, RuleBase: ptr8728777216, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8728779456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052416 := &system.Base{Context: ptr8728779456, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728779392 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729008672 := &system.String_rule{Base: ptr8729052416, RuleBase: ptr8728779392, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729280544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052672 := &system.Base{Context: ptr8729280544, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729280512 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729271696 := &system.JsonBool_rule{Base: ptr8729052672, RuleBase: ptr8729280512}

	ptr8729281088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052800 := &system.Base{Context: ptr8729281088, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729281056 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729281024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052928 := &system.Base{Context: ptr8729281024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728691424 := &system.Rule_rule{Base: ptr8729052928, RuleBase: (*system.RuleBase)(nil)}

	ptr8729177024 := &system.Map_rule{Base: ptr8729052800, RuleBase: ptr8729281056, Items: ptr8728691424, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729281472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729053056 := &system.Base{Context: ptr8729281472, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729281440 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728692736 := &system.Type_rule{Base: ptr8729053056, RuleBase: ptr8729281440}

	ptr8728775712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051776 := &system.Base{Context: ptr8728775712, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728775680 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728775584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051904 := &system.Base{Context: ptr8728775584, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728877392 := &system.Reference_rule{Base: ptr8729051904, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728877312 := &system.Array_rule{Base: ptr8729051776, RuleBase: ptr8728775680, Items: ptr8728877392, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728778080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052160 := &system.Base{Context: ptr8728778080, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728778048 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728778016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052288 := &system.Base{Context: ptr8728778016, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728877632 := &system.Reference_rule{Base: ptr8729052288, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728877552 := &system.Array_rule{Base: ptr8729052160, RuleBase: ptr8728778048, Items: ptr8728877632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728780224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729052544 := &system.Base{Context: ptr8728780224, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728780160 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729271456 := &system.JsonBool_rule{Base: ptr8729052544, RuleBase: ptr8728780160}

	ptr8729008496 := &system.Type{Base: ptr8729051520, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8729269728, "extends": ptr8728877472, "native": ptr8729008672, "exclude": ptr8729271696, "fields": ptr8729177024, "rule": ptr8728692736, "embed": ptr8728877312, "is": ptr8728877552, "interface": ptr8729271456}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728776288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013632 := &system.Base{Context: ptr8728776288, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729004800 := &system.Type{Base: ptr8729013632, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728775136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013376 := &system.Base{Context: ptr8728775136, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728775008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013504 := &system.Base{Context: ptr8728775008, Description: "Default value if this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728774944 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728774848 := &system.Bool_rule{Base: ptr8729013504, RuleBase: ptr8728774944, Default: system.Bool{Value: false, Exists: false}}

	ptr8729004624 := &system.Type{Base: ptr8729013376, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728774848}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729279648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017600 := &system.Base{Context: ptr8729279648, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729276576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017856 := &system.Base{Context: ptr8729276576, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729276544 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729276448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017984 := &system.Base{Context: ptr8729276448, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729007792 := &system.String_rule{Base: ptr8729017984, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728871872 := &system.Array_rule{Base: ptr8729017856, RuleBase: ptr8729276544, Items: ptr8729007792, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729277056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729018112 := &system.Base{Context: ptr8729277056, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729277024 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728876992 := &system.Int_rule{Base: ptr8729018112, RuleBase: ptr8729277024, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729277536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729018240 := &system.Base{Context: ptr8729277536, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729277504 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728877072 := &system.Int_rule{Base: ptr8729018240, RuleBase: ptr8729277504, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729277984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051136 := &system.Base{Context: ptr8729277984, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729277920 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729007968 := &system.String_rule{Base: ptr8729051136, RuleBase: ptr8729277920, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729278496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051264 := &system.Base{Context: ptr8729278496, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729278464 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729008144 := &system.String_rule{Base: ptr8729051264, RuleBase: ptr8729278464, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729279360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051392 := &system.Base{Context: ptr8729279360, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729279328 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729008320 := &system.String_rule{Base: ptr8729051392, RuleBase: ptr8729279328, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729275904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017728 := &system.Base{Context: ptr8729275904, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729275872 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729007616 := &system.String_rule{Base: ptr8729017728, RuleBase: ptr8729275872, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729007440 := &system.Type{Base: ptr8729017600, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"enum": ptr8728871872, "minLength": ptr8728876992, "maxLength": ptr8728877072, "equal": ptr8729007968, "pattern": ptr8729008144, "format": ptr8729008320, "default": ptr8729007616}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728779904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013888 := &system.Base{Context: ptr8728779904, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728779360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014016 := &system.Base{Context: ptr8728779360, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728777920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014144 := &system.Base{Context: ptr8728777920, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728777888 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728872032 := &system.Int_rule{Base: ptr8729014144, RuleBase: ptr8728777888, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728778336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014272 := &system.Base{Context: ptr8728778336, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728778304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728872112 := &system.Int_rule{Base: ptr8729014272, RuleBase: ptr8728778304, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728778784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014400 := &system.Base{Context: ptr8728778784, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728778752 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728872192 := &system.Int_rule{Base: ptr8729014400, RuleBase: ptr8728778752, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728779200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729014528 := &system.Base{Context: ptr8728779200, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728779168 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728872352 := &system.Int_rule{Base: ptr8729014528, RuleBase: ptr8728779168, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729005328 := &system.Type{Base: ptr8729014016, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728872032, "multipleOf": ptr8728872112, "minimum": ptr8728872192, "maximum": ptr8728872352}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729005152 := &system.Type{Base: ptr8729013888, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729005328}

	ptr8729279296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016320 := &system.Base{Context: ptr8729279296, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729278976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016448 := &system.Base{Context: ptr8729278976, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729278848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016576 := &system.Base{Context: ptr8729278848, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729278816 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728874432 := &system.Reference_rule{Base: ptr8729016576, RuleBase: ptr8729278816, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729006384 := &system.Type{Base: ptr8729016448, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728874432}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729006208 := &system.Type{Base: ptr8729016320, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729006384}

	ptr8729280128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017472 := &system.Base{Context: ptr8729280128, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729007264 := &system.Type{Base: ptr8729017472, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729007440}

	ptr8729274752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017344 := &system.Base{Context: ptr8729274752, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729007088 := &system.Type{Base: ptr8729017344, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729282016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729053184 := &system.Base{Context: ptr8729282016, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729008848 := &system.Type{Base: ptr8729053184, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729272512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016832 := &system.Base{Context: ptr8729272512, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729006736 := &system.Type{Base: ptr8729016832, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729274176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016960 := &system.Base{Context: ptr8729274176, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729273440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017088 := &system.Base{Context: ptr8729273440, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729273376 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729264784 := &system.JsonBool_rule{Base: ptr8729017088, RuleBase: ptr8729273376}

	ptr8729274048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729017216 := &system.Base{Context: ptr8729274048, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729273920 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729265344 := &system.JsonString_rule{Base: ptr8729017216, RuleBase: ptr8729273920}

	ptr8729006912 := &system.Type{Base: ptr8729016960, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729264784, "selector": ptr8729265344}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728779840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012224 := &system.Base{Context: ptr8728779840, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728777728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012352 := &system.Base{Context: ptr8728777728, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728871552 := &system.Reference_rule{Base: ptr8729012352, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728778176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012480 := &system.Base{Context: ptr8728778176, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728778144 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728696928 := &system.JsonString_rule{Base: ptr8729012480, RuleBase: ptr8728778144}

	ptr8728778592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012608 := &system.Base{Context: ptr8728778592, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728778560 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728697440 := &system.JsonString_rule{Base: ptr8729012608, RuleBase: ptr8728778560}

	ptr8728778976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012736 := &system.Base{Context: ptr8728778976, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728778944 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728697680 := &system.Context_rule{Base: ptr8729012736, RuleBase: ptr8728778944}

	ptr8728779648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012864 := &system.Base{Context: ptr8728779648, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728779616 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728779584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729012992 := &system.Base{Context: ptr8728779584, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728698704 := &system.Rule_rule{Base: ptr8729012992, RuleBase: (*system.RuleBase)(nil)}

	ptr8728871632 := &system.Array_rule{Base: ptr8729012864, RuleBase: ptr8728779616, Items: ptr8728698704, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729004272 := &system.Type{Base: ptr8729012224, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8728871552, "id": ptr8728696928, "description": ptr8728697440, "context": ptr8728697680, "rules": ptr8728871632}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729277472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015424 := &system.Base{Context: ptr8729277472, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729276512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015936 := &system.Base{Context: ptr8729276512, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729276480 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729276352 := &system.Bool_rule{Base: ptr8729015936, RuleBase: ptr8729276480, Default: system.Bool{Value: false, Exists: true}}

	ptr8729276832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016064 := &system.Base{Context: ptr8729276832, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729276800 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728879680 := &system.Number_rule{Base: ptr8729016064, RuleBase: ptr8729276800, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729277344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016192 := &system.Base{Context: ptr8729277344, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729277312 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729277184 := &system.Bool_rule{Base: ptr8729016192, RuleBase: ptr8729277312, Default: system.Bool{Value: false, Exists: true}}

	ptr8729275104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015552 := &system.Base{Context: ptr8729275104, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729275072 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728879392 := &system.Number_rule{Base: ptr8729015552, RuleBase: ptr8729275072, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729275552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015680 := &system.Base{Context: ptr8729275552, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729275520 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728879488 := &system.Number_rule{Base: ptr8729015680, RuleBase: ptr8729275520, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729276000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015808 := &system.Base{Context: ptr8729276000, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729275968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728879584 := &system.Number_rule{Base: ptr8729015808, RuleBase: ptr8729275968, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729006032 := &system.Type{Base: ptr8729015424, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMinimum": ptr8729276352, "maximum": ptr8728879680, "exclusiveMaximum": ptr8729277184, "default": ptr8728879392, "multipleOf": ptr8728879488, "minimum": ptr8728879584}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729277792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729015296 := &system.Base{Context: ptr8729277792, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729005856 := &system.Type{Base: ptr8729015296, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729006032}

	ptr8728775520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729013248 := &system.Base{Context: ptr8728775520, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729004448 := &system.Type{Base: ptr8729013248, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729004624}

	ptr8729279776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729016704 := &system.Base{Context: ptr8729279776, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729006560 := &system.Type{Base: ptr8729016704, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729004096)

	system.RegisterType("kego.io/system:@base", ptr8729002336)

	system.RegisterType("kego.io/system:@bool", ptr8729004624)

	system.RegisterType("kego.io/system:@context", ptr8729004976)

	system.RegisterType("kego.io/system:@int", ptr8729005328)

	system.RegisterType("kego.io/system:@map", ptr8729005680)

	system.RegisterType("kego.io/system:@number", ptr8729006032)

	system.RegisterType("kego.io/system:@reference", ptr8729006384)

	system.RegisterType("kego.io/system:@rule", ptr8729006736)

	system.RegisterType("kego.io/system:@ruleBase", ptr8729007088)

	system.RegisterType("kego.io/system:@string", ptr8729007440)

	system.RegisterType("kego.io/system:@type", ptr8729008848)

	system.RegisterType("kego.io/system:array", ptr8729002160)

	system.RegisterType("kego.io/system:base", ptr8729004272)

	system.RegisterType("kego.io/system:bool", ptr8729004448)

	system.RegisterType("kego.io/system:context", ptr8729004800)

	system.RegisterType("kego.io/system:int", ptr8729005152)

	system.RegisterType("kego.io/system:map", ptr8729005504)

	system.RegisterType("kego.io/system:number", ptr8729005856)

	system.RegisterType("kego.io/system:reference", ptr8729006208)

	system.RegisterType("kego.io/system:rule", ptr8729006560)

	system.RegisterType("kego.io/system:ruleBase", ptr8729006912)

	system.RegisterType("kego.io/system:string", ptr8729007264)

	system.RegisterType("kego.io/system:type", ptr8729008496)

	// If there's no references to the non types package in here we
	// won't be able to compile, so we force it by calling a noop.
	system.DoNothing()
}
