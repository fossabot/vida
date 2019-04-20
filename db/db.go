package db

import (
	"context"

	"github.com/gangachris/vida/config"
	"github.com/gangachris/vida/db/postgres"
	"github.com/gangachris/vida/entities"
)

// Storage represents vida's storage struct
type Storage struct {
	MovieStore MovieStore
}

// NewPostgres creates a new postgres storage
func NewPostgres(cfg config.Config) (*Storage, error) {
	client, err := postgres.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	storage := &Storage{
		MovieStore: postgres.NewMovieStore(client),
	}
	return storage, nil
}

// MovieStore interface represents a contract for all movie storage operations
type MovieStore interface {
	Store(ctx context.Context, movie *entities.Movie) error
	IMDBJSONExists(ctx context.Context, search string) (bool, error)
}
