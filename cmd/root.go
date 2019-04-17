package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	ErrNotImplemented = errors.New("Not Implemented Yet")
)

var rootCmd = &cobra.Command{
	Use:   "vida",
	Short: "vida is a media player server",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("❓ vida help will print here 😇")
	},
}

// Execute executes the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("%v", err)
		os.Exit(1)
	}
}
