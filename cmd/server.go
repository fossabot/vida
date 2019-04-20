package cmd

import (
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/gangachris/vida/config"
	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/server/grpc"
)

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.AddCommand(grpcCmd)
}

var apiCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the available servers in vida (grpc)",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("not implemented") // should be generated at build time
	},
}

// media actions
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run the grpc server (default port is 50005)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		store, err := db.NewPostgres(cfg)
		if err != nil {
			exit(errors.Wrap(err, "could not instantiate storage"))
		}

		grpcServer := grpc.NewServer(cfg, store)

		color.Blue("GRPC Server listening on %s:%s", "localhost", cfg.GRPCPort())
		if err := grpcServer.Start(); err != nil {
			exit(errors.Wrap(err, "could not start grpc server"))
		}
	},
}
