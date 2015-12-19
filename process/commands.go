package process

import (
	"fmt"
	"os/exec"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/kerr"
	"kego.io/process/validate"
)

func ValidateCommand(ctx context.Context) (typesChanged bool, err error) {

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		return false, kerr.New("MISESATNVN", nil, "No cmd in ctx")
	}

	if cmd.Verbose {
		fmt.Print("Validating... ")
	}
	if err := validate.Validate(ctx); err != nil {
		if _, ok := kerr.Source(err).(validate.TypesChangedError); ok {
			if cmd.Verbose {
				fmt.Println("Types have changed - rebuilding...")
			}
			if err := RunAllCommands(ctx); err != nil {
				return true, err
			}
			return true, nil
		}
		return false, err
	}
	if cmd.Verbose {
		fmt.Println("OK.")
	}
	return false, nil
}

// Ke is the main process fired by the ke command line tool.
func KeCommand(ctx context.Context) error {

	cmd, ok := cmdctx.FromContext(ctx)
	if !ok {
		return kerr.New("UXBXCPJLLN", nil, "No cmd in ctx")
	}

	env, ok := envctx.FromContext(ctx)
	if !ok {
		return kerr.New("GUJHVCFFGJ", nil, "No env in ctx")
	}

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

	if err := RunAllCommands(ctx); err != nil {
		return err
	}
	return nil
}

func RunAllCommands(ctx context.Context) error {
	if err := Run(ctx, C_STRUCTS); err != nil {
		return err
	}
	if err := Run(ctx, C_TYPES); err != nil {
		return err
	}
	if err := Run(ctx, C_LOCAL_KE); err != nil {
		return err
	}
	return nil
}
