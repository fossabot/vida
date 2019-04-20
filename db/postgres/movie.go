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

// All returns all movies in the database
func (m MovieStore) All(ctx context.Context) (movies []entities.Movie, err error) {
	query := `
		SELECT
			id, imdb_id, title, synopsis, image_url, trailer_url, playback_uri, duration, year
		FROM ` + moviesTableName

	err = m.client.Select(&movies, query)
	return
}

// IMDBJSONExists checks wether the potential json to be returned was already searched/request already sent
func (m MovieStore) IMDBJSONExists(ctx context.Context, search string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM ` + moviesTableName + `
		WHERE
			imdb_json ->> 'q' = $1
	`
	var count int
	if err := m.client.QueryRow(query, search).Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
