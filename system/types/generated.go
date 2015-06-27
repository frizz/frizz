package types

import (
	"kego.io/system"
)

func init() {

	ptr8729023872 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729024000 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729333808 := &system.Rule_rule{Base: ptr8729024000, RuleBase: (*system.RuleBase)(nil)}

	ptr8729024128 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728689024 := &system.RuleBase{Optional: true}

	ptr8729060448 := &system.Int_rule{Base: ptr8729024128, RuleBase: ptr8728689024, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729024256 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728689280 := &system.RuleBase{Optional: true}

	ptr8729060528 := &system.Int_rule{Base: ptr8729024256, RuleBase: ptr8728689280, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729024384 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728689568 := &system.RuleBase{Optional: true}

	ptr8729334240 := &system.JsonBool_rule{Base: ptr8729024384, RuleBase: ptr8728689568}

	ptr8729023744 := &system.Type{Base: ptr8729023872, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr8729333808, "minItems": ptr8729060448, "maxItems": ptr8729060528, "uniqueItems": ptr8729334240}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729175168 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175424 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175552 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729328448 := &system.RuleBase{Optional: true}

	ptr8729034112 := &system.Reference_rule{Base: ptr8729175552, RuleBase: ptr8729328448, Default: system.Reference{}}

	ptr8729175296 := &system.Type{Base: ptr8729175424, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729034112}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729175040 := &system.Type{Base: ptr8729175168, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729175296}

	ptr8729022080 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729022208 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729022336 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729297072 := &system.String_rule{Base: ptr8729022336, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729029632 := &system.Map_rule{Base: ptr8729022208, RuleBase: (*system.RuleBase)(nil), Items: ptr8729297072, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729021952 := &system.Type{Base: ptr8729022080, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr8729029632}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729175808 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175680 := &system.Type{Base: ptr8729175808, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729022592 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729022464 := &system.Type{Base: ptr8729022592, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729176320 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729176448 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729321568 := &system.RuleBase{Optional: true}

	ptr8729335104 := &system.JsonBool_rule{Base: ptr8729176448, RuleBase: ptr8729321568}

	ptr8729176576 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729321888 := &system.RuleBase{Optional: true}

	ptr8729329872 := &system.JsonString_rule{Base: ptr8729176576, RuleBase: ptr8729321888}

	ptr8729176192 := &system.Type{Base: ptr8729176320, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr8729335104, "selector": ptr8729329872}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729180288 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729180160 := &system.Type{Base: ptr8729180288, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729176064 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175936 := &system.Type{Base: ptr8729176064, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729026304 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729025280 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729025408 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729332032 := &system.Rule_rule{Base: ptr8729025408, RuleBase: (*system.RuleBase)(nil)}

	ptr8729025536 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729325472 := &system.RuleBase{Optional: true}

	ptr8729061008 := &system.Int_rule{Base: ptr8729025536, RuleBase: ptr8729325472, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729025664 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729325728 := &system.RuleBase{Optional: true}

	ptr8729061088 := &system.Int_rule{Base: ptr8729025664, RuleBase: ptr8729325728, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729025152 := &system.Type{Base: ptr8729025280, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr8729332032, "minItems": ptr8729061008, "maxItems": ptr8729061088}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729026176 := &system.Type{Base: ptr8729026304, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729025152}

	ptr8729026432 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729174144 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729174272 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729326464 := &system.RuleBase{Optional: true}

	ptr8728879200 := &system.Number_rule{Base: ptr8729174272, RuleBase: ptr8729326464, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174400 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729326688 := &system.RuleBase{Optional: true}

	ptr8728879296 := &system.Number_rule{Base: ptr8729174400, RuleBase: ptr8729326688, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174528 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729326912 := &system.RuleBase{Optional: true}

	ptr8728879392 := &system.Number_rule{Base: ptr8729174528, RuleBase: ptr8729326912, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174656 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729327200 := &system.RuleBase{Optional: true}

	ptr8729333376 := &system.JsonBool_rule{Base: ptr8729174656, RuleBase: ptr8729327200}

	ptr8729174784 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729327424 := &system.RuleBase{Optional: true}

	ptr8728879584 := &system.Number_rule{Base: ptr8729174784, RuleBase: ptr8729327424, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174912 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729327712 := &system.RuleBase{Optional: true}

	ptr8729333728 := &system.JsonBool_rule{Base: ptr8729174912, RuleBase: ptr8729327712}

	ptr8729174016 := &system.Type{Base: ptr8729174144, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8728879200, "multipleOf": ptr8728879296, "minimum": ptr8728879392, "exclusiveMinimum": ptr8729333376, "maximum": ptr8728879584, "exclusiveMaximum": ptr8729333728}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729025792 := &system.Type{Base: ptr8729026432, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729174016}

	ptr8729021184 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729021056 := &system.Type{Base: ptr8729021184, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729177344 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729178368 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729324416 := &system.RuleBase{Optional: true}

	ptr8729297776 := &system.String_rule{Base: ptr8729178368, RuleBase: ptr8729324416, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177472 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729322720 := &system.RuleBase{Optional: true}

	ptr8729297248 := &system.String_rule{Base: ptr8729177472, RuleBase: ptr8729322720, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177600 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729323072 := &system.RuleBase{Optional: true}

	ptr8729177728 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729330816 := &system.JsonString_rule{Base: ptr8729177728, RuleBase: (*system.RuleBase)(nil)}

	ptr8729061328 := &system.Array_rule{Base: ptr8729177600, RuleBase: ptr8729323072, Items: ptr8729330816, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729177856 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729323296 := &system.RuleBase{Optional: true}

	ptr8729061488 := &system.Int_rule{Base: ptr8729177856, RuleBase: ptr8729323296, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729177984 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729323584 := &system.RuleBase{Optional: true}

	ptr8729061568 := &system.Int_rule{Base: ptr8729177984, RuleBase: ptr8729323584, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729178112 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729323808 := &system.RuleBase{Optional: true}

	ptr8729297424 := &system.String_rule{Base: ptr8729178112, RuleBase: ptr8729323808, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729178240 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729324096 := &system.RuleBase{Optional: true}

	ptr8729297600 := &system.String_rule{Base: ptr8729178240, RuleBase: ptr8729324096, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177216 := &system.Type{Base: ptr8729177344, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"format": ptr8729297776, "default": ptr8729297248, "enum": ptr8729061328, "minLength": ptr8729061488, "maxLength": ptr8729061568, "equal": ptr8729297424, "pattern": ptr8729297600}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729177088 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729176960 := &system.Type{Base: ptr8729177088, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729177216}

	ptr8729024640 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729018368 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729321760 := &system.RuleBase{Optional: true}

	ptr8729018496 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729329840 := &system.Rule_rule{Base: ptr8729018496, RuleBase: (*system.RuleBase)(nil)}

	ptr8729060608 := &system.Array_rule{Base: ptr8729018368, RuleBase: ptr8729321760, Items: ptr8729329840, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729024768 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729030336 := &system.Reference_rule{Base: ptr8729024768, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729024896 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728690336 := &system.RuleBase{Optional: true}

	ptr8729030464 := &system.Reference_rule{Base: ptr8729024896, RuleBase: ptr8728690336, Default: system.Reference{}}

	ptr8729025024 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729321472 := &system.RuleBase{Optional: true}

	ptr8729334896 := &system.JsonString_rule{Base: ptr8729025024, RuleBase: ptr8729321472}

	ptr8729024512 := &system.Type{Base: ptr8729024640, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"rules": ptr8729060608, "type": ptr8729030336, "id": ptr8729030464, "description": ptr8729334896}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729176832 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729176704 := &system.Type{Base: ptr8729176832, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729022848 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729023104 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729023232 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729323840 := &system.RuleBase{Optional: true}

	ptr8729060688 := &system.Int_rule{Base: ptr8729023232, RuleBase: ptr8729323840, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729023360 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729324064 := &system.RuleBase{Optional: true}

	ptr8729060768 := &system.Int_rule{Base: ptr8729023360, RuleBase: ptr8729324064, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729025920 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729324288 := &system.RuleBase{Optional: true}

	ptr8729060848 := &system.Int_rule{Base: ptr8729025920, RuleBase: ptr8729324288, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729026048 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729324512 := &system.RuleBase{Optional: true}

	ptr8729060928 := &system.Int_rule{Base: ptr8729026048, RuleBase: ptr8729324512, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729022976 := &system.Type{Base: ptr8729023104, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729060688, "multipleOf": ptr8729060768, "minimum": ptr8729060848, "maximum": ptr8729060928}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729022720 := &system.Type{Base: ptr8729022848, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729022976}

	ptr8729023616 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729023488 := &system.Type{Base: ptr8729023616, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729023744}

	ptr8729021440 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729021696 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729021824 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729322528 := &system.RuleBase{Optional: true}

	ptr8729322464 := &system.Bool_rule{Base: ptr8729021824, RuleBase: ptr8729322528, Default: system.Bool{}}

	ptr8729021568 := &system.Type{Base: ptr8729021696, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729322464}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729021312 := &system.Type{Base: ptr8729021440, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729021568}

	ptr8729178624 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729178880 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729325344 := &system.RuleBase{Optional: true}

	ptr8729179008 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729030144 := &system.Reference_rule{Base: ptr8729179008, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729061808 := &system.Array_rule{Base: ptr8729178880, RuleBase: ptr8729325344, Items: ptr8729030144, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729179136 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729325664 := &system.RuleBase{Optional: true}

	ptr8729179264 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729030272 := &system.Reference_rule{Base: ptr8729179264, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729061888 := &system.Array_rule{Base: ptr8729179136, RuleBase: ptr8729325664, Items: ptr8729030272, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729179392 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729326016 := &system.RuleBase{Optional: true}

	ptr8729297952 := &system.String_rule{Base: ptr8729179392, RuleBase: ptr8729326016, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729179520 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729326368 := &system.RuleBase{Optional: true}

	ptr8729333200 := &system.JsonBool_rule{Base: ptr8729179520, RuleBase: ptr8729326368}

	ptr8729179648 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729326720 := &system.RuleBase{Optional: true}

	ptr8729333456 := &system.JsonBool_rule{Base: ptr8729179648, RuleBase: ptr8729326720}

	ptr8729179776 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729327040 := &system.RuleBase{Optional: true}

	ptr8729179904 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729333824 := &system.Rule_rule{Base: ptr8729179904, RuleBase: (*system.RuleBase)(nil)}

	ptr8729030528 := &system.Map_rule{Base: ptr8729179776, RuleBase: ptr8729327040, Items: ptr8729333824, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729180032 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8729327296 := &system.RuleBase{Optional: true}

	ptr8729334032 := &system.Type_rule{Base: ptr8729180032, RuleBase: ptr8729327296}

	ptr8729178752 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729325056 := &system.RuleBase{Optional: true}

	ptr8729332096 := &system.JsonBool_rule{Base: ptr8729178752, RuleBase: ptr8729325056}

	ptr8729178496 := &system.Type{Base: ptr8729178624, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"embed": ptr8729061808, "is": ptr8729061888, "native": ptr8729297952, "interface": ptr8729333200, "exclude": ptr8729333456, "fields": ptr8729030528, "rule": ptr8729334032, "basic": ptr8729332096}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "map", ptr8729026176)

	system.RegisterType("kego.io/system", "@string", ptr8729177216)

	system.RegisterType("kego.io/system", "@map", ptr8729025152)

	system.RegisterType("kego.io/system", "int", ptr8729022720)

	system.RegisterType("kego.io/system", "array", ptr8729023488)

	system.RegisterType("kego.io/system", "type", ptr8729178496)

	system.RegisterType("kego.io/system", "@type", ptr8729180160)

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729176704)

	system.RegisterType("kego.io/system", "@reference", ptr8729175296)

	system.RegisterType("kego.io/system", "bool", ptr8729021312)

	system.RegisterType("kego.io/system", "@int", ptr8729022976)

	system.RegisterType("kego.io/system", "@array", ptr8729023744)

	system.RegisterType("kego.io/system", "imports", ptr8729021952)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729176192)

	system.RegisterType("kego.io/system", "@bool", ptr8729021568)

	system.RegisterType("kego.io/system", "@number", ptr8729174016)

	system.RegisterType("kego.io/system", "base", ptr8729024512)

	system.RegisterType("kego.io/system", "reference", ptr8729175040)

	system.RegisterType("kego.io/system", "rule", ptr8729175680)

	system.RegisterType("kego.io/system", "@imports", ptr8729022464)

	system.RegisterType("kego.io/system", "@rule", ptr8729175936)

	system.RegisterType("kego.io/system", "number", ptr8729025792)

	system.RegisterType("kego.io/system", "@base", ptr8729021056)

	system.RegisterType("kego.io/system", "string", ptr8729176960)

}
