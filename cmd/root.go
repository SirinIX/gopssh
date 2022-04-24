package cmd

import (
	"cmd-scaffold/cmd/deploy"
	"cmd-scaffold/cmd/upgrade"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
		deploy.DeployCmd,
		upgrade.UpgradeCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
