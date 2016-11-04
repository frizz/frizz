package data

import (
	"context"

	"kego.io/system"
)

func (a Alajs) String(ctx context.Context) *system.String {
	out := ""
	for _, v := range a {
		out += v
	}
	return system.NewString(out)
}
