package ping

import "github.com/spf13/cobra"

var PingCmd = &cobra.Command{
	Use:   "ping",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
