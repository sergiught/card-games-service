package deck

import (
	"encoding/json"
	"net/http"

	"github.com/nicklaw5/go-respond"
	"github.com/rs/zerolog"
)

// Service for managing card decks.
type Service struct {
	log        zerolog.Logger
	repository RepositoryOperator
}

// NewService returns a new instance of Service.
func NewService(log zerolog.Logger, repository RepositoryOperator) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

// CreateDeckRequest represents the parameters
// required to create a new deck of cards.
type CreateDeckRequest struct {
	Shuffled bool   `json:"shuffled"`
	DeckType string `json:"deck_type"`
}

// CreateDeckResponse represents the response
// returned after creating a new deck of cards.
type CreateDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// CreateDeck handles the creation of a new card deck.
func (s *Service) CreateDeck(w http.ResponseWriter, req *http.Request) {
	var request CreateDeckRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		s.log.Error().Err(err).Msg("failed to decode request body")
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}

	deck, err := NewWithFrenchCards(request.DeckType, request.Shuffled)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to create a new french card deck")
		respond.NewResponse(w).DefaultMessage().InternalServerError(nil)
	}

	if err := s.repository.CreateDeck(req.Context(), deck); err != nil {
		s.log.Error().Err(err).Msg("failed to persist french card deck to the database")
		respond.NewResponse(w).DefaultMessage().InternalServerError(nil)
	}

	response := CreateDeckResponse{
		DeckID:    deck.ID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}

	respond.NewResponse(w).DefaultMessage().Ok(response)
}
