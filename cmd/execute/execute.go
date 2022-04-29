package execute

import (
	"gopssh/log"
	"gopssh/pkg/cache"
	"gopssh/pkg/config"
	"gopssh/pkg/label"

	"github.com/spf13/cobra"
)

type option struct {
	command      string
	configFile   string
	labels       string
	withoutCache bool
}

var op = &option{}

var ExecuteCmd = &cobra.Command{
	Use:     "execute",
	Short:   "Execute command and return result",
	Example: `  Simple:                 gopssh execute -c 'ls -l'
  Specify config:         gopssh execute -c 'ls -l' -f /sample.yaml
  Select host to execute: gopssh execute -c 'ls -l' -l app=mysql
  Execute without cache:  gopssh execute -c 'ls -l' -n`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	ExecuteCmd.Flags().StringVarP(&op.command, "command", "c", "", "command to execute")
	ExecuteCmd.Flags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
	ExecuteCmd.Flags().BoolVarP(&op.withoutCache, "without-cache", "n", false, "not use cache, default use cache")
	ExecuteCmd.Flags().StringVarP(&op.labels, "labels", "l", "", "label to filter on, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2")

	ExecuteCmd.MarkFlagRequired("command")
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
		done := make(chan bool)
		go func(instance *cache.Instance) {
			instance.SSH.Logger = log.NewCtxLogger(map[string]interface{}{
				"host": instance.SSH.Address.Ip,
			})
			res, err := instance.SSH.Command(op.command)
			if err != nil {
				return
			}
			log.Info("command execute result is: \n%s", res.String())
			done <- true
		}(inst)
		<-done
	}

	return nil
}
