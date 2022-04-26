package template

import (
	"gopssh/pkg/config"

	"github.com/spf13/cobra"
)

type option struct {
	templateType string
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
}

func execute(op *option) error {
	return config.DumpConfigTemplate(op.templateType)
}
