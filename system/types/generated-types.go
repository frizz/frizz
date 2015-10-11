package types

import "kego.io/system"

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
	ptr18 := &system.Object_base{Description: "Special field that should be excluded when marshaling - e.g. package.global", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr19 := &system.Rule_base{Optional: true}
	ptr20 := &system.JsonBool_rule{Object_base: ptr18, Rule_base: ptr19}
	ptr21 := &system.Object_base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr22 := &system.Rule_base{Optional: true}
	ptr23 := &system.JsonBool_rule{Object_base: ptr21, Rule_base: ptr22}
	ptr24 := &system.Object_base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr25 := &system.Rule_base{Optional: true}
	ptr26 := &system.JsonString_rule{Object_base: ptr24, Rule_base: ptr25}
	ptr27 := &system.Type{Object_base: ptr17, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"exclude": ptr20, "optional": ptr23, "selector": ptr26}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr28 := &system.Object_base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr29 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr30 := &system.Rule_base{}
	ptr31 := &system.Rule_rule{Object_base: ptr29, Rule_base: ptr30}
	ptr32 := &system.Object_base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr33 := &system.Rule_base{Optional: true}
	ptr34 := &system.Int_rule{Object_base: ptr32, Rule_base: ptr33, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr35 := &system.Object_base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr36 := &system.Rule_base{Optional: true}
	ptr37 := &system.Int_rule{Object_base: ptr35, Rule_base: ptr36, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr38 := &system.Object_base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr39 := &system.Rule_base{Optional: true}
	ptr40 := &system.JsonBool_rule{Object_base: ptr38, Rule_base: ptr39}
	ptr41 := &system.Type{Object_base: ptr28, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr31, "maxItems": ptr34, "minItems": ptr37, "uniqueItems": ptr40}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr42 := &system.Object_base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr43 := &system.Object_base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr44 := &system.Rule_base{Optional: true}
	ptr45 := &system.Bool_rule{Object_base: ptr43, Rule_base: ptr44, Default: system.Bool{}}
	ptr46 := &system.Type{Object_base: ptr42, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr45}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr47 := &system.Object_base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr48 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr49 := &system.Rule_base{Optional: true}
	ptr50 := &system.Int_rule{Object_base: ptr48, Rule_base: ptr49, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr51 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr52 := &system.Rule_base{Optional: true}
	ptr53 := &system.Int_rule{Object_base: ptr51, Rule_base: ptr52, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr54 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr55 := &system.Rule_base{Optional: true}
	ptr56 := &system.Int_rule{Object_base: ptr54, Rule_base: ptr55, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr57 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr58 := &system.Rule_base{Optional: true}
	ptr59 := &system.Int_rule{Object_base: ptr57, Rule_base: ptr58, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr60 := &system.Type{Object_base: ptr47, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr50, "maximum": ptr53, "minimum": ptr56, "multipleOf": ptr59}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr61 := &system.Object_base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr62 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr63 := &system.Rule_base{}
	ptr64 := &system.Rule_rule{Object_base: ptr62, Rule_base: ptr63}
	ptr65 := &system.Object_base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr66 := &system.Rule_base{Optional: true}
	ptr67 := &system.Int_rule{Object_base: ptr65, Rule_base: ptr66, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr68 := &system.Object_base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr69 := &system.Rule_base{Optional: true}
	ptr70 := &system.Int_rule{Object_base: ptr68, Rule_base: ptr69, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr71 := &system.Type{Object_base: ptr61, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr64, "maxItems": ptr67, "minItems": ptr70}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr72 := &system.Object_base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr73 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr74 := &system.Rule_base{Optional: true}
	ptr75 := &system.Number_rule{Object_base: ptr73, Rule_base: ptr74, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr76 := &system.Object_base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr77 := &system.Rule_base{Optional: true}
	ptr78 := &system.JsonBool_rule{Object_base: ptr76, Rule_base: ptr77}
	ptr79 := &system.Object_base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr80 := &system.Rule_base{Optional: true}
	ptr81 := &system.JsonBool_rule{Object_base: ptr79, Rule_base: ptr80}
	ptr82 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr83 := &system.Rule_base{Optional: true}
	ptr84 := &system.Number_rule{Object_base: ptr82, Rule_base: ptr83, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr85 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr86 := &system.Rule_base{Optional: true}
	ptr87 := &system.Number_rule{Object_base: ptr85, Rule_base: ptr86, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr88 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr89 := &system.Rule_base{Optional: true}
	ptr90 := &system.Number_rule{Object_base: ptr88, Rule_base: ptr89, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr91 := &system.Type{Object_base: ptr72, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr75, "exclusiveMaximum": ptr78, "exclusiveMinimum": ptr81, "maximum": ptr84, "minimum": ptr87, "multipleOf": ptr90}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr92 := &system.Object_base{Description: "Automatically created basic rule for object", Id: system.Reference{Package: "kego.io/system", Name: "@object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr93 := &system.Type{Object_base: ptr92, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr94 := &system.Object_base{Description: "Automatically created basic rule for package", Id: system.Reference{Package: "kego.io/system", Name: "@package", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr95 := &system.Type{Object_base: ptr94, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr96 := &system.Object_base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr97 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr98 := &system.Rule_base{Optional: true}
	ptr99 := &system.Reference_rule{Object_base: ptr97, Rule_base: ptr98, Default: system.Reference{}}
	ptr100 := &system.Type{Object_base: ptr96, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr99}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr101 := &system.Object_base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr102 := &system.Type{Object_base: ptr101, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr103 := &system.Object_base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr104 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr105 := &system.Rule_base{Optional: true}
	ptr106 := &system.String_rule{Object_base: ptr104, Rule_base: ptr105, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr107 := &system.Object_base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr108 := &system.Rule_base{Optional: true}
	ptr109 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr110 := &system.Rule_base{}
	ptr111 := &system.JsonString_rule{Object_base: ptr109, Rule_base: ptr110}
	ptr112 := &system.Array_rule{Object_base: ptr107, Rule_base: ptr108, Items: ptr111, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr113 := &system.Object_base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr114 := &system.Rule_base{Optional: true}
	ptr115 := &system.String_rule{Object_base: ptr113, Rule_base: ptr114, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr116 := &system.Object_base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr117 := &system.Rule_base{Optional: true}
	ptr118 := &system.String_rule{Object_base: ptr116, Rule_base: ptr117, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr119 := &system.Object_base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr120 := &system.Rule_base{Optional: true}
	ptr121 := &system.Int_rule{Object_base: ptr119, Rule_base: ptr120, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr122 := &system.Object_base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr123 := &system.Rule_base{Optional: true}
	ptr124 := &system.Int_rule{Object_base: ptr122, Rule_base: ptr123, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr125 := &system.Object_base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr126 := &system.Rule_base{Optional: true}
	ptr127 := &system.String_rule{Object_base: ptr125, Rule_base: ptr126, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr128 := &system.Type{Object_base: ptr103, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr106, "enum": ptr112, "equal": ptr115, "format": ptr118, "maxLength": ptr121, "minLength": ptr124, "pattern": ptr127}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr129 := &system.Object_base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr130 := &system.Type{Object_base: ptr129, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr131 := &system.Object_base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr132 := &system.Type{Object_base: ptr131, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr41}
	ptr133 := &system.Object_base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr134 := &system.Type{Object_base: ptr133, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr46}
	ptr135 := &system.Object_base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr136 := &system.Type{Object_base: ptr135, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr60}
	ptr137 := &system.Object_base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr138 := &system.Type{Object_base: ptr137, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr71}
	ptr139 := &system.Object_base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr140 := &system.Type{Object_base: ptr139, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr91}
	ptr141 := &system.Object_base{Description: "This is the interface that all ke types should implement", Id: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr142 := &system.Type{Object_base: ptr141, Base: ptr16, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr143 := &system.Object_base{Description: "Package info - forms the root node of the package", Id: system.Reference{Package: "kego.io/system", Name: "package", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr144 := &system.Object_base{Description: "Map of import aliases used in this package: key = package path, value = alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr145 := &system.Rule_base{Optional: true}
	ptr146 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr147 := &system.Rule_base{}
	ptr148 := &system.JsonString_rule{Object_base: ptr146, Rule_base: ptr147}
	ptr149 := &system.Map_rule{Object_base: ptr144, Rule_base: ptr145, Items: ptr148, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr150 := &system.Type{Object_base: ptr143, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"aliases": ptr149}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr151 := &system.Object_base{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr152 := &system.Type{Object_base: ptr151, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr100}
	ptr153 := &system.Object_base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr154 := &system.Type{Object_base: ptr153, Base: ptr27, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr155 := &system.Object_base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr156 := &system.Type{Object_base: ptr155, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr128}
	ptr157 := &system.Object_base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr158 := &system.Object_base{Description: "Optionally this is a type which is embedded in each object that implements this interface.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr159 := &system.Rule_base{Optional: true}
	ptr160 := &system.Type_rule{Object_base: ptr158, Rule_base: ptr159}
	ptr161 := &system.Object_base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr162 := &system.Rule_base{Optional: true}
	ptr163 := &system.JsonBool_rule{Object_base: ptr161, Rule_base: ptr162}
	ptr164 := &system.Object_base{Description: "Types which this should embed - system:object is always added unless basic = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr165 := &system.Rule_base{Optional: true}
	ptr166 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr167 := &system.Rule_base{}
	ptr168 := &system.Reference_rule{Object_base: ptr166, Rule_base: ptr167, Default: system.Reference{}}
	ptr169 := &system.Array_rule{Object_base: ptr164, Rule_base: ptr165, Items: ptr168, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr170 := &system.Object_base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr171 := &system.Rule_base{Optional: true}
	ptr172 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr173 := &system.Rule_base{}
	ptr174 := &system.Rule_rule{Object_base: ptr172, Rule_base: ptr173}
	ptr175 := &system.Map_rule{Object_base: ptr170, Rule_base: ptr171, Items: ptr174, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr176 := &system.Object_base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr177 := &system.Rule_base{Optional: true}
	ptr178 := &system.JsonBool_rule{Object_base: ptr176, Rule_base: ptr177}
	ptr179 := &system.Object_base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr180 := &system.Rule_base{Optional: true}
	ptr181 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr182 := &system.Rule_base{}
	ptr183 := &system.Reference_rule{Object_base: ptr181, Rule_base: ptr182, Default: system.Reference{}}
	ptr184 := &system.Array_rule{Object_base: ptr179, Rule_base: ptr180, Items: ptr183, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr185 := &system.Object_base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr186 := &system.Rule_base{Optional: true}
	ptr187 := &system.String_rule{Object_base: ptr185, Rule_base: ptr186, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr188 := &system.Object_base{Description: "Type that defines restriction rules for this type.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr189 := &system.Rule_base{Optional: true}
	ptr190 := &system.Type_rule{Object_base: ptr188, Rule_base: ptr189}
	ptr191 := &system.Type{Object_base: ptr157, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"base": ptr160, "basic": ptr163, "embed": ptr169, "fields": ptr175, "interface": ptr178, "is": ptr184, "native": ptr187, "rule": ptr190}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/system", "$object", ptr16, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "$rule", ptr27, 0x8365876b3555b5d)
	system.Register("kego.io/system", "@array", ptr41, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "@bool", ptr46, 0x849d95e096ea903a)
	system.Register("kego.io/system", "@int", ptr60, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "@map", ptr71, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "@number", ptr91, 0x14f986941edf71e9)
	system.Register("kego.io/system", "@object", ptr93, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "@package", ptr95, 0x5168e36553335f0b)
	system.Register("kego.io/system", "@reference", ptr100, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "@rule", ptr102, 0x8365876b3555b5d)
	system.Register("kego.io/system", "@string", ptr128, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "@type", ptr130, 0x9a6d7f2bf8effa6f)
	system.Register("kego.io/system", "array", ptr132, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "bool", ptr134, 0x849d95e096ea903a)
	system.Register("kego.io/system", "int", ptr136, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "map", ptr138, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "number", ptr140, 0x14f986941edf71e9)
	system.Register("kego.io/system", "object", ptr142, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "package", ptr150, 0x5168e36553335f0b)
	system.Register("kego.io/system", "reference", ptr152, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "rule", ptr154, 0x8365876b3555b5d)
	system.Register("kego.io/system", "string", ptr156, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "type", ptr191, 0x9a6d7f2bf8effa6f)
}
