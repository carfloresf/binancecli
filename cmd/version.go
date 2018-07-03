package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of binance-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 0.01 by CarlosF")
	},
}
