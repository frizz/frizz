package types

import "kego.io/system"

func init() {

	ptr8728346112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631872 := &system.Object{Context: ptr8728346112, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650448 := &system.Type{Object: ptr8728631872, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728343584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631200 := &system.Object{Context: ptr8728343584, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728343104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631312 := &system.Object{Context: ptr8728343104, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728342880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632768 := &system.Object{Context: ptr8728342880, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728342848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728630752 := &system.Object{Context: ptr8728342848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728258192 := &system.Bool_rule{Object: ptr8728630752, Default: system.Bool{Value: false, Exists: true}}

	ptr8728497920 := &system.Property{Object: ptr8728632768, Defaulter: false, Item: ptr8728258192, Optional: true}

	ptr8728343872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631648 := &system.Object{Context: ptr8728343872, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631984 := &system.Object{Context: ptr8728343840, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728428976 := &system.Rule_rule{Object: ptr8728631984}

	ptr8728496912 := &system.Property{Object: ptr8728631648, Defaulter: false, Item: ptr8728428976, Optional: false}

	ptr8728344288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632096 := &system.Object{Context: ptr8728344288, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632432 := &system.Object{Context: ptr8728344256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728442496 := &system.Int_rule{Object: ptr8728632432, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497488 := &system.Property{Object: ptr8728632096, Defaulter: false, Item: ptr8728442496, Optional: true}

	ptr8728344640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632544 := &system.Object{Context: ptr8728344640, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632656 := &system.Object{Context: ptr8728344608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728442576 := &system.Int_rule{Object: ptr8728632656, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497824 := &system.Property{Object: ptr8728632544, Defaulter: false, Item: ptr8728442576, Optional: true}

	ptr8728655632 := &system.Type{Object: ptr8728631312, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"uniqueItems": ptr8728497920, "items": ptr8728496912, "minItems": ptr8728497488, "maxItems": ptr8728497824}, Rule: (*system.Type)(nil)}

	ptr8728655488 := &system.Type{Object: ptr8728631200, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728655632}

	ptr8728614688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744768 := &system.Object{Context: ptr8728614688, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728614240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745440 := &system.Object{Context: ptr8728614240, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728614048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747904 := &system.Object{Context: ptr8728614048, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728613984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728625600 := &system.Object{Context: ptr8728613984, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728561536 := &system.String_rule{Object: ptr8728625600, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728500464 := &system.Property{Object: ptr8728747904, Defaulter: false, Item: ptr8728561536, Optional: true}

	ptr8728346432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745552 := &system.Object{Context: ptr8728346432, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728346368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728746336 := &system.Object{Context: ptr8728346368, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728560576 := &system.String_rule{Object: ptr8728746336, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728498448 := &system.Property{Object: ptr8728745552, Defaulter: true, Item: ptr8728560576, Optional: true}

	ptr8728610144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728746448 := &system.Object{Context: ptr8728610144, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728609856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728746784 := &system.Object{Context: ptr8728609856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728609760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728746896 := &system.Object{Context: ptr8728609760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728561056 := &system.String_rule{Object: ptr8728746896, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728602496 := &system.Array_rule{Object: ptr8728746784, Items: ptr8728561056, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728498928 := &system.Property{Object: ptr8728746448, Defaulter: false, Item: ptr8728602496, Optional: true}

	ptr8728610816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747008 := &system.Object{Context: ptr8728610816, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728610784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747120 := &system.Object{Context: ptr8728610784, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439616 := &system.Int_rule{Object: ptr8728747120, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728499600 := &system.Property{Object: ptr8728747008, Defaulter: false, Item: ptr8728439616, Optional: true}

	ptr8728611584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747232 := &system.Object{Context: ptr8728611584, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728611456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747344 := &system.Object{Context: ptr8728611456, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728437856 := &system.Int_rule{Object: ptr8728747344, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728499984 := &system.Property{Object: ptr8728747232, Defaulter: false, Item: ptr8728437856, Optional: true}

	ptr8728612288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747456 := &system.Object{Context: ptr8728612288, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728612192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747568 := &system.Object{Context: ptr8728612192, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728561216 := &system.String_rule{Object: ptr8728747568, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728500080 := &system.Property{Object: ptr8728747456, Defaulter: false, Item: ptr8728561216, Optional: true}

	ptr8728612960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747680 := &system.Object{Context: ptr8728612960, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728612928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728747792 := &system.Object{Context: ptr8728612928, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728561376 := &system.String_rule{Object: ptr8728747792, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728500320 := &system.Property{Object: ptr8728747680, Defaulter: false, Item: ptr8728561376, Optional: true}

	ptr8728653184 := &system.Type{Object: ptr8728745440, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"format": ptr8728500464, "default": ptr8728498448, "enum": ptr8728498928, "minLength": ptr8728499600, "maxLength": ptr8728499984, "equal": ptr8728500080, "pattern": ptr8728500320}, Rule: (*system.Type)(nil)}

	ptr8728653040 := &system.Type{Object: ptr8728744768, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728653184}

	ptr8728616448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742864 := &system.Object{Context: ptr8728616448, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728615776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744096 := &system.Object{Context: ptr8728615776, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728615744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744432 := &system.Object{Context: ptr8728615744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728429512 := &system.Context_rule{Object: ptr8728744432}

	ptr8728499024 := &system.Property{Object: ptr8728744096, Defaulter: false, Item: ptr8728429512, Optional: true}

	ptr8728616320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744544 := &system.Object{Context: ptr8728616320, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728616288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744880 := &system.Object{Context: ptr8728616288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728616256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744992 := &system.Object{Context: ptr8728616256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728429696 := &system.Rule_rule{Object: ptr8728744992}

	ptr8728601152 := &system.Map_rule{Object: ptr8728744880, Items: ptr8728429696, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728499312 := &system.Property{Object: ptr8728744544, Defaulter: false, Item: ptr8728601152, Optional: true}

	ptr8728614432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742976 := &system.Object{Context: ptr8728614432, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728614400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743312 := &system.Object{Context: ptr8728614400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728601024 := &system.Reference_rule{Object: ptr8728743312, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728497536 := &system.Property{Object: ptr8728742976, Defaulter: false, Item: ptr8728601024, Optional: false}

	ptr8728614912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743424 := &system.Object{Context: ptr8728614912, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728614880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743760 := &system.Object{Context: ptr8728614880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728429408 := &system.JsonString_rule{Object: ptr8728743760}

	ptr8728498208 := &system.Property{Object: ptr8728743424, Defaulter: false, Item: ptr8728429408, Optional: true}

	ptr8728615360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743872 := &system.Object{Context: ptr8728615360, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728615328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743984 := &system.Object{Context: ptr8728615328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728429488 := &system.JsonString_rule{Object: ptr8728743984}

	ptr8728498976 := &system.Property{Object: ptr8728743872, Defaulter: false, Item: ptr8728429488, Optional: true}

	ptr8728650736 := &system.Type{Object: ptr8728742864, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"context": ptr8728499024, "rules": ptr8728499312, "type": ptr8728497536, "id": ptr8728498208, "description": ptr8728498976}, Rule: (*system.Type)(nil)}

	ptr8728342304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743536 := &system.Object{Context: ptr8728342304, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728341344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743648 := &system.Object{Context: ptr8728341344, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728341312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744208 := &system.Object{Context: ptr8728341312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728601728 := &system.Reference_rule{Object: ptr8728744208, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728496288 := &system.Property{Object: ptr8728743648, Defaulter: true, Item: ptr8728601728, Optional: true}

	ptr8728652320 := &system.Type{Object: ptr8728743536, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728496288}, Rule: (*system.Type)(nil)}

	ptr8728616640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745104 := &system.Object{Context: ptr8728616640, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650880 := &system.Type{Object: ptr8728745104, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728610112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632320 := &system.Object{Context: ptr8728610112, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728609024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632880 := &system.Object{Context: ptr8728609024, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728608992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632992 := &system.Object{Context: ptr8728608992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439056 := &system.Int_rule{Object: ptr8728632992, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728496960 := &system.Property{Object: ptr8728632880, Defaulter: true, Item: ptr8728439056, Optional: true}

	ptr8728609344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728633104 := &system.Object{Context: ptr8728609344, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728609312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728633216 := &system.Object{Context: ptr8728609312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439136 := &system.Int_rule{Object: ptr8728633216, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497104 := &system.Property{Object: ptr8728633104, Defaulter: false, Item: ptr8728439136, Optional: true}

	ptr8728609664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728281200 := &system.Object{Context: ptr8728609664, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728609632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739840 := &system.Object{Context: ptr8728609632, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439216 := &system.Int_rule{Object: ptr8728739840, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497200 := &system.Property{Object: ptr8728281200, Defaulter: false, Item: ptr8728439216, Optional: true}

	ptr8728609984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739952 := &system.Object{Context: ptr8728609984, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728609952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740064 := &system.Object{Context: ptr8728609952, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439296 := &system.Int_rule{Object: ptr8728740064, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497344 := &system.Property{Object: ptr8728739952, Defaulter: false, Item: ptr8728439296, Optional: true}

	ptr8728651168 := &system.Type{Object: ptr8728632320, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728496960, "multipleOf": ptr8728497104, "minimum": ptr8728497200, "maximum": ptr8728497344}, Rule: (*system.Type)(nil)}

	ptr8728345376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728630864 := &system.Object{Context: ptr8728345376, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728345056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631088 := &system.Object{Context: ptr8728345056, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728344928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631424 := &system.Object{Context: ptr8728344928, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631536 := &system.Object{Context: ptr8728344896, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728260720 := &system.Bool_rule{Object: ptr8728631536, Default: system.Bool{Value: false, Exists: false}}

	ptr8728496192 := &system.Property{Object: ptr8728631424, Defaulter: true, Item: ptr8728260720, Optional: true}

	ptr8728650016 := &system.Type{Object: ptr8728631088, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728496192}, Rule: (*system.Type)(nil)}

	ptr8728649872 := &system.Type{Object: ptr8728630864, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728650016}

	ptr8728345024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744656 := &system.Object{Context: ptr8728345024, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652752 := &system.Type{Object: ptr8728744656, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728344128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728744320 := &system.Object{Context: ptr8728344128, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652608 := &system.Type{Object: ptr8728744320, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728341152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728629520 := &system.Object{Context: ptr8728341152, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652896 := &system.Type{Object: ptr8728629520, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728345920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728631760 := &system.Object{Context: ptr8728345920, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650304 := &system.Type{Object: ptr8728631760, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728612992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741184 := &system.Object{Context: ptr8728612992, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728612256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742416 := &system.Object{Context: ptr8728612256, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728611968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742528 := &system.Object{Context: ptr8728611968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728437216 := &system.Number_rule{Object: ptr8728742528, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728496336 := &system.Property{Object: ptr8728742416, Defaulter: false, Item: ptr8728437216, Optional: true}

	ptr8728612736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742640 := &system.Object{Context: ptr8728612736, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728612672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742752 := &system.Object{Context: ptr8728612672, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728855904 := &system.Bool_rule{Object: ptr8728742752, Default: system.Bool{Value: false, Exists: true}}

	ptr8728496672 := &system.Property{Object: ptr8728742640, Defaulter: false, Item: ptr8728855904, Optional: true}

	ptr8728609536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741296 := &system.Object{Context: ptr8728609536, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728609248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741632 := &system.Object{Context: ptr8728609248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439696 := &system.Number_rule{Object: ptr8728741632, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728498256 := &system.Property{Object: ptr8728741296, Defaulter: true, Item: ptr8728439696, Optional: true}

	ptr8728610240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741744 := &system.Object{Context: ptr8728610240, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728610048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741856 := &system.Object{Context: ptr8728610048, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728437056 := &system.Number_rule{Object: ptr8728741856, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728495808 := &system.Property{Object: ptr8728741744, Defaulter: false, Item: ptr8728437056, Optional: true}

	ptr8728610944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741968 := &system.Object{Context: ptr8728610944, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728610752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742080 := &system.Object{Context: ptr8728610752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728437136 := &system.Number_rule{Object: ptr8728742080, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728495904 := &system.Property{Object: ptr8728741968, Defaulter: false, Item: ptr8728437136, Optional: true}

	ptr8728611520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742192 := &system.Object{Context: ptr8728611520, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728611488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728742304 := &system.Object{Context: ptr8728611488, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728855488 := &system.Bool_rule{Object: ptr8728742304, Default: system.Bool{Value: false, Exists: true}}

	ptr8728496048 := &system.Property{Object: ptr8728742192, Defaulter: false, Item: ptr8728855488, Optional: true}

	ptr8728652032 := &system.Type{Object: ptr8728741184, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maximum": ptr8728496336, "exclusiveMaximum": ptr8728496672, "default": ptr8728498256, "multipleOf": ptr8728495808, "minimum": ptr8728495904, "exclusiveMinimum": ptr8728496048}, Rule: (*system.Type)(nil)}

	ptr8728610432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728632208 := &system.Object{Context: ptr8728610432, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728651024 := &system.Type{Object: ptr8728632208, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728651168}

	ptr8728343072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728743088 := &system.Object{Context: ptr8728343072, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652176 := &system.Type{Object: ptr8728743088, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728652320}

	ptr8728613472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741072 := &system.Object{Context: ptr8728613472, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728651888 := &system.Type{Object: ptr8728741072, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728652032}

	ptr8728612544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740176 := &system.Object{Context: ptr8728612544, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728612224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740288 := &system.Object{Context: ptr8728612224, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728611392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740400 := &system.Object{Context: ptr8728611392, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728611360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740512 := &system.Object{Context: ptr8728611360, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728429432 := &system.Rule_rule{Object: ptr8728740512}

	ptr8728497728 := &system.Property{Object: ptr8728740400, Defaulter: false, Item: ptr8728429432, Optional: false}

	ptr8728611744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740624 := &system.Object{Context: ptr8728611744, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728611712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740736 := &system.Object{Context: ptr8728611712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439456 := &system.Int_rule{Object: ptr8728740736, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728498016 := &system.Property{Object: ptr8728740624, Defaulter: false, Item: ptr8728439456, Optional: true}

	ptr8728612096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740848 := &system.Object{Context: ptr8728612096, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728612064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728740960 := &system.Object{Context: ptr8728612064, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728439536 := &system.Int_rule{Object: ptr8728740960, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728497968 := &system.Property{Object: ptr8728740848, Defaulter: false, Item: ptr8728439536, Optional: true}

	ptr8728651600 := &system.Type{Object: ptr8728740288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8728497728, "minItems": ptr8728498016, "maxItems": ptr8728497968}, Rule: (*system.Type)(nil)}

	ptr8728651456 := &system.Type{Object: ptr8728740176, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728651600}

	ptr8728339840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741520 := &system.Object{Context: ptr8728339840, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650592 := &system.Type{Object: ptr8728741520, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728339328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745216 := &system.Object{Context: ptr8728339328, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728339776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745328 := &system.Object{Context: ptr8728339776, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728339744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745664 := &system.Object{Context: ptr8728339744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728429752 := &system.JsonBool_rule{Object: ptr8728745664}

	ptr8728499888 := &system.Property{Object: ptr8728745328, Defaulter: false, Item: ptr8728429752, Optional: true}

	ptr8728340768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745776 := &system.Object{Context: ptr8728340768, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728340736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728745888 := &system.Object{Context: ptr8728340736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728429824 := &system.JsonBool_rule{Object: ptr8728745888}

	ptr8728500224 := &system.Property{Object: ptr8728745776, Defaulter: false, Item: ptr8728429824, Optional: true}

	ptr8728339136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728746000 := &system.Object{Context: ptr8728339136, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728339104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728741408 := &system.Object{Context: ptr8728339104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728428584 := &system.Rule_rule{Object: ptr8728741408}

	ptr8728500272 := &system.Property{Object: ptr8728746000, Defaulter: false, Item: ptr8728428584, Optional: false}

	ptr8728655344 := &system.Type{Object: ptr8728745216, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8728499888, "defaulter": ptr8728500224, "item": ptr8728500272}, Rule: (*system.Type)(nil)}

	ptr8728339616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728625152 := &system.Object{Context: ptr8728339616, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728611008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728625264 := &system.Object{Context: ptr8728611008, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728610880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728625376 := &system.Object{Context: ptr8728610880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728600704 := &system.Reference_rule{Object: ptr8728625376, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8728494272 := &system.Property{Object: ptr8728625264, Defaulter: false, Item: ptr8728600704, Optional: true}

	ptr8728613120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728625488 := &system.Object{Context: ptr8728613120, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728612640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728627616 := &system.Object{Context: ptr8728612640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728612512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728627728 := &system.Object{Context: ptr8728612512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728601216 := &system.Reference_rule{Object: ptr8728627728, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728600896 := &system.Array_rule{Object: ptr8728627616, Items: ptr8728601216, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728494320 := &system.Property{Object: ptr8728625488, Defaulter: false, Item: ptr8728600896, Optional: true}

	ptr8728614528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728627840 := &system.Object{Context: ptr8728614528, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728614464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728627952 := &system.Object{Context: ptr8728614464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728559616 := &system.String_rule{Object: ptr8728627952, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728494464 := &system.Property{Object: ptr8728627840, Defaulter: false, Item: ptr8728559616, Optional: true}

	ptr8728615392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628064 := &system.Object{Context: ptr8728615392, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728615264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628176 := &system.Object{Context: ptr8728615264, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728429208 := &system.JsonBool_rule{Object: ptr8728628176}

	ptr8728494512 := &system.Property{Object: ptr8728628064, Defaulter: false, Item: ptr8728429208, Optional: true}

	ptr8728615968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628288 := &system.Object{Context: ptr8728615968, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728615904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628400 := &system.Object{Context: ptr8728615904, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728429232 := &system.JsonBool_rule{Object: ptr8728628400}

	ptr8728494560 := &system.Property{Object: ptr8728628288, Defaulter: false, Item: ptr8728429232, Optional: true}

	ptr8728338528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628512 := &system.Object{Context: ptr8728338528, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728338496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628624 := &system.Object{Context: ptr8728338496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728338464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728628960 := &system.Object{Context: ptr8728338464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8728429264 := &system.Property_rule{Object: ptr8728628960}

	ptr8728601536 := &system.Map_rule{Object: ptr8728628624, Items: ptr8728429264, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728494656 := &system.Property{Object: ptr8728628512, Defaulter: false, Item: ptr8728601536, Optional: true}

	ptr8728339424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728629072 := &system.Object{Context: ptr8728339424, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728339392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728629408 := &system.Object{Context: ptr8728339392, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8728429760 := &system.Type_rule{Object: ptr8728629408}

	ptr8728497872 := &system.Property{Object: ptr8728629072, Defaulter: false, Item: ptr8728429760, Optional: true}

	ptr8728651312 := &system.Type{Object: ptr8728625152, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"extends": ptr8728494272, "is": ptr8728494320, "native": ptr8728494464, "interface": ptr8728494512, "exclude": ptr8728494560, "properties": ptr8728494656, "rule": ptr8728497872}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8728655632)

	system.RegisterType("kego.io/system:@bool", ptr8728650016)

	system.RegisterType("kego.io/system:@context", ptr8728650448)

	system.RegisterType("kego.io/system:@int", ptr8728651168)

	system.RegisterType("kego.io/system:@map", ptr8728651600)

	system.RegisterType("kego.io/system:@number", ptr8728652032)

	system.RegisterType("kego.io/system:@object", ptr8728650880)

	system.RegisterType("kego.io/system:@property", ptr8728650592)

	system.RegisterType("kego.io/system:@reference", ptr8728652320)

	system.RegisterType("kego.io/system:@rule", ptr8728652752)

	system.RegisterType("kego.io/system:@string", ptr8728653184)

	system.RegisterType("kego.io/system:@type", ptr8728652896)

	system.RegisterType("kego.io/system:array", ptr8728655488)

	system.RegisterType("kego.io/system:bool", ptr8728649872)

	system.RegisterType("kego.io/system:context", ptr8728650304)

	system.RegisterType("kego.io/system:int", ptr8728651024)

	system.RegisterType("kego.io/system:map", ptr8728651456)

	system.RegisterType("kego.io/system:number", ptr8728651888)

	system.RegisterType("kego.io/system:object", ptr8728650736)

	system.RegisterType("kego.io/system:property", ptr8728655344)

	system.RegisterType("kego.io/system:reference", ptr8728652176)

	system.RegisterType("kego.io/system:rule", ptr8728652608)

	system.RegisterType("kego.io/system:string", ptr8728653040)

	system.RegisterType("kego.io/system:type", ptr8728651312)

}
