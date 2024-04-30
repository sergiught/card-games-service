package scenario

import (
	"github.com/cucumber/godog"
)

// Initialize steps for the CreateFrenchDeck.feature scenarios.
func Initialize(ctx *Context) func(*godog.ScenarioContext) {
	return func(godogCtx *godog.ScenarioContext) {
		godogCtx.When(
			`^I create a "(standard|partial)"(?: and "(shuffled)")? deck of French cards$`,
			ctx.iCreateADeckOfFrenchCards,
		)
		godogCtx.When(
			`^I create a "(standard|partial)"(?: and "(shuffled)")? deck of French cards with the following cards in this order:$`,
			ctx.iCreateADeckOfFrenchCardsWithCards,
		)

		godogCtx.Then(`^I should receive "([^"]*)" cards$`, ctx.iShouldReceiveNCards)
		godogCtx.Then(`^I should have the following cards in this order:$`, ctx.iShouldHaveTheFollowingCardsInThisOrder)
		godogCtx.Then(`^I should have the cards in a shuffled order$`, ctx.iShouldHaveTheCardsInAShuffledOrder)
	}
}

func (ctx *Context) iCreateADeckOfFrenchCards(deckType, deckOrder string) error {
	return godog.ErrPending
}

func (ctx *Context) iCreateADeckOfFrenchCardsWithCards(deckType, deckOrder string, cardsTable *godog.Table) error {
	return godog.ErrPending
}

func (ctx *Context) iShouldReceiveNCards(cardsCount string) error {
	return godog.ErrPending
}

func (ctx *Context) iShouldHaveTheFollowingCardsInThisOrder(cardsTable *godog.Table) error {
	return godog.ErrPending
}

func (ctx *Context) iShouldHaveTheCardsInAShuffledOrder() error {
	return godog.ErrPending
}
