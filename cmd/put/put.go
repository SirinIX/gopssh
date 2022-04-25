package put

import "github.com/spf13/cobra"

var PutCmd = &cobra.Command{
	Use:   "put",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
