package deck

import (
	"github.com/google/uuid"
)

// Deck represents a card deck.
type Deck struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}
