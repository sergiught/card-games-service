package deck

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Deck represents a card deck.
type Deck struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []Card    `json:"-"`
}

// Card is the interface that all card types need to implement.
type Card interface{}

// MarshalJSON implements the json.Marshaler interface.
func (d *Deck) MarshalJSON() ([]byte, error) {
	type deck Deck
	type deckWrapper struct {
		*deck
		RawCards []json.RawMessage `json:"cards"`
	}

	wrapper := deckWrapper{(*deck)(d), nil}
	wrapper.RawCards = make([]json.RawMessage, len(d.Cards))

	for index, card := range d.Cards {
		bytes, err := json.Marshal(card)
		if err != nil {
			return nil, err
		}

		wrapper.RawCards[index] = bytes
	}

	return json.Marshal(wrapper)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Deck) UnmarshalJSON(data []byte) error {
	type deck Deck
	type deckWrapper struct {
		*deck
		RawCards []json.RawMessage `json:"cards"`
	}

	wrapper := &deckWrapper{(*deck)(d), nil}

	if err := json.Unmarshal(data, wrapper); err != nil {
		return err
	}

	if wrapper.RawCards == nil {
		return nil
	}

	d.Cards = make([]Card, len(wrapper.RawCards))

	for index, rawCard := range wrapper.RawCards {
		var frenchCard FrenchCard
		if err := json.Unmarshal(rawCard, &frenchCard); err == nil {
			d.Cards[index] = frenchCard
			continue
		}
	}

	return nil
}

// NewWithFrenchCards creates a new deck of cards based on the specified deck type and shuffle option.
func NewWithFrenchCards(deckType string, shuffled bool, customCards []FrenchCard) (*Deck, error) {
	var cards []Card

	switch deckType {
	case "custom":
		cards = make([]Card, len(customCards))
		for index, card := range customCards {
			cards[index] = Card(card)
		}
	case "standard":
		cards = GenerateStandardFrenchCardsForDeck()
	default:
		cards = []Card{}
	}

	if shuffled {
		err := ShuffleCards(cards)
		if err != nil {
			return nil, err
		}
	}

	return &Deck{
		ID:        uuid.New(),
		Remaining: len(cards),
		Shuffled:  shuffled,
		Cards:     cards,
	}, nil
}
