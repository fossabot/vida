package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gangachris/vida/models"
)

type MovieStore struct {
	Client sqlx.DB
}

func (m *MovieStore) Store(ctx context.Context, movie models.Movie) error {
	// some insert statement will go here
	return nil
}