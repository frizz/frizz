package types

import (
	"kego.io/system"
)

func init() {

	ptr8728804704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050240 := &system.Base{Context: ptr8728804704, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728804576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050496 := &system.Base{Context: ptr8728804576, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728804544 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729321456 := &system.JsonString_rule{Base: ptr8729050496, RuleBase: ptr8728804544}

	ptr8728804032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050368 := &system.Base{Context: ptr8728804032, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728804000 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729321216 := &system.JsonBool_rule{Base: ptr8729050368, RuleBase: ptr8728804000}

	ptr8729032016 := &system.Type{Base: ptr8729050240, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"selector": ptr8729321456, "optional": ptr8729321216}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729330912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046400 := &system.Base{Context: ptr8729330912, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729028848 := &system.Type{Base: ptr8729046400, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044736 := &system.Base{Context: ptr8728803616, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728803200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044864 := &system.Base{Context: ptr8728803200, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728802496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045248 := &system.Base{Context: ptr8728802496, Description: "This is the maximum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728802464 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175216 := &system.Int_rule{Base: ptr8729045248, RuleBase: ptr8728802464, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728803040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045376 := &system.Base{Context: ptr8728803040, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728803008 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728802880 := &system.Bool_rule{Base: ptr8729045376, RuleBase: ptr8728803008, Default: system.Bool{Value: false, Exists: true}}

	ptr8728801440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044992 := &system.Base{Context: ptr8728801440, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728720640 := &system.Rule_rule{Base: ptr8729044992, RuleBase: (*system.RuleBase)(nil)}

	ptr8728801952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045120 := &system.Base{Context: ptr8728801952, Description: "This is the minimum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728801920 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175136 := &system.Int_rule{Base: ptr8729045120, RuleBase: ptr8728801920, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729026912 := &system.Type{Base: ptr8729044864, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"maxItems": ptr8729175216, "uniqueItems": ptr8728802880, "items": ptr8728720640, "minItems": ptr8729175136}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729026736 := &system.Type{Base: ptr8729044736, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729026912}

	ptr8729333824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047424 := &system.Base{Context: ptr8729333824, Description: "Lists imports used in this package.", Id: "imports", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729333696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047552 := &system.Base{Context: ptr8729333696, Description: "Map of import name to path.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729333664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048576 := &system.Base{Context: ptr8729333664, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729029904 := &system.String_rule{Base: ptr8729048576, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728884992 := &system.Map_rule{Base: ptr8729047552, RuleBase: (*system.RuleBase)(nil), Items: ptr8729029904, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729029728 := &system.Type{Base: ptr8729047424, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728884992}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728802048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049600 := &system.Base{Context: ptr8728802048, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728801632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049728 := &system.Base{Context: ptr8728801632, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728801504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049856 := &system.Base{Context: ptr8728801504, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728801472 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175936 := &system.Reference_rule{Base: ptr8729049856, RuleBase: ptr8728801472, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729031488 := &system.Type{Base: ptr8729049728, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729175936}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729031312 := &system.Type{Base: ptr8729049600, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729031488}

	ptr8728803168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050112 := &system.Base{Context: ptr8728803168, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729031840 := &system.Type{Base: ptr8729050112, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729333024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047040 := &system.Base{Context: ptr8729333024, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029552 := &system.Type{Base: ptr8729047040, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729307296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050752 := &system.Base{Context: ptr8729307296, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729306848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050880 := &system.Base{Context: ptr8729306848, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729306720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354496 := &system.Base{Context: ptr8729306720, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729306688 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176336 := &system.Int_rule{Base: ptr8729354496, RuleBase: ptr8729306688, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729307040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354624 := &system.Base{Context: ptr8729307040, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729307008 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176416 := &system.Int_rule{Base: ptr8729354624, RuleBase: ptr8729307008, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729305344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354752 := &system.Base{Context: ptr8729305344, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305280 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033072 := &system.String_rule{Base: ptr8729354752, RuleBase: ptr8729305280, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729305792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354880 := &system.Base{Context: ptr8729305792, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305760 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033248 := &system.String_rule{Base: ptr8729354880, RuleBase: ptr8729305760, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729306624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355008 := &system.Base{Context: ptr8729306624, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729306592 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033424 := &system.String_rule{Base: ptr8729355008, RuleBase: ptr8729306592, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729305888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051008 := &system.Base{Context: ptr8729305888, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305856 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729032720 := &system.String_rule{Base: ptr8729051008, RuleBase: ptr8729305856, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729306336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354240 := &system.Base{Context: ptr8729306336, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729306304 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729306272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729354368 := &system.Base{Context: ptr8729306272, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729032896 := &system.String_rule{Base: ptr8729354368, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729176176 := &system.Array_rule{Base: ptr8729354240, RuleBase: ptr8729306304, Items: ptr8729032896, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729032544 := &system.Type{Base: ptr8729050880, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"minLength": ptr8729176336, "maxLength": ptr8729176416, "equal": ptr8729033072, "pattern": ptr8729033248, "format": ptr8729033424, "default": ptr8729032720, "enum": ptr8729176176}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729032368 := &system.Type{Base: ptr8729050752, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729032544}

	ptr8729334240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048704 := &system.Base{Context: ptr8729334240, Description: "Automatically created basic rule for imports", Id: "@imports", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729030080 := &system.Type{Base: ptr8729048704, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729312448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356800 := &system.Base{Context: ptr8729312448, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729033952 := &system.Type{Base: ptr8729356800, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729335904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047296 := &system.Base{Context: ptr8729335904, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729335584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047680 := &system.Base{Context: ptr8729335584, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729334752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047808 := &system.Base{Context: ptr8729334752, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729316080 := &system.Rule_rule{Base: ptr8729047808, RuleBase: (*system.RuleBase)(nil)}

	ptr8729335104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047936 := &system.Base{Context: ptr8729335104, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729335072 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175776 := &system.Int_rule{Base: ptr8729047936, RuleBase: ptr8729335072, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729335456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048064 := &system.Base{Context: ptr8729335456, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729335424 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175856 := &system.Int_rule{Base: ptr8729048064, RuleBase: ptr8729335424, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729030432 := &system.Type{Base: ptr8729047680, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729316080, "minItems": ptr8729175776, "maxItems": ptr8729175856}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729030784 := &system.Type{Base: ptr8729047296, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729030432}

	ptr8729330496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045504 := &system.Base{Context: ptr8729330496, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728804800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045760 := &system.Base{Context: ptr8728804800, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728804768 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728723376 := &system.JsonString_rule{Base: ptr8729045760, RuleBase: ptr8728804768}

	ptr8728805216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045888 := &system.Base{Context: ptr8728805216, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728805184 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729313424 := &system.JsonString_rule{Base: ptr8729045888, RuleBase: ptr8728805184}

	ptr8729329824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046016 := &system.Base{Context: ptr8729329824, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8729329792 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729313664 := &system.Context_rule{Base: ptr8729046016, RuleBase: ptr8729329792}

	ptr8729330304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046144 := &system.Base{Context: ptr8729330304, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729330272 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729330240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046272 := &system.Base{Context: ptr8729330240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729313984 := &system.Rule_rule{Base: ptr8729046272, RuleBase: (*system.RuleBase)(nil)}

	ptr8729175376 := &system.Array_rule{Base: ptr8729046144, RuleBase: ptr8729330272, Items: ptr8729313984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728804416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045632 := &system.Base{Context: ptr8728804416, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729175296 := &system.Reference_rule{Base: ptr8729045632, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729028672 := &system.Type{Base: ptr8729045504, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"id": ptr8728723376, "description": ptr8729313424, "context": ptr8729313664, "rules": ptr8729175376, "type": ptr8729175296}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728805344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050624 := &system.Base{Context: ptr8728805344, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729032192 := &system.Type{Base: ptr8729050624, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729333184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048832 := &system.Base{Context: ptr8729333184, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729332480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729042944 := &system.Base{Context: ptr8729332480, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729331808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044608 := &system.Base{Context: ptr8729331808, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729331680 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175616 := &system.Int_rule{Base: ptr8729044608, RuleBase: ptr8729331680, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729332320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047168 := &system.Base{Context: ptr8729332320, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729332256 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175696 := &system.Int_rule{Base: ptr8729047168, RuleBase: ptr8729332256, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729330624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044224 := &system.Base{Context: ptr8729330624, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729330592 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175456 := &system.Int_rule{Base: ptr8729044224, RuleBase: ptr8729330592, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729331264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044480 := &system.Base{Context: ptr8729331264, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729331200 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175536 := &system.Int_rule{Base: ptr8729044480, RuleBase: ptr8729331200, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729030608 := &system.Type{Base: ptr8729042944, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"minimum": ptr8729175616, "maximum": ptr8729175696, "default": ptr8729175456, "multipleOf": ptr8729175536}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729030256 := &system.Type{Base: ptr8729048832, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729030608}

	ptr8729331936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046656 := &system.Base{Context: ptr8729331936, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729331776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046784 := &system.Base{Context: ptr8729331776, Description: "Default value if this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729331744 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729331648 := &system.Bool_rule{Base: ptr8729046784, RuleBase: ptr8729331744, Default: system.Bool{Value: false, Exists: false}}

	ptr8729029200 := &system.Type{Base: ptr8729046656, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729331648}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729312032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355136 := &system.Base{Context: ptr8729312032, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729308608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355392 := &system.Base{Context: ptr8729308608, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729308576 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729308544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355520 := &system.Base{Context: ptr8729308544, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729176736 := &system.Reference_rule{Base: ptr8729355520, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729176656 := &system.Array_rule{Base: ptr8729355392, RuleBase: ptr8729308576, Items: ptr8729176736, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729310304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356032 := &system.Base{Context: ptr8729310304, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729310272 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033776 := &system.String_rule{Base: ptr8729356032, RuleBase: ptr8729310272, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729310720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356160 := &system.Base{Context: ptr8729310720, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729310688 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728722176 := &system.JsonBool_rule{Base: ptr8729356160, RuleBase: ptr8729310688}

	ptr8729311104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356288 := &system.Base{Context: ptr8729311104, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729311072 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728722656 := &system.JsonBool_rule{Base: ptr8729356288, RuleBase: ptr8729311072}

	ptr8729311584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356416 := &system.Base{Context: ptr8729311584, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729311552 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729311520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356544 := &system.Base{Context: ptr8729311520, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728723344 := &system.Rule_rule{Base: ptr8729356544, RuleBase: (*system.RuleBase)(nil)}

	ptr8728882624 := &system.Map_rule{Base: ptr8729356416, RuleBase: ptr8729311552, Items: ptr8728723344, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729311904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729356672 := &system.Base{Context: ptr8729311904, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729311872 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729313472 := &system.Type_rule{Base: ptr8729356672, RuleBase: ptr8729311872}

	ptr8729308128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355264 := &system.Base{Context: ptr8729308128, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729308096 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728719728 := &system.JsonBool_rule{Base: ptr8729355264, RuleBase: ptr8729308096}

	ptr8729309152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355648 := &system.Base{Context: ptr8729309152, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729309120 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176816 := &system.Reference_rule{Base: ptr8729355648, RuleBase: ptr8729309120, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729309632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355776 := &system.Base{Context: ptr8729309632, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729309600 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729309568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729355904 := &system.Base{Context: ptr8729309568, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729176976 := &system.Reference_rule{Base: ptr8729355904, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729176896 := &system.Array_rule{Base: ptr8729355776, RuleBase: ptr8729309600, Items: ptr8729176976, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729033600 := &system.Type{Base: ptr8729355136, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"embed": ptr8729176656, "native": ptr8729033776, "interface": ptr8728722176, "exclude": ptr8728722656, "fields": ptr8728882624, "rule": ptr8729313472, "basic": ptr8728719728, "extends": ptr8729176816, "is": ptr8729176896}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729332288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046528 := &system.Base{Context: ptr8729332288, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029024 := &system.Type{Base: ptr8729046528, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729029200}

	ptr8728802624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049984 := &system.Base{Context: ptr8728802624, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729031664 := &system.Type{Base: ptr8729049984, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729332832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046912 := &system.Base{Context: ptr8729332832, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029376 := &system.Type{Base: ptr8729046912, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728799904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048320 := &system.Base{Context: ptr8728799904, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728798784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049216 := &system.Base{Context: ptr8728798784, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728798752 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728798496 := &system.Bool_rule{Base: ptr8729049216, RuleBase: ptr8728798752, Default: system.Bool{Value: false, Exists: true}}

	ptr8728799264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049344 := &system.Base{Context: ptr8728799264, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728799232 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904064 := &system.Number_rule{Base: ptr8729049344, RuleBase: ptr8728799232, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728799776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049472 := &system.Base{Context: ptr8728799776, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728799744 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728799616 := &system.Bool_rule{Base: ptr8729049472, RuleBase: ptr8728799744, Default: system.Bool{Value: false, Exists: true}}

	ptr8729336960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048448 := &system.Base{Context: ptr8729336960, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729336928 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904352 := &system.Number_rule{Base: ptr8729048448, RuleBase: ptr8729336928, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729337408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048960 := &system.Base{Context: ptr8729337408, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729337376 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728903872 := &system.Number_rule{Base: ptr8729048960, RuleBase: ptr8729337376, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728797280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049088 := &system.Base{Context: ptr8728797280, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729337824 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728903968 := &system.Number_rule{Base: ptr8729049088, RuleBase: ptr8729337824, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729031136 := &system.Type{Base: ptr8729048320, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMinimum": ptr8728798496, "maximum": ptr8728904064, "exclusiveMaximum": ptr8728799616, "default": ptr8728904352, "multipleOf": ptr8728903872, "minimum": ptr8728903968}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728800224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048192 := &system.Base{Context: ptr8728800224, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729030960 := &system.Type{Base: ptr8729048192, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729031136}

	system.RegisterType("kego.io/system:@array", ptr8729026912)

	system.RegisterType("kego.io/system:@base", ptr8729028848)

	system.RegisterType("kego.io/system:@bool", ptr8729029200)

	system.RegisterType("kego.io/system:@context", ptr8729029552)

	system.RegisterType("kego.io/system:@imports", ptr8729030080)

	system.RegisterType("kego.io/system:@int", ptr8729030608)

	system.RegisterType("kego.io/system:@map", ptr8729030432)

	system.RegisterType("kego.io/system:@number", ptr8729031136)

	system.RegisterType("kego.io/system:@reference", ptr8729031488)

	system.RegisterType("kego.io/system:@rule", ptr8729031840)

	system.RegisterType("kego.io/system:@ruleBase", ptr8729032192)

	system.RegisterType("kego.io/system:@string", ptr8729032544)

	system.RegisterType("kego.io/system:@type", ptr8729033952)

	system.RegisterType("kego.io/system:array", ptr8729026736)

	system.RegisterType("kego.io/system:base", ptr8729028672)

	system.RegisterType("kego.io/system:bool", ptr8729029024)

	system.RegisterType("kego.io/system:context", ptr8729029376)

	system.RegisterType("kego.io/system:imports", ptr8729029728)

	system.RegisterType("kego.io/system:int", ptr8729030256)

	system.RegisterType("kego.io/system:map", ptr8729030784)

	system.RegisterType("kego.io/system:number", ptr8729030960)

	system.RegisterType("kego.io/system:reference", ptr8729031312)

	system.RegisterType("kego.io/system:rule", ptr8729031664)

	system.RegisterType("kego.io/system:ruleBase", ptr8729032016)

	system.RegisterType("kego.io/system:string", ptr8729032368)

	system.RegisterType("kego.io/system:type", ptr8729033600)

}
