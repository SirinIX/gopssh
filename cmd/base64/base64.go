package base64

import "github.com/spf13/cobra"

var Base64Cmd = &cobra.Command{
	Use:   "base64",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
