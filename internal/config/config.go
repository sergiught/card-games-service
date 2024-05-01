package config

import (
	_ "github.com/joho/godotenv/autoload" // Autoload env vars from a .env file.
	"github.com/kelseyhightower/envconfig"

	"github.com/sergiught/card-games-service/internal/database"
	"github.com/sergiught/card-games-service/internal/server"
)

// Specification contains all the config
// parameters that this service uses.
type Specification struct {
	Environment string          `envconfig:"ENVIRONMENT" default:"development"`
	Server      server.Config   `envconfig:"SERVER"`
	Database    database.Config `envconfig:"DATABASE"`
}

// LoadFromEnv will load the env vars from the OS.
func LoadFromEnv() (*Specification, error) {
	spec := &Specification{}
	err := envconfig.Process("", spec)
	return spec, err
}
