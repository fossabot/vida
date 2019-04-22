package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/gangachris/vida/meta"
	"github.com/gangachris/vida/models"
)

func init() {
	rootCmd.AddCommand(mediaCmd)

	mediaCmd.AddCommand(mediaSearchCmd)

	mediaSearchCmd.Flags().String("type", "", "define the type of media to search for: (movie, series)")
	mediaSearchCmd.Flags().String("dir", "", "the directory in which your media files exist")

	_ = mediaSearchCmd.MarkFlagRequired("type")
	_ = mediaSearchCmd.MarkFlagRequired("dir")
}

var mediaCmd = &cobra.Command{
	Use:   "media",
	Short: "Actions to handle media within vida",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("not implemented") // should be generated at build time
	},
}

// media actions
var mediaSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for media",
	Run: func(cmd *cobra.Command, args []string) {
		mediaType, err := cmd.Flags().GetString("type")
		if err != nil || mediaType == "" {
			exit(err)
		}

		dir, err := cmd.Flags().GetString("dir")
		if err != nil || dir == "" {
			exit(err)
		}

		switch mediaType {
		case "movies":
			movieCh := make(chan models.Movie)
			errCh := make(chan error)
			doneCh := make(chan struct{})

			go meta.SearchMovies(dir, nil, doneCh, errCh)

			for {
				select {
				case movie := <-movieCh:
					color.Yellow("%+v/n", movie)
				case err := <-errCh:
					if err != nil {
						exit(err)
					}
				case <-doneCh:
					color.Green("Movies search complete. check the database to see metadata or call the grpc cli")
					return
				}
			}
		case "series":
			exit(ErrNotImplemented)
		}
	},
}
