package views

import (
	"context"
	"math/rand"

	"bytes"

	"frizz.io/editor/client/stores"
	"frizz.io/flux"
)

type View struct {
	*flux.View
	App *stores.App
}

func New(ctx context.Context, self flux.ViewInterface) *View {
	app := stores.FromContext(ctx)
	return &View{
		View: flux.NewView(ctx, self, app),
		App:  app,
	}
}

func randomId() string {
	randInt := func(min int, max int) int {
		return min + rand.Intn(max-min)
	}
	var result bytes.Buffer
	var temp string
	for i := 0; i < 20; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
