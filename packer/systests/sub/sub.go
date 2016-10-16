package sub

import (
	"fmt"

	"context"

	"kego.io/system"
)

func (a *A) GetString(ctx context.Context) *system.String {
	return system.NewString(fmt.Sprintf("%v", *a))
}
