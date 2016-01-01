package process

import (
	"fmt"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/process/validate"
)

func ValidateCommand(ctx context.Context) (err error) {

	cmd := cmdctx.FromContext(ctx)

	if cmd.Verbose {
		fmt.Print("Validating... ")
	}
	if err := validate.Validate(ctx); err != nil {
		return err
	}
	if cmd.Verbose {
		fmt.Println("OK.")
	}
	return nil
}
