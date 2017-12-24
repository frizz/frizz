package cmd

import (
	"log"
	"os"

	"context"

	"os/signal"

	"frizz.io/edit"
	"frizz.io/generate"
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

		ctx, cancel := context.WithCancel(context.Background())

		// trap Ctrl+C and call cancel on the context
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		go func() {
			select {
			case <-signalChannel:
				cancel()
			case <-ctx.Done():
			}
		}()
		defer signal.Stop(signalChannel)
		defer cancel()

		if err := generate.Save(os.Stdout, args[0]); err != nil {
			log.Fatal(err)
		}

		if err := edit.Open(ctx, cancel, vos.Os(), os.Stdout, args[0]); err != nil {
			log.Fatal(err)
		}
	},
}
