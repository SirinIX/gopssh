package base64

import (
	"gopssh/log"
	"gopssh/pkg/base64"

	"github.com/spf13/cobra"
)

var (
	decode  bool
	content string
)

var Base64Cmd = &cobra.Command{
	Use: "base64",
	Example: `  Encode: gopssh base64 -c 'root$123'
  Decode: gopssh base64 -d -c 'cm9vdA=='`,
	Short: "base64 encode or decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(decode, content)
	},
}

func init() {
	Base64Cmd.Flags().BoolVarP(&decode, "decode", "d", false, "decode or encode, default is encode")
	Base64Cmd.Flags().StringVarP(&content, "content", "c", "", "decode / encode data content (required)")

	Base64Cmd.MarkFlagRequired("content")
}

func execute(decode bool, content string) error {
	if decode {
		decStr, err := base64.Decode(content)
		if err != nil {
			return err
		}
		log.Info("the base64 decoded result of %v is: %v", content, decStr)
	} else {
		encStr := base64.Encode(content)
		log.Info("the base64 encoded result of %v is: %v", content, encStr)
	}
	return nil
}
