package deploy

import "github.com/spf13/cobra"

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the easy-manager application product",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
