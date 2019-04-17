package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Version = "unset"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the currently installed version of vida",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue(Version) // should be generated at build time
	},
}
