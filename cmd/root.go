package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// ErrNotImplemented is used to show that a method/functionality is not yet implemented
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

// exit is used to exit the application
func exit(err error) {
	color.Red("%v", err)
	os.Exit(1)
}
