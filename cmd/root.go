package cmd

import (
	"fmt"
	"gopssh/cmd/base64"
	"gopssh/cmd/check"
	"gopssh/cmd/convert"
	"gopssh/cmd/execute"
	"gopssh/cmd/template"
	"gopssh/cmd/upload"
	"gopssh/cmd/version"
	"gopssh/log"

	"github.com/spf13/cobra"
)

const (
	logo = `
                          / 
    _,  __  ,_   (   (   /_ 
   (_)_(_)_/|_)_/_)_/_)_/ /_
    /|     /|               
   (/     (/                 
`
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A high-performance and user friendly ssh tool",
	// Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Welcome to gopssh!" + logo + "\nPlease type `gopssh --help` to see the usage")
	},
}

func init() {
	log.InitLogger()

	rootCmd.AddCommand(
		base64.Base64Cmd,
		check.CheckCmd,
		convert.ConvertCmd,
		execute.ExecuteCmd,
		template.TemplateCmd,
		upload.UploadCmd,
		version.VersionCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
