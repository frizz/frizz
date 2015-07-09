package types

import (
	system "kego.io/system"
)

func init() {
	ptr34902153984 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902154112 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr34902323152 := &system.Rule_rule{Base: ptr34902154112, RuleBase: (*system.RuleBase)(nil)}
	ptr34902154240 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902268096 := &system.RuleBase{Optional: true}
	ptr34902176288 := &system.Int_rule{Base: ptr34902154240, RuleBase: ptr34902268096, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr34902154368 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902268416 := &system.RuleBase{Optional: true}
	ptr34902176368 := &system.Int_rule{Base: ptr34902154368, RuleBase: ptr34902268416, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr34902154496 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902268768 := &system.RuleBase{Optional: true}
	ptr34902324272 := &system.JsonBool_rule{Base: ptr34902154496, RuleBase: ptr34902268768}
	ptr34902153856 := &system.Type{Base: ptr34902153984, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr34902323152, "minItems": ptr34902176288, "maxItems": ptr34902176368, "uniqueItems": ptr34902324272}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902434176 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902434048 := &system.Type{Base: ptr34902434176, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902434688 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902434816 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr34902426752 := &system.RuleBase{Optional: true}
	ptr34902426720 := &system.Bool_rule{Base: ptr34902434816, RuleBase: ptr34902426752, Default: system.Bool{}}
	ptr34902434560 := &system.Type{Base: ptr34902434688, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr34902426720}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902435584 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902435456 := &system.Type{Base: ptr34902435584, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902436096 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902436224 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902428064 := &system.RuleBase{Optional: true}
	ptr34902176528 := &system.Int_rule{Base: ptr34902436224, RuleBase: ptr34902428064, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902436352 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902428352 := &system.RuleBase{Optional: true}
	ptr34902176608 := &system.Int_rule{Base: ptr34902436352, RuleBase: ptr34902428352, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902436480 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902428640 := &system.RuleBase{Optional: true}
	ptr34902176688 := &system.Int_rule{Base: ptr34902436480, RuleBase: ptr34902428640, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902436608 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902428928 := &system.RuleBase{Optional: true}
	ptr34902176768 := &system.Int_rule{Base: ptr34902436608, RuleBase: ptr34902428928, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902435968 := &system.Type{Base: ptr34902436096, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr34902176528, "multipleOf": ptr34902176608, "minimum": ptr34902176688, "maximum": ptr34902176768}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902437120 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902437248 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr34902454112 := &system.Rule_rule{Base: ptr34902437248, RuleBase: (*system.RuleBase)(nil)}
	ptr34902437376 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902429856 := &system.RuleBase{Optional: true}
	ptr34902176848 := &system.Int_rule{Base: ptr34902437376, RuleBase: ptr34902429856, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr34902437504 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902430176 := &system.RuleBase{Optional: true}
	ptr34902176928 := &system.Int_rule{Base: ptr34902437504, RuleBase: ptr34902430176, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr34902436992 := &system.Type{Base: ptr34902437120, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr34902454112, "minItems": ptr34902176848, "maxItems": ptr34902176928}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902438016 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902438144 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr34902430848 := &system.RuleBase{Optional: true}
	ptr34901372960 := &system.Number_rule{Base: ptr34902438144, RuleBase: ptr34902430848, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr34902438272 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr34902431136 := &system.RuleBase{Optional: true}
	ptr34901373056 := &system.Number_rule{Base: ptr34902438272, RuleBase: ptr34902431136, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr34902438400 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr34902431424 := &system.RuleBase{Optional: true}
	ptr34901373152 := &system.Number_rule{Base: ptr34902438400, RuleBase: ptr34902431424, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr34902438528 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902431776 := &system.RuleBase{Optional: true}
	ptr34902456896 := &system.JsonBool_rule{Base: ptr34902438528, RuleBase: ptr34902431776}
	ptr34902438656 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr34902432064 := &system.RuleBase{Optional: true}
	ptr34901373248 := &system.Number_rule{Base: ptr34902438656, RuleBase: ptr34902432064, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr34902438784 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902432416 := &system.RuleBase{Optional: true}
	ptr34902457520 := &system.JsonBool_rule{Base: ptr34902438784, RuleBase: ptr34902432416}
	ptr34902437888 := &system.Type{Base: ptr34902438016, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr34901372960, "multipleOf": ptr34901373056, "minimum": ptr34901373152, "exclusiveMinimum": ptr34902456896, "maximum": ptr34901373248, "exclusiveMaximum": ptr34902457520}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902439296 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902439424 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr34902433088 := &system.RuleBase{Optional: true}
	ptr34902404608 := &system.Reference_rule{Base: ptr34902439424, RuleBase: ptr34902433088, Default: system.Reference{}}
	ptr34902439168 := &system.Type{Base: ptr34902439296, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr34902404608}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902439936 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902439808 := &system.Type{Base: ptr34902439936, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902440704 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902440576 := &system.Type{Base: ptr34902440704, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902441216 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902622336 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34902592064 := &system.RuleBase{Optional: true}
	ptr34901739232 := &system.String_rule{Base: ptr34902622336, RuleBase: ptr34902592064, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902622464 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34902592448 := &system.RuleBase{Optional: true}
	ptr34901739408 := &system.String_rule{Base: ptr34902622464, RuleBase: ptr34902592448, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902441344 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34902590528 := &system.RuleBase{Optional: true}
	ptr34901738880 := &system.String_rule{Base: ptr34902441344, RuleBase: ptr34902590528, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902441472 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr34902590912 := &system.RuleBase{Optional: true}
	ptr34902441600 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr34902576576 := &system.JsonString_rule{Base: ptr34902441600, RuleBase: (*system.RuleBase)(nil)}
	ptr34902177168 := &system.Array_rule{Base: ptr34902441472, RuleBase: ptr34902590912, Items: ptr34902576576, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902441728 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902591200 := &system.RuleBase{Optional: true}
	ptr34902177328 := &system.Int_rule{Base: ptr34902441728, RuleBase: ptr34902591200, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902441856 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr34902591488 := &system.RuleBase{Optional: true}
	ptr34902177408 := &system.Int_rule{Base: ptr34902441856, RuleBase: ptr34902591488, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34902622208 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34902591776 := &system.RuleBase{Optional: true}
	ptr34901739056 := &system.String_rule{Base: ptr34902622208, RuleBase: ptr34902591776, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902441088 := &system.Type{Base: ptr34902441216, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"pattern": ptr34901739232, "format": ptr34901739408, "default": ptr34901738880, "enum": ptr34902177168, "minLength": ptr34902177328, "maxLength": ptr34902177408, "equal": ptr34901739056}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902624384 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902624256 := &system.Type{Base: ptr34902624384, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902153728 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902153600 := &system.Type{Base: ptr34902153728, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr34902153856}
	ptr34902154752 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902155136 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr34902425664 := &system.RuleBase{Optional: true}
	ptr34902325728 := &system.JsonString_rule{Base: ptr34902155136, RuleBase: ptr34902425664}
	ptr34902433792 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr34902426048 := &system.RuleBase{Optional: true}
	ptr34902433920 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr34902326192 := &system.Rule_rule{Base: ptr34902433920, RuleBase: (*system.RuleBase)(nil)}
	ptr34902176448 := &system.Array_rule{Base: ptr34902433792, RuleBase: ptr34902426048, Items: ptr34902326192, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902154880 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr34902294464 := &system.Reference_rule{Base: ptr34902154880, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr34902155008 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr34902269632 := &system.RuleBase{Optional: true}
	ptr34902401024 := &system.Reference_rule{Base: ptr34902155008, RuleBase: ptr34902269632, Default: system.Reference{}}
	ptr34902154624 := &system.Type{Base: ptr34902154752, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"description": ptr34902325728, "rules": ptr34902176448, "type": ptr34902294464, "id": ptr34902401024}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902434432 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902434304 := &system.Type{Base: ptr34902434432, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr34902434560}
	ptr34902435072 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902435200 := &system.Base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr34902435328 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34901738704 := &system.String_rule{Base: ptr34902435328, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902402496 := &system.Map_rule{Base: ptr34902435200, RuleBase: (*system.RuleBase)(nil), Items: ptr34901738704, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902434944 := &system.Type{Base: ptr34902435072, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr34902402496}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902435840 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902435712 := &system.Type{Base: ptr34902435840, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr34902435968}
	ptr34902436864 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902436736 := &system.Type{Base: ptr34902436864, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr34902436992}
	ptr34902437760 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902437632 := &system.Type{Base: ptr34902437760, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr34902437888}
	ptr34902439040 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902438912 := &system.Type{Base: ptr34902439040, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr34902439168}
	ptr34902439680 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902439552 := &system.Type{Base: ptr34902439680, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902440192 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902440320 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902589504 := &system.RuleBase{Optional: true}
	ptr34902574688 := &system.JsonBool_rule{Base: ptr34902440320, RuleBase: ptr34902589504}
	ptr34902440448 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr34902589856 := &system.RuleBase{Optional: true}
	ptr34902575008 := &system.JsonString_rule{Base: ptr34902440448, RuleBase: ptr34902589856}
	ptr34902440064 := &system.Type{Base: ptr34902440192, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr34902574688, "selector": ptr34902575008}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34902440960 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902440832 := &system.Type{Base: ptr34902440960, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr34902441088}
	ptr34902622720 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34902622976 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr34902593472 := &system.RuleBase{Optional: true}
	ptr34902623104 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr34902407104 := &system.Reference_rule{Base: ptr34902623104, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr34902177648 := &system.Array_rule{Base: ptr34902622976, RuleBase: ptr34902593472, Items: ptr34902407104, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902623232 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr34902593856 := &system.RuleBase{Optional: true}
	ptr34902623360 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr34902407232 := &system.Reference_rule{Base: ptr34902623360, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr34902177728 := &system.Array_rule{Base: ptr34902623232, RuleBase: ptr34902593856, Items: ptr34902407232, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902623488 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr34902594272 := &system.RuleBase{Optional: true}
	ptr34901739584 := &system.String_rule{Base: ptr34902623488, RuleBase: ptr34902594272, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr34902623616 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902594656 := &system.RuleBase{Optional: true}
	ptr34902580912 := &system.JsonBool_rule{Base: ptr34902623616, RuleBase: ptr34902594656}
	ptr34902623744 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902595008 := &system.RuleBase{Optional: true}
	ptr34902581232 := &system.JsonBool_rule{Base: ptr34902623744, RuleBase: ptr34902595008}
	ptr34902623872 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr34902595392 := &system.RuleBase{Optional: true}
	ptr34902624000 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr34902680000 := &system.Rule_rule{Base: ptr34902624000, RuleBase: (*system.RuleBase)(nil)}
	ptr34902407360 := &system.Map_rule{Base: ptr34902623872, RuleBase: ptr34902595392, Items: ptr34902680000, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr34902624128 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr34902595680 := &system.RuleBase{Optional: true}
	ptr34902680320 := &system.Type_rule{Base: ptr34902624128, RuleBase: ptr34902595680}
	ptr34902622848 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr34902593088 := &system.RuleBase{Optional: true}
	ptr34902579120 := &system.JsonBool_rule{Base: ptr34902622848, RuleBase: ptr34902593088}
	ptr34902622592 := &system.Type{Base: ptr34902622720, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"embed": ptr34902177648, "is": ptr34902177728, "native": ptr34901739584, "interface": ptr34902580912, "exclude": ptr34902581232, "fields": ptr34902407360, "rule": ptr34902680320, "basic": ptr34902579120}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/system", "@int", ptr34902435968)
	system.RegisterType("kego.io/system", "@rule", ptr34902439808)
	system.RegisterType("kego.io/system", "@ruleBase", ptr34902440576)
	system.RegisterType("kego.io/system", "@string", ptr34902441088)
	system.RegisterType("kego.io/system", "int", ptr34902435712)
	system.RegisterType("kego.io/system", "map", ptr34902436736)
	system.RegisterType("kego.io/system", "string", ptr34902440832)
	system.RegisterType("kego.io/system", "@imports", ptr34902435456)
	system.RegisterType("kego.io/system", "bool", ptr34902434304)
	system.RegisterType("kego.io/system", "ruleBase", ptr34902440064)
	system.RegisterType("kego.io/system", "@base", ptr34902434048)
	system.RegisterType("kego.io/system", "@bool", ptr34902434560)
	system.RegisterType("kego.io/system", "@map", ptr34902436992)
	system.RegisterType("kego.io/system", "@number", ptr34902437888)
	system.RegisterType("kego.io/system", "@type", ptr34902624256)
	system.RegisterType("kego.io/system", "array", ptr34902153600)
	system.RegisterType("kego.io/system", "base", ptr34902154624)
	system.RegisterType("kego.io/system", "reference", ptr34902438912)
	system.RegisterType("kego.io/system", "rule", ptr34902439552)
	system.RegisterType("kego.io/system", "@array", ptr34902153856)
	system.RegisterType("kego.io/system", "@reference", ptr34902439168)
	system.RegisterType("kego.io/system", "imports", ptr34902434944)
	system.RegisterType("kego.io/system", "number", ptr34902437632)
	system.RegisterType("kego.io/system", "type", ptr34902622592)
}
