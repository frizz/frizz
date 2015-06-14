package types

import (
	"kego.io/system"
)

func init() {

	ptr8728989152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569280 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989152, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728988768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569408 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988768, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569536 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988672, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729084544 := &system.Reference_rule{Base: ptr8729569536, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729102928 := &system.Property{Base: ptr8729569408, Defaulter: true, Item: ptr8729084544, Optional: true}

	ptr8729219728 := &system.Type{Base: ptr8729569280, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729102928}, Rule: (*system.Type)(nil)}

	ptr8729601504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236352 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601504, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729601280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236480 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601280, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729601056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237376 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601056, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729601024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237504 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729084464 := &system.Int_rule{Base: ptr8729237504, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729102736 := &system.Property{Base: ptr8729237376, Defaulter: false, Item: ptr8729084464, Optional: true}

	ptr8729601152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236608 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601152, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729601120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236736 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601120, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729085104 := &system.Int_rule{Base: ptr8729236736, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729107632 := &system.Property{Base: ptr8729236608, Defaulter: true, Item: ptr8729085104, Optional: true}

	ptr8729601472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236864 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601472, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729601440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236992 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601440, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729085184 := &system.Int_rule{Base: ptr8729236992, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729107776 := &system.Property{Base: ptr8729236864, Defaulter: false, Item: ptr8729085184, Optional: true}

	ptr8729600544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237120 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729600544, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729600480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237248 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729600480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729084144 := &system.Int_rule{Base: ptr8729237248, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729107872 := &system.Property{Base: ptr8729237120, Defaulter: false, Item: ptr8729084144, Optional: true}

	ptr8729218320 := &system.Type{Base: ptr8729236480, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maximum": ptr8729102736, "default": ptr8729107632, "multipleOf": ptr8729107776, "minimum": ptr8729107872}, Rule: (*system.Type)(nil)}

	ptr8729218144 := &system.Type{Base: ptr8729236352, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729218320}

	ptr8729607584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571328 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607584, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729607424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573120 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607424, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729607392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573248 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607392, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729222192 := &system.String_rule{Base: ptr8729573248, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729477744 := &system.Property{Base: ptr8729573120, Defaulter: false, Item: ptr8729222192, Optional: true}

	ptr8729603936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571456 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603936, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729603904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571584 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603904, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729221488 := &system.String_rule{Base: ptr8729571584, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729108064 := &system.Property{Base: ptr8729571456, Defaulter: true, Item: ptr8729221488, Optional: true}

	ptr8729604704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571712 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729604704, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729604576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571840 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729604576, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729604544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571968 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729604544, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729221664 := &system.String_rule{Base: ptr8729571968, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729085984 := &system.Array_rule{Base: ptr8729571840, RuleBase: (*system.RuleBase)(nil), Items: ptr8729221664, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729477120 := &system.Property{Base: ptr8729571712, Defaulter: false, Item: ptr8729085984, Optional: true}

	ptr8729605152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572096 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605152, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729605088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572224 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605088, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729086144 := &system.Int_rule{Base: ptr8729572224, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729477264 := &system.Property{Base: ptr8729572096, Defaulter: false, Item: ptr8729086144, Optional: true}

	ptr8729605600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572352 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605600, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729605536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572480 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605536, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729086304 := &system.Int_rule{Base: ptr8729572480, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729477360 := &system.Property{Base: ptr8729572352, Defaulter: false, Item: ptr8729086304, Optional: true}

	ptr8729606048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572608 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606048, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729605984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572736 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605984, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729221840 := &system.String_rule{Base: ptr8729572736, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729477456 := &system.Property{Base: ptr8729572608, Defaulter: false, Item: ptr8729221840, Optional: true}

	ptr8729606528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572864 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606528, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729606496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572992 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729222016 := &system.String_rule{Base: ptr8729572992, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729477600 := &system.Property{Base: ptr8729572864, Defaulter: false, Item: ptr8729222016, Optional: true}

	ptr8729221312 := &system.Type{Base: ptr8729571328, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"format": ptr8729477744, "default": ptr8729108064, "enum": ptr8729477120, "minLength": ptr8729477264, "maxLength": ptr8729477360, "equal": ptr8729477456, "pattern": ptr8729477600}, Rule: (*system.Type)(nil)}

	ptr8728991904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233920 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991904, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728991040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234816 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991040, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234944 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991008, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728910368 := &system.Context_rule{Base: ptr8729234944, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105712 := &system.Property{Base: ptr8729234816, Defaulter: false, Item: ptr8728910368, Optional: true}

	ptr8728991776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235072 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991776, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235200 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728991648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235328 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991648, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728911456 := &system.Rule_rule{Base: ptr8729235328, RuleBase: (*system.RuleBase)(nil)}

	ptr8729084384 := &system.Array_rule{Base: ptr8729235200, RuleBase: (*system.RuleBase)(nil), Items: ptr8728911456, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729106048 := &system.Property{Base: ptr8729235072, Defaulter: false, Item: ptr8729084384, Optional: true}

	ptr8728989696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234048 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989696, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728989664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234176 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989664, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729084224 := &system.Reference_rule{Base: ptr8729234176, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729104656 := &system.Property{Base: ptr8729234048, Defaulter: false, Item: ptr8729084224, Optional: false}

	ptr8728990176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234304 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728990176, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728990144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234432 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728990144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728909584 := &system.JsonString_rule{Base: ptr8729234432, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105088 := &system.Property{Base: ptr8729234304, Defaulter: false, Item: ptr8728909584, Optional: true}

	ptr8728990624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234560 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728990624, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728990592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729234688 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728990592, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728910128 := &system.JsonString_rule{Base: ptr8729234688, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105664 := &system.Property{Base: ptr8729234560, Defaulter: false, Item: ptr8728910128, Optional: true}

	ptr8729215328 := &system.Type{Base: ptr8729233920, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"context": ptr8729105712, "rules": ptr8729106048, "type": ptr8729104656, "id": ptr8729105088, "description": ptr8729105664}, Rule: (*system.Type)(nil)}

	ptr8729600032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236096 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729600032, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217792 := &system.Type{Base: ptr8729236096, Basic: false, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729603552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237632 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603552, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729603456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237760 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603456, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729602976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238144 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729602976, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729602944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238272 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729602944, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729084624 := &system.Int_rule{Base: ptr8729238272, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729104512 := &system.Property{Base: ptr8729238144, Defaulter: false, Item: ptr8729084624, Optional: true}

	ptr8729603328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238400 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603328, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729603296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238528 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729603296, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729084704 := &system.Int_rule{Base: ptr8729238528, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729104704 := &system.Property{Base: ptr8729238400, Defaulter: false, Item: ptr8729084704, Optional: true}

	ptr8729602624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729237888 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729602624, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729602592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238016 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729602592, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729445680 := &system.Rule_rule{Base: ptr8729238016, RuleBase: (*system.RuleBase)(nil)}

	ptr8729104416 := &system.Property{Base: ptr8729237888, Defaulter: false, Item: ptr8729445680, Optional: false}

	ptr8729218672 := &system.Type{Base: ptr8729237760, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"minItems": ptr8729104512, "maxItems": ptr8729104704, "items": ptr8729104416}, Rule: (*system.Type)(nil)}

	ptr8729218496 := &system.Type{Base: ptr8729237632, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729218672}

	ptr8729607360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238784 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607360, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729607232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567872 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607232, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729607200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568000 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607200, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729607072 := &system.Bool_rule{Base: ptr8729568000, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: true}}

	ptr8729106864 := &system.Property{Base: ptr8729567872, Defaulter: false, Item: ptr8729607072, Optional: true}

	ptr8729605120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238912 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605120, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729604960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239040 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729604960, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729092384 := &system.Number_rule{Base: ptr8729239040, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729105184 := &system.Property{Base: ptr8729238912, Defaulter: true, Item: ptr8729092384, Optional: true}

	ptr8729605568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239168 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605568, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729605408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239296 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605408, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729092480 := &system.Number_rule{Base: ptr8729239296, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729106288 := &system.Property{Base: ptr8729239168, Defaulter: false, Item: ptr8729092480, Optional: true}

	ptr8729606016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729239424 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606016, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729605856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567232 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729605856, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729092576 := &system.Number_rule{Base: ptr8729567232, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729106384 := &system.Property{Base: ptr8729239424, Defaulter: false, Item: ptr8729092576, Optional: true}

	ptr8729606400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567360 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606400, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729606368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567488 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606368, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729606240 := &system.Bool_rule{Base: ptr8729567488, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: true}}

	ptr8729106528 := &system.Property{Base: ptr8729567360, Defaulter: false, Item: ptr8729606240, Optional: true}

	ptr8729606848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567616 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606848, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729606688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567744 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729606688, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729092672 := &system.Number_rule{Base: ptr8729567744, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729106672 := &system.Property{Base: ptr8729567616, Defaulter: false, Item: ptr8729092672, Optional: true}

	ptr8729219024 := &system.Type{Base: ptr8729238784, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"exclusiveMaximum": ptr8729106864, "default": ptr8729105184, "multipleOf": ptr8729106288, "minimum": ptr8729106384, "exclusiveMinimum": ptr8729106528, "maximum": ptr8729106672}, Rule: (*system.Type)(nil)}

	ptr8729601856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570688 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601856, Description: "All rules should extend this type.", Id: "selector", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729601600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570816 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601600, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729601568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570944 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729601568, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729220784 := &system.String_rule{Base: ptr8729570944, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: ":root", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729106336 := &system.Property{Base: ptr8729570816, Defaulter: false, Item: ptr8729220784, Optional: false}

	ptr8729220608 := &system.Type{Base: ptr8729570688, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"selector": ptr8729106336}, Rule: (*system.Type)(nil)}

	ptr8728993056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569920 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728993056, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728991968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570048 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991968, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570176 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991936, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729446336 := &system.JsonBool_rule{Base: ptr8729570176, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105616 := &system.Property{Base: ptr8729570048, Defaulter: false, Item: ptr8729446336, Optional: true}

	ptr8728992736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570304 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728992736, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570432 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728992672, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729446592 := &system.JsonString_rule{Base: ptr8729570432, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105760 := &system.Property{Base: ptr8729570304, Defaulter: false, Item: ptr8729446592, Optional: true}

	ptr8729220256 := &system.Type{Base: ptr8729569920, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8729105616, "selector": ptr8729105760}, Rule: (*system.Type)(nil)}

	ptr8729607808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571200 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607808, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729221136 := &system.Type{Base: ptr8729571200, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729221312}

	ptr8728988736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729231360 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988736, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728988608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729232640 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988608, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728988384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233664 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988384, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233792 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988352, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728988064 := &system.Bool_rule{Base: ptr8729233792, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: true}}

	ptr8729102544 := &system.Property{Base: ptr8729233664, Defaulter: false, Item: ptr8728988064, Optional: true}

	ptr8728988544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729232896 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988544, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233024 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988512, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728907952 := &system.Rule_rule{Base: ptr8729233024, RuleBase: (*system.RuleBase)(nil)}

	ptr8729105472 := &system.Property{Base: ptr8729232896, Defaulter: false, Item: ptr8728907952, Optional: false}

	ptr8728986592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233152 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728986592, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728986528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233280 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728986528, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729089264 := &system.Int_rule{Base: ptr8729233280, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729106144 := &system.Property{Base: ptr8729233152, Defaulter: false, Item: ptr8729089264, Optional: true}

	ptr8728987424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233408 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728987424, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728987360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729233536 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728987360, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729084064 := &system.Int_rule{Base: ptr8729233536, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729102160 := &system.Property{Base: ptr8729233408, Defaulter: false, Item: ptr8729084064, Optional: true}

	ptr8729217088 := &system.Type{Base: ptr8729232640, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"uniqueItems": ptr8729102544, "items": ptr8729105472, "minItems": ptr8729106144, "maxItems": ptr8729102160}, Rule: (*system.Type)(nil)}

	ptr8729215152 := &system.Type{Base: ptr8729231360, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729217088}

	ptr8728993248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235584 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728993248, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728993152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235712 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728993152, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728993024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235840 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728993024, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235968 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728992992, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728992896 := &system.Bool_rule{Base: ptr8729235968, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729107008 := &system.Property{Base: ptr8729235840, Defaulter: true, Item: ptr8728992896, Optional: true}

	ptr8729217616 := &system.Type{Base: ptr8729235712, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729107008}, Rule: (*system.Type)(nil)}

	ptr8729217440 := &system.Type{Base: ptr8729235584, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729217616}

	ptr8728992096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729235456 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728992096, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217264 := &system.Type{Base: ptr8729235456, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729600224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729236224 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729600224, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729217968 := &system.Type{Base: ptr8729236224, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729607456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729238656 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729607456, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729218848 := &system.Type{Base: ptr8729238656, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729219024}

	ptr8728993696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570560 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728993696, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220432 := &system.Type{Base: ptr8729570560, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729488928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573376 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488928, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729485888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573504 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729485888, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729485856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573632 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729485856, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729452048 := &system.JsonBool_rule{Base: ptr8729573632, RuleBase: (*system.RuleBase)(nil)}

	ptr8729477984 := &system.Property{Base: ptr8729573504, Defaulter: false, Item: ptr8729452048, Optional: true}

	ptr8729486432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573760 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486432, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729486336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573888 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486336, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729486304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574016 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486304, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729086704 := &system.Reference_rule{Base: ptr8729574016, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729086624 := &system.Array_rule{Base: ptr8729573888, RuleBase: (*system.RuleBase)(nil), Items: ptr8729086704, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729478032 := &system.Property{Base: ptr8729573760, Defaulter: false, Item: ptr8729086624, Optional: true}

	ptr8729486944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574144 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486944, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729486912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574272 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486912, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729086784 := &system.Reference_rule{Base: ptr8729574272, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729478128 := &system.Property{Base: ptr8729574144, Defaulter: false, Item: ptr8729086784, Optional: true}

	ptr8729485696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574400 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729485696, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729485568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574528 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729485568, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729485504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574656 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729485504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729086944 := &system.Reference_rule{Base: ptr8729574656, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729086864 := &system.Array_rule{Base: ptr8729574528, RuleBase: (*system.RuleBase)(nil), Items: ptr8729086944, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729478176 := &system.Property{Base: ptr8729574400, Defaulter: false, Item: ptr8729086864, Optional: true}

	ptr8729486688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574784 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486688, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729486656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574912 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729486656, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729222544 := &system.String_rule{Base: ptr8729574912, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729477312 := &system.Property{Base: ptr8729574784, Defaulter: false, Item: ptr8729222544, Optional: true}

	ptr8729487360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729575040 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729487360, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729487328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729575168 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729487328, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728908032 := &system.JsonBool_rule{Base: ptr8729575168, RuleBase: (*system.RuleBase)(nil)}

	ptr8729477408 := &system.Property{Base: ptr8729575040, Defaulter: false, Item: ptr8728908032, Optional: true}

	ptr8729487776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729575296 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729487776, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729487744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729468928 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729487744, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728908640 := &system.JsonBool_rule{Base: ptr8729468928, RuleBase: (*system.RuleBase)(nil)}

	ptr8729477552 := &system.Property{Base: ptr8729575296, Defaulter: false, Item: ptr8728908640, Optional: true}

	ptr8729488384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469056 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488384, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729488352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469184 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488352, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729488320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469312 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488320, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8728909984 := &system.Property_rule{Base: ptr8729469312, RuleBase: (*system.RuleBase)(nil)}

	ptr8729395392 := &system.Map_rule{Base: ptr8729469184, RuleBase: (*system.RuleBase)(nil), Items: ptr8728909984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729477792 := &system.Property{Base: ptr8729469056, Defaulter: false, Item: ptr8729395392, Optional: true}

	ptr8729488800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469440 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488800, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729488768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469568 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729488768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8728910880 := &system.Type_rule{Base: ptr8729469568, RuleBase: (*system.RuleBase)(nil)}

	ptr8729478848 := &system.Property{Base: ptr8729469440, Defaulter: false, Item: ptr8728910880, Optional: true}

	ptr8729222368 := &system.Type{Base: ptr8729573376, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"basic": ptr8729477984, "embed": ptr8729478032, "extends": ptr8729478128, "is": ptr8729478176, "native": ptr8729477312, "interface": ptr8729477408, "exclude": ptr8729477552, "properties": ptr8729477792, "rule": ptr8729478848}, Rule: (*system.Type)(nil)}

	ptr8729489344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729469696 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729489344, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729222720 := &system.Type{Base: ptr8729469696, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728991104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569792 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728991104, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220080 := &system.Type{Base: ptr8729569792, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728990464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569664 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728990464, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729219904 := &system.Type{Base: ptr8729569664, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729602496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571072 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729602496, Description: "Automatically created basic rule for selector", Id: "@selector", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729220960 := &system.Type{Base: ptr8729571072, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728989120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568128 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989120, Description: "This is a property of a type", Id: "property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728986656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568256 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728986656, Description: "This specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728986432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568384 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728986432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729449248 := &system.JsonBool_rule{Base: ptr8729568384, RuleBase: (*system.RuleBase)(nil)}

	ptr8729107344 := &system.Property{Base: ptr8729568256, Defaulter: false, Item: ptr8729449248, Optional: true}

	ptr8728988160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568512 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988160, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568640 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988096, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729449792 := &system.JsonBool_rule{Base: ptr8729568640, RuleBase: (*system.RuleBase)(nil)}

	ptr8729108160 := &system.Property{Base: ptr8729568512, Defaulter: false, Item: ptr8729449792, Optional: true}

	ptr8728988864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568768 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988864, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568896 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728988832, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729450032 := &system.Rule_rule{Base: ptr8729568896, RuleBase: (*system.RuleBase)(nil)}

	ptr8729108208 := &system.Property{Base: ptr8729568768, Defaulter: false, Item: ptr8729450032, Optional: false}

	ptr8729219200 := &system.Type{Base: ptr8729568128, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8729107344, "defaulter": ptr8729108160, "item": ptr8729108208}, Rule: (*system.Type)(nil)}

	ptr8728989312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569152 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989312, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729219552 := &system.Type{Base: ptr8729569152, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729219728}

	ptr8728989760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569024 := &system.Base{Base: (*system.Base)(nil), Context: ptr8728989760, Description: "Automatically created basic rule for property", Id: "@property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729219376 := &system.Type{Base: ptr8729569024, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729217088)

	system.RegisterType("kego.io/system:@base", ptr8729217264)

	system.RegisterType("kego.io/system:@bool", ptr8729217616)

	system.RegisterType("kego.io/system:@context", ptr8729217968)

	system.RegisterType("kego.io/system:@int", ptr8729218320)

	system.RegisterType("kego.io/system:@map", ptr8729218672)

	system.RegisterType("kego.io/system:@number", ptr8729219024)

	system.RegisterType("kego.io/system:@property", ptr8729219376)

	system.RegisterType("kego.io/system:@reference", ptr8729219728)

	system.RegisterType("kego.io/system:@rule", ptr8729220080)

	system.RegisterType("kego.io/system:@ruleBase", ptr8729220432)

	system.RegisterType("kego.io/system:@selector", ptr8729220960)

	system.RegisterType("kego.io/system:@string", ptr8729221312)

	system.RegisterType("kego.io/system:@type", ptr8729222720)

	system.RegisterType("kego.io/system:array", ptr8729215152)

	system.RegisterType("kego.io/system:base", ptr8729215328)

	system.RegisterType("kego.io/system:bool", ptr8729217440)

	system.RegisterType("kego.io/system:context", ptr8729217792)

	system.RegisterType("kego.io/system:int", ptr8729218144)

	system.RegisterType("kego.io/system:map", ptr8729218496)

	system.RegisterType("kego.io/system:number", ptr8729218848)

	system.RegisterType("kego.io/system:property", ptr8729219200)

	system.RegisterType("kego.io/system:reference", ptr8729219552)

	system.RegisterType("kego.io/system:rule", ptr8729219904)

	system.RegisterType("kego.io/system:ruleBase", ptr8729220256)

	system.RegisterType("kego.io/system:selector", ptr8729220608)

	system.RegisterType("kego.io/system:string", ptr8729221136)

	system.RegisterType("kego.io/system:type", ptr8729222368)

}
