package process

import (
	"fmt"
	"os/exec"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/editor/server"
	"kego.io/kerr"
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

type CommandType int

const (
	Ke CommandType = iota
	Ked
)

// Ke is the main process fired by the ke command line tool.
func KeCommand(ctx context.Context, typ CommandType) error {

	cmd := cmdctx.FromContext(ctx)
	env := envctx.FromContext(ctx)

	for p, a := range env.Aliases {
		if cmd.Verbose {
			if cmd.Update {
				fmt.Print("Updating package ", a, "... ")
			} else {
				fmt.Print("Getting package ", a, "... ")
			}
		}
		params := []string{"get"}
		if cmd.Update {
			params = append(params, "-u")
		}
		if cmd.Verbose {
			params = append(params, "-v")
		}
		params = append(params, p)

		combined, stdout, stderr := logger(cmd.Verbose)
		exe := exec.Command("go", params...)
		exe.Stdout = stdout
		exe.Stderr = stderr

		if err := exe.Run(); err != nil {
			return kerr.New("HHKSTQMAKG", err, "go get command: %s", combined.String())
		} else {
			if cmd.Verbose {
				fmt.Println("OK.")
			}
		}
	}

	switch typ {
	case Ke:
		if err := RunAllCommands(ctx); err != nil {
			return err
		}
	case Ked:
		// Something
		if err := server.Start(ctx); err != nil {
			return kerr.New("PEUTNVGTSX", err, "server.Start")
		}
	default:
		panic(kerr.New("NVFYQAMGLW", nil, "Invalid typ"))
	}
	return nil
}

func RunAllCommands(ctx context.Context) error {

	//if err := Run(ctx, C_STRUCTS); err != nil {
	//	return err
	//}
	if err := Run(ctx, C_TYPES); err != nil {
		return err
	}
	if err := Run(ctx, C_LOCAL_KE); err != nil {
		return err
	}
	return nil
}
