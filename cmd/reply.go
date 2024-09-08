package cmd

import (
	"github.com/spf13/cobra"
	"pidy/wx"
)

var replyCmd = &cobra.Command{
	Use: "reply",
	Run: func(cmd *cobra.Command, args []string) {
		wx.HandleMessage()
	},
}
