package types

import (
	"kego.io/system"
)

func init() {

	ptr8728967680 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729235840 := &system.Type{Base: ptr8728967680, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728979584 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729237280 := &system.Type{Base: ptr8728979584, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728962944 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963072 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963200 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729274688 := &system.Rule_rule{Base: ptr8728963200, RuleBase: (*system.RuleBase)(nil)}

	ptr8728963328 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729283520 := &system.RuleBase{Optional: true}

	ptr8729109040 := &system.Int_rule{Base: ptr8728963328, RuleBase: ptr8729283520, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728963456 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729283776 := &system.RuleBase{Optional: true}

	ptr8729109120 := &system.Int_rule{Base: ptr8728963456, RuleBase: ptr8729283776, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728963584 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729284064 := &system.RuleBase{Optional: true}

	ptr8729283968 := &system.Bool_rule{Base: ptr8728963584, RuleBase: ptr8729284064, Default: system.Bool{Exists: true}}

	ptr8729231680 := &system.Type{Base: ptr8728963072, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729274688, "minItems": ptr8729109040, "maxItems": ptr8729109120, "uniqueItems": ptr8729283968}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729231520 := &system.Type{Base: ptr8728962944, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729231680}

	ptr8728967936 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728968064 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728732960 := &system.RuleBase{Optional: true}

	ptr8729276976 := &system.JsonBool_rule{Base: ptr8728968064, RuleBase: ptr8728732960}

	ptr8728968192 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8728733408 := &system.RuleBase{Optional: true}

	ptr8729277136 := &system.JsonString_rule{Base: ptr8728968192, RuleBase: ptr8728733408}

	ptr8729236160 := &system.Type{Base: ptr8728967936, Basic: true, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729276976, "selector": ptr8729277136}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728968576 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728968832 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728734560 := &system.RuleBase{Optional: true}

	ptr8728968960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729166176 := &system.String_rule{Base: ptr8728968960, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729109920 := &system.Array_rule{Base: ptr8728968832, RuleBase: ptr8728734560, Items: ptr8729166176, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728969088 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728734848 := &system.RuleBase{Optional: true}

	ptr8729110080 := &system.Int_rule{Base: ptr8728969088, RuleBase: ptr8728734848, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728977408 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728731808 := &system.RuleBase{Optional: true}

	ptr8729110160 := &system.Int_rule{Base: ptr8728977408, RuleBase: ptr8728731808, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728977536 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728732544 := &system.RuleBase{Optional: true}

	ptr8729165824 := &system.String_rule{Base: ptr8728977536, RuleBase: ptr8728732544, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728977664 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728732992 := &system.RuleBase{Optional: true}

	ptr8729166352 := &system.String_rule{Base: ptr8728977664, RuleBase: ptr8728732992, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728977792 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728733696 := &system.RuleBase{Optional: true}

	ptr8729166528 := &system.String_rule{Base: ptr8728977792, RuleBase: ptr8728733696, Default: system.String{}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728968704 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728734304 := &system.RuleBase{Optional: true}

	ptr8729166704 := &system.String_rule{Base: ptr8728968704, RuleBase: ptr8728734304, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729236640 := &system.Type{Base: ptr8728968576, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"enum": ptr8729109920, "minLength": ptr8729110080, "maxLength": ptr8729110160, "equal": ptr8729165824, "pattern": ptr8729166352, "format": ptr8729166528, "default": ptr8729166704}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728964608 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728964736 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728964864 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729286720 := &system.RuleBase{Optional: true}

	ptr8729286656 := &system.Bool_rule{Base: ptr8728964864, RuleBase: ptr8729286720, Default: system.Bool{}}

	ptr8729232640 := &system.Type{Base: ptr8728964736, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729286656}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729232480 := &system.Type{Base: ptr8728964608, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729232640}

	ptr8728968320 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729236320 := &system.Type{Base: ptr8728968320, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728967424 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728967552 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729288000 := &system.RuleBase{Optional: true}

	ptr8728817344 := &system.Reference_rule{Base: ptr8728967552, RuleBase: ptr8729288000, Default: system.Reference{}}

	ptr8729235360 := &system.Type{Base: ptr8728967424, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728817344}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728967808 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729236000 := &system.Type{Base: ptr8728967808, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728966400 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728965504 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728966912 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729285920 := &system.RuleBase{Optional: true}

	ptr8729285824 := &system.Bool_rule{Base: ptr8728966912, RuleBase: ptr8729285920, Default: system.Bool{Exists: true}}

	ptr8728967040 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729286176 := &system.RuleBase{Optional: true}

	ptr8728838336 := &system.Number_rule{Base: ptr8728967040, RuleBase: ptr8729286176, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728967168 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729286624 := &system.RuleBase{Optional: true}

	ptr8729286528 := &system.Bool_rule{Base: ptr8728967168, RuleBase: ptr8729286624, Default: system.Bool{Exists: true}}

	ptr8728966528 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729284608 := &system.RuleBase{Optional: true}

	ptr8728838816 := &system.Number_rule{Base: ptr8728966528, RuleBase: ptr8729284608, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728966656 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729285024 := &system.RuleBase{Optional: true}

	ptr8728838912 := &system.Number_rule{Base: ptr8728966656, RuleBase: ptr8729285024, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728966784 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729285472 := &system.RuleBase{Optional: true}

	ptr8728838240 := &system.Number_rule{Base: ptr8728966784, RuleBase: ptr8729285472, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729234720 := &system.Type{Base: ptr8728965504, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMinimum": ptr8729285824, "maximum": ptr8728838336, "exclusiveMaximum": ptr8729286528, "default": ptr8728838816, "multipleOf": ptr8728838912, "minimum": ptr8728838240}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729234560 := &system.Type{Base: ptr8728966400, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729234720}

	ptr8728964480 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729232320 := &system.Type{Base: ptr8728964480, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728964992 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728965248 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728965376 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729166000 := &system.String_rule{Base: ptr8728965376, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728817536 := &system.Map_rule{Base: ptr8728965248, RuleBase: (*system.RuleBase)(nil), Items: ptr8729166000, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729233760 := &system.Type{Base: ptr8728964992, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728817536}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728962560 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728962688 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729281120 := &system.RuleBase{Optional: true}

	ptr8729109280 := &system.Int_rule{Base: ptr8728962688, RuleBase: ptr8729281120, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728962816 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729281344 := &system.RuleBase{Optional: true}

	ptr8729109360 := &system.Int_rule{Base: ptr8728962816, RuleBase: ptr8729281344, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728965120 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729281568 := &system.RuleBase{Optional: true}

	ptr8729109440 := &system.Int_rule{Base: ptr8728965120, RuleBase: ptr8729281568, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728965632 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729281792 := &system.RuleBase{Optional: true}

	ptr8729109520 := &system.Int_rule{Base: ptr8728965632, RuleBase: ptr8729281792, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729233920 := &system.Type{Base: ptr8728962560, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729109280, "multipleOf": ptr8729109360, "minimum": ptr8729109440, "maximum": ptr8729109520}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728977920 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728978048 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728734720 := &system.RuleBase{Optional: true}

	ptr8729273632 := &system.JsonBool_rule{Base: ptr8728978048, RuleBase: ptr8728734720}

	ptr8728978432 := &system.Base{Description: "Type which this should extend", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728735488 := &system.RuleBase{Optional: true}

	ptr8728816512 := &system.Reference_rule{Base: ptr8728978432, RuleBase: ptr8728735488, Default: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}}

	ptr8728978816 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728736384 := &system.RuleBase{Optional: true}

	ptr8729166880 := &system.String_rule{Base: ptr8728978816, RuleBase: ptr8728736384, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728979200 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728737280 := &system.RuleBase{Optional: true}

	ptr8728979328 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729275664 := &system.Rule_rule{Base: ptr8728979328, RuleBase: (*system.RuleBase)(nil)}

	ptr8728816704 := &system.Map_rule{Base: ptr8728979200, RuleBase: ptr8728737280, Items: ptr8729275664, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728979456 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8728737504 := &system.RuleBase{Optional: true}

	ptr8729275872 := &system.Type_rule{Base: ptr8728979456, RuleBase: ptr8728737504}

	ptr8728978176 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728735072 := &system.RuleBase{Optional: true}

	ptr8728978304 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728816448 := &system.Reference_rule{Base: ptr8728978304, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729110400 := &system.Array_rule{Base: ptr8728978176, RuleBase: ptr8728735072, Items: ptr8728816448, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728978560 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728735776 := &system.RuleBase{Optional: true}

	ptr8728978688 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728816640 := &system.Reference_rule{Base: ptr8728978688, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8729110480 := &system.Array_rule{Base: ptr8728978560, RuleBase: ptr8728735776, Items: ptr8728816640, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728978944 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728736704 := &system.RuleBase{Optional: true}

	ptr8729275024 := &system.JsonBool_rule{Base: ptr8728978944, RuleBase: ptr8728736704}

	ptr8728979072 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728736992 := &system.RuleBase{Optional: true}

	ptr8729275312 := &system.JsonBool_rule{Base: ptr8728979072, RuleBase: ptr8728736992}

	ptr8729237120 := &system.Type{Base: ptr8728977920, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8729273632, "extends": ptr8728816512, "native": ptr8729166880, "fields": ptr8728816704, "rule": ptr8729275872, "embed": ptr8729110400, "is": ptr8729110480, "interface": ptr8729275024, "exclude": ptr8729275312}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728961024 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729231360 := &system.Type{Base: ptr8728961024, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728963712 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963840 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729140992 := &system.Reference_rule{Base: ptr8728963840, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728963968 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729285120 := &system.RuleBase{Optional: true}

	ptr8729141056 := &system.Reference_rule{Base: ptr8728963968, RuleBase: ptr8729285120, Default: system.Reference{}}

	ptr8728964096 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729285440 := &system.RuleBase{Optional: true}

	ptr8729275776 := &system.JsonString_rule{Base: ptr8728964096, RuleBase: ptr8729285440}

	ptr8728964224 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729285728 := &system.RuleBase{Optional: true}

	ptr8728964352 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729276000 := &system.Rule_rule{Base: ptr8728964352, RuleBase: (*system.RuleBase)(nil)}

	ptr8729109200 := &system.Array_rule{Base: ptr8728964224, RuleBase: ptr8729285728, Items: ptr8729276000, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729232160 := &system.Type{Base: ptr8728963712, Basic: true, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729140992, "id": ptr8729141056, "description": ptr8729275776, "rules": ptr8729109200}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728962304 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729233120 := &system.Type{Base: ptr8728962304, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729233920}

	ptr8728965888 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728966016 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729273824 := &system.Rule_rule{Base: ptr8728966016, RuleBase: (*system.RuleBase)(nil)}

	ptr8728966144 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729283104 := &system.RuleBase{Optional: true}

	ptr8729109600 := &system.Int_rule{Base: ptr8728966144, RuleBase: ptr8729283104, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728966272 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729283360 := &system.RuleBase{Optional: true}

	ptr8729109680 := &system.Int_rule{Base: ptr8728966272, RuleBase: ptr8729283360, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729233440 := &system.Type{Base: ptr8728965888, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729273824, "minItems": ptr8729109600, "maxItems": ptr8729109680}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728968448 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729236480 := &system.Type{Base: ptr8728968448, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729236640}

	ptr8728967296 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729235200 := &system.Type{Base: ptr8728967296, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729235360}

	ptr8728965760 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729233280 := &system.Type{Base: ptr8728965760, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729233440}

	system.RegisterType("kego.io/system", "rule", ptr8729235840)

	system.RegisterType("kego.io/system", "@array", ptr8729231680)

	system.RegisterType("kego.io/system", "@int", ptr8729233920)

	system.RegisterType("kego.io/system", "map", ptr8729233280)

	system.RegisterType("kego.io/system", "@type", ptr8729237280)

	system.RegisterType("kego.io/system", "@rule", ptr8729236000)

	system.RegisterType("kego.io/system", "number", ptr8729234560)

	system.RegisterType("kego.io/system", "imports", ptr8729233760)

	system.RegisterType("kego.io/system", "@imports", ptr8729231360)

	system.RegisterType("kego.io/system", "string", ptr8729236480)

	system.RegisterType("kego.io/system", "reference", ptr8729235200)

	system.RegisterType("kego.io/system", "int", ptr8729233120)

	system.RegisterType("kego.io/system", "array", ptr8729231520)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729236160)

	system.RegisterType("kego.io/system", "@string", ptr8729236640)

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729236320)

	system.RegisterType("kego.io/system", "@bool", ptr8729232640)

	system.RegisterType("kego.io/system", "@base", ptr8729232320)

	system.RegisterType("kego.io/system", "base", ptr8729232160)

	system.RegisterType("kego.io/system", "@number", ptr8729234720)

	system.RegisterType("kego.io/system", "bool", ptr8729232480)

	system.RegisterType("kego.io/system", "@reference", ptr8729235360)

	system.RegisterType("kego.io/system", "type", ptr8729237120)

	system.RegisterType("kego.io/system", "@map", ptr8729233440)

}
