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
	labels     *[]string
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
	// https://github.com/kubernetes/kubernetes/blob/ea07644522/staging/src/k8s.io/kubectl/pkg/cmd/label/label.go
	// cmd.Flags().StringVarP(&o.selector, "selector", "l", o.selector, "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2).")
	op.labels = CheckCmd.Flags().StringArrayP("labels", "l", []string{"all=all"}, "host labels")
}

func execute(op *option) error {
	instances, err := config.ConfigFileToInstances(op.configFile)
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
