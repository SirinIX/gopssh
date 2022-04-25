package ping

import "github.com/spf13/cobra"

type option struct {
}

var op = &option{}

var PingCmd = &cobra.Command{
	Use:   "ping",
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
}

func execute(op *option) error {
	return nil
}