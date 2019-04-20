package config

import (
	"reflect"
)

// Config represents the configurations for the server
type Config struct {
	database Database
	grpcPort string `default:"50005"`
}

// Database returns the config database
func (c Config) Database() Database {
	return c.database
}

// GRPCPort returns the port to be used by the GRPC server
func (c Config) GRPCPort() string {
	return getDefaultValue(c, "grpcPort")
}

// Load returns the configuration for the server
func Load() Config {
	postgres := Postgres{}
	db := Database{
		postgres: postgres,
	}

	return Config{
		database: db,
	}
}

// Database represents the database configurations for the server
type Database struct {
	postgres Postgres
}

// Postgres returns the postgres db config
func (d Database) Postgres() Postgres {
	return d.postgres
}

func getDefaultValue(i interface{}, prop string) string {
	field, ok := reflect.TypeOf(i).FieldByName(prop)
	if !ok {
		return ""
	}
	return field.Tag.Get("default")
}
