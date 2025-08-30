package cmd

import (
	"messenger/bootstrap"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "messenger",
	Short: "Messenger backend",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			bootstrap.CommonModules,
		)
		app.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
