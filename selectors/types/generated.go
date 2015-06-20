package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8729044224 := &system.Base{Description: "Automatically created basic rule for kid", Id: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729139296 := &system.Type{Base: ptr8729044224, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729163520 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729163648 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461120 := &system.Number_rule{Base: ptr8729163648, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729163776 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461216 := &system.Number_rule{Base: ptr8729163776, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729163904 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729143888 := &system.String_rule{Base: ptr8729163904, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729164032 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144064 := &system.String_rule{Base: ptr8729164032, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729164160 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144240 := &system.String_rule{Base: ptr8729164160, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729164288 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729257728 := &system.Bool_rule{Base: ptr8729164288, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729164416 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729257984 := &system.Bool_rule{Base: ptr8729164416, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{}}

	ptr8729138176 := &system.Type{Base: ptr8729163520, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"int": ptr8729461120, "float": ptr8729461216, "string": ptr8729143888, "string2": ptr8729144064, "null": ptr8729144240, "true": ptr8729257728, "false": ptr8729257984}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729529856 := &system.Base{Description: "Automatically created basic rule for typed", Id: system.Reference{Package: "kego.io/selectors", Name: "@typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729140736 := &system.Type{Base: ptr8729529856, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729165312 := &system.Base{Description: "Automatically created basic rule for image", Id: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729138976 := &system.Type{Base: ptr8729165312, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729161728 := &system.Base{Description: "Automatically created basic rule for basic", Id: system.Reference{Package: "kego.io/selectors", Name: "@basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729136896 := &system.Type{Base: ptr8729161728, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729526272 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729526400 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729256384 := &system.RuleBase{Optional: true}

	ptr8729143360 := &system.String_rule{Base: ptr8729526400, RuleBase: ptr8729256384, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729526528 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144768 := &system.String_rule{Base: ptr8729526528, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729526656 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144944 := &system.String_rule{Base: ptr8729526656, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729139456 := &system.Type{Base: ptr8729526272, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8729143360, "server": ptr8729144768, "path": ptr8729144944}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729163008 := &system.Base{Description: "Automatically created basic rule for collision", Id: system.Reference{Package: "kego.io/selectors", Name: "@collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729137536 := &system.Type{Base: ptr8729163008, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729528448 := &system.Base{Description: "Automatically created basic rule for sibling", Id: system.Reference{Package: "kego.io/selectors", Name: "@sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729140416 := &system.Type{Base: ptr8729528448, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729161856 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729161984 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729460832 := &system.Number_rule{Base: ptr8729161984, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729162112 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729460928 := &system.Number_rule{Base: ptr8729162112, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729162240 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729162368 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461024 := &system.Number_rule{Base: ptr8729162368, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728906112 := &system.Map_rule{Base: ptr8729162240, RuleBase: (*system.RuleBase)(nil), Items: ptr8729461024, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729137056 := &system.Type{Base: ptr8729161856, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729460832, "b": ptr8729460928, "c": ptr8728906112}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729526912 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729527040 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729527168 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729348208 := &selectors.Kid_rule{Base: ptr8729527168, RuleBase: (*system.RuleBase)(nil)}

	ptr8729363232 := &system.Array_rule{Base: ptr8729527040, RuleBase: (*system.RuleBase)(nil), Items: ptr8729348208, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729139936 := &system.Type{Base: ptr8729526912, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729363232}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729162496 := &system.Base{Description: "Automatically created basic rule for c", Id: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729137216 := &system.Type{Base: ptr8729162496, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729165056 := &system.Base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "kego.io/selectors", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729138656 := &system.Type{Base: ptr8729165056, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729164672 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729164800 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729164928 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729351328 := &selectors.Image_rule{Base: ptr8729164928, RuleBase: (*system.RuleBase)(nil)}

	ptr8728906304 := &system.Map_rule{Base: ptr8729164800, RuleBase: (*system.RuleBase)(nil), Items: ptr8729351328, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729138496 := &system.Type{Base: ptr8729164672, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728906304}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729165184 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729138816 := &system.Type{Base: ptr8729165184, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729165440 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "kid", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729165568 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144416 := &system.String_rule{Base: ptr8729165568, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729165696 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729144592 := &system.String_rule{Base: ptr8729165696, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729042944 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}

	ptr8729260448 := &system.RuleBase{Optional: true}

	ptr8729260352 := &system.Bool_rule{Base: ptr8729042944, RuleBase: ptr8729260448, Default: system.Bool{}}

	ptr8729139136 := &system.Type{Base: ptr8729165440, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729144416, "level": ptr8729144592, "preferred": ptr8729260352}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729163136 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729163264 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729143712 := &system.String_rule{Base: ptr8729163264, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729137696 := &system.Type{Base: ptr8729163136, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729143712}, Is: []system.Reference{system.Reference{Package: "kego.io/selectors", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729163392 := &system.Base{Description: "Automatically created basic rule for diagram", Id: system.Reference{Package: "kego.io/selectors", Name: "@diagram", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729138016 := &system.Type{Base: ptr8729163392, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729527296 := &system.Base{Description: "Automatically created basic rule for polykids", Id: system.Reference{Package: "kego.io/selectors", Name: "@polykids", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729140096 := &system.Type{Base: ptr8729527296, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729528576 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "typed", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729529344 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@image", Exists: true}}

	ptr8729350720 := &selectors.Image_rule{Base: ptr8729529344, RuleBase: (*system.RuleBase)(nil)}

	ptr8729529472 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729529600 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729145648 := &system.String_rule{Base: ptr8729529600, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729364112 := &system.Array_rule{Base: ptr8729529472, RuleBase: (*system.RuleBase)(nil), Items: ptr8729145648, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729529728 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461696 := &system.Number_rule{Base: ptr8729529728, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729528704 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729528832 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729145296 := &system.String_rule{Base: ptr8729528832, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728905984 := &system.Map_rule{Base: ptr8729528704, RuleBase: (*system.RuleBase)(nil), Items: ptr8729145296, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729528960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729145472 := &system.String_rule{Base: ptr8729528960, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729529088 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729529216 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@kid", Exists: true}}

	ptr8729350608 := &selectors.Kid_rule{Base: ptr8729529216, RuleBase: (*system.RuleBase)(nil)}

	ptr8728906048 := &system.Map_rule{Base: ptr8729529088, RuleBase: (*system.RuleBase)(nil), Items: ptr8729350608, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729140576 := &system.Type{Base: ptr8729528576, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"avatar": ptr8729350720, "drinkPreference": ptr8729364112, "weight": ptr8729461696, "name": ptr8728905984, "favoriteColor": ptr8729145472, "kids": ptr8728906048}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729527424 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "sibling", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729527808 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "@c", Exists: true}}

	ptr8729348880 := &selectors.C_rule{Base: ptr8729527808, RuleBase: (*system.RuleBase)(nil)}

	ptr8729527936 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729528064 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461504 := &system.Number_rule{Base: ptr8729528064, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728905856 := &system.Map_rule{Base: ptr8729527936, RuleBase: (*system.RuleBase)(nil), Items: ptr8729461504, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729528192 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729528320 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461600 := &system.Number_rule{Base: ptr8729528320, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8728905920 := &system.Map_rule{Base: ptr8729528192, RuleBase: (*system.RuleBase)(nil), Items: ptr8729461600, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729527552 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461312 := &system.Number_rule{Base: ptr8729527552, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729527680 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729461408 := &system.Number_rule{Base: ptr8729527680, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729140256 := &system.Type{Base: ptr8729527424, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"c": ptr8729348880, "d": ptr8728905856, "e": ptr8728905920, "a": ptr8729461312, "b": ptr8729461408}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729160192 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "basic", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729160576 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729142656 := &system.String_rule{Base: ptr8729160576, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729160704 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729160832 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729160960 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729142832 := &system.String_rule{Base: ptr8729160960, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728905280 := &system.Map_rule{Base: ptr8729160832, RuleBase: (*system.RuleBase)(nil), Items: ptr8729142832, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729362912 := &system.Array_rule{Base: ptr8729160704, RuleBase: (*system.RuleBase)(nil), Items: ptr8728905280, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729161088 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729161216 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729143008 := &system.String_rule{Base: ptr8729161216, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729362992 := &system.Array_rule{Base: ptr8729161088, RuleBase: (*system.RuleBase)(nil), Items: ptr8729143008, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729161344 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}

	ptr8729161472 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729143184 := &system.String_rule{Base: ptr8729161472, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8729363072 := &system.Array_rule{Base: ptr8729161344, RuleBase: (*system.RuleBase)(nil), Items: ptr8729143184, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}, UniqueItems: system.Bool{Exists: true}}

	ptr8729161600 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}

	ptr8729460736 := &system.Number_rule{Base: ptr8729161600, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, ExclusiveMaximum: system.Bool{Exists: true}, ExclusiveMinimum: system.Bool{Exists: true}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}

	ptr8729160320 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729160448 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729142480 := &system.String_rule{Base: ptr8729160448, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728905216 := &system.Map_rule{Base: ptr8729160320, RuleBase: (*system.RuleBase)(nil), Items: ptr8729142480, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729136736 := &system.Type{Base: ptr8729160192, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"favoriteColor": ptr8729142656, "languagesSpoken": ptr8729362912, "seatingPreference": ptr8729362992, "drinkPreference": ptr8729363072, "weight": ptr8729460736, "name": ptr8728905216}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729526784 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/selectors", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729139776 := &system.Type{Base: ptr8729526784, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729162624 := &system.Base{Id: system.Reference{Package: "kego.io/selectors", Name: "collision", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729162752 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}

	ptr8729162880 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}

	ptr8729143536 := &system.String_rule{Base: ptr8729162880, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []system.String(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}

	ptr8728906176 := &system.Map_rule{Base: ptr8729162752, RuleBase: (*system.RuleBase)(nil), Items: ptr8729143536, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}

	ptr8729137376 := &system.Type{Base: ptr8729162624, Embed: []system.Reference(nil), Extends: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8728906176}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729164544 := &system.Base{Description: "Automatically created basic rule for expr", Id: system.Reference{Package: "kego.io/selectors", Name: "@expr", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}

	ptr8729138336 := &system.Type{Base: ptr8729164544, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Extends: system.Reference{}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors", "@kid", ptr8729139296)

	system.RegisterType("kego.io/selectors", "@basic", ptr8729136896)

	system.RegisterType("kego.io/selectors", "@c", ptr8729137216)

	system.RegisterType("kego.io/selectors", "@polykids", ptr8729140096)

	system.RegisterType("kego.io/selectors", "typed", ptr8729140576)

	system.RegisterType("kego.io/selectors", "@expr", ptr8729138336)

	system.RegisterType("kego.io/selectors", "@sibling", ptr8729140416)

	system.RegisterType("kego.io/selectors", "c", ptr8729137056)

	system.RegisterType("kego.io/selectors", "polykids", ptr8729139936)

	system.RegisterType("kego.io/selectors", "gallery", ptr8729138496)

	system.RegisterType("kego.io/selectors", "diagram", ptr8729137696)

	system.RegisterType("kego.io/selectors", "expr", ptr8729138176)

	system.RegisterType("kego.io/selectors", "@collision", ptr8729137536)

	system.RegisterType("kego.io/selectors", "kid", ptr8729139136)

	system.RegisterType("kego.io/selectors", "sibling", ptr8729140256)

	system.RegisterType("kego.io/selectors", "basic", ptr8729136736)

	system.RegisterType("kego.io/selectors", "@photo", ptr8729139776)

	system.RegisterType("kego.io/selectors", "collision", ptr8729137376)

	system.RegisterType("kego.io/selectors", "@typed", ptr8729140736)

	system.RegisterType("kego.io/selectors", "@image", ptr8729138976)

	system.RegisterType("kego.io/selectors", "photo", ptr8729139456)

	system.RegisterType("kego.io/selectors", "@gallery", ptr8729138656)

	system.RegisterType("kego.io/selectors", "image", ptr8729138816)

	system.RegisterType("kego.io/selectors", "@diagram", ptr8729138016)

}

var _ selectors.Nothing
