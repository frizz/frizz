package types

import (
	"kego.io/system"
)

func init() {

	ptr8729621632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609216 := &system.Base{Context: ptr8729621632, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729618400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609728 := &system.Base{Context: ptr8729618400, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729618336 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729084544 := &system.Reference_rule{Base: ptr8729609728, RuleBase: ptr8729618336, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729621504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610752 := &system.Base{Context: ptr8729621504, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729621472 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729255968 := &system.Type_rule{Base: ptr8729610752, RuleBase: ptr8729621472}

	ptr8729619104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609856 := &system.Base{Context: ptr8729619104, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729619072 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729619040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609984 := &system.Base{Context: ptr8729619040, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729085824 := &system.Reference_rule{Base: ptr8729609984, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729084784 := &system.Array_rule{Base: ptr8729609856, RuleBase: ptr8729619072, Items: ptr8729085824, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729619776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610112 := &system.Base{Context: ptr8729619776, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729619744 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729222016 := &system.String_rule{Base: ptr8729610112, RuleBase: ptr8729619744, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729620192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610240 := &system.Base{Context: ptr8729620192, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729620160 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728910592 := &system.JsonBool_rule{Base: ptr8729610240, RuleBase: ptr8729620160}

	ptr8729620576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610368 := &system.Base{Context: ptr8729620576, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729620544 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728910832 := &system.JsonBool_rule{Base: ptr8729610368, RuleBase: ptr8729620544}

	ptr8729621120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610496 := &system.Base{Context: ptr8729621120, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729621088 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729621056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610624 := &system.Base{Context: ptr8729621056, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728911616 := &system.Rule_rule{Base: ptr8729610624, RuleBase: (*system.RuleBase)(nil)}

	ptr8729069504 := &system.Map_rule{Base: ptr8729610496, RuleBase: ptr8729621088, Items: ptr8728911616, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729616992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609344 := &system.Base{Context: ptr8729616992, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729616960 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728905712 := &system.JsonBool_rule{Base: ptr8729609344, RuleBase: ptr8729616960}

	ptr8729617664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609472 := &system.Base{Context: ptr8729617664, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729617632 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729617600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609600 := &system.Base{Context: ptr8729617600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729084464 := &system.Reference_rule{Base: ptr8729609600, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729084224 := &system.Array_rule{Base: ptr8729609472, RuleBase: ptr8729617632, Items: ptr8729084464, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729221840 := &system.Type{Base: ptr8729609216, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"extends": ptr8729084544, "rule": ptr8729255968, "is": ptr8729084784, "native": ptr8729222016, "interface": ptr8728910592, "exclude": ptr8728910832, "fields": ptr8729069504, "basic": ptr8728905712, "embed": ptr8729084224}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728992448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234944 := &system.Base{Context: ptr8728992448, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217968 := &system.Type{Base: ptr8729234944, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729507872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237504 := &system.Base{Context: ptr8729507872, Description: "This is a property of a type", Id: "property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729507072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237632 := &system.Base{Context: ptr8729507072, Description: "This specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729507040 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729258336 := &system.JsonBool_rule{Base: ptr8729237632, RuleBase: ptr8729507040}

	ptr8729507456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237760 := &system.Base{Context: ptr8729507456, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729507424 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729258976 := &system.JsonBool_rule{Base: ptr8729237760, RuleBase: ptr8729507424}

	ptr8729507744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237888 := &system.Base{Context: ptr8729507744, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729259216 := &system.Rule_rule{Base: ptr8729237888, RuleBase: (*system.RuleBase)(nil)}

	ptr8729219200 := &system.Type{Base: ptr8729237504, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729258336, "defaulter": ptr8729258976, "item": ptr8729259216}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729505856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236608 := &system.Base{Context: ptr8729505856, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729503552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236992 := &system.Base{Context: ptr8729503552, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729503520 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729092576 := &system.Number_rule{Base: ptr8729236992, RuleBase: ptr8729503520, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729504224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237120 := &system.Base{Context: ptr8729504224, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729504192 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729504064 := &system.Bool_rule{Base: ptr8729237120, RuleBase: ptr8729504192, Default: system.Bool{Value: false, Exists: true}}

	ptr8729504736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237248 := &system.Base{Context: ptr8729504736, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729504672 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729092672 := &system.Number_rule{Base: ptr8729237248, RuleBase: ptr8729504672, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729505600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237376 := &system.Base{Context: ptr8729505600, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729505568 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729505376 := &system.Bool_rule{Base: ptr8729237376, RuleBase: ptr8729505568, Default: system.Bool{Value: false, Exists: true}}

	ptr8729502240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236736 := &system.Base{Context: ptr8729502240, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729502208 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729092384 := &system.Number_rule{Base: ptr8729236736, RuleBase: ptr8729502208, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729502944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236864 := &system.Base{Context: ptr8729502944, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729502816 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729092480 := &system.Number_rule{Base: ptr8729236864, RuleBase: ptr8729502816, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729219024 := &system.Type{Base: ptr8729236608, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"minimum": ptr8729092576, "exclusiveMinimum": ptr8729504064, "maximum": ptr8729092672, "exclusiveMaximum": ptr8729505376, "default": ptr8729092384, "multipleOf": ptr8729092480}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729622048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729610880 := &system.Base{Context: ptr8729622048, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729222192 := &system.Type{Base: ptr8729610880, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729502304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235200 := &system.Base{Context: ptr8729502304, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728993312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235328 := &system.Base{Context: ptr8728993312, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728993280 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729084864 := &system.Int_rule{Base: ptr8729235328, RuleBase: ptr8728993280, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728993632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235456 := &system.Base{Context: ptr8728993632, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728993600 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729084944 := &system.Int_rule{Base: ptr8729235456, RuleBase: ptr8728993600, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729501856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235584 := &system.Base{Context: ptr8729501856, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729501824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085024 := &system.Int_rule{Base: ptr8729235584, RuleBase: ptr8729501824, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729502176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235712 := &system.Base{Context: ptr8729502176, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729502144 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085104 := &system.Int_rule{Base: ptr8729235712, RuleBase: ptr8729502144, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729218320 := &system.Type{Base: ptr8729235200, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729084864, "multipleOf": ptr8729084944, "minimum": ptr8729085024, "maximum": ptr8729085104}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729506176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236480 := &system.Base{Context: ptr8729506176, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729218848 := &system.Type{Base: ptr8729236480, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729219024}

	ptr8728991104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239168 := &system.Base{Context: ptr8728991104, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220432 := &system.Type{Base: ptr8729239168, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728991392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234560 := &system.Base{Context: ptr8728991392, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728991264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234688 := &system.Base{Context: ptr8728991264, Description: "Default value if this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728991232 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728991136 := &system.Bool_rule{Base: ptr8729234688, RuleBase: ptr8728991232, Default: system.Bool{Value: false, Exists: false}}

	ptr8729217616 := &system.Type{Base: ptr8729234560, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728991136}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729618144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239424 := &system.Base{Context: ptr8729618144, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728992704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608192 := &system.Base{Context: ptr8728992704, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728992640 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729220960 := &system.String_rule{Base: ptr8729608192, RuleBase: ptr8728992640, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728993408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608320 := &system.Base{Context: ptr8728993408, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728993376 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728993344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608448 := &system.Base{Context: ptr8728993344, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729221136 := &system.String_rule{Base: ptr8729608448, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729085184 := &system.Array_rule{Base: ptr8729608320, RuleBase: ptr8728993376, Items: ptr8729221136, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729616480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608576 := &system.Base{Context: ptr8729616480, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729616448 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085904 := &system.Int_rule{Base: ptr8729608576, RuleBase: ptr8729616448, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729616800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608704 := &system.Base{Context: ptr8729616800, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729616768 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085984 := &system.Int_rule{Base: ptr8729608704, RuleBase: ptr8729616768, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729617120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608832 := &system.Base{Context: ptr8729617120, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729617088 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729221312 := &system.String_rule{Base: ptr8729608832, RuleBase: ptr8729617088, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729617440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729608960 := &system.Base{Context: ptr8729617440, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729617408 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729221488 := &system.String_rule{Base: ptr8729608960, RuleBase: ptr8729617408, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729618016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729609088 := &system.Base{Context: ptr8729618016, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729617984 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729221664 := &system.String_rule{Base: ptr8729609088, RuleBase: ptr8729617984, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729220784 := &system.Type{Base: ptr8729239424, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729220960, "enum": ptr8729085184, "minLength": ptr8729085904, "maxLength": ptr8729085984, "equal": ptr8729221312, "pattern": ptr8729221488, "format": ptr8729221664}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728990368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238784 := &system.Base{Context: ptr8728990368, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728989184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238912 := &system.Base{Context: ptr8728989184, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728988928 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729261344 := &system.JsonBool_rule{Base: ptr8729238912, RuleBase: ptr8728988928}

	ptr8728990144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239040 := &system.Base{Context: ptr8728990144, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728989824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729261568 := &system.JsonString_rule{Base: ptr8729239040, RuleBase: ptr8728989824}

	ptr8729220256 := &system.Type{Base: ptr8729238784, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729261344, "selector": ptr8729261568}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728989952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729231360 := &system.Base{Context: ptr8728989952, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728989632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729232640 := &system.Base{Context: ptr8728989632, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728988288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729232896 := &system.Base{Context: ptr8728988288, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728907264 := &system.Rule_rule{Base: ptr8729232896, RuleBase: (*system.RuleBase)(nil)}

	ptr8728988704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233024 := &system.Base{Context: ptr8728988704, Description: "This is the minimum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728988672 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729087984 := &system.Int_rule{Base: ptr8729233024, RuleBase: ptr8728988672, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728989056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233152 := &system.Base{Context: ptr8728989056, Description: "This is the maximum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728989024 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729088064 := &system.Int_rule{Base: ptr8729233152, RuleBase: ptr8728989024, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728989504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233280 := &system.Base{Context: ptr8728989504, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728989472 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728989280 := &system.Bool_rule{Base: ptr8729233280, RuleBase: ptr8728989472, Default: system.Bool{Value: false, Exists: true}}

	ptr8729217088 := &system.Type{Base: ptr8729232640, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728907264, "minItems": ptr8729087984, "maxItems": ptr8729088064, "uniqueItems": ptr8728989280}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729215152 := &system.Type{Base: ptr8729231360, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729217088}

	ptr8728990016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234304 := &system.Base{Context: ptr8728990016, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729215328 := &system.Type{Base: ptr8729234304, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728991712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234432 := &system.Base{Context: ptr8728991712, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217440 := &system.Type{Base: ptr8729234432, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729217616}

	ptr8728992256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234816 := &system.Base{Context: ptr8728992256, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217792 := &system.Type{Base: ptr8729234816, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729502624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235072 := &system.Base{Context: ptr8729502624, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729218144 := &system.Type{Base: ptr8729235072, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729218320}

	ptr8728986592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238528 := &system.Base{Context: ptr8728986592, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729219904 := &system.Type{Base: ptr8729238528, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729504800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235840 := &system.Base{Context: ptr8729504800, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729504480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235968 := &system.Base{Context: ptr8729504480, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729503648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236096 := &system.Base{Context: ptr8729503648, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729257312 := &system.Rule_rule{Base: ptr8729236096, RuleBase: (*system.RuleBase)(nil)}

	ptr8729504000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236224 := &system.Base{Context: ptr8729504000, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729503968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085344 := &system.Int_rule{Base: ptr8729236224, RuleBase: ptr8729503968, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729504352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236352 := &system.Base{Context: ptr8729504352, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729504320 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729085664 := &system.Int_rule{Base: ptr8729236352, RuleBase: ptr8729504320, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729218672 := &system.Type{Base: ptr8729235968, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729257312, "minItems": ptr8729085344, "maxItems": ptr8729085664}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729218496 := &system.Type{Base: ptr8729235840, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729218672}

	ptr8728988064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238656 := &system.Base{Context: ptr8728988064, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220080 := &system.Type{Base: ptr8729238656, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729509568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238144 := &system.Base{Context: ptr8729509568, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729509248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238272 := &system.Base{Context: ptr8729509248, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729509120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238400 := &system.Base{Context: ptr8729509120, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729509088 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729084384 := &system.Reference_rule{Base: ptr8729238400, RuleBase: ptr8729509088, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729219728 := &system.Type{Base: ptr8729238272, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729084384}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729219552 := &system.Type{Base: ptr8729238144, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729219728}

	ptr8729508288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238016 := &system.Base{Context: ptr8729508288, Description: "Automatically created basic rule for property", Id: "@property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729219376 := &system.Type{Base: ptr8729238016, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729618464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239296 := &system.Base{Context: ptr8729618464, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220608 := &system.Type{Base: ptr8729239296, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729220784}

	ptr8728989376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233408 := &system.Base{Context: ptr8728989376, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728987424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233792 := &system.Base{Context: ptr8728987424, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728987264 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728905104 := &system.JsonString_rule{Base: ptr8729233792, RuleBase: ptr8728987264}

	ptr8728988224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233920 := &system.Base{Context: ptr8728988224, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728988192 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728906992 := &system.Context_rule{Base: ptr8729233920, RuleBase: ptr8728988192}

	ptr8728989152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234048 := &system.Base{Context: ptr8728989152, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728989120 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728989088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234176 := &system.Base{Context: ptr8728989088, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728908960 := &system.Rule_rule{Base: ptr8729234176, RuleBase: (*system.RuleBase)(nil)}

	ptr8729084064 := &system.Array_rule{Base: ptr8729234048, RuleBase: ptr8728989120, Items: ptr8728908960, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728990784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233536 := &system.Base{Context: ptr8728990784, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729088224 := &system.Reference_rule{Base: ptr8729233536, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728986400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233664 := &system.Base{Context: ptr8728986400, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728986336 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728910048 := &system.JsonString_rule{Base: ptr8729233664, RuleBase: ptr8728986336}

	ptr8729217264 := &system.Type{Base: ptr8729233408, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"description": ptr8728905104, "context": ptr8728906992, "rules": ptr8729084064, "type": ptr8729088224, "id": ptr8728910048}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729217088)

	system.RegisterType("kego.io/system:@base", ptr8729215328)

	system.RegisterType("kego.io/system:@bool", ptr8729217616)

	system.RegisterType("kego.io/system:@context", ptr8729217968)

	system.RegisterType("kego.io/system:@int", ptr8729218320)

	system.RegisterType("kego.io/system:@map", ptr8729218672)

	system.RegisterType("kego.io/system:@number", ptr8729219024)

	system.RegisterType("kego.io/system:@property", ptr8729219376)

	system.RegisterType("kego.io/system:@reference", ptr8729219728)

	system.RegisterType("kego.io/system:@rule", ptr8729220080)

	system.RegisterType("kego.io/system:@ruleBase", ptr8729220432)

	system.RegisterType("kego.io/system:@string", ptr8729220784)

	system.RegisterType("kego.io/system:@type", ptr8729222192)

	system.RegisterType("kego.io/system:array", ptr8729215152)

	system.RegisterType("kego.io/system:base", ptr8729217264)

	system.RegisterType("kego.io/system:bool", ptr8729217440)

	system.RegisterType("kego.io/system:context", ptr8729217792)

	system.RegisterType("kego.io/system:int", ptr8729218144)

	system.RegisterType("kego.io/system:map", ptr8729218496)

	system.RegisterType("kego.io/system:number", ptr8729218848)

	system.RegisterType("kego.io/system:property", ptr8729219200)

	system.RegisterType("kego.io/system:reference", ptr8729219552)

	system.RegisterType("kego.io/system:rule", ptr8729219904)

	system.RegisterType("kego.io/system:ruleBase", ptr8729220256)

	system.RegisterType("kego.io/system:string", ptr8729220608)

	system.RegisterType("kego.io/system:type", ptr8729221840)

}
