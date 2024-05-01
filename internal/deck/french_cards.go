package deck

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

var (
	frenchSuitMap = map[string]string{
		"S": "SPADES",
		"H": "HEARTS",
		"D": "DIAMONDS",
		"C": "CLUBS",
	}
	frenchValueMap = map[string]string{
		"A":  "ACE",
		"2":  "2",
		"3":  "3",
		"4":  "4",
		"5":  "5",
		"6":  "6",
		"7":  "7",
		"8":  "8",
		"9":  "9",
		"10": "10",
		"J":  "JACK",
		"Q":  "QUEEN",
		"K":  "KING",
	}

	// ErrInvalidCardCode is an error that indicates that the card code
	// provided to the system does not conform to the expected format.
	ErrInvalidCardCode = errors.New("invalid card code")

	// ErrInvalidSuitCode is an error that signifies a non-matching
	// suit code was provided during card creation or processing.
	ErrInvalidSuitCode = errors.New("invalid suit code")

	// ErrInvalidValueCode is an error indicating that the value code portion
	// of a card code does not match any of the standard card values.
	ErrInvalidValueCode = errors.New("invalid value code")
)

// FrenchCard is a type that implements the Card interface.
type FrenchCard struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// FromCode uses maps to determine the full value and suit from a card code.
func (c *FrenchCard) FromCode(cardCode string) error {
	if len(cardCode) < 2 || len(cardCode) > 3 {
		return fmt.Errorf("%w: must have 2 or 3 letters", ErrInvalidCardCode)
	}

	// Extract value and suit from the code.
	// Takes into account the fact that 10 of a suit is 3 letters.
	valueCode := cardCode[:len(cardCode)-1]
	suitCode := cardCode[len(cardCode)-1:]

	var ok bool
	if c.Value, ok = frenchValueMap[valueCode]; !ok {
		return fmt.Errorf("%w: %s", ErrInvalidValueCode, valueCode)
	}
	if c.Suit, ok = frenchSuitMap[suitCode]; !ok {
		return fmt.Errorf("%w: %s", ErrInvalidSuitCode, valueCode)
	}

	c.Code = cardCode

	return nil
}

// GenerateStandardFrenchCardsForDeck creates a complete deck of French playing cards in a
// sequential order: A-spades, 2-spades, 3-spades... Followed by diamonds, clubs, then hearts.
func GenerateStandardFrenchCardsForDeck() []Card {
	var cards []Card

	for _, suitCode := range []string{"S", "D", "C", "H"} {
		suit := frenchSuitMap[suitCode]

		for _, valueCode := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			value := frenchValueMap[valueCode]

			card := FrenchCard{
				Value: value,
				Suit:  suit,
				Code:  fmt.Sprintf("%s%s", valueCode, suitCode),
			}

			cards = append(cards, card)
		}
	}

	return cards
}

// ShuffleCards takes a slice of cards and shuffles
// it in place using cryptographic randomness.
func ShuffleCards(cards []Card) error {
	for i := range cards {
		bi, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := bi.Int64()

		cards[i], cards[j] = cards[j], cards[i]
	}
	return nil
}
