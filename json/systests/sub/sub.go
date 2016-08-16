package sub

import (
	"fmt"

	"context"

	"kego.io/system"
)

type A float64

func (a *A) GetString(ctx context.Context) *system.String {
	return system.NewString(fmt.Sprintf("%v", *a))
}
