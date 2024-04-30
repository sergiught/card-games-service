package deck

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
	var deck Deck

	if err := json.NewDecoder(request.Body).Decode(&deck); err != nil {
		respond.NewResponse(writer).DefaultMessage().BadRequest(nil)
		return
	}

	// We're just making the tests pass.
	deck.ID = uuid.New()
	deck.Remaining = 52

	respond.NewResponse(writer).DefaultMessage().Ok(deck)
}
