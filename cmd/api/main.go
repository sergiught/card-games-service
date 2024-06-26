package main

import (
	stdLog "log"

	"github.com/sergiught/card-games-service/internal/config"
	"github.com/sergiught/card-games-service/internal/database"
	"github.com/sergiught/card-games-service/internal/logger"
	"github.com/sergiught/card-games-service/internal/router"
	"github.com/sergiught/card-games-service/internal/server"
)

func main() {
	configuration, err := config.LoadFromEnv()
	if err != nil {
		stdLog.Fatalf("failed to load env vars: %v", err)
	}

	log := logger.New(configuration.Environment)

	db, err := database.Connect(configuration.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to the database")
	}

	httpRouter := router.New(log, db)

	httpServer := server.New(configuration.Server, log, httpRouter)

	if err := httpServer.Start(); err != nil {
		log.Fatal().Err(err).Msg("server failure")
	}
}
