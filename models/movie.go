package models

import (
	"context"

	"github.com/pkg/errors"

	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/entities"
)

// Movie represents a movie
type Movie struct {
	ID          int64  `json:"id,omitempty"` // struct tags
	IMDBID      string `json:"imdb_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Synopsis    string `json:"synopsis,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	TrailerURL  string `json:"trailer_url,omitempty"`
	PlaybackURI string `json:"playback_uri,omitempty"`
	Starring    string `json:"starring,omitempty"`
	Duration    string `json:"duration,omitempty"`
	Year        int64  `json:"year,omitempty"`
	ReleaseDate int64  `json:"release_date,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	IMDBJSON    string `json:"imdbjson,omitempty"`
	Search      string `json:"search"`
}

// Store save a movie to the storage
func (m *Movie) Store(ctx context.Context, store *db.Storage) error {
	movie := entities.Movie(*m)
	return store.MovieStore.Store(ctx, &movie)
}

// FindMovieBySearchTerm will search the db for a search term specified
func FindMovieBySearchTerm(ctx context.Context, store *db.Storage, search string) (*Movie, error) {
	em, err := store.MovieStore.FindMovieBySearchTerm(ctx, search)
	if err != nil {
		return nil, errors.Wrap(err, "could not get movie by search term")
	}
	if em == nil {
		return nil, nil
	}
	movie := Movie(*em)
	return &movie, err
}

// AllMovies returns all movies in th e database
func AllMovies(ctx context.Context, store *db.Storage) ([]Movie, error) {
	em, err := store.MovieStore.All(ctx)
	if err != nil {
		return nil, err
	}

	movies := make([]Movie, len(em))
	for idx, movie := range em {
		movies[idx] = Movie(movie)
	}
	return movies, nil
}
