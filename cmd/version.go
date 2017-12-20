package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(VersionCommand)
}

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Version short description",
	Long:  "Version long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version")
	},
}
