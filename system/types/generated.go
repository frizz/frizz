package types

import (
	"kego.io/system"
)

func init() {

	ptr8729412576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022720 := &system.Base{Context: ptr8729412576, Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729411680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022848 := &system.Base{Context: ptr8729411680, Description: "Default value if this property is omitted", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729411648 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454304 := &system.Int_rule{Base: ptr8729022848, RuleBase: ptr8729411648, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729411936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022976 := &system.Base{Context: ptr8729411936, Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729411904 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454384 := &system.Int_rule{Base: ptr8729022976, RuleBase: ptr8729411904, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729412192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023104 := &system.Base{Context: ptr8729412192, Description: "This provides a lower bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729412160 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454464 := &system.Int_rule{Base: ptr8729023104, RuleBase: ptr8729412160, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729412448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023232 := &system.Base{Context: ptr8729412448, Description: "This provides an upper bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729412416 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454544 := &system.Int_rule{Base: ptr8729023232, RuleBase: ptr8729412416, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729331424 := &system.Type{Base: ptr8729022720, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729454304, "multipleOf": ptr8729454384, "minimum": ptr8729454464, "maximum": ptr8729454544}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021952 := &system.Base{Context: ptr8728803936, Description: "Automatically created basic rule for context", Id: system.Reference{Package: "kego.io/system", Name: "@context", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729333344 := &system.Type{Base: ptr8729021952, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729417472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025664 := &system.Base{Context: ptr8729417472, Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729417024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025792 := &system.Base{Context: ptr8729417024, Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8729416992 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729397936 := &system.JsonBool_rule{Base: ptr8729025792, RuleBase: ptr8729416992}

	ptr8729417344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025920 := &system.Base{Context: ptr8729417344, Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8729417312 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729398128 := &system.JsonString_rule{Base: ptr8729025920, RuleBase: ptr8729417312}

	ptr8729334784 := &system.Type{Base: ptr8729025664, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729397936, "selector": ptr8729398128}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728798784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729026304 := &system.Base{Context: ptr8728798784, Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729418336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729026432 := &system.Base{Context: ptr8729418336, Description: "Default value of this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729418304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729059680 := &system.String_rule{Base: ptr8729026432, RuleBase: ptr8729418304, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729418656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354240 := &system.Base{Context: ptr8729418656, Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729418624 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729418592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354368 := &system.Base{Context: ptr8729418592, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729059856 := &system.String_rule{Base: ptr8729354368, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729454944 := &system.Array_rule{Base: ptr8729354240, RuleBase: ptr8729418624, Items: ptr8729059856, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729418976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354496 := &system.Base{Context: ptr8729418976, Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729418944 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729455104 := &system.Int_rule{Base: ptr8729354496, RuleBase: ptr8729418944, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729419232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354624 := &system.Base{Context: ptr8729419232, Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729419200 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729455184 := &system.Int_rule{Base: ptr8729354624, RuleBase: ptr8729419200, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729419488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354752 := &system.Base{Context: ptr8729419488, Description: "This is a string that the value must match", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729419456 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729060032 := &system.String_rule{Base: ptr8729354752, RuleBase: ptr8729419456, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729419744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354880 := &system.Base{Context: ptr8729419744, Description: "This is a regex to match the value to", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729419712 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729060208 := &system.String_rule{Base: ptr8729354880, RuleBase: ptr8729419712, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728798528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355008 := &system.Base{Context: ptr8728798528, Description: "This restricts the value to one of several built-in formats", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728798496 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729060384 := &system.String_rule{Base: ptr8729355008, RuleBase: ptr8728798496, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729335264 := &system.Type{Base: ptr8729026304, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729059680, "enum": ptr8729454944, "minLength": ptr8729455104, "maxLength": ptr8729455184, "equal": ptr8729060032, "pattern": ptr8729060208, "format": ptr8729060384}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021824 := &system.Base{Context: ptr8728803840, Description: "Unmarshal context.", Id: system.Reference{Package: "kego.io/system", Name: "context", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729333184 := &system.Type{Base: ptr8729021824, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Package: "kego.io/system", Name: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729415776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025152 := &system.Base{Context: ptr8729415776, Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729415648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025280 := &system.Base{Context: ptr8729415648, Description: "Default value of this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729415616 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728883392 := &system.Reference_rule{Base: ptr8729025280, RuleBase: ptr8729415616, Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729333984 := &system.Type{Base: ptr8729025152, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728883392}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728804832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022464 := &system.Base{Context: ptr8728804832, Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729331104 := &system.Type{Base: ptr8729022464, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728799936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729019648 := &system.Base{Context: ptr8728799936, Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728799040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020032 := &system.Base{Context: ptr8728799040, Description: "This is the minimum number of items allowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728799008 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454064 := &system.Int_rule{Base: ptr8729020032, RuleBase: ptr8728799008, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728799488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020160 := &system.Base{Context: ptr8728799488, Description: "This is the maximum number of items allowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8728799456 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454144 := &system.Int_rule{Base: ptr8729020160, RuleBase: ptr8728799456, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728799808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020288 := &system.Base{Context: ptr8728799808, Description: "If this is true, each item must be unique", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728799776 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728799680 := &system.Bool_rule{Base: ptr8729020288, RuleBase: ptr8728799776, Default: system.Bool{Value: false, Exists: true}}

	ptr8728798592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729019904 := &system.Base{Context: ptr8728798592, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729395632 := &system.Rule_rule{Base: ptr8729019904, RuleBase: (*system.RuleBase)(nil)}

	ptr8729329824 := &system.Type{Base: ptr8729019648, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"minItems": ptr8729454064, "maxItems": ptr8729454144, "uniqueItems": ptr8728799680, "items": ptr8729395632}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728802144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020416 := &system.Base{Context: ptr8728802144, Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728800736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020544 := &system.Base{Context: ptr8728800736, Description: "Type of the object.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8729267456 := &system.Reference_rule{Base: ptr8729020544, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8728800992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020672 := &system.Base{Context: ptr8728800992, Description: "All global objects should have an id.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728800960 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729267584 := &system.Reference_rule{Base: ptr8729020672, RuleBase: ptr8728800960, Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8728801344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020800 := &system.Base{Context: ptr8728801344, Description: "Description for the developer", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}

	ptr8728801312 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729396944 := &system.JsonString_rule{Base: ptr8729020800, RuleBase: ptr8728801312}

	ptr8728801600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729020928 := &system.Base{Context: ptr8728801600, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@context", Exists: true}}

	ptr8728801568 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729397120 := &system.Context_rule{Base: ptr8729020928, RuleBase: ptr8728801568}

	ptr8728801952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021056 := &system.Base{Context: ptr8728801952, Description: "Extra validation rules for this object or descendants", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728801920 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728801888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021184 := &system.Base{Context: ptr8728801888, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729397312 := &system.Rule_rule{Base: ptr8729021184, RuleBase: (*system.RuleBase)(nil)}

	ptr8729454224 := &system.Array_rule{Base: ptr8729021056, RuleBase: ptr8728801920, Items: ptr8729397312, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729330304 := &system.Type{Base: ptr8729020416, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729267456, "id": ptr8729267584, "description": ptr8729396944, "context": ptr8729397120, "rules": ptr8729454224}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728802400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021312 := &system.Base{Context: ptr8728802400, Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729330464 := &system.Type{Base: ptr8729021312, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728804512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022080 := &system.Base{Context: ptr8728804512, Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728804384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022208 := &system.Base{Context: ptr8728804384, Description: "Map of import name to path.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728804352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022336 := &system.Base{Context: ptr8728804352, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729059504 := &system.String_rule{Base: ptr8729022336, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729271424 := &system.Map_rule{Base: ptr8729022208, RuleBase: (*system.RuleBase)(nil), Items: ptr8729059504, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729333504 := &system.Type{Base: ptr8729022080, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8729271424}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728799168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729026176 := &system.Base{Context: ptr8728799168, Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729335104 := &system.Type{Base: ptr8729026176, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729335264}

	ptr8728804480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356800 := &system.Base{Context: ptr8728804480, Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729335904 := &system.Type{Base: ptr8729356800, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729414432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024128 := &system.Base{Context: ptr8729414432, Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729414240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024896 := &system.Base{Context: ptr8729414240, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729414208 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729414048 := &system.Bool_rule{Base: ptr8729024896, RuleBase: ptr8729414208, Default: system.Bool{Value: false, Exists: true}}

	ptr8729411808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024256 := &system.Base{Context: ptr8729411808, Description: "Default value if this property is omitted", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729411776 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729026848 := &system.Number_rule{Base: ptr8729024256, RuleBase: ptr8729411776, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729412320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024384 := &system.Base{Context: ptr8729412320, Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729412288 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729026944 := &system.Number_rule{Base: ptr8729024384, RuleBase: ptr8729412288, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729412864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024512 := &system.Base{Context: ptr8729412864, Description: "This provides a lower bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729412800 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729027040 := &system.Number_rule{Base: ptr8729024512, RuleBase: ptr8729412800, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729413376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024640 := &system.Base{Context: ptr8729413376, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729413344 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729413248 := &system.Bool_rule{Base: ptr8729024640, RuleBase: ptr8729413344, Default: system.Bool{Value: false, Exists: true}}

	ptr8729413664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024768 := &system.Base{Context: ptr8729413664, Description: "This provides an upper bound for the restriction", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729413632 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729027136 := &system.Number_rule{Base: ptr8729024768, RuleBase: ptr8729413632, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729332704 := &system.Type{Base: ptr8729024128, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMaximum": ptr8729414048, "default": ptr8729026848, "multipleOf": ptr8729026944, "minimum": ptr8729027040, "exclusiveMinimum": ptr8729413248, "maximum": ptr8729027136}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728804128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355136 := &system.Base{Context: ptr8728804128, Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728802304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356032 := &system.Base{Context: ptr8728802304, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8728802240 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729060560 := &system.String_rule{Base: ptr8729356032, RuleBase: ptr8728802240, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728803168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356288 := &system.Base{Context: ptr8728803168, Description: "Exclude this type from the generated json?", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728803136 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729401808 := &system.JsonBool_rule{Base: ptr8729356288, RuleBase: ptr8728803136}

	ptr8728803648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356416 := &system.Base{Context: ptr8728803648, Description: "Each field is listed with it's type", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8728803616 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728803552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356544 := &system.Base{Context: ptr8728803552, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729402016 := &system.Rule_rule{Base: ptr8729356544, RuleBase: (*system.RuleBase)(nil)}

	ptr8728881664 := &system.Map_rule{Base: ptr8729356416, RuleBase: ptr8728803616, Items: ptr8729402016, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728800416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355392 := &system.Base{Context: ptr8728800416, Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728800384 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728800352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355520 := &system.Base{Context: ptr8728800352, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728881152 := &system.Reference_rule{Base: ptr8729355520, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729455424 := &system.Array_rule{Base: ptr8729355392, RuleBase: ptr8728800384, Items: ptr8728881152, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728801024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355648 := &system.Base{Context: ptr8728801024, Description: "Type which this should extend", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728800928 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728881280 := &system.Reference_rule{Base: ptr8729355648, RuleBase: ptr8728800928, Default: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}}

	ptr8728802752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356160 := &system.Base{Context: ptr8728802752, Description: "Is this type an interface?", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728802720 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729401632 := &system.JsonBool_rule{Base: ptr8729356160, RuleBase: ptr8728802720}

	ptr8728804000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356672 := &system.Base{Context: ptr8728804000, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}

	ptr8728803968 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729402160 := &system.Type_rule{Base: ptr8729356672, RuleBase: ptr8728803968}

	ptr8728799968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355264 := &system.Base{Context: ptr8728799968, Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}

	ptr8728799904 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729400496 := &system.JsonBool_rule{Base: ptr8729355264, RuleBase: ptr8728799904}

	ptr8728801472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355776 := &system.Base{Context: ptr8728801472, Description: "Array of interface types that this type should support", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8728801440 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728801408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355904 := &system.Base{Context: ptr8728801408, Description: "", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}

	ptr8728881536 := &system.Reference_rule{Base: ptr8729355904, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Package: "", Name: "", Exists: false}}

	ptr8729455504 := &system.Array_rule{Base: ptr8729355776, RuleBase: ptr8728801440, Items: ptr8728881536, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729335744 := &system.Type{Base: ptr8729355136, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"native": ptr8729060560, "exclude": ptr8729401808, "fields": ptr8728881664, "embed": ptr8729455424, "extends": ptr8728881280, "interface": ptr8729401632, "rule": ptr8729402160, "basic": ptr8729400496, "is": ptr8729455504}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729416608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025536 := &system.Base{Context: ptr8729416608, Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729334624 := &system.Type{Base: ptr8729025536, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021440 := &system.Base{Context: ptr8728803456, Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728803200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021568 := &system.Base{Context: ptr8728803200, Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8728803072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729021696 := &system.Base{Context: ptr8728803072, Description: "Default value if this is missing or null", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8728803040 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728802976 := &system.Bool_rule{Base: ptr8729021696, RuleBase: ptr8728803040, Default: system.Bool{Value: false, Exists: false}}

	ptr8729330784 := &system.Type{Base: ptr8729021568, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728802976}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729330624 := &system.Type{Base: ptr8729021440, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729330784}

	ptr8729414272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023488 := &system.Base{Context: ptr8729414272, Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729413568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023616 := &system.Base{Context: ptr8729413568, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}

	ptr8729400688 := &system.Rule_rule{Base: ptr8729023616, RuleBase: (*system.RuleBase)(nil)}

	ptr8729413856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023744 := &system.Base{Context: ptr8729413856, Description: "This is the minimum number of items alowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729413824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454624 := &system.Int_rule{Base: ptr8729023744, RuleBase: ptr8729413824, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729414144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023872 := &system.Base{Context: ptr8729414144, Description: "This is the maximum number of items alowed in the array", Id: system.Reference{Package: "", Name: "", Exists: false}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}

	ptr8729414112 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729454704 := &system.Int_rule{Base: ptr8729023872, RuleBase: ptr8729414112, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729332064 := &system.Type{Base: ptr8729023488, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729400688, "minItems": ptr8729454624, "maxItems": ptr8729454704}, Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729417728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729026048 := &system.Base{Context: ptr8729417728, Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729334944 := &system.Type{Base: ptr8729026048, Basic: false, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Package: "", Name: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729414528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729023360 := &system.Base{Context: ptr8729414528, Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729331904 := &system.Type{Base: ptr8729023360, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729332064}

	ptr8729416032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025024 := &system.Base{Context: ptr8729416032, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729333824 := &system.Type{Base: ptr8729025024, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729333984}

	ptr8729416352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729025408 := &system.Base{Context: ptr8729416352, Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729334464 := &system.Type{Base: ptr8729025408, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729412832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729022592 := &system.Base{Context: ptr8729412832, Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729331264 := &system.Type{Base: ptr8729022592, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729331424}

	ptr8729414816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729024000 := &system.Base{Context: ptr8729414816, Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729332544 := &system.Type{Base: ptr8729024000, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729332704}

	ptr8728800192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729018368 := &system.Base{Context: ptr8728800192, Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729329664 := &system.Type{Base: ptr8729018368, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729329824}

	system.RegisterType("kego.io/system", "@ruleBase", ptr8729334944)

	system.RegisterType("kego.io/system", "map", ptr8729331904)

	system.RegisterType("kego.io/system", "number", ptr8729332544)

	system.RegisterType("kego.io/system", "array", ptr8729329664)

	system.RegisterType("kego.io/system", "@string", ptr8729335264)

	system.RegisterType("kego.io/system", "@reference", ptr8729333984)

	system.RegisterType("kego.io/system", "@base", ptr8729330464)

	system.RegisterType("kego.io/system", "string", ptr8729335104)

	system.RegisterType("kego.io/system", "@rule", ptr8729334624)

	system.RegisterType("kego.io/system", "@bool", ptr8729330784)

	system.RegisterType("kego.io/system", "@map", ptr8729332064)

	system.RegisterType("kego.io/system", "int", ptr8729331264)

	system.RegisterType("kego.io/system", "imports", ptr8729333504)

	system.RegisterType("kego.io/system", "@type", ptr8729335904)

	system.RegisterType("kego.io/system", "@number", ptr8729332704)

	system.RegisterType("kego.io/system", "type", ptr8729335744)

	system.RegisterType("kego.io/system", "@imports", ptr8729331104)

	system.RegisterType("kego.io/system", "@array", ptr8729329824)

	system.RegisterType("kego.io/system", "base", ptr8729330304)

	system.RegisterType("kego.io/system", "bool", ptr8729330624)

	system.RegisterType("kego.io/system", "@int", ptr8729331424)

	system.RegisterType("kego.io/system", "@context", ptr8729333344)

	system.RegisterType("kego.io/system", "ruleBase", ptr8729334784)

	system.RegisterType("kego.io/system", "context", ptr8729333184)

	system.RegisterType("kego.io/system", "reference", ptr8729333824)

	system.RegisterType("kego.io/system", "rule", ptr8729334464)

}
