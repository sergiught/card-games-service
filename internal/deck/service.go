package deck

import (
	"net/http"

	"github.com/nicklaw5/go-respond"
)

// Service for managing card decks.
type Service struct{}

// NewService returns a new instance of Service.
func NewService() *Service {
	return &Service{}
}

// CreateDeck handles the creation of a new card deck.
func (s *Service) CreateDeck(writer http.ResponseWriter, request *http.Request) {
	respond.NewResponse(writer).DefaultMessage().Ok(nil)
}
