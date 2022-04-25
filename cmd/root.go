package cmd

import (
	"fmt"
	"gopssh/cmd/execute"
	"gopssh/cmd/put"
	"gopssh/cmd/ping"

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
	rootCmd.AddCommand(
		versionCmd,
		ping.PingCmd,
		execute.ExecuteCmd,
		put.PutCmd,
	)

	// A flag can be 'persistent', meaning that this flag will be available to the command it's assigned to as well as every command under that command.
	// For global flags, assign a flag as a persistent flag on the root.
	loggerLevel := rootCmd.PersistentFlags().IntP("log-level", "v", 4, "the level of logger, debug 5, info 4, warn 3, error 2")
	cobra.OnlyValidArgs(rootCmd, []string{"deploy"})
	fmt.Println(loggerLevel)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
