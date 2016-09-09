//go:generate ke kego.io/system
package system // import "kego.io/system"

// ke: {"package": {"complete": true}}

import (
	"strings"

	"kego.io/context/jsonctx"
)

func GoName(id string) string {

	if len(id) == 0 {
		return ""
	}

	rule := id[0:1] == jsonctx.RULE_PREFIX

	if rule {
		id = id[1:]
	}

	words := strings.Split(id, "-")
	var name string
	for _, word := range words {
		name = name + strings.Title(word)
	}
	if rule {
		name = name + "Rule"
	}
	return name
}

func GoInterfaceName(id string) string {

	if len(id) == 0 {
		return ""
	}

	rule := id[0:1] == jsonctx.RULE_PREFIX

	if rule {
		id = id[1:]
	}

	words := strings.Split(id, "-")
	var name string
	for _, word := range words {
		name = name + strings.Title(word)
	}
	return name + "Interface"
}
