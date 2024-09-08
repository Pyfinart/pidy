package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "say greet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("你好，需要什么帮助\n输入--help以获取更多信息")
	},
}
