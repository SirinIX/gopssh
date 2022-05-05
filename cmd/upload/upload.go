package upload

import "github.com/spf13/cobra"

type option struct {
	configFile   string
	labels       string
	withoutCache bool

	uploadFile string
	outputPath string
}

var op = &option{}

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Upload file to remote",
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

	UploadCmd.MarkFlagRequired("upload-file")
	UploadCmd.MarkFlagRequired("output-path")
}

func execute(op *option) error {
	return nil
}
