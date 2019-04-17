package config

import "fmt"

// Postgres represents the postgres config for the server
type Postgres struct {
	host              string `default:"localhost"`
	port              string `default:"5432"`
	database          string `default:"vida"`
	user              string `default:"vida"`
	password          string `default:"vida"`
	sslmode           string `default:"disable"`
	connectionTimeout int64
	sslcert           string
	sslkey            string
	sslrootcert       string
	drivername        string `default:"postgres"`
	schema            string `default:"vida"`
}

// DSN returns the Postgres db connection dsn
func (p Postgres) DSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		p.DriverName(),
		p.User(),
		p.Password(),
		p.Host(),
		p.Port(),
		p.Database(),
		p.SSLMode(),
	)
}

// Host returns postgres host
func (p Postgres) Host() string {
	if p.host != "" {
		return p.host
	}
	return getDefaultValue(p, "host")
}

// Port returns postgres port
func (p Postgres) Port() string {
	if p.port != "" {
		return p.port
	}
	return getDefaultValue(p, "port")
}

// Database returns postgres database
func (p Postgres) Database() string {
	if p.database != "" {
		return p.database
	}
	return getDefaultValue(p, "database")
}

// User returns postgres user
func (p Postgres) User() string {
	if p.user != "" {
		return p.user
	}
	return getDefaultValue(p, "user")
}

// Password returns postgres password
func (p Postgres) Password() string {
	if p.password != "" {
		return p.password
	}
	return getDefaultValue(p, "password")
}

// SSLMode returns postgres sslmode
func (p Postgres) SSLMode() string {
	if p.sslmode != "" {
		return p.sslmode
	}
	return getDefaultValue(p, "sslmode")
}

func (p Postgres) DriverName() string {
	if p.drivername != "" {
		return p.drivername
	}
	return getDefaultValue(p, "drivername")
}

func (p Postgres) Schema() string {
	if p.schema != "" {
		return p.schema
	}
	return getDefaultValue(p, "schema")
}
