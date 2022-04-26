package cmd

import (
	"gopssh/cmd/base64"
	"gopssh/cmd/execute"
	"gopssh/cmd/ping"
	"gopssh/cmd/put"
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

	

	rootCmd.AddCommand(
		versionCmd,
		ping.PingCmd,
		execute.ExecuteCmd,
		put.PutCmd,
		base64.Base64Cmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
