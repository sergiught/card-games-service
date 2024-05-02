package deck

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	Shuffled bool         `json:"shuffled"`
	DeckType string       `json:"deck_type"`
	Cards    []FrenchCard `json:"cards"`
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

	s.log.Debug().Interface("request", request).Msg("received create deck request")

	deck, err := NewWithFrenchCards(request.DeckType, request.Shuffled, request.Cards)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to create a new french card deck")
		respond.NewResponse(w).DefaultMessage().InternalServerError(nil)
		return
	}

	if err := s.repository.CreateDeck(req.Context(), deck); err != nil {
		s.log.Error().Err(err).Msg("failed to persist french card deck to the database")
		respond.NewResponse(w).DefaultMessage().InternalServerError(nil)
		return
	}

	response := CreateDeckResponse{
		DeckID:    deck.ID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}

	s.log.Debug().Interface("response", response).Msg("sending create deck response")

	respond.NewResponse(w).DefaultMessage().Ok(response)
}

// OpenDeck handles the creation of a new card deck.
func (s *Service) OpenDeck(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	deckID := params.ByName("id")

	deck, err := s.repository.OpenDeck(req.Context(), deckID)
	if err != nil {
		s.log.Error().Err(err).Str("deck_id", deckID).Msg("failed to open a french card deck")
		respond.NewResponse(w).DefaultMessage().NotFound(nil)
		return
	}

	respond.NewResponse(w).DefaultMessage().Ok(deck)
}

// DrawCardsRequest represents the parameters
// required to draw cards from a deck.
type DrawCardsRequest struct {
	Cards int `json:"cards"`
}

// DrawCardsResponse represents the response
// returned after drawing cards from a deck.
type DrawCardsResponse struct {
	Cards []FrenchCard `json:"cards"`
}

// DrawCards draws N cards from a given card deck.
func (s *Service) DrawCards(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	deckID := params.ByName("id")

	var request DrawCardsRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		s.log.Error().Err(err).Msg("failed to decode request body")
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}

	drawnCards, err := s.repository.DrawCards(req.Context(), deckID, request.Cards)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to draw french cards from deck")
		respond.NewResponse(w).DefaultMessage().InternalServerError(nil)
		return
	}

	response := DrawCardsResponse{
		Cards: drawnCards,
	}

	s.log.Debug().Interface("response", response).Msg("sending draw cards response")

	respond.NewResponse(w).DefaultMessage().Ok(response)
}
