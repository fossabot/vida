package meta

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"

	"github.com/gangachris/vida/config"
	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/models"
)

// SearchMovies is used to search for movies, save them in the database and pass the results to a channel
func SearchMovies(dir string, movieCh chan<- models.Movie, doneCh chan struct{}, errCh chan<- error) {
	defer func() {
		if movieCh != nil {
			close(movieCh)
		}
		close(doneCh)
		close(errCh)
	}()

	cfg := config.Load()
	store, err := db.NewPostgres(cfg)
	if err != nil {
		errCh <- errors.Wrap(err, "could not initialize the data store")
	}

	metaHelper := Meta{Client: http.DefaultClient}

	absolutePath, err := filepath.Abs(dir)
	if err != nil {
		errCh <- errors.Wrap(err, fmt.Sprint("could not get absolute path of %q"+dir))
	}
	err = godirwalk.Walk(absolutePath, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, directoryEntry *godirwalk.Dirent) error {
			ctx := context.Background()
			if directoryEntry.IsDir() {
				return nil
			}

			// get the last part with the name of the file without the .mp4 extension
			_, file := filepath.Split(osPathname)
			search := strings.TrimSuffix(file, ".mp4") // TODO: should support other formats

			// check if search term was already searched
			movie, err := models.FindMovieBySearchTerm(ctx, store, search)
			if err != nil {
				return err
			}
			if movie != nil {
				if movieCh != nil {
					movieCh <- *movie
				}
				return nil
			}

			imdbSuggestion, err := metaHelper.SearchIMDB(search)
			if err != nil {
				return err
			}

			mv, err := imdbSuggestion.ToMovie(osPathname)
			if err != nil {
				return err
			}

			// this will be removed to a central location for the API
			if err := mv.Store(ctx, store); err != nil { // TODO: get the result from the DB
				color.Green("%+v", err)
				return nil
			}

			if movieCh != nil {
				movieCh <- mv
			}

			return nil
		},
	})

	errCh <- err
	doneCh <- struct{}{}
}
