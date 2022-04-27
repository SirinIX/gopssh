package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"

	latestCommitDate = "2022-04-27 10:04:57"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version and latest commit time of gopssh",
	Example: "  gopssh version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gopssh version %v\nlatest commit time %v", version, latestCommitDate)
	},
}
