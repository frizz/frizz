package types

import (
	"kego.io/system"
)

func init() {

	ptr8729142400 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142656 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142784 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729324704 := &system.RuleBase{Optional: true}

	ptr8728842176 := &system.Reference_rule{Base: ptr8729142784, RuleBase: ptr8729324704, Default: system.Reference{}}

	ptr8729142528 := &system.Type{Base: ptr8729142656, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8728842176}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729142272 := &system.Type{Base: ptr8729142400, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729142528}

	ptr8728985728 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728988416 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728988672 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729322336 := &system.RuleBase{Optional: true}

	ptr8728834912 := &system.Int_rule{Base: ptr8728988672, RuleBase: ptr8729322336, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728988800 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729322592 := &system.RuleBase{Optional: true}

	ptr8728834992 := &system.Int_rule{Base: ptr8728988800, RuleBase: ptr8729322592, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728988928 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729322880 := &system.RuleBase{Optional: true}

	ptr8729322784 := &system.Bool_rule{Base: ptr8728988928, RuleBase: ptr8729322880, Default: system.Bool{Exists: true}}

	ptr8728988544 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729002432 := &system.Rule_rule{Base: ptr8728988544, RuleBase: (*system.RuleBase)(nil)}

	ptr8728988288 := &system.Type{Base: ptr8728988416, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"minItems": ptr8728834912, "maxItems": ptr8728834992, "uniqueItems": ptr8729322784, "items": ptr8729002432}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728985600 := &system.Type{Base: ptr8728985728, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8728988288}

	ptr8729144320 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729144576 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729145472 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729328032 := &system.RuleBase{Optional: true}

	ptr8729109360 := &system.String_rule{Base: ptr8729145472, RuleBase: ptr8729328032, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729145600 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729328608 := &system.RuleBase{Optional: true}

	ptr8729109536 := &system.String_rule{Base: ptr8729145600, RuleBase: ptr8729328608, Default: system.String{}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729144704 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729326688 := &system.RuleBase{Optional: true}

	ptr8729108832 := &system.String_rule{Base: ptr8729144704, RuleBase: ptr8729326688, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729144832 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729326976 := &system.RuleBase{Optional: true}

	ptr8729144960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729109008 := &system.String_rule{Base: ptr8729144960, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728831152 := &system.Array_rule{Base: ptr8729144832, RuleBase: ptr8729326976, Items: ptr8729109008, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729145088 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729327296 := &system.RuleBase{Optional: true}

	ptr8728831392 := &system.Int_rule{Base: ptr8729145088, RuleBase: ptr8729327296, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729145216 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729327552 := &system.RuleBase{Optional: true}

	ptr8728831712 := &system.Int_rule{Base: ptr8729145216, RuleBase: ptr8729327552, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729145344 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729327776 := &system.RuleBase{Optional: true}

	ptr8729109184 := &system.String_rule{Base: ptr8729145344, RuleBase: ptr8729327776, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729144448 := &system.Type{Base: ptr8729144576, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"pattern": ptr8729109360, "format": ptr8729109536, "default": ptr8729108832, "enum": ptr8728831152, "minLength": ptr8728831392, "maxLength": ptr8728831712, "equal": ptr8729109184}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729144192 := &system.Type{Base: ptr8729144320, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729144448}

	ptr8728992000 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728992128 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729326368 := &system.RuleBase{Optional: true}

	ptr8728835152 := &system.Int_rule{Base: ptr8728992128, RuleBase: ptr8729326368, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728992256 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729326592 := &system.RuleBase{Optional: true}

	ptr8728835232 := &system.Int_rule{Base: ptr8728992256, RuleBase: ptr8729326592, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728992384 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729326816 := &system.RuleBase{Optional: true}

	ptr8728835312 := &system.Int_rule{Base: ptr8728992384, RuleBase: ptr8729326816, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728992512 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729327040 := &system.RuleBase{Optional: true}

	ptr8728835392 := &system.Int_rule{Base: ptr8728992512, RuleBase: ptr8729327040, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728991872 := &system.Type{Base: ptr8728992000, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8728835152, "multipleOf": ptr8728835232, "minimum": ptr8728835312, "maximum": ptr8728835392}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729141376 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729141760 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729322624 := &system.RuleBase{Optional: true}

	ptr8728994176 := &system.Number_rule{Base: ptr8729141760, RuleBase: ptr8729322624, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141888 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729323104 := &system.RuleBase{Optional: true}

	ptr8729323008 := &system.Bool_rule{Base: ptr8729141888, RuleBase: ptr8729323104, Default: system.Bool{Exists: true}}

	ptr8729142016 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729323392 := &system.RuleBase{Optional: true}

	ptr8728994272 := &system.Number_rule{Base: ptr8729142016, RuleBase: ptr8729323392, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729142144 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729323872 := &system.RuleBase{Optional: true}

	ptr8729323744 := &system.Bool_rule{Base: ptr8729142144, RuleBase: ptr8729323872, Default: system.Bool{Exists: true}}

	ptr8729141504 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729321856 := &system.RuleBase{Optional: true}

	ptr8728993984 := &system.Number_rule{Base: ptr8729141504, RuleBase: ptr8729321856, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141632 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729322208 := &system.RuleBase{Optional: true}

	ptr8728994080 := &system.Number_rule{Base: ptr8729141632, RuleBase: ptr8729322208, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141248 := &system.Type{Base: ptr8729141376, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"minimum": ptr8728994176, "exclusiveMinimum": ptr8729323008, "maximum": ptr8728994272, "exclusiveMaximum": ptr8729323744, "default": ptr8728993984, "multipleOf": ptr8728994080}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729143040 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142912 := &system.Type{Base: ptr8729143040, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728993664 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728993536 := &system.Type{Base: ptr8728993664, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729141248}

	ptr8729143552 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143680 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729325536 := &system.RuleBase{Optional: true}

	ptr8729004416 := &system.JsonBool_rule{Base: ptr8729143680, RuleBase: ptr8729325536}

	ptr8729143808 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729325824 := &system.RuleBase{Optional: true}

	ptr8729004640 := &system.JsonString_rule{Base: ptr8729143808, RuleBase: ptr8729325824}

	ptr8729143424 := &system.Type{Base: ptr8729143552, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr8729004416, "selector": ptr8729004640}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728993024 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728993152 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729006640 := &system.Rule_rule{Base: ptr8728993152, RuleBase: (*system.RuleBase)(nil)}

	ptr8728993280 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729328000 := &system.RuleBase{Optional: true}

	ptr8728835472 := &system.Int_rule{Base: ptr8728993280, RuleBase: ptr8729328000, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728993408 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729328256 := &system.RuleBase{Optional: true}

	ptr8728835552 := &system.Int_rule{Base: ptr8728993408, RuleBase: ptr8729328256, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728992896 := &system.Type{Base: ptr8728993024, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr8729006640, "minItems": ptr8728835472, "maxItems": ptr8728835552}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728990976 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728991104 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728991232 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729108656 := &system.String_rule{Base: ptr8728991232, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728841536 := &system.Map_rule{Base: ptr8728991104, RuleBase: (*system.RuleBase)(nil), Items: ptr8729108656, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728990848 := &system.Type{Base: ptr8728990976, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr8728841536}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729144064 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143936 := &system.Type{Base: ptr8729144064, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729147520 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729147392 := &system.Type{Base: ptr8729147520, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728992768 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728992640 := &system.Type{Base: ptr8728992768, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8728992896}

	ptr8728990336 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728990592 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728990720 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729325056 := &system.RuleBase{Optional: true}

	ptr8729324992 := &system.Bool_rule{Base: ptr8728990720, RuleBase: ptr8729325056, Default: system.Bool{}}

	ptr8728990464 := &system.Type{Base: ptr8728990592, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729324992}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728990208 := &system.Type{Base: ptr8728990336, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8728990464}

	ptr8729143296 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143168 := &system.Type{Base: ptr8729143296, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728990080 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728989952 := &system.Type{Base: ptr8728990080, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729145856 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729147008 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728734048 := &system.RuleBase{Optional: true}

	ptr8729147136 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729008224 := &system.Rule_rule{Base: ptr8729147136, RuleBase: (*system.RuleBase)(nil)}

	ptr8728840960 := &system.Map_rule{Base: ptr8729147008, RuleBase: ptr8728734048, Items: ptr8729008224, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729147264 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8728734272 := &system.RuleBase{Optional: true}

	ptr8729008368 := &system.Type_rule{Base: ptr8729147264, RuleBase: ptr8728734272}

	ptr8729145984 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729329184 := &system.RuleBase{Optional: true}

	ptr8729006912 := &system.JsonBool_rule{Base: ptr8729145984, RuleBase: ptr8729329184}

	ptr8729146112 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729329472 := &system.RuleBase{Optional: true}

	ptr8729146240 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728840576 := &system.Reference_rule{Base: ptr8729146240, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728835872 := &system.Array_rule{Base: ptr8729146112, RuleBase: ptr8729329472, Items: ptr8728840576, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729146368 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728731808 := &system.RuleBase{Optional: true}

	ptr8729146496 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728840832 := &system.Reference_rule{Base: ptr8729146496, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728835952 := &system.Array_rule{Base: ptr8729146368, RuleBase: ptr8728731808, Items: ptr8728840832, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729146624 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728732992 := &system.RuleBase{Optional: true}

	ptr8729109712 := &system.String_rule{Base: ptr8729146624, RuleBase: ptr8728732992, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729146752 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728733472 := &system.RuleBase{Optional: true}

	ptr8729007840 := &system.JsonBool_rule{Base: ptr8729146752, RuleBase: ptr8728733472}

	ptr8729146880 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728733760 := &system.RuleBase{Optional: true}

	ptr8729008016 := &system.JsonBool_rule{Base: ptr8729146880, RuleBase: ptr8728733760}

	ptr8729145728 := &system.Type{Base: ptr8729145856, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"fields": ptr8728840960, "rule": ptr8729008368, "basic": ptr8729006912, "embed": ptr8728835872, "is": ptr8728835952, "native": ptr8729109712, "interface": ptr8729007840, "exclude": ptr8729008016}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728991744 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728991616 := &system.Type{Base: ptr8728991744, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8728991872}

	ptr8728989184 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728989696 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729324224 := &system.RuleBase{Optional: true}

	ptr8728989824 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729004448 := &system.Rule_rule{Base: ptr8728989824, RuleBase: (*system.RuleBase)(nil)}

	ptr8728835072 := &system.Array_rule{Base: ptr8728989696, RuleBase: ptr8729324224, Items: ptr8729004448, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728989312 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729172992 := &system.Reference_rule{Base: ptr8728989312, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728989440 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729323616 := &system.RuleBase{Optional: true}

	ptr8729173056 := &system.Reference_rule{Base: ptr8728989440, RuleBase: ptr8729323616, Default: system.Reference{}}

	ptr8728989568 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729323936 := &system.RuleBase{Optional: true}

	ptr8729004224 := &system.JsonString_rule{Base: ptr8728989568, RuleBase: ptr8729323936}

	ptr8728989056 := &system.Type{Base: ptr8728989184, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"rules": ptr8728835072, "type": ptr8729172992, "id": ptr8729173056, "description": ptr8729004224}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728991488 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728991360 := &system.Type{Base: ptr8728991488, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729143936)

	system.RegisterType("kego.io/system", "@base", ptr8728989952)

	system.RegisterType("kego.io/system", "@string", ptr8729144448)

	system.RegisterType("kego.io/system", "reference", ptr8729142272)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729143424)

	system.RegisterType("kego.io/system", "number", ptr8728993536)

	system.RegisterType("kego.io/system", "bool", ptr8728990208)

	system.RegisterType("kego.io/system", "array", ptr8728985600)

	system.RegisterType("kego.io/system", "@number", ptr8729141248)

	system.RegisterType("kego.io/system", "@bool", ptr8728990464)

	system.RegisterType("kego.io/system", "type", ptr8729145728)

	system.RegisterType("kego.io/system", "base", ptr8728989056)

	system.RegisterType("kego.io/system", "@array", ptr8728988288)

	system.RegisterType("kego.io/system", "string", ptr8729144192)

	system.RegisterType("kego.io/system", "@reference", ptr8729142528)

	system.RegisterType("kego.io/system", "@map", ptr8728992896)

	system.RegisterType("kego.io/system", "imports", ptr8728990848)

	system.RegisterType("kego.io/system", "@type", ptr8729147392)

	system.RegisterType("kego.io/system", "map", ptr8728992640)

	system.RegisterType("kego.io/system", "@rule", ptr8729143168)

	system.RegisterType("kego.io/system", "int", ptr8728991616)

	system.RegisterType("kego.io/system", "@int", ptr8728991872)

	system.RegisterType("kego.io/system", "rule", ptr8729142912)

	system.RegisterType("kego.io/system", "@imports", ptr8728991360)

}
