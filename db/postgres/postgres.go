package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/gangachris/vida/config"
)

// NewClient creates a new instance of the db
func NewClient(cfg config.Config) (*sqlx.DB, error) {
	pg := cfg.Database().Postgres()
	return sqlx.Open(pg.DriverName(), pg.DSN())
}
