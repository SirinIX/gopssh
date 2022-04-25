package execute

import "github.com/spf13/cobra"

var ExecuteCmd = &cobra.Command{
	Use:   "execute",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
