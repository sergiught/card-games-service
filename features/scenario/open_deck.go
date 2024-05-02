package scenario

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RegisterOpenDeckSteps registers steps for the OpenFrenchDeck.feature scenarios.
func RegisterOpenDeckSteps(ctx *godog.ScenarioContext, deckCtx *DeckContext) {
	ctx.Step(
		`^I open the deck$`,
		deckCtx.iOpenTheDeck,
	)
	ctx.Step(
		`^I should see that the deck has the following "(shuffled|sorted)" cards:$`,
		deckCtx.iShouldSeeThatTheDeckHasTheFollowingCards,
	)
}

func (deckCtx *DeckContext) iOpenTheDeck(ctx context.Context) error {
	uri := "/decks/" + deckCtx.response.DeckID

	deckCtx.rawResponse = deckCtx.sendRequest(ctx, http.MethodGet, uri, nil)

	return nil
}

func (deckCtx *DeckContext) iShouldSeeThatTheDeckHasTheFollowingCards(
	ctx context.Context,
	deckOrder string,
	cardsTable *godog.Table,
) error {
	assert.Equal(godog.T(ctx), http.StatusOK, deckCtx.rawResponse.StatusCode)

	err := json.NewDecoder(deckCtx.rawResponse.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	cardsFromTable := deckCtx.parseCardsTable(ctx, cardsTable)

	if deckOrder == "sorted" {
		assert.Equal(godog.T(ctx), cardsFromTable, deckCtx.response.Cards)
		return nil
	}

	assert.ElementsMatch(godog.T(ctx), deckCtx.response.Cards, cardsFromTable)

	return nil
}
