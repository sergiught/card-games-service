package deck

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeck_JSON(t *testing.T) {
	deckID, err := uuid.Parse("718a08ff-d1a2-47ea-9168-b949d11d4545")
	require.NoError(t, err)

	var testCases = []struct {
		deck Deck
		json string
	}{
		{
			deck: Deck{
				ID:        deckID,
				Shuffled:  true,
				Remaining: 3,
				Cards: []Card{
					FrenchCard{
						Value: "ACE",
						Suit:  "SPADES",
						Code:  "AS",
					},
					FrenchCard{
						Value: "KING",
						Suit:  "HEARTS",
						Code:  "KH",
					},
					FrenchCard{
						Value: "8",
						Suit:  "CLUBS",
						Code:  "8C",
					},
				},
			},
			json: `{"deck_id":"718a08ff-d1a2-47ea-9168-b949d11d4545","shuffled":true,"remaining":3,"cards":[{"value":"ACE","suit":"SPADES","code":"AS"},{"value":"KING","suit":"HEARTS","code":"KH"},{"value":"8","suit":"CLUBS","code":"8C"}]}`,
		},
		{
			deck: Deck{
				ID:        deckID,
				Shuffled:  false,
				Remaining: 0,
				Cards:     []Card{},
			},
			json: `{"deck_id":"718a08ff-d1a2-47ea-9168-b949d11d4545","shuffled":false,"remaining":0,"cards":[]}`,
		},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			t.Run("json encoding", func(t *testing.T) {
				payload, err := json.Marshal(&test.deck)
				require.NoError(t, err)
				assert.Equal(t, test.json, string(payload))
			})

			t.Run("json decoding", func(t *testing.T) {
				var deck Deck
				err = json.Unmarshal([]byte(test.json), &deck)
				require.NoError(t, err)
				assert.Equal(t, test.deck, deck)
			})
		})
	}
}
