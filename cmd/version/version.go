package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "1.0.0"

	latestCommitDate = "2022-05-06 15:05:56"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version and latest commit time of gopssh",
	Example: `  Simple: gopssh version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gopssh version %v\nlatest commit time %v", version, latestCommitDate)
	},
}
