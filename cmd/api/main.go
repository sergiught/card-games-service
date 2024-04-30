package main

import (
	"log"

	"github.com/sergiught/card-games-service/internal/config"
	"github.com/sergiught/card-games-service/internal/server"
)

func main() {
	configuration, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	httpServer := server.New(configuration.Server, nil)

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
