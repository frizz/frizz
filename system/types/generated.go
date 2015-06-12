package types

import "kego.io/system"

func init() {

	ptr8729589472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095344 := &system.Object{Context: ptr8729589472, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729356688 := &system.Type{Object: ptr8729095344, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729584768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324720 := &system.Object{Context: ptr8729584768, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729584448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324832 := &system.Object{Context: ptr8729584448, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729584320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324944 := &system.Object{Context: ptr8729584320, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729584288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325056 := &system.Object{Context: ptr8729584288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728909856 := &system.Bool_rule{Object: ptr8729325056, Default: system.Bool{Value: false, Exists: false}}

	ptr8729147296 := &system.Property{Object: ptr8729324944, Defaulter: true, Item: ptr8728909856, Optional: true}

	ptr8729358416 := &system.Type{Object: ptr8729324832, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729147296}, Rule: (*system.Type)(nil)}

	ptr8729358272 := &system.Type{Object: ptr8729324720, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729358416}

	ptr8728992096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729328080 := &system.Object{Context: ptr8728992096, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728991936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729329536 := &system.Object{Context: ptr8728991936, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092432 := &system.Object{Context: ptr8728991584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728991552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092544 := &system.Object{Context: ptr8728991552, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076472 := &system.Rule_rule{Object: ptr8729092544}

	ptr8729387200 := &system.Map_rule{Object: ptr8729092432, Items: ptr8729076472, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729146000 := &system.Property{Object: ptr8729329536, Defaulter: false, Item: ptr8729387200, Optional: true}

	ptr8728986208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729328304 := &system.Object{Context: ptr8728986208, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728986144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326400 := &system.Object{Context: ptr8728986144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729388480 := &system.Reference_rule{Object: ptr8729326400, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729147824 := &system.Property{Object: ptr8729328304, Defaulter: false, Item: ptr8729388480, Optional: false}

	ptr8728988064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326512 := &system.Object{Context: ptr8728988064, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728987520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729328752 := &system.Object{Context: ptr8728987520, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729075816 := &system.JsonString_rule{Object: ptr8729328752}

	ptr8729142832 := &system.Property{Object: ptr8729326512, Defaulter: false, Item: ptr8729075816, Optional: true}

	ptr8728988512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729328864 := &system.Object{Context: ptr8728988512, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728988480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729328976 := &system.Object{Context: ptr8728988480, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729076048 := &system.JsonString_rule{Object: ptr8729328976}

	ptr8729145184 := &system.Property{Object: ptr8729328864, Defaulter: false, Item: ptr8729076048, Optional: true}

	ptr8728989920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729329088 := &system.Object{Context: ptr8728989920, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728989888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729329424 := &system.Object{Context: ptr8728989888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8729076136 := &system.Context_rule{Object: ptr8729329424}

	ptr8729145280 := &system.Property{Object: ptr8729329088, Defaulter: false, Item: ptr8729076136, Optional: true}

	ptr8729355680 := &system.Type{Object: ptr8729328080, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"rules": ptr8729146000, "type": ptr8729147824, "id": ptr8729142832, "description": ptr8729145184, "context": ptr8729145280}, Rule: (*system.Type)(nil)}

	ptr8728993216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322480 := &system.Object{Context: ptr8728993216, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728991904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322928 := &system.Object{Context: ptr8728991904, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323264 := &system.Object{Context: ptr8728991872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076144 := &system.Rule_rule{Object: ptr8729323264}

	ptr8729145664 := &system.Property{Object: ptr8729322928, Defaulter: false, Item: ptr8729076144, Optional: false}

	ptr8728992320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323376 := &system.Object{Context: ptr8728992320, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323824 := &system.Object{Context: ptr8728992288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729473248 := &system.Int_rule{Object: ptr8729323824, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146240 := &system.Property{Object: ptr8729323376, Defaulter: false, Item: ptr8729473248, Optional: true}

	ptr8728992672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323936 := &system.Object{Context: ptr8728992672, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324048 := &system.Object{Context: ptr8728992640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729473328 := &system.Int_rule{Object: ptr8729324048, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146576 := &system.Property{Object: ptr8729323936, Defaulter: false, Item: ptr8729473328, Optional: true}

	ptr8728993088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324160 := &system.Object{Context: ptr8728993088, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324608 := &system.Object{Context: ptr8728993056, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728908816 := &system.Bool_rule{Object: ptr8729324608, Default: system.Bool{Value: false, Exists: true}}

	ptr8729146672 := &system.Property{Object: ptr8729324160, Defaulter: false, Item: ptr8728908816, Optional: true}

	ptr8729357984 := &system.Type{Object: ptr8729322480, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8729145664, "minItems": ptr8729146240, "maxItems": ptr8729146576, "uniqueItems": ptr8729146672}, Rule: (*system.Type)(nil)}

	ptr8729466976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729097808 := &system.Object{Context: ptr8729466976, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729464384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098256 := &system.Object{Context: ptr8729464384, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729464288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098368 := &system.Object{Context: ptr8729464288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729464256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098480 := &system.Object{Context: ptr8729464256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729388352 := &system.Reference_rule{Object: ptr8729098480, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729388224 := &system.Array_rule{Object: ptr8729098368, Items: ptr8729388352, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729107152 := &system.Property{Object: ptr8729098256, Defaulter: false, Item: ptr8729388224, Optional: true}

	ptr8729465024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098816 := &system.Object{Context: ptr8729465024, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729464992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098928 := &system.Object{Context: ptr8729464992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729210144 := &system.String_rule{Object: ptr8729098928, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729107248 := &system.Property{Object: ptr8729098816, Defaulter: false, Item: ptr8729210144, Optional: true}

	ptr8729465472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729099264 := &system.Object{Context: ptr8729465472, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729465440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729099376 := &system.Object{Context: ptr8729465440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076784 := &system.JsonBool_rule{Object: ptr8729099376}

	ptr8729107296 := &system.Property{Object: ptr8729099264, Defaulter: false, Item: ptr8729076784, Optional: true}

	ptr8729465888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729099600 := &system.Object{Context: ptr8729465888, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729465856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729099824 := &system.Object{Context: ptr8729465856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076808 := &system.JsonBool_rule{Object: ptr8729099824}

	ptr8729107344 := &system.Property{Object: ptr8729099600, Defaulter: false, Item: ptr8729076808, Optional: true}

	ptr8729466432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729099936 := &system.Object{Context: ptr8729466432, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729466400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729100048 := &system.Object{Context: ptr8729466400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729466368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729329200 := &system.Object{Context: ptr8729466368, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8729076840 := &system.Property_rule{Object: ptr8729329200}

	ptr8729388544 := &system.Map_rule{Object: ptr8729100048, Items: ptr8729076840, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729107440 := &system.Property{Object: ptr8729099936, Defaulter: false, Item: ptr8729388544, Optional: true}

	ptr8729466848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729329312 := &system.Object{Context: ptr8729466848, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729466816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729274784 := &system.Object{Context: ptr8729466816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729076976 := &system.Type_rule{Object: ptr8729274784}

	ptr8729108400 := &system.Property{Object: ptr8729329312, Defaulter: false, Item: ptr8729076976, Optional: true}

	ptr8729463840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098032 := &system.Object{Context: ptr8729463840, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729463808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729098144 := &system.Object{Context: ptr8729463808, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729388032 := &system.Reference_rule{Object: ptr8729098144, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729107104 := &system.Property{Object: ptr8729098032, Defaulter: false, Item: ptr8729388032, Optional: true}

	ptr8729355968 := &system.Type{Object: ptr8729097808, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"is": ptr8729107152, "native": ptr8729107248, "interface": ptr8729107296, "exclude": ptr8729107344, "properties": ptr8729107440, "rule": ptr8729108400, "extends": ptr8729107104}, Rule: (*system.Type)(nil)}

	ptr8729585376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092768 := &system.Object{Context: ptr8729585376, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729585120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093776 := &system.Object{Context: ptr8729585120, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729584864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093888 := &system.Object{Context: ptr8729584864, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076712 := &system.Rule_rule{Object: ptr8729093888}

	ptr8729148112 := &system.Property{Object: ptr8729093776, Defaulter: false, Item: ptr8729076712, Optional: false}

	ptr8728993408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092880 := &system.Object{Context: ptr8728993408, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093216 := &system.Object{Context: ptr8728993344, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076528 := &system.JsonBool_rule{Object: ptr8729093216}

	ptr8729147488 := &system.Property{Object: ptr8729092880, Defaulter: false, Item: ptr8729076528, Optional: true}

	ptr8729584064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093328 := &system.Object{Context: ptr8729584064, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729584032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093552 := &system.Object{Context: ptr8729584032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076688 := &system.JsonBool_rule{Object: ptr8729093552}

	ptr8729147920 := &system.Property{Object: ptr8729093328, Defaulter: false, Item: ptr8729076688, Optional: true}

	ptr8729359280 := &system.Type{Object: ptr8729092768, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"item": ptr8729148112, "optional": ptr8729147488, "defaulter": ptr8729147920}, Rule: (*system.Type)(nil)}

	ptr8729587584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729094672 := &system.Object{Context: ptr8729587584, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729587360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729094784 := &system.Object{Context: ptr8729587360, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729587328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729094896 := &system.Object{Context: ptr8729587328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729390656 := &system.Reference_rule{Object: ptr8729094896, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729148352 := &system.Property{Object: ptr8729094784, Defaulter: true, Item: ptr8729390656, Optional: true}

	ptr8729356256 := &system.Type{Object: ptr8729094672, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729148352}, Rule: (*system.Type)(nil)}

	ptr8728993536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322368 := &system.Object{Context: ptr8728993536, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729357840 := &system.Type{Object: ptr8729322368, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729357984}

	ptr8729588288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323712 := &system.Object{Context: ptr8729588288, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729587456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324272 := &system.Object{Context: ptr8729587456, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729587424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729324496 := &system.Object{Context: ptr8729587424, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076344 := &system.Rule_rule{Object: ptr8729324496}

	ptr8729144896 := &system.Property{Object: ptr8729324272, Defaulter: false, Item: ptr8729076344, Optional: false}

	ptr8729587808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325504 := &system.Object{Context: ptr8729587808, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729587776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325728 := &system.Object{Context: ptr8729587776, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469328 := &system.Int_rule{Object: ptr8729325728, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729145232 := &system.Property{Object: ptr8729325504, Defaulter: false, Item: ptr8729469328, Optional: true}

	ptr8729588160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325840 := &system.Object{Context: ptr8729588160, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729588128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325952 := &system.Object{Context: ptr8729588128, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469408 := &system.Int_rule{Object: ptr8729325952, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729145328 := &system.Property{Object: ptr8729325840, Defaulter: false, Item: ptr8729469408, Optional: true}

	ptr8729354672 := &system.Type{Object: ptr8729323712, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8729144896, "minItems": ptr8729145232, "maxItems": ptr8729145328}, Rule: (*system.Type)(nil)}

	ptr8729585312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325168 := &system.Object{Context: ptr8729585312, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729358704 := &system.Type{Object: ptr8729325168, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729585504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325280 := &system.Object{Context: ptr8729585504, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729358848 := &system.Type{Object: ptr8729325280, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728992384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092656 := &system.Object{Context: ptr8728992384, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729354384 := &system.Type{Object: ptr8729092656, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729588800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095008 := &system.Object{Context: ptr8729588800, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729356544 := &system.Type{Object: ptr8729095008, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729462880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095456 := &system.Object{Context: ptr8729462880, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729462560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095568 := &system.Object{Context: ptr8729462560, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729590560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095680 := &system.Object{Context: ptr8729590560, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729590528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096240 := &system.Object{Context: ptr8729590528, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729210464 := &system.String_rule{Object: ptr8729096240, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729148880 := &system.Property{Object: ptr8729095680, Defaulter: true, Item: ptr8729210464, Optional: true}

	ptr8729591616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096352 := &system.Object{Context: ptr8729591616, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729591456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096912 := &system.Object{Context: ptr8729591456, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729591392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729097024 := &system.Object{Context: ptr8729591392, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729211424 := &system.String_rule{Object: ptr8729097024, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729389248 := &system.Array_rule{Object: ptr8729096912, Items: ptr8729211424, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729149264 := &system.Property{Object: ptr8729096352, Defaulter: false, Item: ptr8729389248, Optional: true}

	ptr8729460896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729097136 := &system.Object{Context: ptr8729460896, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729460864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092096 := &system.Object{Context: ptr8729460864, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469248 := &system.Int_rule{Object: ptr8729092096, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729102832 := &system.Property{Object: ptr8729097136, Defaulter: false, Item: ptr8729469248, Optional: true}

	ptr8729461216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092320 := &system.Object{Context: ptr8729461216, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729461184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729092992 := &system.Object{Context: ptr8729461184, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469888 := &system.Int_rule{Object: ptr8729092992, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729102208 := &system.Property{Object: ptr8729092320, Defaulter: false, Item: ptr8729469888, Optional: true}

	ptr8729461536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729093104 := &system.Object{Context: ptr8729461536, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729461504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729095792 := &system.Object{Context: ptr8729461504, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729206784 := &system.String_rule{Object: ptr8729095792, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729102304 := &system.Property{Object: ptr8729093104, Defaulter: false, Item: ptr8729206784, Optional: true}

	ptr8729461856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096128 := &system.Object{Context: ptr8729461856, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729461824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096576 := &system.Object{Context: ptr8729461824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729209024 := &system.String_rule{Object: ptr8729096576, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729102496 := &system.Property{Object: ptr8729096128, Defaulter: false, Item: ptr8729209024, Optional: true}

	ptr8729462432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729096688 := &system.Object{Context: ptr8729462432, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729462400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729097472 := &system.Object{Context: ptr8729462400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729209664 := &system.String_rule{Object: ptr8729097472, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729102640 := &system.Property{Object: ptr8729096688, Defaulter: false, Item: ptr8729209664, Optional: true}

	ptr8729357120 := &system.Type{Object: ptr8729095568, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729148880, "enum": ptr8729149264, "minLength": ptr8729102832, "maxLength": ptr8729102208, "equal": ptr8729102304, "pattern": ptr8729102496, "format": ptr8729102640}, Rule: (*system.Type)(nil)}

	ptr8729356976 := &system.Type{Object: ptr8729095456, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729357120}

	ptr8729586496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729325392 := &system.Object{Context: ptr8729586496, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729586176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729321472 := &system.Object{Context: ptr8729586176, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729584544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729321584 := &system.Object{Context: ptr8729584544, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729584512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729321696 := &system.Object{Context: ptr8729584512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729468928 := &system.Int_rule{Object: ptr8729321696, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729142592 := &system.Property{Object: ptr8729321584, Defaulter: true, Item: ptr8729468928, Optional: true}

	ptr8729585088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729321808 := &system.Object{Context: ptr8729585088, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729584960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322144 := &system.Object{Context: ptr8729584960, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469008 := &system.Int_rule{Object: ptr8729322144, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729143024 := &system.Property{Object: ptr8729321808, Defaulter: false, Item: ptr8729469008, Optional: true}

	ptr8729585728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322592 := &system.Object{Context: ptr8729585728, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729585696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729322816 := &system.Object{Context: ptr8729585696, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469088 := &system.Int_rule{Object: ptr8729322816, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729143216 := &system.Property{Object: ptr8729322592, Defaulter: false, Item: ptr8729469088, Optional: true}

	ptr8729586048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323040 := &system.Object{Context: ptr8729586048, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729586016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323152 := &system.Object{Context: ptr8729586016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469168 := &system.Int_rule{Object: ptr8729323152, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729143600 := &system.Property{Object: ptr8729323040, Defaulter: false, Item: ptr8729469168, Optional: true}

	ptr8729354240 := &system.Type{Object: ptr8729321472, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729142592, "multipleOf": ptr8729143024, "minimum": ptr8729143216, "maximum": ptr8729143600}, Rule: (*system.Type)(nil)}

	ptr8729359424 := &system.Type{Object: ptr8729325392, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729354240}

	ptr8729588608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729323488 := &system.Object{Context: ptr8729588608, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729354528 := &system.Type{Object: ptr8729323488, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729354672}

	ptr8728990016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326064 := &system.Object{Context: ptr8728990016, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728987488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326176 := &system.Object{Context: ptr8728987488, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729590016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326288 := &system.Object{Context: ptr8729590016, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729589856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326624 := &system.Object{Context: ptr8729589856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729469568 := &system.Number_rule{Object: ptr8729326624, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729145808 := &system.Property{Object: ptr8729326288, Defaulter: true, Item: ptr8729469568, Optional: true}

	ptr8729590464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326848 := &system.Object{Context: ptr8729590464, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729590304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729326960 := &system.Object{Context: ptr8729590304, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729469648 := &system.Number_rule{Object: ptr8729326960, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729146528 := &system.Property{Object: ptr8729326848, Defaulter: false, Item: ptr8729469648, Optional: true}

	ptr8729590912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327072 := &system.Object{Context: ptr8729590912, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729590752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327184 := &system.Object{Context: ptr8729590752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729469728 := &system.Number_rule{Object: ptr8729327184, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729146768 := &system.Property{Object: ptr8729327072, Defaulter: false, Item: ptr8729469728, Optional: true}

	ptr8729591264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327296 := &system.Object{Context: ptr8729591264, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729591232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327408 := &system.Object{Context: ptr8729591232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728911712 := &system.Bool_rule{Object: ptr8729327408, Default: system.Bool{Value: false, Exists: true}}

	ptr8729146912 := &system.Property{Object: ptr8729327296, Defaulter: false, Item: ptr8728911712, Optional: true}

	ptr8729591712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327520 := &system.Object{Context: ptr8729591712, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729591552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327632 := &system.Object{Context: ptr8729591552, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729469808 := &system.Number_rule{Object: ptr8729327632, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729147008 := &system.Property{Object: ptr8729327520, Defaulter: false, Item: ptr8729469808, Optional: true}

	ptr8728986272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327744 := &system.Object{Context: ptr8728986272, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728986240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729327856 := &system.Object{Context: ptr8728986240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729510144 := &system.Bool_rule{Object: ptr8729327856, Default: system.Bool{Value: false, Exists: true}}

	ptr8729147152 := &system.Property{Object: ptr8729327744, Defaulter: false, Item: ptr8729510144, Optional: true}

	ptr8729355104 := &system.Type{Object: ptr8729326176, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729145808, "multipleOf": ptr8729146528, "minimum": ptr8729146768, "exclusiveMinimum": ptr8729146912, "maximum": ptr8729147008, "exclusiveMaximum": ptr8729147152}, Rule: (*system.Type)(nil)}

	ptr8729354960 := &system.Type{Object: ptr8729326064, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729355104}

	ptr8729588096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729094448 := &system.Object{Context: ptr8729588096, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729356112 := &system.Type{Object: ptr8729094448, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729356256}

	ptr8729467392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729274896 := &system.Object{Context: ptr8729467392, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729357408 := &system.Type{Object: ptr8729274896, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729586240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729094224 := &system.Object{Context: ptr8729586240, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729355824 := &system.Type{Object: ptr8729094224, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729357984)

	system.RegisterType("kego.io/system:@bool", ptr8729358416)

	system.RegisterType("kego.io/system:@context", ptr8729358848)

	system.RegisterType("kego.io/system:@int", ptr8729354240)

	system.RegisterType("kego.io/system:@map", ptr8729354672)

	system.RegisterType("kego.io/system:@number", ptr8729355104)

	system.RegisterType("kego.io/system:@object", ptr8729354384)

	system.RegisterType("kego.io/system:@property", ptr8729355824)

	system.RegisterType("kego.io/system:@reference", ptr8729356256)

	system.RegisterType("kego.io/system:@rule", ptr8729356688)

	system.RegisterType("kego.io/system:@string", ptr8729357120)

	system.RegisterType("kego.io/system:@type", ptr8729357408)

	system.RegisterType("kego.io/system:array", ptr8729357840)

	system.RegisterType("kego.io/system:bool", ptr8729358272)

	system.RegisterType("kego.io/system:context", ptr8729358704)

	system.RegisterType("kego.io/system:int", ptr8729359424)

	system.RegisterType("kego.io/system:map", ptr8729354528)

	system.RegisterType("kego.io/system:number", ptr8729354960)

	system.RegisterType("kego.io/system:object", ptr8729355680)

	system.RegisterType("kego.io/system:property", ptr8729359280)

	system.RegisterType("kego.io/system:reference", ptr8729356112)

	system.RegisterType("kego.io/system:rule", ptr8729356544)

	system.RegisterType("kego.io/system:string", ptr8729356976)

	system.RegisterType("kego.io/system:type", ptr8729355968)

}
