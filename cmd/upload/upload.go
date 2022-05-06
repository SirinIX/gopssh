package upload

import (
	"gopssh/log"
	"gopssh/pkg/cache"
	"gopssh/pkg/config"
	"gopssh/pkg/label"

	"github.com/spf13/cobra"
)

type option struct {
	configFile   string
	labels       string
	withoutCache bool

	uploadFile string
	outputPath string
}

var op = &option{}

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file to remote",
	Example: ` Simple:                 gopssh upload -i sample.txt -o /tmp/upload.txt
  Specify config:         gopssh execute -i sample.txt -o /tmp/upload.txt -f /sample.yaml
  Select host to execute: gopssh execute -i sample.txt -o /tmp/upload.txt -l app=mysql
  Execute without cache:  gopssh execute -i sample.txt -o /tmp/upload.txt -n`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	UploadCmd.Flags().StringVarP(&op.configFile, "config-file", "f", "", "config file path")
	UploadCmd.Flags().BoolVarP(&op.withoutCache, "without-cache", "n", false, "not use cache, default use cache")
	UploadCmd.Flags().StringVarP(&op.labels, "labels", "l", "", "label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2")

	UploadCmd.Flags().StringVarP(&op.uploadFile, "upload-file", "i", "", "the file to upload")
	UploadCmd.Flags().StringVarP(&op.outputPath, "output-path", "o", "", "upload file download path")

	_ = UploadCmd.MarkFlagRequired("upload-file")
	_ = UploadCmd.MarkFlagRequired("output-path")
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

	done := make(chan bool, len(instances))
	for _, inst := range instances {
		go func(instance *cache.Instance) {
			instance.SSH.Logger = log.NewCtxLogger(map[string]interface{}{
				"host": instance.SSH.Address.Ip,
			})
			if err := instance.SSH.CopyFile(op.uploadFile, op.outputPath); err != nil {
				done <- false
				return
			}
			done <- true
		}(inst)
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}

	return nil
}
