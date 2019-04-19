package cmd

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/gangachris/vida/config"
	"github.com/gangachris/vida/meta"
	"github.com/gangachris/vida/storage"
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
			if err := searchMovies(dir); err != nil {
				exit(errors.Wrap(err, "could not search for movies"))
			}
		case "series":
			exit(ErrNotImplemented)
		}

		// walk dir and recursively search

		//suggestion, err := meta.SearchIMDB(search)
		//if err != nil {
		//	exit(err)
		//}
		//
		//movie, err := suggestion.ToMovie()
		//if err != nil {
		//	exit(err)
		//}

		//color.Blue("%+v", "movie") // should be generated at build time
	},
}

func searchMovies(dir string) error {
	cfg := config.Load()
	store, err := storage.NewPostgres(cfg)
	if err != nil {
		exit(errors.Wrap(err, "could not initialize the data store"))
	}
	metaHelper := meta.Meta{Client: http.DefaultClient}

	absolutePath, err := filepath.Abs(dir)
	if err != nil {
		exit(errors.Wrap(err, fmt.Sprint("could not get absolute path of %q"+dir)))
	}
	return godirwalk.Walk(absolutePath, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, directoryEntry *godirwalk.Dirent) error {
			// get the last part with the name of the file without the .mp4 extension
			_, file := filepath.Split(osPathname)
			if file == strings.TrimPrefix(dir, "./") {
				return nil
			}
			search := strings.TrimSuffix(file, ".mp4") // TODO: should support other formats

			imdbSuggestion, err := metaHelper.SearchIMDB(search)
			if err != nil {
				return err
			}

			movie, err := imdbSuggestion.ToMovie(osPathname)
			if err != nil {
				return err
			}

			// this will be removed to a central location for the API
			if err := movie.Store(context.Background(), store); err != nil {
				color.Green("%+v", err)
				return nil
			}

			return nil
		},
		ErrorCallback: func(s string, e error) godirwalk.ErrorAction {
			//color.Red("%s", errors.Wrap(e, s))
			return godirwalk.SkipNode
		},
	})
}
