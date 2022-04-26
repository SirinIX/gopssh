package cmd

import (
	"gopssh/cmd/base64"
	"gopssh/cmd/check"
	"gopssh/cmd/execute"
	"gopssh/cmd/put"
	"gopssh/cmd/template"
	"gopssh/cmd/version"
	"gopssh/log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "A brief description of your application",
	// Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
	},
}

func init() {
	log.InitLogger()

	// rootCmd.PersistentFlags().StringP("config-file", "f", config.GetDefaultConfigFilePath(), "config file path")

	rootCmd.AddCommand(
		base64.Base64Cmd,
		check.CheckCmd,
		execute.ExecuteCmd,
		put.PutCmd,
		template.TemplateCmd,
		version.VersionCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
