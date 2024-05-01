package database

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver.
)

// Config holds the configuration settings for the database.
type Config struct {
	DSN                   string        `envconfig:"DSN" default:"postgresql://dealer:password@localhost:5432/card_decks?sslmode=disable"`
	MaxOpenConnections    int           `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"10"`
	MaxIdleConnections    int           `envconfig:"DATABASE_MAX_IDLE_CONNECTIONS" default:"10"`
	ConnectionMaxLifetime time.Duration `envconfig:"DATABASE_CONNECTION_MAX_LIFETIME" default:"10m"`
}

// Connect returns a postgres database connection.
func Connect(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)

	return db, nil
}
