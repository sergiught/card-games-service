package main

import (
	"log"
	"time"

	"github.com/sergiught/card-games-service/internal/server"
)

func main() {
	configuration := server.Config{
		Address:         "0.0.0.0:8000",
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
		ShutdownTimeout: 5 * time.Second,
	}

	httpServer := server.New(configuration, nil)

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
