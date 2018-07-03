package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "binance-cli",
	Short: "binance-cli is a command line app to access binance data (more like a test for me)",
	Long:  "Short, fast application for getting public information from Binance API built with spf13/cobra as an exercise by CarlosF.",
}
