package scenario

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RegisterDrawCardsSteps registers steps for the DrawCardsFromFrenchDeck.feature scenarios.
func RegisterDrawCardsSteps(ctx *godog.ScenarioContext, deckCtx *DeckContext) {
	ctx.Step(
		`^I draw "(\d+)" cards from the deck$`,
		deckCtx.iDrawCardsFromTheDeck,
	)
	ctx.Step(
		`^I should see that I received "(\d+)" cards from the deck$`,
		deckCtx.iShouldSeeThatIReceivedCardsFromTheDeck,
	)
}

func (deckCtx *DeckContext) iDrawCardsFromTheDeck(ctx context.Context, cardsCount int) error {
	uri := "/decks/" + deckCtx.response.DeckID + "/draw"
	payload := []byte(fmt.Sprintf(`{"cards":%d}`, cardsCount))

	deckCtx.rawResponse = deckCtx.sendRequest(ctx, http.MethodPost, uri, payload)

	return nil
}

func (deckCtx *DeckContext) iShouldSeeThatIReceivedCardsFromTheDeck(ctx context.Context, cardsCount int) error {
	assert.Equal(godog.T(ctx), http.StatusOK, deckCtx.rawResponse.StatusCode)

	err := json.NewDecoder(deckCtx.rawResponse.Body).Decode(&deckCtx.response)
	require.NoError(godog.T(ctx), err)

	require.NotEmpty(godog.T(ctx), deckCtx.response)
	assert.Len(godog.T(ctx), deckCtx.response.Cards, cardsCount)

	return nil
}
