package types

import (
	"kego.io/system"
)

func init() {

	ptr8729147520 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729147392 := &system.Type{Base: ptr8729147520, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729142656 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142784 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729364320 := &system.RuleBase{Optional: true}

	ptr8728988480 := &system.Reference_rule{Base: ptr8729142784, RuleBase: ptr8729364320, Default: system.Reference{}}

	ptr8729142528 := &system.Type{Base: ptr8729142656, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8728988480}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729144320 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729144576 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729144704 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729366080 := &system.RuleBase{Optional: true}

	ptr8729108832 := &system.String_rule{Base: ptr8729144704, RuleBase: ptr8729366080, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729144832 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729366400 := &system.RuleBase{Optional: true}

	ptr8729144960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729292080 := &system.JsonString_rule{Base: ptr8729144960, RuleBase: (*system.RuleBase)(nil)}

	ptr8729028560 := &system.Array_rule{Base: ptr8729144832, RuleBase: ptr8729366400, Items: ptr8729292080, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729145088 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729366624 := &system.RuleBase{Optional: true}

	ptr8729028720 := &system.Int_rule{Base: ptr8729145088, RuleBase: ptr8729366624, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729145216 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729366848 := &system.RuleBase{Optional: true}

	ptr8729028800 := &system.Int_rule{Base: ptr8729145216, RuleBase: ptr8729366848, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729145344 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729367072 := &system.RuleBase{Optional: true}

	ptr8729109008 := &system.String_rule{Base: ptr8729145344, RuleBase: ptr8729367072, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729145472 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729367296 := &system.RuleBase{Optional: true}

	ptr8729109184 := &system.String_rule{Base: ptr8729145472, RuleBase: ptr8729367296, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729145600 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729367584 := &system.RuleBase{Optional: true}

	ptr8729109360 := &system.String_rule{Base: ptr8729145600, RuleBase: ptr8729367584, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729144448 := &system.Type{Base: ptr8729144576, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8729108832, "enum": ptr8729028560, "minLength": ptr8729028720, "maxLength": ptr8729028800, "equal": ptr8729109008, "pattern": ptr8729109184, "format": ptr8729109360}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729144192 := &system.Type{Base: ptr8729144320, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729144448}

	ptr8728998016 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728998272 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728998400 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728768512 := &system.RuleBase{Optional: true}

	ptr8728768448 := &system.Bool_rule{Base: ptr8728998400, RuleBase: ptr8728768512, Default: system.Bool{}}

	ptr8728998144 := &system.Type{Base: ptr8728998272, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr8728768448}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728997888 := &system.Type{Base: ptr8728998016, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8728998144}

	ptr8728999680 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729000448 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728770496 := &system.RuleBase{Optional: true}

	ptr8729028160 := &system.Int_rule{Base: ptr8729000448, RuleBase: ptr8728770496, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728999808 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728769824 := &system.RuleBase{Optional: true}

	ptr8729027920 := &system.Int_rule{Base: ptr8728999808, RuleBase: ptr8728769824, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729001216 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728770048 := &system.RuleBase{Optional: true}

	ptr8729028000 := &system.Int_rule{Base: ptr8729001216, RuleBase: ptr8728770048, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729001344 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728770272 := &system.RuleBase{Optional: true}

	ptr8729028080 := &system.Int_rule{Base: ptr8729001344, RuleBase: ptr8728770272, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728999552 := &system.Type{Base: ptr8728999680, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"maximum": ptr8729028160, "default": ptr8729027920, "multipleOf": ptr8729028000, "minimum": ptr8729028080}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729141376 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729141632 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729362560 := &system.RuleBase{Optional: true}

	ptr8728871200 := &system.Number_rule{Base: ptr8729141632, RuleBase: ptr8729362560, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141760 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729362784 := &system.RuleBase{Optional: true}

	ptr8728871296 := &system.Number_rule{Base: ptr8729141760, RuleBase: ptr8729362784, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141888 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729363072 := &system.RuleBase{Optional: true}

	ptr8729289216 := &system.JsonBool_rule{Base: ptr8729141888, RuleBase: ptr8729363072}

	ptr8729142016 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729363296 := &system.RuleBase{Optional: true}

	ptr8728871392 := &system.Number_rule{Base: ptr8729142016, RuleBase: ptr8729363296, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729142144 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729363584 := &system.RuleBase{Optional: true}

	ptr8729289600 := &system.JsonBool_rule{Base: ptr8729142144, RuleBase: ptr8729363584}

	ptr8729141504 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8728772512 := &system.RuleBase{Optional: true}

	ptr8728871104 := &system.Number_rule{Base: ptr8729141504, RuleBase: ptr8728772512, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729141248 := &system.Type{Base: ptr8729141376, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"multipleOf": ptr8728871200, "minimum": ptr8728871296, "exclusiveMinimum": ptr8729289216, "maximum": ptr8728871392, "exclusiveMaximum": ptr8729289600, "default": ptr8728871104}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729143040 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142912 := &system.Type{Base: ptr8729143040, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729001856 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729001728 := &system.Type{Base: ptr8729001856, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729141248}

	ptr8728998656 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728998784 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728998912 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729108656 := &system.String_rule{Base: ptr8728998912, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728990016 := &system.Map_rule{Base: ptr8728998784, RuleBase: (*system.RuleBase)(nil), Items: ptr8729108656, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728998528 := &system.Type{Base: ptr8728998656, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr8728990016}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729000064 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729000320 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728993792 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729288880 := &system.Rule_rule{Base: ptr8728993792, RuleBase: (*system.RuleBase)(nil)}

	ptr8728993920 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728765696 := &system.RuleBase{Optional: true}

	ptr8729027680 := &system.Int_rule{Base: ptr8728993920, RuleBase: ptr8728765696, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728996480 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728766080 := &system.RuleBase{Optional: true}

	ptr8729027760 := &system.Int_rule{Base: ptr8728996480, RuleBase: ptr8728766080, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728996608 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728766400 := &system.RuleBase{Optional: true}

	ptr8729289312 := &system.JsonBool_rule{Base: ptr8728996608, RuleBase: ptr8728766400}

	ptr8729000192 := &system.Type{Base: ptr8729000320, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr8729288880, "minItems": ptr8729027680, "maxItems": ptr8729027760, "uniqueItems": ptr8729289312}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728999936 := &system.Type{Base: ptr8729000064, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729000192}

	ptr8729144064 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143936 := &system.Type{Base: ptr8729144064, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729145856 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729146368 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729368736 := &system.RuleBase{Optional: true}

	ptr8729146496 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728989312 := &system.Reference_rule{Base: ptr8729146496, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729029120 := &system.Array_rule{Base: ptr8729146368, RuleBase: ptr8729368736, Items: ptr8728989312, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729146624 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729369056 := &system.RuleBase{Optional: true}

	ptr8729109536 := &system.String_rule{Base: ptr8729146624, RuleBase: ptr8729369056, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729146752 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729369376 := &system.RuleBase{Optional: true}

	ptr8729294352 := &system.JsonBool_rule{Base: ptr8729146752, RuleBase: ptr8729369376}

	ptr8729146880 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729369664 := &system.RuleBase{Optional: true}

	ptr8729294528 := &system.JsonBool_rule{Base: ptr8729146880, RuleBase: ptr8729369664}

	ptr8729147008 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729369952 := &system.RuleBase{Optional: true}

	ptr8729147136 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729294736 := &system.Rule_rule{Base: ptr8729147136, RuleBase: (*system.RuleBase)(nil)}

	ptr8728989568 := &system.Map_rule{Base: ptr8729147008, RuleBase: ptr8729369952, Items: ptr8729294736, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729147264 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8729370176 := &system.RuleBase{Optional: true}

	ptr8729294880 := &system.Type_rule{Base: ptr8729147264, RuleBase: ptr8729370176}

	ptr8729145984 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729368160 := &system.RuleBase{Optional: true}

	ptr8729293616 := &system.JsonBool_rule{Base: ptr8729145984, RuleBase: ptr8729368160}

	ptr8729146112 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729368448 := &system.RuleBase{Optional: true}

	ptr8729146240 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728989056 := &system.Reference_rule{Base: ptr8729146240, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729029040 := &system.Array_rule{Base: ptr8729146112, RuleBase: ptr8729368448, Items: ptr8728989056, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729145728 := &system.Type{Base: ptr8729145856, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"is": ptr8729029120, "native": ptr8729109536, "interface": ptr8729294352, "exclude": ptr8729294528, "fields": ptr8728989568, "rule": ptr8729294880, "basic": ptr8729293616, "embed": ptr8729029040}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728999168 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728999040 := &system.Type{Base: ptr8728999168, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729000704 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729000960 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729001088 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729292576 := &system.Rule_rule{Base: ptr8729001088, RuleBase: (*system.RuleBase)(nil)}

	ptr8729001472 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728771488 := &system.RuleBase{Optional: true}

	ptr8729028240 := &system.Int_rule{Base: ptr8729001472, RuleBase: ptr8728771488, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729001600 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728771776 := &system.RuleBase{Optional: true}

	ptr8729028320 := &system.Int_rule{Base: ptr8729001600, RuleBase: ptr8728771776, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729000832 := &system.Type{Base: ptr8729000960, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr8729292576, "minItems": ptr8729028240, "maxItems": ptr8729028320}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729000576 := &system.Type{Base: ptr8729000704, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729000832}

	ptr8729143296 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143168 := &system.Type{Base: ptr8729143296, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728999424 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728999296 := &system.Type{Base: ptr8728999424, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8728999552}

	ptr8728996864 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728997120 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728767136 := &system.RuleBase{Optional: true}

	ptr8728987904 := &system.Reference_rule{Base: ptr8728997120, RuleBase: ptr8728767136, Default: system.Reference{}}

	ptr8728997248 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8728767456 := &system.RuleBase{Optional: true}

	ptr8729289968 := &system.JsonString_rule{Base: ptr8728997248, RuleBase: ptr8728767456}

	ptr8728997376 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728767744 := &system.RuleBase{Optional: true}

	ptr8728997504 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729290192 := &system.Rule_rule{Base: ptr8728997504, RuleBase: (*system.RuleBase)(nil)}

	ptr8729027840 := &system.Array_rule{Base: ptr8728997376, RuleBase: ptr8728767744, Items: ptr8729290192, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728996992 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728987840 := &system.Reference_rule{Base: ptr8728996992, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728996736 := &system.Type{Base: ptr8728996864, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"id": ptr8728987904, "description": ptr8729289968, "rules": ptr8729027840, "type": ptr8728987840}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728997760 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728997632 := &system.Type{Base: ptr8728997760, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729142400 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729142272 := &system.Type{Base: ptr8729142400, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729142528}

	ptr8729143552 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729143680 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729365056 := &system.RuleBase{Optional: true}

	ptr8729290976 := &system.JsonBool_rule{Base: ptr8729143680, RuleBase: ptr8729365056}

	ptr8729143808 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729365344 := &system.RuleBase{Optional: true}

	ptr8729291200 := &system.JsonString_rule{Base: ptr8729143808, RuleBase: ptr8729365344}

	ptr8729143424 := &system.Type{Base: ptr8729143552, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr8729290976, "selector": ptr8729291200}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "type", ptr8729145728)

	system.RegisterType("kego.io/system", "int", ptr8728999296)

	system.RegisterType("kego.io/system", "@reference", ptr8729142528)

	system.RegisterType("kego.io/system", "bool", ptr8728997888)

	system.RegisterType("kego.io/system", "number", ptr8729001728)

	system.RegisterType("kego.io/system", "imports", ptr8728998528)

	system.RegisterType("kego.io/system", "array", ptr8728999936)

	system.RegisterType("kego.io/system", "@rule", ptr8729143168)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729143424)

	system.RegisterType("kego.io/system", "string", ptr8729144192)

	system.RegisterType("kego.io/system", "@number", ptr8729141248)

	system.RegisterType("kego.io/system", "@string", ptr8729144448)

	system.RegisterType("kego.io/system", "base", ptr8728996736)

	system.RegisterType("kego.io/system", "@map", ptr8729000832)

	system.RegisterType("kego.io/system", "@bool", ptr8728998144)

	system.RegisterType("kego.io/system", "@type", ptr8729147392)

	system.RegisterType("kego.io/system", "@int", ptr8728999552)

	system.RegisterType("kego.io/system", "@array", ptr8729000192)

	system.RegisterType("kego.io/system", "@imports", ptr8728999040)

	system.RegisterType("kego.io/system", "map", ptr8729000576)

	system.RegisterType("kego.io/system", "@base", ptr8728997632)

	system.RegisterType("kego.io/system", "reference", ptr8729142272)

	system.RegisterType("kego.io/system", "rule", ptr8729142912)

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729143936)

}
