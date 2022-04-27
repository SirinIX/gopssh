package template

import (
	"fmt"
	"gopssh/pkg/file"
	"gopssh/pkg/template"

	"github.com/spf13/cobra"
)

type option struct {
	templateType string
	outputPath   string
}

var op = &option{}

var TemplateCmd = &cobra.Command{
	Use:     "template",
	Short:   "Dump config template, yaml or json",
	Example: "  gopssh template\n  gopssh template -t json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	TemplateCmd.Flags().StringVarP(&op.templateType, "type", "t", "yaml", "config template type, yaml or json")
	TemplateCmd.Flags().StringVarP(&op.outputPath, "output-path", "o", "", "output file path")
}

func execute(op *option) error {
	var cfgStr string

	if op.templateType == "yaml" {
		cfgStr = template.GetYAMLConfigTemplate()
	} else if op.templateType == "json" {
		cfgStr = template.GetJSONConfigTemplate()
	} else {
		err := fmt.Errorf("only support yaml or json")
		return err
	}
	fmt.Printf("The config template is:\n\n%v\n\n", cfgStr)

	// Save
	if op.outputPath != "" {
		if err := file.SaveStringAsFile(op.outputPath, cfgStr); err != nil {
			return err
		}
		fmt.Printf("The result is saved to %v\n", op.outputPath)
	}

	return nil
}
