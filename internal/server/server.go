package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type (
	// Server encapsulates the details of an HTTP server.
	Server struct {
		config  Config
		handler http.Handler
		logger  zerolog.Logger
	}

	// Config holds the configuration settings for the HTTP Server.
	Config struct {
		Address         string        `envconfig:"ADDRESS" default:"0.0.0.0:8000"`
		ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s"`
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"5s"`
	}
)

// New returns a new HTTP Server.
func New(config Config, log zerolog.Logger, handler http.Handler) *Server {
	return &Server{
		config:  config,
		handler: handler,
		logger:  log,
	}
}

// Start will start the HTTP Server and will handle shutdowns gracefully.
func (s *Server) Start() error {
	// Make a channel to listen for an interrupt or a terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := &http.Server{
		Addr:         s.config.Address,
		Handler:      s.handler,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}

	// Make a channel to listen for errors coming from the listener.
	// Use a buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		s.logger.Info().Msgf("server listening on port %q", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Block and wait for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server encountered an error: %w", err)
	case sig := <-shutdown:
		s.logger.Info().Msgf("server shutting down after receiving %+v", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		// Ask listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			_ = api.Close()
			return fmt.Errorf("server failed to shutdown gracefully: %w", err)
		}
	}

	return nil
}
