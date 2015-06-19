package types

import (
	"kego.io/system"
)

func init() {

	ptr8729332736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047424 := &system.Base{Context: ptr8729332736, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029552 := &system.Type{Base: ptr8729047424, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729330208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045376 := &system.Base{Context: ptr8729330208, Description: "This is the most basic type.", Id: "base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729330016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046016 := &system.Base{Context: ptr8729330016, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729329984 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729329952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046144 := &system.Base{Context: ptr8729329952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729313760 := &system.Rule_rule{Base: ptr8729046144, RuleBase: (*system.RuleBase)(nil)}

	ptr8729175376 := &system.Array_rule{Base: ptr8729046016, RuleBase: ptr8729329984, Items: ptr8729313760, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728804128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045504 := &system.Base{Context: ptr8728804128, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729175296 := &system.Reference_rule{Base: ptr8729045504, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728804512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045632 := &system.Base{Context: ptr8728804512, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728804480 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728722784 := &system.JsonString_rule{Base: ptr8729045632, RuleBase: ptr8728804480}

	ptr8728804928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045760 := &system.Base{Context: ptr8728804928, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728804896 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728723376 := &system.JsonString_rule{Base: ptr8729045760, RuleBase: ptr8728804896}

	ptr8728805248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045888 := &system.Base{Context: ptr8728805248, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728805216 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729313440 := &system.Context_rule{Base: ptr8729045888, RuleBase: ptr8728805216}

	ptr8729028672 := &system.Type{Base: ptr8729045376, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"rules": ptr8729175376, "type": ptr8729175296, "id": ptr8728722784, "description": ptr8728723376, "context": ptr8729313440}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728801888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049600 := &system.Base{Context: ptr8728801888, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728801568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049728 := &system.Base{Context: ptr8728801568, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728801344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049856 := &system.Base{Context: ptr8728801344, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728801312 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175936 := &system.Reference_rule{Base: ptr8729049856, RuleBase: ptr8728801312, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729031488 := &system.Type{Base: ptr8729049728, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729175936}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729031312 := &system.Type{Base: ptr8729049600, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729031488}

	ptr8728803104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050112 := &system.Base{Context: ptr8728803104, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729031840 := &system.Type{Base: ptr8729050112, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729332352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048960 := &system.Base{Context: ptr8729332352, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729330464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729042944 := &system.Base{Context: ptr8729330464, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729330400 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175456 := &system.Int_rule{Base: ptr8729042944, RuleBase: ptr8729330400, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729331072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044224 := &system.Base{Context: ptr8729331072, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729331040 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175536 := &system.Int_rule{Base: ptr8729044224, RuleBase: ptr8729331040, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729331584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044480 := &system.Base{Context: ptr8729331584, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729331552 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175616 := &system.Int_rule{Base: ptr8729044480, RuleBase: ptr8729331552, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729332096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046912 := &system.Base{Context: ptr8729332096, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729332064 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175696 := &system.Int_rule{Base: ptr8729046912, RuleBase: ptr8729332064, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729030432 := &system.Type{Base: ptr8729048960, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729175456, "multipleOf": ptr8729175536, "minimum": ptr8729175616, "maximum": ptr8729175696}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729332960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048832 := &system.Base{Context: ptr8729332960, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729030256 := &system.Type{Base: ptr8729048832, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729030432}

	ptr8729306784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050880 := &system.Base{Context: ptr8729306784, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729305280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362944 := &system.Base{Context: ptr8729305280, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305216 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033072 := &system.String_rule{Base: ptr8729362944, RuleBase: ptr8729305216, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729305728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363072 := &system.Base{Context: ptr8729305728, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305696 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033248 := &system.String_rule{Base: ptr8729363072, RuleBase: ptr8729305696, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729306560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363200 := &system.Base{Context: ptr8729306560, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729306528 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033424 := &system.String_rule{Base: ptr8729363200, RuleBase: ptr8729306528, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729305824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729051008 := &system.Base{Context: ptr8729305824, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729305792 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729032720 := &system.String_rule{Base: ptr8729051008, RuleBase: ptr8729305792, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729306272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362432 := &system.Base{Context: ptr8729306272, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729306240 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729306208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362560 := &system.Base{Context: ptr8729306208, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729032896 := &system.String_rule{Base: ptr8729362560, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729176176 := &system.Array_rule{Base: ptr8729362432, RuleBase: ptr8729306240, Items: ptr8729032896, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729306656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362688 := &system.Base{Context: ptr8729306656, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729306624 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176336 := &system.Int_rule{Base: ptr8729362688, RuleBase: ptr8729306624, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729306976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362816 := &system.Base{Context: ptr8729306976, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729306944 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176416 := &system.Int_rule{Base: ptr8729362816, RuleBase: ptr8729306944, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729032544 := &system.Type{Base: ptr8729050880, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"equal": ptr8729033072, "pattern": ptr8729033248, "format": ptr8729033424, "default": ptr8729032720, "enum": ptr8729176176, "minLength": ptr8729176336, "maxLength": ptr8729176416}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729331616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046528 := &system.Base{Context: ptr8729331616, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729331488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046656 := &system.Base{Context: ptr8729331488, Description: "Default value if this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729331456 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729331360 := &system.Bool_rule{Base: ptr8729046656, RuleBase: ptr8729331456, Default: system.Bool{Value: false, Exists: false}}

	ptr8729029200 := &system.Type{Base: ptr8729046528, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729331360}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729307232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050752 := &system.Base{Context: ptr8729307232, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729032368 := &system.Type{Base: ptr8729050752, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr8729032544}

	ptr8729335744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047040 := &system.Base{Context: ptr8729335744, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729335424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047168 := &system.Base{Context: ptr8729335424, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729334592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047296 := &system.Base{Context: ptr8729334592, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729315984 := &system.Rule_rule{Base: ptr8729047296, RuleBase: (*system.RuleBase)(nil)}

	ptr8729334944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047936 := &system.Base{Context: ptr8729334944, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729334912 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175776 := &system.Int_rule{Base: ptr8729047936, RuleBase: ptr8729334912, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729335296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048064 := &system.Base{Context: ptr8729335296, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729335264 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175856 := &system.Int_rule{Base: ptr8729048064, RuleBase: ptr8729335264, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729030784 := &system.Type{Base: ptr8729047168, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729315984, "minItems": ptr8729175776, "maxItems": ptr8729175856}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729030608 := &system.Type{Base: ptr8729047040, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr8729030784}

	ptr8729312384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364992 := &system.Base{Context: ptr8729312384, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729033952 := &system.Type{Base: ptr8729364992, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729330624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046272 := &system.Base{Context: ptr8729330624, Description: "Automatically created basic rule for base", Id: "@base", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729028848 := &system.Type{Base: ptr8729046272, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728800064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048192 := &system.Base{Context: ptr8728800064, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728799744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048320 := &system.Base{Context: ptr8728799744, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728799040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049344 := &system.Base{Context: ptr8728799040, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728799008 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904064 := &system.Number_rule{Base: ptr8729049344, RuleBase: ptr8728799008, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728799616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049472 := &system.Base{Context: ptr8728799616, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728799584 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728799456 := &system.Bool_rule{Base: ptr8729049472, RuleBase: ptr8728799584, Default: system.Bool{Value: false, Exists: true}}

	ptr8729336800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048448 := &system.Base{Context: ptr8729336800, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729336768 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728904352 := &system.Number_rule{Base: ptr8729048448, RuleBase: ptr8729336768, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729337248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048576 := &system.Base{Context: ptr8729337248, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729337216 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728903872 := &system.Number_rule{Base: ptr8729048576, RuleBase: ptr8729337216, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729337696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049088 := &system.Base{Context: ptr8729337696, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729337664 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728903968 := &system.Number_rule{Base: ptr8729049088, RuleBase: ptr8729337664, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728798496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049216 := &system.Base{Context: ptr8728798496, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728798240 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728797984 := &system.Bool_rule{Base: ptr8729049216, RuleBase: ptr8728798240, Default: system.Bool{Value: false, Exists: true}}

	ptr8729031136 := &system.Type{Base: ptr8729048320, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"maximum": ptr8728904064, "exclusiveMaximum": ptr8728799456, "default": ptr8728904352, "multipleOf": ptr8728903872, "minimum": ptr8728903968, "exclusiveMinimum": ptr8728797984}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729030960 := &system.Type{Base: ptr8729048192, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr8729031136}

	ptr8729331968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046400 := &system.Base{Context: ptr8729331968, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029024 := &system.Type{Base: ptr8729046400, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr8729029200}

	ptr8729332544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729046784 := &system.Base{Context: ptr8729332544, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729029376 := &system.Type{Base: ptr8729046784, Basic: true, Embed: []system.Reference(nil), Exclude: true, Extends: system.Reference{Value: "kego.io/system:", Package: "kego.io/system", Type: "", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729311968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363328 := &system.Base{Context: ptr8729311968, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729308064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363456 := &system.Base{Context: ptr8729308064, Description: "Basic types don't have system:object added by default to the embedded types.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729308032 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728719712 := &system.JsonBool_rule{Base: ptr8729363456, RuleBase: ptr8729308032}

	ptr8729309088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363840 := &system.Base{Context: ptr8729309088, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729309056 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729176816 := &system.Reference_rule{Base: ptr8729363840, RuleBase: ptr8729309056, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729311520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364608 := &system.Base{Context: ptr8729311520, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729311488 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729311456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364736 := &system.Base{Context: ptr8729311456, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728723344 := &system.Rule_rule{Base: ptr8729364736, RuleBase: (*system.RuleBase)(nil)}

	ptr8728882624 := &system.Map_rule{Base: ptr8729364608, RuleBase: ptr8729311488, Items: ptr8728723344, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729311040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364480 := &system.Base{Context: ptr8729311040, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729311008 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728722640 := &system.JsonBool_rule{Base: ptr8729364480, RuleBase: ptr8729311008}

	ptr8729311840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364864 := &system.Base{Context: ptr8729311840, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729311808 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729313488 := &system.Type_rule{Base: ptr8729364864, RuleBase: ptr8729311808}

	ptr8729308544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363584 := &system.Base{Context: ptr8729308544, Description: "Types which this should embed - system:object is always added unless base = true.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729308512 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729308480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363712 := &system.Base{Context: ptr8729308480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729176736 := &system.Reference_rule{Base: ptr8729363712, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729176656 := &system.Array_rule{Base: ptr8729363584, RuleBase: ptr8729308512, Items: ptr8729176736, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729309568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729363968 := &system.Base{Context: ptr8729309568, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729309536 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729309504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364096 := &system.Base{Context: ptr8729309504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729176976 := &system.Reference_rule{Base: ptr8729364096, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729176896 := &system.Array_rule{Base: ptr8729363968, RuleBase: ptr8729309536, Items: ptr8729176976, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729310240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364224 := &system.Base{Context: ptr8729310240, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729310208 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729033776 := &system.String_rule{Base: ptr8729364224, RuleBase: ptr8729310208, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729310656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729364352 := &system.Base{Context: ptr8729310656, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729310624 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728722160 := &system.JsonBool_rule{Base: ptr8729364352, RuleBase: ptr8729310624}

	ptr8729033600 := &system.Type{Base: ptr8729363328, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8728719712, "extends": ptr8729176816, "fields": ptr8728882624, "exclude": ptr8728722640, "rule": ptr8729313488, "embed": ptr8729176656, "is": ptr8729176896, "native": ptr8729033776, "interface": ptr8728722160}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729333952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729048704 := &system.Base{Context: ptr8729333952, Description: "Automatically created basic rule for imports", Id: "@imports", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729030080 := &system.Type{Base: ptr8729048704, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728803168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044608 := &system.Base{Context: ptr8728803168, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728802784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044736 := &system.Base{Context: ptr8728802784, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728801024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044864 := &system.Base{Context: ptr8728801024, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728720416 := &system.Rule_rule{Base: ptr8729044864, RuleBase: (*system.RuleBase)(nil)}

	ptr8728801536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729044992 := &system.Base{Context: ptr8728801536, Description: "This is the minimum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728801504 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175136 := &system.Int_rule{Base: ptr8729044992, RuleBase: ptr8728801504, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728802080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045120 := &system.Base{Context: ptr8728802080, Description: "This is the maximum number of items allowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728802048 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729175216 := &system.Int_rule{Base: ptr8729045120, RuleBase: ptr8728802048, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728802624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729045248 := &system.Base{Context: ptr8728802624, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728802592 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8728802464 := &system.Bool_rule{Base: ptr8729045248, RuleBase: ptr8728802592, Default: system.Bool{Value: false, Exists: true}}

	ptr8729026912 := &system.Type{Base: ptr8729044736, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728720416, "minItems": ptr8729175136, "maxItems": ptr8729175216, "uniqueItems": ptr8728802464}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729026736 := &system.Type{Base: ptr8729044608, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr8729026912}

	ptr8728804640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050240 := &system.Base{Context: ptr8728804640, Description: "All rules should embed this type.", Id: "ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728803936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050368 := &system.Base{Context: ptr8728803936, Description: "If this rule is a field, this specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728803904 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729321088 := &system.JsonBool_rule{Base: ptr8729050368, RuleBase: ptr8728803904}

	ptr8728804416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050496 := &system.Base{Context: ptr8728804416, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728804384 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729321344 := &system.JsonString_rule{Base: ptr8729050496, RuleBase: ptr8728804384}

	ptr8729032016 := &system.Type{Base: ptr8729050240, Basic: true, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729321088, "selector": ptr8729321344}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729333536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047552 := &system.Base{Context: ptr8729333536, Description: "Lists imports used in this package.", Id: "imports", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729333408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047680 := &system.Base{Context: ptr8729333408, Description: "Map of import name to path.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729333376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729047808 := &system.Base{Context: ptr8729333376, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729029904 := &system.String_rule{Base: ptr8729047808, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728884992 := &system.Map_rule{Base: ptr8729047680, RuleBase: (*system.RuleBase)(nil), Items: ptr8729029904, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729029728 := &system.Type{Base: ptr8729047552, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728884992}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728802496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729049984 := &system.Base{Context: ptr8728802496, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729031664 := &system.Type{Base: ptr8729049984, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728805280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729050624 := &system.Base{Context: ptr8728805280, Description: "Automatically created basic rule for ruleBase", Id: "@ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729032192 := &system.Type{Base: ptr8729050624, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729026912)

	system.RegisterType("kego.io/system:@base", ptr8729028848)

	system.RegisterType("kego.io/system:@bool", ptr8729029200)

	system.RegisterType("kego.io/system:@context", ptr8729029552)

	system.RegisterType("kego.io/system:@imports", ptr8729030080)

	system.RegisterType("kego.io/system:@int", ptr8729030432)

	system.RegisterType("kego.io/system:@map", ptr8729030784)

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

	system.RegisterType("kego.io/system:map", ptr8729030608)

	system.RegisterType("kego.io/system:number", ptr8729030960)

	system.RegisterType("kego.io/system:reference", ptr8729031312)

	system.RegisterType("kego.io/system:rule", ptr8729031664)

	system.RegisterType("kego.io/system:ruleBase", ptr8729032016)

	system.RegisterType("kego.io/system:string", ptr8729032368)

	system.RegisterType("kego.io/system:type", ptr8729033600)

}
