package cmd

import (
	"log"
	"os"

	"frizz.io/generate"

	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(GenerateCommand)
}

var GenerateCommand = &cobra.Command{
	Use:   "generate [packages]",
	Short: "Generate code",
	Long:  "Generate code for specified package(s)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Save(os.Stdout, args...); err != nil {
			log.Fatal(err)
		}
	},
}
