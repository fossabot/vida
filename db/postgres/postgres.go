package postgres

import (
	"github.com/gangachris/vida/config"

	"github.com/jmoiron/sqlx"

	// implements db/sql
	_ "github.com/lib/pq"
)

// NewClient returns a new Psotgres client
func NewClient(cfg config.Config) (*sqlx.DB, error) {
	return sqlx.Connect(cfg.Database().Postgres().DriverName(), cfg.Database().Postgres().DSN())
}
