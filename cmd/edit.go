package cmd

import (
	"log"
	"os"

	"frizz.io/generate"

	"context"

	"frizz.io/edit"
	"github.com/dave/patsy/vos"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(EditCommand)
}

var EditCommand = &cobra.Command{
	Use:   "edit [package]",
	Short: "Edit package",
	Long:  "Open editor for specified package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Save(os.Stdout, args[0]); err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		if err := edit.Open(ctx, cancel, vos.Os(), os.Stdout, args[0]); err != nil {
			log.Fatal(err)
		}
	},
}
