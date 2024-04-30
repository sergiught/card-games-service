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

// Initialize steps for the CreateFrenchDeck.feature scenarios.
func Initialize(deckCtx *DeckContext) func(*godog.ScenarioContext) {
	return func(ctx *godog.ScenarioContext) {
		ctx.When(
			`^I create a "(standard|custom)"(?: and "(shuffled)")? deck of French cards$`,
			deckCtx.iCreateADeckOfFrenchCards,
		)
		ctx.When(
			`^I create a "(standard|custom)"(?: and "(shuffled)")? deck of French cards with the following cards in this order:$`,
			deckCtx.iCreateADeckOfFrenchCardsWithCards,
		)

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
	return godog.ErrPending
}

func (deckCtx *DeckContext) iShouldReceiveNCards(ctx context.Context, cardsCount int) error {
	require.NotEmpty(godog.T(ctx), deckCtx.response)

	assert.Contains(godog.T(ctx), deckCtx.response, "deck_id")
	assert.Contains(godog.T(ctx), deckCtx.response, "shuffled")
	assert.Contains(godog.T(ctx), deckCtx.response, "remaining")
	assert.Equal(godog.T(ctx), cardsCount, deckCtx.response["remaining"])

	return nil
}

func (deckCtx *DeckContext) iShouldHaveTheFollowingCardsInThisOrder(ctx context.Context, cardsTable *godog.Table) error {
	return godog.ErrPending
}

func (deckCtx *DeckContext) iShouldHaveTheCardsInAShuffledOrder(ctx context.Context) error {
	return godog.ErrPending
}
