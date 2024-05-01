package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFrenchCard_FromCode(t *testing.T) {
	var testCases = []struct {
		name          string
		givenCode     string
		expected      *FrenchCard
		expectedError error
	}{
		{
			name:      "valid card code: Ace of Spades",
			givenCode: "AS",
			expected:  &FrenchCard{Value: "ACE", Suit: "SPADES", Code: "AS"},
		},
		{
			name:      "valid card code: 10 of Hearts",
			givenCode: "10H",
			expected:  &FrenchCard{Value: "10", Suit: "HEARTS", Code: "10H"},
		},
		{
			name:          "invalid card code: too short",
			givenCode:     "A",
			expected:      &FrenchCard{},
			expectedError: ErrInvalidCardCode,
		},
		{
			name:          "invalid card code: too long",
			givenCode:     "AAAAA",
			expected:      &FrenchCard{},
			expectedError: ErrInvalidCardCode,
		},
		{
			name:          "invalid card code: wrong value",
			givenCode:     "1S",
			expected:      &FrenchCard{},
			expectedError: ErrInvalidValueCode,
		},
		{
			name:          "invalid card code: wrong suit",
			givenCode:     "AK",
			expected:      &FrenchCard{},
			expectedError: ErrInvalidSuitCode,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			card := &FrenchCard{}
			err := card.FromCode(test.givenCode)

			if test.expectedError != nil {
				assert.ErrorIs(t, err, test.expectedError)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.expected, card)
		})
	}
}

func TestGenerateStandardFrenchCardsForDeck(t *testing.T) {
	cards := GenerateStandardFrenchCardsForDeck()
	expectedCount := 52 // 13 values across 4 suits.

	assert.Len(t, cards, expectedCount, "should generate 52 cards")

	// Test some specific cards to ensure correct order and values.
	expectedFirst := FrenchCard{Value: "ACE", Suit: "SPADES", Code: "AS"}
	expectedLast := FrenchCard{Value: "KING", Suit: "HEARTS", Code: "KH"}

	assert.Equal(t, expectedFirst, cards[0], "the first card should be the Ace of Spades")
	assert.Equal(t, expectedLast, cards[expectedCount-1], "the last card should be the King of Hearts")
}

func TestShuffleCards(t *testing.T) {
	initialDeck := GenerateStandardFrenchCardsForDeck()
	copiedDeck := make([]Card, len(initialDeck))
	copy(copiedDeck, initialDeck)

	err := ShuffleCards(copiedDeck)
	require.NoError(t, err)

	t.Run("test if all cards are still present", func(t *testing.T) {
		require.Len(t, copiedDeck, len(initialDeck), "shuffled deck must have the same number of cards")

		// Create maps to count occurrences of each
		// card in original and shuffled decks.
		originalCardMap := make(map[string]int)
		shuffledCardMap := make(map[string]int)

		for i, card := range initialDeck {
			originalCardMap[card.(FrenchCard).Code]++
			shuffledCardMap[copiedDeck[i].(FrenchCard).Code]++
		}

		for code, count := range originalCardMap {
			require.Equal(t, count, shuffledCardMap[code], fmt.Sprintf("card %s should appear the same number of times in both decks", code))
		}
	})

	t.Run("test if order has changed", func(t *testing.T) {
		// It's possible but highly unlikely that a shuffle results in the same order.
		// Thus, we assume the test is valid if we detect any position where the cards differ.
		sameOrder := true
		for i, card := range initialDeck {
			if card.(FrenchCard).Code != copiedDeck[i].(FrenchCard).Code {
				sameOrder = false
				break
			}
		}
		assert.False(t, sameOrder, "the order of cards should be different after shuffling")
	})
}
