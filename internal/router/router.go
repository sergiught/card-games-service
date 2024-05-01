package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"

	"github.com/sergiught/card-games-service/internal/deck"
)

// New instantiates a new http router and
// configures the endpoints of the service.
func New(log zerolog.Logger) http.Handler {
	router := httprouter.New()

	deckService := deck.NewService(log)

	router.HandlerFunc(http.MethodPost, "/decks", deckService.CreateDeck)

	return router
}
