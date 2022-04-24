package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version          = "0.0.1"
	latestCommitDate = "2022-04-22 15:04:27"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version and latest commit time of cmd-scaffold",
	Example: "cmd-scaffold version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cmd-scaffold version %v\nlatest commit time %v", version, latestCommitDate)
	},
}
