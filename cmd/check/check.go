package check

import (
	"fmt"
	"gopssh/log"
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
	Short:   "Check all IP ports in the configuration file for connectivity",
	Example: "  gopssh check -f config.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	CheckCmd.Flags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
}

func execute(op *option) error {
	instances, err := config.ConfigFileToInstances(op.configFile)
	if err != nil {
		return err
	}

	cantConnected := []string{}
	connected := []string{}
	for _, inst := range instances {
		if !port.CheckPort(inst.Address) {
			cantConnected = append(cantConnected, inst.Address.String())
			log.Warning("failed  to connect %s", inst.Address.String())
		} else {
			connected = append(connected, inst.Address.String())
			log.Info("succeed to connect %s", inst.Address.String())
		}
	}
	
	fmt.Printf("\nSuccess %d IPs: %v\n", len(connected), connected)
	fmt.Printf("Fail %d IPs: %v\n", len(cantConnected), cantConnected)

	return nil
}
