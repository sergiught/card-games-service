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

// RegisterCreateDeckSteps registers steps for the CreateFrenchDeck.feature scenarios.
func RegisterCreateDeckSteps(ctx *godog.ScenarioContext, deckCtx *DeckContext) {
	ctx.Step(
		`^I create a "(standard|custom)" and "(shuffled|sorted)" deck of French cards$`,
		deckCtx.iCreateADeckOfFrenchCards,
	)
	ctx.Step(
		`^I create a "(standard|custom)" and "(shuffled|sorted)" deck of French cards with the following cards:$`,
		deckCtx.iCreateADeckOfFrenchCardsWithCards,
	)
	ctx.Step(
		`^I should see that a deck was created with "(\d+)" "(shuffled|sorted)" cards$`,
		deckCtx.iShouldSeeThatADeckWasCreatedWithCards,
	)
	ctx.Step(
		`^the deck should have the following cards in "(shuffled|sorted)" order:$`,
		deckCtx.theDeckShouldHaveTheFollowingCardsInOrder,
	)
}

func (deckCtx *DeckContext) iCreateADeckOfFrenchCards(ctx context.Context, deckType, deckOrder string) error {
	shuffled := "false"
	if deckOrder == "shuffled" {
		shuffled = "true"
	}

	payload := []byte(fmt.Sprintf(`{"deck_type":%q,"shuffled":%s}`, deckType, shuffled))

	deckCtx.rawResponse = deckCtx.sendRequest(ctx, http.MethodPost, "/decks", payload)

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

	deckCtx.rawResponse = deckCtx.sendRequest(ctx, http.MethodPost, "/decks", payload)

	return nil
}

func (deckCtx *DeckContext) iShouldSeeThatADeckWasCreatedWithCards(
	ctx context.Context,
	cardsCount int,
	deckOrder string,
) error {
	assert.Equal(godog.T(ctx), http.StatusOK, deckCtx.rawResponse.StatusCode)

	err := json.NewDecoder(deckCtx.rawResponse.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	shuffled := false
	if deckOrder == "shuffled" {
		shuffled = true
	}

	require.NotEmpty(godog.T(ctx), deckCtx.response)
	assert.NotEmpty(godog.T(ctx), deckCtx.response.DeckID)
	assert.Equal(godog.T(ctx), cardsCount, deckCtx.response.Remaining)
	assert.Equal(godog.T(ctx), shuffled, deckCtx.response.Shuffled)

	return nil
}

func (deckCtx *DeckContext) theDeckShouldHaveTheFollowingCardsInOrder(
	ctx context.Context,
	deckOrder string,
	cardsTable *godog.Table,
) error {
	query := `SELECT cards FROM card_decks WHERE deck_id = $1`

	row := deckCtx.database.QueryRowContext(ctx, query, deckCtx.response.DeckID)

	var cardsJSON string

	err := row.Scan(&cardsJSON)
	require.NoError(godog.T(ctx), err)

	var cards []deck.FrenchCard
	err = json.Unmarshal([]byte(cardsJSON), &cards)
	require.NoError(godog.T(ctx), err)

	cardsFromTable := deckCtx.parseCardsTable(ctx, cardsTable)

	if deckOrder == "sorted" {
		assert.Equal(godog.T(ctx), cardsFromTable, cards)
		return nil
	}

	assert.ElementsMatch(godog.T(ctx), cards, cardsFromTable)

	return nil
}
