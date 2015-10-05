package types

import (
	"kego.io/system"
)

func init() {
	ptr0 := &system.Object_base{Description: "This is the base type for the object interface. All ke objects have this type embedded.", Id: system.Reference{Package: "kego.io/system", Name: "$object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Object_base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr2 := &system.Rule_base{Optional: true}
	ptr3 := &system.JsonString_rule{Object_base: ptr1, Rule_base: ptr2}
	ptr4 := &system.Object_base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr5 := &system.Rule_base{Optional: true}
	ptr6 := &system.Reference_rule{Object_base: ptr4, Rule_base: ptr5, Default: system.Reference{}}
	ptr7 := &system.Object_base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr8 := &system.Rule_base{Optional: true}
	ptr9 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr10 := &system.Rule_base{}
	ptr11 := &system.Rule_rule{Object_base: ptr9, Rule_base: ptr10}
	ptr12 := &system.Array_rule{Object_base: ptr7, Rule_base: ptr8, Items: ptr11, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr13 := &system.Object_base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr14 := &system.Rule_base{}
	ptr15 := &system.Reference_rule{Object_base: ptr13, Rule_base: ptr14, Default: system.Reference{}}
	ptr16 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"description": ptr3, "id": ptr6, "rules": ptr12, "type": ptr15}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr17 := &system.Object_base{Description: "All rules will have this embedded in them.", Id: system.Reference{Package: "kego.io/system", Name: "$rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr18 := &system.Object_base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr19 := &system.Rule_base{Optional: true}
	ptr20 := &system.JsonBool_rule{Object_base: ptr18, Rule_base: ptr19}
	ptr21 := &system.Object_base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr22 := &system.Rule_base{Optional: true}
	ptr23 := &system.JsonString_rule{Object_base: ptr21, Rule_base: ptr22}
	ptr24 := &system.Type{Object_base: ptr17, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr20, "selector": ptr23}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr25 := &system.Object_base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr26 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr27 := &system.Rule_base{}
	ptr28 := &system.Rule_rule{Object_base: ptr26, Rule_base: ptr27}
	ptr29 := &system.Object_base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr30 := &system.Rule_base{Optional: true}
	ptr31 := &system.Int_rule{Object_base: ptr29, Rule_base: ptr30, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr32 := &system.Object_base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr33 := &system.Rule_base{Optional: true}
	ptr34 := &system.Int_rule{Object_base: ptr32, Rule_base: ptr33, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr35 := &system.Object_base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr36 := &system.Rule_base{Optional: true}
	ptr37 := &system.JsonBool_rule{Object_base: ptr35, Rule_base: ptr36}
	ptr38 := &system.Type{Object_base: ptr25, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr28, "maxItems": ptr31, "minItems": ptr34, "uniqueItems": ptr37}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr39 := &system.Object_base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr40 := &system.Object_base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr41 := &system.Rule_base{Optional: true}
	ptr42 := &system.Bool_rule{Object_base: ptr40, Rule_base: ptr41, Default: system.Bool{}}
	ptr43 := &system.Type{Object_base: ptr39, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr42}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr44 := &system.Object_base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr45 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr46 := &system.Rule_base{Optional: true}
	ptr47 := &system.Int_rule{Object_base: ptr45, Rule_base: ptr46, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr48 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr49 := &system.Rule_base{Optional: true}
	ptr50 := &system.Int_rule{Object_base: ptr48, Rule_base: ptr49, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr51 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr52 := &system.Rule_base{Optional: true}
	ptr53 := &system.Int_rule{Object_base: ptr51, Rule_base: ptr52, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr54 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr55 := &system.Rule_base{Optional: true}
	ptr56 := &system.Int_rule{Object_base: ptr54, Rule_base: ptr55, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr57 := &system.Type{Object_base: ptr44, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr47, "maximum": ptr50, "minimum": ptr53, "multipleOf": ptr56}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr58 := &system.Object_base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr59 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr60 := &system.Rule_base{}
	ptr61 := &system.Rule_rule{Object_base: ptr59, Rule_base: ptr60}
	ptr62 := &system.Object_base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr63 := &system.Rule_base{Optional: true}
	ptr64 := &system.Int_rule{Object_base: ptr62, Rule_base: ptr63, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr65 := &system.Object_base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr66 := &system.Rule_base{Optional: true}
	ptr67 := &system.Int_rule{Object_base: ptr65, Rule_base: ptr66, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr68 := &system.Type{Object_base: ptr58, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr61, "maxItems": ptr64, "minItems": ptr67}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr69 := &system.Object_base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr70 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr71 := &system.Rule_base{Optional: true}
	ptr72 := &system.Number_rule{Object_base: ptr70, Rule_base: ptr71, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr73 := &system.Object_base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr74 := &system.Rule_base{Optional: true}
	ptr75 := &system.JsonBool_rule{Object_base: ptr73, Rule_base: ptr74}
	ptr76 := &system.Object_base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr77 := &system.Rule_base{Optional: true}
	ptr78 := &system.JsonBool_rule{Object_base: ptr76, Rule_base: ptr77}
	ptr79 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr80 := &system.Rule_base{Optional: true}
	ptr81 := &system.Number_rule{Object_base: ptr79, Rule_base: ptr80, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr82 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr83 := &system.Rule_base{Optional: true}
	ptr84 := &system.Number_rule{Object_base: ptr82, Rule_base: ptr83, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr85 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr86 := &system.Rule_base{Optional: true}
	ptr87 := &system.Number_rule{Object_base: ptr85, Rule_base: ptr86, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr88 := &system.Type{Object_base: ptr69, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr72, "exclusiveMaximum": ptr75, "exclusiveMinimum": ptr78, "maximum": ptr81, "minimum": ptr84, "multipleOf": ptr87}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr89 := &system.Object_base{Description: "Automatically created basic rule for object", Id: system.Reference{Package: "kego.io/system", Name: "@object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr90 := &system.Type{Object_base: ptr89, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr91 := &system.Object_base{Description: "Automatically created basic rule for package", Id: system.Reference{Package: "kego.io/system", Name: "@package", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr92 := &system.Type{Object_base: ptr91, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr93 := &system.Object_base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr94 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr95 := &system.Rule_base{Optional: true}
	ptr96 := &system.Reference_rule{Object_base: ptr94, Rule_base: ptr95, Default: system.Reference{}}
	ptr97 := &system.Type{Object_base: ptr93, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr96}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr98 := &system.Object_base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr99 := &system.Type{Object_base: ptr98, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr100 := &system.Object_base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr101 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr102 := &system.Rule_base{Optional: true}
	ptr103 := &system.String_rule{Object_base: ptr101, Rule_base: ptr102, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr104 := &system.Object_base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr105 := &system.Rule_base{Optional: true}
	ptr106 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr107 := &system.Rule_base{}
	ptr108 := &system.JsonString_rule{Object_base: ptr106, Rule_base: ptr107}
	ptr109 := &system.Array_rule{Object_base: ptr104, Rule_base: ptr105, Items: ptr108, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr110 := &system.Object_base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr111 := &system.Rule_base{Optional: true}
	ptr112 := &system.String_rule{Object_base: ptr110, Rule_base: ptr111, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr113 := &system.Object_base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr114 := &system.Rule_base{Optional: true}
	ptr115 := &system.String_rule{Object_base: ptr113, Rule_base: ptr114, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr116 := &system.Object_base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr117 := &system.Rule_base{Optional: true}
	ptr118 := &system.Int_rule{Object_base: ptr116, Rule_base: ptr117, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr119 := &system.Object_base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr120 := &system.Rule_base{Optional: true}
	ptr121 := &system.Int_rule{Object_base: ptr119, Rule_base: ptr120, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr122 := &system.Object_base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr123 := &system.Rule_base{Optional: true}
	ptr124 := &system.String_rule{Object_base: ptr122, Rule_base: ptr123, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr125 := &system.Type{Object_base: ptr100, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr103, "enum": ptr109, "equal": ptr112, "format": ptr115, "maxLength": ptr118, "minLength": ptr121, "pattern": ptr124}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr126 := &system.Object_base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr127 := &system.Type{Object_base: ptr126, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr128 := &system.Object_base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr129 := &system.Type{Object_base: ptr128, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr38}
	ptr130 := &system.Object_base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr131 := &system.Type{Object_base: ptr130, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr43}
	ptr132 := &system.Object_base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr133 := &system.Type{Object_base: ptr132, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr57}
	ptr134 := &system.Object_base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr135 := &system.Type{Object_base: ptr134, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr68}
	ptr136 := &system.Object_base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr137 := &system.Type{Object_base: ptr136, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr88}
	ptr138 := &system.Object_base{Description: "This is the interface that all ke types should implement", Id: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr139 := &system.Type{Object_base: ptr138, Base: ptr16, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr140 := &system.Object_base{Description: "Package info - forms the root node of the package", Id: system.Reference{Package: "kego.io/system", Name: "package", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr141 := &system.Object_base{Description: "The global map is populated with all the global objects in the package", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr142 := &system.Rule_base{Optional: true}
	ptr143 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@object", Exists: true}}
	ptr144 := &system.Rule_base{}
	ptr145 := &system.Object_rule{Object_base: ptr143, Rule_base: ptr144}
	ptr146 := &system.Map_rule{Object_base: ptr141, Rule_base: ptr142, Items: ptr145, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr147 := &system.Object_base{Description: "Map of import aliases used in this package: key = package path, value = alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr148 := &system.Rule_base{Optional: true}
	ptr149 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr150 := &system.Rule_base{}
	ptr151 := &system.JsonString_rule{Object_base: ptr149, Rule_base: ptr150}
	ptr152 := &system.Map_rule{Object_base: ptr147, Rule_base: ptr148, Items: ptr151, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr153 := &system.Type{Object_base: ptr140, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"global": ptr146, "import": ptr152}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr154 := &system.Object_base{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr155 := &system.Type{Object_base: ptr154, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr97}
	ptr156 := &system.Object_base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr157 := &system.Type{Object_base: ptr156, Base: ptr24, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr158 := &system.Object_base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr159 := &system.Type{Object_base: ptr158, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr125}
	ptr160 := &system.Object_base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr161 := &system.Object_base{Description: "Optionally this is a type which is embedded in each object that implements this interface.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr162 := &system.Rule_base{Optional: true}
	ptr163 := &system.Type_rule{Object_base: ptr161, Rule_base: ptr162}
	ptr164 := &system.Object_base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr165 := &system.Rule_base{Optional: true}
	ptr166 := &system.JsonBool_rule{Object_base: ptr164, Rule_base: ptr165}
	ptr167 := &system.Object_base{Description: "Types which this should embed - system:object is always added unless basic = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr168 := &system.Rule_base{Optional: true}
	ptr169 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr170 := &system.Rule_base{}
	ptr171 := &system.Reference_rule{Object_base: ptr169, Rule_base: ptr170, Default: system.Reference{}}
	ptr172 := &system.Array_rule{Object_base: ptr167, Rule_base: ptr168, Items: ptr171, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr173 := &system.Object_base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr174 := &system.Rule_base{Optional: true}
	ptr175 := &system.JsonBool_rule{Object_base: ptr173, Rule_base: ptr174}
	ptr176 := &system.Object_base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr177 := &system.Rule_base{Optional: true}
	ptr178 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr179 := &system.Rule_base{}
	ptr180 := &system.Rule_rule{Object_base: ptr178, Rule_base: ptr179}
	ptr181 := &system.Map_rule{Object_base: ptr176, Rule_base: ptr177, Items: ptr180, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr182 := &system.Object_base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr183 := &system.Rule_base{Optional: true}
	ptr184 := &system.JsonBool_rule{Object_base: ptr182, Rule_base: ptr183}
	ptr185 := &system.Object_base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr186 := &system.Rule_base{Optional: true}
	ptr187 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr188 := &system.Rule_base{}
	ptr189 := &system.Reference_rule{Object_base: ptr187, Rule_base: ptr188, Default: system.Reference{}}
	ptr190 := &system.Array_rule{Object_base: ptr185, Rule_base: ptr186, Items: ptr189, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr191 := &system.Object_base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr192 := &system.Rule_base{Optional: true}
	ptr193 := &system.String_rule{Object_base: ptr191, Rule_base: ptr192, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr194 := &system.Object_base{Description: "Type that defines restriction rules for this type.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr195 := &system.Rule_base{Optional: true}
	ptr196 := &system.Type_rule{Object_base: ptr194, Rule_base: ptr195}
	ptr197 := &system.Type{Object_base: ptr160, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"base": ptr163, "basic": ptr166, "embed": ptr172, "exclude": ptr175, "fields": ptr181, "interface": ptr184, "is": ptr190, "native": ptr193, "rule": ptr196}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/system", "$object", ptr16, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "$rule", ptr24, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "@array", ptr38, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "@bool", ptr43, 0x849d95e096ea903a)
	system.Register("kego.io/system", "@int", ptr57, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "@map", ptr68, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "@number", ptr88, 0x14f986941edf71e9)
	system.Register("kego.io/system", "@object", ptr90, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "@package", ptr92, 0x45625b24c593f855)
	system.Register("kego.io/system", "@reference", ptr97, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "@rule", ptr99, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "@string", ptr125, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "@type", ptr127, 0xe6e176a06eb36092)
	system.Register("kego.io/system", "array", ptr129, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "bool", ptr131, 0x849d95e096ea903a)
	system.Register("kego.io/system", "int", ptr133, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "map", ptr135, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "number", ptr137, 0x14f986941edf71e9)
	system.Register("kego.io/system", "object", ptr139, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "package", ptr153, 0x45625b24c593f855)
	system.Register("kego.io/system", "reference", ptr155, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "rule", ptr157, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "string", ptr159, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "type", ptr197, 0xe6e176a06eb36092)
}
