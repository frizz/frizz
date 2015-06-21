package types

import (
	"kego.io/system"
)

func init() {

	ptr8729175168 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175424 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175552 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729320256 := &system.RuleBase{Optional: true}

	ptr8729025920 := &system.Reference_rule{Base: ptr8729175552, RuleBase: ptr8729320256, Default: system.Reference{}}

	ptr8729175296 := &system.Type{Base: ptr8729175424, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729025920}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729175040 := &system.Type{Base: ptr8729175168, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729175296}

	ptr8728887168 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729174144 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729174272 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729318272 := &system.RuleBase{Optional: true}

	ptr8729026656 := &system.Number_rule{Base: ptr8729174272, RuleBase: ptr8729318272, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174400 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729318496 := &system.RuleBase{Optional: true}

	ptr8729026752 := &system.Number_rule{Base: ptr8729174400, RuleBase: ptr8729318496, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174528 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729318720 := &system.RuleBase{Optional: true}

	ptr8729026848 := &system.Number_rule{Base: ptr8729174528, RuleBase: ptr8729318720, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174656 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729319008 := &system.RuleBase{Optional: true}

	ptr8729325184 := &system.JsonBool_rule{Base: ptr8729174656, RuleBase: ptr8729319008}

	ptr8729174784 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729319232 := &system.RuleBase{Optional: true}

	ptr8729027040 := &system.Number_rule{Base: ptr8729174784, RuleBase: ptr8729319232, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729174912 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729319520 := &system.RuleBase{Optional: true}

	ptr8729325536 := &system.JsonBool_rule{Base: ptr8729174912, RuleBase: ptr8729319520}

	ptr8729174016 := &system.Type{Base: ptr8729174144, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729026656, "multipleOf": ptr8729026752, "minimum": ptr8729026848, "exclusiveMinimum": ptr8729325184, "maximum": ptr8729027040, "exclusiveMaximum": ptr8729325536}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728886528 := &system.Type{Base: ptr8728887168, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729174016}

	ptr8729176832 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729176704 := &system.Type{Base: ptr8729176832, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728886016 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728886272 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729317280 := &system.RuleBase{Optional: true}

	ptr8729061008 := &system.Int_rule{Base: ptr8728886272, RuleBase: ptr8729317280, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728886400 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729317536 := &system.RuleBase{Optional: true}

	ptr8729061088 := &system.Int_rule{Base: ptr8728886400, RuleBase: ptr8729317536, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728886144 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729323840 := &system.Rule_rule{Base: ptr8728886144, RuleBase: (*system.RuleBase)(nil)}

	ptr8728885888 := &system.Type{Base: ptr8728886016, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"minItems": ptr8729061008, "maxItems": ptr8729061088, "items": ptr8729323840}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728883840 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728886656 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729316096 := &system.RuleBase{Optional: true}

	ptr8729060848 := &system.Int_rule{Base: ptr8728886656, RuleBase: ptr8729316096, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728886784 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729316320 := &system.RuleBase{Optional: true}

	ptr8729060928 := &system.Int_rule{Base: ptr8728886784, RuleBase: ptr8729316320, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728883968 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729315648 := &system.RuleBase{Optional: true}

	ptr8729060688 := &system.Int_rule{Base: ptr8728883968, RuleBase: ptr8729315648, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728884096 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729315872 := &system.RuleBase{Optional: true}

	ptr8729060768 := &system.Int_rule{Base: ptr8728884096, RuleBase: ptr8729315872, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728883712 := &system.Type{Base: ptr8728883840, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"minimum": ptr8729060848, "maximum": ptr8729060928, "default": ptr8729060688, "multipleOf": ptr8729060768}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728885376 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728879104 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729313568 := &system.RuleBase{Optional: true}

	ptr8728879232 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729321648 := &system.Rule_rule{Base: ptr8728879232, RuleBase: (*system.RuleBase)(nil)}

	ptr8729060608 := &system.Array_rule{Base: ptr8728879104, RuleBase: ptr8729313568, Items: ptr8729321648, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728885504 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729022144 := &system.Reference_rule{Base: ptr8728885504, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728885632 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728690336 := &system.RuleBase{Optional: true}

	ptr8729022272 := &system.Reference_rule{Base: ptr8728885632, RuleBase: ptr8728690336, Default: system.Reference{}}

	ptr8728885760 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729313280 := &system.RuleBase{Optional: true}

	ptr8729326704 := &system.JsonString_rule{Base: ptr8728885760, RuleBase: ptr8729313280}

	ptr8728885248 := &system.Type{Base: ptr8728885376, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"rules": ptr8729060608, "type": ptr8729022144, "id": ptr8729022272, "description": ptr8729326704}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729177088 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729177344 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729178368 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729316224 := &system.RuleBase{Optional: true}

	ptr8729297776 := &system.String_rule{Base: ptr8729178368, RuleBase: ptr8729316224, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177472 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729314528 := &system.RuleBase{Optional: true}

	ptr8729297248 := &system.String_rule{Base: ptr8729177472, RuleBase: ptr8729314528, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177600 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729314880 := &system.RuleBase{Optional: true}

	ptr8729177728 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729322624 := &system.JsonString_rule{Base: ptr8729177728, RuleBase: (*system.RuleBase)(nil)}

	ptr8729061328 := &system.Array_rule{Base: ptr8729177600, RuleBase: ptr8729314880, Items: ptr8729322624, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729177856 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729315104 := &system.RuleBase{Optional: true}

	ptr8729061488 := &system.Int_rule{Base: ptr8729177856, RuleBase: ptr8729315104, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729177984 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729315392 := &system.RuleBase{Optional: true}

	ptr8729061568 := &system.Int_rule{Base: ptr8729177984, RuleBase: ptr8729315392, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729178112 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729315616 := &system.RuleBase{Optional: true}

	ptr8729297424 := &system.String_rule{Base: ptr8729178112, RuleBase: ptr8729315616, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729178240 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729315904 := &system.RuleBase{Optional: true}

	ptr8729297600 := &system.String_rule{Base: ptr8729178240, RuleBase: ptr8729315904, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729177216 := &system.Type{Base: ptr8729177344, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"format": ptr8729297776, "default": ptr8729297248, "enum": ptr8729061328, "minLength": ptr8729061488, "maxLength": ptr8729061568, "equal": ptr8729297424, "pattern": ptr8729297600}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729176960 := &system.Type{Base: ptr8729177088, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729177216}

	ptr8728882176 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728882432 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728882560 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729314336 := &system.RuleBase{Optional: true}

	ptr8729314272 := &system.Bool_rule{Base: ptr8728882560, RuleBase: ptr8729314336, Default: system.Bool{}}

	ptr8728882304 := &system.Type{Base: ptr8728882432, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729314272}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728882048 := &system.Type{Base: ptr8728882176, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8728882304}

	ptr8729176320 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729176576 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729313696 := &system.RuleBase{Optional: true}

	ptr8729321680 := &system.JsonString_rule{Base: ptr8729176576, RuleBase: ptr8729313696}

	ptr8729176448 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729313376 := &system.RuleBase{Optional: true}

	ptr8729326912 := &system.JsonBool_rule{Base: ptr8729176448, RuleBase: ptr8729313376}

	ptr8729176192 := &system.Type{Base: ptr8729176320, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"selector": ptr8729321680, "optional": ptr8729326912}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728884352 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728884608 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728884992 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728689280 := &system.RuleBase{Optional: true}

	ptr8729060528 := &system.Int_rule{Base: ptr8728884992, RuleBase: ptr8728689280, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728885120 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728689568 := &system.RuleBase{Optional: true}

	ptr8729326048 := &system.JsonBool_rule{Base: ptr8728885120, RuleBase: ptr8728689568}

	ptr8728884736 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729325616 := &system.Rule_rule{Base: ptr8728884736, RuleBase: (*system.RuleBase)(nil)}

	ptr8728884864 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728689024 := &system.RuleBase{Optional: true}

	ptr8729060448 := &system.Int_rule{Base: ptr8728884864, RuleBase: ptr8728689024, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728884480 := &system.Type{Base: ptr8728884608, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"maxItems": ptr8729060528, "uniqueItems": ptr8729326048, "items": ptr8729325616, "minItems": ptr8729060448}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728884224 := &system.Type{Base: ptr8728884352, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8728884480}

	ptr8729180288 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729180160 := &system.Type{Base: ptr8729180288, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728882816 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728882944 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728883072 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729297072 := &system.String_rule{Base: ptr8728883072, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729021440 := &system.Map_rule{Base: ptr8728882944, RuleBase: (*system.RuleBase)(nil), Items: ptr8729297072, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728882688 := &system.Type{Base: ptr8728882816, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr8729021440}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729176064 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175936 := &system.Type{Base: ptr8729176064, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728883584 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728883456 := &system.Type{Base: ptr8728883584, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8728883712}

	ptr8729175808 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729175680 := &system.Type{Base: ptr8729175808, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728887040 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728886912 := &system.Type{Base: ptr8728887040, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8728885888}

	ptr8728881920 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728881792 := &system.Type{Base: ptr8728881920, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728883328 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728883200 := &system.Type{Base: ptr8728883328, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729178624 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729178752 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729316864 := &system.RuleBase{Optional: true}

	ptr8729323904 := &system.JsonBool_rule{Base: ptr8729178752, RuleBase: ptr8729316864}

	ptr8729178880 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729317152 := &system.RuleBase{Optional: true}

	ptr8729179008 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729021952 := &system.Reference_rule{Base: ptr8729179008, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729061808 := &system.Array_rule{Base: ptr8729178880, RuleBase: ptr8729317152, Items: ptr8729021952, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729179136 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729317472 := &system.RuleBase{Optional: true}

	ptr8729179264 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729022080 := &system.Reference_rule{Base: ptr8729179264, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729061888 := &system.Array_rule{Base: ptr8729179136, RuleBase: ptr8729317472, Items: ptr8729022080, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729179392 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729317824 := &system.RuleBase{Optional: true}

	ptr8729297952 := &system.String_rule{Base: ptr8729179392, RuleBase: ptr8729317824, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729179520 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729318176 := &system.RuleBase{Optional: true}

	ptr8729325008 := &system.JsonBool_rule{Base: ptr8729179520, RuleBase: ptr8729318176}

	ptr8729179648 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729318528 := &system.RuleBase{Optional: true}

	ptr8729325264 := &system.JsonBool_rule{Base: ptr8729179648, RuleBase: ptr8729318528}

	ptr8729179776 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729318848 := &system.RuleBase{Optional: true}

	ptr8729179904 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729325632 := &system.Rule_rule{Base: ptr8729179904, RuleBase: (*system.RuleBase)(nil)}

	ptr8729022336 := &system.Map_rule{Base: ptr8729179776, RuleBase: ptr8729318848, Items: ptr8729325632, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729180032 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8729319104 := &system.RuleBase{Optional: true}

	ptr8729325840 := &system.Type_rule{Base: ptr8729180032, RuleBase: ptr8729319104}

	ptr8729178496 := &system.Type{Base: ptr8729178624, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"basic": ptr8729323904, "embed": ptr8729061808, "is": ptr8729061888, "native": ptr8729297952, "interface": ptr8729325008, "exclude": ptr8729325264, "fields": ptr8729022336, "rule": ptr8729325840}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "bool", ptr8728882048)

	system.RegisterType("kego.io/system", "@reference", ptr8729175296)

	system.RegisterType("kego.io/system", "imports", ptr8728882688)

	system.RegisterType("kego.io/system", "reference", ptr8729175040)

	system.RegisterType("kego.io/system", "@number", ptr8729174016)

	system.RegisterType("kego.io/system", "array", ptr8728884224)

	system.RegisterType("kego.io/system", "@array", ptr8728884480)

	system.RegisterType("kego.io/system", "map", ptr8728886912)

	system.RegisterType("kego.io/system", "@imports", ptr8728883200)

	system.RegisterType("kego.io/system", "type", ptr8729178496)

	system.RegisterType("kego.io/system", "number", ptr8728886528)

	system.RegisterType("kego.io/system", "@map", ptr8728885888)

	system.RegisterType("kego.io/system", "@string", ptr8729177216)

	system.RegisterType("kego.io/system", "rule", ptr8729175680)

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729176704)

	system.RegisterType("kego.io/system", "@int", ptr8728883712)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729176192)

	system.RegisterType("kego.io/system", "@bool", ptr8728882304)

	system.RegisterType("kego.io/system", "@type", ptr8729180160)

	system.RegisterType("kego.io/system", "@rule", ptr8729175936)

	system.RegisterType("kego.io/system", "int", ptr8728883456)

	system.RegisterType("kego.io/system", "@base", ptr8728881792)

	system.RegisterType("kego.io/system", "base", ptr8728885248)

	system.RegisterType("kego.io/system", "string", ptr8729176960)

}
