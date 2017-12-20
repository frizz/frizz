package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "frizz",
	Short: "Frizz short description",
	Long:  "Frizz long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Frizz")
	},
}

func init() {
	/*
		rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
		rootCommand.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
		rootCommand.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
		rootCommand.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
		rootCommand.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
		viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
		viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
		viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
		viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
		viper.SetDefault("license", "apache")
	*/
}
