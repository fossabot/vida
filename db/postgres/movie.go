package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gangachris/vida/entities"
)

const moviesTableName = "movies"

// MovieStore represents a postgres movie store
type MovieStore struct {
	client *sqlx.DB
}

// NewMovieStore instantiates a postgres movie store
func NewMovieStore(client *sqlx.DB) *MovieStore {
	return &MovieStore{client}
}

// Store implements the store interface for storing movies
func (m MovieStore) Store(ctx context.Context, movie *entities.Movie) error {
	query := `
	INSERT INTO ` + moviesTableName + `
		(imdb_id, title, synopsis, image_url, trailer_url, playback_uri, duration, year, imdb_json)
	VALUES 
		(:imdb_id, :title, :synopsis, :image_url, :trailer_url, :playback_uri, :duration, :year, :imdb_json)`

	_, err := m.client.NamedExec(query, movie)
	return err
}
