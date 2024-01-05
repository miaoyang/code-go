package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var HelpCmd = &cobra.Command{
	Use:     "help",
	Long:    "Print help info",
	Example: "code-go help -h",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	fmt.Println("-----------查看help信息-----------")
}

func init() {
	HelpCmd.Flags().StringP("help", "h", "help", "print help info")

}
