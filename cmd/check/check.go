package check

import (
	"gopssh/pkg/config"
	"gopssh/pkg/port"

	"github.com/spf13/cobra"
)

type option struct {
	configFile string
}

var op = &option{}

var CheckCmd = &cobra.Command{
	Use:     "check",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	CheckCmd.PersistentFlags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
}

func execute(op *option) error {
	instances, err := config.ConfigFileToInstances(op.configFile)
	if err != nil {
		return err
	}

	for _, inst := range instances {
		port.CheckPort(inst.Address)
	}

	return nil
}
