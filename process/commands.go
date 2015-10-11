package process

import (
	"fmt"
	"os/exec"

	"kego.io/kerr"
	"kego.io/process/settings"
	"kego.io/process/validate"
)

func ValidateCommand(set *settings.Settings) (typesChanged bool, err error) {
	if set.Verbose {
		fmt.Print("Validating... ")
	}
	if err := validate.Validate(set); err != nil {
		if _, ok := kerr.Source(err).(validate.TypesChangedError); ok {
			if set.Verbose {
				fmt.Println("Types have changed - rebuilding...")
			}
			if err := RunAllCommands(set); err != nil {
				return true, err
			}
			return true, nil
		}
		return false, err
	}
	if set.Verbose {
		fmt.Println("OK.")
	}
	return false, nil
}

// Ke is the main process fired by the ke command line tool.
func KeCommand(set *settings.Settings) error {

	for p, a := range set.Aliases {
		if set.Verbose {
			if set.Update {
				fmt.Print("Updating package ", a, "... ")
			} else {
				fmt.Print("Getting package ", a, "... ")
			}
		}
		params := []string{"get"}
		if set.Update {
			params = append(params, "-u")
		}
		if set.Verbose {
			params = append(params, "-v")
		}
		params = append(params, p)

		combined, stdout, stderr := logger(set.Verbose)
		cmd := exec.Command("go", params...)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return kerr.New("HHKSTQMAKG", err, "go get command: %s", combined.String())
		} else {
			if set.Verbose {
				fmt.Println("OK.")
			}
		}
	}

	if err := RunAllCommands(set); err != nil {
		return err
	}
	return nil
}

func RunAllCommands(set *settings.Settings) error {
	if err := Run(C_STRUCTS, set); err != nil {
		return err
	}
	if err := Run(C_TYPES, set); err != nil {
		return err
	}
	if err := Run(C_KE, set); err != nil {
		return err
	}
	return nil
}
