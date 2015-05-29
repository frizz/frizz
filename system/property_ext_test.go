package system_test

import (
	"testing"

	"kego.io/system"
	_ "kego.io/system/types"
)

func TestGoTypeDescriptorErrors(t *testing.T) {
	system.GoTypeDescriptorErrors_NeedsTypes(t)
}
