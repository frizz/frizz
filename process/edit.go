package process

import (
	"kego.io/kerr"
	"kego.io/process/editor"
)

func EditCommand(set settings) error {
	err := editor.StartServer(set.path)
	if err != nil {
		return kerr.New("EGGSFTSGWM", err, "editor.StartServer")
	}
	return nil
}
