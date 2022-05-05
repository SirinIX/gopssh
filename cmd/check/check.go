package check

import (
	"fmt"
	"gopssh/log"
	"gopssh/pkg/config"
	"gopssh/pkg/label"
	"gopssh/pkg/port"

	"github.com/spf13/cobra"
)

type option struct {
	configFile   string
	labels       string
	withoutCache bool
}

var op = &option{}

var CheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check all IP ports in the configuration file for connectivity",
	Example: `  Simple:               gopssh check
  Specify config:       gopssh check -f config.yaml
  Select host to check: gopssh check -l app=mysql
  Check without cache:  gopssh check -f config.yaml -n`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	CheckCmd.Flags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
	CheckCmd.Flags().BoolVarP(&op.withoutCache, "without-cache", "n", false, "not use cache, default use cache")
	CheckCmd.Flags().StringVarP(&op.labels, "labels", "l", "", "label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2")
}

func execute(op *option) error {
	instances, err := config.ConfigFileToInstances(op.configFile, op.withoutCache)
	if err != nil {
		return err
	}
	instances, err = label.SelectInstances(op.labels, instances)
	if err != nil {
		return err
	}

	cantConnected := []string{}
	connected := []string{}
	for _, inst := range instances {
		if !port.CheckPort(inst.SSH.Address) {
			cantConnected = append(cantConnected, inst.SSH.Address.String())
			log.Warning("failed  to connect %s", inst.SSH.Address.String())
		} else {
			connected = append(connected, inst.SSH.Address.String())
			log.Info("succeed to connect %s", inst.SSH.Address.String())
		}
	}

	fmt.Printf("\nSucceed to connect %d IPs: %v\n", len(connected), connected)
	fmt.Printf("Failed  to connect %d IPs: %v\n", len(cantConnected), cantConnected)

	return nil
}
