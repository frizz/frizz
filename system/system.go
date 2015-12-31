//go:generate ke kego.io/system
package system // import "kego.io/system"

import (
	"strings"

	"kego.io/json"
)

func GoName(id string) string {

	if len(id) == 0 {
		return ""
	}

	switch id[0:1] {
	case json.RULE_PREFIX:
		return capitalise(id, true) + "Rule"
	default:
		return capitalise(id, false)
	}
}

func GoInterfaceName(typeOrRuleName string) string {

	if len(typeOrRuleName) == 0 {
		return ""
	}

	switch typeOrRuleName[0:1] {
	case json.RULE_PREFIX:
		return capitalise(typeOrRuleName, true) + "Interface"
	default:
		return capitalise(typeOrRuleName, false) + "Interface"
	}
}

func capitalise(s string, skipFirstCharacter bool) string {
	if skipFirstCharacter {
		return strings.ToUpper(s[1:2]) + s[2:]
	} else {
		return strings.ToUpper(s[0:1]) + s[1:]
	}
}
