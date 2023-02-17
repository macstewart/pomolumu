package cmd

import (
	"os"

	"github.com/macstewart/pomolumu/ui"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "pomolumu",
	Run: func(cmd *cobra.Command, args []string) { ui.Run(args) },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
