package convert

import (
	"fmt"

	"gopssh/pkg/convert"
	"gopssh/pkg/file"

	"github.com/spf13/cobra"
)

type option struct {
	yamlPath   string
	jsonPath   string
	outputPath string
}

var op = &option{}

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert json to yaml, or yaml to json",
	Example: `  Convert YAML to JSON: gopssh convert -y sample.yaml
  Convert JSON to YAML: gopssh convert -j sample.json
  Convert and save:     gopssh convert -j sample.yaml -o convert.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&op.outputPath, "output-path", "o", "", "output file path")
	ConvertCmd.Flags().StringVarP(&op.yamlPath, "yaml-path", "y", "", "yaml file path")
	ConvertCmd.Flags().StringVarP(&op.jsonPath, "json-path", "j", "", "json file path")
}

func execute(op *option) error {
	if op.yamlPath != "" &&  op.jsonPath != ""{
		err := fmt.Errorf("only one of yaml-path or json-path can be specified")
		return err
	}
	if op.yamlPath == "" && op.jsonPath == "" {
		err := fmt.Errorf("one of yaml-path or json-path must be specified")
		return err
	}

	// Convert 
	convertFunc := func (f convert.Convert, inputPath , outputPath string) error {
		bs, err := f(inputPath)
		if err != nil {
			return err
		}
		fmt.Printf("File %v convert result is: \n\n%v\n\n", inputPath, string(bs))

		if outputPath != "" {
			if err := file.SaveBytesAsFile(outputPath, bs); err != nil {
				return err
			}
			fmt.Printf("The result is saved to %v\n", outputPath)
		}

		return nil
	}

	if op.yamlPath != "" {
		if err := convertFunc(convert.YAMLFileToJSONBytes, op.yamlPath, op.outputPath); err != nil {
			return err
		}
	}
	if op.jsonPath != "" {
		if err := convertFunc(convert.JSONFileToYAMLBytes, op.jsonPath, op.outputPath); err != nil {
			return err
		}
	}

	return nil
}
