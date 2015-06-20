package types

import (
	"kego.io/system"
)

func init() {

	ptr8728957440 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728957568 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728957696 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729207744 := &system.Rule_rule{Base: ptr8728957696, RuleBase: (*system.RuleBase)(nil)}

	ptr8728957824 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728734816 := &system.RuleBase{Optional: true}

	ptr8728835472 := &system.Int_rule{Base: ptr8728957824, RuleBase: ptr8728734816, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728957952 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728735136 := &system.RuleBase{Optional: true}

	ptr8728835552 := &system.Int_rule{Base: ptr8728957952, RuleBase: ptr8728735136, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8729256736 := &system.Type{Base: ptr8728957568, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729207744, "minItems": ptr8728835472, "maxItems": ptr8728835552}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729256576 := &system.Type{Base: ptr8728957440, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729256736}

	ptr8728960128 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729258976 := &system.Type{Base: ptr8728960128, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728958208 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728958464 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8728736800 := &system.RuleBase{Optional: true}

	ptr8728838432 := &system.Number_rule{Base: ptr8728958464, RuleBase: ptr8728736800, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728958592 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8728737184 := &system.RuleBase{Optional: true}

	ptr8728838528 := &system.Number_rule{Base: ptr8728958592, RuleBase: ptr8728737184, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728958720 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728737664 := &system.RuleBase{Optional: true}

	ptr8728737536 := &system.Bool_rule{Base: ptr8728958720, RuleBase: ptr8728737664, Default: system.Bool{Exists: true}}

	ptr8728958848 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8728737984 := &system.RuleBase{Optional: true}

	ptr8728838624 := &system.Number_rule{Base: ptr8728958848, RuleBase: ptr8728737984, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728958976 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728739680 := &system.RuleBase{Optional: true}

	ptr8728739040 := &system.Bool_rule{Base: ptr8728958976, RuleBase: ptr8728739680, Default: system.Bool{Exists: true}}

	ptr8728958336 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8728736352 := &system.RuleBase{Optional: true}

	ptr8728838336 := &system.Number_rule{Base: ptr8728958336, RuleBase: ptr8728736352, Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729257376 := &system.Type{Base: ptr8728958208, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"multipleOf": ptr8728838432, "minimum": ptr8728838528, "exclusiveMinimum": ptr8728737536, "maximum": ptr8728838624, "exclusiveMaximum": ptr8728739040, "default": ptr8728838336}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728954112 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728954624 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728733824 := &system.RuleBase{Optional: true}

	ptr8728834992 := &system.Int_rule{Base: ptr8728954624, RuleBase: ptr8728733824, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728954752 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728739456 := &system.RuleBase{Optional: true}

	ptr8728738688 := &system.Bool_rule{Base: ptr8728954752, RuleBase: ptr8728739456, Default: system.Bool{Exists: true}}

	ptr8728954368 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8728654528 := &system.Rule_rule{Base: ptr8728954368, RuleBase: (*system.RuleBase)(nil)}

	ptr8728954496 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728733568 := &system.RuleBase{Optional: true}

	ptr8728834912 := &system.Int_rule{Base: ptr8728954496, RuleBase: ptr8728733568, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}

	ptr8728964064 := &system.Type{Base: ptr8728954112, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"maxItems": ptr8728834992, "uniqueItems": ptr8728738688, "items": ptr8728654528, "minItems": ptr8728834912}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728956544 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728969024 := &system.Type{Base: ptr8728956544, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728956800 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728956928 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728731776 := &system.RuleBase{Optional: true}

	ptr8728835152 := &system.Int_rule{Base: ptr8728956928, RuleBase: ptr8728731776, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728957056 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728732480 := &system.RuleBase{Optional: true}

	ptr8728835232 := &system.Int_rule{Base: ptr8728957056, RuleBase: ptr8728732480, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728957184 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728732960 := &system.RuleBase{Optional: true}

	ptr8728835312 := &system.Int_rule{Base: ptr8728957184, RuleBase: ptr8728732960, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8728957312 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728733312 := &system.RuleBase{Optional: true}

	ptr8728835392 := &system.Int_rule{Base: ptr8728957312, RuleBase: ptr8728733312, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729256096 := &system.Type{Base: ptr8728956800, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728835152, "multipleOf": ptr8728835232, "minimum": ptr8728835312, "maximum": ptr8728835392}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728960384 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728960512 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729290912 := &system.RuleBase{Optional: true}

	ptr8729117024 := &system.String_rule{Base: ptr8728960512, RuleBase: ptr8729290912, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728960640 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729291168 := &system.RuleBase{Optional: true}

	ptr8728960768 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729117200 := &system.String_rule{Base: ptr8728960768, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728831152 := &system.Array_rule{Base: ptr8728960640, RuleBase: ptr8729291168, Items: ptr8729117200, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728960896 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729291456 := &system.RuleBase{Optional: true}

	ptr8728831392 := &system.Int_rule{Base: ptr8728960896, RuleBase: ptr8729291456, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729321472 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729291680 := &system.RuleBase{Optional: true}

	ptr8728831712 := &system.Int_rule{Base: ptr8729321472, RuleBase: ptr8729291680, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}

	ptr8729321600 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729291904 := &system.RuleBase{Optional: true}

	ptr8729117376 := &system.String_rule{Base: ptr8729321600, RuleBase: ptr8729291904, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729321728 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729292128 := &system.RuleBase{Optional: true}

	ptr8729117552 := &system.String_rule{Base: ptr8729321728, RuleBase: ptr8729292128, Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729321856 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729292608 := &system.RuleBase{Optional: true}

	ptr8729117728 := &system.String_rule{Base: ptr8729321856, RuleBase: ptr8729292608, Default: system.String{}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729259296 := &system.Type{Base: ptr8728960384, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729117024, "enum": ptr8728831152, "minLength": ptr8728831392, "maxLength": ptr8728831712, "equal": ptr8729117376, "pattern": ptr8729117552, "format": ptr8729117728}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728959232 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728959360 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728739616 := &system.RuleBase{Optional: true}

	ptr8729110720 := &system.Reference_rule{Base: ptr8728959360, RuleBase: ptr8728739616, Default: system.Reference{}}

	ptr8729258016 := &system.Type{Base: ptr8728959232, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729110720}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728952832 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728963424 := &system.Type{Base: ptr8728952832, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8728964064}

	ptr8728956672 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729255936 := &system.Type{Base: ptr8728956672, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729256096}

	ptr8728959744 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728959872 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729289728 := &system.RuleBase{Optional: true}

	ptr8729210512 := &system.JsonBool_rule{Base: ptr8728959872, RuleBase: ptr8729289728}

	ptr8728960000 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729290016 := &system.RuleBase{Optional: true}

	ptr8729210672 := &system.JsonString_rule{Base: ptr8728960000, RuleBase: ptr8729290016}

	ptr8729258816 := &system.Type{Base: ptr8728959744, Basic: true, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729210512, "selector": ptr8729210672}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728960256 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729259136 := &system.Type{Base: ptr8728960256, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729259296}

	ptr8728954880 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728955008 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729108928 := &system.Reference_rule{Base: ptr8728955008, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728955136 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728734848 := &system.RuleBase{Optional: true}

	ptr8729108992 := &system.Reference_rule{Base: ptr8728955136, RuleBase: ptr8728734848, Default: system.Reference{}}

	ptr8728955264 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8728735168 := &system.RuleBase{Optional: true}

	ptr8728656320 := &system.JsonString_rule{Base: ptr8728955264, RuleBase: ptr8728735168}

	ptr8728955392 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728735456 := &system.RuleBase{Optional: true}

	ptr8728955520 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8728656544 := &system.Rule_rule{Base: ptr8728955520, RuleBase: (*system.RuleBase)(nil)}

	ptr8728835072 := &system.Array_rule{Base: ptr8728955392, RuleBase: ptr8728735456, Items: ptr8728656544, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8728966144 := &system.Type{Base: ptr8728954880, Basic: true, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729108928, "id": ptr8729108992, "description": ptr8728656320, "rules": ptr8728835072}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728958080 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729257216 := &system.Type{Base: ptr8728958080, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729257376}

	ptr8728956160 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728956288 := &system.Base{Description: "Map of import name to path.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728956416 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729116848 := &system.String_rule{Base: ptr8728956416, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729110656 := &system.Map_rule{Base: ptr8728956288, RuleBase: (*system.RuleBase)(nil), Items: ptr8729116848, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8728968704 := &system.Type{Base: ptr8728956160, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8729110656}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728959488 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729258496 := &system.Type{Base: ptr8728959488, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728955776 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728955904 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728956032 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728736480 := &system.RuleBase{Optional: true}

	ptr8728736416 := &system.Bool_rule{Base: ptr8728956032, RuleBase: ptr8728736480, Default: system.Bool{}}

	ptr8728967424 := &system.Type{Base: ptr8728955904, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728736416}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728967104 := &system.Type{Base: ptr8728955776, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8728967424}

	ptr8728959104 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729257856 := &system.Type{Base: ptr8728959104, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729258016}

	ptr8728955648 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728966784 := &system.Type{Base: ptr8728955648, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728959616 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729258656 := &system.Type{Base: ptr8728959616, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729323648 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729259936 := &system.Type{Base: ptr8729323648, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729321984 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729322624 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729290048 := &system.RuleBase{Optional: true}

	ptr8729322752 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729108800 := &system.Reference_rule{Base: ptr8729322752, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728835952 := &system.Array_rule{Base: ptr8729322624, RuleBase: ptr8729290048, Items: ptr8729108800, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729323136 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729291360 := &system.RuleBase{Optional: true}

	ptr8729208368 := &system.JsonBool_rule{Base: ptr8729323136, RuleBase: ptr8729291360}

	ptr8729323264 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729291712 := &system.RuleBase{Optional: true}

	ptr8729323392 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729208800 := &system.Rule_rule{Base: ptr8729323392, RuleBase: (*system.RuleBase)(nil)}

	ptr8729108864 := &system.Map_rule{Base: ptr8729323264, RuleBase: ptr8729291712, Items: ptr8729208800, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729323520 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8729291968 := &system.RuleBase{Optional: true}

	ptr8729208992 := &system.Type_rule{Base: ptr8729323520, RuleBase: ptr8729291968}

	ptr8729322112 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729288928 := &system.RuleBase{Optional: true}

	ptr8729206912 := &system.JsonBool_rule{Base: ptr8729322112, RuleBase: ptr8729288928}

	ptr8729322496 := &system.Base{Description: "Type which this should extend", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729289696 := &system.RuleBase{Optional: true}

	ptr8729108672 := &system.Reference_rule{Base: ptr8729322496, RuleBase: ptr8729289696, Default: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}}

	ptr8729322880 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729290688 := &system.RuleBase{Optional: true}

	ptr8729116672 := &system.String_rule{Base: ptr8729322880, RuleBase: ptr8729290688, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729323008 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729291040 := &system.RuleBase{Optional: true}

	ptr8729208160 := &system.JsonBool_rule{Base: ptr8729323008, RuleBase: ptr8729291040}

	ptr8729322240 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729289280 := &system.RuleBase{Optional: true}

	ptr8729322368 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729108608 := &system.Reference_rule{Base: ptr8729322368, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}

	ptr8728835872 := &system.Array_rule{Base: ptr8729322240, RuleBase: ptr8729289280, Items: ptr8729108608, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729259776 := &system.Type{Base: ptr8729321984, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"is": ptr8728835952, "exclude": ptr8729208368, "fields": ptr8729108864, "rule": ptr8729208992, "basic": ptr8729206912, "extends": ptr8729108672, "native": ptr8729116672, "interface": ptr8729208160, "embed": ptr8728835872}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729258976)

	system.RegisterType("kego.io/system", "@array", ptr8728964064)

	system.RegisterType("kego.io/system", "base", ptr8728966144)

	system.RegisterType("kego.io/system", "imports", ptr8728968704)

	system.RegisterType("kego.io/system", "reference", ptr8729257856)

	system.RegisterType("kego.io/system", "@base", ptr8728966784)

	system.RegisterType("kego.io/system", "@type", ptr8729259936)

	system.RegisterType("kego.io/system", "@number", ptr8729257376)

	system.RegisterType("kego.io/system", "@imports", ptr8728969024)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729258816)

	system.RegisterType("kego.io/system", "string", ptr8729259136)

	system.RegisterType("kego.io/system", "number", ptr8729257216)

	system.RegisterType("kego.io/system", "@bool", ptr8728967424)

	system.RegisterType("kego.io/system", "int", ptr8729255936)

	system.RegisterType("kego.io/system", "bool", ptr8728967104)

	system.RegisterType("kego.io/system", "@rule", ptr8729258656)

	system.RegisterType("kego.io/system", "type", ptr8729259776)

	system.RegisterType("kego.io/system", "map", ptr8729256576)

	system.RegisterType("kego.io/system", "@int", ptr8729256096)

	system.RegisterType("kego.io/system", "@map", ptr8729256736)

	system.RegisterType("kego.io/system", "@string", ptr8729259296)

	system.RegisterType("kego.io/system", "@reference", ptr8729258016)

	system.RegisterType("kego.io/system", "array", ptr8728963424)

	system.RegisterType("kego.io/system", "rule", ptr8729258496)

}
