//go:generate ke kego.io/system
package system // import "kego.io/system"

import "strings"

const (
	INTERFACE_PREFIX string = "^"
	RULE_PREFIX             = "£"
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
	case "@":
		return cap(id, true) + "_rule"
	case "$":
		return cap(id, true) + "_base"
	default:
		return cap(id, false)
	}
}
