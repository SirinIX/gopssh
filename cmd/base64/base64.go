package base64

import (
	"fmt"
	"gopssh/pkg/base64"

	"github.com/spf13/cobra"
)

type option struct {
	decode  bool
	content string
}

var op = &option{}

var Base64Cmd = &cobra.Command{
	Use: "base64",
	Short: "Encode or decode content with base64",
	Example: `  Encode: gopssh base64 -c 'root$123'
  Decode: gopssh base64 -d -c 'cm9vdA=='`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute(op)
	},
}

func init() {
	Base64Cmd.Flags().BoolVarP(&op.decode, "decode", "d", false, "decode or encode, default is encode")
	Base64Cmd.Flags().StringVarP(&op.content, "content", "c", "", "decode / encode data content (required)")

	Base64Cmd.MarkFlagRequired("content")
}

func execute(op *option) error {
	if op.decode {
		decStr, err := base64.Decode(op.content)
		if err != nil {
			return err
		}
		fmt.Printf("the base64 decoded result of %v is: %v", op.content, decStr)
	} else {
		encStr := base64.Encode(op.content)
		fmt.Printf("the base64 encoded result of %v is: %v", op.content, encStr)
	}
	return nil
}
