package types

import (
	"kego.io/system"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for aliases", Id: system.Reference{Package: "kego.io/system", Name: "@aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr4 := &system.Rule_rule{Base: ptr3, RuleBase: (*system.RuleBase)(nil)}
	ptr5 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr6 := &system.RuleBase{Optional: true}
	ptr7 := &system.Int_rule{Base: ptr5, RuleBase: ptr6, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr8 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr9 := &system.RuleBase{Optional: true}
	ptr10 := &system.Int_rule{Base: ptr8, RuleBase: ptr9, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr11 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr12 := &system.RuleBase{Optional: true}
	ptr13 := &system.JsonBool_rule{Base: ptr11, RuleBase: ptr12}
	ptr14 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr4, "maxItems": ptr7, "minItems": ptr10, "uniqueItems": ptr13}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr15 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr16 := &system.Type{Base: ptr15, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr17 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr18 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr19 := &system.RuleBase{Optional: true}
	ptr20 := &system.Bool_rule{Base: ptr18, RuleBase: ptr19, Default: system.Bool{}}
	ptr21 := &system.Type{Base: ptr17, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr20}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Type{Base: ptr22, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr26 := &system.RuleBase{Optional: true}
	ptr27 := &system.Int_rule{Base: ptr25, RuleBase: ptr26, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr28 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr29 := &system.RuleBase{Optional: true}
	ptr30 := &system.Int_rule{Base: ptr28, RuleBase: ptr29, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr31 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr32 := &system.RuleBase{Optional: true}
	ptr33 := &system.Int_rule{Base: ptr31, RuleBase: ptr32, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr34 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr35 := &system.RuleBase{Optional: true}
	ptr36 := &system.Int_rule{Base: ptr34, RuleBase: ptr35, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr37 := &system.Type{Base: ptr24, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr27, "maximum": ptr30, "minimum": ptr33, "multipleOf": ptr36}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr38 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr39 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr40 := &system.Rule_rule{Base: ptr39, RuleBase: (*system.RuleBase)(nil)}
	ptr41 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr42 := &system.RuleBase{Optional: true}
	ptr43 := &system.Int_rule{Base: ptr41, RuleBase: ptr42, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr44 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr45 := &system.RuleBase{Optional: true}
	ptr46 := &system.Int_rule{Base: ptr44, RuleBase: ptr45, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr47 := &system.Type{Base: ptr38, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr40, "maxItems": ptr43, "minItems": ptr46}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr48 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr49 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr50 := &system.RuleBase{Optional: true}
	ptr51 := &system.Number_rule{Base: ptr49, RuleBase: ptr50, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr52 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr53 := &system.RuleBase{Optional: true}
	ptr54 := &system.JsonBool_rule{Base: ptr52, RuleBase: ptr53}
	ptr55 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr56 := &system.RuleBase{Optional: true}
	ptr57 := &system.JsonBool_rule{Base: ptr55, RuleBase: ptr56}
	ptr58 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr59 := &system.RuleBase{Optional: true}
	ptr60 := &system.Number_rule{Base: ptr58, RuleBase: ptr59, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr61 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr62 := &system.RuleBase{Optional: true}
	ptr63 := &system.Number_rule{Base: ptr61, RuleBase: ptr62, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr64 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr65 := &system.RuleBase{Optional: true}
	ptr66 := &system.Number_rule{Base: ptr64, RuleBase: ptr65, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr67 := &system.Type{Base: ptr48, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr51, "exclusiveMaximum": ptr54, "exclusiveMinimum": ptr57, "maximum": ptr60, "minimum": ptr63, "multipleOf": ptr66}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr68 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr69 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr70 := &system.RuleBase{Optional: true}
	ptr71 := &system.Reference_rule{Base: ptr69, RuleBase: ptr70, Default: system.Reference{}}
	ptr72 := &system.Type{Base: ptr68, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr71}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr73 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr74 := &system.Type{Base: ptr73, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Type{Base: ptr75, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr77 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr78 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr79 := &system.RuleBase{Optional: true}
	ptr80 := &system.String_rule{Base: ptr78, RuleBase: ptr79, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr81 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr82 := &system.RuleBase{Optional: true}
	ptr83 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr84 := &system.JsonString_rule{Base: ptr83, RuleBase: (*system.RuleBase)(nil)}
	ptr85 := &system.Array_rule{Base: ptr81, RuleBase: ptr82, Items: ptr84, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr86 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr87 := &system.RuleBase{Optional: true}
	ptr88 := &system.String_rule{Base: ptr86, RuleBase: ptr87, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr89 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr90 := &system.RuleBase{Optional: true}
	ptr91 := &system.String_rule{Base: ptr89, RuleBase: ptr90, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr92 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr93 := &system.RuleBase{Optional: true}
	ptr94 := &system.Int_rule{Base: ptr92, RuleBase: ptr93, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr95 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr96 := &system.RuleBase{Optional: true}
	ptr97 := &system.Int_rule{Base: ptr95, RuleBase: ptr96, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr98 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr99 := &system.RuleBase{Optional: true}
	ptr100 := &system.String_rule{Base: ptr98, RuleBase: ptr99, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr101 := &system.Type{Base: ptr77, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr80, "enum": ptr85, "equal": ptr88, "format": ptr91, "maxLength": ptr94, "minLength": ptr97, "pattern": ptr100}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Type{Base: ptr102, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr104 := &system.Base{Description: "Lists aliases used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr105 := &system.Base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr106 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr107 := &system.JsonString_rule{Base: ptr106, RuleBase: (*system.RuleBase)(nil)}
	ptr108 := &system.Map_rule{Base: ptr105, RuleBase: (*system.RuleBase)(nil), Items: ptr107, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr109 := &system.Type{Base: ptr104, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"aliases": ptr108}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr110 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr111 := &system.Type{Base: ptr110, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr14}
	ptr112 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr113 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr114 := &system.RuleBase{Optional: true}
	ptr115 := &system.JsonString_rule{Base: ptr113, RuleBase: ptr114}
	ptr116 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr117 := &system.RuleBase{Optional: true}
	ptr118 := &system.Reference_rule{Base: ptr116, RuleBase: ptr117, Default: system.Reference{}}
	ptr119 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr120 := &system.RuleBase{Optional: true}
	ptr121 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr122 := &system.Rule_rule{Base: ptr121, RuleBase: (*system.RuleBase)(nil)}
	ptr123 := &system.Array_rule{Base: ptr119, RuleBase: ptr120, Items: ptr122, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr124 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr125 := &system.Reference_rule{Base: ptr124, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr126 := &system.Type{Base: ptr112, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"description": ptr115, "id": ptr118, "rules": ptr123, "type": ptr125}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr127 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr128 := &system.Type{Base: ptr127, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr21}
	ptr129 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr130 := &system.Base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr131 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr132 := &system.JsonString_rule{Base: ptr131, RuleBase: (*system.RuleBase)(nil)}
	ptr133 := &system.Map_rule{Base: ptr130, RuleBase: (*system.RuleBase)(nil), Items: ptr132, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr134 := &system.Type{Base: ptr129, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr133}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr135 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr136 := &system.Type{Base: ptr135, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr37}
	ptr137 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr138 := &system.Type{Base: ptr137, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr47}
	ptr139 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr140 := &system.Type{Base: ptr139, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr67}
	ptr141 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr142 := &system.Type{Base: ptr141, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr72}
	ptr143 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr144 := &system.Type{Base: ptr143, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr145 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr146 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr147 := &system.RuleBase{Optional: true}
	ptr148 := &system.JsonBool_rule{Base: ptr146, RuleBase: ptr147}
	ptr149 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr150 := &system.RuleBase{Optional: true}
	ptr151 := &system.JsonString_rule{Base: ptr149, RuleBase: ptr150}
	ptr152 := &system.Type{Base: ptr145, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr148, "selector": ptr151}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr153 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr154 := &system.Type{Base: ptr153, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr101}
	ptr155 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr156 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr157 := &system.RuleBase{Optional: true}
	ptr158 := &system.JsonBool_rule{Base: ptr156, RuleBase: ptr157}
	ptr159 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr160 := &system.RuleBase{Optional: true}
	ptr161 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr162 := &system.Reference_rule{Base: ptr161, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr163 := &system.Array_rule{Base: ptr159, RuleBase: ptr160, Items: ptr162, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr164 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr165 := &system.RuleBase{Optional: true}
	ptr166 := &system.JsonBool_rule{Base: ptr164, RuleBase: ptr165}
	ptr167 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr168 := &system.RuleBase{Optional: true}
	ptr169 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr170 := &system.Rule_rule{Base: ptr169, RuleBase: (*system.RuleBase)(nil)}
	ptr171 := &system.Map_rule{Base: ptr167, RuleBase: ptr168, Items: ptr170, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr172 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr173 := &system.RuleBase{Optional: true}
	ptr174 := &system.JsonBool_rule{Base: ptr172, RuleBase: ptr173}
	ptr175 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr176 := &system.RuleBase{Optional: true}
	ptr177 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr178 := &system.Reference_rule{Base: ptr177, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr179 := &system.Array_rule{Base: ptr175, RuleBase: ptr176, Items: ptr178, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr180 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr181 := &system.RuleBase{Optional: true}
	ptr182 := &system.String_rule{Base: ptr180, RuleBase: ptr181, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr183 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr184 := &system.RuleBase{Optional: true}
	ptr185 := &system.Type_rule{Base: ptr183, RuleBase: ptr184}
	ptr186 := &system.Type{Base: ptr155, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"basic": ptr158, "embed": ptr163, "exclude": ptr166, "fields": ptr171, "interface": ptr174, "is": ptr179, "native": ptr182, "rule": ptr185}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("kego.io/system", "@aliases", ptr1, 0xa201259ad19e56d5)
	system.RegisterType("kego.io/system", "@array", ptr14, 0x9396ab3adc9b26c2)
	system.RegisterType("kego.io/system", "@base", ptr16, 0x55165bf6126129f)
	system.RegisterType("kego.io/system", "@bool", ptr21, 0x18e72064215d8f71)
	system.RegisterType("kego.io/system", "@imports", ptr23, 0x9bb5dc657899f549)
	system.RegisterType("kego.io/system", "@int", ptr37, 0xe7b98d2f66bb5b1)
	system.RegisterType("kego.io/system", "@map", ptr47, 0x68d11e9fab3ed4bb)
	system.RegisterType("kego.io/system", "@number", ptr67, 0x8c2b0d23320b8aeb)
	system.RegisterType("kego.io/system", "@reference", ptr72, 0x2a5ddf0f63b58876)
	system.RegisterType("kego.io/system", "@rule", ptr74, 0x50ad825a1446eccd)
	system.RegisterType("kego.io/system", "@ruleBase", ptr76, 0xab213a8c8218179a)
	system.RegisterType("kego.io/system", "@string", ptr101, 0xae4613a6af2bb7d6)
	system.RegisterType("kego.io/system", "@type", ptr103, 0xc1bba2c7da9518b7)
	system.RegisterType("kego.io/system", "aliases", ptr109, 0xa201259ad19e56d5)
	system.RegisterType("kego.io/system", "array", ptr111, 0x9396ab3adc9b26c2)
	system.RegisterType("kego.io/system", "base", ptr126, 0x55165bf6126129f)
	system.RegisterType("kego.io/system", "bool", ptr128, 0x18e72064215d8f71)
	system.RegisterType("kego.io/system", "imports", ptr134, 0x9bb5dc657899f549)
	system.RegisterType("kego.io/system", "int", ptr136, 0xe7b98d2f66bb5b1)
	system.RegisterType("kego.io/system", "map", ptr138, 0x68d11e9fab3ed4bb)
	system.RegisterType("kego.io/system", "number", ptr140, 0x8c2b0d23320b8aeb)
	system.RegisterType("kego.io/system", "reference", ptr142, 0x2a5ddf0f63b58876)
	system.RegisterType("kego.io/system", "rule", ptr144, 0x50ad825a1446eccd)
	system.RegisterType("kego.io/system", "ruleBase", ptr152, 0xab213a8c8218179a)
	system.RegisterType("kego.io/system", "string", ptr154, 0xae4613a6af2bb7d6)
	system.RegisterType("kego.io/system", "type", ptr186, 0xc1bba2c7da9518b7)
}
