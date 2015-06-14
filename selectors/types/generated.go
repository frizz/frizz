package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/jsonselect"
)

func init() {

	ptr8729930048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730132736 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729930048, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729939920 := &system.Type{Base: ptr8730132736, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729928928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729714176 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928928, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729928800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729714304 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729928768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729714432 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729928736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729714560 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928736, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729374256 := &jsonselect.Image_rule{Base: ptr8729714560, RuleBase: (*system.RuleBase)(nil)}

	ptr8729806208 := &system.Map_rule{Base: ptr8729714432, RuleBase: (*system.RuleBase)(nil), Items: ptr8729374256, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729844816 := &system.Property{Base: ptr8729714304, Defaulter: false, Item: ptr8729806208, Optional: false}

	ptr8729939392 := &system.Type{Base: ptr8729714176, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"images": ptr8729844816}, Rule: (*system.Type)(nil)}

	ptr8730043584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137216 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043584, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729942208 := &system.Type{Base: ptr8730137216, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729932800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135424 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729932800, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729941856 := &system.Type{Base: ptr8730135424, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729929344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730132480 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729929344, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729939568 := &system.Type{Base: ptr8730132480, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729929728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133888 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729929728, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729929536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134528 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729929536, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729929504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134656 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729929504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729941328 := &system.String_rule{Base: ptr8730134656, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729570304 := &system.Property{Base: ptr8730134528, Defaulter: false, Item: ptr8729941328, Optional: false}

	ptr8729928064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134016 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928064, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729928032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134144 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928032, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729940976 := &system.String_rule{Base: ptr8730134144, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729570016 := &system.Property{Base: ptr8730134016, Defaulter: false, Item: ptr8729940976, Optional: true}

	ptr8729928608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134272 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928608, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729928576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134400 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729928576, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729941152 := &system.String_rule{Base: ptr8730134400, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729569776 := &system.Property{Base: ptr8730134272, Defaulter: false, Item: ptr8729941152, Optional: false}

	ptr8729940800 := &system.Type{Base: ptr8730133888, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"path": ptr8729570304, "protocol": ptr8729570016, "server": ptr8729569776}, Rule: (*system.Type)(nil)}

	ptr8730044960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711616 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044960, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8730044704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711744 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044704, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730044672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711872 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044672, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729938160 := &system.String_rule{Base: ptr8729711872, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729843328 := &system.Property{Base: ptr8729711744, Defaulter: false, Item: ptr8729938160, Optional: false}

	ptr8729937984 := &system.Type{Base: ptr8729711616, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"url": ptr8729843328}, Rule: (*system.Type)(nil)}

	ptr8730045536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712000 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045536, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729938336 := &system.Type{Base: ptr8729712000, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729460224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710592 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729460224, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729458720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729707392 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729458720, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729458560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729707520 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729458560, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919488 := &system.Number_rule{Base: ptr8729707520, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729842416 := &system.Property{Base: ptr8729707392, Defaulter: false, Item: ptr8729919488, Optional: false}

	ptr8729459328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729707648 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729459328, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729459168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729707776 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729459168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919584 := &system.Number_rule{Base: ptr8729707776, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729842560 := &system.Property{Base: ptr8729707648, Defaulter: false, Item: ptr8729919584, Optional: false}

	ptr8729460096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708928 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729460096, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729460064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709056 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729460064, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729459904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710720 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729459904, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919680 := &system.Number_rule{Base: ptr8729710720, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729804864 := &system.Map_rule{Base: ptr8729709056, RuleBase: (*system.RuleBase)(nil), Items: ptr8729919680, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729842752 := &system.Property{Base: ptr8729708928, Defaulter: false, Item: ptr8729804864, Optional: false}

	ptr8729937104 := &system.Type{Base: ptr8729710592, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729842416, "b": ptr8729842560, "c": ptr8729842752}, Rule: (*system.Type)(nil)}

	ptr8730050368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137344 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050368, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8730045632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137472 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045632, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730045376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137600 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045376, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730045344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137728 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045344, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729942560 := &system.String_rule{Base: ptr8730137728, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729807232 := &system.Map_rule{Base: ptr8730137600, RuleBase: (*system.RuleBase)(nil), Items: ptr8729942560, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8730091952 := &system.Property{Base: ptr8730137472, Defaulter: false, Item: ptr8729807232, Optional: false}

	ptr8730046464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137856 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046464, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730046336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137984 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046336, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729942736 := &system.String_rule{Base: ptr8730137984, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8730092048 := &system.Property{Base: ptr8730137856, Defaulter: false, Item: ptr8729942736, Optional: false}

	ptr8730047264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138112 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047264, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730047200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138240 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047200, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730047168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138368 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8730211456 := &jsonselect.Kid_rule{Base: ptr8730138368, RuleBase: (*system.RuleBase)(nil)}

	ptr8729807360 := &system.Map_rule{Base: ptr8730138240, RuleBase: (*system.RuleBase)(nil), Items: ptr8730211456, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8730092096 := &system.Property{Base: ptr8730138112, Defaulter: false, Item: ptr8729807360, Optional: false}

	ptr8730047776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138496 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047776, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730047744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138624 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047744, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8730211664 := &jsonselect.Image_rule{Base: ptr8730138624, RuleBase: (*system.RuleBase)(nil)}

	ptr8730092192 := &system.Property{Base: ptr8730138496, Defaulter: false, Item: ptr8730211664, Optional: false}

	ptr8730048960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138752 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048960, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730048864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730138880 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048864, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8730048800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730139008 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729942912 := &system.String_rule{Base: ptr8730139008, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729962048 := &system.Array_rule{Base: ptr8730138880, RuleBase: (*system.RuleBase)(nil), Items: ptr8729942912, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8730092240 := &system.Property{Base: ptr8730138752, Defaulter: false, Item: ptr8729962048, Optional: false}

	ptr8730050176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730139136 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050176, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730049792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730139264 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049792, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729920352 := &system.Number_rule{Base: ptr8730139264, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8730092336 := &system.Property{Base: ptr8730139136, Defaulter: false, Item: ptr8729920352, Optional: false}

	ptr8729942384 := &system.Type{Base: ptr8730137344, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8730091952, "favoriteColor": ptr8730092048, "kids": ptr8730092096, "avatar": ptr8730092192, "drinkPreference": ptr8730092240, "weight": ptr8730092336}, Rule: (*system.Type)(nil)}

	ptr8729929632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730132608 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729929632, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729939744 := &system.Type{Base: ptr8730132608, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729931968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134912 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931968, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729931808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135040 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931808, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729931680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135168 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729931552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135296 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8730207568 := &jsonselect.Kid_rule{Base: ptr8730135296, RuleBase: (*system.RuleBase)(nil)}

	ptr8729961168 := &system.Array_rule{Base: ptr8730135168, RuleBase: (*system.RuleBase)(nil), Items: ptr8730207568, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729843232 := &system.Property{Base: ptr8730135040, Defaulter: false, Item: ptr8729961168, Optional: false}

	ptr8729941680 := &system.Type{Base: ptr8730134912, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729843232}, Rule: (*system.Type)(nil)}

	ptr8730050240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729707904 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050240, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8730045280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708416 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045280, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730045248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708544 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730045248, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729936224 := &system.String_rule{Base: ptr8729708544, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729843664 := &system.Property{Base: ptr8729708416, Defaulter: false, Item: ptr8729936224, Optional: false}

	ptr8730046368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708672 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046368, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730046272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708800 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046272, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8730046240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709184 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730046208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709312 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046208, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729936400 := &system.String_rule{Base: ptr8729709312, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729805632 := &system.Map_rule{Base: ptr8729709184, RuleBase: (*system.RuleBase)(nil), Items: ptr8729936400, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729960688 := &system.Array_rule{Base: ptr8729708800, RuleBase: (*system.RuleBase)(nil), Items: ptr8729805632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729843712 := &system.Property{Base: ptr8729708672, Defaulter: false, Item: ptr8729960688, Optional: false}

	ptr8730048416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709440 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048416, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730048320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709568 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048320, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8730048288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709696 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729936576 := &system.String_rule{Base: ptr8729709696, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729960848 := &system.Array_rule{Base: ptr8729709568, RuleBase: (*system.RuleBase)(nil), Items: ptr8729936576, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729844432 := &system.Property{Base: ptr8729709440, Defaulter: false, Item: ptr8729960848, Optional: false}

	ptr8730049152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709824 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049152, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730049056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729709952 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049056, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8730049024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710080 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729936752 := &system.String_rule{Base: ptr8729710080, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729960928 := &system.Array_rule{Base: ptr8729709952, RuleBase: (*system.RuleBase)(nil), Items: ptr8729936752, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729844528 := &system.Property{Base: ptr8729709824, Defaulter: false, Item: ptr8729960928, Optional: false}

	ptr8730050080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710208 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050080, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730049888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710336 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049888, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729926880 := &system.Number_rule{Base: ptr8729710336, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729844624 := &system.Property{Base: ptr8729710208, Defaulter: false, Item: ptr8729926880, Optional: false}

	ptr8730044832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708032 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044832, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730044800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708160 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730044768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729708288 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730044768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729936048 := &system.String_rule{Base: ptr8729708288, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729805440 := &system.Map_rule{Base: ptr8729708160, RuleBase: (*system.RuleBase)(nil), Items: ptr8729936048, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729842656 := &system.Property{Base: ptr8729708032, Defaulter: false, Item: ptr8729805440, Optional: false}

	ptr8729689840 := &system.Type{Base: ptr8729707904, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"favoriteColor": ptr8729843664, "languagesSpoken": ptr8729843712, "seatingPreference": ptr8729844432, "drinkPreference": ptr8729844528, "weight": ptr8729844624, "name": ptr8729842656}, Rule: (*system.Type)(nil)}

	ptr8729460640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710848 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729460640, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729937280 := &system.Type{Base: ptr8729710848, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8730050304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712128 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050304, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8730047680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712768 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730047648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712896 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047648, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729938688 := &system.String_rule{Base: ptr8729712896, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729843808 := &system.Property{Base: ptr8729712768, Defaulter: false, Item: ptr8729938688, Optional: false}

	ptr8730048128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713024 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048128, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730048096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713152 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048096, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729938864 := &system.String_rule{Base: ptr8729713152, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729843856 := &system.Property{Base: ptr8729713024, Defaulter: false, Item: ptr8729938864, Optional: false}

	ptr8730048768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713280 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730048736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713408 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730048736, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729939040 := &system.String_rule{Base: ptr8729713408, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729843904 := &system.Property{Base: ptr8729713280, Defaulter: false, Item: ptr8729939040, Optional: false}

	ptr8730049504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713536 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730049472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713664 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730049472, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8730049248 := &system.Bool_rule{Base: ptr8729713664, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729843952 := &system.Property{Base: ptr8729713536, Defaulter: false, Item: ptr8730049248, Optional: false}

	ptr8730050144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713792 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730050112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729713920 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730050112, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8730049952 := &system.Bool_rule{Base: ptr8729713920, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729844384 := &system.Property{Base: ptr8729713792, Defaulter: false, Item: ptr8730049952, Optional: false}

	ptr8730046624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712256 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046624, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730046432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712384 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730046432, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919776 := &system.Number_rule{Base: ptr8729712384, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729843568 := &system.Property{Base: ptr8729712256, Defaulter: false, Item: ptr8729919776, Optional: false}

	ptr8730047232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712512 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047232, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730047072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729712640 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730047072, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919872 := &system.Number_rule{Base: ptr8729712640, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729843616 := &system.Property{Base: ptr8729712512, Defaulter: false, Item: ptr8729919872, Optional: false}

	ptr8729938512 := &system.Type{Base: ptr8729712128, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"string": ptr8729843808, "string2": ptr8729843856, "null": ptr8729843904, "true": ptr8729843952, "false": ptr8729844384, "int": ptr8729843568, "float": ptr8729843616}, Rule: (*system.Type)(nil)}

	ptr8729458464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730139392 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729458464, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729943088 := &system.Type{Base: ptr8730139392, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729932192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133760 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729932192, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729940624 := &system.Type{Base: ptr8730133760, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729927904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729714048 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729927904, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729939216 := &system.Type{Base: ptr8729714048, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729931776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730132864 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931776, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729930720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730132992 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729930720, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729930688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133120 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729930688, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729940272 := &system.String_rule{Base: ptr8730133120, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729568912 := &system.Property{Base: ptr8730132992, Defaulter: false, Item: ptr8729940272, Optional: false}

	ptr8729931168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133248 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729931136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133376 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729940448 := &system.String_rule{Base: ptr8730133376, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729569104 := &system.Property{Base: ptr8730133248, Defaulter: false, Item: ptr8729940448, Optional: false}

	ptr8729931648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133504 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931648, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729931616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730133632 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729931616, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729931488 := &system.Bool_rule{Base: ptr8730133632, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729569296 := &system.Property{Base: ptr8730133504, Defaulter: false, Item: ptr8729931488, Optional: true}

	ptr8729940096 := &system.Type{Base: ptr8730132864, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"language": ptr8729568912, "level": ptr8729569104, "preferred": ptr8729569296}, Rule: (*system.Type)(nil)}

	ptr8729457984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710464 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729457984, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729936928 := &system.Type{Base: ptr8729710464, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8730043296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729710976 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043296, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8730043168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711104 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730043136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711232 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730043104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711360 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043104, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729937632 := &system.String_rule{Base: ptr8729711360, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729805312 := &system.Map_rule{Base: ptr8729711232, RuleBase: (*system.RuleBase)(nil), Items: ptr8729937632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729843040 := &system.Property{Base: ptr8729711104, Defaulter: false, Item: ptr8729805312, Optional: false}

	ptr8729937456 := &system.Type{Base: ptr8729710976, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8729843040}, Rule: (*system.Type)(nil)}

	ptr8729930432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730134784 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729930432, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729941504 := &system.Type{Base: ptr8730134784, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8730043808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729711488 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730043808, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729937808 := &system.Type{Base: ptr8729711488, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8730042912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135552 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730042912, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729933632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135680 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729933632, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729933472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135808 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729933472, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729919968 := &system.Number_rule{Base: ptr8730135808, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729845008 := &system.Property{Base: ptr8730135680, Defaulter: false, Item: ptr8729919968, Optional: false}

	ptr8729934208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730135936 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729934208, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729934048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136064 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729934048, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729920064 := &system.Number_rule{Base: ptr8730136064, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729845056 := &system.Property{Base: ptr8730135936, Defaulter: false, Item: ptr8729920064, Optional: false}

	ptr8729934656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136192 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729934656, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729934624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136320 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729934624, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@c", Package: "kego.io/jsonselect", Type: "@c", Exists: true}}

	ptr8730209040 := &jsonselect.C_rule{Base: ptr8730136320, RuleBase: (*system.RuleBase)(nil)}

	ptr8729845152 := &system.Property{Base: ptr8730136192, Defaulter: false, Item: ptr8730209040, Optional: false}

	ptr8729935424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136448 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729935424, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729935392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136576 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729935392, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729935232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136704 := &system.Base{Base: (*system.Base)(nil), Context: ptr8729935232, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729920160 := &system.Number_rule{Base: ptr8730136704, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729806784 := &system.Map_rule{Base: ptr8730136576, RuleBase: (*system.RuleBase)(nil), Items: ptr8729920160, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8730091568 := &system.Property{Base: ptr8730136448, Defaulter: false, Item: ptr8729806784, Optional: false}

	ptr8730042752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136832 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730042752, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8730042720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730136960 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730042720, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8730042496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8730137088 := &system.Base{Base: (*system.Base)(nil), Context: ptr8730042496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729920256 := &system.Number_rule{Base: ptr8730137088, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729806912 := &system.Map_rule{Base: ptr8730136960, RuleBase: (*system.RuleBase)(nil), Items: ptr8729920256, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8730091664 := &system.Property{Base: ptr8730136832, Defaulter: false, Item: ptr8729806912, Optional: false}

	ptr8729942032 := &system.Type{Base: ptr8730135552, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729845008, "b": ptr8729845056, "c": ptr8729845152, "d": ptr8730091568, "e": ptr8730091664}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8729936928)

	system.RegisterType("kego.io/jsonselect:@c", ptr8729937280)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8729937808)

	system.RegisterType("kego.io/jsonselect:@diagram", ptr8729938336)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8729939216)

	system.RegisterType("kego.io/jsonselect:@gallery", ptr8729939568)

	system.RegisterType("kego.io/jsonselect:@image", ptr8729939920)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8729940624)

	system.RegisterType("kego.io/jsonselect:@photo", ptr8729941504)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8729941856)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8729942208)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8729943088)

	system.RegisterType("kego.io/jsonselect:basic", ptr8729689840)

	system.RegisterType("kego.io/jsonselect:c", ptr8729937104)

	system.RegisterType("kego.io/jsonselect:collision", ptr8729937456)

	system.RegisterType("kego.io/jsonselect:diagram", ptr8729937984)

	system.RegisterType("kego.io/jsonselect:expr", ptr8729938512)

	system.RegisterType("kego.io/jsonselect:gallery", ptr8729939392)

	system.RegisterType("kego.io/jsonselect:image", ptr8729939744)

	system.RegisterType("kego.io/jsonselect:kid", ptr8729940096)

	system.RegisterType("kego.io/jsonselect:photo", ptr8729940800)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8729941680)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8729942032)

	system.RegisterType("kego.io/jsonselect:typed", ptr8729942384)

}
