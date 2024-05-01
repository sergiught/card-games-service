package scenario

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/internal/deck"
)

// Initialize steps for the CreateFrenchDeck.feature scenarios.
func Initialize(deckCtx *DeckContext) func(*godog.ScenarioContext) {
	return func(ctx *godog.ScenarioContext) {
		ctx.Step(
			`^I create a "(standard|custom)"(?: and "(shuffled)")? deck of French cards$`,
			deckCtx.iCreateADeckOfFrenchCards,
		)
		ctx.Step(
			`^I create a "(standard|custom)"(?: and "(shuffled)")? deck of French cards with the following cards in this order:$`,
			deckCtx.iCreateADeckOfFrenchCardsWithCards,
		)
		ctx.When(`^I open the deck$`, deckCtx.iOpenTheDeck)

		ctx.Then(`^I should receive "(\d+)" cards$`, deckCtx.iShouldReceiveNCards)
		ctx.Then(`^I should have the following cards in this order:$`, deckCtx.iShouldHaveTheFollowingCardsInThisOrder)
		ctx.Then(`^I should have the cards in a shuffled order$`, deckCtx.iShouldHaveTheCardsInAShuffledOrder)
	}
}

func (deckCtx *DeckContext) iCreateADeckOfFrenchCards(ctx context.Context, deckType, deckOrder string) error {
	shuffled := "false"
	if deckOrder == "shuffled" {
		shuffled = "true"
	}

	payload := []byte(fmt.Sprintf(`{"deck_type":%q,"shuffled":%s}`, deckType, shuffled))

	response := deckCtx.sendRequest(ctx, http.MethodPost, "/decks", payload)

	assert.Equal(godog.T(ctx), http.StatusOK, response.StatusCode)

	err := json.NewDecoder(response.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	return nil
}

func (deckCtx *DeckContext) iCreateADeckOfFrenchCardsWithCards(
	ctx context.Context,
	deckType string,
	deckOrder string,
	cardsTable *godog.Table,
) error {
	shuffled := "false"
	if deckOrder == "shuffled" {
		shuffled = "true"
	}

	cardsFromTable := deckCtx.parseCardsTable(ctx, cardsTable)

	cardsJSON, err := json.Marshal(cardsFromTable)
	require.NoError(godog.T(ctx), err)

	payload := []byte(fmt.Sprintf(`{"deck_type":%q,"shuffled":%s,"cards":%s}`, deckType, shuffled, cardsJSON))

	response := deckCtx.sendRequest(ctx, http.MethodPost, "/decks", payload)

	assert.Equal(godog.T(ctx), http.StatusOK, response.StatusCode)

	err = json.NewDecoder(response.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	return nil
}

func (deckCtx *DeckContext) iOpenTheDeck(ctx context.Context) error {
	require.NotEmpty(godog.T(ctx), deckCtx.response)
	assert.Contains(godog.T(ctx), deckCtx.response, "deck_id")

	deckID := deckCtx.response["deck_id"].(string)
	response := deckCtx.sendRequest(ctx, http.MethodGet, "/decks/"+deckID, nil)

	assert.Equal(godog.T(ctx), http.StatusOK, response.StatusCode)

	err := json.NewDecoder(response.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	deckCtx.checkCardsInResponse = true

	return nil
}

func (deckCtx *DeckContext) iShouldReceiveNCards(ctx context.Context, cardsCount float64) error {
	require.NotEmpty(godog.T(ctx), deckCtx.response)
	assert.Contains(godog.T(ctx), deckCtx.response, "deck_id")
	assert.Contains(godog.T(ctx), deckCtx.response, "shuffled")
	assert.Contains(godog.T(ctx), deckCtx.response, "remaining")
	assert.Equal(godog.T(ctx), cardsCount, deckCtx.response["remaining"])

	if deckCtx.checkCardsInResponse {
		assert.Contains(godog.T(ctx), deckCtx.response, "cards")
	}

	return nil
}

func (deckCtx *DeckContext) iShouldHaveTheFollowingCardsInThisOrder(ctx context.Context, cardsTable *godog.Table) error {
	query := `SELECT cards FROM card_decks WHERE deck_id = $1`

	row := deckCtx.database.QueryRowContext(ctx, query, deckCtx.response["deck_id"])

	var cardsJSON string

	err := row.Scan(&cardsJSON)
	require.NoError(godog.T(ctx), err)

	var cards []deck.FrenchCard
	err = json.Unmarshal([]byte(cardsJSON), &cards)
	require.NoError(godog.T(ctx), err)

	cardsFromTable := deckCtx.parseCardsTable(ctx, cardsTable)

	assert.Equal(godog.T(ctx), cardsFromTable, cards)

	return nil
}

func (deckCtx *DeckContext) iShouldHaveTheCardsInAShuffledOrder(ctx context.Context) error {
	query := `SELECT cards FROM card_decks WHERE deck_id = $1`

	row := deckCtx.database.QueryRowContext(ctx, query, deckCtx.response["deck_id"])

	var cardsJSON string

	err := row.Scan(&cardsJSON)
	require.NoError(godog.T(ctx), err)

	var cards []deck.FrenchCard
	err = json.Unmarshal([]byte(cardsJSON), &cards)
	require.NoError(godog.T(ctx), err)

	standardCardsInDeck := deck.GenerateStandardFrenchCardsForDeck()

	assert.NotEqual(godog.T(ctx), standardCardsInDeck, cards)

	return nil
}
