package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/sergiught/card-games-service/internal/deck"
)

// New instantiates a new http router and
// configures the endpoints of the service.
func New() http.Handler {
	router := httprouter.New()

	deckService := deck.NewService()

	router.HandlerFunc(http.MethodPost, "/decks", deckService.CreateDeck)

	return router
}
