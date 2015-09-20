//go:generate ke kego.io/system
package system // import "kego.io/system"

import (
	"fmt"
	"strings"
)

func GoName(id string) string {
	if id == "" {
		return ""
	}
	if strings.HasPrefix(id, "@") {
		// Rules are prefixed with @ in the json ID, and suffixed with _rule in the
		// golang type name.
		return fmt.Sprintf("%s%s_rule", strings.ToUpper(id[1:2]), id[2:])
	} else if strings.HasPrefix(id, "$") {
		// Base types are prefixed with $ in the json ID, and suffixed with _base in the
		// golang type name.
		return fmt.Sprintf("%s%s_base", strings.ToUpper(id[1:2]), id[2:])
	} else {
		return fmt.Sprintf("%s%s", strings.ToUpper(id[0:1]), id[1:])
	}
}
