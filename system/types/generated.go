package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728591744 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings", Rules: []system.Rule(nil)}
	ptr8728846464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match", Rules: []system.Rule(nil)}
	ptr8728395648 := &system.RuleBase{Optional: true}
	ptr8728723984 := &system.String_rule{Base: ptr8728846464, RuleBase: ptr8728395648, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728846592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to", Rules: []system.Rule(nil)}
	ptr8728387712 := &system.RuleBase{Optional: true}
	ptr8728724160 := &system.String_rule{Base: ptr8728846592, RuleBase: ptr8728387712, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728846720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil)}
	ptr8728388032 := &system.RuleBase{Optional: true}
	ptr8728723456 := &system.String_rule{Base: ptr8728846720, RuleBase: ptr8728388032, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728591872 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728960192 := &system.RuleBase{Optional: true}
	ptr8728723632 := &system.String_rule{Base: ptr8728591872, RuleBase: ptr8728960192, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728592000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil)}
	ptr8728960512 := &system.RuleBase{Optional: true}
	ptr8728592128 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729024048 := &system.JsonString_rule{Base: ptr8728592128, RuleBase: (*system.RuleBase)(nil)}
	ptr8728863840 := &system.Array_rule{Base: ptr8728592000, RuleBase: ptr8728960512, Items: ptr8729024048, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728592256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil)}
	ptr8728960736 := &system.RuleBase{Optional: true}
	ptr8728864000 := &system.Int_rule{Base: ptr8728592256, RuleBase: ptr8728960736, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728846336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil)}
	ptr8728960960 := &system.RuleBase{Optional: true}
	ptr8728864080 := &system.Int_rule{Base: ptr8728846336, RuleBase: ptr8728960960, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728794256 := &system.Type{Base: ptr8728591744, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"equal": ptr8728723984, "pattern": ptr8728724160, "format": ptr8728723456, "default": ptr8728723632, "enum": ptr8728863840, "minLength": ptr8728864000, "maxLength": ptr8728864080}}
	ptr8728585472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers", Rules: []system.Rule(nil)}
	ptr8728585600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728953312 := &system.RuleBase{Optional: true}
	ptr8728863200 := &system.Int_rule{Base: ptr8728585600, RuleBase: ptr8728953312, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728953536 := &system.RuleBase{Optional: true}
	ptr8728863280 := &system.Int_rule{Base: ptr8728585728, RuleBase: ptr8728953536, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728953760 := &system.RuleBase{Optional: true}
	ptr8728863360 := &system.Int_rule{Base: ptr8728585856, RuleBase: ptr8728953760, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585984 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728953984 := &system.RuleBase{Optional: true}
	ptr8728863440 := &system.Int_rule{Base: ptr8728585984, RuleBase: ptr8728953984, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728788992 := &system.Type{Base: ptr8728585472, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728863200, "multipleOf": ptr8728863280, "minimum": ptr8728863360, "maximum": ptr8728863440}}
	ptr8728588544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package.", Rules: []system.Rule(nil)}
	ptr8728588672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of path to alias.", Rules: []system.Rule(nil)}
	ptr8728588800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728723808 := &system.String_rule{Base: ptr8728588800, RuleBase: (*system.RuleBase)(nil), Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728563008 := &system.Map_rule{Base: ptr8728588672, RuleBase: (*system.RuleBase)(nil), Items: ptr8728723808, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728790336 := &system.Type{Base: ptr8728588544, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728563008}}
	ptr8728586240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps", Rules: []system.Rule(nil)}
	ptr8728586368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil)}
	ptr8729020016 := &system.Rule_rule{Base: ptr8728586368, RuleBase: (*system.RuleBase)(nil)}
	ptr8728589312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728954944 := &system.RuleBase{Optional: true}
	ptr8728863520 := &system.Int_rule{Base: ptr8728589312, RuleBase: ptr8728954944, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728589440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728955232 := &system.RuleBase{Optional: true}
	ptr8728863600 := &system.Int_rule{Base: ptr8728589440, RuleBase: ptr8728955232, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728790896 := &system.Type{Base: ptr8728586240, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729020016, "minItems": ptr8728863520, "maxItems": ptr8728863600}}
	ptr8728848384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type", Rules: []system.Rule(nil)}
	ptr8728792688 := &system.Type{Base: ptr8728848384, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728591488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil)}
	ptr8728791904 := &system.Type{Base: ptr8728591488, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728846848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728846976 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil)}
	ptr8728388960 := &system.RuleBase{Optional: true}
	ptr8729019376 := &system.JsonBool_rule{Base: ptr8728846976, RuleBase: ptr8728388960}
	ptr8728847104 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil)}
	ptr8728389376 := &system.RuleBase{Optional: true}
	ptr8728847232 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728562560 := &system.Reference_rule{Base: ptr8728847232, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728864320 := &system.Array_rule{Base: ptr8728847104, RuleBase: ptr8728389376, Items: ptr8728562560, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728847360 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support", Rules: []system.Rule(nil)}
	ptr8728389760 := &system.RuleBase{Optional: true}
	ptr8728847488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728562688 := &system.Reference_rule{Base: ptr8728847488, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728864400 := &system.Array_rule{Base: ptr8728847360, RuleBase: ptr8728389760, Items: ptr8728562688, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728847616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil)}
	ptr8728390112 := &system.RuleBase{Optional: true}
	ptr8728724336 := &system.String_rule{Base: ptr8728847616, RuleBase: ptr8728390112, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}}
	ptr8728847744 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?", Rules: []system.Rule(nil)}
	ptr8728390432 := &system.RuleBase{Optional: true}
	ptr8729020480 := &system.JsonBool_rule{Base: ptr8728847744, RuleBase: ptr8728390432}
	ptr8728847872 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil)}
	ptr8728390720 := &system.RuleBase{Optional: true}
	ptr8729020704 := &system.JsonBool_rule{Base: ptr8728847872, RuleBase: ptr8728390720}
	ptr8728848000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type", Rules: []system.Rule(nil)}
	ptr8728391008 := &system.RuleBase{Optional: true}
	ptr8728848128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729021008 := &system.Rule_rule{Base: ptr8728848128, RuleBase: (*system.RuleBase)(nil)}
	ptr8728562816 := &system.Map_rule{Base: ptr8728848000, RuleBase: ptr8728391008, Items: ptr8729021008, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728848256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil)}
	ptr8728391232 := &system.RuleBase{Optional: true}
	ptr8729021264 := &system.Type_rule{Base: ptr8728848256, RuleBase: ptr8728391232}
	ptr8728794928 := &system.Type{Base: ptr8728846848, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8729019376, "embed": ptr8728864320, "is": ptr8728864400, "native": ptr8728724336, "interface": ptr8729020480, "exclude": ptr8729020704, "fields": ptr8728562816, "rule": ptr8729021264}}
	ptr8728586112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type.", Rules: []system.Rule(nil)}
	ptr8728789216 := &system.Type{Base: ptr8728586112, Rule: ptr8728790896, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728588928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil)}
	ptr8728790448 := &system.Type{Base: ptr8728588928, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728588032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base", Rules: []system.Rule(nil)}
	ptr8728789888 := &system.Type{Base: ptr8728588032, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728589696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers", Rules: []system.Rule(nil)}
	ptr8728589184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728956032 := &system.RuleBase{Optional: true}
	ptr8728567904 := &system.Number_rule{Base: ptr8728589184, RuleBase: ptr8728956032, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728589824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728956288 := &system.RuleBase{Optional: true}
	ptr8728568576 := &system.Number_rule{Base: ptr8728589824, RuleBase: ptr8728956288, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728589952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728956544 := &system.RuleBase{Optional: true}
	ptr8728568096 := &system.Number_rule{Base: ptr8728589952, RuleBase: ptr8728956544, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728590080 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil)}
	ptr8728956864 := &system.RuleBase{Optional: true}
	ptr8729021296 := &system.JsonBool_rule{Base: ptr8728590080, RuleBase: ptr8728956864}
	ptr8728590208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728957152 := &system.RuleBase{Optional: true}
	ptr8728568288 := &system.Number_rule{Base: ptr8728590208, RuleBase: ptr8728957152, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}}
	ptr8728590336 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil)}
	ptr8728957472 := &system.RuleBase{Optional: true}
	ptr8729021696 := &system.JsonBool_rule{Base: ptr8728590336, RuleBase: ptr8728957472}
	ptr8728792464 := &system.Type{Base: ptr8728589696, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728567904, "multipleOf": ptr8728568576, "minimum": ptr8728568096, "exclusiveMinimum": ptr8729021296, "maximum": ptr8728568288, "exclusiveMaximum": ptr8729021696}}
	ptr8728586624 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays", Rules: []system.Rule(nil)}
	ptr8728586752 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil)}
	ptr8729021088 := &system.Rule_rule{Base: ptr8728586752, RuleBase: (*system.RuleBase)(nil)}
	ptr8728586880 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728955616 := &system.RuleBase{Optional: true}
	ptr8728862960 := &system.Int_rule{Base: ptr8728586880, RuleBase: ptr8728955616, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728955872 := &system.RuleBase{Optional: true}
	ptr8728863040 := &system.Int_rule{Base: ptr8728587008, RuleBase: ptr8728955872, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587136 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique", Rules: []system.Rule(nil)}
	ptr8728956160 := &system.RuleBase{Optional: true}
	ptr8729021520 := &system.JsonBool_rule{Base: ptr8728587136, RuleBase: ptr8728956160}
	ptr8728789440 := &system.Type{Base: ptr8728586624, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729021088, "minItems": ptr8728862960, "maxItems": ptr8728863040, "uniqueItems": ptr8729021520}}
	ptr8728590464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil)}
	ptr8728590592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references", Rules: []system.Rule(nil)}
	ptr8728590720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728958336 := &system.RuleBase{Optional: true}
	ptr8728562880 := &system.Reference_rule{Base: ptr8728590720, RuleBase: ptr8728958336, Default: system.Reference{}}
	ptr8728791344 := &system.Type{Base: ptr8728590592, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728562880}}
	ptr8728791008 := &system.Type{Base: ptr8728590464, Rule: ptr8728791344, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728590976 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil)}
	ptr8728791568 := &system.Type{Base: ptr8728590976, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728591104 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type.", Rules: []system.Rule(nil)}
	ptr8728591232 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil)}
	ptr8728959168 := &system.RuleBase{Optional: true}
	ptr8729023040 := &system.JsonBool_rule{Base: ptr8728591232, RuleBase: ptr8728959168}
	ptr8728591360 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil)}
	ptr8728959456 := &system.RuleBase{Optional: true}
	ptr8729023216 := &system.JsonString_rule{Base: ptr8728591360, RuleBase: ptr8728959456}
	ptr8728791680 := &system.Type{Base: ptr8728591104, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729023040, "selector": ptr8729023216}}
	ptr8728586496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type", Rules: []system.Rule(nil)}
	ptr8728789328 := &system.Type{Base: ptr8728586496, Rule: ptr8728789440, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728587264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728587776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil)}
	ptr8728957504 := &system.RuleBase{Optional: true}
	ptr8728587904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729022400 := &system.Rule_rule{Base: ptr8728587904, RuleBase: (*system.RuleBase)(nil)}
	ptr8728863120 := &system.Array_rule{Base: ptr8728587776, RuleBase: ptr8728957504, Items: ptr8729022400, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728587392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object.", Rules: []system.Rule(nil)}
	ptr8728936064 := &system.Reference_rule{Base: ptr8728587392, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728587520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id.", Rules: []system.Rule(nil)}
	ptr8728956896 := &system.RuleBase{Optional: true}
	ptr8728936128 := &system.Reference_rule{Base: ptr8728587520, RuleBase: ptr8728956896, Default: system.Reference{}}
	ptr8728587648 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer", Rules: []system.Rule(nil)}
	ptr8728957216 := &system.RuleBase{Optional: true}
	ptr8729022176 := &system.JsonString_rule{Base: ptr8728587648, RuleBase: ptr8728957216}
	ptr8728789664 := &system.Type{Base: ptr8728587264, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"rules": ptr8728863120, "type": ptr8728936064, "id": ptr8728936128, "description": ptr8729022176}}
	ptr8728590848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil)}
	ptr8728791456 := &system.Type{Base: ptr8728590848, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil)}
	ptr8728589568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type", Rules: []system.Rule(nil)}
	ptr8728792128 := &system.Type{Base: ptr8728589568, Rule: ptr8728792464, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728588288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools", Rules: []system.Rule(nil)}
	ptr8728588416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null", Rules: []system.Rule(nil)}
	ptr8728958272 := &system.RuleBase{Optional: true}
	ptr8728958208 := &system.Bool_rule{Base: ptr8728588416, RuleBase: ptr8728958272, Default: system.Bool{}}
	ptr8728790224 := &system.Type{Base: ptr8728588288, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728958208}}
	ptr8728589056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type", Rules: []system.Rule(nil)}
	ptr8728790672 := &system.Type{Base: ptr8728589056, Rule: ptr8728788992, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728591616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type", Rules: []system.Rule(nil)}
	ptr8728792016 := &system.Type{Base: ptr8728591616, Rule: ptr8728794256, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728588160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type", Rules: []system.Rule(nil)}
	ptr8728790112 := &system.Type{Base: ptr8728588160, Rule: ptr8728790224, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil)}
	system.RegisterType("kego.io/system", "@base", ptr8728789888)
	system.RegisterType("kego.io/system", "reference", ptr8728791008)
	system.RegisterType("kego.io/system", "@rule", ptr8728791568)
	system.RegisterType("kego.io/system", "array", ptr8728789328)
	system.RegisterType("kego.io/system", "@type", ptr8728792688)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728791904)
	system.RegisterType("kego.io/system", "type", ptr8728794928)
	system.RegisterType("kego.io/system", "rule", ptr8728791456)
	system.RegisterType("kego.io/system", "number", ptr8728792128)
	system.RegisterType("kego.io/system", "bool", ptr8728790112)
	system.RegisterType("kego.io/system", "map", ptr8728789216)
	system.RegisterType("kego.io/system", "@imports", ptr8728790448)
	system.RegisterType("kego.io/system", "@number", ptr8728792464)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728791680)
	system.RegisterType("kego.io/system", "@bool", ptr8728790224)
	system.RegisterType("kego.io/system", "string", ptr8728792016)
	system.RegisterType("kego.io/system", "@string", ptr8728794256)
	system.RegisterType("kego.io/system", "@int", ptr8728788992)
	system.RegisterType("kego.io/system", "@array", ptr8728789440)
	system.RegisterType("kego.io/system", "@reference", ptr8728791344)
	system.RegisterType("kego.io/system", "int", ptr8728790672)
	system.RegisterType("kego.io/system", "imports", ptr8728790336)
	system.RegisterType("kego.io/system", "@map", ptr8728790896)
	system.RegisterType("kego.io/system", "base", ptr8728789664)
}
