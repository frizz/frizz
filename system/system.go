//go:generate ke kego.io/system
package system // import "kego.io/system"

import "strings"

const (
	INTERFACE_PREFIX string = "$"
	RULE_PREFIX             = "@"
)

func GoName(id string) string {

	if len(id) == 0 {
		return ""
	}

	cap := func(s string, skipFirstCharacter bool) string {
		if skipFirstCharacter {
			return strings.ToUpper(s[1:2]) + s[2:]
		} else {
			return strings.ToUpper(s[0:1]) + s[1:]
		}
	}

	switch id[0:1] {
	case INTERFACE_PREFIX:
		return cap(id, true) + "Interface"
	case RULE_PREFIX:
		return cap(id, true) + "Rule"
	default:
		return cap(id, false)
	}
}

func GoInterfaceName(typeOrRuleName string) string {

	if len(typeOrRuleName) == 0 {
		return ""
	}

	if strings.HasPrefix(typeOrRuleName, RULE_PREFIX) {
		typeOrRuleName = typeOrRuleName[1:]
	}

	return GoName(INTERFACE_PREFIX + typeOrRuleName)
}
