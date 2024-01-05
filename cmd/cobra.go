package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "code-go",
	Short:        "code-go",
	SilenceUsage: true,
	Long:         `code-go`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("执行出错")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Println("-----------欢迎使用cobra-----------")
	help := "code-go -h or code-go -help 查看帮助信息"
	fmt.Println(help)
}

func init() {
	rootCmd.AddCommand(HelpCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
