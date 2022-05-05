package get

import (
	"gopssh/log"
	"gopssh/pkg/config"
	"gopssh/pkg/label"

	"github.com/spf13/cobra"
)

type option struct {
	configFile   string
	labels       string
	withoutCache bool
	// decodePassword bool
}

var op = &option{}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get connection instance of config file",
	Example: `  Simple:                 gopssh get
  Specify config:         gopssh get -f /sample.yaml
  Select host to execute: gopssh get -l app=mysql
  Execute without cache:  gopssh get -n`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	GetCmd.Flags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
	GetCmd.Flags().BoolVarP(&op.withoutCache, "without-cache", "n", false, "not use cache, default use cache")
	GetCmd.Flags().StringVarP(&op.labels, "labels", "l", "", "label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2")

	// GetCmd.Flags().BoolVarP(&op.decodePassword, "decode-password", "d", false, "decode password with base64")
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

	for _, inst := range instances {
		log.Info(inst.String())
	}

	return nil
}
