package router

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"

	"github.com/sergiught/card-games-service/internal/deck"
)

// New instantiates a new http router and
// configures the endpoints of the service.
func New(log zerolog.Logger, db *sql.DB) http.Handler {
	router := httprouter.New()

	deckRepository := deck.NewRepository(db)
	deckService := deck.NewService(log, deckRepository)

	router.HandlerFunc(http.MethodPost, "/decks", deckService.CreateDeck)

	return router
}
