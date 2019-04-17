package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mediaCmd)

	mediaCmd.AddCommand(mediaSearchCmd)
}

var mediaCmd = &cobra.Command{
	Use:   "media",
	Short: "actions to handle media within vida",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("not implemented") // should be generated at build time
	},
}

// media actions
var mediaSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for media",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("not implemented") // should be generated at build time
	},
}
